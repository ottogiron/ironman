package cmd

import (
	"errors"
	"fmt"
	"io"

	"github.com/ironman-project/ironman/pkg/ironman"

	"github.com/spf13/cobra"
)

type uninstallCmd struct {
	out        io.Writer
	client     *ironman.Ironman
	templateID string
}

func newUninstallCmd(client *ironman.Ironman, out io.Writer) *cobra.Command {
	uninstall := &uninstallCmd{
		out:    out,
		client: client,
	}
	// uninstallCmd represents the uninstall command
	var uninstallCmd = &cobra.Command{
		Use: "uninstall <template_ID>",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return errors.New("Template ID is required")
			}

			if len(args) > 1 {
				return errors.New("Invalid number of arguments")
			}

			return nil
		},
		Short: "Uninstalls a template by ID",
		Long: `Uninstall a template by ID
Example:

ironman uninstall my-template-id
`,
		RunE: func(cmd *cobra.Command, args []string) error {
			uninstall.templateID = args[0]
			uninstall.client, uninstall.out = ensureIronmanClientAndOutput(uninstall.client, uninstall.out)
			return uninstall.run()
		},
	}
	return uninstallCmd
}

func (u *uninstallCmd) run() error {
	fmt.Fprintln(u.out, "Uninstalling template", u.templateID, "...")
	err := u.client.Uninstall(u.templateID)
	if err != nil {
		return err
	}
	fmt.Fprintln(u.out, "done")
	return nil
}
