package cmd

import (
	"errors"
	"fmt"
	"io"

	"github.com/ironman-project/ironman/pkg/ironman"

	"github.com/spf13/cobra"
)

type unlinkCmd struct {
	out        io.Writer
	client     *ironman.Ironman
	templateID string
}

func newUnlinkCmd(client *ironman.Ironman, out io.Writer) *cobra.Command {
	unlink := &unlinkCmd{
		out:    out,
		client: client,
	}
	// unlinkCmd represents the unlink command
	var unlinkCmd = &cobra.Command{
		Use: "unlink <template_ID>",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return errors.New("Template ID is required")
			}

			return nil
		},
		Short: "Removes a symlink from the ironman repository",
		Long: `Removes a symlink from the ironman repository

Example:
ironman unlink my-template-id
	`,
		RunE: func(cmd *cobra.Command, args []string) error {
			unlink.templateID = args[0]
			unlink.client, unlink.out = ensureIronmanClientAndOutput(unlink.client, unlink.out)
			return unlink.run()
		},
	}
	return unlinkCmd
}

func (u *unlinkCmd) run() error {
	err := u.client.Unlink(u.templateID)
	fmt.Fprintf(u.out, "Unlinking template from repository with ID %s...", u.templateID)
	if err != nil {
		return err
	}
	fmt.Fprintln(u.out, "Done")
	return nil
}
