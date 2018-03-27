package template

import (
	"io"
	gtemplate "text/template"
)

var _ Engine = (*goEngine)(nil)

//Engine represents a template engine
type Engine interface {
	Parse(text string) (Engine, error)
	Execute(writer io.Writer, data interface{}) error
}

type goEngine struct {
	template *gtemplate.Template
}

//NewGoEngine returns a new instance of a go template engine
func NewGoEngine(name string) Engine {
	template := gtemplate.New(name)
	return &goEngine{template}
}

func (g *goEngine) Parse(text string) (Engine, error) {
	var err error
	g.template, err = g.template.Parse(text)
	if err != nil {
		return nil, err
	}

	return g, nil
}

func (g *goEngine) Execute(writer io.Writer, data interface{}) error {
	return g.template.Execute(writer, data)
}
