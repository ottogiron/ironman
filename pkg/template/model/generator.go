package model

//FileTypeOptions  options for file type generator
type FileTypeOptions struct {
	DefaultTemplateFile        string `json:"defaultTemplateFile,omitempty" yaml:"defaultTemplateFile,omitempty"`
	FileGenerationRelativePath string `json:"fileGenerationRelativePath,omitempty" yaml:"fileGenerationRelativePath,omitempty"`
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
	ID              string          `json:"id" yaml:"id"`
	TType           GeneratorType   `json:"type" yaml:"type"`
	Name            string          `json:"name" yaml:"name"`
	Description     string          `json:"description" yaml:"description"`
	DirectoryName   string          `json:"-" yaml:"-"`
	FileTypeOptions FileTypeOptions `json:"-" yaml:"-"`
	Hooks           *GeneratorHooks `json:"hooks,omitempty" yaml:"hooks,omitempty"`
}

//Type Simple type serialization for generator model
func (g *Generator) Type() string {
	return "model.generator"
}

type GeneratorHooks struct {
	PreGenerate  []*Command `json:"postGenerate,omitempty" yaml:"pre_generate,omitempty"`
	PostGenerate []*Command `json:"preGenerate,omitempty" yaml:"post_generate,omitempty"`
}
