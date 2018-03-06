package acceptance

import (
	"fmt"
	"path/filepath"

	"github.com/DATA-DOG/godog"
	"github.com/ironman-project/ironman/testutils"
	"github.com/rendon/testcli"
)

func aTemplateToUninstallURL(URL string) error {
	testcli.Run(testutils.ExecutablePath(), "install", "--ironman-home="+ironmanTestDir, URL)
	if !testcli.Success() {
		return fmt.Errorf("Template with url %s should be installed", URL)
	}
	return nil
}

func itRunsWithCorrectID(ID string) error {
	testcli.Run(testutils.ExecutablePath(), "uninstall", "--ironman-home="+ironmanTestDir, ID)
	return nil
}

func theUninstallProcessStateShouldBeSuccess() error {
	if !testcli.Success() {
		return fmt.Errorf("Uninstall command did not succeded %s %s", testcli.Error(), testcli.Stderr())
	}
	return nil
}

func theUninstallOutputShouldContainAnd(out1, out2 string) error {
	if !(testcli.StdoutContains(out1) && testcli.StdoutContains(out2)) {
		return fmt.Errorf("output => %s", testcli.Stdout())
	}
	return nil
}

func anUninstalledTemplateWithIDShouldNotExists(templateID string) error {
	templatePath := filepath.Join(ironmanTemplatesDir, templateID)
	if testutils.FileExists(templatePath) {
		return fmt.Errorf("Directory should not exists for template ID %s", templateID)
	}
	return nil
}

//UninstallContext context for uninstall feature
func UninstallContext(s *godog.Suite) {
	s.Step(`^A template to uninstall URL "([^"]*)"$`, aTemplateToUninstallURL)
	s.Step(`^It runs with correct ID "([^"]*)"$`, itRunsWithCorrectID)
	s.Step(`^The Uninstall process state should be success$`, theUninstallProcessStateShouldBeSuccess)
	s.Step(`^The Uninstall output should contain "([^"]*)" and "([^"]*)"$`, theUninstallOutputShouldContainAnd)
	s.Step(`^An Uninstalled template with ID "([^"]*)" should not exists$`, anUninstalledTemplateWithIDShouldNotExists)
}
