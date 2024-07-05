package startup

import (
	"os"

	"cudos-task/contract"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// init function to shut down std err writes, done by cosmos sdk, without using logger functions
// see go/pkg/mod/github.com/cosmos/cosmos-sdk@v0.45.16/client/tx/tx.go:101
// nolint:gochecknoinits // workaround, see previous line
func init() {
	devNull, err := os.OpenFile(os.DevNull, os.O_WRONLY, 0755)
	if err != nil {
		panic(err)
	}
	os.Stderr = devNull
}

const (
	AccountPubKeyPrefix    = contract.AccountAddressPrefix + "pub"
	ValidatorAddressPrefix = contract.AccountAddressPrefix + "valoper"
	ValidatorPubKeyPrefix  = contract.AccountAddressPrefix + "valoperpub"
	ConsNodeAddressPrefix  = contract.AccountAddressPrefix + "valcons"
	ConsNodePubKeyPrefix   = contract.AccountAddressPrefix + "valconspub"
)

// SetConfig applies cudos specific customizations
func SetConfig() {
	config := sdk.GetConfig()
	config.SetBech32PrefixForAccount(contract.AccountAddressPrefix, AccountPubKeyPrefix)
	config.SetBech32PrefixForValidator(ValidatorAddressPrefix, ValidatorPubKeyPrefix)
	config.SetBech32PrefixForConsensusNode(ConsNodeAddressPrefix, ConsNodePubKeyPrefix)
	config.Seal()
}
