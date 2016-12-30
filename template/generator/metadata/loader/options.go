package loader

import "github.com/ottogiron/ironman/template/unmarshall"

//Option a loader option
type Option func(*Loader)

//Path path of the file to load
func Path(path string) Option {
	return func(l *Loader) {
		l.path = path
	}
}

//FileUnmarshaller file unmarshaller for this loader
func FileUnmarshaller(unmarshaller template.Unmarshaller) Option {
	return func(l *Loader) {
		l.fileUnmarshaller = unmarshaller
	}
}
