package model

import "time"

//SourceType represents how the template has been installed
type SourceType string

const (
	//SourceTypeURL the template has been installed from a remote source
	SourceTypeURL SourceType = "URL"
	//SourceTypeLink the template has been installed as a file system link
	SourceTypeLink = "Link"
)

//Mantainer  type for a template mantainer
type Mantainer struct {
	Name  string `json:"name" yaml:"name"`
	Email string `json:"email" yaml:"email"`
	URL   string `json:"url" yaml:"url"`
}

//Template template metadata definition
type Template struct {
	ID            string       `json:"id" yaml:"id" storm:"id"` //contains an special storm annotation
	SourceType    SourceType   `json:"sourceType,omitempty" yaml:"sourceType,omitempty"`
	Source        string       `json:"source,omitempty" yaml:"source,omitempty"`
	Version       string       `json:"version" yaml:"version"`
	Name          string       `json:"name" yaml:"name"`
	Description   string       `json:"description" yaml:"description"`
	Generators    []*Generator `json:"generators" yaml:"generators"`
	DirectoryName string       `json:"directoryName" yaml:"-"`
	HomeURL       string       `json:"home,omitempty" yaml:"home,omitempty"`
	Sources       []string     `json:"sources,omitempty" yaml:"sources,omitempty"`
	Mantainers    []*Mantainer `json:"mantainers,omitempty" yaml:"mantainers,omitempty"`
	AppVersion    string       `json:"appVersion,omitempty" yaml:"appVersion,omitempty"`
	Deprecated    bool         `json:"deprecated,omitempty" yaml:"deprecated,omitempty"`
	CreatedAt     time.Time    `json:"createdAt" yaml:"-"`
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
