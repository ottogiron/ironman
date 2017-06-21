package repository

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

//BaseRepository implements basic generic repository operations
type BaseRepository struct {
}

//Install not implemented for base repository since it depends on specific provider
func (b *BaseRepository) Install(templateLocator string) error {
	panic("not implemented")
}

//Update not implemented for base repository since it depend on specific provider
func (b *BaseRepository) Update(templateID string) error {
	panic("not implemented")
}

//Uninstall uninstalls a template
func (b *BaseRepository) Uninstall(templateID string) error {
	panic("not implemented")
}

//Find finds a template in the repository
func (b *BaseRepository) Find(templateID string) error {
	panic("not implemented")
}

//IsInstalled verifies if template is installed
func (b *BaseRepository) IsInstalled(templateID string) (bool, error) {
	panic("not implemented")
}

//Installed returns a lists of installed templates
func (b *BaseRepository) Installed() ([]string, error) {
	panic("not implemented")
}

//Link links a template on a path to the repository
func (b *BaseRepository) Link(templatePath string, templateName string) error {
	panic("not implemented")
}

//Unlink unlinks a linked template
func (b *BaseRepository) Unlink(templateName string) error {
	panic("not implemented")
}
