package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "vaultik",
	Short:   "A brief description of your application",
	Long:    `this is VAULTIK, a CLI interface for add variables in a secure way`,
	Example: "vaultik init this_is_a_secure_key",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {

}
