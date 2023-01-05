package cmd

import (
	"_spike-signature-server/config"
	"_spike-signature-server/constant"
	"_spike-signature-server/initialize"
	"github.com/spf13/cobra"
)

func StartCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "start",
		Short: "Start Spike Signature Server",
		Long:  `According to the configuration item of config.toml, start this project`,
		Run: func(cmd *cobra.Command, args []string) {
			constant.Viper = config.InitViper()
			config.InitWallet()
			initialize.RunServer()
		},
	}

	return cmd
}
