package yaml

import (
	"io/ioutil"

	"github.com/ottogiron/ironman/template/generator/metadata"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

//New returns a new yaml unmarshaller
func New() *Unmarshaller {
	return &Unmarshaller{}
}

//Unmarshaller defines a  unmarshaller
type Unmarshaller struct {
}

//Unmarshall unmarshall a yaml file from a file
func (u *Unmarshaller) Unmarshall(path string, out interface{}) error {
	b, err := ioutil.ReadFile(path)

	if err != nil {
		return errors.Wrapf(err, "Failed to load yaml file %s", path)
	}

	m := &metadata.Metadata{}

	err = yaml.Unmarshal(b, m)

	if err != nil {
		return errors.Wrapf(err, "Failed to unmarshall yaml file %s", path)
	}

	return nil
}
