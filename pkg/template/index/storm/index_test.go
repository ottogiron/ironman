package storm

import (
	"io/ioutil"
	"path/filepath"
	"reflect"
	"testing"
	"time"

	"github.com/ironman-project/ironman/pkg/template/model"
	"github.com/ironman-project/ironman/pkg/testutils"
)

func tempIndexPath(t *testing.T) string {
	dir, err := ioutil.TempDir("", "ironman-storm-test")
	if err != nil {
		t.Fatal("failed to create test bleve index directory", err)
	}

	indexPath := filepath.Join(dir, "index")
	return indexPath
}

func TestIndex_Index(t *testing.T) {
	type args struct {
		model *model.Template
	}
	tests := []struct {
		name    string
		args    args
		want    string
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
					CreatedAt:   time.Now(),
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
			"test-template",
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			path := tempIndexPath(t)
			dbFactory := DefaultDBFactory(path)
			i := New(dbFactory)
			got, err := i.Index(tt.args.model)
			if (err != nil) != tt.wantErr {
				t.Errorf("Index.Index() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Index.Index() = %v, want %v", got, tt.want)
			}

			//fetch stored object
			db, _ := dbFactory()
			defer db.Close()
			var stored model.Template
			err = db.One("ID", tt.args.model.ID, &stored)

			if (err != nil) != tt.wantErr {
				t.Errorf("Index.Index() error = %v, wantErr %v", err, tt.wantErr)
				return

			}
			//Marshal to compare
			marshalGot := testutils.Marshal(stored, t)
			marghalWant := testutils.Marshal(tt.args.model, t)

			if marshalGot != marghalWant {
				t.Errorf("Index.Index() template = %s want %s", marshalGot, marghalWant)
			}

		})
	}
}

func TestIndex_Update(t *testing.T) {
	type args struct {
		model *model.Template
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "",
			args: args{
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
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			path := tempIndexPath(t)
			dbFactory := DefaultDBFactory(path)
			i := New(dbFactory)

			//make independent db access for each transaction
			func() {
				//First store the template to update
				db, err := dbFactory()
				if (err != nil) != tt.wantErr {
					t.Errorf("Index.Update() error = %v, wantErr %v", err, tt.wantErr)
				}
				defer db.Close()
				err = db.Save(tt.args.model)
				if (err != nil) != tt.wantErr {
					t.Errorf("Index.Update() error = %v, wantErr %v", err, tt.wantErr)
					return
				}

			}()

			//Actually update the template
			tt.args.model.Version = "0.2.0"
			if err := i.Update(tt.args.model); (err != nil) != tt.wantErr {
				t.Errorf("Index.Update() error = %v, wantErr %v", err, tt.wantErr)
			}

			func() {
				//Check the stored object matches the updated in memory "Version" field
				db, _ := dbFactory()
				defer db.Close()
				var stored model.Template
				err := db.One("ID", tt.args.model.ID, &stored)
				if (err != nil) != tt.wantErr {
					t.Errorf("Index.Update() error = %v, wantErr %v", err, tt.wantErr)
					return
				}

				marshalGot := testutils.Marshal(stored, t)
				marshalWant := testutils.Marshal(tt.args.model, t)

				if marshalGot != marshalWant {
					t.Errorf("Index.Update() template = %s, want %s", marshalGot, marshalWant)
				}
			}()

		})
	}
}

func TestIndex_Delete(t *testing.T) {
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
			path := tempIndexPath(t)
			dbFactory := DefaultDBFactory(path)
			i := New(dbFactory)

			func() {
				db, _ := dbFactory()
				defer db.Close()
				template := model.Template{ID: tt.args.ID}
				err := db.Save(&template)
				if (err != nil) != tt.wantErr {
					t.Errorf("Index.Delete() error = %v, wantErr %v", err, tt.wantErr)
					return
				}

			}()
			got, err := i.Delete(tt.args.ID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Index.Delete() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Index.Delete() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIndex_List(t *testing.T) {
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
			path := tempIndexPath(t)
			dbFactory := DefaultDBFactory(path)
			i := New(dbFactory)

			func() {
				db, _ := dbFactory()
				defer db.Close()
				for _, template := range tt.want {
					err := db.Save(template)
					if (err != nil) != tt.wantErr {
						t.Errorf("Index.List() error = %v, wantErr %v", err, tt.wantErr)
						break
					}
				}
			}()

			got, err := i.List()
			if (err != nil) != tt.wantErr {
				t.Errorf("Index.List() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Index.List() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIndex_FindTemplateByID(t *testing.T) {
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
			name: "Find template by ID",
			args: args{
				ID: "search-template",
			},
			want: &model.Template{
				ID: "search-template",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			path := tempIndexPath(t)
			dbFactory := DefaultDBFactory(path)
			i := New(dbFactory)

			func() {
				db, _ := dbFactory()
				defer db.Close()
				template := &model.Template{ID: tt.args.ID}
				err := db.Save(template)
				if (err != nil) != tt.wantErr {
					t.Errorf("Index.FindTemplateByID() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
			}()
			got, err := i.FindTemplateByID(tt.args.ID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Index.FindTemplateByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Index.FindTemplateByID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIndex_Exists(t *testing.T) {
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
			name: "Template Exists",
			args: args{
				ID: "exits-test",
			},
			want:    true,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			path := tempIndexPath(t)

			dbFactory := DefaultDBFactory(path)
			i := New(dbFactory)

			func() {
				db, _ := dbFactory()
				defer db.Close()
				template := &model.Template{ID: tt.args.ID}
				err := db.Save(template)
				if (err != nil) != tt.wantErr {
					t.Errorf("Index.Exists() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
			}()

			got, err := i.Exists(tt.args.ID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Index.Exists() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Index.Exists() = %v, want %v", got, tt.want)
			}
		})
	}
}
