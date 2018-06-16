package testutils

import (
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

//CreateTempDir Creates a temporary directory with an specified name
func CreateTempDir(prefix string, t *testing.T) string {
	name, err := ioutil.TempDir("", prefix)
	if err != nil {
		t.Fatalf("Could not create temp directory %s %s", err, name)
	}
	return name
}

//CreateDir Creates a directory in the path
func CreateDir(name string, t *testing.T) {
	err := os.Mkdir(name, os.ModePerm)
	if err != nil {
		t.Fatalf("Could not create temp directory %s %s", err, name)
	}
}

//ReadFile reads a fixture from  a "fixtures" directory relative to the current file
func ReadFile(t *testing.T, path ...string) string {
	b, err := ioutil.ReadFile(filepath.Join(path...))
	if err != nil {
		t.Fatal(err)
	}
	return string(b)
}

//FileExists verifies if file exists
func FileExists(path string) bool {
	if _, err := os.Stat(path); err == nil {
		return true
	}
	return false
}

//CopyDir copies a directory recursively to a destination directory
func CopyDir(source string, dest string, t *testing.T) {

	// get properties of source dir
	sourceinfo, err := os.Stat(source)
	if err != nil {
		t.Fatal(err)
	}

	// create dest dir

	err = os.MkdirAll(dest, sourceinfo.Mode())
	if err != nil {
		t.Fatal(err)
	}

	directory, _ := os.Open(source)

	objects, err := directory.Readdir(-1)

	for _, obj := range objects {

		sourcefilepointer := source + "/" + obj.Name()

		destinationfilepointer := dest + "/" + obj.Name()

		if obj.IsDir() {
			// create sub-directories - recursively
			CopyDir(sourcefilepointer, destinationfilepointer, t)

		} else {
			// perform copy
			CopyFile(sourcefilepointer, destinationfilepointer, t)
		}

	}
	return
}

//CopyFile copies a source file to a destination file
func CopyFile(source string, dest string, t *testing.T) {
	sourcefile, err := os.Open(source)
	if err != nil {
		t.Fatal(err)
	}

	defer sourcefile.Close()

	destfile, err := os.Create(dest)
	if err != nil {
		t.Fatal(err)
	}

	defer destfile.Close()

	_, err = io.Copy(destfile, sourcefile)
	if err == nil {
		sourceinfo, err := os.Stat(source)
		if err != nil {
			err = os.Chmod(dest, sourceinfo.Mode())
		}

	}

	return
}
