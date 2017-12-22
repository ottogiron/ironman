package git

import (
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
	templatePath := r.templatePathFromLocation(location)
	gitRepo, err := gogit.NewFilesystemRepository(templatePath)

	if err != nil {
		return errors.Wrapf(err, "Failed to get template repository %s", location)
	}
	err = gitRepo.Clone(&gogit.CloneOptions{
		URL: location,
	})

	if err != nil {
		return errors.Wrapf(err, "Failed to install template  %s", location)
	}
	return nil
}

//Update updates a template from a git repository
func (r *Repository) Update(location string) error {
	templatePath := r.templatePathFromLocation(location)
	gitRepo, err := gogit.NewFilesystemRepository(templatePath)

	if err != nil {
		return errors.Wrapf(err, "Failed to get template repository %s", location)
	}

	err = gitRepo.Pull(&gogit.PullOptions{})

	if err != nil {
		return errors.Wrapf(err, "Failed to update template  %s", location)
	}

	return nil
}

func (r *Repository) templatePathFromLocation(location string) string {
	templateID := path.Base(strings.TrimSuffix(location, ".git"))
	templatePath := r.TemplatePath(templateID)
	return templatePath
}
