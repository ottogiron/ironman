package field

//Type defines the type of a field
type Type string

const (
	//TypeArray represents a fixed field type array
	TypeArray Type = "array"
	//TypeFixedList represents a fixed list of fields
	TypeFixedList Type = "fixedlist"
	//TypeText represents a text input field
	TypeText Type = "text"
)
