package template

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/ironman-project/ironman/template/engine"
	"github.com/ironman-project/ironman/template/model"
	"github.com/ironman-project/ironman/template/values"
	"github.com/pkg/errors"
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
	path           string
	generationPath string
	ignore         []string
	data           GeneratorData
	engine         engine.Factory
}

//NewGenerator returns a new instance of a generator
func NewGenerator(path string, generationPath string, ignore []string, data GeneratorData, engineFactory engine.Factory) Generator {
	return &generator{
		path,
		generationPath,
		ignore,
		data,
		engineFactory,
	}
}

type processResult struct {
	bytes []byte
	path  string
	err   error
}

type writeResult struct {
	pathFrom string
	pathTo   string
	err      error
}

func (g *generator) Generate(ctx context.Context) error {

	childCtx, cancelFunc := context.WithCancel(ctx)
	defer cancelFunc()

	paths, errc := g.walkFiles(childCtx)

	presults := make(chan processResult)

	workersExecute(1, func(w int, wg *sync.WaitGroup) {
		g.processor(childCtx, paths, presults)
		fmt.Println("Worker process ID Leaving:", w)
		wg.Done()
	}, func() {
		close(presults)
		fmt.Println("Process workers closed")
	})

	wresults := make(chan writeResult)
	workersExecute(1, func(w int, wg *sync.WaitGroup) {
		g.write(childCtx, presults, wresults)
		fmt.Println("Worker write ID Leaving:", w)
		wg.Done()
	},
		func() {
			close(wresults)
			fmt.Println("Write workers closed")
		},
	)

	for wresult := range wresults {
		if wresult.err != nil {
			cancelFunc()
			return wresult.err
		}
		fmt.Printf("\nPath to is:%v\n", wresult)
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

func (g *generator) walkFiles(context context.Context) (<-chan string, <-chan error) {
	errc := make(chan error, 1)
	paths := make(chan string)

	go func() {
		defer close(paths)
		defer close(errc)
		errc <- filepath.Walk(g.path, func(path string, info os.FileInfo, err error) error {

			if err != nil {
				return err
			}

			if !info.Mode().IsRegular() {
				return nil
			}

			if info.IsDir() {
				return nil
			}

			if g.ignoreFile(filepath.Base(path)) {
				return nil
			}

			select {
			case paths <- path:
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

func (g *generator) processor(context context.Context, paths <-chan string, result chan<- processResult) {
	for path := range paths {
		data, err := g.processFile(path)
		select {
		case result <- processResult{data, path, err}:
		case <-context.Done():
			return
		}
	}
}

func (g *generator) processFile(path string) ([]byte, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, errors.Wrapf(err, "Failed to read template contents", path)
	}
	engine := g.engine()
	tmpl, err := engine.Parse(string(data))
	var buffer bytes.Buffer
	err = tmpl.Execute(&buffer, g.data)

	if err != nil {
		return nil, errors.Wrapf(err, "Failed to execute template processing %s", path)
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

	toRelativePath := strings.TrimPrefix(presult.path, g.path)

	toPath := filepath.Join(g.generationPath, toRelativePath)

	err := ioutil.WriteFile(toPath, presult.bytes, os.ModePerm)

	if err != nil {
		return writeResult{err: err}
	}

	return writeResult{pathFrom: presult.path, pathTo: toPath}
}
