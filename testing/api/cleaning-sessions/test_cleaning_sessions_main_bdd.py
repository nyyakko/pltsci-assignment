from pytest_bdd import scenario, scenarios, given, when, step, then, parsers
import json
import requests

BASE = "http://localhost:8080/v1"

scenarios('features/cleaning_sessions.feature')

@step(parsers.parse("I send \"{method}\" request with body to \"{path}\":"), target_fixture="response")
def when_i_send_request_with_body_to(method, path, docstring):
    response = requests.post(f"{BASE}/cleaning-sessions", headers={"Content-Type": "application/json"}, data=docstring, timeout=5)
    return response

@step(parsers.parse("the response status code should be {code:d}"))
def then_the_response_code_should_be(response, code):
    assert response.status_code == code
    return response

@step(parsers.parse("the response content should be:"))
def and_the_response_body_should_be(response, docstring):
    pass
