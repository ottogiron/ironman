package bleve

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/blevesearch/bleve"
	"github.com/ironman-project/ironman/template/model"
	"github.com/ironman-project/ironman/template/repository"
	uuid "github.com/satori/go.uuid"
)

func tempIndexPath(t *testing.T) string {
	dir, err := ioutil.TempDir("", "ironman-bleve-test")
	if err != nil {
		t.Fatal("Failed to create test bleve index directory", err)
	}

	indexPath := filepath.Join(dir, "index")
	return indexPath
}

func newTestRepository(t *testing.T) (repository.Repository, bleve.Index, func()) {
	path := tempIndexPath(t)

	index, err := BuildIndex(path)

	if err != nil {
		t.Fatalf("Failed to open test index %s", err)
	}

	r := New(SetIndex(index))

	return r, index, func() {
		err := index.Close()

		if err != nil {
			t.Fatal("Failed to close bleve index", err)
		}
		err = os.RemoveAll(path)
		if err != nil {
			t.Fatal("Failed to clean bleve index", err)
		}
	}
}

func Test_bleeveRepository_Index(t *testing.T) {

	type args struct {
		template *model.Template
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"Index a template",
			args{
				&model.Template{
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
			r, index, clean := newTestRepository(t)
			defer clean()
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

			got, err := deserialize(doc)

			if err != nil {
				t.Errorf("bleeveRepository.Index()  error = %v", err)
			}

			if got != nil && !reflect.DeepEqual(got, tt.args.template) {
				t.Errorf("bleeveRepository.FindTemplateByID() = %v, want %v", got, tt.args.template)
			}

		})
	}
}

func Test_bleeveRepository_Update(t *testing.T) {

	type args struct {
		template *model.Template
	}
	tests := []struct {
		name     string
		args     args
		template *model.Template
		wantErr  bool
	}{
		{
			"Update template index",
			args{
				&model.Template{
					ID:            "template-id",
					Name:          "Updated name",
					Description:   "Updated description",
					DirectoryName: "test",
					Generators: []*model.Generator{
						&model.Generator{
							ID:            "test-generator",
							Name:          "Test generator",
							Description:   "This is a test generator",
							DirectoryName: "test",
						},
						&model.Generator{
							ID:            "test-generator2",
							Name:          "Test generator2",
							Description:   "This is a test generator 2",
							DirectoryName: "test2",
						},
					},
				},
			},
			&model.Template{
				ID: "template-id",
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r, index, clean := newTestRepository(t)
			defer clean()
			id := uuid.NewV4().String()
			tt.template.IID = id
			err := index.Index(id, tt.template)

			if err != nil {
				t.Error("Failed to index template to update", err)
			}

			if err := r.Update(tt.args.template); (err != nil) != tt.wantErr {
				t.Errorf("bleeveRepository.Update() error = %v, wantErr %v", err, tt.wantErr)
			}

			doc, err := index.Document(id)

			if doc == nil {
				t.Error("failed to retreive indexed document with ID", id)
			}

			if err != nil {
				t.Error("failed to retreive indexed document", tt.template, err)
			}

			for _, field := range doc.Fields {

				value := string(field.Value())
				switch field.Name() {
				case "iid":
					if string(value) == "" || (value != id) {
						t.Errorf("bleveRepository.Update() IID = %v want %v", value, id)
					}
				case "id":
					if string(value) == "" || (value != tt.args.template.ID) {
						t.Errorf("bleveRepository.Update() templateID = %v want %v", value, tt.args.template.ID)
					}
				case "name":
					if string(value) == "" || (value != tt.args.template.Name) {
						t.Errorf("bleveRepository.Update() templateName = %v want %v", value, tt.args.template.Name)
					}
				case "description":
					if string(value) == "" || (value != tt.args.template.Description) {
						t.Errorf("bleveRepository.Update() templateDescription = %v want %v", value, tt.args.template.Description)
					}
				case "directory_name":
					if string(value) == "" || (value != tt.args.template.DirectoryName) {
						t.Errorf("bleveRepository.Update() templateDirectoryName = %v want %v", value, tt.args.template.DirectoryName)
					}
				case "generators.id":
					pos := field.ArrayPositions()[0]
					expectedID := tt.args.template.Generators[pos].ID
					if value != expectedID {
						t.Errorf("bleveRepository.Update() template.Generators[%d].ID = %v want %v", pos, value, expectedID)
					}
				case "generators.name":
					pos := field.ArrayPositions()[0]
					expectedName := tt.args.template.Generators[pos].Name
					if value != expectedName {
						t.Errorf("bleveRepository.Update() template.Generators[%d].Name = %v want %v", pos, value, expectedName)
					}
				case "generators.description":
					pos := field.ArrayPositions()[0]
					expectedDescription := tt.args.template.Generators[pos].Description
					if value != expectedDescription {
						t.Errorf("bleveRepository.Update() template.Generators[%d].Description = %v want %v", pos, value, expectedDescription)
					}
				case "generators.directory_name":
					pos := field.ArrayPositions()[0]
					expectedDirectoryName := tt.args.template.Generators[pos].DirectoryName
					if value != expectedDirectoryName {
						t.Errorf("bleveRepository.Update() template.Generators[%d].DirectoryName = %v want %v", pos, value, expectedDirectoryName)
					}
				default:
					t.Error("doc.Fields should assert field", field.Name(), string(field.Value()))
				}
			}

		})
	}
}

