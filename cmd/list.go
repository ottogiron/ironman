package cmd

import (
	"fmt"
	"os"

	"github.com/ironman-project/ironman/template/manager/git"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists the available installed templates with its ID and description",
	Long: `Lists the available installed templates with its ID description:

Example:
ironman list

+--------------------+--------------------+-------------+
|         ID         |        NAME        | DESCRIPTION |
+--------------------+--------------------+-------------+
| template-example   | template-example   | TODO        |
+--------------------+--------------------+-------------+
`,
	RunE: func(cmd *cobra.Command, args []string) error {

		manager := git.New(ironmanHome)
		fmt.Println("Installed templates")
		installedList, err := manager.Installed()
		if err != nil {
			return err
		}

		if len(installedList) < 0 {
			fmt.Println("None")
		}
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"ID", "Name", "Description"})

		for _, installed := range installedList {
			table.Append([]string{installed.ID, installed.ID, "TODO"})
		}
		table.Render() // Send output
		return nil
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

}
