package field

//Array a fixed size array of  fields
type Array struct {
	Field
	size            int
	fieldDefinition interface{}
}

//NewArray returns a new initialized array field
func NewArray(f Field, size int, fieldDefinition interface{}) *Array {
	fieldArr := &Array{Field: f, size: size, fieldDefinition: fieldDefinition}
	return fieldArr
}
