package template

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/ironman-project/ironman/pkg/template/engine"
	"github.com/ironman-project/ironman/pkg/template/engine/goengine"
	"github.com/ironman-project/ironman/pkg/template/model"
	"github.com/ironman-project/ironman/pkg/template/values"
	"github.com/pkg/errors"
)

//arbitrary number
const noGeneratorWorkers = 20
const (
	preGenerateLabel  = "pre-generate"
	postGenerateLabel = "post-generate"
)

//GeneratorData represents the data to be passed to each generator file template
type GeneratorData struct {
	Template  *model.Template
	Generator *model.Generator
	Values    values.Values
}

var _ Generator = (*generator)(nil)

//Generator defines a template generator
type Generator interface {
	Generate(context context.Context) error
}

type generator struct {
	path                  string
	generationPath        string
	ignore                []string
	data                  GeneratorData
	engineFactory         engine.Factory
	out                   io.Writer
	withPreGenerateHooks  bool
	withPostGenerateHooks bool
}

//NewGenerator returns a new instance of a generator
func NewGenerator(path string, generationPath string, data GeneratorData, options ...GeneratorOption) Generator {

	g := &generator{
		path:           path,
		generationPath: generationPath,
		data:           data,
		ignore:         []string{".ironman.yaml"},
		engineFactory: func() engine.Engine {
			return goengine.New("ironman")
		},
		out:                   os.Stdout,
		withPreGenerateHooks:  true,
		withPostGenerateHooks: true,
	}

	for _, option := range options {
		option(g)
	}

	return g
}

type processResult struct {
	bytes              []byte
	templatePathResult templatePathResult
	err                error
}

type writeResult struct {
	pathFrom string
	pathTo   string
	err      error
}

type templatePathResult struct {
	path  string
	isDir bool
}

func (g *generator) Generate(ctx context.Context) error {
	gdata := g.data.Generator

	if g.withPreGenerateHooks {
		err := g.runPreGenerateHooks()
		if err != nil {
			return errors.Errorf("failed to run %s hooks", preGenerateLabel)
		}
	}

	//Generate a file only if the generator type is file
	if g.data.Generator.TType == model.GeneratorTypeFile {
		if gdata.FileTypeOptions.DefaultTemplateFile == "" {
			return errors.Errorf("The default template file for the file generator %s is not set", gdata.ID)
		}
		templateFilePath := filepath.Join(g.path, gdata.FileTypeOptions.DefaultTemplateFile)
		presult := templatePathResult{templateFilePath, false}
		bytes, err := g.processFile(presult)
		if err != nil {
			return errors.Wrapf(err, "failed to process generator %s for template %s", gdata.ID, templateFilePath)
		}

		wr := g.writeFile(processResult{
			bytes,
			presult,
			nil,
		})

		if wr.err != nil {
			return wr.err
		}

		return nil
	}

	//The default if type is empty is directory
	childCtx, cancelFunc := context.WithCancel(ctx)
	defer cancelFunc()

	paths, errc := g.walkTemplateFiles(childCtx)

	presults := make(chan processResult)

	workersExecute(noGeneratorWorkers, func(w int, wg *sync.WaitGroup) {
		g.processor(childCtx, paths, presults)
		wg.Done()
	}, func() {
		close(presults)

	})

	wresults := make(chan writeResult)
	workersExecute(noGeneratorWorkers, func(w int, wg *sync.WaitGroup) {
		g.write(childCtx, presults, wresults)
		wg.Done()
	},
		func() {
			close(wresults)

		},
	)

	for wresult := range wresults {

		if wresult.err != nil {
			cancelFunc()
			return wresult.err
		}
	}

	err := <-errc

	if err != nil {
		return errors.Wrapf(err, "failed to process generator path templates: %s", g.path)
	}

	if g.withPostGenerateHooks {
		err := g.runPostGenerateHooks()
		if err != nil {
			return errors.Errorf("faield to run %s hooks", postGenerateLabel)
		}
	}

	return nil
}

func workersExecute(number int, work func(workerID int, wg *sync.WaitGroup), done func()) {
	var wg sync.WaitGroup
	wg.Add(number)
	for i := 0; i < number; i++ {
		go work(i, &wg)
	}
	go func() {
		wg.Wait()
		done()
	}()
}

func (g *generator) walkTemplateFiles(context context.Context) (<-chan templatePathResult, <-chan error) {
	errc := make(chan error, 1)
	paths := make(chan templatePathResult)

	go func() {
		defer close(paths)
		defer close(errc)
		errc <- filepath.Walk(g.path, func(path string, info os.FileInfo, err error) error {

			if err != nil {
				return err
			}

			if info.IsDir() && path == g.path {
				return nil
			}

			if !info.IsDir() && !info.Mode().IsRegular() {
				return nil
			}

			if g.ignoreFile(filepath.Base(path)) {
				return nil
			}

			select {
			case paths <- templatePathResult{path, info.IsDir()}:
			case <-context.Done():
				return errors.New("Walk canceled")

			}
			return nil
		})
	}()

	return paths, errc
}

func (g *generator) ignoreFile(fileName string) bool {
	for _, ignore := range g.ignore {
		if ignore == fileName {
			return true
		}
	}
	return false
}

