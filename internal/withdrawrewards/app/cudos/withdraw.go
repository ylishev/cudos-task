package cudos

import (
	"fmt"

	"cudos-task/contract"

	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	distribtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
)

func (cc CudosClient) Withdraw() (sdk.Coin, sdk.TxResponse, error) {
	res := sdk.TxResponse{}
	coin := sdk.Coin{}
	delAddr := cc.clientCtx.GetFromAddress()
	queryClient := distribtypes.NewQueryClient(cc.clientCtx)
	delValsRes, err := queryClient.DelegatorValidators(cc.cobraCmd.Context(), &distribtypes.QueryDelegatorValidatorsRequest{DelegatorAddress: delAddr.String()})
	if err != nil {
		return coin, res, fmt.Errorf("failed to obtain the staking validators: %v", err)
	}

	validators := delValsRes.Validators
	// build multi-message transaction
	msgs := make([]sdk.Msg, 0, len(validators))
	for _, valAddr := range validators {
		val, err := sdk.ValAddressFromBech32(valAddr)
		if err != nil {
			return coin, res, fmt.Errorf("failed to check the validator address: %v", err)
		}

		msg := distribtypes.NewMsgWithdrawDelegatorReward(delAddr, val)
		if err := msg.ValidateBasic(); err != nil {
			return coin, res, fmt.Errorf("failed to check the withdraw message: %v", err)
		}
		msgs = append(msgs, msg)
	}

	// clean the buffer for future use
	_ = cc.outWriter.Flush()
	cc.outBuffer.Reset()

	err = tx.GenerateOrBroadcastTxCLI(cc.clientCtx, cc.cobraCmd.Flags(), msgs...)
	if err != nil {
		return coin, res, fmt.Errorf("failed to broadcast withdraw rewards tx: %v", err)
	}

	_ = cc.outWriter.Flush()
	err = cc.clientCtx.Codec.UnmarshalJSON(cc.outBuffer.Bytes(), &res)

	// clean the buffer for future use
	cc.outBuffer.Reset()

	if err != nil {
		return coin, res, fmt.Errorf("failed to unmarshal the result of send tx: %v", err)
	}
	if res.Code != 0 {
		return coin, res, fmt.Errorf("widthdraw rewards tx faild: %s", res.RawLog)
	}

	return cc.CalculateWithdrawRewards(res), res, nil
}

// CalculateWithdrawRewards calculate the total 'amount' of rewards
func (cc CudosClient) CalculateWithdrawRewards(res sdk.TxResponse) sdk.Coin {
	total := sdk.NewInt64Coin(contract.Denom, 0)
	if res.Code != 0 {
		return total
	}

	var received []sdk.Coin

	for _, e := range res.Events {
		if e.Type == distribtypes.EventTypeWithdrawRewards {
			for _, a := range e.Attributes {
				if string(a.Key) == sdk.AttributeKeyAmount {
					coin, err := sdk.ParseCoinNormalized(string(a.Value))
					if err == nil && coin.Denom == contract.Denom {
						received = append(received, coin)
					}
				}
			}
		}
	}

	for _, c := range received {
		total = total.AddAmount(c.Amount)
	}
	return total
}

// FormatWithdrawRewards calculate the total 'amount' of rewards
func (cc CudosClient) FormatWithdrawRewards(res sdk.TxResponse, totalAmount sdk.Coin) string {
	return fmt.Sprintf("tx hash %s, gas used %d, withdraw rewards collected %v", res.TxHash, res.GasUsed, totalAmount)
}
