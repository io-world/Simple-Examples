# This script is designed to execute against TestPlant publicly available Device Automation cloud.
# Use the connection instructions provided with your Trial license request to connect to one of those devices

# This Feature calls a series of Eggplant Functional Gherkin handlers to open the Amazon search page, click on some page elements, verify that the elements were found, and then close the browser. The browser specified in this feature file is currently Chrome, but you can modify this feature and corresponding handlers to use any browser present on your SUT. 

# To inspect the Gherkin script handlers, open the "Gherkin_Desktop" script in the Scripts section on the left side of this window. (If scripts are hidden, click the pencil icon to show them.) 

Feature: AmazonBasket 
#  This feature is specific to a desktop device  

  Scenario: OpenAmazonSite
# -----------------Change the device string to match the desktop device with which you are currently connected --------------------    
    Given I am connected to a device
    When I enter the amazon URL on "chrome"
    Then I should see the amazon home page

  Scenario: SelectBasket
    Given I should see the amazon home page
    When I click the basket icon
    Then I should see the shopping page
      
  Scenario: CloseBrowser
    Given I close the browser
    Then The test should pass
