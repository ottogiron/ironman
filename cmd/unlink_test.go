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
func setUpUnlinkCmd(t *testing.T, client *ironman.Ironman, testCase testhelpers.CmdTestCase) {
	linkCmd := newLinkCmd(client, ioutil.Discard)
	//equivalent to "ironman install https://github.com/ironman-project/template-example.git"
	args := []string{filepath.Join("testing", "templates", "linkable-template"), "linked-template"}
	err := testhelpers.RunTestCmd(linkCmd, args, nil)
	if err != nil {
		t.Fatalf("failed to setUp link tests %s", err)
	}
}

func TestUnlinkCmd(t *testing.T) {
	tempGenerateDir := testutils.CreateTempDir("temp-generate", t)
	defer func() {
		_ = os.RemoveAll(tempGenerateDir)
	}()

	//Tests with pre installed template
	tests := []testhelpers.CmdTestCase{
		{
			Name:     "Unlink existing templates",
			Args:     []string{"linked-template"},
			Flags:    []string{""},
			Expected: "Unlinking template from repository with ID linked-template",
			Err:      false,
		},
		{
			Name:     "Unlink non existing ID",
			Args:     []string{},
			Flags:    []string{""},
			Expected: "Template ID is required",
			Err:      true,
		},
	}
	testhelpers.RunCmdTests(t, tests, func(client *ironman.Ironman, out io.Writer) *cobra.Command {
		return newUnlinkCmd(client, out)
	}, setUpUnlinkCmd, nil)

}
