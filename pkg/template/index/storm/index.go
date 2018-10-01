package storm

import (
	"time"

	"github.com/asdine/storm"
	"github.com/ironman-project/ironman/pkg/template/index"
	"github.com/ironman-project/ironman/pkg/template/model"
	"github.com/pkg/errors"
)

var _ index.Index = (*Index)(nil)

//DBFactory represents a *storm.DB factory
type DBFactory func() (*storm.DB, error)

func DefaultDBFactory(path string) DBFactory {
	return func() (*storm.DB, error) {
		return storm.Open(path)
	}
}

func New(dbFactory DBFactory) *Index {
	return &Index{
		dbFactory: dbFactory,
	}
}

type Index struct {
	dbFactory DBFactory
}

func (i *Index) Index(model *model.Template) (string, error) {
	db, err := i.dbFactory()
	if err != nil {
		return "", errors.Errorf("failed to index template %s %s", model.ID, err)
	}
	defer db.Close()
	model.CreatedAt = time.Now()
	err = db.Save(model)

	if err != nil {
		return "", errors.Errorf("failed to index template %s %s", model.ID, err)
	}
	return model.ID, nil
}

func (i *Index) Update(model *model.Template) error {
	db, err := i.dbFactory()
	if err != nil {
		return errors.Errorf("failed to update template %s %s", model.ID, err)
	}
	defer db.Close()

	err = db.Save(model)

	if err != nil {
		return errors.Errorf("failed to update template %s %s", model.ID, err)
	}

	return nil
}

func (i *Index) Delete(ID string) (bool, error) {
	db, err := i.dbFactory()
	if err != nil {
		return false, errors.Errorf("failed to delete template %s %s", ID, err)
	}
	defer db.Close()
	template := model.Template{ID: ID}

	err = db.DeleteStruct(&template)
	if err != nil {
		return false, errors.Errorf("faield to delete template %s %s", ID, err)
	}

	return true, nil

}

func (i *Index) List() ([]*model.Template, error) {
	db, err := i.dbFactory()
	if err != nil {
		return nil, errors.Errorf("failed to get list of templates %s", err)
	}
	defer db.Close()
	var templates []*model.Template
	err = db.All(&templates)
	if err != nil {
		return nil, errors.Errorf("failed to get list of templates %s", err)
	}
	return templates, nil
}

func (i *Index) FindTemplateByID(ID string) (*model.Template, error) {
	db, err := i.dbFactory()
	if err != nil {
		return nil, errors.Errorf("failed to find template by ID %s %s", ID, err)
	}
	defer db.Close()

	var template model.Template
	err = db.One("ID", ID, &template)
	if err != nil {
		return nil, errors.Errorf("failed to find template by ID %s %s", ID, err)
	}

	return &template, nil
}

func (i *Index) Exists(ID string) (bool, error) {

	if _, err := i.FindTemplateByID(ID); err != nil {
		return false, errors.Errorf("failed to verify if teplate exists %s %s", ID, err)
	}
	return true, nil
}
