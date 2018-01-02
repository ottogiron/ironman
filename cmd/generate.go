package cmd

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
)

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generates a new project based on an installed template",
	Long: `Generates a new project based on an installed template,

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("generate called")
		return errors.New("hola")
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)

}
