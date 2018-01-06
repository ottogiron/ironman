
Feature: Uninstall a template
  In order to Uninstall a template
  As a developer
  I need to run the Uninstall command 

  Scenario: Uninstalling a template with correct ID
    Given A template to uninstall URL "https://github.com/ottogiron/wizard-hello-world.git"
    When It runs with correct ID "wizard-hello-world"
    Then The Uninstall process state should be success 
    And The Uninstall output should contain "Uninstallinging template" and "done"
    And An Uninstalled template with ID "wizard-hello-world" should not exists
  

  