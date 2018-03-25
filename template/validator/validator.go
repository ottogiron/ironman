package validator

import "github.com/ironman-project/ironman/template/model"

//Validator validates  if a model is valid
type Validator interface {
	Validate(model *model.Template) (valid bool, errors []string, err error)
}
