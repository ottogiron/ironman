package cmd

import (
	"io"
	"io/ioutil"
	"os"
	"testing"

	testhelpers "github.com/ironman-project/ironman/cmd/testing"
	"github.com/ironman-project/ironman/pkg/ironman"
	"github.com/ironman-project/ironman/pkg/testutils"
	"github.com/spf13/cobra"
)

//Pre-installs a template for running tests
func setUpUninstallCmd(t *testing.T, client *ironman.Ironman, testCase testhelpers.CmdTestCase) {
	installCmd := newInstallCommand(client, ioutil.Discard)
	//equivalent to "ironman install https://github.com/ironman-project/template-example.git"
	args := []string{"https://github.com/ironman-project/template-example.git"}
	testhelpers.RunTestCmd(installCmd, t, args, nil)
}

func TestUninstallCmd(t *testing.T) {
	tempGenerateDir := testutils.CreateTempDir("temp-generate", t)
	defer func() {
		_ = os.RemoveAll(tempGenerateDir)
	}()

	//Tests with pre installed template
	tests := []testhelpers.CmdTestCase{
		{
			Name:     "Uninstall existing templates",
			Args:     []string{"template-example"},
			Flags:    []string{""},
			Expected: "Uninstalling template template-example",
			Err:      false,
		},
		{
			Name:     "Uninstall non existing ID",
			Args:     []string{},
			Flags:    []string{""},
			Expected: "Template ID is required",
			Err:      true,
		},
	}
	testhelpers.RunCmdTests(t, tests, func(client *ironman.Ironman, out io.Writer) *cobra.Command {
		return newUninstallCmd(client, out)
	}, setUpUninstallCmd, nil)

}
