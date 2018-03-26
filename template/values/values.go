package values

//Values Represents the values read from a reader
type Values map[string]interface{}

//Reader values reader interface
type Reader interface {
	Read() (Values, error)
}
