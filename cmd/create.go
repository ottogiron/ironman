package cmd

import (
	"fmt"
	"io"

	"github.com/ironman-project/ironman/pkg/ironman"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

type createCmd struct {
	out    io.Writer
	client *ironman.Ironman
	path   string
}

func newCreateCmd(client *ironman.Ironman, out io.Writer) *cobra.Command {
	create := &createCmd{
		out:    out,
		client: client,
	}
	// createCmd represents the create command
	var createCmd = &cobra.Command{
		Use: "create",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return errors.New("template path is required")
			}

			if len(args) > 1 {
				return errors.New("invalid number of arguments")
			}

			return nil
		},
		Short: "Creates a new ironman base template",
		Long: `Creates anew ironman base template. For example:

ironman create mytemplate`,
		RunE: func(cmd *cobra.Command, args []string) error {
			create.path = args[0]
			create.client, create.out = ensureIronmanClientAndOutput(create.client, create.out)
			return create.run()
		},
	}

	return createCmd
}

func (c *createCmd) run() error {
	fmt.Fprintf(c.out, "Creating new template %s... \n", c.path)
	err := c.client.Create(c.path)

	if err != nil {
		return err
	}

	fmt.Fprintln(c.out, "Done")
	return nil
}
