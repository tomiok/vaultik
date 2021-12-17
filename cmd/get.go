package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var decrypted bool

var getCmd = &cobra.Command{
	Use:   "get [-d | --decrypt (true|false)] {key} | will return the value of the given key. Use -d for read the value in plain text",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		vault := getVaultikData()
		if len(args) == 0 {
			fmt.Println("please provide a key")
			return
		}

		key := args[0]

		res, err := vault.getValue(key, decrypted)

		if err != nil {
			fmt.Println(fmt.Sprintf("error: %v, please try again", err))
			return
		}
		fmt.Println(res)
	},
}

func init() {
	getCmd.Flags().BoolVarP(&decrypted, "decrypted", "d", false, "read the  value in plain text")
	rootCmd.AddCommand(getCmd)
}
