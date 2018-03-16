package model

//Template template metadata definition
type Template struct {
	//IID internal database ID
	IID         string       `json:"iid,omitempty" yaml:"iid,omitempty"`
	ID          string       `json:"id" yaml:"id"`
	Name        string       `json:"name" yaml:"name"`
	Description string       `json:"description" yaml:"description"`
	Generators  []*Generator `json:"generators" yaml:"generators"`
}

//Type Simple type serialization for template model
func (t *Template) Type() string {
	return "model.template"
}

//Generator generator metadata definition
type Generator struct {
	ID          string `json:"id" yaml:"id"`
	Name        string `json:"name" yaml:"name"`
	Description string `json:"description" yaml:"description"`
}

//Type Simple type serialization for generator model
func (g *Generator) Type() string {
	return "model.generator"
}
