package metadata

//ReaderType defines the different types of metadata readers
type ReaderType string

const (
	//Yaml Reader type yaml
	Yaml = "yaml"
	//JSON reader type json
	JSON = "json"
	//Toml reader type toml
	Toml = "toml"
)

//Reader  template metadata reader
type Reader interface {
	Read() (*Template, error)
}

//NewReader returns a new reader based on the type. Defaults to yaml
func NewReader(path string, t ReaderType) Reader {
	var reader Reader
	switch t {
	default:
		reader = &yamlReader{}
	}
	return reader
}

type yamlReader struct {
}

func (yr *yamlReader) Read() (*Template, error) {
	return nil, nil
}
