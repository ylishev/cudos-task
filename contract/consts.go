package contract

import (
	"time"

	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
)

const (
	AccountAddressPrefix    = "cudos"
	Denom                   = "acudos"
	CudosPowerReduction     = 1000000000000000000
	NodeAddressDefault      = "https://rpc.cudos.org:443"
	ChainIDDefault          = "cudos-1"
	GasAdjDefault           = 1.3
	GasDefault              = flags.GasFlagAuto
	GasPricesDefault        = "5000000000000acudos"
	NoteDefault             = "Tx via withdraw-rewards app"
	KeyringBackendDefault   = keyring.BackendTest
	SkipConfirmDefault      = true
	ScheduleIntervalDefault = 5 * time.Minute
	ShutdownMaxWaitTime     = time.Minute
	WithdrawRewardsCmdName  = "withdraw-rewards"
)

const (
	ToAddressFlagName        = "to-address"
	ConfigFlagName           = "config"
	ScheduleIntervalFlagName = "interval"
	ReportBackFlagName       = "report-back"
)
