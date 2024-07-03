package cudos

import (
	"bufio"
	"bytes"
	"log"
	"os"

	"cudos-task/contract"
	"cudos-task/internal/withdrawrewards/app/cudos/startup"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/simapp"
	"github.com/cosmos/cosmos-sdk/std"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtx "github.com/cosmos/cosmos-sdk/x/auth/tx"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type CudosClient struct {
	cobraCmd  *cobra.Command
	shutdown  contract.ShutdownReady
	clientCtx client.Context
	outBuffer *bytes.Buffer
	outWriter *bufio.Writer
	vp        *viper.Viper
}

func NewCudosClient(cc *cobra.Command, vp *viper.Viper, shutdown contract.ShutdownReady) *CudosClient {
	cmd := CudosClient{
		cobraCmd: cc,
		shutdown: shutdown,
		vp:       vp,
	}

	sdk.DefaultPowerReduction = sdk.NewIntFromUint64(contract.CudosPowerReduction)
	startup.SetConfig()

	aminoCodec := codec.NewLegacyAmino()
	interfaceRegistry := codectypes.NewInterfaceRegistry()
	marshaler := codec.NewProtoCodec(interfaceRegistry)
	txConfig := authtx.NewTxConfig(marshaler, authtx.DefaultSignModes)

	std.RegisterLegacyAminoCodec(aminoCodec)
	std.RegisterInterfaces(interfaceRegistry)
	simapp.ModuleBasics.RegisterLegacyAminoCodec(aminoCodec)
	simapp.ModuleBasics.RegisterInterfaces(interfaceRegistry)

	var outBuffer bytes.Buffer
	cmd.outBuffer = &outBuffer
	outWriter := bufio.NewWriter(&outBuffer)
	cmd.outWriter = outWriter

	clientCtx, err := client.GetClientTxContext(cc)
	if err != nil {
		log.Fatalf("failed to obtain client context: %v", err)
	}
	clientCtx.Viper = vp

	clientCtx = clientCtx.
		WithCodec(marshaler).
		WithJSONCodec(marshaler).
		WithInterfaceRegistry(interfaceRegistry).
		WithTxConfig(txConfig).
		WithLegacyAmino(aminoCodec).
		WithInput(os.Stdin).
		WithOutput(outWriter).
		WithBroadcastMode(flags.BroadcastBlock).
		WithAccountRetriever(types.AccountRetriever{})

	cmd.clientCtx = clientCtx

	return &cmd
}

func (cc CudosClient) Context() client.Context {
	return cc.clientCtx
}
