package main

import "github.com/DATA-DOG/godog"

func iRunInstallCommand() error {
	return godog.ErrPending
}

func itRunsWithCorrectParameters() error {
	return godog.ErrPending
}

func aTemplateShouldBeInstalled() error {
	return godog.ErrPending
}

func iRunInstallCommandWithIncorrectParameters() error {
	return godog.ErrPending
}

func itRunsWithIncorrectParameters() error {
	return godog.ErrPending
}

func anErrorMessageShouldBeShown() error {
	return godog.ErrPending
}

func FeatureContext(s *godog.Suite) {
	s.Step(`^I run install command$`, iRunInstallCommand)
	s.Step(`^It runs with correct parameters$`, itRunsWithCorrectParameters)
	s.Step(`^A template should be installed$`, aTemplateShouldBeInstalled)
	s.Step(`^I run install command with incorrect parameters$`, iRunInstallCommandWithIncorrectParameters)
	s.Step(`^It runs with incorrect parameters$`, itRunsWithIncorrectParameters)
	s.Step(`^An error message should be shown$`, anErrorMessageShouldBeShown)
}
