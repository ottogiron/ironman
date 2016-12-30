package field

const (
	fieldIDKey    = "id"
	fieldTypeKey  = "type"
	fieldLabelKey = "label"
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
