package cmd

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/ironman-project/ironman/pkg/ironman"
	"github.com/mitchellh/go-homedir"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

//Commands global variables
var cfgFile string
var ironmanHome string
var verbose bool

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
	err := ironman.InitIronmanHome(ironmanHome)
	if err != nil {
		fmt.Printf("failed to create ironman home directory %s %s", ironmanHome, err)
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
		fmt.Printf("failed to get the current user home dir %s ", err)
		os.Exit(-1)
	}
	defaultIronmanHomeDir := filepath.Join(defaultHomeDir, ".ironman")
	rootCmd.PersistentFlags().StringVar(&ironmanHome, "ironman-home", defaultIronmanHomeDir, "ironman home directory")
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", true, "verbose output e.g --verbose false")
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

var thman *ironman.Ironman

func iironman() *ironman.Ironman {
	if thman == nil {
		thman = ironman.New(ironmanHome, ironman.SetOutput(ironmanOutput()))
	}

	return thman
}

var logger *log.Logger

func ilogger() *log.Logger {
	if logger == nil {
		logger = log.New(ironmanOutput(), "", 0)
	}
	return logger
}

func ironmanOutput() io.Writer {
	var output io.Writer = os.Stdout
	if !verbose {
		output = ioutil.Discard
	}
	return output
}
