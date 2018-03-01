package git

import (
	"path/filepath"
	"testing"

	"github.com/ironman-project/ironman/testutils"

	"github.com/ironman-project/ironman/template/manager"
)

func newTestGitManager() manager.Manager {

	return New("testing")
}

func TestManager_Install(t *testing.T) {

	type args struct {
		location string
	}
	tests := []struct {
		name               string
		args               args
		expectedTemplateID string
		expectedFilesPaths []string
		wantErr            bool
	}{
		{
			"Install template",
			args{"https://github.com/ironman-project/template-example.git"},
			"template-example",
			[]string{".ironman.yml"},
			false,
		},
		{
			"Install unexisting template",
			args{"https://github.com/ironman-project/unexisting-template"},
			"",
			[]string{},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := newTestGitManager()
			defer func() {
				_ = r.Uninstall(tt.expectedTemplateID)
			}()
			if err := r.Install(tt.args.location); (err != nil) != tt.wantErr {
				t.Errorf("Manager.Install() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			expectedTemplatePath := r.TemplatePath(tt.expectedTemplateID)

			if !testutils.FileExists(expectedTemplatePath) {
				t.Errorf("Manager.Install() template was not installed want path %v", expectedTemplatePath)
			}

			for _, fileRelativePath := range tt.expectedFilesPaths {
				filePath := filepath.Join(expectedTemplatePath, fileRelativePath)
				if !testutils.FileExists(filePath) {
					t.Errorf("Manager.Install() expected file was not found, path %v", filePath)
				}
			}
		})
	}
}

func TestManager_Update(t *testing.T) {

	type args struct {
		id       string
		location string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"Update template",
			args{"template-example", "https://github.com/ironman-project/template-example.git"},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := newTestGitManager()

			err := r.Install(tt.args.location)

			defer func() {
				r.Uninstall(tt.args.id)
			}()

			if err != nil {
				t.Errorf("Manager.Update() error = %v", err)
			}

			if err := r.Update(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("Manager.Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
