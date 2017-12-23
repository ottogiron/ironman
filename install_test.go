package main

import "github.com/DATA-DOG/godog"

func iRunInstallCommand() error {
	return godog.ErrPending
}

func itRunsWithCorrectURL() error {
	return godog.ErrPending
}

func aTemplateShouldBeInstalled() error {
	return godog.ErrPending
}

func theExitOutputShouldBe(arg1 int) error {
	return godog.ErrPending
}

func iRunInstallCommandWithIncorrectURL() error {
	return godog.ErrPending
}

func itRunsWithIncorrectIncorrectURL() error {
	return godog.ErrPending
}

func anErrorMessageShouldBeShown() error {
	return godog.ErrPending
}

func theErrorOutputShouldBeNot(arg1 int) error {
	return godog.ErrPending
}

func FeatureContext(s *godog.Suite) {
	s.Step(`^I run install command$`, iRunInstallCommand)
	s.Step(`^It runs with correct URL$`, itRunsWithCorrectURL)
	s.Step(`^A template should be installed$`, aTemplateShouldBeInstalled)
	s.Step(`^The exit output should be (\d+)$`, theExitOutputShouldBe)
	s.Step(`^I run install command with incorrect URL$`, iRunInstallCommandWithIncorrectURL)
	s.Step(`^It runs with incorrect incorrect URL$`, itRunsWithIncorrectIncorrectURL)
	s.Step(`^An error message should be shown$`, anErrorMessageShouldBeShown)
	s.Step(`^The error output should be not (\d+)$`, theErrorOutputShouldBeNot)
}
