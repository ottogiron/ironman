package mapper

import (
	"github.com/ironman-project/ironman/template/generator/metadata/field"
)

//TextMapper a text field mapper
type TextMapper struct {
}

//Map maps an unstructerd text  to an internal text representation
func (t *TextMapper) Map(f field.Field) (interface{}, error) {

	//Example
	//id: myText
	//type: text
	//label: Enter some text
	//default: some default

	return field.NewText(f), nil
}
