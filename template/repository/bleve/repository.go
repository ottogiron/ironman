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

func (r *bleeveRepository) Update(template model.Template) error {
	query := bleve.NewMatchQuery(template.ID)
	search := bleve.NewSearchRequest(query)
	searchResults, err := r.index.Search(search)

	if err != nil {
		return errors.Wrapf(err, "Search for template failed id: %s")
	}

	if searchResults.Total != 1 {
		return errors.Errorf("Could not update template. %s not found. Hits: %d", template.ID, searchResults.Total)
	}

	id := searchResults.Hits[0].ID
	err = r.index.Index(id, template)

	if err != nil {
		return errors.Errorf("Failed to update template %s", template.ID)
	}

	return nil
}

func (r *bleeveRepository) Delete(ID string) error {
	panic("not implemented")
}

func (r *bleeveRepository) List() ([]model.Template, error) {
	panic("not implemented")
}
