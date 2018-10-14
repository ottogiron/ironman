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
func NewFSReader(ignoreFiles []string, fileExtension MetadataFileExtension, decoder Decoder, generatorsPath string) Reader {
	reader := &fsReader{
		fileExtension,
		decoder,
		ignoreFiles,
		generatorsPath,
	}

	return reader
}

type fsReader struct {
	fileExtension  MetadataFileExtension
	decoder        Decoder
	ignoreFiles    []string
	generatorsPath string
}

func (r *fsReader) Read(path string) (*Template, error) {
	rootIronmanMetadataPath := filepath.Join(path, meatadataFileName+"."+string(r.fileExtension))
	rootIronmanTemplateFile, err := os.Open(rootIronmanMetadataPath)

	if err != nil {
		if os.IsNotExist(err) {
			return nil, errors.Wrap(err, rootIronmanMetadataPath)
		}

		return nil, errors.Wrapf(err, "failed to read metadata file %s", rootIronmanMetadataPath)
	}
	defer rootIronmanTemplateFile.Close()

	var templateModel Template
	err = r.decoder.Decode(&templateModel, rootIronmanTemplateFile)

	if err != nil {
		return nil, errors.Wrapf(err, "failed to decode template information from %s", rootIronmanMetadataPath)
	}

	absolutePath, err := filepath.Abs(path)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to get absolute path from template path %s", path)
	}
	templateModel.DirectoryName = filepath.Base(absolutePath)
	generatorsPath := filepath.Join(path, r.generatorsPath)
	generatorFiles, err := ioutil.ReadDir(generatorsPath)

	if err != nil {
		return nil, errors.Wrapf(err, "failed to read available generators for %s", path)
	}

	for _, generatorFile := range generatorFiles {
		if generatorFile.IsDir() && !r.ignore(generatorFile.Name()) {
			generatorMetadataPath := filepath.Join(generatorsPath, generatorFile.Name(), meatadataFileName+"."+string(r.fileExtension))
			generatorMetadataFile, err := os.Open(generatorMetadataPath)
			if err != nil {
				if os.IsNotExist(err) {
					return nil, errors.Wrap(err, generatorMetadataPath)
				}

				return nil, errors.Wrapf(err, "failed to read metadata file %s", rootIronmanMetadataPath)
			}
			defer generatorMetadataFile.Close()
			var generatorModel Generator
			err = r.decoder.Decode(&generatorModel, generatorMetadataFile)
			if err != nil {
				return nil, errors.Wrapf(err, "failed to decode generator information from %s", generatorMetadataPath)
			}
			generatorModel.DirectoryName = generatorFile.Name()
			//Make the generator id optional. Use the directory ID if the .ironman.yaml doesn't contain an ID
			if generatorModel.ID == "" {
				generatorModel.ID = generatorFile.Name()
			}

			if string(generatorModel.TType) == "" {
				generatorModel.TType = GeneratorTypeDirectory
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
