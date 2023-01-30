package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "tri",
	Short: "tri is a todo application",
	Long:  `Tri will help you get more done in less time. It's a designed to be as simple as possible to help you accomplish your goals.`,
}

var dataFile string
var cfgFile string

func init() {
	home, err := homedir.Dir()

	if err != nil {
		log.Println("Unable to detect home directory. Please set data using --datafile")
	}

	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.PersistentFlags().StringVar(&dataFile, "datafile", home+string(os.PathSeparator)+".tridos.json", "data file to stroe todods")
}

// Read in config file and ENV variables if set.
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

		// Search config in home directory with name ".tri" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".tri")
	}

	viper.SetEnvPrefix("tri")
	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

func Execute() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.tri.yaml)")
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
