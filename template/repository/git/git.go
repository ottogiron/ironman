package git

import "github.com/ironman-project/ironman/template/repository"

var _ *repository.Repository = (*repository.Repository)(nil)

//Repository represents an implementation of a ironman Repository
type Repository struct {
	*repository.BaseRepository
}

//New returns a new instance of the git repository
func New(baseRepository *repository.BaseRepository) repository.Repository {
	return &Repository{baseRepository}
}

//Install installs a template from a git url
func (r *Repository) Install(location string) error {
	panic("not implemented")
}

//Update updates a template from a git repository
func (r *Repository) Update(name string) error {
	panic("not implemented")
}
