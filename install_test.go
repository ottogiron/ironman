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
	testcli.Run(testutils.ExecutablePath(), "install", URL, "--ironman-home="+ironmanTestDir)
	return nil
}

func theProcessStateShouldBeSuccess() error {
	if !testcli.Success() {
		return fmt.Errorf("Install command did not succeded %s", testcli.Error())
	}
	return nil
}

func aTemplateShouldBeInstalledInPath(arg1 string) error {
	return godog.ErrPending
}

func theOutputShouldContainAnd(out1, out2 string) error {
	if !strings.Contains(testcli.Stdout(), out1) && !strings.Contains(testcli.Stdout(), out2) {
		return fmt.Errorf("output => %s", testcli.Stdout())
	}
	return godog.ErrPending
}

func itRunsWithIncorrectURL(URL string) error {
	return godog.ErrPending
}

func theProcessStateShouldBeFailure() error {
	return godog.ErrPending
}

func theOutputShouldCointain(arg1 string) error {
	return godog.ErrPending
}

func FeatureContext(s *godog.Suite) {
	s.Step(`^It runs with correct URL "([^"]*)"$`, itRunsWithCorrectURL)
	s.Step(`^The process state should be success$`, theProcessStateShouldBeSuccess)
	s.Step(`^A template should be installed in path "([^"]*)"$`, aTemplateShouldBeInstalledInPath)
	s.Step(`^The output should contain "([^"]*)" and "([^"]*)"$`, theOutputShouldContainAnd)
	s.Step(`^It runs with incorrect URL "([^"]*)"$`, itRunsWithIncorrectURL)
	s.Step(`^The process state should be failure$`, theProcessStateShouldBeFailure)
	s.Step(`^The output should cointain "([^"]*)"$`, theOutputShouldCointain)
	s.BeforeScenario(func(i interface{}) {
		_ = os.RemoveAll(ironmanTestDir)
	})
}
