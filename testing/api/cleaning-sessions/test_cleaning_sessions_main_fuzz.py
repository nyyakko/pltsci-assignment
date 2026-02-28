from dotenv import load_dotenv
from hypothesis import given, settings, strategies as st, Verbosity
from hypothesis.strategies import text, integers, lists
import json
import os
import requests
import time

load_dotenv()

json_body = st.fixed_dictionaries({
    "roomSize": st.lists(st.integers(min_value=1, max_value=1000), min_size=2, max_size=2),
    "coords": st.lists(st.integers(min_value=1, max_value=1000), min_size=2, max_size=2),
    "patches": st.lists(st.lists(st.integers(min_value=1, max_value=1000), min_size=2, max_size=2), min_size=0, max_size=50),
    "instructions": st.text(alphabet=list("NESW") + [chr(x) for x in range(32, 127)], min_size=0)
})

@settings(max_examples=500, deadline=None)
@given(body=json_body)
def test_cleaning_sessions(body):
    response = requests.post(f"{os.getenv("API_BASE_URL")}/cleaning-sessions", headers={"Content-Type": "application/json"}, data=json.dumps(body), timeout=5)
    assert not (500 <= response.status_code <= 599)
    # time.sleep(0.5)
