package template

import (
	"bytes"
	"context"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/ironman-project/ironman/pkg/template/engine"
	"github.com/ironman-project/ironman/pkg/template/engine/goengine"
	"github.com/ironman-project/ironman/pkg/template/model"
	"github.com/ironman-project/ironman/pkg/template/values"
	"github.com/ironman-project/ironman/pkg/testutils"
)

func engineFactory() engine.Engine {
	return goengine.New("test_valid")
}

type fileResult struct {
	relativePath string
	contents     string
}

func Test_generator_Generate(t *testing.T) {
	type fields struct {
		path           string
		data           GeneratorData
		generationPath string
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		wantFiles []fileResult
		wantErr   bool
	}{
		{
			"Generate template with directory generator",
			fields{
				path: filepath.Join("testing", "templates", "valid", "app"),
				data: GeneratorData{
					&model.Template{
						Name: "test",
					},
					&model.Generator{
						Name: "app",
					},
					values.Values{
						"foo": "bar",
						"bar": "foo",
					},
				},
			},
			args{context.Background()},
			[]fileResult{
				fileResult{
					relativePath: "hi.js",
					contents:     testutils.ReadFile(t, "testing", "expected", "templates", "valid", "app", "hi.js"),
				},
				fileResult{
					relativePath: "internal/hi.js",
					contents:     testutils.ReadFile(t, "testing", "expected", "templates", "valid", "app", "hi.js"),
				},
			},
			false,
		},
		{
			"Generate template with file generator relative path",
			fields{
				path: filepath.Join("testing", "templates", "valid", "controller"),
				data: GeneratorData{
					&model.Template{
						Name: "test",
					},
					&model.Generator{
						Name:  "controller",
						TType: model.GeneratorTypeFile,
						FileTypeOptions: model.FileTypeOptions{
							DefaultTemplateFile: "Controller.java",
						},
					},
					values.Values{
						"Name": "Foo",
					},
				},
				generationPath: "NewController.java",
			},
			args{context.Background()},
			[]fileResult{
				fileResult{
					relativePath: "NewController.java",
					contents:     testutils.ReadFile(t, "testing", "expected", "templates", "valid", "controller", "Controller.java"),
				},
			},
			false,
		},
		{
			"Generate template with file generator on internal directory",
			fields{
				path: filepath.Join("testing", "templates", "valid", "controller"),
				data: GeneratorData{
					&model.Template{
						Name: "test",
					},
					&model.Generator{
						Name:  "controller",
						TType: model.GeneratorTypeFile,
						FileTypeOptions: model.FileTypeOptions{
							DefaultTemplateFile:        "Controller.java",
							FileGenerationRelativePath: "controllers",
						},
					},
					values.Values{
						"Name": "Foo",
					},
				},
				generationPath: "NewController.java",
			},
			args{context.Background()},
			[]fileResult{
				fileResult{
					relativePath: "controllers/NewController.java",
					contents:     testutils.ReadFile(t, "testing", "expected", "templates", "valid", "controller", "Controller.java"),
				},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tempDir := testutils.CreateTempDir("test_valid", t)
			generationDir := filepath.Join(tempDir, filepath.Dir(tt.fields.generationPath))
			generationPath := filepath.Join(tempDir, tt.fields.generationPath)
			_ = os.MkdirAll(generationDir, os.ModePerm)
			defer func() {
				_ = os.RemoveAll(tempDir)

			}()

			g := NewGenerator(
				tt.fields.path,
				generationPath,
				tt.fields.data,
				SetGeneratorOutput(ioutil.Discard),
			)
			if err := g.Generate(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("generator.Generate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			for _, wantFile := range tt.wantFiles {
				file, err := ioutil.ReadFile(filepath.Join(tempDir, wantFile.relativePath))
				if err != nil {
					t.Errorf("generator.Generate() error = %v file should exists", wantFile.relativePath)
					continue
				}

				if string(file) != wantFile.contents {
					t.Errorf("Generator.Generate() \ncontents\n %s\n want \n%s\n", string(file), wantFile.contents)
				}
			}
		})
	}
}

func Test_generator_runHooks(t *testing.T) {
	type fields struct {
		data GeneratorData
	}
	type args struct {
		name  string
		hooks []*model.Command
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantOutput string
		wantErr    bool
	}{
		{
			name: "Run Hooks Including templates",
			fields: fields{
				data: GeneratorData{
					Template: &model.Template{
						ID:      "test",
						Version: "1.0",
						Name:    "test",
					},
					Generator: &model.Generator{
						ID: "gen-test",
					},
					Values: map[string]interface{}{
						"somevalue": "value",
					},
				},
			},
			args: args{
				name: "pre-generate",
				hooks: []*model.Command{
					&model.Command{
						Name: "echo",
						Args: []string{
							"-n",
							"Template {{.Template.ID}} Version {{.Template.Version}} Generator {{.Generator.ID}} with values with some-value {{.Values.somevalue}}",
						},
					},
				},
			},
			wantOutput: "Running pre-generate hooks\nTemplate test Version 1.0 Generator gen-test with values with some-value value\n...Running pre-generate hooks done\n",
			wantErr:    false,
		},
		{
			name: "Failed command",
			fields: fields{
				data: GeneratorData{},
			},
			args: args{
				name: "nonexistingcommand",
				hooks: []*model.Command{
					&model.Command{},
				},
			},
			wantOutput: "Running nonexistingcommand hooks\n",
			wantErr:    true,
		},
		{
			name: "Empty Command",
			fields: fields{
				data: GeneratorData{},
			},
			args: args{
				name: "",
				hooks: []*model.Command{
					&model.Command{},
				},
			},
			wantOutput: "Running  hooks\n",
			wantErr:    true,
		},
		{
			name: "Empty Hooks",
			fields: fields{
				data: GeneratorData{},
			},
			args: args{
				name:  "echo",
				hooks: nil,
			},
			wantOutput: "",
			wantErr:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			var output bytes.Buffer
			g := &generator{
				data:          tt.fields.data,
				engineFactory: engineFactory,
				out:           &output,
			}

			if err := g.runHooks(tt.args.name, tt.args.hooks); (err != nil) != tt.wantErr {
				t.Errorf("generator.runHooks() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if gotOutput := output.String(); gotOutput != tt.wantOutput {
				t.Errorf("generator.runHooks() = %v, want %v", gotOutput, tt.wantOutput)
			}

		})
	}
}
