
Feature: Unlink a template
  In order to Unlink a template from the ironman repository
  As a developer
  I need to run the Unlink command 

  Scenario: Unlink a template from the ironman repository
    Given Theres a  link to "acceptance/testing/templates/linkable-template" with ID "dev-linkable-template" 
    When Unlink runs with correct ID "dev-linkable-template"
    Then The Unlink process state should be success 
    And The Unlink output should contain "Unlinking template from repository with ID dev-linkable-template" and "done"
    And A template link with ID "dev-linkable-template" should not exists
  