func (g *generator) processor(context context.Context, paths <-chan templatePathResult, result chan<- processResult) {
	for path := range paths {
		data, err := g.processFile(path)
		select {
		case result <- processResult{data, path, err}:
		case <-context.Done():
			return
		}
	}
}

func (g *generator) processFile(templatePathResult templatePathResult) ([]byte, error) {

	if templatePathResult.isDir {
		return nil, nil
	}

	data, err := ioutil.ReadFile(templatePathResult.path)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to read template contents %s", templatePathResult.path)
	}
	engine := g.engineFactory()
	tmpl, err := engine.Parse(string(data))

	if err != nil {
		return nil, errors.Wrapf(err, "failed to parse template %s %s ", templatePathResult.path, err)
	}

	var buffer bytes.Buffer
	err = tmpl.Execute(&buffer, g.data)

	if err != nil {
		return nil, errors.Wrapf(err, "failed to execute template processing %s", templatePathResult.path)
	}
	return buffer.Bytes(), nil
}

func (g *generator) write(context context.Context, processResults <-chan processResult, result chan<- writeResult) {

	for processResult := range processResults {
		select {
		case result <- g.writeFile(processResult):
		case <-context.Done():
			return
		}
	}

}

func (g *generator) writeFile(presult processResult) writeResult {

	if presult.err != nil {
		return writeResult{err: presult.err}
	}

	toRelativePath := strings.TrimPrefix(presult.templatePathResult.path, g.path)
	generationDir := g.generationPath
	if g.data.Generator.TType == model.GeneratorTypeFile {
		//Join relative extra path from the defined generation path
		//e.g ironman generate template:controller /path/to/newController.go
		//Generation path => controller.go
		//Base Path => /path/to
		//Generator defined Relative path to base path controllers (directory)
		//output should be /path/to/controllers/newController.go
		basePath := filepath.Dir(toRelativePath)
		fileName := filepath.Base(g.generationPath)
		newPath := filepath.Join(basePath, g.data.Generator.FileTypeOptions.FileGenerationRelativePath, fileName)
		toRelativePath = newPath
		generationDir = filepath.Dir(generationDir)
	}

	toPath := filepath.Join(generationDir, toRelativePath)

	if presult.templatePathResult.isDir {

		return writeResult{pathFrom: presult.templatePathResult.path, pathTo: toPath}
	}

	fmt.Fprintln(g.out, "Writing... ", toPath)

	//Create directory
	dir := filepath.Dir(toPath)
	if _, err := os.Stat(dir); os.IsNotExist(err) {

		err := os.MkdirAll(dir, os.ModePerm)
		if err != nil && !os.IsExist(err) {
			return writeResult{err: errors.Wrap(err, "failed to create generation directory")}
		}

	}

	err := ioutil.WriteFile(toPath, presult.bytes, os.ModePerm)

	if err != nil {
		return writeResult{err: err}
	}
	return writeResult{pathFrom: presult.templatePathResult.path, pathTo: toPath}
}

func (g *generator) runPreGenerateHooks() error {
	hooks := g.data.Generator.Hooks
	if hooks != nil {
		return g.runHooks(preGenerateLabel, hooks.PreGenerate)
	}
	return nil
}

func (g *generator) runPostGenerateHooks() error {
	hooks := g.data.Generator.Hooks
	if hooks != nil {
		return g.runHooks(postGenerateLabel, hooks.PostGenerate)
	}
	return nil
}

func (g *generator) runHooks(name string, hooks []*model.Command) error {

	if len(hooks) < 1 { //if it doesn't have at least one hook
		return nil // do nothing
	}

	fmt.Fprintf(g.out, "Running %s hooks\n", name)
	for _, hookCommand := range hooks {
		if err := g.executeCommand(hookCommand); err != nil {
			return errors.Errorf("failed to execute %s hook %s %s", name, hookCommand.Name, err)
		}
	}
	fmt.Fprintf(g.out, "\n...Running %s hooks done\n", name)

	return nil
}

func (g *generator) executeCommand(command *model.Command) error {

	engine := g.engineFactory()
	nameTmpl, err := engine.Parse(command.Name)
	if err != nil {
		return errors.Errorf("failed to parse command hook name template %s", command.Name)
	}

	var buffer bytes.Buffer
	if err := nameTmpl.Execute(&buffer, g.data); err != nil {
		return errors.Errorf("failed to execute template for command hook name %s", command.Name)
	}

	command.Name = buffer.String() //extract the parsed command name

	//now do the same for each argument
	var cmdArgs []string

	for _, argument := range command.Args {

		argTmpl, err := engine.Parse(argument)
		if err != nil {
			return errors.Errorf("failed to parse command hook argument template %s, %s", argument, err)
		}

		buffer.Reset() // reset the buffer to store the new parsed template
		if err := argTmpl.Execute(&buffer, g.data); err != nil {
			return errors.Errorf("failed to execute template for command hook argument %s %s", argument, err)
		}

		cmdArg := buffer.String()
		cmdArgs = append(cmdArgs, cmdArg)
	}

	command.Args = cmdArgs

	if err := ExecuteCommand(command, g.out); err != nil {
		return err // not wrapping just return the original error an wrap in the calling function above
	}
	return nil
}
