
Feature: Install a template
  In order to install a template
  As a developer
  I need to run the install command 

  Scenario: Install a template with correct URL
    When It runs with correct URL
    Then A template should be installed
    And A sucess message should be shown
    And The exit output should be 0

  Scenario: Install a template with incorrect URL 
    When It runs with incorrect URL
    Then An error message should be shown
    And The error output should be not 0