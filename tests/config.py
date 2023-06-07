import logging
import json
import unittest

import vcr

import yaylib


json_path = "../config.json"
cassette_library_dir = "cassettes"

with open(json_path, "r") as f:
    config = json.load(f)

email = config.get("EMAIL", "")
password = config.get("PASSWORD", "")
user_id = config.get("USER_ID", "")
opponent_id = config.get("OPPONENT_ID", "")
access_token = config.get("ACCESS_TOKEN", "")


def before_record_response(response):
    body = response["content"]
    if body:
        json_body = json.loads(body)
        if "token" in json_body:
            del json_body["token"]
        response["content"] = json.dumps(json_body)
    return response


tape = vcr.VCR(
    cassette_library_dir=cassette_library_dir,
    filter_headers=["Authorization", "X-Jwt"],
    filter_query_parameters=["token", "access_token"],
    before_record_response=before_record_response,
)


class YaylibTestCase(unittest.TestCase):

    def setUp(self):

        if access_token:
            self.api = yaylib.Client(
                access_token=access_token,
                # loglevel_stream=logging.DEBUG,
            )
        else:
            self.api = yaylib.Client(
                # loglevel_stream=logging.DEBUG,
            )
            self.api.login_with_email(email, password)

            config["ACCESS_TOKEN"] = self.api.login_data.access_token
            with open(json_path, "w") as f:
                f.write(json.dumps(config))
