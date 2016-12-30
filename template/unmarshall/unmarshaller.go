package template

//Unmarshaller unmarshalls from a file
type Unmarshaller interface {
	Unmarshall(bytes []byte, out interface{}) error
}
