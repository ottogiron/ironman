package yaml

import (
	"github.com/ottogiron/ironman/template/generator/metadata"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

//NewUnmarshaller returns a new yaml unmarshaller
func NewUnmarshaller() *Unmarshaller {
	return &Unmarshaller{}
}

//Unmarshaller defines a  unmarshaller
type Unmarshaller struct {
}

//Unmarshall unmarshall a yaml file from a file
func (u *Unmarshaller) Unmarshall(bytes []byte, out interface{}) error {

	m := &metadata.Metadata{}

	err := yaml.Unmarshal(bytes, m)

	if err != nil {
		return errors.Wrap(err, "Failed to unmarshall yaml  metadata")
	}

	return nil
}
