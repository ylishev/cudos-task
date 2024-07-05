package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func initConfig(cfgFile *string, vp *viper.Viper, rootCmd *cobra.Command) func() {
	return func() {
		flagSet := rootCmd.PersistentFlags()
		if *cfgFile != "" {
			// use config file from the flag
			vp.SetConfigFile(*cfgFile)
		} else {
			home, err := os.UserHomeDir()
			cobra.CheckErr(err)

			// search config in home directory with name ".withdraw"
			vp.AddConfigPath(home)
			vp.SetConfigName(".withdraw")
			vp.SetConfigType("yaml")
		}

		vp.AutomaticEnv()
		vp.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

		if err := vp.ReadInConfig(); err == nil {
			_, _ = fmt.Fprintf(rootCmd.OutOrStdout(), "using config file:%s\n", vp.ConfigFileUsed())
		}

		// figure out the default keyring directory
		if vp.GetString(flags.FlagKeyringDir) == "" {
			pwd, err := os.Getwd()
			if err == nil {
				err = flagSet.Set(flags.FlagKeyringDir, pwd)
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
