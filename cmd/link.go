package cmd

import (
	"fmt"

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
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("link called")
	},
}

func init() {
	rootCmd.AddCommand(linkCmd)

}
