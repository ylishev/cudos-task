package cmd

import (
	"context"
	"fmt"
	"log"
	"time"

	"cudos-task/contract"

	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

//go:generate mockery --name CommandRunner
type CommandRunner interface {
	Run(cmd *cobra.Command, args []string)
}

// WithdrawRewardsCommandAttach provides a way to decouple outer layer (Cobra command) configuration
// from business layer by using CommandRunner interface
func WithdrawRewardsCommandAttach(ctx context.Context, runner CommandRunner, vp *viper.Viper) (*cobra.Command, error) {
	withdrawRewardsCmd := &cobra.Command{
		Use:   "withdraw-rewards",
		Short: "withdraw-rewards is a command for automatically withdraw all rewards and send to an address",
		Long: `
withdraw-rewards is a command for automatically collecting
of all rewards available from user's staked assets with
validators and sending them to an address.
	`,
		Run: func(cmd *cobra.Command, args []string) {
			_ = cmd.Flags().Set(contract.ReportBackFlagName, "true")
			runner.Run(cmd, args)
		},
		TraverseChildren: true,
	}

	// define sub-command specific flags
	var toAddress, from, memo string
	var interval time.Duration
	var reportBack bool
	withdrawRewardsCmd.PersistentFlags().StringVar(&toAddress, contract.ToAddressFlagName, "", fmt.Sprintf("address to send the rewards"))
	err := withdrawRewardsCmd.MarkPersistentFlagRequired(contract.ToAddressFlagName)
	if err != nil {
		return nil, err
	}

	withdrawRewardsCmd.PersistentFlags().StringVar(&from, flags.FlagFrom, "", fmt.Sprintf("keyring from"))
	err = withdrawRewardsCmd.MarkPersistentFlagRequired(flags.FlagFrom)
	if err != nil {
		return nil, err
	}

	withdrawRewardsCmd.PersistentFlags().StringVar(&memo, flags.FlagNote, contract.NoteDefault, fmt.Sprintf("memo"))

	withdrawRewardsCmd.PersistentFlags().DurationVar(&interval, contract.ScheduleIntervalFlagName, contract.ScheduleIntervalDefault,
		fmt.Sprintf("schedule interval duration"))

	withdrawRewardsCmd.PersistentFlags().BoolVar(&reportBack, contract.ReportBackFlagName, false, "")
	err = withdrawRewardsCmd.PersistentFlags().MarkHidden(contract.ReportBackFlagName)
	if err != nil {
		log.Fatalf("failed to mark flag as hidden: %v", err)
	}

	err = vp.BindPFlag(contract.ToAddressFlagName, withdrawRewardsCmd.PersistentFlags().Lookup(contract.ToAddressFlagName))
	if err != nil {
		return nil, err
	}
	err = vp.BindPFlag(flags.FlagFrom, withdrawRewardsCmd.PersistentFlags().Lookup(flags.FlagFrom))
	if err != nil {
		return nil, err
	}
	err = vp.BindPFlag(flags.FlagNote, withdrawRewardsCmd.PersistentFlags().Lookup(flags.FlagNote))
	if err != nil {
		return nil, err
	}
	err = vp.BindPFlag(contract.ScheduleIntervalFlagName, withdrawRewardsCmd.PersistentFlags().Lookup(contract.ScheduleIntervalFlagName))
	if err != nil {
		return nil, err
	}

	withdrawRewardsCmd.SetContext(ctx)

	return withdrawRewardsCmd, nil
}
