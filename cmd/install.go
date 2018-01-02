package cmd

import (
	"fmt"
	"net/url"

	"github.com/ironman-project/ironman/template/repository/git"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

// installCmd represents the install command
var installCmd = &cobra.Command{
	Use: "install <url>",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("url arg is required")
		}
		templateURL := args[0]
		if _, err := url.Parse(args[0]); err != nil {
			return errors.Wrapf(err, "Invalid URL %s", templateURL)
		}
		return nil
	},
	Short: "Installs a template using a git URL",
	Long: `Installs a template using a git URL:

Example:
iroman install https://github.com/ottogiron/wizard-hello-world.git
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		templateURL := args[0]
		repository := git.New(ironmanHome)
		fmt.Println("Installing template", templateURL, "...")
		err := repository.Install(templateURL)
		if err != nil {
			return err
		}
		fmt.Println("Done")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(installCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// installCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// installCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
