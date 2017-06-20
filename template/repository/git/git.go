package git

import "github.com/ironman-project/ironman/template/repository"

var _ *repository.Repository = (*repository.Repository)(nil)

//Repository represents an implementation of a ironman Repository
type Repository struct {
}

//Install installs a template from a git url
func (r *Repository) Install(location string) error {
	panic("not implemented")
}

//Update updates a template from a git repository
func (r *Repository) Update(name string) error {
	panic("not implemented")
}

//Uninstall uninstalls a template from a git repository
func (r *Repository) Uninstall(name string) error {
	panic("not implemented")
}

//Find finds an already installed template in this repository
func (r *Repository) Find(name string) error {
	panic("not implemented")
}

//IsInstalled verifies if a template is installed
func (r *Repository) IsInstalled(name string) bool {
	panic("not implemented")
}

//Installed returns a list of installed templates
func (r *Repository) Installed() ([]string, error) {
	panic("not implemented")
}

//Link links a template anywhere in the OS as an installed template
func (r *Repository) Link(templatePath string, templateName string) error {
	panic("not implemented")
}

//Unlink a linked template
func (r *Repository) Unlink(templateName string) error {
	panic("not implemented")
}
