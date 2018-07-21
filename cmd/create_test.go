package cmd

import (
	"io"
	"os"
	"path/filepath"
	"testing"

	"github.com/spf13/cobra"

	"github.com/ironman-project/ironman/pkg/ironman"

	testhelpers "github.com/ironman-project/ironman/cmd/testing"
	"github.com/ironman-project/ironman/pkg/testutils"
)

func TestCreateCmd(t *testing.T) {
	tempCreateDir := testutils.CreateTempDir("temp-generate", t)
	defer func() {
		_ = os.RemoveAll(tempCreateDir)
	}()
	tests := []testhelpers.CmdTestCase{
		{
			Name:     "successful generate",
			Args:     []string{filepath.Join(tempCreateDir, "test-gen")},
			Flags:    []string{""},
			Expected: "Creating new template",
			Err:      false,
		},
		{
			Name:     "template path is required",
			Args:     []string{},
			Flags:    []string{},
			Expected: "template path is required",
			Err:      true,
		},
	}
	testhelpers.RunCmdTests(t, tests, func(client *ironman.Ironman, out io.Writer) *cobra.Command {
		return newCreateCmd(client, out)
	}, nil, nil)

}
