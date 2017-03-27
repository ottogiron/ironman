package mapper

import "github.com/pkg/errors"
import "github.com/ironman-project/ironman/template/generator/metadata/field"

//ArrayMapper maps unstructured array to an internal representation
type ArrayMapper struct {
}

//Map maps a field to an array field
func (a *ArrayMapper) Map(f field.Field) (interface{}, error) {
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
	fieldDefinitionMap[field.FieldIDKey] = "placeholder"
	arrField, err := MapUnstructuredToField(fieldDefinitionMap)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to map %s form \n%v", fieldDefinitionKey, f)
	}
	fieldArr := field.NewArray(f, size, arrField)

	return fieldArr, nil
}
