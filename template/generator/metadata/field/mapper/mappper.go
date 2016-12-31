package mapper

import (
	"github.com/ottogiron/ironman/template/generator/metadata/field"
	"github.com/pkg/errors"
)

const (
	fieldDefinitionKey = "field_definition"
	fieldsKey          = "fields"
	sizeDefinitionKey  = "size"
)

//Mapper defines a mapper from and unstructured field to a internal field definition
type Mapper interface {
	Map(f field.Field) (interface{}, error)
}

//New returns a nes Mapper based on a field type
func New(fieldType field.Type) (Mapper, error) {
	switch fieldType {
	case field.TypeText:
		return &TextMapper{}, nil
	case field.TypeArray:
		return &ArrayMapper{}, nil
	case field.TypeFixedList:
		return &FixedListMapper{}, nil
	default:
		return nil, errors.Errorf("mapper for %s is not implemented", fieldType)
	}
}

//MapUnstructuredToField maps and unmarshalls map[string]interface{} to an internal field definition
func MapUnstructuredToField(unstructuredField interface{}) (interface{}, error) {
	var ma map[string]interface{}
	var ok bool
	if ma, ok = unstructuredField.(map[string]interface{}); !ok {
		return nil, errors.Errorf("Can't map the defined field to an internal field definition \n%v", unstructuredField)
	}

	f := field.Field(ma)

	if err := field.ValidateMandatoryFieldAttributes(f); err != nil {
		return nil, errors.Wrap(err, "Failed to map field ")
	}

	mapper, err := New(f.Type())

	if err != nil {
		return nil, errors.Wrapf(err, "Failed to create mapper for field type %s", f.Type())
	}

	return mapper.Map(f)
}
