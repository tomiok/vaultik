package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set the value and save it in the file system, the encodingKey is required",
	Long:  `Set the value and save it in the file system, the encodingKey is required (Long)`,
	Run: func(cmd *cobra.Command, args []string) {

		vault := newVaultik(encodingKey, filename)

		if len(args) < 2 {
			fmt.Println("please provide the key and de actual API key")
			return
		}

		if err := vault.setValue(args[0], args[1]); err != nil {
			fmt.Printf("error: %s, please try again", err.Error())
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(setCmd)
}
