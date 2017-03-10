package loader

import "github.com/ironman-project/ironman/text"

//Option a loader option
type Option func(*Loader)

//FileUnmarshaller file unmarshaller for this loader
func FileUnmarshaller(unmarshaller text.Unmarshaller) Option {
	return func(l *Loader) {
		l.fileUnmarshaller = unmarshaller
	}
}
