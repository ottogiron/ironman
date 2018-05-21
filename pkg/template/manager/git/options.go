package git

import "io"

//Option represents a git manager setter
type Option func(mananger *Manager)

//SetOutput sets the writer output for this manager
func SetOutput(output io.Writer) Option {
	return func(manager *Manager) {
		manager.output = output
	}
}
