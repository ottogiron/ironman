
Feature: Update a template
  In order to Update a template
  As a developer
  I need to run the Update command 

    Scenario: Update a template with an ID
    Given A template is installed with URL "https://github.com/ironman-project/template-example.git"
    When Update runs with correct template ID "template-example"
    Then The Update process state should be success 
    And The Update output should contain "Updating template template-example" and "done"
