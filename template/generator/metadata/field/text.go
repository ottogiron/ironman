package field

//Text represents a text field
type Text struct {
	field Field
}

//NewText returns a new text field
func NewText(field Field) *Text {
	return &Text{field}
}
