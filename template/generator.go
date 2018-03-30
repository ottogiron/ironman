package template

import (
	"bytes"
	"context"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/ironman-project/ironman/template/engine"
	"github.com/ironman-project/ironman/template/engine/goengine"
	"github.com/ironman-project/ironman/template/model"
	"github.com/ironman-project/ironman/template/values"
	"github.com/pkg/errors"
)

//arbitrary number
const noGeneratorWorkers = 20

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
	path           string
	generationPath string
	ignore         []string
	data           GeneratorData
	engine         engine.Factory
	logger         *log.Logger
	force          bool
}

//NewGenerator returns a new instance of a generator
func NewGenerator(path string, generationPath string, data GeneratorData, options ...GeneratorOption) Generator {

	g := &generator{
		path:           path,
		generationPath: generationPath,
		data:           data,
		ignore:         []string{".ironman.yaml"},
		engine: func() engine.Engine {
			return goengine.New("ironman")
		},
		logger: log.New(os.Stdout, "", 0),
		force:  false,
	}

	for _, option := range options {
		option(g)
	}

	return g
}

type processResult struct {
	bytes      []byte
	pathResult pathResult
	err        error
}

type writeResult struct {
	pathFrom string
	pathTo   string
	err      error
}

type pathResult struct {
	path  string
	isDir bool
}

func (g *generator) Generate(ctx context.Context) error {

	childCtx, cancelFunc := context.WithCancel(ctx)
	defer cancelFunc()

	paths, errc := g.walkFiles(childCtx)

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
			g.logger.Print("Processing failed for", wresult.pathTo)
			return wresult.err
		}
	}

	err := <-errc

	if err != nil {
		return errors.Wrapf(err, "Failed to process generator path templates: %s", g.path)
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

func (g *generator) walkFiles(context context.Context) (<-chan pathResult, <-chan error) {
	errc := make(chan error, 1)
	paths := make(chan pathResult)

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
			case paths <- pathResult{path, info.IsDir()}:
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

func (g *generator) processor(context context.Context, paths <-chan pathResult, result chan<- processResult) {
	for path := range paths {
		data, err := g.processFile(path)
		select {
		case result <- processResult{data, path, err}:
		case <-context.Done():
			return
		}
	}
}

func (g *generator) processFile(pathResult pathResult) ([]byte, error) {

	if pathResult.isDir {
		return nil, nil
	}

	data, err := ioutil.ReadFile(pathResult.path)
	if err != nil {
		return nil, errors.Wrapf(err, "Failed to read template contents", pathResult.path)
	}
	engine := g.engine()
	tmpl, err := engine.Parse(string(data))
	var buffer bytes.Buffer
	err = tmpl.Execute(&buffer, g.data)

	if err != nil {
		return nil, errors.Wrapf(err, "Failed to execute template processing %s", pathResult.path)
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

	toRelativePath := strings.TrimPrefix(presult.pathResult.path, g.path)

	toPath := filepath.Join(g.generationPath, toRelativePath)

	if presult.pathResult.isDir {

		//remove the directory
		if g.force {
			err := os.RemoveAll(toPath)
			if err != nil {
				return writeResult{err: errors.Wrap(err, "Failed to force  generation")}
			}
		}

		err := os.Mkdir(toPath, os.ModePerm)
		if err != nil {
			return writeResult{err: errors.Wrap(err, "Failed to create directory")}
		}

		return writeResult{pathFrom: presult.pathResult.path, pathTo: toPath}
	}
	g.logger.Print("Writing... ", toPath)
	err := ioutil.WriteFile(toPath, presult.bytes, os.ModePerm)

	if err != nil {
		return writeResult{err: err}
	}

	return writeResult{pathFrom: presult.pathResult.path, pathTo: toPath}
}
