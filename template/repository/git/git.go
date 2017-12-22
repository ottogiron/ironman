package git

import (
	"os"
	"path"
	"strings"

	"github.com/ironman-project/ironman/template/repository"
	"github.com/pkg/errors"
	gogit "gopkg.in/src-d/go-git.v4"
)

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
	templateID := path.Base(strings.TrimSuffix(location, ".git"))
	templatePath := r.TemplatePath(templateID)
	_, err := gogit.PlainClone(templatePath, false, &gogit.CloneOptions{
		URL:      location,
		Progress: os.Stdout,
	})

	if err != nil {
		return errors.Wrapf(err, "Failed to clone template repository %s", location)
	}
	return nil
}

//Update updates a template from a git repository
func (r *Repository) Update(name string) error {
	panic("not implemented")
}
