package field

import (
	"github.com/pkg/errors"
)

const (
	//FieldIDKey key of the id on a field
	FieldIDKey = "id"
	//FieldTypeKey key of the type on a field
	FieldTypeKey = "type"
	//FieldLabelKey key of the label on a field
	FieldLabelKey = "label"
)

//Field represents a field which implements basic input methods
type Field map[string]interface{}

//ID returns de id of a base field
func (f Field) ID() string {
	return f.stringValue(FieldIDKey)
}

//Label returns the label of a base field
func (f Field) Label() string {
	return f.stringValue(FieldLabelKey)
}

//Type returns the type of the field
func (f Field) Type() Type {
	return Type(f.stringValue(FieldTypeKey))
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
	if err := validateMandatoryFieldProperty(f, FieldIDKey); err != nil {
		return err
	}
	if err := validateMandatoryFieldProperty(f, FieldTypeKey); err != nil {
		return err
	}
	if err := validateMandatoryFieldProperty(f, FieldLabelKey); err != nil {
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
