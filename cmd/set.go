package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var setCmd = &cobra.Command{
	Use:   "set",
	Short: " ",
	Long: ` `,
	Run: func(cmd *cobra.Command, args []string) {
		for _, s := range args {
			fmt.Println(s)
		}
	},
}

func init() {
	rootCmd.AddCommand(setCmd)
	setCmd.Flags().StringP("key", "k", "", "The key to encrypt the secret")
}
