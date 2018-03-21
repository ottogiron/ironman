package manager

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/ironman-project/ironman/template"
	"github.com/ironman-project/ironman/testutils"
)

var (
	testManagerPath   = "testing/repository"
	testTemplatesPath = filepath.Join(testManagerPath, managerTemplatesDirectory)
)

func createTestTemplate(t *testing.T, names ...string) (string, func()) {
	tempManager, err := ioutil.TempDir("", "ironman-test-manager")
	if err != nil {
		t.Fatalf("Failed to create test manager %s", err)
	}
	sourcePath := filepath.Join(testManagerPath, "templates", "base")
	for _, name := range names {
		destPath := filepath.Join(tempManager, name)
		err = testutils.CopyDir(sourcePath, destPath)
		if err != nil {
			t.Fatalf("Failed to create test template %s", err)
		}
	}

	return tempManager, func() {
		err := os.RemoveAll(tempManager)
		if err != nil {
			t.Fatalf("Failed to clean test manager %s", err)
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
			if got := NewBaseManager(tt.args.path); !reflect.DeepEqual(got, tt.want) {
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
			b := NewBaseManager(managerPath)
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

func TestBaseManager_IsInstalled(t *testing.T) {
	type args struct {
		templateID string
	}
	tests := []struct {
		name        string
		managerPath string
		args        args
		want        bool
		wantErr     bool
	}{
		{"Template is installed", testManagerPath, args{"valid"}, true, false},
		{"Template is not installed", testManagerPath, args{"not_installed"}, false, false},
		{"Template invalid empty name", testManagerPath, args{""}, false, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := NewBaseManager(tt.managerPath)
			got, err := b.IsInstalled(tt.args.templateID)
			if (err != nil) != tt.wantErr {
				t.Errorf("BaseManager.IsInstalled() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("BaseManager.IsInstalled() = %v, want %v", got, tt.want)
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
			b := NewBaseManager(tt.managerPath)
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

			b := NewBaseManager(testManagerPath)
			createdLinkPath := filepath.Join(testTemplatesPath, tt.args.templateID)
			defer func() {
				_ = os.Remove(createdLinkPath)
			}()

			if err := b.Link(tt.args.templatePath, tt.args.templateID); (err != nil) != tt.wantErr {
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

			b := NewBaseManager(testManagerPath)
			createdLinkPath := filepath.Join(testTemplatesPath, tt.args.templateID)
			defer func() {
				_ = os.Remove(createdLinkPath)
			}()
			_ = b.Link(tt.args.templatePath, tt.args.templateID)

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

func TestInitIronmanHome(t *testing.T) {
	type args struct {
		ironmanHome string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := InitIronmanHome(tt.args.ironmanHome); (err != nil) != tt.wantErr {
				t.Errorf("InitIronmanHome() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
