package acceptance

/**
Used to initialize and clean variables common to all acceptance tests
**/
import (
	"os"
	"path/filepath"

	"github.com/DATA-DOG/godog"

	homedir "github.com/mitchellh/go-homedir"
)

var ironmanTestDir string
var ironmanTemplatesDir string

func init() {
	var err error
	ironmanTestDir, err = homedir.Dir()

	if err != nil {
		os.Exit(-1)
	}
	ironmanTestDir = filepath.Join(ironmanTestDir, ".ironman_test")
	ironmanTemplatesDir = filepath.Join(ironmanTestDir, "templates")
}

//VarsContext context for vars file
func VarsContext(s *godog.Suite) {

	s.AfterSuite(func() {
		_ = os.RemoveAll(ironmanTestDir)
	})
}
