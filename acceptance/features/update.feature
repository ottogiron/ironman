
Feature: Update a template
  In order to Update a template
  As a developer
  I need to run the Update command 

    Scenario: Update a template with an ID
    Given A template is installed with URL "https://github.com/ottogiron/wizard-hello-world.git"
    When Update runs with correct template ID "wizard-hello-world"
    Then The Update process state should be success 
    And The Update output should contain "Updating template wizard-hello-world" and "done"
