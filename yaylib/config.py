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

from .models import Device


class Configs:
    API_VERSION_NAME = "3.40"
    VERSION_NAME = "3.40.0"
    API_VERSION_KEY = "d175f967bb6444f49ff08b3b6d646527"
    API_KEY = "ccd59ee269c01511ba763467045c115779fcae3050238a252f1bd1a4b65cfec6"
    SHARED_KEY = "yayZ1"
    STORE_KEY = "yayZ1payment"
    ID_CARD_CHECK_SECRET_KEY = "4aa6d1c301a97154bc1098c2"
    REVIEW_HOST_1 = "review.yay.space"
    REVIEW_HOST_2 = "cas-stg.yay.space"
    STAGING_HOST_1 = "stg.yay.space"
    STAGING_HOST_2 = "cas.yay.space"
    CABLE_HOST = "cable.yay.space"
    CONFIG_HOST = "settings.yay.space"
    PRODUCTION_HOST = "api.yay.space"
    ID_CARD_CHECK_HOST_PRODUCTION = "idcardcheck.com"
    ID_CARD_CHECK_HOST_STAGING = "stg.idcardcheck.com"
    DEVICE = Device(
        device_type="android",
        os_version="11",
        screen_density="3.5",
        screen_size="1440x2960",
        model="Galaxy S9",
    )
    USER_AGENT = f"{DEVICE.device_type} {DEVICE.os_version} ({DEVICE.screen_density}x {DEVICE.screen_size} {DEVICE.model})"
    DEVICE_INFO = "yay " + VERSION_NAME + " " + USER_AGENT
