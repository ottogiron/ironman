package cmd

import (
	"io"

	"github.com/ironman-project/ironman/pkg/ironman"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

type describeCmd struct {
	out        io.Writer
	client     *ironman.Ironman
	resourceID string
	format     string
}

func newDescribeCmd(client *ironman.Ironman, out io.Writer) *cobra.Command {

	describe := &describeCmd{
		out:    out,
		client: client,
	}

	var describeCmd = &cobra.Command{
		Use: "describe",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return errors.New("resource ID is required")
			}
			if len(args) > 1 {
				return errors.New("invalid number of arguments")
			}
			return nil
		},
		Short: "Displays some useful information about a resource. A resource can be a template or a generator",
		Long: `Displays some useful information about a resource. A resource can be a template or a generator
 Example:
 	$ ironman describe template-id # would display information about a template 
	$ ironman describe template-id:generator-id # would display information about a generator
		`,
		RunE: func(cmd *cobra.Command, args []string) error {
			describe.resourceID = args[0]
			describe.client, describe.out = ensureIronmanClientAndOutput(describe.client, describe.out)
			return describe.run()
		},
	}

	f := describeCmd.Flags()
	f.StringVar(&describe.format, "format", ironman.FormatYAML, "ouput format posible values yaml | json")
	return describeCmd
}

func (d *describeCmd) run() error {
	//a resource ID can be a template <template-id> or a generator <template-id>:<generator-id>
	err := d.client.Describe(d.resourceID, d.format, d.out)
	if err != nil {
		return err
	}
	return nil
}
