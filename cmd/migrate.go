package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// method is for
var method string

var migrateCmd = &cobra.Command{
	Use:     "migrate {user} {path to id_rsa} {destination}",
	Short:   "migrate should send in scp protocol all the secured and hashed pair of key-values",
	Long:    ``,
	Example: "migrate tomas p45w0rd1 example.com:22",
	Args:    cobra.MinimumNArgs(3),
	Run: func(cmd *cobra.Command, args []string) {

		username := args[0]
		pass := args[1]
		host := args[3]

		sshClientConfig := withPassword(username, pass, nil)

		client := NewClient(host, &sshClientConfig)

		err := client.Connect()

		if err != nil {
			fmt.Println(err.Error())
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(migrateCmd)
}
