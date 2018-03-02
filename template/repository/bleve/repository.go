package bleve

import (
	"github.com/blevesearch/bleve"
	"github.com/ironman-project/ironman/template/model"
	"github.com/ironman-project/ironman/template/repository"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
)

var _ repository.Repository = (*bleeveRepository)(nil)

type bleeveRepository struct {
	index bleve.Index
}

//New creates a new instance of a bleeve repository
func New(options ...Option) repository.Repository {
	r := &bleeveRepository{}

	for _, option := range options {
		option(r)
	}
	return r
}

func (r *bleeveRepository) Index(template model.Template) (string, error) {
	id := uuid.NewV4()
	err := r.index.Index(id.String(), template)
	if err != nil {
		return "", errors.Wrapf(err, "Failed to index template %s", template.ID)
	}
	return id.String(), nil
}

func (r *bleeveRepository) Update(model.Template) error {
	panic("not implemented")
}

func (r *bleeveRepository) Delete(ID string) error {
	panic("not implemented")
}

func (r *bleeveRepository) List() ([]model.Template, error) {
	panic("not implemented")
}
