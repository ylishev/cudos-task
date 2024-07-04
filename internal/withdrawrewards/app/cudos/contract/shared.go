package contract

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

//go:generate mockery --name CudosWithdrawSender
type CudosWithdrawSender interface {
	Withdraw() (sdk.Coin, *sdk.TxResponse, error)
	Send(amount sdk.Coin) (sdk.Coin, *sdk.TxResponse, error)
}
