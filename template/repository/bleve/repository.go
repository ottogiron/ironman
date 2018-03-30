package bleve

import (
	"github.com/blevesearch/bleve"
	"github.com/blevesearch/bleve/document"
	//Registers the bleve keyword analyzer
	_ "github.com/blevesearch/bleve/analysis/analyzer/keyword"
	"github.com/ironman-project/ironman/template/model"
	"github.com/ironman-project/ironman/template/repository"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
)

const (
	idKey = "id"
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

//BuildIndex builds a nex index based on a path
func BuildIndex(path string) (bleve.Index, error) {
	indexMapping := bleve.NewIndexMapping()

	templateIDMapping := bleve.NewTextFieldMapping()
	templateIDMapping.Analyzer = "keyword"

	templateDocMapping := bleve.NewDocumentMapping()

	generatorsMapping := bleve.NewDocumentMapping()

	templateDocMapping.AddFieldMappingsAt(idKey, templateIDMapping)
	templateDocMapping.AddSubDocumentMapping("model.generator", generatorsMapping)

	indexMapping.AddDocumentMapping("model.template", templateDocMapping)

	index, err := bleve.New(path, indexMapping)

	if err != nil {
		return nil, err
	}

	return index, nil
}

func (r *bleeveRepository) Index(template *model.Template) (string, error) {
	id := uuid.NewV4()
	template.IID = id.String()
	err := r.index.Index(id.String(), template)
	if err != nil {
		return "", errors.Wrapf(err, "Failed to index template %s", template.ID)
	}
	return id.String(), nil
}

func (r *bleeveRepository) Update(template *model.Template) error {
	query := bleve.NewTermQuery(template.ID)
	query.SetField(idKey)
	search := bleve.NewSearchRequest(query)
	searchResults, err := r.index.Search(search)

	if err != nil {
		return errors.Wrapf(err, "Search for template failed id: %s")
	}

	if searchResults.Total != 1 {
		return errors.Errorf("Could not update template. %s not found. Hits: %d", template.ID, searchResults.Total)
	}

	id := searchResults.Hits[0].ID
	template.IID = id
	err = r.index.Index(id, template)

	if err != nil {
		return errors.Errorf("Failed to update template %s", template.ID)
	}

	return nil
}

func (r *bleeveRepository) FindTemplateByID(ID string) (*model.Template, error) {
	query := bleve.NewTermQuery(ID)
	query.SetField(idKey)
	search := bleve.NewSearchRequest(query)

	searchResults, err := r.index.Search(search)

	if err != nil {
		return nil, errors.Wrapf(err, "Search for template failed id: %s")
	}

	if searchResults.Total != 1 {
		return nil, nil
	}

	match := searchResults.Hits[0]
	doc, err := r.index.Document(match.ID)

	if err != nil {
		return nil, errors.Wrapf(err, "Failed to get template document %s", ID)
	}

	t, err := deserialize(doc)

	if err != nil {
		return nil, err
	}

	return t, nil

}

func deserialize(doc *document.Document) (*model.Template, error) {
	template := &model.Template{}
	var currGenerator *model.Generator
	var generators []*model.Generator
	for _, field := range doc.Fields {
		value := string(field.Value())
		switch field.Name() {
		case "iid":
			template.IID = value
		case "id":
			template.ID = value
		case "version":
			template.Version = value
		case "name":
			template.Name = value
		case "description":
			template.Description = value
		case "directory_name":
			template.DirectoryName = value
		case "generators.id":
			currGenerator = &model.Generator{}
			currGenerator.ID = value
		case "generators.name":
			currGenerator.Name = value
		case "generators.description":
			currGenerator.Description = value
		case "generators.directory_name":
			currGenerator.DirectoryName = value
			generators = append(generators, currGenerator)
		default:
			return nil, errors.Errorf("Could not deserialize template from bleve document. Field %s must be explicitly processed", field.Name())
		}
	}
	template.Generators = generators
	return template, nil
}

func (r *bleeveRepository) Delete(ID string) (bool, error) {
	t, err := r.FindTemplateByID(ID)
	if err != nil {
		return false, errors.Wrapf(err, "Failed to delete template %s", ID)
	}

	//It doesn't exists
	if t == nil {
		return false, nil
	}

	err = r.index.Delete(t.IID)

	if err != nil {
		return false, errors.Wrapf(err, "Failed to delete template with id %s", ID)
	}

	return true, nil
}

func (r *bleeveRepository) List() ([]*model.Template, error) {
	var results []*model.Template
	query := bleve.NewMatchAllQuery()
	search := bleve.NewSearchRequest(query)

	searchResults, err := r.index.Search(search)

	if err != nil {
		return nil, errors.Wrap(err, "Failed to list all the available templates")
	}

	if searchResults.Total > 0 {
		for _, result := range searchResults.Hits {
			doc, err := r.index.Document(result.ID)

			if err != nil {
				return nil, errors.Wrapf(err, "Failed to retrieve template from document IID:%s", doc.ID)
			}

			templ, err := deserialize(doc)

			if err != nil {
				return nil, errors.Wrapf(err, "Failed to deserialize template from document IID:%s", doc.ID)
			}
			results = append(results, templ)
		}
	}
	return results, nil
}

func (r *bleeveRepository) Exists(ID string) (bool, error) {
	templ, err := r.FindTemplateByID(ID)

	if err != nil {
		return false, errors.Wrapf(err, "Failed to verify if template exists %s", ID)
	}

	if templ != nil {
		return true, nil
	}

	return false, nil
}
