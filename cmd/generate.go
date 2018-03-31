package cmd

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/ironman-project/ironman/template/values/strvals"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	yaml "gopkg.in/yaml.v2"
	helmstrvals "k8s.io/helm/pkg/strvals"
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

var values []string
var stringValues []string
var forceGeneration bool
var valFiles valueFiles

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

		valuesReader := strvals.New(valFiles, values)
		values, err := valuesReader.Read()

		if err != nil {
			return err
		}

		ilogger().Println("Running template generator", generatorID)
		err = iironman().Generate(context.Background(), templateID, generatorID, path, values, forceGeneration)
		if err != nil {
			return err
		}
		ilogger().Println("Done")

		return nil
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		//TODO: we need to run the "post generate" commands
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
	f := generateCmd.Flags()
	f.StringArrayVar(&values, "set", []string{}, "set values on the command line (can specify multiple or separate values with commas: key1=val1,key2=val2)")
	f.VarP(&valFiles, "values", "f", "specify values in a YAML file (can specify multiple)")
	f.BoolVar(&forceGeneration, "force", false, "Forces generation even if directory or file exists. e.g ironman generate --force template /generation/path")

}

// vals merges values from files specified via -f/--values and
// directly via --set, marshaling them to YAML
func vals(valueFiles valueFiles, values []string) ([]byte, error) {
	base := map[string]interface{}{}

	// User specified a values files via -f/--values
	for _, filePath := range valueFiles {
		currentMap := map[string]interface{}{}

		var bytes []byte
		var err error
		if strings.TrimSpace(filePath) == "-" {
			bytes, err = ioutil.ReadAll(os.Stdin)
		} else {
			bytes, err = readFile(filePath)
		}

		if err != nil {
			return []byte{}, err
		}

		if err := yaml.Unmarshal(bytes, &currentMap); err != nil {
			return []byte{}, fmt.Errorf("failed to parse %s: %s", filePath, err)
		}
		// Merge with the previous map
		base = mergeValues(base, currentMap)
	}

	// User specified a value via --set
	for _, value := range values {
		if err := helmstrvals.ParseInto(value, base); err != nil {
			return []byte{}, fmt.Errorf("failed parsing --set data: %s", err)
		}
	}

	return yaml.Marshal(base)
}

//readFile load a file from the local directory or a remote file with a url.
func readFile(filePath string) ([]byte, error) {
	return ioutil.ReadFile(filePath)
}

// Merges source and destination map, preferring values from the source map
func mergeValues(dest map[string]interface{}, src map[string]interface{}) map[string]interface{} {
	for k, v := range src {
		// If the key doesn't exist already, then just set the key to that value
		if _, exists := dest[k]; !exists {
			dest[k] = v
			continue
		}
		nextMap, ok := v.(map[string]interface{})
		// If it isn't another map, overwrite the value
		if !ok {
			dest[k] = v
			continue
		}
		// If the key doesn't exist already, then just set the key to that value
		if _, exists := dest[k]; !exists {
			dest[k] = nextMap
			continue
		}
		// Edge case: If the key exists in the destination, but isn't a map
		destMap, isMap := dest[k].(map[string]interface{})
		// If the source map has a map for this key, prefer it
		if !isMap {
			dest[k] = v
			continue
		}
		// If we got to this point, it is a map in both, so merge them
		dest[k] = mergeValues(destMap, nextMap)
	}
	return dest
}
