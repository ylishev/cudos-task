package cudos

import (
	"fmt"

	"cudos-task/contract"

	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
)

func (cc CudosClient) Send(amount sdk.Coin) (sdk.Coin, sdk.TxResponse, error) {
	res := sdk.TxResponse{}
	coin := sdk.Coin{}

	toAddress, err := sdk.AccAddressFromBech32(cc.vp.GetString(contract.ToAddressFlagName))
	if err != nil {
		return coin, res, fmt.Errorf("failed to obtain the to-address: %v", err)
	}

	msg := &banktypes.MsgSend{
		FromAddress: cc.clientCtx.GetFromAddress().String(),
		ToAddress:   toAddress.String(),
		Amount:      sdk.NewCoins(amount),
	}
	if err := msg.ValidateBasic(); err != nil {
		return coin, res, fmt.Errorf("failed to parse the amount to send: %v", err)
	}

	// clean the buffer for future use
	_ = cc.outWriter.Flush()
	cc.outBuffer.Reset()
	err = tx.GenerateOrBroadcastTxCLI(cc.clientCtx, cc.cobraCmd.Flags(), msg)

	if err != nil {
		return coin, res, fmt.Errorf("failed to broadcast send tx: %v", err)
	}

	_ = cc.outWriter.Flush()
	err = cc.clientCtx.Codec.UnmarshalJSON(cc.outBuffer.Bytes(), &res)
	if err != nil {
		return coin, res, fmt.Errorf("failed to unmarshal the result of send tx: %v", err)
	}

	// clean the buffer for future use
	cc.outBuffer.Reset()
	if res.Code != 0 {
		return coin, res, fmt.Errorf("send tx faild: %s", res.RawLog)
	}

	return cc.CalculateSentAmount(res, toAddress), res, nil
}

// FormatSend formats the sent 'amount'
func (cc CudosClient) FormatSend(res sdk.TxResponse, totalAmount sdk.Coin) string {
	return fmt.Sprintf("tx hash %s, gas used %d, sent coins %v", res.TxHash, res.GasUsed, totalAmount)
}

// CalculateSentAmount calculate the total 'amount' of sent coins
func (cc CudosClient) CalculateSentAmount(res sdk.TxResponse, to sdk.AccAddress) sdk.Coin {
	total := sdk.NewInt64Coin(contract.Denom, 0)
	if res.Code != 0 {
		return total
	}

	var received []sdk.Coin
NextEvent:
	for _, e := range res.Events {
		if e.Type == banktypes.EventTypeCoinReceived {
			for _, a := range e.Attributes {
				if string(a.Key) == banktypes.AttributeKeyReceiver && string(a.Value) != to.String() {
					continue NextEvent
				}
			}
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
