package acceptance

import (
	"fmt"
	"path/filepath"

	"github.com/DATA-DOG/godog"
	"github.com/ironman-project/ironman/testutils"
	"github.com/rendon/testcli"
)

func theresALinkToWithID(templatePath, ID string) error {
	testcli.Run(testutils.ExecutablePath(), "link", "--ironman-home="+ironmanTestDir, templatePath, ID)
	if !testcli.Success() {
		return fmt.Errorf("Failed to link test template %s %s", templatePath, testcli.Stderr())
	}
	return nil
}

func unlinkRunsWithCorrectID(templateID string) error {
	testcli.Run(testutils.ExecutablePath(), "unlink", "--ironman-home="+ironmanTestDir, templateID)
	return nil
}

func theUnlinkProcessStateShouldBeSuccess() error {
	if !testcli.Success() {
		return fmt.Errorf("Failed to unlink test template %s", testcli.Stderr())
	}
	return nil
}

func theUnlinkOutputShouldContainAnd(out1, out2 string) error {
	if !(testcli.StdoutContains(out1) && testcli.StdoutContains(out2)) {
		return fmt.Errorf("output => %s", testcli.Stdout())
	}
	return nil
}

func aTemplateLinkWithIDShouldNotExists(templateLinkID string) error {
	linkPath := filepath.Join(ironmanTemplatesDir, templateLinkID)
	if testutils.FileExists(linkPath) {
		return fmt.Errorf("Link should not exists for template ID %s", templateLinkID)
	}
	return nil
}

//UnlinkContext context for unlink command
func UnlinkContext(s *godog.Suite) {
	s.Step(`^Theres a  link to "([^"]*)" with ID "([^"]*)"$`, theresALinkToWithID)
	s.Step(`^Unlink runs with correct ID "([^"]*)"$`, unlinkRunsWithCorrectID)
	s.Step(`^The Unlink process state should be success$`, theUnlinkProcessStateShouldBeSuccess)
	s.Step(`^The Unlink output should contain "([^"]*)" and "([^"]*)"$`, theUnlinkOutputShouldContainAnd)
	s.Step(`^A template link with ID "([^"]*)" should not exists$`, aTemplateLinkWithIDShouldNotExists)
}
