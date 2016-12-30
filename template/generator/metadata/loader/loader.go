package loader

import "github.com/ottogiron/ironman/template/generator/metadata"

//Loader metadata loader
type Loader interface {
	Load() *metadata.Metadata
}
