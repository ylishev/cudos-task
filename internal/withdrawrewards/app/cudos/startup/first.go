package startup

import (
	"os"

	"cudos-task/contract"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// init function to shut down std err writes, done by cosmos sdk, without using logger functions
func init() {
	devNull, err := os.OpenFile(os.DevNull, os.O_WRONLY, 0755)
	if err != nil {
		panic(err)
	}
	os.Stderr = devNull
}

var (
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
