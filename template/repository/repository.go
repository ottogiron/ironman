package repository

import (
	"os"
	"path/filepath"

	"io/ioutil"

	"github.com/pkg/errors"
)

//Repository represents a local ironman repository
type Repository interface {
	Install(templateLocator string) error
	Update(templateID string) error
	Uninstall(templateID string) error
	Find(templateID string) error
	IsInstalled(templateID string) (bool, error)
	Installed() ([]string, error)
	Link(templatePath string, templateName string) error
	Unlink(templateName string) error
}

const (
	repositoryTemplatesDirectory = "templates"
)

//BaseRepository implements basic generic repository operations
type BaseRepository struct {
	path          string
	templatesPath string
}

//NewBaseRepository returns a new instance of a base repository
func NewBaseRepository(path string) Repository {
	templatesPath := filepath.Join(path, repositoryTemplatesDirectory)
	return &BaseRepository{path, templatesPath}
}

//Uninstall uninstalls a template
func (b *BaseRepository) Uninstall(templateID string) error {
	if err := validateTemplateID(templateID); err != nil {
		return err
	}
	templatePath := b.templatePath(templateID)
	err := os.RemoveAll(templatePath)
	if err != nil {
		return errors.Wrapf(err, "Failed to remove template %s", templateID)
	}
	return nil
}

//Find finds a template in the repository
func (b *BaseRepository) Find(templateID string) error {
	panic("not implemented")
}

//IsInstalled verifies if template is installed
func (b *BaseRepository) IsInstalled(templateID string) (bool, error) {
	if err := validateTemplateID(templateID); err != nil {
		return false, err
	}
	templatePath := b.templatePath(templateID)
	_, err := os.Stat(templatePath)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, errors.Wrapf(err, "verifying template installation ID: %s", templateID)
	}
	return true, nil
}

func validateTemplateID(templateID string) error {
	if templateID == "" {
		return errors.Errorf("a templateID cannot be empty")
	}
	return nil
}

func (b *BaseRepository) templatePath(templateID string) string {
	return filepath.Join(b.path, repositoryTemplatesDirectory, templateID)
}

//Installed returns a lists of installed templates
func (b *BaseRepository) Installed() ([]string, error) {

	files, err := ioutil.ReadDir(b.templatesPath)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to list al the available templates")
	}

	var templatesList []string
	for _, f := range files {
		templatesList = append(templatesList, f.Name())
	}

	return templatesList, nil
}

//Link links a template on a path to the repository
func (b *BaseRepository) Link(templatePath string, templateName string) error {
	panic("not implemented")
}

//Unlink unlinks a linked template
func (b *BaseRepository) Unlink(templateName string) error {
	panic("not implemented")
}

//Install not implemented for base repository since it depends on specific provider
func (b *BaseRepository) Install(templateLocator string) error {
	panic("not implemented")
}

//Update not implemented for base repository since it depend on specific provider
func (b *BaseRepository) Update(templateID string) error {
	panic("not implemented")
}
