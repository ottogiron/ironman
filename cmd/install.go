package cmd

import (
	"fmt"
	"io"

	"github.com/ironman-project/ironman/pkg/ironman"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

type installCmd struct {
	out             io.Writer
	client          *ironman.Ironman
	templateLocator string
}

func newInstallCommand(client *ironman.Ironman, out io.Writer) *cobra.Command {
	install := &installCmd{
		out:    out,
		client: client,
	}
	// installCmd represents the install command
	var installCmd = &cobra.Command{
		Use: "install <url>",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return errors.New("url arg is required")
			}

			if len(args) > 1 {
				return errors.New("Invalid number of arguments")
			}

			return nil
		},
		Short: "Installs a template using a git URL",
		Long: `Installs a template using a git URL:

Example:
iroman install https://github.com/ironman-project/template-example.git
`,
		RunE: func(cmd *cobra.Command, args []string) error {
			install.templateLocator = args[0]
			install.client, install.out = ensureIronmanClientAndOutput(install.client, install.out)
			return install.run()
		},
	}
	return installCmd
}

func (i *installCmd) run() error {
	fmt.Fprintln(i.out, "Installing template", i.templateLocator, "...")
	err := i.client.Install(i.templateLocator)
	if err != nil {
		return err
	}
	fmt.Fprintln(i.out, "Done")
	return nil
}
