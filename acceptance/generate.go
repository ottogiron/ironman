package acceptance

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"

	"github.com/DATA-DOG/godog"
	"github.com/ironman-project/ironman/testutils"
	"github.com/rendon/testcli"
)

var generatedPath string

func aTemplateToGenerateInstalledWithURL(URL string) error {

	testcli.Run(testutils.ExecutablePath(), "install", "--ironman-home="+ironmanTestDir, URL)
	if !testcli.Success() {
		return fmt.Errorf("Failed to install test template %s", URL)
	}
	return nil
}

func generateRunsWithIDGeneratorIDAndFlags(templateID, generatorID, flags string) error {
	generatedPath = filepath.Join(ironmanTestDir, "test")
	testcli.Run(testutils.ExecutablePath(), "generate", templateID+":"+generatorID, generatedPath, "--ironman-home="+ironmanTestDir, "--set", flags)
	return nil
}

func theGenerateProcessStateShouldBeSuccess() error {
	if !testcli.Success() {
		return fmt.Errorf("Generate command did not succeeded %s %s", testcli.Error(), testcli.Stderr())
	}
	return nil
}

func theGenerateOutputShouldContainAnd(out1, out2 string) error {
	if !(testcli.StdoutContains(out1) && testcli.StdoutContains(out2)) {
		return fmt.Errorf("output => %s", testcli.Stdout())
	}
	return nil
}

func aFileUnderTheGeneratedPathShouldContain(file, contents string) error {
	filePath := filepath.Join(generatedPath, file)
	if !testutils.FileExists(filePath) {
		return fmt.Errorf("Expected file don't exists %s", filePath)
	}
	fileContent, err := ioutil.ReadFile(filePath)

	if err != nil {
		return fmt.Errorf("Failed to read file contents %s", err)
	}

	if string(fileContent) != contents {
		return fmt.Errorf("File content %s want %s", fileContent, contents)
	}
	return nil
}

func generateWithNonExistingIDRunsWithIDGeneratorID(templateID, generatorID string) error {
	testcli.Run(testutils.ExecutablePath(), "generate", templateID+":"+generatorID, generatedPath, "--ironman-home="+ironmanTestDir)
	return nil
}

func theGenerateWithNonExistingIDProcessStateShouldBeFailure() error {
	if !testcli.Failure() {
		return fmt.Errorf("Generate command did not failed %s", testcli.Stdout())
	}
	return nil
}

func theGenerateWithNonExistingIDOutputShouldCointain(expectedOutput string) error {
	if !strings.Contains(testcli.Stderr(), expectedOutput) {
		return fmt.Errorf("output => %s", testcli.Stderr())
	}
	return nil
}

//GenerateContext context for generate command
func GenerateContext(s *godog.Suite) {
	s.Step(`^A template to generate installed with URL "([^"]*)"$`, aTemplateToGenerateInstalledWithURL)
	s.Step(`^Generate runs with ID "([^"]*)" generator ID "([^"]*)" and flags "([^"]*)"$`, generateRunsWithIDGeneratorIDAndFlags)
	s.Step(`^The generate process state should be success$`, theGenerateProcessStateShouldBeSuccess)
	s.Step(`^The generate output should contain "([^"]*)" and "([^"]*)"$`, theGenerateOutputShouldContainAnd)
	s.Step(`^A file "([^"]*)" under the generated path should contain "([^"]*)"$`, aFileUnderTheGeneratedPathShouldContain)
	s.Step(`^Generate with non existing ID runs with ID "([^"]*)" generator ID "([^"]*)"$`, generateWithNonExistingIDRunsWithIDGeneratorID)
	s.Step(`^The generate with non existing ID process state should be failure$`, theGenerateWithNonExistingIDProcessStateShouldBeFailure)
	s.Step(`^The generate with non existing ID output should cointain "([^"]*)"$`, theGenerateWithNonExistingIDOutputShouldCointain)
}
