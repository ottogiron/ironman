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
