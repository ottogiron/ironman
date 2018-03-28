package model

import (
	"testing"

	"github.com/ironman-project/ironman/testutils"
)

func Test_fsReader_Read(t *testing.T) {
	type fields struct {
		ignore        []string
		path          string
		fileExtension MetadataFileExtension
		decoder       Decoder
	}
	tests := []struct {
		name    string
		fields  fields
		want    *Template
		wantErr bool
	}{
		{
			"Read template metadata from file system",
			fields{
				[]string{".git"},
				"testing/test_read_template",
				"yaml",
				NewDecoder(DecoderTypeYAML),
			},
			&Template{
				ID:            "test-read-template",
				Name:          "Test Read Template",
				Description:   "This is a test template",
				DirectoryName: "test_read_template",
				Generators: []*Generator{
					&Generator{
						ID:            "generator",
						Name:          "Test Generator",
						Description:   "This is a test generator",
						DirectoryName: "generator",
					},
				},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &fsReader{
				ignoreFiles:   tt.fields.ignore,
				fileExtension: tt.fields.fileExtension,
				decoder:       tt.fields.decoder,
			}
			got, err := r.Read(tt.fields.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("fsReader.Read() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if got.DirectoryName != tt.want.DirectoryName {
				t.Errorf("fsReader.Read() directory_name = %s want %s", got.DirectoryName, tt.want.DirectoryName)
			}

			for i, generator := range tt.want.Generators {
				gotGenerator := got.Generators[i]
				if generator.DirectoryName != gotGenerator.DirectoryName {
					t.Errorf("fsReader.Read() generator directory_name = %s want %s", gotGenerator.DirectoryName, generator.DirectoryName)
				}
			}

			gotM := testutils.Marshal(got, t)
			wantM := testutils.Marshal(tt.want, t)
			if gotM != wantM {
				t.Errorf("fsReader.Read() = %v, want %v", gotM, wantM)
			}

		})
	}
}
