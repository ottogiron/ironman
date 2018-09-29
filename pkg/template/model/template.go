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
	//IID internal database ID
	IID        string     `json:"iid,omitempty" yaml:"iid,omitempty"`
	SourceType SourceType `json:"sourceType,omitempty" yaml:"sourceType,omitempty"`
	ID         string     `json:"id" yaml:"id" storm:"id"` //contains an special storm annotation

	Version       string       `json:"version" yaml:"version"`
	Name          string       `json:"name" yaml:"name"`
	Description   string       `json:"description" yaml:"description"`
	Generators    []*Generator `json:"generators" yaml:"generators"`
	DirectoryName string       `json:"directoryName,omitempty" yaml:"directoryName,omitempty"`
	HomeURL       string       `json:"home" yaml:"home"`
	Sources       []string     `json:"sources" yaml:"sources"`
	Mantainers    []*Mantainer `json:"mantainers" yaml:"mantainers"`
	AppVersion    string       `json:"appVersion" yaml:"appVersion"`
	Deprecated    bool         `json:"deprecated" yaml:"deprecated"`
	CreatedAt     time.Time    `json:"created_at" yaml:"created_at"`
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
