package cmd

import (
	"errors"

	"github.com/spf13/cobra"
)

// unlinkCmd represents the unlink command
var unlinkCmd = &cobra.Command{
	Use: "unlink <template_ID>",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("Template ID is required")
		}

		return nil
	},
	Short: "Removes a symlink from the ironman repository",
	Long: `Removes a symlink from the ironman repository

Example:
ironman unlink my-template-id
	`,
	RunE: func(cmd *cobra.Command, args []string) error {
		templateID := args[0]

		err := iironman().Unlink(templateID)
		ilogger().Printf("Unlinking template from repository with ID %s...", templateID)
		if err != nil {
			return err
		}
		ilogger().Println("Done")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(unlinkCmd)
}
