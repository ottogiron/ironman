package template

//Unmarshaller unmarshalls from a file
type Unmarshaller interface {
	Unmarshall(path string, out interface{}) error
}
