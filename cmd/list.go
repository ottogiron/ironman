package cmd

import (
	"fmt"
	"os"

	"github.com/ironman-project/ironman/template/repository/git"
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
+------+-----------------------+--------------+
| ID   |      Name             |  Description |
+------+-----------------------+--------------+
|  A   |       The Good        |    500 	  |
|  B   | The Very very Bad Man |    288       |
|  C   |       The Ugly        |    120       |
|  D   |      The Gopher       |    800       |
+------+-----------------------+--------------+

`,
	RunE: func(cmd *cobra.Command, args []string) error {

		repository := git.New(ironmanHome)
		fmt.Println("Installed templates")
		installedList, err := repository.Installed()
		if err != nil {
			return err
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
