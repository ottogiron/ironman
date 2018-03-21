package acceptance

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/DATA-DOG/godog"
	"github.com/ironman-project/ironman/testutils"
	"github.com/rendon/testcli"
)

func itRunsWithCorrectURL(URL string) error {
	testcli.Run(testutils.ExecutablePath(), "install", "--ironman-home="+ironmanTestDir, URL)
	return nil
}

func theProcessStateShouldBeSuccess() error {
	if !testcli.Success() {
		return fmt.Errorf("Install command did not succeeded %s %s", testcli.Error(), testcli.Stderr())
	}
	return nil
}

func aTemplateShouldBeInstalledWithID(id string) error {
	path := filepath.Join(ironmanTestDir, "templates", id)
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return fmt.Errorf("%s does not exists", path)
	}
	return nil
}

func theOutputShouldContainAnd(out1, out2 string) error {
	if !(testcli.StdoutContains(out1) && testcli.StdoutContains(out2)) {
		return fmt.Errorf("output => %s", testcli.Stdout())
	}
	return nil
}

func itRunsWithUnreachableURL(URL string) error {
	testcli.Run(testutils.ExecutablePath(), "install", "--ironman-home="+ironmanTestDir, URL)
	return nil
}

func theProcessStateShouldBeFailure() error {
	if !testcli.Failure() {
		return fmt.Errorf("Install command did not failed %s", testcli.Stdout())
	}
	return nil
}

func theOutputShouldCointain(expectedOutput string) error {
	if !strings.Contains(testcli.Stderr(), expectedOutput) {
		return fmt.Errorf("output => %s", testcli.Stderr())
	}
	return nil
}

//InstallContext context for install command
func InstallContext(s *godog.Suite) {
	s.Step(`^It runs with correct URL "([^"]*)"$`, itRunsWithCorrectURL)
	s.Step(`^The process state should be success$`, theProcessStateShouldBeSuccess)
	s.Step(`^A template should be installed with ID "([^"]*)"$`, aTemplateShouldBeInstalledWithID)
	s.Step(`^The output should contain "([^"]*)" and "([^"]*)"$`, theOutputShouldContainAnd)
	s.Step(`^It runs with unreachable URL "([^"]*)"$`, itRunsWithUnreachableURL)
	s.Step(`^The process state should be failure$`, theProcessStateShouldBeFailure)
	s.Step(`^The output should cointain "([^"]*)"$`, theOutputShouldCointain)

}
