package testutils

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

//CreateTempFile create a new temporary file for tests
func CreateTempFile(contents string, t *testing.T) (string, func()) {
	path := CreateTempFileInDir(os.TempDir(), contents, t)
	return path, func() {
		os.Remove(path)
	}
}

//CreateTempFileInDir creates a file in a directory
func CreateTempFileInDir(dir string, contents string, t *testing.T) (path string) {
	f, err := ioutil.TempFile(dir, "test_file_")

	if err != nil {
		t.Fatalf("Couldn't create the temporary file %s", err)
	}

	path, err = getFilePath(f.Name())

	if err != nil {
		t.Fatalf("Couldn't create the temporary file %s", err)
	}

	b := []byte(contents)
	_, err = f.Write(b)

	if err != nil {
		t.Fatalf("Couldn't create the temporary file %s", err)
	}

	return
}

func getFilePath(filename string) (path string, err error) {
	path, err = filepath.Abs(filename)
	return
}

//ReadFixture reads a fixture from  a "fixtures" directory relative to the current file
func ReadFixture(name string, t *testing.T) string {
	b, err := ioutil.ReadFile("./fixtures/" + name + ".json")
	if err != nil {
		t.Fatal(err)
	}
	return string(b)
}
