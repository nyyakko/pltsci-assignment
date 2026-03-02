Feature: Cleaning Sessions
    Scenario: Do nothing
        When I send "POST" request with body to "/v1/cleaning-sessions":
            """
            {"roomSize":[5,5],"coords":[1,2],"patches":[[1,0],[2,2],[2,3]],"instructions":""}
            """

        Then the response status code should be 200

        And the response content should be:
            """
            {"coords":[1,2],"patches":0}
            """

    Scenario: Clean normally
        When I send "POST" request with body to "/v1/cleaning-sessions":
            """
            {"roomSize":[5,5],"coords":[1,2],"patches":[[1,0],[2,2],[2,3]],"instructions":"NNESEESWNWW"}
            """

        Then the response status code should be 200

        And the response content should be:
            """
            {"coords":[1,3],"patches":1}
            """

    Scenario: Clean patch underneath
        When I send "POST" request with body to "/v1/cleaning-sessions":
            """
            {"roomSize":[5,5],"coords":[2,2],"patches":[[2,2]],"instructions":""}
            """

        Then the response status code should be 200

        And the response content should be:
            """
            {"coords":[2,2],"patches":1}
            """
