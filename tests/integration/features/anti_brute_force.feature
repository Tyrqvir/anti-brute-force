# file: anti_brute_force.feature
Feature: anti bruteforce service
  As grpc client
  In order to check what an user is blocked
  I want send a request to service

  Scenario: limit by login
    Given I set login "login"
    And I set password "password"
    And I set ip "192.168.1.1"
    When I send 10 requests to authorisation
    Then the requests are not blocked

  Scenario: limit by password
    Given I set login "login"
    And I set password "password"
    And I set ip "192.168.1.1"
    When I send 10 requests to authorisation
    And I set login "login2"
    When  I send 10 requests to authorisation
    And I set login "login3"
    When  I send 10 requests to authorisation
    And I set login "login4"
    When  I send 10 requests to authorisation
    Then the requests are not blocked

  Scenario: limit by ip
    Given I set login "login"
    And I set password "password"
    And I set ip "192.168.1.1"
    When I send 10 requests to authorisation
    And I set login "login2"
    And I set password "password2"
    When  I send 10 requests to authorisation
    And I set login "login3"
    And I set password "password3"
    When  I send 10 requests to authorisation
    And I set login "login4"
    And I set password "password4"
    When  I send 10 requests to authorisation
    Then the requests are not blocked

  Scenario: limit by login with reset
    Given I set login "login"
    And I set password "password"
    And I set ip "192.168.1.1"
    When I send 10 requests to authorisation
    And I reset bucket for login "login" and password "password"
    When I send 10 requests to authorisation
    Then the requests are not blocked