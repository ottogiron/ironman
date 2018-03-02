package model

//Template template metadata definition
type Template struct {
	ID          string
	Name        string
	Description string
	Generators  []*Generator
}

//Type Simple type serialization for template model
func (t *Template) Type() string {
	return "model.template"
}

//Generator generator metadata definition
type Generator struct {
	ID          string
	Name        string
	Description string
}

//Type Simple type serialization for generator model
func (g *Generator) Type() string {
	return "model.generator"
}
