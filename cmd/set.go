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
		for _, s := range args {
			fmt.Println(s)
		}
	},
}

func init() {
	rootCmd.AddCommand(setCmd)
}
