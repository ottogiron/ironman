package template

var _ Generator = (*generator)(nil)

//Generator defines a template generator
type Generator interface {
	Generate() error
}

type generator struct {
}

func (g *generator) Generate() error {
	return nil
}
