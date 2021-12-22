package cmd

import (
	"github.com/spf13/cobra"
)

var migrateCmd = &cobra.Command{
	Use:   "migrate {user} {path to id_rsa} {destination}",
	Short: "migrate should send in scp protocol all the secured and hashed pair of key-values",
	Long:  ``,
	Example: "migrate tomas p45w0rd1 example.com:22",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {
	rootCmd.AddCommand(migrateCmd)
}
