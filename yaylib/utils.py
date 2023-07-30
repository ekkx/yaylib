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

import base64
import hashlib
import hmac
import json
import os
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


def generate_uuid(uuid_type: bool = True) -> tuple:
    generated_uuid = str(uuid.uuid4())
    if uuid_type:
        return generated_uuid
    else:
        return generated_uuid.replace("-", "")


def parse_datetime(timestamp: int) -> str:
    if timestamp is not None:
        return str(datetime.fromtimestamp(timestamp))
    return timestamp


def encrypt(fernet, cookies: dict):
    cookies.update(
        {
            "access_token": fernet.encrypt(
                cookies.get("access_token").encode()
            ).decode(),
            "refresh_token": fernet.encrypt(
                cookies.get("refresh_token").encode()
            ).decode(),
        }
    )
    return cookies


def decrypt(fernet, cookies: dict):
    cookies.update(
        {
            "access_token": fernet.decrypt(cookies.get("access_token")).decode(),
            "refresh_token": fernet.decrypt(cookies.get("refresh_token")).decode(),
        }
    )
    return cookies


def save_cookies(
    base_path: str,
    cookie_filename: str,
    fernet,
    access_token: str,
    refresh_token: str,
    user_id: int,
    email: str = None,
):
    cookies = load_cookies(base_path=base_path, cookie_filename=cookie_filename)
    updated_cookies = {
        "access_token": access_token,
        "refresh_token": refresh_token,
        "user_id": user_id,
        "email": email,
    }
    if email is None:
        updated_cookies["email"] = cookies.get("email")

    updated_cookies = encrypt(fernet, updated_cookies)

    with open(base_path + cookie_filename + ".json", "w") as f:
        json.dump(updated_cookies, f, indent=4)


def load_cookies(
    base_path: str, cookie_filename: str, fernet=None, check_email: str = None
):
    if not os.path.exists(base_path + cookie_filename + ".json"):
        return None

    with open(base_path + cookie_filename + ".json", "r") as f:
        cookies = json.load(f)

    result = all(
        key in cookies for key in ("access_token", "refresh_token", "user_id", "email")
    )
    cookies = None if result is False else cookies

    if check_email is not None and cookies is not None:
        cookies = None if check_email != cookies["email"] else cookies

    if fernet is not None and cookies is not None:
        cookies = decrypt(fernet, cookies)

    return cookies
