package main

import (
	"bytes"
	"context"
	"io"
	"testing"

	"cudos-task/cmd/cudos-task/cmd"
	"cudos-task/cmd/cudos-task/cmd/mocks"
	"cudos-task/contract"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestCudosTaskCommand(t *testing.T) {
	tests := []struct {
		name  string
		args  []string
		want  string
		fails bool
	}{
		{
			name:  "runs withdraw command and returns help",
			args:  []string{contract.WithdrawRewardsCmdName, "--help"},
			want:  "withdraw-rewards is a command for automatically collecting",
			fails: false,
		},
		{
			name:  "runs withdraw command without params and returns missing flags error",
			args:  []string{contract.WithdrawRewardsCmdName},
			want:  `required flag(s) "from", "to-address" not set`,
			fails: true,
		},
		{
			name:  "runs withdraw command successfully",
			args:  []string{contract.WithdrawRewardsCmdName, "--from", "test", "--to-address", "cudos1fefgllqh9qpjn3laz3lmhl68kvlfhwzmq6clfq"},
			want:  ``,
			fails: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vp := viper.New()
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			// build and run Cobra commands
			withdrawAwardsCmd := new(mocks.CommandRunner)
			withdrawAwardsCmd.On("Run", mock.AnythingOfType("*cobra.Command"), mock.AnythingOfType("[]string"))

			rootCmd := cmd.InitRootCmd(ctx, vp, withdrawAwardsCmd)
			tBuf := bytes.NewBufferString("")
			rootCmd.SetOut(tBuf)
			rootCmd.SetArgs(tt.args)
			executedCMD, err := rootCmd.ExecuteC()
			if err != nil {
				cancel()
				if !tt.fails {
					assert.FailNow(t, "unexpected error executing command")
				}
				assert.Contains(t, err.Error(), tt.want)
				return
			}
			require.NotNil(t, executedCMD, "cobra command should not be nil")

			out, err := io.ReadAll(tBuf)
			require.Nil(t, err)
			assert.Contains(t, string(out), tt.want)
		})
	}
}
