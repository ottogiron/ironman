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

+--------------------+--------------------+-------------+
|         ID         |        NAME        | DESCRIPTION |
+--------------------+--------------------+-------------+
| template-example   | template-example   | TODO        |
+--------------------+--------------------+-------------+
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
	table.SetHeader([]string{"ID", "Name", "Description"})

	for _, installed := range installedList {
		table.Append([]string{installed.ID, installed.Name, installed.Description})
	}
	table.Render() // Send output
	return nil
}
