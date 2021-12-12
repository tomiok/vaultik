package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set the value and save it in the file system, the encodingKey and filename are required",
	Long:  `Set the value and save it in the file system, the encodingKey and filename are required (Long)`,
	Run: func(cmd *cobra.Command, args []string) {
		vault := getVaultikData()

		if len(args) < 2 {
			fmt.Println("please provide the key and de actual API key")
			return
		}

		if err := vault.setValue(args[0], args[1]); err != nil {
			fmt.Println(fmt.Printf("error: %v, please try again", err))
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(setCmd)
}
