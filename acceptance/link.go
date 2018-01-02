package acceptance

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/DATA-DOG/godog"
	"github.com/ironman-project/ironman/testutils"
	"github.com/rendon/testcli"
)

func itRunsWithCorrectATemplatePathToLinkWithID(templatePath, ID string) error {

	testcli.Run(testutils.ExecutablePath(), "link", "--ironman-home="+ironmanTestDir, templatePath, ID)
	return nil
}

func theLinkProcessStateShouldBeSuccess() error {
	if !testcli.Success() {
		return fmt.Errorf("Link command did not succeded %s %s %s", testcli.Error(), testcli.Stderr(), testcli.Stdout())
	}
	return nil
}

func theLinkOutputShouldContainAnd(out1, out2 string) error {
	if !testcli.StdoutContains(out1) && !testcli.StdoutContains(out2) {
		return fmt.Errorf("output => %s", testcli.Stdout())
	}
	return nil
}

func aTemplateLinkShouldExistsBeInstalledWithID(templateLinkID string) error {
	linkPath := filepath.Join(ironmanTemplatesDir, templateLinkID)
	if _, err := os.Stat(linkPath); os.IsNotExist(err) {
		return fmt.Errorf("Link for %s not found", templateLinkID)
	}
	return nil
}

//LinkContext context for install command
func LinkContext(s *godog.Suite) {
	s.Step(`^It runs with correct a template path to link "([^"]*)" with ID "([^"]*)"$`, itRunsWithCorrectATemplatePathToLinkWithID)
	s.Step(`^The Link process state should be success$`, theLinkProcessStateShouldBeSuccess)
	s.Step(`^The Link output should contain "([^"]*)" and "([^"]*)"$`, theLinkOutputShouldContainAnd)
	s.Step(`^A template link should exists be installed with ID "([^"]*)"$`, aTemplateLinkShouldExistsBeInstalledWithID)
}
