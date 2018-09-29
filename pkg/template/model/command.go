package model

//Command represents a command to be run
type Command struct {
	Name string   `json:"name"`
	Args []string `json:"args"`
}
