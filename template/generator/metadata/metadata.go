package metadata

//Metadata Defines a generator metadata
type Metadata struct {
	Name         string
	Description  string
	Fields       []interface{}
	Pregenerate  []string
	Postgenerate []string
}
