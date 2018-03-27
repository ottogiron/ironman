package template

import (
	"context"
	"os"
	"path/filepath"

	"github.com/ironman-project/ironman/template/model"
	"github.com/ironman-project/ironman/template/values"
	"github.com/pkg/errors"
)

var _ Generator = (*generator)(nil)

//Generator defines a template generator
type Generator interface {
	Generate(context context.Context) error
}

type generator struct {
	path           string
	generationPath string
	values         values.Values
	engine         Engine
}

//NewGenerator returns a new instance of a generator
func NewGenerator(path string, generationPath string, templateModel *model.Template, values values.Values, engine Engine) Generator {
	return &generator{
		path,
		generationPath,
		values,
		engine,
	}
}

type processResult struct {
	text string
	path string
	err  error
}

func (g *generator) Generate(ctx context.Context) error {
	childContext, cancelFunc := context.WithCancel(ctx)
	defer cancelFunc()

	processResult, errc := g.process(childContext)

	err := <-errc

	if err != nil {
		return errors.Wrapf(err, "Failed to process generator path templates %s", g.path)
	}

	for result := range processResult {
		if result.err != nil {
			cancelFunc()
		}
	}
	return nil
}

func (g *generator) write(context context.Context, procesResults <-chan processResult) {

}

func (g *generator) process(context context.Context) (<-chan processResult, <-chan error) {
	errc := make(chan error, 1)
	go func() {
		err := filepath.Walk(g.path, func(path string, info os.FileInfo, err error) error {
			return nil
		})
		errc <- err
	}()

	return nil, errc
}
