package ironman

import (
	"io"

	"github.com/ironman-project/ironman/pkg/template/index"
	"github.com/ironman-project/ironman/pkg/template/manager"
	"github.com/ironman-project/ironman/pkg/template/validator"
)

//Option represents an ironman options
type Option func(*Ironman)

//SetTemplateManager sets ironman's template manager
func SetTemplateManager(manager manager.Manager) Option {
	return func(i *Ironman) {
		i.manager = manager
	}
}

//SetTemplateIndex sets the ironman template index
func SetTemplateIndex(index index.Index) Option {
	return func(i *Ironman) {
		i.index = index
	}
}

//SetValidators sets the model validators
func SetValidators(validators ...validator.Validator) Option {
	return func(i *Ironman) {
		for _, validator := range validators {
			i.validators = append(i.validators, validator)
		}
	}
}

//SetOutput sets ironman output writer
func SetOutput(output io.Writer) Option {
	return func(i *Ironman) {
		i.output = output
	}
}
