package cmd

import (
	"fmt"

	"github.com/ironman-project/ironman/template/repository/git"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use: "update <template_ID>",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("ID arg is required")
		}

		return nil
	},
	Short: "Updates a template given an ID",
	Long: `Updates a template given an ID
Example:

ironman update my-template-id`,
	RunE: func(cmd *cobra.Command, args []string) error {
		templateID := args[0]
		repository := git.New(ironmanHome)
		fmt.Println("Updating template", templateID, "...")
		err := repository.Update(templateID)
		if err != nil {
			return err
		}
		fmt.Println("Done")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)

}
