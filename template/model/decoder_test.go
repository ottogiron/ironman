package model

import (
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func getTestingFile(name string, t *testing.T) *os.File {
	path := filepath.Join("testing", name)
	f, err := os.Open(path)
	if err != nil {
		t.Fatalf("Failed to open test file %s", path)
	}
	return f
}

func Test_yamlDecoder_Decode(t *testing.T) {
	type args struct {
		testFile string
	}
	tests := []struct {
		name    string
		args    args
		want    Template
		wantErr bool
	}{
		{
			"decode yaml file",
			args{"decode.yaml"},
			Template{
				ID:          "template-id",
				Name:        "the template name",
				Description: "the template description",
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			yr := &yamlDecoder{}
			reader := getTestingFile(tt.args.testFile, t)
			var got Template
			err := yr.Decode(&got, reader)
			if (err != nil) != tt.wantErr {
				t.Errorf("yamlDecoder.Decode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("yamlDecoder.Decode() = %v, want %v", &got, tt.want)
			}
		})
	}
}
