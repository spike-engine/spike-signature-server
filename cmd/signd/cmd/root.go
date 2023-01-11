package cmd

import (
	"github.com/spf13/cobra"
)

var (
	rootCmd *cobra.Command
)

func RootCommands() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "signd",
		Short: "Spike Signature Machine",
	}
	return cmd
}

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd = RootCommands()

	rootCmd.AddCommand(
		WalletCommands(),
		StartCommand())

}
