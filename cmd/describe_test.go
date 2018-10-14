package cmd

import (
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	testhelpers "github.com/ironman-project/ironman/cmd/testing"
	"github.com/ironman-project/ironman/pkg/ironman"
	"github.com/ironman-project/ironman/pkg/testutils"
	"github.com/spf13/cobra"
)

//Pre-installs a template for running tests
func setUpDescribe(t *testing.T, client *ironman.Ironman, testCase testhelpers.CmdTestCase) {
	installCmd := newInstallCommand(client, ioutil.Discard)
	//equivalent to "ironman install https://github.com/ironman-project/template-example.git"
	args := []string{"https://github.com/ironman-project/template-example.git"}
	err := testhelpers.RunTestCmd(installCmd, args, nil)
	if err != nil {
		t.Fatalf("failed to setUp describe tests %s", err)
	}
}

func TestDescribeCmd(t *testing.T) {
	tempGenerateDir := testutils.CreateTempDir("temp-generate", t)
	defer func() {
		_ = os.RemoveAll(tempGenerateDir)
	}()
	tests := []testhelpers.CmdTestCase{
		{
			Name:     "describe a template with default yaml format",
			Args:     []string{"template-example"},
			Flags:    []string{""},
			Expected: testutils.ReadFile(t, filepath.Join("testing", "expected_describe.yaml")),
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
		return newDescribeCmd(client, out)
	}, setUpDescribe, nil)

}
