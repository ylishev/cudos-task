package app

import (
	"context"
	"errors"
	"fmt"
	"time"

	"cudos-task/contract"
	cudoscontract "cudos-task/internal/withdrawrewards/app/cudos/contract"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type CudosCommand struct {
	shutdown contract.ShutdownReady
	cudos    cudoscontract.CudosWithdrawSender
}

func NewCudosCommand(shutdown contract.ShutdownReady, cudos cudoscontract.CudosWithdrawSender) *CudosCommand {
	cmd := CudosCommand{
		shutdown: shutdown,
		cudos:    cudos,
	}

	return &cmd
}

func (cc CudosCommand) RunSchedule(ctx context.Context, out chan<- string, interval time.Duration) {
	go func() {
		ticker := time.NewTicker(time.Millisecond)
		defer ticker.Stop()
		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				func() {
					defer func() {
						cc.shutdown.SetReady(true)
						printWithChan(ctx, fmt.Sprintf("in progress; scheduling again in %s", interval), out)
						ticker.Reset(interval)
					}()
					// SetReady makes sure that both withdraw and send will be executed without being interrupted by Ctrl+C
					cc.shutdown.SetReady(false)
					ticker.Stop()
					if !printWithChan(ctx, "withdraw and send command", out) {
						return
					}

					withdrawRewardAmount, res, err := cc.cudos.Withdraw()
					if checkErrWithChan(ctx, err, out) {
						return
					}

					// don't try to broadcast send tx with zero amount
					if withdrawRewardAmount.IsZero() {
						return
					}

					// decouple domain-specific (Cosmos SDK) types to prevent leaking those to the outside layer
					// by generating a human-like reporting message
					if !printWithChan(ctx, cc.formatWithdrawRewards(res, withdrawRewardAmount), out) {
						return
					}

					// send the withdrawn reward (caveat: tx fees + gas >> withdrawn reward which makes no financial sense)
					sentAmount, res, err := cc.cudos.Send(withdrawRewardAmount)
					if checkErrWithChan(ctx, err, out) {
						return
					}
					if !printWithChan(ctx, cc.formatSend(res, sentAmount), out) {
						return
					}

					if sentAmount.String() != withdrawRewardAmount.String() {
						checkErrWithChan(ctx, fmt.Errorf("sent amont %s differs from withdraw reward %s",
							sentAmount, withdrawRewardAmount), out)
					}
				}()
			}
		}
	}()
}

func checkErrWithChan(ctx context.Context, err error, errChan chan<- string) bool {
	if err != nil {
		select {
		case <-ctx.Done():
		case errChan <- err.Error():
		}
		return true
	}
	return false
}

func printWithChan(ctx context.Context, msg string, outChan chan<- string) bool {
	select {
	case <-ctx.Done():
		return false
	case outChan <- msg:
		if errors.Is(ctx.Err(), context.Canceled) {
			return false
		}
	}
	return true
}

// formatWithdrawRewards calculate the total 'amount' of rewards
func (cc CudosCommand) formatWithdrawRewards(res *sdk.TxResponse, totalAmount sdk.Coin) string {
	return fmt.Sprintf("tx hash %s, gas used %d, withdraw rewards collected %v", res.TxHash, res.GasUsed, totalAmount)
}

// formatSend formats the sent 'amount'
func (cc CudosCommand) formatSend(res *sdk.TxResponse, totalAmount sdk.Coin) string {
	return fmt.Sprintf("tx hash %s, gas used %d, sent coins %v", res.TxHash, res.GasUsed, totalAmount)
}
