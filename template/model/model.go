package model

//Mantainer  type for a template mantainer
type Mantainer struct {
	Name  string `json:"name" yaml:"name"`
	Email string `json:"email" yaml:"email"`
	URL   string `json:"url" yaml:"url"`
}

//Template template metadata definition
type Template struct {
	//IID internal database ID
	IID string `json:"iid,omitempty" yaml:"iid,omitempty"`
	ID  string `json:"id" yaml:"id"`

	Version       string       `json:"version" yaml:"version"`
	Name          string       `json:"name" yaml:"name"`
	Description   string       `json:"description" yaml:"description"`
	Generators    []*Generator `json:"generators" yaml:"generators"`
	DirectoryName string       `json:"directory_name,omitempty" yaml:"directory_name,omitempty"`
	HomeURL       string       `json:"home" yaml:"home"`
	Sources       []string     `json:"sources" yaml:"sources"`
	Mantainers    []*Mantainer `json:"mantainers" yaml:"mantainers"`
	AppVersion    string       `json:"app_version" yaml:"app_version"`
	Deprecated    bool         `json:"deprecated" yaml:"deprecated"`
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

//GeneratorType represents a generator type, directory or file
type GeneratorType string

const (
	//GeneratorTypeDirectory represents the type of a directory generator
	GeneratorTypeDirectory GeneratorType = "directory"
	//GeneratorTypeFile represents the type of a file generator
	GeneratorTypeFile GeneratorType = "file"
)

//Generator generator metadata definition
type Generator struct {
	ID                  string        `json:"id" yaml:"id"`
	TType               GeneratorType `json:"type" yaml:"type"`
	Name                string        `json:"name" yaml:"name"`
	Description         string        `json:"description" yaml:"description"`
	DirectoryName       string        `json:"directory_name,omitempty" yaml:"directory_name,omitempty"`
	FileTypeDefaultFile string        `json:"file_type_default_file" yaml:"file_type_default_file"`
}

//Type Simple type serialization for generator model
func (g *Generator) Type() string {
	return "model.generator"
}
