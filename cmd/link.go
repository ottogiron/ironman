package cmd

import (
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

// linkCmd represents the link command
var linkCmd = &cobra.Command{
	Use: "link <template_path> <template_ID>",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 2 {
			return errors.New("template path and symlink name are required")
		}

		return nil
	},
	Short: "Creates a symlink of a ironman template to the ironman repository",
	Long: `Creates a symlink of a ironman template to the ironman repository:

Example:
ironman link /path/to/template dev-template

If you run "ironman list" you should see the symlink of your template created.
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		templatePath := args[0]
		templateID := args[1]

		ilogger().Printf("Linking template to repository with ID %s...", templateID)
		err := iironman().Link(templatePath, templateID)
		if err != nil {
			return err
		}
		ilogger().Println("Done")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(linkCmd)

}
