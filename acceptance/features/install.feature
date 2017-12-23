
Feature: Install a template
  In order to do something
  As a developer
  I need to do something

  Scenario: Correctly install a template
    Given I run install command
    When It runs with correct parameters
    Then A template should be installed

  Scenario: Incorectly install a template 2
    Given I run install command with incorrect parameters
    When It runs with incorrect parameters
    Then An error message shoudl be shown