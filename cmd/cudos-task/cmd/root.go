package cmd

import (
	"context"
	"log"

	"cudos-task/contract"
	"cudos-task/internal/withdrawrewards/cmd"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func InitRootCmd(ctx context.Context, vp *viper.Viper, shutdown contract.ShutdownReady, outChannel chan string) *cobra.Command {
	var rootCmd = &cobra.Command{
		Use:              "cudos-task",
		Short:            "cudos-task is a CLI tool that provides command for automatically withdraw all rewards and send them to an address",
		RunE:             client.ValidateCmd,
		TraverseChildren: true,
	}

	var cfgFile, node, chainId, keyringBackend, keyringDir, gasPrices, gas string
	var gasAdjustment float64
	var yes bool

	// handle bootstrap the loading of the configuration file
	cobra.OnInitialize(initConfig(&cfgFile, vp, rootCmd.PersistentFlags()))

	// define global flags
	rootCmd.PersistentFlags().StringVar(&cfgFile, contract.ConfigFlagName, "", "config file (default is $HOME/.withdraw.yaml)")
	rootCmd.PersistentFlags().StringVar(&node, flags.FlagNode, contract.NodeAddressDefault, "cudos node")
	rootCmd.PersistentFlags().StringVar(&chainId, flags.FlagChainID, contract.ChainIDDefault, "cudos chain id")
	rootCmd.PersistentFlags().StringVar(&keyringBackend, flags.FlagKeyringBackend, contract.KeyringBackendDefault, "keyring backend")
	rootCmd.PersistentFlags().StringVar(&keyringDir, flags.FlagKeyringDir, "", "keyring dir (default is $PWD/)")
	rootCmd.PersistentFlags().StringVar(&gasPrices, flags.FlagGasPrices, contract.GasPricesDefault, "gas prices")
	rootCmd.PersistentFlags().Float64Var(&gasAdjustment, flags.FlagGasAdjustment, contract.GasAdjDefault, "gas adjustment")
	rootCmd.PersistentFlags().StringVar(&gas, flags.FlagGas, contract.GasDefault, "gas")
	rootCmd.PersistentFlags().BoolVar(&yes, flags.FlagSkipConfirmation, contract.SkipConfirmDefault, "")

	err := rootCmd.PersistentFlags().MarkHidden(flags.FlagSkipConfirmation)
	if err != nil {
		log.Fatalf("failed to mark flag as hidden: %v", err)
	}

	// bind them to the viper
	err = vp.BindPFlag(flags.FlagNode, rootCmd.PersistentFlags().Lookup(flags.FlagNode))
	if err != nil {
		log.Fatalf("failed to bind flag: %v", err)
	}
	err = vp.BindPFlag(flags.FlagChainID, rootCmd.PersistentFlags().Lookup(flags.FlagChainID))
	if err != nil {
		log.Fatalf("failed to bind flag: %v", err)
	}
	err = vp.BindPFlag(flags.FlagKeyringBackend, rootCmd.PersistentFlags().Lookup(flags.FlagKeyringBackend))
	if err != nil {
		log.Fatalf("failed to bind flag: %v", err)
	}
	err = vp.BindPFlag(flags.FlagKeyringDir, rootCmd.PersistentFlags().Lookup(flags.FlagKeyringDir))
	if err != nil {
		log.Fatalf("failed to bind flag: %v", err)
	}
	err = vp.BindPFlag(flags.FlagGasPrices, rootCmd.PersistentFlags().Lookup(flags.FlagGasPrices))
	if err != nil {
		log.Fatalf("failed to bind flag: %v", err)
	}
	err = vp.BindPFlag(flags.FlagGasAdjustment, rootCmd.PersistentFlags().Lookup(flags.FlagGasAdjustment))
	if err != nil {
		log.Fatalf("failed to bind flag: %v", err)
	}
	err = vp.BindPFlag(flags.FlagGas, rootCmd.PersistentFlags().Lookup(flags.FlagGas))
	if err != nil {
		log.Fatalf("failed to bind flag: %v", err)
	}
	err = vp.BindPFlag(flags.FlagSkipConfirmation, rootCmd.PersistentFlags().Lookup(flags.FlagSkipConfirmation))
	if err != nil {
		log.Fatalf("failed to bind flag: %v", err)
	}

	// define withdraw-rewards Cobra command and bridge to the business layer
	wrCmd, err := WithdrawRewardsCommandAttach(ctx, cmd.NewWithdrawRewardsCommand(ctx, vp, shutdown, outChannel), vp)
	if err != nil {
		log.Fatalf("failed to attach withdraw send command: %v", err)
	}
	rootCmd.SetContext(ctx)
	rootCmd.AddCommand(wrCmd)
	return rootCmd
}
