package testutils

import (
	"fmt"
	"io"
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

//CreateTempDir Creates a temporary directory with an specified name
func CreateTempDir(prefix string, t *testing.T) string {
	name, err := ioutil.TempDir("", prefix)
	if err != nil {
		t.Fatalf("Could not create temp directory %s %s", err, name)
	}
	return name
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

func CopyDir(source string, dest string) (err error) {

	// get properties of source dir
	sourceinfo, err := os.Stat(source)
	if err != nil {
		return err
	}

	// create dest dir

	err = os.MkdirAll(dest, sourceinfo.Mode())
	if err != nil {
		return err
	}

	directory, _ := os.Open(source)

	objects, err := directory.Readdir(-1)

	for _, obj := range objects {

		sourcefilepointer := source + "/" + obj.Name()

		destinationfilepointer := dest + "/" + obj.Name()

		if obj.IsDir() {
			// create sub-directories - recursively
			err = CopyDir(sourcefilepointer, destinationfilepointer)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			// perform copy
			err = CopyFile(sourcefilepointer, destinationfilepointer)
			if err != nil {
				fmt.Println(err)
			}
		}

	}
	return
}

func CopyFile(source string, dest string) (err error) {
	sourcefile, err := os.Open(source)
	if err != nil {
		return err
	}

	defer sourcefile.Close()

	destfile, err := os.Create(dest)
	if err != nil {
		return err
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

//FileExists verifies if file exists
func FileExists(path string) bool {
	if _, err := os.Stat(path); err == nil {
		return true
	}
	return false
}
