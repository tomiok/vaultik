package cmd

import (
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "vaultik",
	Short:   "A brief description of your application",
	Long:    `this is VAULTIK, a CLI interface for add variables in a secure way`,
	Example: "vaultik init this_is_a_secure_key",
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {}
