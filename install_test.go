package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/mitchellh/go-homedir"

	"github.com/DATA-DOG/godog"
	"github.com/ironman-project/ironman/testutils"
	"github.com/rendon/testcli"
)

var ironmanTestDir string

func init() {
	var err error
	ironmanTestDir, err = homedir.Dir()

	if err != nil {
		os.Exit(-1)
	}
	ironmanTestDir = filepath.Join(ironmanTestDir, ".ironman_test")
}

func itRunsWithCorrectURL(URL string) error {
	testcli.Run(testutils.ExecutablePath(), "install", "--ironman-home="+ironmanTestDir, URL)
	return nil
}

func theProcessStateShouldBeSuccess() error {
	if !testcli.Success() {
		return fmt.Errorf("Install command did not succeded %s", testcli.Error())
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
	if !strings.Contains(testcli.Stdout(), out1) && !strings.Contains(testcli.Stdout(), out2) {
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

func FeatureContext(s *godog.Suite) {
	s.Step(`^It runs with correct URL "([^"]*)"$`, itRunsWithCorrectURL)
	s.Step(`^The process state should be success$`, theProcessStateShouldBeSuccess)
	s.Step(`^A template should be installed with ID "([^"]*)"$`, aTemplateShouldBeInstalledWithID)
	s.Step(`^The output should contain "([^"]*)" and "([^"]*)"$`, theOutputShouldContainAnd)
	s.Step(`^It runs with unreachable URL "([^"]*)"$`, itRunsWithUnreachableURL)
	s.Step(`^The process state should be failure$`, theProcessStateShouldBeFailure)
	s.Step(`^The output should cointain "([^"]*)"$`, theOutputShouldCointain)
	s.BeforeScenario(func(i interface{}) {
		_ = os.RemoveAll(ironmanTestDir)
	})
}
