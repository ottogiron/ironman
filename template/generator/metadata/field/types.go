package field

//Type defines the type of a field
type Type string

const (
	//TypeFixedArray represents a fixed field type array
	TypeFixedArray Type = "fixedarray"
	//TypeFixedList represents a fixed list of fields
	TypeFixedList Type = "fixedlist"
	//TypeText represents a text input field
	TypeText Type = "text"
)
