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
func setUpListCmd(t *testing.T, client *ironman.Ironman, testCase testhelpers.CmdTestCase) {
	installCmd := newInstallCommand(client, ioutil.Discard)
	//equivalent to "ironman install https://github.com/ironman-project/template-example.git"
	args := []string{"https://github.com/ironman-project/template-example.git"}
	testhelpers.RunTestCmd(installCmd, t, args, nil)
}

func TestListCmd(t *testing.T) {
	tempGenerateDir := testutils.CreateTempDir("temp-generate", t)
	defer func() {
		_ = os.RemoveAll(tempGenerateDir)
	}()

	//Tests with pre installed template
	tests := []testhelpers.CmdTestCase{
		{
			Name:     "List existing templates",
			Args:     []string{},
			Flags:    []string{""},
			Expected: "Installed templates",
			Err:      false,
		},
	}
	testhelpers.RunCmdTests(t, tests, func(client *ironman.Ironman, out io.Writer) *cobra.Command {
		return newListCmd(client, out)
	}, setUpListCmd, nil)

}

func TestListCmdNoPreinstalled(t *testing.T) {

	tempGenerateDir := testutils.CreateTempDir("temp-generate", t)
	defer func() {
		_ = os.RemoveAll(tempGenerateDir)
	}()

	//Tests without pre installed templates
	testsNoInstall := []testhelpers.CmdTestCase{
		{
			Name:     "List existing templates",
			Args:     []string{},
			Flags:    []string{""},
			Expected: "Installed templates\nNone\n",
			Err:      false,
		},
	}
	testhelpers.RunCmdTests(t, testsNoInstall, func(client *ironman.Ironman, out io.Writer) *cobra.Command {
		return newListCmd(client, out)
	}, nil, nil)

}
