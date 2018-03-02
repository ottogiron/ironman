package model

//Template template metadata definition
type Template struct {
	ID          string
	Name        string
	Description string
	Generators  []*Generator
}

//Generator generator metadata definition
type Generator struct {
	ID          string
	Name        string
	Description string
}
