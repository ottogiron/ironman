package cmd

import (
	"fmt"
	"io"

	"github.com/spf13/cobra"
)

type versionCmd struct {
	out          io.Writer
	buildVersion string
	buildCommit  string
	buildDate    string
}

func newVersionCmd(buildVersion string, buildCommit string, buildDate string, out io.Writer) *cobra.Command {
	version := &versionCmd{
		out:          out,
		buildVersion: buildVersion,
		buildCommit:  buildCommit,
		buildDate:    buildDate,
	}

	// versionCmd represents the version command
	var versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Print the version number of Ironman",
		Long:  `All software has versions. This is Ironman's.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return version.Run()
		},
	}
	return versionCmd

}

func (v *versionCmd) Run() error {
	fmt.Fprintf(v.out, "Ironman %s-%s Build date: %s\n", v.buildVersion, v.buildCommit, v.buildDate)
	return nil
}
