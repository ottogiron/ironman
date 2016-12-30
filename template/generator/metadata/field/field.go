package field

import "github.com/pkg/errors"

const (
	fieldIDKey    = "id"
	fieldTypeKey  = "type"
	fieldLabelKey = "label"
)

const (
	fieldDefinitionKey = "field_definition"
	fieldsKey          = "fields"
	sizeDefinitionKey  = "size"
)

//Field represents a field which implements basic input methods
type Field map[string]interface{}

//ID returns de id of a base field
func (f Field) ID() string {
	return f.stringValue(fieldIDKey)
}

//Label returns the label of a base field
func (f Field) Label() string {
	return f.stringValue(fieldLabelKey)
}

//Type returns the type of the field
func (f Field) Type() string {
	return f.stringValue(fieldTypeKey)
}

func (f Field) stringValue(fieldName string) string {
	if f[fieldName] == nil {
		return ""
	}
	var val string
	var ok bool
	if val, ok = f[fieldName].(string); !ok {
		return ""
	}
	return val
}

//ValidateMandatoryFieldAttributes validates if all the mandatory attributes of a field are present
func ValidateMandatoryFieldAttributes(f Field) error {
	if f["id"] == nil || f["type"] == nil || f["label"] == nil {
		return errors.Errorf("%s, %s, and %s are mandatory", fieldIDKey, fieldTypeKey, fieldLabelKey)
	}
	return nil
}

//MapUnstructuredToField maps and unmarshalls map[string]interface{} to an internal field definition
func MapUnstructuredToField(unstructuredField interface{}) (interface{}, error) {
	var ma map[string]interface{}
	var ok bool
	if ma, ok = unstructuredField.(map[string]interface{}); !ok {
		return nil, errors.Errorf("Can't map the defined field to an internal field definition %v", unstructuredField)
	}

	f := Field(ma)

	if err := ValidateMandatoryFieldAttributes(f); err != nil {
		return nil, errors.Wrap(err, "Failed to map field ")
	}
	var mappedField interface{}
	switch Type(f.Type()) {

	case TypeText:

		//Example
		//id: myText
		//type: text
		//label: Enter some text
		//default: some default

		mappedField = NewText(f)
	case TypeFixedArray:

		//Example
		// id: myListOfThings
		// type: fieldarray
		// label: My List of things
		// size: 3
		// field_definition:
		// 	type: text
		//	label: Some label
		//  default: Something default

		var fieldDefinitionMap map[string]interface{}
		var ok bool

		if f[sizeDefinitionKey] == nil {
			return nil, errors.Errorf("%s is mandatory for %v", sizeDefinitionKey, f)
		}

		var size int
		if size, ok = f[sizeDefinitionKey].(int); !ok {
			return nil, errors.Errorf("Could not map mandatory %s from array field %v ", sizeDefinitionKey, f)
		}

		if f[fieldDefinitionKey] == nil {
			return nil, errors.Errorf("%s is mandatory for %v", fieldDefinitionKey, f)
		}
		if fieldDefinitionMap, ok = f[fieldDefinitionKey].(map[string]interface{}); !ok {
			return nil, errors.Errorf("Could not map mandatory %s from array field %v", fieldDefinitionKey, f)
		}
		arrField, err := MapUnstructuredToField(fieldDefinitionMap)
		if err != nil {
			return nil, errors.Errorf("Failed to map field definition for Array Field %v", f)
		}
		fieldArr := NewFixedArray(f, size, arrField)

		mappedField = fieldArr
	case TypeFixedList:
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

		var list []interface{}
		var ok bool

		if list, ok = f[fieldsKey].([]interface{}); !ok {
			return nil, errors.Errorf("Could not map mandatory fields definition for %v", f)
		}
		mappedFieldList := make([]interface{}, len(list))

		for i, fieldToMap := range list {
			mapped, err := MapUnstructuredToField(fieldToMap)
			if err != nil {
				return nil, errors.Errorf("Could not map field for field list %v %v", f, fieldToMap)
			}
			mappedFieldList[i] = mapped
		}

		fieldList := NewFixedList(f, mappedFieldList)
		mappedField = fieldList

	default:
		return nil, errors.Errorf("Could not find right type mapping for %v", f)
	}

	return mappedField, nil
}
