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
func setUpUpdateCmd(t *testing.T, client *ironman.Ironman, testCase testhelpers.CmdTestCase) {
	installCmd := newInstallCommand(client, ioutil.Discard)
	//equivalent to "ironman install https://github.com/ironman-project/template-example.git"
	args := []string{"https://github.com/ironman-project/template-example.git"}
	testhelpers.RunTestCmd(installCmd, t, args, nil)
}

func TestUpdateCmd(t *testing.T) {
	tempUpdateDir := testutils.CreateTempDir("temp-generate", t)
	defer func() {
		_ = os.RemoveAll(tempUpdateDir)
	}()
	tests := []testhelpers.CmdTestCase{
		{
			Name:     "successful generate",
			Args:     []string{"template-example"},
			Flags:    []string{""},
			Expected: "Updating template template-example",
			Err:      false,
		},
		{
			Name:     "template id required",
			Args:     []string{},
			Flags:    []string{},
			Expected: "",
			Err:      true,
		},
	}
	testhelpers.RunCmdTests(t, tests, func(client *ironman.Ironman, out io.Writer) *cobra.Command {
		return newUpdateCmd(client, out)
	}, setUpGenerateCmd, nil)

}
