package cmd

import (
	"fmt"
	"os"

	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func initConfig(cfgFile *string, vp *viper.Viper, flagSet *pflag.FlagSet) func() {
	return func() {
		if *cfgFile != "" {
			// Use config file from the flag.
			vp.SetConfigFile(*cfgFile)
		} else {
			// Find home directory.
			home, err := os.UserHomeDir()
			cobra.CheckErr(err)

			// Search config in home directory with name ".withdraw"
			vp.AddConfigPath(home)
			vp.SetConfigName(".withdraw")
			vp.SetConfigType("yaml")
		}

		vp.AutomaticEnv()

		if err := vp.ReadInConfig(); err == nil {
			fmt.Println("using config file:", vp.ConfigFileUsed())
		}

		if vp.GetString(flags.FlagKeyringDir) == "" {
			pwd, err := os.Getwd()
			if err == nil {
				err = flagSet.Set("keyring-dir", pwd)
				cobra.CheckErr(err)
			}
		}
		// sync back the flags in case they are still with the default values
		for _, key := range vp.AllKeys() {
			if vp.InConfig(key) {
				flag := flagSet.Lookup(key)
				if flag != nil && !flag.Changed {
					err := flagSet.Set(key, vp.GetString(key))
					cobra.CheckErr(err)
				}
			}
		}
	}
}
