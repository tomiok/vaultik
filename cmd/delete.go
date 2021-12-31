package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var delCmd = &cobra.Command{
	Use:   "del {key}",
	Short: "delete a file given a key in plain text",
	Long:  `delete a file given a key in plain text`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		key := args[0]
		err := deleteFile(key)

		if err != nil {
			fmt.Println(fmt.Sprintf("cannot delete file %s, due to error %s", key, err.Error()))
			return
		}

		fmt.Println(fmt.Sprintf("file %s deleted", key))
	},
}

func deleteFile(key string) error {
	p, err := getSecretPath(key)

	if err != nil {
		return nil
	}

	return os.Remove(p)
}

func init() {
	rootCmd.AddCommand(delCmd)
}
