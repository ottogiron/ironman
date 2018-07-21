package template

import (
	"bytes"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/ironman-project/ironman/pkg/template/engine"
	"github.com/ironman-project/ironman/pkg/template/engine/goengine"

	"github.com/ironman-project/ironman/pkg/template/values"
	"github.com/pkg/errors"
)

const ironmanConfigFileName = ".ironman.yaml"

var rootIronmanYamlTmpl = `
version: 1.0.0
id: template-example
name: Template Example
description: This is an example of a valid template.
`

var appGeneratorYamlTmpl = `
id: app
description: Application Generator
`

var appGeneratorReadmeTmpl = `
# The App generator or "default" generator

When running the generate command without explicitly defining a generator this will be the called generator.
`

var templateReadmeTmpl = `
# Template created with ironman's create command.

Use the ***ironman link /path/to/template \<dev-template-id\>*** command to edit and run commands for this template from your file system.

There are 2 types of generators

 * Directory: Generates a directory with multiple files or directories inside.
 * File: Generates a single file.

In this template the "app" generators is of type "directory" and the "single" generator is of type "file"

## Running the app generator

The app generator is the default generator for a template.

	ironman generate <dev-template-id> /generation/path/directory

The above will run the default app generator under /generation/path/directory

## Running other generators

You must select an specific generator. The template contains a "single" generator example which will generate a single file
check the [generator .ironman.yaml file](generators/single/.ironman.yaml)

	ironman generate <dev-template-id>:single /path/to/myfilename

`

var singleFileGeneratorYamlTmpl = `
id: single
type: file
name: Single file Generator
description: Generates a controller
fileTypeOptions:
  defaultTemplateFile: file.txt
`

var singleFileTmpl = `
This is an example of a file generator
`

//Create creates a new template
func Create(templatePath string, values values.Values) error {

	var err error
	if err = createTemplateDirectories(templatePath); err != nil {
		return err
	}

	engine := goengine.New("create-template")

	var filesToWrite = []fileInfo{
		fileInfo{
			path:     filepath.Join(templatePath, ironmanConfigFileName),
			template: rootIronmanYamlTmpl,
		},
		fileInfo{
			path:     filepath.Join(templatePath, "README.md"),
			template: templateReadmeTmpl,
		},
		fileInfo{
			path:     filepath.Join(templatePath, "generators", "app", ironmanConfigFileName),
			template: appGeneratorYamlTmpl,
		},
		fileInfo{
			path:     filepath.Join(templatePath, "generators", "app", "README.md"),
			template: appGeneratorReadmeTmpl,
		},
		fileInfo{
			path:     filepath.Join(templatePath, "generators", "single", ironmanConfigFileName),
			template: singleFileGeneratorYamlTmpl,
		},
		fileInfo{
			path:     filepath.Join(templatePath, "generators", "single", "file.txt"),
			template: singleFileTmpl,
		},
	}

	err = writeFiles(engine, values, filesToWrite)

	if err != nil {
		return err
	}

	return nil

}

type fileInfo struct {
	path     string
	template string
}

func writeFiles(engine engine.Engine, values values.Values, files []fileInfo) error {
	for _, file := range files {
		err := processFile(engine, file.template, file.path, values)
		if err != nil {
			return errors.Wrapf(err, "failed to generate file %s", file.path)
		}
	}
	return nil
}

func processFile(engine engine.Engine, template string, path string, values values.Values) error {
	nEngine, err := engine.Parse(template)

	if err != nil {
		return errors.Wrapf(err, "failed to parse template %s", path)
	}

	var buffer bytes.Buffer

	err = nEngine.Execute(&buffer, values)

	if err != nil {
		return errors.Wrapf(err, "failed to process template %s", path)
	}

	err = ioutil.WriteFile(path, buffer.Bytes(), os.ModePerm)

	if err != nil {
		return errors.Wrapf(err, "failed to write template file %s", path)
	}
	return nil
}

func createTemplateDirectories(templatePath string) error {

	if templatePath == "" {
		templatePath = "."
	}

	err := os.Mkdir(templatePath, os.ModePerm)

	if err != nil {
		return errors.Wrapf(err, "failed to create template directory in path %s", templatePath)
	}

	appGeneratorPath := filepath.Join(templatePath, "generators", "app")
	err = os.MkdirAll(appGeneratorPath, os.ModePerm)
	if err != nil {
		return errors.Wrapf(err, "failed to create template app generator directory in path", appGeneratorPath)
	}

	singleFileGeneratorPath := filepath.Join(templatePath, "generators", "single")
	err = os.MkdirAll(singleFileGeneratorPath, os.ModePerm)
	if err != nil {
		return errors.Wrapf(err, "failed to create template single file generator directory in templatePath", appGeneratorPath)
	}

	return nil
}
