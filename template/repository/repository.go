package repository

import "github.com/ironman-project/ironman/template/model"

//Repository defines basic operations for a template model
type Repository interface {
	Index(model.Template) (string, error)
	Update(model.Template) error
	Delete(templateID string) error
	List() ([]model.Template, error)
}
