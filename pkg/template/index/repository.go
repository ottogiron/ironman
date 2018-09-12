package index

import "github.com/ironman-project/ironman/pkg/template/model"

//index defines basic operations for a template model
type Index interface {
	Index(*model.Template) (string, error)
	Update(*model.Template) error
	Delete(templateID string) (bool, error)
	List() ([]*model.Template, error)
	FindTemplateByID(string) (*model.Template, error)
	Exists(string) (bool, error)
}
