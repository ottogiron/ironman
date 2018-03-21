
Feature: List Available Templates
  In order to list the available a templates
  As a developer
  I need to run the list command

  Scenario: List available templates
    Given A template to list is installed with URL "https://github.com/ironman-project/template-example.git"
    When List runs 
    Then The List process state should be success 
    And The List output should contain "Installed templates" and 
    """
+------------------+------------------+--------------------------------+
|        ID        |       NAME       |          DESCRIPTION           |
+------------------+------------------+--------------------------------+
| template-example | Template Example | This is an example of a valid  |
|                  |                  | template.                      |
+------------------+------------------+--------------------------------+
    """

    Scenario: List available templates when none available 
    When List runs an no template is available
    Then The List with no template process state should be success 
    And The List output with no template should contain "Installed templates" and "none"