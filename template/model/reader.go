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

//MetadataFileExtension metadata file extension type
type MetadataFileExtension string

const (
	//MetadataFileExtensionYAML file extension for yaml metadata files
	MetadataFileExtensionYAML = "yaml"
)

//Reader  template metadata reader
type Reader interface {
	Read(location string) (*Template, error)
}

//NewFSReader returns a new reader based on the type.
func NewFSReader(ignoreFiles []string, fileExtension MetadataFileExtension, decoder Decoder) Reader {
	reader := &fsReader{
		fileExtension,
		decoder,
		ignoreFiles,
	}

	return reader
}

type fsReader struct {
	fileExtension MetadataFileExtension
	decoder       Decoder
	ignoreFiles   []string
}

func (r *fsReader) Read(path string) (*Template, error) {
	rootIronmanMetadataPath := filepath.Join(path, meatadataFileName+"."+string(r.fileExtension))
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

	generatorFiles, err := ioutil.ReadDir(path)

	if err != nil {
		return nil, errors.Wrapf(err, "Failed to read available generators for %s", path)
	}

	for _, generatorFile := range generatorFiles {
		if generatorFile.IsDir() && !r.ignore(generatorFile.Name()) {
			generatorMetadataPath := filepath.Join(path, generatorFile.Name(), meatadataFileName+"."+string(r.fileExtension))
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

func (r *fsReader) ignore(fileName string) bool {
	for _, ignore := range r.ignoreFiles {
		if ignore == fileName {
			return true
		}
	}
	return false
}
