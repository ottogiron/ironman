package cmd

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
)

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use: "generate <template>:<generator>",
	Short: `Generates a new project based on an installed template using a template generator.
			If no generator was given, it will use 'app' by default.`,
	Long: `Generates a new project based on an installed template using a template generator.
If no generator was given, it will use 'app' by default.

Example:
ironman generate template-example
ironman generate template:example:controller
`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		//TODO: validate we can create the project folder and if exists it should be empty
	},
	PreRun: func(cmd *cobra.Command, args []string) {
		//TODO: we need to run the "pre generate" commands
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		//TODO: Render the template
		fmt.Println("generate called")
		return errors.New("hola")
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		//TODO: we need to run the "post generate" commands
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)

}
