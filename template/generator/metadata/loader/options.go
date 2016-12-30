package loader

import "github.com/ottogiron/ironman/template/unmarshall"

//Option a loader option
type Option func(*Loader)

//FileUnmarshaller file unmarshaller for this loader
func FileUnmarshaller(unmarshaller template.Unmarshaller) Option {
	return func(l *Loader) {
		l.fileUnmarshaller = unmarshaller
	}
}
