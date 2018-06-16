package cmd

import (
	"io"
	"os"
	"testing"

	testhelpers "github.com/ironman-project/ironman/cmd/testing"
	"github.com/ironman-project/ironman/pkg/ironman"
	"github.com/ironman-project/ironman/pkg/testutils"
	"github.com/spf13/cobra"
)

func TestVersionCmd(t *testing.T) {

	tempVersionDir := testutils.CreateTempDir("temp-generate", t)
	defer func() {
		_ = os.RemoveAll(tempVersionDir)
	}()
	tests := []testhelpers.CmdTestCase{
		{
			Name:     "successful generate",
			Args:     []string{},
			Flags:    []string{""},
			Expected: "Ironman 0.1.0-xyz Build date: 2087-08-01",
			Err:      false,
		},
	}
	testhelpers.RunCmdTests(t, tests, func(client *ironman.Ironman, out io.Writer) *cobra.Command {
		return newVersionCmd("0.1.0", "xyz", "2087-08-01", out)
	}, nil, nil)

}
