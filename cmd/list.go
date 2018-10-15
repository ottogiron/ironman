package cmd

import (
	"fmt"
	"io"
	"os"

	"github.com/ironman-project/ironman/pkg/ironman"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

type listCmd struct {
	out    io.Writer
	client *ironman.Ironman
}

func newListCmd(client *ironman.Ironman, out io.Writer) *cobra.Command {
	list := &listCmd{
		out:    out,
		client: client,
	}
	// listCmd represents the list command
	var listCmd = &cobra.Command{
		Use:   "list",
		Short: "Lists the available installed templates with its ID and description",
		Long: `Lists the available installed templates with its ID description:

Example:
ironman list
+------------------+------------------+--------------------------------+-------------+---------------------------------------------------------+
|        ID        |       NAME       |          DESCRIPTION           | SOURCE TYPE |                         SOURCE                          |
+------------------+------------------+--------------------------------+-------------+---------------------------------------------------------+
| template-example | Template Example | This is an example of a valid  | URL         | https://github.com/ironman-project/template-example.git |
|                  |                  | template.                      |             |                                                         |
+------------------+------------------+--------------------------------+-------------+---------------------------------------------------------+
`,
		RunE: func(cmd *cobra.Command, args []string) error {
			list.client, list.out = ensureIronmanClientAndOutput(list.client, list.out)
			return list.run()
		},
	}
	return listCmd
}

func (l *listCmd) run() error {
	fmt.Fprintln(l.out, "Installed templates")
	installedList, err := l.client.List()

	if err != nil {
		return err
	}

	if len(installedList) == 0 {
		fmt.Fprintln(l.out, "None")
		return nil

	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "Name", "Description", "Source Type", "Source"})

	for _, installed := range installedList {
		source := truncateString(installed.Source, 50) //50 is an arbitrary size
		table.Append([]string{installed.ID, installed.Name, installed.Description, string(installed.SourceType), source})
	}
	table.Render() // Send output
	return nil
}

func truncateString(str string, num int) string {
	bnoden := str
	if len(str) > num {
		if num > 3 {
			num -= 3
		}
		bnoden = str[0:num] + "..."
	}
	return bnoden
}
