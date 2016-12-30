package field

//Type defines the type of a field
type Type string

const (
	//TypeFieldArray represents the type of a field array
	TypeFieldArray Type = "fieldarray"
	//TypeFieldList represents the type of a field list
	TypeFieldList Type = "fieldlist"
	//TypeText represents a text input field
	TypeText Type = "text"
)
