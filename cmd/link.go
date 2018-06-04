package cmd

import (
	"fmt"
	"io"

	"github.com/ironman-project/ironman/pkg/ironman"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

type linkCmd struct {
	out          io.Writer
	client       *ironman.Ironman
	templatePath string
	templateID   string
}

func newLinkCmd(client *ironman.Ironman, out io.Writer) *cobra.Command {
	link := &linkCmd{
		out:    out,
		client: client,
	}
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
			link.templatePath = args[0]
			link.templateID = args[1]
			link.client, link.out = ensureIronmanClientAndOutput(link.client, link.out)
			return link.run()

		},
	}
	return linkCmd
}

func (l *linkCmd) run() error {

	fmt.Fprintf(l.out, "Linking template to repository with ID %s...", l.templateID)
	err := l.client.Link(l.templatePath, l.templateID)
	if err != nil {
		return err
	}
	fmt.Fprintln(l.out, "Done")
	return nil
}
