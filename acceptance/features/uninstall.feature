
Feature: Uninstall a template
  In order to Uninstall a template
  As a developer
  I need to run the Uninstall command 

  Scenario: Uninstalling a template with correct ID
    Given A template to uninstall URL "https://github.com/ironman-project/template-example.git"
    When It runs with correct ID "template-example"
    Then The Uninstall process state should be success 
    And The Uninstall output should contain "Uninstalling template" and "done"
    And An Uninstalled template with ID "template-example" should not exists
  

  