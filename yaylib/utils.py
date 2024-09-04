"""
MIT License

Copyright (c) 2023 ekkx

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
import logging
import re
import uuid
from base64 import urlsafe_b64encode
from datetime import datetime
from json import dumps
from typing import Any, Optional

from . import config
from .constants import Color
from .models import Attachment


class CustomFormatter(logging.Formatter):
    """ログ用意のフォーマッター"""

    @staticmethod
    def __get_formats() -> dict:
        date = Color.HEADER + "%(asctime)s " + Color.RESET
        level = Color.UNDERLINE + "%(levelname)s" + Color.RESET
        body = " » %(message)s"

        return {
            logging.DEBUG: date + Color.OKGREEN + level + Color.RESET + body,
            logging.INFO: date + Color.OKBLUE + level + Color.RESET + body,
            logging.WARNING: date + Color.WARNING + level + Color.RESET + body,
            logging.ERROR: date + Color.FAIL + level + Color.RESET + body,
            logging.CRITICAL: date + Color.FAIL + level + Color.RESET + body,
        }

    def format(self, record):
        fmt = self.__get_formats().get(record.levelno)
        formatter = logging.Formatter(fmt)
        return formatter.format(record)


def mention(user_id: int, display_name: str) -> str:
    """

    ユーザーをメンションします

    ※ メンションするには相手をフォローする必要があります。

    #### Useage

        >>> import yaylib
        >>> from yaylib import mention
        >>> api = yaylib.Client()
        >>> api.login(email, password)
        >>> api.create_post(f"こんにちは、{mention(user_id=15184, display_name='アルパカ')}さん！")

    """
    if not len(display_name):
        raise ValueError("display_nameは空白にできません。")
    return f"<@>{user_id}:@{display_name}<@/>"


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

        text = re.sub(r"<@>(\d+):([^<]+)<@/>", r"\2", text)
        return text, message_tags


def get_post_type(**kwargs) -> str:
    if kwargs.get("choices"):
        return "survey"
    elif kwargs.get("shared_url"):
        return "shareable_url"
    elif kwargs.get("video_file_name"):
        return "video"
    elif kwargs.get("attachment_filename"):
        return "image"
    else:
        return "text"


def filter_dict(params: Optional[dict] = None) -> Optional[dict]:
    if params is None:
        return None
    new_params = {}
    for k in params:
        if params[k] is not None:
            new_params[k] = params[k]
    return new_params


def generate_uuid(uuid_type=True):
    generated_uuid = str(uuid.uuid4())
    if uuid_type:
        return generated_uuid
    else:
        return generated_uuid.replace("-", "")


def generate_jwt() -> str:
    timestamp = int(datetime.now().timestamp())
    encoded_headers = (
        urlsafe_b64encode(dumps({"alg": "HS256"}, separators=(",", ":")).encode())
        .decode()
        .strip("=")
    )
    encoded_payload = (
        urlsafe_b64encode(
            dumps(
                {"iat": timestamp, "exp": timestamp + 5}, separators=(",", ":")
            ).encode()
        )
        .decode()
        .strip("=")
    )
    payload = encoded_headers + "." + encoded_payload
    sig = (
        urlsafe_b64encode(
            hmac.new(
                key=config.API_VERSION_KEY.encode(),
                msg=payload.encode(),
                digestmod=hashlib.sha256,
            ).digest()
        )
        .decode()
        .strip("=")
    )
    return payload + "." + sig


def is_valid_image_format(image_format: str):
    allowed_formats = [".jpg", ".jpeg", ".png", ".gif"]
    return image_format in allowed_formats


def is_valid_video_format(video_format: str):
    allowed_formats = [".mp4"]
    return video_format in allowed_formats


def get_hashed_filename(att: Attachment, file_type: str, key: int, uuid_str: str):
    today = datetime.now()
    full_date = today.strftime("%Y/%m/%d")
    thumbnail = "thumb_" if att.is_thumb else ""
    file_name = f"{thumbnail}{uuid_str}_{int(today.timestamp())}_{key}"
    sizes = f"_size_{att.natural_width}x{att.natural_height}"
    extension = f"{att.original_file_extension}"

    hashed_filename = f"{file_type}/{full_date}/{file_name}{sizes}{extension}"

    return hashed_filename


def md5(device_uuid: str, timestamp: int, require_shared_key: bool) -> str:
    shared_key: str = config.SHARED_KEY if require_shared_key else ""
    return hashlib.md5(
        (config.API_KEY + device_uuid + str(timestamp) + shared_key).encode()
    ).hexdigest()


def sha256() -> str:
    return base64.b64encode(
        hmac.new(
            config.API_VERSION_KEY.encode(),
            "yay_android/{}".format(config.API_VERSION_NAME).encode(),
            hashlib.sha256,
        ).digest()
    ).decode("utf-8")
