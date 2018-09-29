package template

import (
	"bytes"
	"testing"

	"github.com/ironman-project/ironman/pkg/template/model"
)

func TestExecuteCommand(t *testing.T) {
	type args struct {
		command *model.Command
	}
	tests := []struct {
		name       string
		args       args
		wantOutput string
		wantErr    bool
	}{
		{
			name: "Execute echo",
			args: args{
				command: &model.Command{
					Name: "echo",
					Args: []string{"-n", "Hello world!"},
				},
			},
			wantOutput: "Hello world!",
			wantErr:    false,
		},
		{
			name: "General run command failure",
			args: args{
				command: &model.Command{
					Name: "randomcommand",
					Args: nil,
				},
			},
			wantOutput: "",
			wantErr:    true,
		},
		{
			name: "Command name empty",
			args: args{
				command: &model.Command{
					Name: "",
					Args: nil,
				},
			},
			wantOutput: "",
			wantErr:    true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := &bytes.Buffer{}
			if err := ExecuteCommand(tt.args.command, output); (err != nil) != tt.wantErr {
				t.Errorf("ExecuteCommand() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotOutput := output.String(); gotOutput != tt.wantOutput {
				t.Errorf("ExecuteCommand() = %v, want %v", gotOutput, tt.wantOutput)
			}
		})
	}
}
