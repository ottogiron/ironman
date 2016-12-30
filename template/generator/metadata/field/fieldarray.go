package field

//Array a fixed size array of  fields
type Array struct {
	Field
	Size            int
	FieldDefinition interface{}
}
