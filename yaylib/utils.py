"""
MIT License

Copyright (c) 2023 qvco

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

import re
import jwt
import uuid

from datetime import datetime
from typing import Any

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


def mention(user_id: int, display_name: str) -> str:
    """

    ユーザーをメンションします

    #### Useage

        >>> import yaylib
        >>> from yaylib import mention
        >>> api = yaylib.Client()
        >>> api.login(email, password)
        >>> api.create_post(f"こんにちは、{mention(user_id=15184, display_name='アルパカ')}さん！")

    """
    if not len(display_name):
        raise ValueError("display_nameは空白にできません。")
    return f"<@>{user_id}:{display_name}<@/>"


def build_message_tags(text: str) -> tuple[str, list[dict[str, Any]]]:
    if "<@>" in text and "<@/>" in text:
        message_tags = []
        regex = re.compile(r"<@>(\d+):([^<]+)<@/>")
        offset_adjustment = 0

        for result in regex.finditer(text):
            full_match_length = len(result.group(0))
            display_name_length = len(result.group(2))

            message_tags.append(
                {
                    "type": "user",
                    "user_id": int(result.group(1)),
                    "offset": result.start() - offset_adjustment,
                    "length": display_name_length,
                }
            )

            offset_adjustment += full_match_length - display_name_length

    text: str = re.sub(r"<@>(\d+):([^<]+)<@/>", r"\2", text)
    return text, message_tags


def generate_uuid(uuid_type=True):
    generated_uuid = str(uuid.uuid4())
    if uuid_type:
        return generated_uuid
    else:
        return generated_uuid.replace("-", "")


def generate_jwt() -> str:
    timestamp = int(datetime.now().timestamp())
    return jwt.encode(
        payload={"exp": timestamp + 5, "iat": timestamp},
        key=Configs.API_VERSION_KEY.encode("utf-8"),
    )


def is_valid_image_format(format):
    is_valid = bool(re.match(r".jpg|.jpeg|.png|.gif", format))
    return is_valid


def get_hashed_filename(att, type, key, uuid):
    today = datetime.now()
    full_date = today.strftime("%Y/%m/%d")
    thumbnail = "thumb_" if att.is_thumb else ""
    file_name = f"{thumbnail}{uuid}_{int(today.timestamp())}_{key}"
    sizes = f"_size_{att.natural_width}x{att.natural_height}"
    extension = f"{att.original_file_extension}"

    hashed_filename = f"{type}/{full_date}/{file_name}{sizes}{extension}"

    return hashed_filename
