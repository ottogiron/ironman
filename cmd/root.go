package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/mitchellh/go-homedir"

	"github.com/ironman-project/ironman/template/repository"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string
var ironmanHome string

// RootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:           "ironman",
	Short:         "`Template manager and engine",
	Long:          `Template manager and engine`,
	SilenceUsage:  true,
	SilenceErrors: true,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		initializeIronmanHome()
	},
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, "There was an error", err)
		os.Exit(-1)
	}
}

func initializeIronmanHome() {
	err := repository.InitIronmanHome(ironmanHome)
	if err != nil {
		fmt.Printf("Failed to create ironman home directory %s %s", ironmanHome, err)
		os.Exit(-1)
	}

}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports Persistent Flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.ironman.yaml)")
	defaultHomeDir, err := homedir.Dir()
	if err != nil {
		fmt.Printf("Failed to get the current user home dir %s ", err)
		os.Exit(-1)
	}
	defaultIronmanHomeDir := filepath.Join(defaultHomeDir, ".ironman")
	rootCmd.PersistentFlags().StringVar(&ironmanHome, "ironman-home", defaultIronmanHomeDir, "ironman home directory")

}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" { // enable ability to specify config file via flag
		viper.SetConfigFile(cfgFile)
	}

	viper.SetConfigName(".ironman") // name of config file (without extension)
	viper.AddConfigPath("$HOME")    // adding home directory as first search path
	viper.AutomaticEnv()            // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
