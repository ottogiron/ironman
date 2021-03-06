package git

import (
	"io"
	"os"
	"path"
	"strings"

	"github.com/ironman-project/ironman/pkg/template/manager"
	"github.com/pkg/errors"
	gogit "gopkg.in/src-d/go-git.v4"
)

var templateID_ manager.Manager = (*Manager)(nil)

//Manager represents an implementation of a ironman Manager
type Manager struct {
	*manager.BaseManager
	output io.Writer
}

//New returns a new instance of the git Manager
func New(path string, templatesDirectory string, options ...Option) manager.Manager {
	BaseManager := manager.NewBaseManager(path, templatesDirectory)
	m := &Manager{
		BaseManager: BaseManager,
		output:      os.Stdout,
	}
	return m
}

//Install installs a template from a git url
func (r *Manager) Install(location string) (string, error) {
	id := templateIDFromLocation(location)
	templatePath := r.templatePathFromID(id)

	_, err := gogit.PlainClone(templatePath, false,
		&gogit.CloneOptions{
			URL:      location,
			Progress: os.Stdout,
		},
	)

	if err != nil {
		return "", errors.Wrapf(err, "failed to install template  %s", location)
	}
	return id, nil
}

//Update updates a template from a git Manager
func (r *Manager) Update(id string) error {

	templatePath := r.templatePathFromID(id)

	gitRepo, err := gogit.PlainOpen(templatePath)

	if err != nil {
		return errors.Wrapf(err, "failed to open repository %s", id)
	}

	// Get the working directory for the Manager
	w, err := gitRepo.Worktree()

	if err != nil {
		return errors.Wrapf(err, "failed to get template working tree %s", id)
	}

	err = w.Pull(&gogit.PullOptions{
		Progress: os.Stdout,
	})

	if gogit.NoErrAlreadyUpToDate != err && err != nil {
		return errors.Wrapf(err, "failed to Update template  %s", id)
	}
	return nil
}

func (r *Manager) templatePathFromID(templateID string) string {

	templatePath := r.TemplateLocation(templateID)
	return templatePath
}

func templateIDFromLocation(location string) string {
	return path.Base(strings.TrimSuffix(location, ".git"))
}
