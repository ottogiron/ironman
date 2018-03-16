package model

import (
	"os"
	"path/filepath"

	"github.com/pkg/errors"
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
	rootIronmanMetadataPath := filepath.Join(r.path, ".ironman."+r.fileExtension)
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
	}

	return &templateModel, nil
}
