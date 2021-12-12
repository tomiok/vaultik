package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get [-d | --decrypt] {key} | will return the value of the given key. Use -d for read the value in plain text",
	Short: "",
	Long:  `Get the value, the cipher encodingKey is needed (Long)`,
	Run: func(cmd *cobra.Command, args []string) {
		vault := getVaultikData()
		if len(args) == 0 {
			fmt.Println("please provide a key")
			return
		}

		key := args[0]

		res, err := vault.getValue(key)

		if err != nil {
			fmt.Println(fmt.Sprintf("error: %v, please try again", err))
			return
		}

		fmt.Println("result: " + res)
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
}
