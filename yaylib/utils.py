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
import uuid
from datetime import datetime


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


def generate_uuid(uuid_type=True):
    generated_uuid = str(uuid.uuid4())
    if uuid_type:
        return generated_uuid
    else:
        return generated_uuid.replace("-", "")


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
