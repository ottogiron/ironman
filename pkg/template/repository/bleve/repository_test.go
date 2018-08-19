package bleve

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"strconv"
	"testing"

	"github.com/ironman-project/ironman/pkg/testutils"

	"github.com/blevesearch/bleve"
	"github.com/ironman-project/ironman/pkg/template/model"
	"github.com/ironman-project/ironman/pkg/template/repository"
	uuid "github.com/satori/go.uuid"
)

func tempIndexPath(t *testing.T) string {
	dir, err := ioutil.TempDir("", "ironman-bleve-test")
	if err != nil {
		t.Fatal("failed to create test bleve index directory", err)
	}

	indexPath := filepath.Join(dir, "index")
	return indexPath
}

func newTestRepository(t *testing.T) (repository.Repository, bleve.Index, func()) {
	path := tempIndexPath(t)

	index, err := BuildIndex(path)

	if err != nil {
		t.Fatalf("failed to open test index %s", err)
	}

	r := New(SetIndex(index))

	return r, index, func() {
		err := index.Close()

		if err != nil {
			t.Fatal("failed to close bleve index", err)
		}
		err = os.RemoveAll(path)
		if err != nil {
			t.Fatal("failed to clean bleve index", err)
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
					SourceType:  model.SourceTypeURL,
					Name:        "Test template",
					Description: "This is a test template",
					HomeURL:     "https://ironman-project.io",
					Sources: []string{
						"https://github.com/ironman-project/ironman",
						"https://ironman-project.io",
					},
					Mantainers: []*model.Mantainer{
						&model.Mantainer{Name: "Otto Giron", Email: "otto"},
						&model.Mantainer{Name: "Gepser Hoil", Email: "jepzer@gepser.com"},
					},
					Generators: []*model.Generator{
						&model.Generator{
							ID:          "test-generator",
							Name:        "Test generator",
							Description: "This is a test generator",
						},
						&model.Generator{
							ID:          "test-generator2",
							Name:        "Test generator2",
							Description: "This is a test generator2",
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

			gotJ := testutils.Marshal(got, t)
			wantJ := testutils.Marshal(tt.args.template, t)

			if got != nil && (gotJ != wantJ) {
				t.Errorf("bleeveRepository.FindTemplateByID() = \n%v, want \n%v", gotJ, wantJ)
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
					SourceType:    model.SourceTypeLink,
					Version:       "0.1.0",
					Name:          "Updated name",
					Description:   "Updated description",
					DirectoryName: "test",
					HomeURL:       "http://template.com",
					AppVersion:    "0.1.0",
					Sources: []string{
						"https://github.com/ironman-project/ironman",
						"https://ironman-project.io",
					},
					Mantainers: []*model.Mantainer{
						&model.Mantainer{Name: "Otto Giron", Email: "otto"},
						&model.Mantainer{Name: "Gepser Hoil", Email: "jepzer@gepser.com"},
					},
					Generators: []*model.Generator{
						&model.Generator{
							ID:            "test-generator",
							TType:         model.GeneratorTypeDirectory,
							Name:          "Test generator",
							Description:   "This is a test generator",
							DirectoryName: "test",
							FileTypeOptions: model.FileTypeOptions{
								DefaultTemplateFile:        "controller.go",
								FileGenerationRelativePath: "controllers",
							},
						},
						&model.Generator{
							ID:            "test-generator2",
							TType:         model.GeneratorTypeFile,
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
				t.Error("failed to index template to update", err)
			}

			if err := r.Update(tt.args.template); (err != nil) != tt.wantErr {
				t.Errorf("bleeveRepository.Update() error = %v, wantErr %v", err, tt.wantErr)
			}

			doc, err := index.Document(id)

			if doc == nil {
				t.Error("failed to retrieve indexed document with ID", id)
			}

			if err != nil {
				t.Error("failed to retrieve indexed document", tt.template, err)
			}

			for _, field := range doc.Fields {

				value := string(field.Value())
				switch field.Name() {
				case "iid":
					if string(value) == "" || (value != id) {
						t.Errorf("bleveRepository.Update() IID = %v want %v", value, id)
					}
				case "sourceType":
					if string(value) == "" || (value != string(tt.args.template.SourceType)) {
						t.Errorf("bleveRepository.Update() sourceType = %v want %v", value, id)
					}
				case "id":
					if string(value) == "" || (value != tt.args.template.ID) {
						t.Errorf("bleveRepository.Update() templateID = %v want %v", value, tt.args.template.ID)
					}
				case "version":
					if string(value) == "" || (value != tt.args.template.Version) {
						t.Errorf("bleveRepository.Update() templateVersion = %v want %v", value, tt.args.template.Version)
					}

				case "name":
					if string(value) == "" || (value != tt.args.template.Name) {
						t.Errorf("bleveRepository.Update() templateName = %v want %v", value, tt.args.template.Name)
					}
				case "description":
					if string(value) == "" || (value != tt.args.template.Description) {
						t.Errorf("bleveRepository.Update() templateDescription = %v want %v", value, tt.args.template.Description)
					}
				case "directoryName":
					if string(value) == "" || (value != tt.args.template.DirectoryName) {
						t.Errorf("bleveRepository.Update() templateDirectoryName = %v want %v", value, tt.args.template.DirectoryName)
					}
				case "home":
					if string(value) == "" || (value != tt.args.template.HomeURL) {
						t.Errorf("bleveRepository.Update() templateHomeURL = %v want %v", value, tt.args.template.HomeURL)
					}
				case "appVersion":
					if string(value) == "" || (value != tt.args.template.AppVersion) {
						t.Errorf("bleveRepository.Update() templateAppVersion = %v want %v", value, tt.args.template.AppVersion)
					}
				case "deprecated":
					boolValue, _ := strconv.ParseBool(value)
					if string(value) == "" || (boolValue != tt.args.template.Deprecated) {
						t.Errorf("bleveRepository.Update() templateDeprecated = %v want %v", value, tt.args.template.Deprecated)
					}

				case "generators.id":
					pos := field.ArrayPositions()[0]
					expectedID := tt.args.template.Generators[pos].ID
					if value != expectedID {
						t.Errorf("bleveRepository.Update() template.Generators[%d].ID = %v want %v", pos, value, expectedID)
					}
				case "generators.type":
					pos := field.ArrayPositions()[0]
					expectedType := tt.args.template.Generators[pos].TType
					if model.GeneratorType(value) != expectedType {
						t.Errorf("bleveRepository.Update() template.Generators[%d].Type = %v want %v", pos, value, expectedType)
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
				case "generators.directoryName":
					pos := field.ArrayPositions()[0]
					expectedDirectoryName := tt.args.template.Generators[pos].DirectoryName
					if value != expectedDirectoryName {
						t.Errorf("bleveRepository.Update() template.Generators[%d].DirectoryName = %v want %v", pos, value, expectedDirectoryName)
					}
				case "generators.fileTypeOptions.defaultTemplateFile":
					pos := field.ArrayPositions()[0]
					expectedFileTypeDefaultTemplateFile := tt.args.template.Generators[pos].FileTypeOptions.DefaultTemplateFile
					if value != expectedFileTypeDefaultTemplateFile {
						t.Errorf("bleveRepository.Update() template.Generators[%d].FileTypeDefaultTemplateFile = %v want %v", pos, value, expectedFileTypeDefaultTemplateFile)
					}
				case "generators.fileTypeOptions.fileGenerationRelativePath":
					pos := field.ArrayPositions()[0]
					expectedFileTypeFileGenerationRelativePath := tt.args.template.Generators[pos].FileTypeOptions.FileGenerationRelativePath
					if value != expectedFileTypeFileGenerationRelativePath {
						t.Errorf("bleveRepository.Update() template.Generators[%d].FileTypeFileGenerationRelativePath = %v want %v", pos, value, expectedFileTypeFileGenerationRelativePath)
					}
				case "sources":
					expectedSource := tt.args.template.Sources[field.ArrayPositions()[0]]
					if value != expectedSource {
						t.Errorf("bleveRepository.Update() template.Sources[index].expectedSource = %v want %v", value, expectedSource)

					}
				case "mantainers.name":
					expectedMantainerName := tt.args.template.Mantainers[field.ArrayPositions()[0]].Name
					if value != expectedMantainerName {
						t.Errorf("bleveRepository.Update() template.Sources[index].expectedMantainerName = %v want %v", value, expectedMantainerName)

					}
				case "mantainers.email":
					expectedMantainerEmail := tt.args.template.Mantainers[field.ArrayPositions()[0]].Email
					if value != expectedMantainerEmail {
						t.Errorf("bleveRepository.Update() template.Sources[index].expectedMantainerEmail = %v want %v", value, expectedMantainerEmail)

					}
				case "mantainers.url":
					expectedMantainerURL := tt.args.template.Mantainers[field.ArrayPositions()[0]].URL
					if value != expectedMantainerURL {
						t.Errorf("bleveRepository.Update() template.Sources[index].expectedMantainerURL = %v want %v", value, expectedMantainerURL)

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
				t.Errorf("failed to index template model")
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
				t.Error("failed to index template to update", err)
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
					t.Error("failed to index template to list", err)
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
				t.Error("failed to index template to verify existence", err)
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
