package cmd

import (
	"context"

	"cudos-task/contract"
	"cudos-task/internal/withdrawrewards/app"
	"cudos-task/internal/withdrawrewards/app/cudos"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type WithdrawRewardsCommand struct {
	ctx      context.Context
	vp       *viper.Viper
	shutdown contract.ShutdownReady
	out      chan string
}

// NewWithdrawRewardsCommand creates a new command and provides the dependencies
func NewWithdrawRewardsCommand(ctx context.Context, vp *viper.Viper, shutdown contract.ShutdownReady,
	outChannel chan string) WithdrawRewardsCommand {
	return WithdrawRewardsCommand{
		ctx:      ctx,
		vp:       vp,
		shutdown: shutdown,
		out:      outChannel,
	}
}

func (wc WithdrawRewardsCommand) Run(cmd *cobra.Command, _ []string) {
	cudosWithdrawSender := cudos.NewClient(cmd, wc.vp, wc.shutdown)
	cc := app.NewCudosCommand(wc.shutdown, cudosWithdrawSender)
	cc.RunSchedule(wc.ctx, wc.out, wc.vp.GetDuration(contract.ScheduleIntervalFlagName))
}
