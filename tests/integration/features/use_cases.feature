# file: use_cases.feature
Feature: cases for anti bruteforce service
  As grpc client
  In order to working with black/white lists
  I want send a request to service

  Scenario: add network to black list
    Given I add network "192.168.1.1/24" to black list
    Then I see network "192.168.1.1/24" in black list

  Scenario: add network to white list
    Given I add network "192.168.1.1/24" to white list
    Then I see network "192.168.1.1/24" in white list

  Scenario: remove network from white list
    Given I remove network "192.168.1.1/24" from white list
    Then I don't see network "192.168.1.1/24" in white list

  Scenario: remove network from black list
    Given I remove network "192.168.1.1/24" from black list
    Then I don't see network "192.168.1.1/24" in black list

  Scenario: reset bucket
    Given I set login "login"
    And I set password "password"
    When I send 11 requests to authorisation
    Then the requests are blocked
    When I reset bucket for login "login" and password "password"
    And I send 10 requests to authorisation
    Then the requests are not blocked