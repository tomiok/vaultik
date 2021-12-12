package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var (
	filename    string
	encodingKey string
	cfgFile     string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:                        "vaultik",
	Short:                      "A brief description of your application",
	Long:                       `this is VAULTIK`,
	Example:                    "",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// tests flags
	if b, _ := rootCmd.Flags().GetBool("help"); b {
		fmt.Print("toggled2")
	}

	if s, _ := rootCmd.Flags().GetString("example"); s == "test" {
		fmt.Print("example2")
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVarP(&filename, "file", "f", "", "File vault")
	rootCmd.PersistentFlags().StringVarP(&encodingKey, "encodingKey", "k", "", "Encoding encodingKey")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		viper.AddConfigPath(home)
		viper.SetConfigName(".vaultik")
	}

	// read in environment variables that match
	viper.AutomaticEnv()

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
