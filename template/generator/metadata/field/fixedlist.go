package field

import "github.com/ironman-project/ironman/text/yaml"

//FixedList defines a list with arbitrary text fields
type FixedList struct {
	Field
	Fields []interface{}
}

//NewFixedList returns a new fixed list of fields
func NewFixedList(f Field, fields []interface{}) *FixedList {
	return &FixedList{Field: f, Fields: fields}
}

func (l *FixedList) String() string {
	return yaml.Print(l)
}
