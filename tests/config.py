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

import os
import vcr
import json
import logging
import unittest

import yaylib


base_path = os.path.abspath(os.getcwd()) + "/config/"
cassette_library_dir = "cassettes"


with open(base_path + "test_config.json", "r") as f:
    test_config = json.load(f)


email = test_config.get("email", "")
password = test_config.get("password", "")
secret_key = test_config.get("secret_key", "")
opponent_id = test_config.get("opponent_id", "")


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
        self.api = yaylib.Client(loglevel=logging.DEBUG)
        self.api.login(email, password, secret_key)
