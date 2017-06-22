package git

import (
	"reflect"
	"testing"

	"github.com/ironman-project/ironman/template/repository"
)

func TestNew(t *testing.T) {
	type args struct {
		baseRepository *repository.BaseRepository
	}
	tests := []struct {
		name string
		args args
		want repository.Repository
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.baseRepository); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRepository_Install(t *testing.T) {
	type fields struct {
		BaseRepository *repository.BaseRepository
	}
	type args struct {
		location string
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
			if err := r.Install(tt.args.location); (err != nil) != tt.wantErr {
				t.Errorf("Repository.Install() error = %v, wantErr %v", err, tt.wantErr)
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
