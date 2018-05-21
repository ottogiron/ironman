package engine

import "io"

//Engine represents a template engine
type Engine interface {
	Parse(text string) (Engine, error)
	Execute(writer io.Writer, data interface{}) error
}

//Factory definition of an engine factory
type Factory func() Engine
