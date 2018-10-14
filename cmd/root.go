package cmd

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/ironman-project/ironman/pkg/ironman"
	homedir "github.com/mitchellh/go-homedir"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	buildVersion string
	buildCommit  string
	buildDate    string
)

//Commands global variables
var cfgFile string
var ironmanHome string
var verbose bool

type commandFactory func(client *ironman.Ironman, out io.Writer) *cobra.Command

func newRootCmd() *cobra.Command {

	// RootCmd represents the base command when called without any subcommands
	var rootCmd = &cobra.Command{
		Use:           "ironman",
		Short:         "`Template manager and engine",
		Long:          `Template manager and engine`,
		SilenceUsage:  true,
		SilenceErrors: true,
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			ironman.EnsureIronmanHome(ironmanHome)
		},
	}

	//Add the command factory here
	commandFactories := []commandFactory{
		newGenerateCommand,
		newInstallCommand,
		newLinkCmd,
		newListCmd,
		newUninstallCmd,
		newUnlinkCmd,
		newUpdateCmd,
		newCreateCmd,
		newDescribeCmd,
	}

	//add all commands
	for _, cmdFactory := range commandFactories {
		rootCmd.AddCommand(cmdFactory(nil, nil))
	}

	//Version command special case
	rootCmd.AddCommand(newVersionCmd(buildVersion, buildCommit, buildDate, os.Stdout))

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
	return rootCmd
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	rootCmd := newRootCmd()
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, "There was an error", err)
		os.Exit(-1)
	}
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

func ensureIronmanClientAndOutput(client *ironman.Ironman, out io.Writer) (*ironman.Ironman, io.Writer) {
	return ensureIronmanClient(client), ensureIronmanOutput(out)
}

func ensureIronmanClient(client *ironman.Ironman) *ironman.Ironman {
	if client == nil {
		return ironman.New(ironmanHome)
	}
	return client
}

func ensureIronmanOutput(out io.Writer) io.Writer {
	if out == nil {
		return ironmanOutput()
	}
	return out
}

func iironman(home string, out io.Writer) *ironman.Ironman {
	thman := ironman.New(home, ironman.SetOutput(out))
	return thman
}

func ironmanOutput() io.Writer {
	var output io.Writer = os.Stdout
	if !verbose {
		output = ioutil.Discard
	}
	return output
}
