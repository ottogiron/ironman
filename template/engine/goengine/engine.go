package goengine

import (
	"io"
	gtemplate "text/template"

	"github.com/ironman-project/ironman/template/engine"
)

var _ engine.Engine = (*goEngine)(nil)

type goEngine struct {
	template *gtemplate.Template
}

//New returns a new instance of a go template engine
func New(name string) engine.Engine {
	template := gtemplate.New(name)
	return &goEngine{template: template}
}

func (g *goEngine) Parse(text string) (engine.Engine, error) {
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
