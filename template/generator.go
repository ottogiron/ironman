package template

import (
	"github.com/ironman-project/ironman/template/values"
)

var _ Generator = (*generator)(nil)

//Generator defines a template generator
type Generator interface {
	Generate() error
}

type generator struct {
	path           string
	generationPath string
	values         values.Values
}

//NewGenerator returns a new instance of a generator
func NewGenerator(path string, generationPath string, values values.Values) Generator {
	return &generator{
		path,
		generationPath,
		values,
	}
}

func (g *generator) Generate() error {
	return nil
}
