
Feature: Unlink a template
  In order to Unlink a template from the ironman repository
  As a developer
  I need to run the Unlink command 

  Scenario: Unlink a template from the ironman repository
    Given There is a template installed with URL "https://github.com/ottogiron/wizard-hello-world.git" 
    When Unlink runs with correct ID "wizard-hello-world"
    Then The Unlink process state should be success 
    And The Unlink output should contain "Unlinking template from repository with ID wizard-hello-world" and "done"
    And A template link with ID "wizard-hello-world" should not exists
  