#script Handlers/Background_Handler
#script Handlers/Eggplant_Handler

Feature: 1. Opportunities Fully automated integration (Jan Release) – SF to CPO Integration

  Background: Get the SUT ready for Eggplant Automation
# Given the test credentials are loaded from here.xlsx"
#And a SUT is up, ready and initalized to run Eggplant script

# === QA =====================================================================================

@QA
  Scenario: QA Test
    Given  I navigate to Eggplant "QA"
    And I close website
      


# === UAT =====================================================================================

@UAT
  Scenario: UAT Test
    Given I navigate to Eggplant "UAT"
    And I close website


