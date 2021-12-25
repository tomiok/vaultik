package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// method is for
var method string

var migrateCmd = &cobra.Command{
	Use:     "migrate {user} {path to id_rsa} {destination} [-m=privateKey]",
	Short:   "migrate should send in scp protocol all the secured and hashed pair of key-values",
	Long:    ``,
	Example: "migrate tomas p45w0rd1 example.com:22",
	Args:    cobra.MinimumNArgs(3),
	Run: func(cmd *cobra.Command, args []string) {

		username := args[0]
		access := args[1]
		host := args[3]

		if method == "" {
			sshClientConfig := withPassword(username, access, nil)

			client := NewClient(host, &sshClientConfig)

			err := client.Connect()

			if err != nil {
				fmt.Println(err.Error())
				return
			}
		} else {
			sshClientConfig, err := withPrivateKey(username, access, nil)
			if err != nil {
				fmt.Println(fmt.Sprintf("cannot access to remote address %s", err.Error()))
				return
			}

			client := NewClient(host, &sshClientConfig)

			err = client.Connect()

			if err != nil {
				fmt.Println(err.Error())
				return
			}
		}
	},
}

func init() {
	migrateCmd.Flags().StringVarP(&method, "method", "m", "", "method is the auth method for" +
		"ths SSH connection. use privateKey and provide the PATH to your private key, or leave it blank to use " +
		"standard user + password connection")
	rootCmd.AddCommand(migrateCmd)
}
