package mapper

import (
	"github.com/ottogiron/ironman/template/generator/metadata/field"
)

//TextMapper maps an unstructerd text  to an internal text representation
func TextMapper(f field.Field) (interface{}, error) {

	//Example
	//id: myText
	//type: text
	//label: Enter some text
	//default: some default

	return field.NewText(f), nil
}
