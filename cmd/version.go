package cmd

import (
	"github.com/spf13/cobra"
)

var (
	buildVersion string
	buildCommit  string
	buildDate    string
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Ironman",
	Long:  `All software has versions. This is Ironman's.`,
	Run: func(cmd *cobra.Command, args []string) {
		ilogger().Printf("Ironman %s-%s Build date: %s\n", buildVersion, buildCommit, buildDate)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
