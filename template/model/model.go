package model

//Template template metadata definition
type Template struct {
	//IID internal database ID
	IID           string       `json:"iid,omitempty" yaml:"iid,omitempty"`
	ID            string       `json:"id" yaml:"id"`
	Version       string       `json:"version" yaml:"version"`
	Name          string       `json:"name" yaml:"name"`
	Description   string       `json:"description" yaml:"description"`
	Generators    []*Generator `json:"generators" yaml:"generators"`
	DirectoryName string       `json:"directory_name,omitempty" yaml:"directory_name,omitempty"`
}

//Type Simple type serialization for template model
func (t *Template) Type() string {
	return "model.template"
}

//Generator returns a generator by ID, nil  if not exists
func (t *Template) Generator(ID string) *Generator {
	for _, generator := range t.Generators {
		if generator.ID == ID {
			return generator
		}
	}
	return nil
}

//Generator generator metadata definition
type Generator struct {
	ID            string `json:"id" yaml:"id"`
	Name          string `json:"name" yaml:"name"`
	Description   string `json:"description" yaml:"description"`
	DirectoryName string `json:"directory_name,omitempty" yaml:"directory_name,omitempty"`
}

//Type Simple type serialization for generator model
func (g *Generator) Type() string {
	return "model.generator"
}
