package field

import (
	"github.com/pkg/errors"
)

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
	if err := validateMandatoryFieldProperty(f, fieldIDKey); err != nil {
		return err
	}
	if err := validateMandatoryFieldProperty(f, fieldTypeKey); err != nil {
		return err
	}
	if err := validateMandatoryFieldProperty(f, fieldLabelKey); err != nil {
		return err
	}
	return nil
}

func validateMandatoryFieldProperty(f Field, property string) error {
	if f[property] == nil {
		return errors.Errorf("%s is mandatory", property)
	}
	return nil
}

//MapUnstructuredToField maps and unmarshalls map[string]interface{} to an internal field definition
func MapUnstructuredToField(unstructuredField interface{}) (interface{}, error) {
	var ma map[string]interface{}
	var ok bool
	if ma, ok = unstructuredField.(map[string]interface{}); !ok {
		return nil, errors.Errorf("Can't map the defined field to an internal field definition \n%v", unstructuredField)
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
	case TypeArray:

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
			return nil, errors.Errorf("%s is mandatory for \n%v", sizeDefinitionKey, f)
		}

		var size int
		if size, ok = f[sizeDefinitionKey].(int); !ok {
			return nil, errors.Errorf("%s should be type int for \n%v ", sizeDefinitionKey, f)
		}

		if f[fieldDefinitionKey] == nil {
			return nil, errors.Errorf("%s is mandatory for \n%v", fieldDefinitionKey, f)
		}
		if fieldDefinitionMap, ok = f[fieldDefinitionKey].(map[string]interface{}); !ok {
			return nil, errors.Errorf("%s should be and object for \n%v", fieldDefinitionKey, f)
		}

		//This placeholder makes this fieldDefinitionMap to pass the ValidateMandatoryFieldAttributes validation
		fieldDefinitionMap[fieldIDKey] = "placeholder"
		arrField, err := MapUnstructuredToField(fieldDefinitionMap)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to map %s form \n%v", fieldDefinitionKey, f)
		}
		fieldArr := NewArray(f, size, arrField)

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

		if f[fieldsKey] == nil {
			return nil, errors.Errorf("%s is mandatory for \n%v", fieldsKey, f)
		}

		if list, ok = f[fieldsKey].([]interface{}); !ok {
			return nil, errors.Errorf("%s should be a list for \n%v", fieldsKey, f)
		}

		mappedFieldList := make([]interface{}, len(list))

		for i, fieldToMap := range list {
			mapped, err := MapUnstructuredToField(fieldToMap)
			if err != nil {
				return nil, errors.Wrapf(err, "Failed to map %s item \n%v for \n%v", TypeFixedList, fieldToMap, f)
			}
			mappedFieldList[i] = mapped
		}
		fieldList := NewFixedList(f, mappedFieldList)
		mappedField = fieldList

	default:
		return nil, errors.Errorf("Could not find right type mapping for \n%v", f)
	}

	return mappedField, nil
}
