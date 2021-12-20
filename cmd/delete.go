package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var delCmd = &cobra.Command{
	Use:   "del {key}",
	Short: "delete a file given a key in plain text",
	Long:  `delete a file given a key in plain text`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("deleting...")
	},
}

func init() {
	rootCmd.AddCommand(delCmd)
}
