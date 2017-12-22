package git

import (
	"os"
	"testing"

	"github.com/ironman-project/ironman/template/repository"
)

func newTestGitRepository() repository.Repository {
	baseRepository := repository.NewBaseRepository("testing")
	return New(baseRepository)
}

func TestRepository_Install(t *testing.T) {

	type args struct {
		location string
	}
	tests := []struct {
		name               string
		args               args
		expectedTemplateID string
		wantErr            bool
	}{
		{
			"Install template",
			args{"https://github.com/ottogiron/wizard-hello-world.git"},
			"wizard-hello-world",
			false,
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
			if _, err := os.Stat(expectedTemplatePath); os.IsNotExist(err) {
				t.Errorf("Repository.Install() template was not installed want path %v", expectedTemplatePath)
			}
		})
	}
}

func TestRepository_Update(t *testing.T) {
	type fields struct {
		BaseRepository *repository.BaseRepository
	}
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Repository{
				BaseRepository: tt.fields.BaseRepository,
			}
			if err := r.Update(tt.args.name); (err != nil) != tt.wantErr {
				t.Errorf("Repository.Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
