package git

import (
	"path/filepath"
	"testing"

	"github.com/ironman-project/ironman/testutils"

	"github.com/ironman-project/ironman/template/repository"
)

func newTestGitRepository() repository.Repository {

	return New("testing")
}

func TestRepository_Install(t *testing.T) {

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
			args{"https://github.com/ottogiron/wizard-hello-world.git"},
			"wizard-hello-world",
			[]string{".xwizard.yml"},
			false,
		},
		{
			"Install unexisting template",
			args{"https://github.com/ottogiron/unexisting-template"},
			"",
			[]string{},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := newTestGitRepository()
			defer func() {
				_ = r.Uninstall(tt.expectedTemplateID)
			}()
			if err := r.Install(tt.args.location); (err != nil) != tt.wantErr {
				t.Errorf("Repository.Install() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			expectedTemplatePath := r.TemplatePath(tt.expectedTemplateID)

			if !testutils.FileExists(expectedTemplatePath) {
				t.Errorf("Repository.Install() template was not installed want path %v", expectedTemplatePath)
			}

			for _, fileRelativePath := range tt.expectedFilesPaths {
				filePath := filepath.Join(expectedTemplatePath, fileRelativePath)
				if !testutils.FileExists(filePath) {
					t.Errorf("Repository.Install() expected file was not found, path %v", filePath)
				}
			}
		})
	}
}

func TestRepository_Update(t *testing.T) {

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
			args{"wizard-hello-world", "https://github.com/ottogiron/wizard-hello-world.git"},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := newTestGitRepository()

			err := r.Install(tt.args.location)

			defer func() {
				r.Uninstall(tt.args.id)
			}()

			if err != nil {
				t.Errorf("Repository.Update() error = %v", err)
			}

			if err := r.Update(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("Repository.Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
