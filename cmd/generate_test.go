package cmd

import (
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/spf13/cobra"

	"github.com/ironman-project/ironman/pkg/ironman"

	testhelpers "github.com/ironman-project/ironman/cmd/testing"
	"github.com/ironman-project/ironman/pkg/testutils"
)

//Pre-installs a template for running tests
func setUpGenerateCmd(t *testing.T, client *ironman.Ironman, testCase testhelpers.CmdTestCase) {
	installCmd := newInstallCommand(client, ioutil.Discard)
	//equivalent to "ironman install https://github.com/ironman-project/template-example.git"
	args := []string{"https://github.com/ironman-project/template-example.git"}
	testhelpers.RunTestCmd(installCmd, t, args, nil)
}

func TestGenerateCmd(t *testing.T) {
	tempGenerateDir := testutils.CreateTempDir("temp-generate", t)
	defer func() {
		_ = os.RemoveAll(tempGenerateDir)
	}()
	tests := []testhelpers.CmdTestCase{
		{
			Name:     "successful generate",
			Args:     []string{"template-example", filepath.Join(tempGenerateDir, "test-gen")},
			Flags:    []string{""},
			Expected: "Running template generator app",
			Err:      false,
		},
		{
			Name:     "successful generate with parameters",
			Args:     []string{"template-example", filepath.Join(tempGenerateDir, "test-gen-with-parameters")},
			Flags:    []string{"--set", "key=value"},
			Expected: "Running template generator app",
			Err:      false,
		},
		{
			Name:     "template id required",
			Args:     []string{},
			Flags:    []string{},
			Expected: "",
			Err:      true,
		},
		{
			Name:     "successful generate with hooks",
			Args:     []string{"template-example:with_hooks", filepath.Join(tempGenerateDir, "test-gen-hooks")},
			Flags:    []string{""},
			Expected: "Running template generator with_hooks\nRunning pre-generate hooks",
			Err:      false,
		},
	}
	testhelpers.RunCmdTests(t, tests, func(client *ironman.Ironman, out io.Writer) *cobra.Command {
		return newGenerateCommand(client, out)
	}, setUpGenerateCmd, nil)

}
