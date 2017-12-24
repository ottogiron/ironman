
Feature: Install a template
  In order to install a template
  As a developer
  I need to run the install command 

  Scenario: Install a template with correct URL
    When It runs with correct URL "https://github.com/ottogiron/wizard-hello-world.git"
    Then The process state should be success 
    And The output should contain "Installing template" and "done"
    And A template should be installed with ID "wizard-hello-world"
   
  Scenario: Install a template with unreachable URL 
    When It runs with unreachable URL "http://hola"
    Then The process state should be failure
    And The output should cointain "Failed to install template"