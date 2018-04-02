package testutils

import (
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
