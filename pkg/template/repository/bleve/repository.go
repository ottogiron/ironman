package bleve

import (
	"strconv"

	"github.com/blevesearch/bleve"
	"github.com/blevesearch/bleve/document"

	//Registers the bleve keyword analyzer
	_ "github.com/blevesearch/bleve/analysis/analyzer/keyword"
	"github.com/ironman-project/ironman/pkg/template/model"
	"github.com/ironman-project/ironman/pkg/template/repository"
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
		return "", errors.Wrapf(err, "failed to index template %s", template.ID)
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
		return errors.Errorf("failed to update template %s", template.ID)
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
		return nil, errors.Wrapf(err, "failed to get template document %s", ID)
	}

	t, err := deserialize(doc)

	if err != nil {
		return nil, err
	}

	return t, nil

}

func deserialize(doc *document.Document) (*model.Template, error) {
	template := &model.Template{}
	var generators []*model.Generator
	var mantainers []*model.Mantainer
	var sources []string
	for _, field := range doc.Fields {
		value := string(field.Value())
		switch field.Name() {
		case "iid":
			template.IID = value
		case "sourceType":
			template.SourceType = model.SourceType(value)
		case "id":
			template.ID = value
		case "version":
			template.Version = value
		case "name":
			template.Name = value
		case "description":
			template.Description = value
		case "directoryName":
			template.DirectoryName = value
		case "home":
			template.HomeURL = value
		case "sources":
			sources = append(sources, value)
		case "appVersion":
			template.AppVersion = value
		case "deprecated":
			template.Deprecated, _ = strconv.ParseBool(value)
		case "generators.id":
			generators = append(generators, &model.Generator{
				FileTypeOptions: model.FileTypeOptions{},
			})
			generators[field.ArrayPositions()[0]].ID = value
		case "generators.type":
			generators[field.ArrayPositions()[0]].TType = model.GeneratorType(value)
		case "generators.name":
			generators[field.ArrayPositions()[0]].Name = value
		case "generators.description":
			generators[field.ArrayPositions()[0]].Description = value
		case "generators.directoryName":
			generators[field.ArrayPositions()[0]].DirectoryName = value
		case "generators.fileTypeOptions.defaultTemplateFile":
			generators[field.ArrayPositions()[0]].FileTypeOptions.DefaultTemplateFile = value
		case "generators.fileTypeOptions.fileGenerationRelativePath":
			generators[field.ArrayPositions()[0]].FileTypeOptions.FileGenerationRelativePath = value
		case "mantainers.name":
			mantainers = append(mantainers, &model.Mantainer{})
			mantainers[field.ArrayPositions()[0]].Name = value
		case "mantainers.email":
			mantainers[field.ArrayPositions()[0]].Email = value
		case "mantainers.url":
			mantainers[field.ArrayPositions()[0]].URL = value

		default:
			return nil, errors.Errorf("Could not deserialize template from bleve document. Field %s must be explicitly processed", field.Name())
		}
	}
	template.Generators = generators
	template.Sources = sources
	template.Mantainers = mantainers
	return template, nil
}

func (r *bleeveRepository) Delete(ID string) (bool, error) {
	t, err := r.FindTemplateByID(ID)
	if err != nil {
		return false, errors.Wrapf(err, "failed to delete template %s", ID)
	}

	//It doesn't exists
	if t == nil {
		return false, nil
	}

	err = r.index.Delete(t.IID)

	if err != nil {
		return false, errors.Wrapf(err, "failed to delete template with id %s", ID)
	}

	return true, nil
}

func (r *bleeveRepository) List() ([]*model.Template, error) {
	var results []*model.Template
	query := bleve.NewMatchAllQuery()
	search := bleve.NewSearchRequest(query)

	searchResults, err := r.index.Search(search)

	if err != nil {
		return nil, errors.Wrap(err, "failed to list all the available templates")
	}

	if searchResults.Total > 0 {
		for _, result := range searchResults.Hits {
			doc, err := r.index.Document(result.ID)

			if err != nil {
				return nil, errors.Wrapf(err, "failed to retrieve template from document IID:%s", doc.ID)
			}

			templ, err := deserialize(doc)

			if err != nil {
				return nil, errors.Wrapf(err, "failed to deserialize template from document IID:%s", doc.ID)
			}
			results = append(results, templ)
		}
	}
	return results, nil
}

func (r *bleeveRepository) Exists(ID string) (bool, error) {
	templ, err := r.FindTemplateByID(ID)

	if err != nil {
		return false, errors.Wrapf(err, "failed to verify if template exists %s", ID)
	}

	if templ != nil {
		return true, nil
	}

	return false, nil
}
