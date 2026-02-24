Feature: Hoover
    Scenario: Walk normally
        When I send "POST" request with body to "/v1/cleaning-sessions":
            """
            {"roomSize":[5,5],"coords":[1,2],"patches":[[1,0],[2,2],[2,3]],"instructions":"NNESEESWNWW"}
            """
        Then the response code should be 200
        And the response body should be:
            """
            {"coords":[1,3],"patches":1}
            """

    Scenario: Do nothing
        When I send "POST" request with body to "/v1/cleaning-sessions":
            """
            {"roomSize":[5,5],"coords":[1,2],"patches":[[1,0],[2,2],[2,3]],"instructions":""}
            """
        Then the response code should be 200
        And the response body should be:
            """
            {"coords":[1,2],"patches":0}
            """

    Scenario: Walk out of bounds
        When I send "POST" request with body to "/v1/cleaning-sessions":
            """
            {"roomSize":[5,5],"coords":[1,2],"patches":[[1,0],[2,2],[2,3]],"instructions":"NNNESEESWNWW"}
            """
        Then the response code should be 500
        And the response body should be:
            """
            {"error":"HOOVER_WENT_OUT_OF_BOUNDS"}
            """

    Scenario: Send unknown instruction
        When I send "POST" request with body to "/v1/cleaning-sessions":
            """
            {"roomSize":[5,5],"coords":[1,2],"patches":[[1,0],[2,2],[2,3]],"instructions":"NNESEGSWNWW"}
            """
        Then the response code should be 500
        And the response body should be:
            """
            {"error":"UNKNOWN_INSTRUCTION_RECEIVED"}
            """
