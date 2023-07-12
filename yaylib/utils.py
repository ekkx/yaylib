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

import hmac
import hashlib
import base64
import uuid

from datetime import datetime
from .config import Configs


class Colors:
    HEADER = "\033[95m"
    OKBLUE = "\033[94m"
    OKCYAN = "\033[96m"
    OKGREEN = "\033[92m"
    WARNING = "\033[93m"
    FAIL = "\033[91m"
    RESET = "\033[0m"
    BOLD = "\033[1m"
    UNDERLINE = "\033[4m"


def console_print(*args):
    print("\n")
    for arg in args:
        print(arg)
    print("\n")


def generate_uuid() -> tuple:
    generated_uuid = str(uuid.uuid4())
    url_uuid = generated_uuid.replace("-", "")
    return generated_uuid, url_uuid


def parse_datetime(timestamp: int) -> str:
    if timestamp is not None:
        return str(datetime.fromtimestamp(timestamp))
    return timestamp


def signed_info_calculating(uuid: str, timestamp: int, shared_key: bool = False) -> str:
    """
    Pass the device_uuid when shared_key is False.
    """
    shared_key = Configs.YAY_SHARED_KEY if shared_key is True else ""
    return hashlib.md5(
        (Configs.YAY_API_KEY + uuid + str(timestamp) + shared_key).encode()
    ).hexdigest()


def signed_version_calculating() -> str:
    hash_object = hmac.new(
        Configs.YAY_API_VERSION_KEY.encode(),
        "yay_android/{}".format(Configs.YAY_API_VERSION).encode(),
        hashlib.sha256,
    )
    return base64.b64encode(hash_object.digest()).decode("utf-8")
