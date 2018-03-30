package goengine

import (
	"html/template"
	"io"
	gtemplate "text/template"

	"github.com/Masterminds/sprig"
	"github.com/ironman-project/ironman/template/engine"
	"github.com/kubernetes/helm/pkg/chartutil"
)

var _ engine.Engine = (*goEngine)(nil)

type goEngine struct {
	template *gtemplate.Template
}

//New returns a new instance of a go template engine
func New(name string) engine.Engine {
	template := gtemplate.New(name)
	template.Funcs(FuncMap())
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

// FuncMap returns a mapping of all of the functions that Engine has.
//
// Because some functions are late-bound (e.g. contain context-sensitive
// data), the functions may not all perform identically outside of an
// Engine as they will inside of an Engine.
//
// Known late-bound functions:
//
//	- "include": This is late-bound in Engine.Render(). The version
//	   included in the FuncMap is a placeholder.
//      - "required": This is late-bound in Engine.Render(). The version
//	   included in the FuncMap is a placeholder.
//      - "tpl": This is late-bound in Engine.Render(). The version
//	   included in the FuncMap is a placeholder.
func FuncMap() template.FuncMap {
	f := sprig.TxtFuncMap()
	delete(f, "env")
	delete(f, "expandenv")

	// Add some extra functionality
	extra := template.FuncMap{
		"toToml":   chartutil.ToToml,
		"toYaml":   chartutil.ToYaml,
		"fromYaml": chartutil.FromYaml,
		"toJson":   chartutil.ToJson,
		"fromJson": chartutil.FromJson,

		// This is a placeholder for the "include" function, which is
		// late-bound to a template. By declaring it here, we preserve the
		// integrity of the linter.
		"include":  func(string, interface{}) string { return "not implemented" },
		"required": func(string, interface{}) interface{} { return "not implemented" },
		"tpl":      func(string, interface{}) interface{} { return "not implemented" },
	}

	for k, v := range extra {
		f[k] = v
	}

	return f
}
