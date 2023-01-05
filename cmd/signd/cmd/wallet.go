package cmd

import (
	"_spike-signature-server/constant"
	"github.com/spf13/cobra"
)

func WalletCommands() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "wallet",
		Short: "Manage the wallet of the signature machine.",
	}

	cmd.AddCommand(
		AddWalletCommand(),
		ReSetWalletPassWordCommand())

	cmd.PersistentFlags().String(constant.FlagWalletDir, "", "The Hot wallet json file directory; can't omitted")
	cmd.PersistentFlags().String(constant.FlagWalletFileName, "", "The Hot wallet json file name; can't omitted")

	return cmd
}
