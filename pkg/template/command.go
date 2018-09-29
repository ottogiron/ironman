package template

import (
	"io"
	"os/exec"

	"github.com/ironman-project/ironman/pkg/template/model"
	"github.com/pkg/errors"
)

//ExecuteCommand executes an ironman model command
func ExecuteCommand(command *model.Command, output io.Writer) error {
	name := command.Name
	if name == "" {
		return errors.New("the command name cannot be empty")
	}
	cmd := exec.Command(command.Name, command.Args...)
	cmd.Stdout = output

	if err := cmd.Run(); err != nil {
		return errors.Errorf("failed to run command %s with args %v", name, command.Args)
	}
	return nil
}
