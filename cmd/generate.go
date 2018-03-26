package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var values []string

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use: "generate <template>:<generator> <destination_path>",
	Short: `Generates a new project based on an installed template using a template generator.
			If no generator was given, it will use 'app' by default.
			It will generate the project on the destination path received (it should not exists) and
			if no destination path was given it will generate the project on the current directory (it should be empty).`,
	Long: `Generates a new project based on an installed template using a template generator.
If no generator was given, it will use 'app' by default.
It will generate the project on the destination path received (it should not exists) and
if no destination path was given it will generate the project on the current directory (it should be empty).

Example:

# This generates a project based on template-example template, based on the 'app' controller since it is the default 
# and it will generate the files on the current directory (it should be empty).
ironman generate template-example

# This generates a project based on template-example template, based on the 'controller' controller
# and it will generate the files on the current directory (it should be empty).
ironman generate template:example:controller

# This generates a project based on template-example template, based on the 'app' controller since it is the default 
# and it will generate the files on the '~/mynewapp' directory (it should not exists since it will be created now).
ironman generate template-example ~/mynewapp

# This generates a project based on template-example template, based on the 'controller' controller
# and it will generate the files on the '~/mynewapp' directory (it should not exists since it will be created now).
ironman generate template:example:controller ~/mynewapp
`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		//TODO: validate we can create the project folder and if exists it should be empty

		//We need a destination path variable (defaults to current folder)
		//If we use current folder, then it should be empty

		//If destination path was given:
		//It should not exists or
		//It can exists but it should be empty (?)

		//Find template

		//Load template

		//Gatter user input

	},
	PreRun: func(cmd *cobra.Command, args []string) {
		//TODO: we need to run the "pre generate" commands
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		//TODO: Render the template
		fmt.Println("generate called", values)
		return nil
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		//TODO: we need to run the "post generate" commands
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
	generateCmd.Flags().StringArrayVarP(&values, "set", "s", []string{}, "Coma separated list of values --set key=value, key2=value2")
}
