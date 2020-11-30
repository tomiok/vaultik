package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get the value, the cipher encodingKey is needed",
	Long:  `Get the value, the cipher encodingKey is needed (Long)`,
	Run: func(cmd *cobra.Command, args []string) {
		vault := newVaultik(encodingKey, filename)

		if len(args) == 0 {
			fmt.Println("Please provide a key")
			return
		}

		key := args[0]

		res, err := vault.getValue(key)

		if err != nil {
			fmt.Println(fmt.Sprintf("error: %s, please try again", err.Error()))
			return
		}

		fmt.Println("result: " + res)
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
}
