package cmd

import (
	"fmt"

	"github.com/ironman-project/ironman/template/repository/git"
	"github.com/spf13/cobra"
)

// linkCmd represents the link command
var linkCmd = &cobra.Command{
	Use:   "link",
	Short: "Creates a symlink of a ironman template to the ironman repository",
	Long: `Creates a symlink of a ironman template to the ironman repository:

Example:
ironman link /path/to/template dev-template

If you run "ironman list" you should see the symlink of your template created.
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		templatePath := args[0]
		templateID := args[1]
		repository := git.New(ironmanHome)
		fmt.Printf("Linking template to repository with ID %s...", templateID)
		err := repository.Link(templatePath, templateID)
		if err != nil {
			return err
		}
		fmt.Println("Done")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(linkCmd)

}
