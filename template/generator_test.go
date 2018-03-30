package template

import (
	"context"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/ironman-project/ironman/template/engine"
	"github.com/ironman-project/ironman/template/engine/goengine"
	"github.com/ironman-project/ironman/template/model"
	"github.com/ironman-project/ironman/template/values"
	"github.com/ironman-project/ironman/testutils"
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
		path string
		data GeneratorData
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
			"Generate template",
			fields{
				filepath.Join("testing", "templates", "valid", "app"),
				GeneratorData{
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
					contents: `//This a generated file from template test and generator app
console.log("Foo is bar bar is foo")`,
				},
				fileResult{
					relativePath: "internal/hi.js",
					contents: `//This a generated file from template test and generator app
console.log("Foo is bar bar is foo")`,
				},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tempDir := testutils.CreateTempDir("test_valid", t)
			defer func() {
				_ = os.RemoveAll(tempDir)
			}()
			g := NewGenerator(
				tt.fields.path,
				tempDir,
				[]string{".ironman.yaml"},
				tt.fields.data,
				engineFactory,
			)
			if err := g.Generate(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("generator.Generate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			for _, wantFile := range tt.wantFiles {
				file, err := ioutil.ReadFile(filepath.Join(tempDir, wantFile.relativePath))
				if err != nil {
					t.Errorf("generator.Generate() error = %v file should exists", wantFile.relativePath)
				}

				if string(file) != wantFile.contents {
					t.Errorf("Generator.Generate() \ncontents %s\n want \n%s\n", string(file), wantFile.contents)
				}
			}
		})
	}
}
