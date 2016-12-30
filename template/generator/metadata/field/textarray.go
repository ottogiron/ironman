package field

import "github.com/ottogiron/ironman/template/generator/metadata"

//TextArray a fixed size array of text fields
type TextArray struct {
	metadata.Field
	Size  int
	field *Text
}
