package loader

import (
	"github.com/ottogiron/ironman/template/generator/metadata"
	"github.com/ottogiron/ironman/template/generator/metadata/field"
	"github.com/ottogiron/ironman/template/unmarshall"
	"github.com/pkg/errors"
)

const (
	fieldDefinitionKey = "field_definition"
	fieldsKey          = "fields"
)

//Loader loads metadata from a  file
type Loader struct {
	fileUnmarshaller template.Unmarshaller
}

//New creates a new instance Loaders
func New(options ...Option) *Loader {
	y := &Loader{}

	for _, option := range options {
		option(y)
	}
	return y
}

//Load loads metadata from a  file
func (l *Loader) Load(bytes []byte) (*metadata.Metadata, error) {
	m := &metadata.Metadata{}
	err := l.fileUnmarshaller.Unmarshall(bytes, m)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to load generator file metadata")
	}
	return m, nil
}

//maps and unmarshalled map[string]interface{} to an internal field definition (from field package)
func mapUnstructuredToField(unstructuredField interface{}) (interface{}, error) {
	var ma map[string]interface{}
	var ok bool
	if ma, ok = unstructuredField.(map[string]interface{}); !ok {
		return nil, errors.Errorf("Can't map the defined field to an internal field definition %v", unstructuredField)
	}
	f := field.Field(ma)
	var mappedField interface{}
	switch field.Type(f.Type()) {

	case field.TypeText:

		//Example
		//id: myText
		//type: text
		//label: Enter some text
		//default: some default

		mappedField = field.Text{Field: f}
	case field.TypeFieldArray:

		//Example
		// id: myListOfThings
		// type: fieldarray
		// label: My List of things
		// size: 3
		// field_definition:
		// 	type: text
		//	label: Some label
		//  default: Something default

		fieldArr := field.Array{Field: f}
		var fieldDefinitionMap map[string]interface{}
		var ok bool

		if fieldDefinitionMap, ok = f[fieldDefinitionKey].(map[string]interface{}); !ok {
			return nil, errors.Errorf("Could not map mandatory field definition from array field %v", f)
		}
		arrField, err := mapUnstructuredToField(fieldDefinitionMap)
		if err != nil {
			return nil, errors.Errorf("Failed to map field definition for Array Field %v", f)
		}
		fieldArr.FieldDefinition = arrField
		mappedField = fieldArr
	case field.TypeFieldList:
		//Example
		//id: myListWithDifferntThings
		//type: fieldlist
		//label: My List of differnt things
		//fields:
		// 	- id: myThing1
		//	  type: text
		//    label: My Thing 1
		// 	- id: myThing2
		//	  type: text
		//    label: My Thing 2
		// 	- id: myThing3
		//	  type: text
		//    label: My Thing 3

		fieldList := field.List{Field: f}
		var list []interface{}
		var ok bool

		if list, ok = f[fieldsKey].([]interface{}); !ok {
			return nil, errors.Errorf("Could not map mandatory fields definition for %v", f)
		}
		mappedFieldList := make([]interface{}, len(list))

		for i, fieldToMap := range list {
			mapped, err := mapUnstructuredToField(fieldToMap)
			if err != nil {
				return nil, errors.Errorf("Could not map field for field list %v %v", f, fieldToMap)
			}
			mappedFieldList[i] = mapped
		}
		fieldList.Fields = mappedFieldList
		mappedField = fieldList

	default:
		return nil, errors.Errorf("Could not find right type mapping for %v", f)
	}

	return mappedField, nil
}
