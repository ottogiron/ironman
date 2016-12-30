package loader

import (
	"github.com/ottogiron/ironman/template/generator/metadata"
	"github.com/ottogiron/ironman/template/unmarshall"
	"github.com/pkg/errors"
)

//Loader loads metadata from a  file
type Loader struct {
	fileUnmarshaller template.Unmarshaller
}

//New creates a new instance Loaders
func New(options ...Option) *Loader {
	y := &Loader{}

	for _, option := range options {
		option(y)
	}
	return y
}

//Load loads metadata from a  file
func (l *Loader) Load(bytes []byte) (*metadata.Metadata, error) {
	m := &metadata.Metadata{}
	err := l.fileUnmarshaller.Unmarshall(bytes, m)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to load generator file metadata")
	}
	return m, nil
}
