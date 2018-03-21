package acceptance

import (
	"fmt"

	"github.com/DATA-DOG/godog"
	"github.com/ironman-project/ironman/testutils"
	"github.com/rendon/testcli"
)

func aTemplateIsInstalledWithURL(URL string) error {
	testcli.Run(testutils.ExecutablePath(), "install", "--ironman-home="+ironmanTestDir, URL)
	if !testcli.Success() {
		return fmt.Errorf("Template with url %s should be installed", URL)
	}
	return nil
}

func updateRunsWithCorrectTemplateID(templateID string) error {
	testcli.Run(testutils.ExecutablePath(), "update", "--ironman-home="+ironmanTestDir, templateID)
	return nil
}

func theUpdateProcessStateShouldBeSuccess() error {
	if !testcli.Success() {
		return fmt.Errorf("Update command did not succeeded %s %s", testcli.Error(), testcli.Stderr())
	}
	return nil
}

func theUpdateOutputShouldContainAnd(out1, out2 string) error {
	if !(testcli.StdoutContains(out1) && testcli.StdoutContains(out2)) {
		return fmt.Errorf("output => %s", testcli.Stdout())
	}
	return nil
}

//UpdateContext context for update command
func UpdateContext(s *godog.Suite) {
	s.Step(`^A template is installed with URL "([^"]*)"$`, aTemplateIsInstalledWithURL)
	s.Step(`^Update runs with correct template ID "([^"]*)"$`, updateRunsWithCorrectTemplateID)
	s.Step(`^The Update process state should be success$`, theUpdateProcessStateShouldBeSuccess)
	s.Step(`^The Update output should contain "([^"]*)" and "([^"]*)"$`, theUpdateOutputShouldContainAnd)
}
