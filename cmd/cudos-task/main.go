package main

import (
	"context"
	"log"

	"cudos-task/cmd/cudos-task/cmd"
	"cudos-task/contract"
	"cudos-task/internal/shutdown"

	"github.com/spf13/viper"
)

func main() {
	// prepare dependencies
	vp := viper.New()
	outChannel := make(chan string, 1)
	ctx, cancel := context.WithCancel(context.Background())

	// handle graceful shutdown of the application
	shutdownReady := shutdown.NewShutdown(cancel)

	// build and run Cobra commands
	rootCmd := cmd.InitRootCmd(ctx, vp, shutdownReady, outChannel)
	executedCMD, err := rootCmd.ExecuteC()
	if err != nil {
		log.Printf("error executing command: %v\nUse --help for details\n", err)
		cancel()
	}

	// report back the scheduled operations to the user
	for flg := executedCMD.Flags().Lookup(contract.ReportBackFlagName); flg != nil && flg.Changed; {
		select {
		// or wait for Ctrl+C
		case <-ctx.Done():
			return
		case msg := <-outChannel:
			log.Printf("[schedule] %v\n", msg)
		}
	}
}
