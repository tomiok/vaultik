package cmd

import (
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init {cipher key}",
	Short: "init the vaultik",
	Long:  `init the vaultik`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		newVaultik(args[0])
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
