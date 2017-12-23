package main

import (
	"fmt"
	"os/exec"

	"github.com/DATA-DOG/godog"
	"github.com/ironman-project/ironman/testutils"
	"github.com/rendon/testcli"
)

var installCommand *exec.Cmd

func itRunsWithCorrectURL() error {
	testcli.Run(testutils.ExecutablePath(), "install")
	if !testcli.Success() {
		return fmt.Errorf("Failed to run install command %s", testcli.Error())
	}
	return nil
}

func aTemplateShouldBeInstalled() error {
	return godog.ErrPending
}

func aSucessMessageShouldBeShown() error {
	return godog.ErrPending
}

func theExitOutputShouldBe(arg1 int) error {
	return godog.ErrPending
}

func itRunsWithIncorrectURL() error {
	return godog.ErrPending
}

func anErrorMessageShouldBeShown() error {
	return godog.ErrPending
}

func theErrorOutputShouldBeNot(arg1 int) error {
	return godog.ErrPending
}

func FeatureContext(s *godog.Suite) {

	s.Step(`^It runs with correct URL$`, itRunsWithCorrectURL)
	s.Step(`^A template should be installed$`, aTemplateShouldBeInstalled)
	s.Step(`^A sucess message should be shown$`, aSucessMessageShouldBeShown)
	s.Step(`^The exit output should be (\d+)$`, theExitOutputShouldBe)
	s.Step(`^It runs with incorrect URL$`, itRunsWithIncorrectURL)
	s.Step(`^An error message should be shown$`, anErrorMessageShouldBeShown)
	s.Step(`^The error output should be not (\d+)$`, theErrorOutputShouldBeNot)
}
