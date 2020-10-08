# This feature is designed to execute against a variety of devices in your local environment or the TestPlant Automation Cloud.
# Use the connection instructions provided with your Trial license request to connect one of those devices

# This Feature calls a series of Eggplant Functional Gherkin handlers to open the Amazon search page, click on some page elements, verify that the elements were found, and then close the browser

# To inspect the Gherkin script handlers, open the "Gherkin_Mobile" script in the Scripts section on the left side of this window. (If scripts are hidden, click the pencil icon to show them.) 


Feature: AmazonApp
#   This feature is specific to a mobile device

  Scenario: Amazon_OpenApp
# -----------------Change the device string to match the mobile device with which you are currently connected --------------------
    Given device is connected
    When Environment is setup
    Then Launch Amazon App
      
  Scenario: SearchItem
    Given I am on home screen
    When I search for "rene fleming mozart"
    Then I should see filter
      
  Scenario: HomeScreen
    Given I should see filter
    When I got to the home screen
    Then I should see home screen icon
      
  Scenario: CloseApp
    Given I should see home screen icon
    When I close the app
    Then The test is over









