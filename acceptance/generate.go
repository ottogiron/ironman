package acceptance

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/DATA-DOG/godog"
	"github.com/DATA-DOG/godog/gherkin"
	"github.com/ironman-project/ironman/pkg/testutils"
	"github.com/rendon/testcli"
)

var generatedPath string
var generateFilePath string

func aTemplateToGenerateInstalledWithURL(URL string) error {

	testcli.Run(testutils.ExecutablePath(), "install", "--ironman-home="+ironmanTestDir, URL)
	if !testcli.Success() {
		return fmt.Errorf("failed to install test template %s", URL)
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
		return fmt.Errorf("Generate command did not succeeded %s\n %s %s", testcli.Stdout(), testcli.Error(), testcli.Stderr())
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
		return fmt.Errorf("failed to read file contents %s", err)
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

func aTemplateToGenerateAFileInstalledWithURL(URL string) error {
	testcli.Run(testutils.ExecutablePath(), "install", "--ironman-home="+ironmanTestDir, URL)
	if !testcli.Success() {
		return fmt.Errorf("failed to install test template %s", URL)
	}
	return nil
}

func generateFileRunsWithIDGeneratorIDAndFileNameAndFlags(templateID, generatorID, fileName, flags string) error {
	generateFilePath = filepath.Join(ironmanTestDir, "testfile", fileName)
	_ = os.Mkdir(filepath.Dir(generateFilePath), os.ModePerm)
	testcli.Run(testutils.ExecutablePath(), "generate", templateID+":"+generatorID, generateFilePath, "--ironman-home="+ironmanTestDir, "--set", flags)
	return nil
}

func theGenerateFileProcessStateShouldBeSuccess() error {
	if !testcli.Success() {
		return fmt.Errorf("Generate command did not succeeded %s\n %s %s", testcli.Stdout(), testcli.Error(), testcli.Stderr())
	}
	return nil
}

func theGenerateFileOutputShouldContainAnd(out1, out2 string) error {
	if !(testcli.StdoutContains(out1) && testcli.StdoutContains(out2)) {
		return fmt.Errorf("output => %s", testcli.Stdout())
	}
	return nil
}

func aFileFromAFileGeneratorUnderTheGeneratedPathShouldContain(contents *gherkin.DocString) error {

	if !testutils.FileExists(generateFilePath) {
		return fmt.Errorf("Expected file don't exists %s", generateFilePath)
	}
	fileContent, err := ioutil.ReadFile(generateFilePath)

	if err != nil {
		return fmt.Errorf("failed to read file contents %s", err)
	}

	if string(fileContent) != contents.Content {
		return fmt.Errorf("file content \n%s\n want \n%s", fileContent, contents.Content)
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

	s.Step(`^A template to generate a file installed with URL "([^"]*)"$`, aTemplateToGenerateAFileInstalledWithURL)
	s.Step(`^Generate file runs with ID "([^"]*)" generator ID "([^"]*)"  and fileName "([^"]*)"  and flags "([^"]*)"$`, generateFileRunsWithIDGeneratorIDAndFileNameAndFlags)
	s.Step(`^The generate file process state should be success$`, theGenerateFileProcessStateShouldBeSuccess)
	s.Step(`^The generate file output should contain "([^"]*)" and "([^"]*)"$`, theGenerateFileOutputShouldContainAnd)
	s.Step(`^A file under the generated path should contain$`, aFileFromAFileGeneratorUnderTheGeneratedPathShouldContain)
}
