package cmd

import (
	"errors"
	"fmt"

	"github.com/ironman-project/ironman/ironman"
	"github.com/spf13/cobra"
)

// uninstallCmd represents the uninstall command
var uninstallCmd = &cobra.Command{
	Use: "uninstall <template_ID>",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("ID arg is required")
		}

		return nil
	},
	Short: "Uninstalls a template by ID",
	Long: `Uninstall a template by ID
Example:

ironman uninstall my-template-id
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		templateID := args[0]
		ironman := ironman.New(ironmanHome)
		fmt.Println("Uninstalling template", templateID, "...")
		err := ironman.Uninstall(templateID)
		if err != nil {
			return err
		}
		fmt.Println("done")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(uninstallCmd)

}
