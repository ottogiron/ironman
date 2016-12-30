package field

//FixedArray a fixed size array of  fields
type FixedArray struct {
	Field
	size            int
	fieldDefinition interface{}
}

//NewFixedArray returns a new initialized array field
func NewFixedArray(f Field, size int, fieldDefinition interface{}) *FixedArray {
	fieldArr := &FixedArray{Field: f, size: size, fieldDefinition: fieldDefinition}
	return fieldArr
}
