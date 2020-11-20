package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "testvault",
	Short: "A brief description of your application",
	Long:  `this is VAULTIK`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		//TODO change this with tomatto-logger
		fmt.Println(err)
		os.Exit(1)
	}
	if b, _ := rootCmd.Flags().GetBool("help"); b {
		fmt.Print("toggled2")
	}

	if s, _ := rootCmd.Flags().GetString("example"); s == "test" {
		fmt.Print("example2")
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	//rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.vaultik.yaml)")

	//
	rootCmd.Flags().BoolP("help", "h", false, "Call the help for using Vaultik")
	//
	rootCmd.Flags().StringP("example", "e", "", "check out a simple example")
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

		// Search config in home directory with name ".vaultik" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".vaultik")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
