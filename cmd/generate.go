package cmd

import (
	"context"
	"fmt"
	"io"
	"strings"

	"github.com/ironman-project/ironman/pkg/ironman"

	"github.com/ironman-project/ironman/pkg/template/values/strvals"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

type valueFiles []string

func (v *valueFiles) String() string {
	return fmt.Sprint(*v)
}

func (v *valueFiles) Type() string {
	return "valueFiles"
}

func (v *valueFiles) Set(value string) error {
	for _, filePath := range strings.Split(value, ",") {
		*v = append(*v, filePath)
	}
	return nil
}

type generateCmd struct {
	out             io.Writer
	client          *ironman.Ironman
	templateID      string
	generatorID     string
	path            string
	values          []string
	stringValues    []string
	forceGeneration bool
	valFiles        valueFiles
}

func newGenerateCommand(client *ironman.Ironman, out io.Writer) *cobra.Command {
	generate := &generateCmd{
		out:    out,
		client: client,
	}
	// generateCmd represents the generate command
	var generateCmd = &cobra.Command{
		Use: "generate <template>:<generator> <destination_path>",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return errors.New("template ID arg is required")
			}
			return nil
		},
		Short: `Generates a new project based on an installed template using a template generator.
			If no generator was given, it will use 'app' by default.
			It will generate the project on the destination path received (it should not exists) and
			if no destination path was given it will generate the project on the current directory.`,
		Long: `Generates a new project based on an installed template using a template generator.
If no generator was given, it will use 'app' by default.
It will generate the project on the destination path received (it should not exists) and
if no destination path was given it will generate the project on the current directory.

Example:

# This generates a project based on template-example template, based on the 'app' controller since it is the default 
# and it will generate the files on the current directory.
ironman generate template-example

# This generates a project based on template-example template, based on the 'controller' controller
# and it will generate the files on the current directory.
ironman generate template-example:controller

# This generates a project based on template-example template, based on the 'app' controller since it is the default 
# and it will generate the files on the '~/mynewapp' directory (it should not exists since it will be created now).
ironman generate template-example ~/mynewapp

# This generates a project based on template-example template, based on the 'controller' controller
# and it will generate the files on the '~/mynewapp' directory.
ironman generate template:example:controller ~/mynewapp
`,
		RunE: func(cmd *cobra.Command, args []string) error {

			templateTokens := strings.Split(args[0], ":")
			templateID := templateTokens[0]
			generatorID := "app"
			path := "."
			if len(templateTokens) > 2 {
				return errors.Errorf("The generator format should be <template>:<generator>")
			}

			if len(templateTokens) == 2 {
				generatorID = templateTokens[1]
			}

			if len(args) == 2 {
				path = args[1]
			}

			generate.templateID = templateID
			generate.generatorID = generatorID
			generate.path = path
			generate.client, generate.out = ensureIronmanClientAndOutput(generate.client, generate.out)
			return generate.run()
		},
	}

	f := generateCmd.Flags()
	f.StringArrayVar(&generate.values, "set", []string{}, "set values on the command line (can specify multiple or separate values with commas: key1=val1,key2=val2)")
	f.VarP(&generate.valFiles, "values", "f", "specify values in a YAML file (can specify multiple)")
	f.BoolVar(&generate.forceGeneration, "force", false, "Forces generation even if directory or file exists. e.g ironman generate --force template /generation/path")

	return generateCmd
}

func (g *generateCmd) run() error {
	valuesReader := strvals.New(g.valFiles, g.values)
	values, err := valuesReader.Read()
	if err != nil {
		return err
	}
	fmt.Fprintln(g.out, "Running template generator", g.generatorID)
	err = g.client.Generate(context.Background(), g.templateID, g.generatorID, g.path, values, g.forceGeneration)
	if err != nil {
		return err
	}
	fmt.Fprintln(g.out, "Done")
	return nil
}
