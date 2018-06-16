package cmd

import (
	"io"
	"os"
	"path/filepath"
	"testing"

	testhelpers "github.com/ironman-project/ironman/cmd/testing"
	"github.com/ironman-project/ironman/pkg/ironman"
	"github.com/ironman-project/ironman/pkg/testutils"
	"github.com/spf13/cobra"
)

func TestLinkCmd(t *testing.T) {
	tempGenerateDir := testutils.CreateTempDir("temp-generate", t)
	defer func() {
		_ = os.RemoveAll(tempGenerateDir)
	}()
	tests := []testhelpers.CmdTestCase{
		{
			Name:     "link a template",
			Args:     []string{filepath.Join("testing", "templates", "linkable-template"), "test-link"},
			Flags:    []string{""},
			Expected: "Linking template to repository with ID test-link...",
			Err:      false,
		},
		{
			Name:     "required template path and symlink name",
			Args:     []string{},
			Flags:    []string{},
			Expected: "template path and symlink name are required",
			Err:      true,
		},
		{
			Name:     "required template path and symlink name",
			Args:     []string{"non-existing-path", "test-link"},
			Flags:    []string{},
			Expected: "failed to create symlink to ironman manager ",
			Err:      true,
		},
	}
	testhelpers.RunCmdTests(t, tests, func(client *ironman.Ironman, out io.Writer) *cobra.Command {
		return newLinkCmd(client, out)
	}, nil, nil)

}
