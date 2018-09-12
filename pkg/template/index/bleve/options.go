package bleve

import (
	"github.com/blevesearch/bleve"
)

//Option defines a bleeve repository option
type Option func(*bleeveIndex)

//SetIndex sets the index name for the repository
func SetIndex(index bleve.Index) Option {
	return func(r *bleeveIndex) {
		r.index = index
	}
}
