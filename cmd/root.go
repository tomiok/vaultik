package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var (
	filename    string
	encodingKey string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "vaultik",
	Short:   "A brief description of your application",
	Long:    `this is VAULTIK, a CLI interface for add variables in a secure way`,
	Example: "",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// tests flags
	if b, _ := rootCmd.Flags().GetBool("help"); b {
		fmt.Print("toggled2")
	}

	if s, _ := rootCmd.Flags().GetString("example"); s == "test" {
		fmt.Print("example2")
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&encodingKey, "encodingKey", "k", "", "Encoding key")
}
