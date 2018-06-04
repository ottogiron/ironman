package cmd

import (
	"fmt"
	"io"

	"github.com/ironman-project/ironman/pkg/ironman"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

type updateCmd struct {
	out        io.Writer
	client     *ironman.Ironman
	templateID string
}

func newUpdateCmd(client *ironman.Ironman, out io.Writer) *cobra.Command {
	update := &updateCmd{
		out:    out,
		client: client,
	}
	// updateCmd represents the update command
	var updateCmd = &cobra.Command{
		Use: "update <template_ID>",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return errors.New("ID arg is required")
			}

			return nil
		},
		Short: "Updates a template given an ID",
		Long: `Updates a template given an ID
Example:

ironman update my-template-id`,
		RunE: func(cmd *cobra.Command, args []string) error {
			update.templateID = args[0]
			update.client, update.out = ensureIronmanClientAndOutput(update.client, update.out)
			return update.run()
		},
	}
	return updateCmd
}

func (u *updateCmd) run() error {
	fmt.Fprintln(u.out, "Updating template", u.templateID, "...")
	err := u.client.Update(u.templateID)
	if err != nil {
		return err
	}
	fmt.Fprintln(u.out, "Done")
	return nil
}
