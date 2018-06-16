package testing

import (
	"bytes"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"testing"

	"github.com/ironman-project/ironman/pkg/ironman"
	"github.com/ironman-project/ironman/pkg/testutils"
	"github.com/spf13/cobra"
)

//CmdTestCase represents a command test case
type CmdTestCase struct {
	Name     string
	Args     []string
	Flags    []string
	Expected string
	Err      bool
}

type testCmdFactory func(ironman *ironman.Ironman, out io.Writer) *cobra.Command

//CmdTestCaseSetUpTearDown function for setting up or tearing down command tests
type CmdTestCaseSetUpTearDown func(*testing.T, *ironman.Ironman, CmdTestCase)

//RunCmdTests runs commands test cases
func RunCmdTests(t *testing.T, tests []CmdTestCase, cmdFactory testCmdFactory, setUp CmdTestCaseSetUpTearDown, tearDown CmdTestCaseSetUpTearDown) {
	var buf bytes.Buffer
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			tempHome := testutils.CreateTempDir("ihome", t)
			testutils.CreateDir(filepath.Join(tempHome, "templates"), t)
			client := ironman.New(tempHome)
			defer func() {
				_ = os.RemoveAll(tempHome)
			}()

			if setUp != nil {
				setUp(t, client, tt)
			}
			if tearDown != nil {
				defer tearDown(t, client, tt)
			}

			cmd := cmdFactory(client, &buf)
			err := RunTestCmd(cmd, t, tt.Args, tt.Flags)

			if (err != nil) != tt.Err {
				t.Errorf("expected error, got '%v'", err)
			}

			//In case theres an error, the ouput of the error is expeted to be something specific
			if err != nil {
				re := regexp.MustCompile(tt.Expected)
				if !re.Match([]byte(err.Error())) {
					t.Errorf("expected\n%q\ngot\n%q", tt.Expected, err)
				}
				return
			}

			re := regexp.MustCompile(tt.Expected)
			if !re.Match(buf.Bytes()) {
				t.Errorf("expected\n%q\ngot\n%q", tt.Expected, buf.String())
			}
			buf.Reset()
		})
	}
}

//RunTestCmd runs a test command
func RunTestCmd(cmd *cobra.Command, t *testing.T, args []string, flags []string) error {
	err := cmd.ParseFlags(flags)
	if err != nil {
		return err
	}

	err = cmd.ValidateArgs(args)
	if err != nil {
		return err
	}

	err = cmd.RunE(cmd, args)
	return err
}
