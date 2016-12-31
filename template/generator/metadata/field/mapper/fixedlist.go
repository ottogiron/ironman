package mapper

import (
	"github.com/ottogiron/ironman/template/generator/metadata/field"
	"github.com/pkg/errors"
)

//FixedListMapper maps an unstructured fixedlist to an internal representation
func FixedListMapper(f field.Field) (interface{}, error) {
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
			return nil, errors.Wrapf(err, "Failed to map %s item \n%v for \n%v", field.TypeFixedList, fieldToMap, f)
		}
		mappedFieldList[i] = mapped
	}
	fieldList := field.NewFixedList(f, mappedFieldList)
	return fieldList, nil
}
