package model

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
)

const (
	meatadataFileName = ".ironman"
)

//Reader  template metadata reader
type Reader interface {
	Read() (*Template, error)
}

//NewFSReader returns a new reader based on the type. Defaults to yaml
func NewFSReader(path string, fileExtension string, decoder Decoder) Reader {
	reader := &fsReader{
		path,
		fileExtension,
		decoder,
	}

	return reader
}

type fsReader struct {
	path          string
	fileExtension string
	decoder       Decoder
}

func (r *fsReader) Read() (*Template, error) {
	rootIronmanMetadataPath := filepath.Join(r.path, meatadataFileName+"."+r.fileExtension)
	rootIronmanTemplateFile, err := os.Open(rootIronmanMetadataPath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, errors.Wrap(err, rootIronmanMetadataPath)
		}

		return nil, errors.Wrapf(err, "Failed to read metadata file %s", rootIronmanMetadataPath)
	}
	var templateModel Template
	err = r.decoder.Decode(&templateModel, rootIronmanTemplateFile)

	if err != nil {
		return nil, errors.Wrapf(err, "Failed to decode template information from %s", rootIronmanMetadataPath)
	}

	generatorFiles, err := ioutil.ReadDir(r.path)

	if err != nil {
		return nil, errors.Wrapf(err, "Failed to read available generators for %s", r.path)
	}

	for _, generatorFile := range generatorFiles {
		if generatorFile.IsDir() {
			generatorMetadataPath := filepath.Join(r.path, generatorFile.Name(), meatadataFileName+"."+r.fileExtension)
			generatorMetadataFile, err := os.Open(generatorMetadataPath)
			if err != nil {
				if os.IsNotExist(err) {
					return nil, errors.Wrap(err, generatorMetadataPath)
				}

				return nil, errors.Wrapf(err, "Failed to read metadata file %s", rootIronmanMetadataPath)
			}
			var generatorModel Generator
			err = r.decoder.Decode(&generatorModel, generatorMetadataFile)
			if err != nil {
				return nil, errors.Wrapf(err, "Failed to decode generator information from %s", generatorMetadataPath)
			}
			templateModel.Generators = append(templateModel.Generators, &generatorModel)
		}
	}

	return &templateModel, nil
}
