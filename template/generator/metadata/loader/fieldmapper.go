package loader

import (
	"github.com/ottogiron/ironman/template/generator/metadata"
)

//FieldMapper maps a loaded field from map[string]interface{} to metadata/field
type FieldMapper func(*metadata.Metadata)
