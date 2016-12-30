package loader

import "github.com/ottogiron/ironman/template/generator/metadata"
import "io/ioutil"
import "github.com/pkg/errors"
import "gopkg.in/yaml.v2"

//YAMLLoader loads metadata from a Yaml file
type YAMLLoader struct {
	fieldMapper FieldMapper
	path        string
}

//NewYAMLLoader creates a new instance YAMLLoaders
func NewYAMLLoader(filePath string, fieldMapper FieldMapper) *YAMLLoader {
	y := YAMLLoader{
		fieldMapper: fieldMapper,
		path:        filePath,
	}
	return &y
}

//Load loads metadata from a Yaml file
func (y *YAMLLoader) Load() (*metadata.Metadata, error) {
	b, err := ioutil.ReadFile(y.path)

	if err != nil {
		return nil, errors.Wrapf(err, "Failed to load yaml file %s", y.path)
	}

	m := &metadata.Metadata{}

	err = yaml.Unmarshal(b, m)

	if err != nil {
		return nil, errors.Wrapf(err, "Failed to unmarshall yaml file %s", y.path)
	}
	return nil, nil
}
