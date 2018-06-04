package cmd

import (
	"fmt"
	"io"

	"github.com/spf13/cobra"
)

var (
	buildVersion string
	buildCommit  string
	buildDate    string
)

type versionCmd struct {
}

func newVersionCmd(out io.Writer) *cobra.Command {
	// versionCmd represents the version command
	var versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Print the version number of Ironman",
		Long:  `All software has versions. This is Ironman's.`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Fprintf(out, "Ironman %s-%s Build date: %s\n", buildVersion, buildCommit, buildDate)
		},
	}
	return versionCmd

}
