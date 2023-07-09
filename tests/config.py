"""
MIT License

Copyright (c) 2023-present Qvco, Konn

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
"""

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
                loglevel_stream=logging.DEBUG,
            )
        else:
            self.api = yaylib.Client(
                loglevel_stream=logging.DEBUG,
            )
            self.api.login(email, password)

            config["ACCESS_TOKEN"] = self.api.login_data.access_token
            with open(json_path, "w") as f:
                f.write(json.dumps(config))