func Test_bleeveRepository_FindTemplateByID(t *testing.T) {

	type args struct {
		ID string
	}
	tests := []struct {
		name    string
		args    args
		want    *model.Template
		wantErr bool
	}{
		{
			"Find by id",
			args{
				"test-template-id",
			},
			&model.Template{
				ID:          "test-template-id",
				Description: "Some description",
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r, index, clean := newTestRepository(t)
			defer clean()
			err := index.Index(uuid.NewV4().String(), tt.want)

			if err != nil {
				t.Errorf("Failed to index template model")
			}

			got, err := r.FindTemplateByID(tt.args.ID)

			if (err != nil) != tt.wantErr {
				t.Errorf("bleeveRepository.FindTemplateByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if got != nil && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("bleeveRepository.FindTemplateByID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_bleeveRepository_Delete(t *testing.T) {

	type args struct {
		ID string
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{
			"Delete template from index",
			args{
				"template-id",
			},
			true,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r, index, clean := newTestRepository(t)
			defer clean()
			id := uuid.NewV4().String()
			templ := &model.Template{
				IID: id,
				ID:  "template-id",
			}
			err := index.Index(id, templ)

			if err != nil {
				t.Error("Failed to index template to update", err)
			}
			got, err := r.Delete(tt.args.ID)
			if (err != nil) != tt.wantErr {
				t.Errorf("bleeveRepository.Delete() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("bleeveRepository.Delete() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_bleeveRepository_List(t *testing.T) {

	tests := []struct {
		name    string
		want    []*model.Template
		wantErr bool
	}{
		{
			"List all templates",
			[]*model.Template{
				&model.Template{ID: "template-id1"},
				&model.Template{ID: "template-id2"},
				&model.Template{ID: "template-id3"},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r, index, clean := newTestRepository(t)
			defer clean()

			for _, templ := range tt.want {
				templ.IID = uuid.NewV4().String()
				err := index.Index(templ.IID, templ)

				if err != nil {
					t.Error("Failed to index template to list", err)
				}
			}
			got, err := r.List()
			if (err != nil) != tt.wantErr {
				t.Errorf("bleeveRepository.List() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			for _, templ := range got {
				if !containsTemplate(templ, got) {
					t.Errorf("bleeveRepository.List() = %v, want %v", got, templ)
				}

			}

		})
	}
}

func containsTemplate(templ *model.Template, templates []*model.Template) bool {
	for _, ltempl := range templates {
		if reflect.DeepEqual(templ, ltempl) {
			return true
		}
	}
	return false
}

func Test_bleeveRepository_Exists(t *testing.T) {

	type args struct {
		ID string
	}
	tests := []struct {
		name string

		args    args
		want    bool
		wantErr bool
	}{
		{"Template exists", args{"template-id"}, true, false},
		{"Template exists", args{"template-id-dont-exists"}, false, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r, index, clean := newTestRepository(t)
			defer clean()
			id := uuid.NewV4().String()
			templ := &model.Template{
				IID: id,
				ID:  "template-id",
			}
			err := index.Index(id, templ)

			if err != nil {
				t.Error("Failed to index template to verify existence", err)
			}
			got, err := r.Exists(tt.args.ID)
			if (err != nil) != tt.wantErr {
				t.Errorf("bleeveRepository.Exists() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("bleeveRepository.Exists() = %v, want %v", got, tt.want)
			}
		})
	}
}
