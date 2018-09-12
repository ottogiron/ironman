package manager

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/ironman-project/ironman/pkg/template"
	"github.com/ironman-project/ironman/pkg/testutils"
)

var (
	testManagerPath        = "testing/repository"
	testTemplatesDirectory = "templates"
	testTemplatesPath      = filepath.Join(testManagerPath, testTemplatesDirectory)
)

func createTestTemplate(t *testing.T, names ...string) (string, func()) {
	tempManager, err := ioutil.TempDir("", "ironman-test-manager")
	if err != nil {
		t.Fatalf("failed to create test manager %s", err)
	}
	sourcePath := filepath.Join(testManagerPath, testTemplatesDirectory, "base")
	for _, name := range names {
		destPath := filepath.Join(tempManager, name)
		testutils.CopyDir(sourcePath, destPath, t)

	}

	return tempManager, func() {
		err := os.RemoveAll(tempManager)
		if err != nil {
			t.Fatalf("failed to clean test manager %s", err)
		}
	}
}

func TestNewBaseManager(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want Manager
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewBaseManager(tt.args.path, "templates"); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBaseManager() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBaseManager_Uninstall(t *testing.T) {
	type args struct {
		templateID string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"Uninstall template", args{"valid_removable"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			managerPath, clean := createTestTemplate(t, tt.args.templateID)
			defer clean()
			b := NewBaseManager(managerPath, testTemplatesDirectory)
			if err := b.Uninstall(tt.args.templateID); (err != nil) != tt.wantErr {
				t.Errorf("BaseManager.Uninstall() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestBaseManager_Find(t *testing.T) {
	type args struct {
		templateID string
	}
	tests := []struct {
		name    string
		b       *BaseManager
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &BaseManager{}
			if err := b.Find(tt.args.templateID); (err != nil) != tt.wantErr {
				t.Errorf("BaseManager.Find() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestBaseManager_Installed(t *testing.T) {
	tests := []struct {
		name        string
		managerPath string
		want        []*template.Metadata
		wantErr     bool
	}{
		{"All the installed templates", testManagerPath, []*template.Metadata{&template.Metadata{ID: "base"}, &template.Metadata{ID: "valid"}}, false},
		{"Non existing manager path", "unexistingPath", nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := NewBaseManager(tt.managerPath, testTemplatesDirectory)
			got, err := b.Installed()
			if (err != nil) != tt.wantErr {
				t.Errorf("BaseManager.Installed() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BaseManager.Installed() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBaseManager_Link(t *testing.T) {
	type args struct {
		templatePath string
		templateID   string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"Link a template",
			args{
				templatePath: filepath.Join("testing", "repository", "templates", "valid"),
				templateID:   "dev-valid",
			},
			false,
		},
		{
			"Link a template with non existing path",
			args{
				templatePath: filepath.Join("nonexisting", "repository", "templates", "valid"),
				templateID:   "dev-nonexisting",
			},
			true,
		},
		{
			"Link a template with invalid ID",
			args{
				templatePath: filepath.Join("nonexisting", "repository", "templates", "valid"),
				templateID:   "",
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			b := NewBaseManager(testManagerPath, testTemplatesDirectory)
			createdLinkPath := filepath.Join(testTemplatesPath, tt.args.templateID)
			defer func() {
				_ = os.Remove(createdLinkPath)
			}()

			if _, err := b.Link(tt.args.templatePath, tt.args.templateID); (err != nil) != tt.wantErr {
				t.Errorf("BaseManager.Link() error = %v, wantErr %v", err, tt.wantErr)
				return
			} else if tt.wantErr {
				return
			}

			if !testutils.FileExists(createdLinkPath) {
				t.Errorf("BaseManager.Link() %s file should exists", createdLinkPath)
				return
			}

			ymlFilePath := filepath.Join(createdLinkPath, ".ironman.yaml")
			if !testutils.FileExists(ymlFilePath) {
				t.Errorf("BaseManager.Link() %s file should exists", ymlFilePath)
			}
		})
	}
}

func TestBaseManager_Unlink(t *testing.T) {
	type args struct {
		templatePath     string
		templateID       string
		unlinkTemplateID string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"Unlink template",
			args{
				templatePath:     filepath.Join("testing", "repository", "templates", "valid"),
				templateID:       "dev-valid",
				unlinkTemplateID: "dev-valid",
			},
			false,
		},
		{
			"Unlink template with non existing id",
			args{
				templatePath:     filepath.Join("testing", "repository", "templates", "valid"),
				templateID:       "dev-valid",
				unlinkTemplateID: "non-existing",
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			b := NewBaseManager(testManagerPath, testTemplatesDirectory)
			createdLinkPath := filepath.Join(testTemplatesPath, tt.args.templateID)
			defer func() {
				_ = os.Remove(createdLinkPath)
			}()
			_, _ = b.Link(tt.args.templatePath, tt.args.templateID)

			if err := b.Unlink(tt.args.unlinkTemplateID); (err != nil) != tt.wantErr {
				t.Errorf("BaseManager.Unlink() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestBaseManager_Install(t *testing.T) {
	type args struct {
		templateLocator string
	}
	tests := []struct {
		name    string
		b       *BaseManager
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &BaseManager{}
			if err := b.Install(tt.args.templateLocator); (err != nil) != tt.wantErr {
				t.Errorf("BaseManager.Install() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestBaseManager_Update(t *testing.T) {
	type args struct {
		templateID string
	}
	tests := []struct {
		name    string
		b       *BaseManager
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &BaseManager{}
			if err := b.Update(tt.args.templateID); (err != nil) != tt.wantErr {
				t.Errorf("BaseManager.Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
