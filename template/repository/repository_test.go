package repository

import (
	"reflect"
	"testing"
)

const (
	testRepositoryPath = "testing/repository"
)

func TestNewBaseRepository(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want Repository
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewBaseRepository(tt.args.path); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBaseRepository() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBaseRepository_Uninstall(t *testing.T) {
	type args struct {
		templateID string
	}
	tests := []struct {
		name    string
		b       *BaseRepository
		args    args
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &BaseRepository{}
			if err := b.Uninstall(tt.args.templateID); (err != nil) != tt.wantErr {
				t.Errorf("BaseRepository.Uninstall() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestBaseRepository_Find(t *testing.T) {
	type args struct {
		templateID string
	}
	tests := []struct {
		name    string
		b       *BaseRepository
		args    args
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &BaseRepository{}
			if err := b.Find(tt.args.templateID); (err != nil) != tt.wantErr {
				t.Errorf("BaseRepository.Find() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestBaseRepository_IsInstalled(t *testing.T) {
	type args struct {
		templateID string
	}
	tests := []struct {
		name           string
		repositoryPath string
		args           args
		want           bool
		wantErr        bool
	}{
		{"Template is installed", testRepositoryPath, args{"valid"}, true, false},
		{"Template is not installed", testRepositoryPath, args{"not_installed"}, false, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := NewBaseRepository(tt.repositoryPath)
			got, err := b.IsInstalled(tt.args.templateID)
			if (err != nil) != tt.wantErr {
				t.Errorf("BaseRepository.IsInstalled() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("BaseRepository.IsInstalled() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBaseRepository_Installed(t *testing.T) {
	tests := []struct {
		name    string
		b       *BaseRepository
		want    []string
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &BaseRepository{}
			got, err := b.Installed()
			if (err != nil) != tt.wantErr {
				t.Errorf("BaseRepository.Installed() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BaseRepository.Installed() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBaseRepository_Link(t *testing.T) {
	type args struct {
		templatePath string
		templateName string
	}
	tests := []struct {
		name    string
		b       *BaseRepository
		args    args
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &BaseRepository{}
			if err := b.Link(tt.args.templatePath, tt.args.templateName); (err != nil) != tt.wantErr {
				t.Errorf("BaseRepository.Link() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestBaseRepository_Unlink(t *testing.T) {
	type args struct {
		templateName string
	}
	tests := []struct {
		name    string
		b       *BaseRepository
		args    args
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &BaseRepository{}
			if err := b.Unlink(tt.args.templateName); (err != nil) != tt.wantErr {
				t.Errorf("BaseRepository.Unlink() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestBaseRepository_Install(t *testing.T) {
	type args struct {
		templateLocator string
	}
	tests := []struct {
		name    string
		b       *BaseRepository
		args    args
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &BaseRepository{}
			if err := b.Install(tt.args.templateLocator); (err != nil) != tt.wantErr {
				t.Errorf("BaseRepository.Install() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestBaseRepository_Update(t *testing.T) {
	type args struct {
		templateID string
	}
	tests := []struct {
		name    string
		b       *BaseRepository
		args    args
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &BaseRepository{}
			if err := b.Update(tt.args.templateID); (err != nil) != tt.wantErr {
				t.Errorf("BaseRepository.Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
