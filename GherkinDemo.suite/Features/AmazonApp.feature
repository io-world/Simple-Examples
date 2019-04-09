Feature: AmazonApp

  Scenario: Open the amazon application
    Given Iamonthehomescreen
    When I click the amazon I app icon
    Then I should be on the amazon home screen
      
  Scenario: SearchItem
    Given I should be on the amazon home screen
    when I enter the search "playstation"
    Then I should see the filter button
      
  Scenario: SearchPage
    Given I should see the filter button
    When I click the back button
    Then I should be on the amazon home screen
      
  Scenario: AmazonHomeScreen
    Given I should be on the amazon home screen
    when I click the home button
    then I am on the home screen   
      
      
      
      
Scenario: AmazonHomeScreen
    Given I should be on the amazon home screen
    when I click the home button
    then I am on the home screen   
      
      





