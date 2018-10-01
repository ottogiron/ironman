package index

import "github.com/ironman-project/ironman/pkg/template/model"

//Index defines basic operations for a template model inside an index.
type Index interface {
	Index(model *model.Template) (string, error)
	Update(model *model.Template) error
	Delete(ID string) (bool, error)
	List() ([]*model.Template, error)
	FindTemplateByID(ID string) (*model.Template, error)
	Exists(ID string) (bool, error)
}
