package manager

import (
	"os"
	"path/filepath"

	"io/ioutil"

	"github.com/ironman-project/ironman/template"
	"github.com/pkg/errors"
)

//Manager represents a local ironman manager
type Manager interface {
	Install(templateLocator string) (ID string, err error)
	Update(templateID string) error
	Uninstall(templateID string) error
	Find(templateID string) error
	IsInstalled(templateID string) (bool, error)
	Installed() ([]*template.Metadata, error)
	Link(templatePath string, templateID string) (string, error)
	Unlink(templateID string) error
	TemplateLocation(templateID string) string
}

//BaseManager implements basic generic manager operations
type BaseManager struct {
	path               string
	templatesPath      string
	templatesDirectory string
}

//NewBaseManager returns a new instance of a base manager
func NewBaseManager(path string, managerTemplatesDirectory string) *BaseManager {
	templatesPath := filepath.Join(path, managerTemplatesDirectory)
	return &BaseManager{path, templatesPath, managerTemplatesDirectory}
}

//Uninstall uninstalls a template
func (b *BaseManager) Uninstall(templateID string) error {
	if err := validateTemplateID(templateID); err != nil {
		return err
	}
	templatePath := b.TemplateLocation(templateID)
	err := os.RemoveAll(templatePath)
	if err != nil {
		return errors.Wrapf(err, "Failed to remove template %s", templateID)
	}
	return nil
}

//Find finds a template in the manager
func (b *BaseManager) Find(templateID string) error {
	panic("not implemented")
}

//IsInstalled verifies if template is installed
func (b *BaseManager) IsInstalled(templateID string) (bool, error) {
	if err := validateTemplateID(templateID); err != nil {
		return false, err
	}
	templatePath := b.TemplateLocation(templateID)
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

//TemplateLocation returns the file system path of a template based on the ID
func (b *BaseManager) TemplateLocation(templateID string) string {
	return filepath.Join(b.path, b.templatesDirectory, templateID)
}

//Installed returns a lists of installed templates
func (b *BaseManager) Installed() ([]*template.Metadata, error) {

	files, err := ioutil.ReadDir(b.templatesPath)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to list al the available templates")
	}

	var templatesList []*template.Metadata
	for _, f := range files {
		templatesList = append(templatesList, &template.Metadata{ID: f.Name()})
	}

	return templatesList, nil
}

//Link links a template on a path to the manager
func (b *BaseManager) Link(templatePath string, templateID string) (string, error) {

	if err := validateTemplateID(templateID); err != nil {
		return "", err
	}

	linkPath := b.TemplateLocation(templateID)

	if _, err := os.Stat(templatePath); os.IsNotExist(err) {
		return "", errors.Wrapf(err, "Failed to create symlink to ironman manager path should %s exists ", templatePath)
	}

	absTemplatePath, err := filepath.Abs(templatePath)

	if err != nil {
		return "", errors.Wrapf(err, "Failed to create symlink to ironman manager for %s with ID %s", templatePath, templateID)
	}

	err = os.Symlink(absTemplatePath, linkPath)
	if err != nil {
		return "", errors.Wrapf(err, "Failed to create symlink to ironman manager for %s with ID %s", templatePath, templateID)
	}

	return linkPath, nil
}

//Unlink unlinks a linked template
func (b *BaseManager) Unlink(templateID string) error {

	if err := validateTemplateID(templateID); err != nil {
		return err
	}

	templatePath := b.TemplateLocation(templateID)

	if _, err := os.Stat(templatePath); os.IsNotExist(err) {
		return errors.Wrapf(err, "Failed to remove symlink for template ID %s", err)
	}

	err := os.Remove(templatePath)
	if err != nil {
		return errors.Wrapf(err, "Failed to remove symlink for template ID %s", templateID)
	}
	return nil
}

//Install not implemented for base manager since it depends on specific provider
func (b *BaseManager) Install(templateLocator string) error {
	panic("not implemented")
}

//Update not implemented for base manager since it depend on specific provider
func (b *BaseManager) Update(templateID string) error {
	panic("not implemented")
}
