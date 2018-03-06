package acceptance

import (
	"fmt"

	"github.com/DATA-DOG/godog"
	"github.com/DATA-DOG/godog/gherkin"
	"github.com/ironman-project/ironman/testutils"
	"github.com/rendon/testcli"
)

func aTemplateToListIsInstalledWithURL(URL string) error {
	testcli.Run(testutils.ExecutablePath(), "install", "--ironman-home="+ironmanTestDir, URL)
	if !testcli.Success() {
		return fmt.Errorf("Template with url %s should be installed", URL)
	}
	return nil
}

func listRuns() error {
	testcli.Run(testutils.ExecutablePath(), "list", "--ironman-home="+ironmanTestDir)
	return nil
}

func theListProcessStateShouldBeSuccess() error {
	if !testcli.Success() {
		return fmt.Errorf("List command did not succeded %s %s", testcli.Error(), testcli.Stderr())
	}
	return nil
}

func theListOutputShouldContainAnd(out1 string, out2 *gherkin.DocString) error {

	if !(testcli.StdoutContains(out1) && testcli.StdoutContains(out2.Content)) {
		return fmt.Errorf("output => %s", testcli.Stdout())
	}
	return nil
}

func listRunsAnNoTemplateIsAvailable() error {
	testcli.Run(testutils.ExecutablePath(), "list", "--ironman-home="+ironmanTestDir)
	return nil
}

func theListWithNoTemplateProcessStateShouldBeSuccess() error {
	if !testcli.Success() {
		return fmt.Errorf("List command did not succeded %s %s", testcli.Error(), testcli.Stderr())
	}
	return nil
}

func theListOutputWithNoTemplateShouldContainAnd(out1, out2 string) error {
	if !(testcli.StdoutContains(out1) && testcli.StdoutContains(out2)) {
		return fmt.Errorf("output => %s", testcli.Stdout())
	}
	return nil
}

//ListContext list feature context
func ListContext(s *godog.Suite) {
	s.Step(`^A template to list is installed with URL "([^"]*)"$`, aTemplateToListIsInstalledWithURL)
	s.Step(`^List runs$`, listRuns)
	s.Step(`^The List process state should be success$`, theListProcessStateShouldBeSuccess)
	s.Step(`^The List output should contain "([^"]*)" and$`, theListOutputShouldContainAnd)
	s.Step(`^List runs an no template is available$`, listRunsAnNoTemplateIsAvailable)
	s.Step(`^The List with no template process state should be success$`, theListWithNoTemplateProcessStateShouldBeSuccess)
	s.Step(`^The List output with no template should contain "([^"]*)" and "([^"]*)"$`, theListOutputWithNoTemplateShouldContainAnd)
}
