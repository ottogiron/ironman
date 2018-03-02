package repository

import "github.com/ironman-project/ironman/template/model"

//Repository defines basic operations for a template model
type Repository interface {
	Index(model.Template) error
	Update(model.Template) error
	Delete(ID string) error
	List() ([]model.Template, error)
}
