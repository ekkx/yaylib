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

import random
from typing import Optional

from . import config

DEVICES = [
    {
        "device_type": "android",
        "os_version": "11",
        "screen_density": "3.5",
        "screen_size": "1440x2960",
        "model": "Galaxy S9",
    },
]


class Device:
    """端末情報を管理するクラス"""

    def __init__(
        self,
        device_type: str,
        os_version: str,
        screen_density: str,
        screen_size: str,
        model: str,
    ):
        self.device_type = device_type
        self.os_version = os_version
        self.screen_density = screen_density
        self.screen_size = screen_size
        self.model = model

    @classmethod
    def create(cls, device: Optional[dict] = None, model: Optional[str] = None):
        """端末を生成する"""
        if device is None:
            if model is not None:
                device = next(
                    (device for device in DEVICES if device["model"] == model), None
                )
                if device is None:
                    raise ValueError("No device found with model: " + model)
            else:
                device = random.choice(DEVICES)
        return cls(**device)

    def get_user_agent(self) -> str:
        """ユーザーエージェントを取得する"""
        # pylint: disable=line-too-long
        return f"{self.device_type} {self.os_version} ({self.screen_density}x {self.screen_size} {self.model})"

    def get_device_info(self) -> str:
        """端末情報を取得する"""
        return f"yay {config.VERSION_NAME} {self.get_user_agent()}"
