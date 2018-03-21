package ironman

import (
	"log"
	"path/filepath"

	"github.com/blevesearch/bleve"
	"github.com/ironman-project/ironman/template/manager"
	"github.com/ironman-project/ironman/template/manager/git"
	"github.com/ironman-project/ironman/template/model"
	"github.com/ironman-project/ironman/template/repository"
	brepository "github.com/ironman-project/ironman/template/repository/bleve"
)

const (
	indexName = "templates.index"
)

//Ironman is the one administering the local
type Ironman struct {
	manager     manager.Manager
	modelReader model.Reader
	repository  repository.Repository
	home        string
}

//New returns a new instance of ironman
func New(home string, options ...Option) *Ironman {
	ir := &Ironman{}
	for _, option := range options {
		option(ir)
	}
	if ir.manager == nil {
		manager := git.New(home)
		ir.manager = manager
	}

	if ir.repository == nil {
		indexPath := filepath.Join(home, indexName)
		index, err := buildIndex(indexPath)
		if err != nil {
			log.Fatal("Failed to create ironman templates index", err)
		}
		ir.repository = brepository.New(
			brepository.SetIndex(index),
		)
	}

	if ir.modelReader == nil {
		decoder := model.NewDecoder(model.DecoderTypeYAML)
		modelReader := model.NewFSReader([]string{".git"}, model.MetadataFileExtensionYAML, decoder)
		ir.modelReader = modelReader
	}
	return ir
}

func buildIndex(path string) (bleve.Index, error) {
	// open the index
	index, err := bleve.Open(path)
	if err == bleve.ErrorIndexPathDoesNotExist {
		index, err = brepository.BuildIndex(path)
		if err != nil {
			return nil, err
		}
	} else if err != nil {
		return nil, err
	}
	return index, nil
}

//Install installs a new template based on a template locator
func (i *Ironman) Install(templateLocator string) error {
	ID, err := i.manager.Install(templateLocator)
	if err != nil {
		return err
	}

	templatePath := i.manager.TemplateLocation(ID)
	model, err := i.modelReader.Read(templatePath)

	if err != nil {
		return err
	}

	_, err = i.repository.Index(model)

	if err != nil {
		//rollback manager installation
		_ = i.manager.Uninstall(ID)
		return err
	}

	return nil
}

//Link Creates a symlink to the ironman repository from any path in the filesystem
func (i *Ironman) Link(templatePath, templateID string) error {
	err := i.manager.Link(templatePath, templateID)

	if err != nil {
		return err
	}

	return nil
}

//List returns a list of all the installed ironman templates
func (i *Ironman) List() ([]*model.Template, error) {
	results, err := i.repository.List()
	if err != nil {
		return nil, err
	}

	return results, nil
}

//Uninstall uninstalls an ironman template
func (i *Ironman) Uninstall(templateID string) error {
	err := i.manager.Uninstall(templateID)
	if err != nil {
		return err
	}
	return nil
}

//Unlink unlinks a previously linked ironman template
func (i *Ironman) Unlink(templateID string) error {
	err := i.manager.Unlink(templateID)
	if err != nil {
		return err
	}
	return nil
}

//Update updates an iroman template
func (i *Ironman) Update(templateID string) error {
	err := i.manager.Update(templateID)
	if err != nil {
		return err
	}
	return nil
}
