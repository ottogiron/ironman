
Feature: Link a template
  In order to Link a template from the filesystem to the ironman repository
  As a developer
  I need to run the Link command 

  Scenario: Link a template to the ironman repository
    When It runs with correct a template path to link "testing/linkable-template"
    Then The process state should be success 
    And The output should contain "Linking template to repository with name dev-linkable-template" and "done"
    And A template link should exists be installed with ID "dev-linkable-template"
  