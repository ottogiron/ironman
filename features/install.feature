
Feature: Install a template
  In order to install a template
  As a developer
  I need to run the install command 

  Scenario: Install a template with correct URL
    Given I run install command
    When It runs with correct URL
    Then A template should be installed
    Then The exit output should be 0

  Scenario: Install a template with incorrect URL 
    Given I run install command with incorrect URL
    When It runs with incorrect incorrect URL
    Then An error message should be shown
    Then The error output should be not 0