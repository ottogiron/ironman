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
	return f[fieldIDKey].(string)
}

//Label returns the label of a base field
func (f Field) Label() string {
	return f[fieldLabelKey].(string)
}

//Type returns the type of the field
func (f Field) Type() string {
	return f[fieldTypeKey].(string)
}
