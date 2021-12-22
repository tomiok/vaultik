package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var readDecrypted bool

var allCmd = &cobra.Command{
	Use:   "all [-d true/false]",
	Short: "read all the variables saved",
	Long:  `Read all the variables saved, could be decrypted with the cipher key`,
	Run: func(cmd *cobra.Command, args []string) {
		vault := getVaultikData()

		cmd.Flags().BoolVarP(&readDecrypted, "decrypted", "d", false, "read the values in plain text")
		err := vault.printAll()

		if err != nil {
			fmt.Println(fmt.Sprintf("cannot read values: %s", err.Error()))
		}
	},
}

func init() {
	rootCmd.AddCommand(allCmd)
}
