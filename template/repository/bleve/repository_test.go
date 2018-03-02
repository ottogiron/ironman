package bleve

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/blevesearch/bleve"
	"github.com/ironman-project/ironman/template/model"
)

func newBleveTestIndex(t *testing.T) (bleve.Index, func()) {
	dir, err := ioutil.TempDir("", "ironman-bleve-test")
	if err != nil {
		t.Fatal("Failed to create test bleve index directory", err)
	}

	indexPath := filepath.Join(dir, "index")
	mapping := bleve.NewIndexMapping()
	index, err := bleve.New(indexPath, mapping)

	if err != nil {
		t.Fatal("Failed to create test bleve index", err)
	}
	return index, func() {
		err := index.Close()

		if err != nil {
			t.Fatal("Failed to close bleve index", err)
		}
		err = os.RemoveAll(dir)
		if err != nil {
			t.Fatal("Failed to clean bleve index", err)
		}
	}
}

func Test_bleeveRepository_Index(t *testing.T) {

	type args struct {
		template model.Template
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"Index a template",
			args{
				model.Template{
					ID:          "test-template",
					Name:        "Test template",
					Description: "This is a test template",
					Generators: []*model.Generator{
						&model.Generator{
							ID:          "test-generator",
							Name:        "Test generator",
							Description: "This is a test generator",
						},
					},
				},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			index, clean := newBleveTestIndex(t)
			defer clean()
			r := New(SetIndex(index))
			var id string
			var err error
			if id, err = r.Index(tt.args.template); (err != nil) != tt.wantErr {
				t.Errorf("bleeveRepository.Index() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			doc, err := index.Document(id)

			if err != nil {
				t.Errorf("bleeveRepository.Index() error = %v, wantErr %v", err, tt.wantErr)
			}

			if doc == nil {
				t.Errorf("bleeveRepository.Index() nil , want %v", tt.args.template)
			}
		})
	}
}
