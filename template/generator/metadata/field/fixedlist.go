package field

//FixedList defines a list with arbitrary text fields
type FixedList struct {
	Field
	fields []interface{}
}

//NewFixedList returns a new fixed list of fields
func NewFixedList(f Field, fields []interface{}) *FixedList {
	return &FixedList{Field: f, fields: fields}
}
