package ironman

import (
	"github.com/ironman-project/ironman/template/manager"
	"github.com/ironman-project/ironman/template/repository"
)

//Option represents an ironman options
type Option func(*Ironman)

//SetTemplateManager sets ironman's template manager
func SetTemplateManager(manager manager.Manager) Option {
	return func(i *Ironman) {
		i.manager = manager
	}
}

//SetTemplateRepository sets the ironman template repository
func SetTemplateRepository(repository repository.Repository) Option {
	return func(i *Ironman) {
		i.repository = repository
	}
}
