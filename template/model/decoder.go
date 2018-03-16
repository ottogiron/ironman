package model

import (
	"io"
	"io/ioutil"

	"github.com/pkg/errors"
	yaml "gopkg.in/yaml.v2"
)

//DecoderType defines the different types of metadata decoders
type DecoderType string

const (
	//Yaml decoder type yaml
	Yaml = "yaml"
	//JSON decoder type json
	JSON = "json"
	//Toml decoder type toml
	Toml = "toml"
)

//Decoder  template metadata reader
type Decoder interface {
	Decode(model interface{}, reader io.Reader) error
}

//NewDecoder returns a new decoder based on the type. Defaults to yaml
func NewDecoder(t DecoderType) Decoder {
	var decoder Decoder
	switch t {
	default:
		decoder = &yamlDecoder{}
	}
	return decoder
}

type yamlDecoder struct {
}

func (yr *yamlDecoder) Decode(model interface{}, reader io.Reader) error {
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return errors.Wrap(err, "Failed to decode template model")
	}
	err = yaml.Unmarshal(data, model)
	if err != nil {
		return errors.Wrap(err, "Failed to decode template model")
	}
	return nil
}
