package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var allCmd = &cobra.Command{
	Use:   "all [-d true/false]",
	Short: "read all the variables saved",
	Long:  `Read all the variables saved, could be decrypted with the `,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		vault := getVaultikData()

		err := vault.readAll()

		if err != nil {
			fmt.Println(fmt.Sprintf("cannot read values: %s", err.Error()))
		}
	},
}

func init() {
	rootCmd.AddCommand(allCmd)
}
