package cmd

import (
	"bytes"
	"io"
	"os"
	"regexp"
	"testing"

	"github.com/spf13/cobra"

	"github.com/ironman-project/ironman/pkg/ironman"

	"github.com/ironman-project/ironman/pkg/testutils"
)

type cmdTestCase struct {
	name     string
	args     []string
	flags    []string
	expected string
	err      bool
}

type testCmdFactory func(ironman *ironman.Ironman, out io.Writer) *cobra.Command

func TestGenerateCmd(t *testing.T) {

	tests := []cmdTestCase{
		{
			"successful generate",
			[]string{"valid", "test-gen"},
			[]string{""},
			"",
			false,
		},
	}
	runCmdTests(t, tests, func(client *ironman.Ironman, out io.Writer) *cobra.Command {
		return newGenerateCommand(client, out)
	})

}

func runCmdTests(t *testing.T, tests []cmdTestCase, cmdFactory testCmdFactory) {
	var buf bytes.Buffer
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tempHome := testutils.CreateTempDir("ihome", t)
			testutils.CopyDir("testing/ihome", tempHome, t)
			client := ironman.New(tempHome)
			defer func() {
				err := os.RemoveAll(tempHome)
				if err != nil {
					t.Fatal("Failed to remove test ironman home", err)
				}

			}()
			cmd := cmdFactory(client, &buf)
			cmd.ParseFlags(tt.flags)
			err := cmd.RunE(cmd, tt.args)
			if (err != nil) != tt.err {
				t.Errorf("expected error, got '%v'", err)
			}
			re := regexp.MustCompile(tt.expected)
			if !re.Match(buf.Bytes()) {
				t.Errorf("expected\n%q\ngot\n%q", tt.expected, buf.String())
			}
			buf.Reset()
		})
	}
}
