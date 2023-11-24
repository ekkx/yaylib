"""
MIT License

Copyright (c) 2023-present qvco

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
    YAYLIB_VERSION = "1.3.4"
    API_VERSION_NAME = "3.29"
    VERSION_NAME = "3.29.0"
    YAY_API_VERSION_KEY = "6fa97fc2c3d04955bb8320f2d080593a"
    YAY_API_KEY = "ccd59ee269c01511ba763467045c115779fcae3050238a252f1bd1a4b65cfec6"
    YAY_SHARED_KEY = "yayZ1"
    YAY_STORE_KEY = "yayZ1payment"
    ID_CARD_CHECK_SECRET_KEY = "4aa6d1c301a97154bc1098c2"
    YAY_REVIEW_HOST_1 = "review.yay.space"
    YAY_REVIEW_HOST_2 = "cas-stg.yay.space"
    YAY_STAGING_HOST_1 = "stg.yay.space"
    YAY_STAGING_HOST_2 = "cas.yay.space"
    YAY_CABLE_HOST = "cable.yay.space"
    YAY_CONFIG_HOST = "settings.yay.space"
    YAY_PRODUCTION_HOST = "api.yay.space"
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


class Endpoints:
    BASE_API_URL = "https://" + Configs.YAY_PRODUCTION_HOST

    # api v1 endpoints
    USERS_V1 = BASE_API_URL + "/v1/users"
    PAYMENTS_V1 = BASE_API_URL + "/v1/payments"
    THREADS_V1 = BASE_API_URL + "/v1/threads"
    PINNED_V1 = BASE_API_URL + "/v1/pinned"
    POSTS_V1 = BASE_API_URL + "/v1/posts"
    CONVERSATIONS_V1 = BASE_API_URL + "/v1/conversations"
    HIDDEN_V1 = BASE_API_URL + "/v1/hidden"
    GROUPS_V1 = BASE_API_URL + "/v1/groups"
    CHAT_ROOMS_V1 = BASE_API_URL + "/v1/chat_rooms"
    CALLS_V1 = BASE_API_URL + "/v1/calls"
    SURVEYS_V1 = BASE_API_URL + "/v1/surveys"
    FRIENDS_V1 = BASE_API_URL + "/v1/friends"
    GAMES_V1 = BASE_API_URL + "/v1/games"
    GENRES_V1 = BASE_API_URL + "/v1/genres"
    ANONYMOUS_CALLS_V1 = BASE_API_URL + "/v1/anonymous_calls"
    NOTIFICATION_SETTINGS_V1 = BASE_API_URL + "/v1/notification_settings"
    STICKER_PACKS_V1 = BASE_API_URL + "/v1/sticker_packs"
    SNS_THUMBNAIL_V1 = BASE_API_URL + "/v1/sns_thumbnail"
    EMAIL_VERIFICATION_URL_V1 = BASE_API_URL + "/v1/email_verification_urls"
    BUCKETS_V1 = BASE_API_URL + "/v1/buckets"
    ID_CHECK_V1 = BASE_API_URL + "/v1/id_check"
    PROMOTIONS_V1 = BASE_API_URL + "/v1/promotions"
    SKYFALL_V1 = BASE_API_URL + "/v1/skyfall"
    GENUINE_DEVICES_V1 = BASE_API_URL + "/v1/genuine_devices"
    LOBI_FRIENDS_V1 = BASE_API_URL + "/v1/lobi_friends"
    WEB_V1 = BASE_API_URL + "/v1/web"

    # api v2 endpoints
    USERS_V2 = BASE_API_URL + "/v2/users"
    PAYMENTS_V2 = BASE_API_URL + "/v2/payments"
    THREADS_V2 = BASE_API_URL + "/v2/threads"
    PINNED_V2 = BASE_API_URL + "/v2/pinned"
    POSTS_V2 = BASE_API_URL + "/v2/posts"
    CONVERSATIONS_V2 = BASE_API_URL + "/v2/conversations"
    HIDDEN_V2 = BASE_API_URL + "/v2/hidden"
    GROUPS_V2 = BASE_API_URL + "/v2/groups"
    CHAT_ROOMS_V2 = BASE_API_URL + "/v2/chat_rooms"
    CALLS_V2 = BASE_API_URL + "/v2/calls"
    SURVEYS_V2 = BASE_API_URL + "/v2/surveys"
    FRIENDS_V2 = BASE_API_URL + "/v2/friends"
    GAMES_V2 = BASE_API_URL + "/v2/games"
    GENRES_V2 = BASE_API_URL + "/v2/genres"
    ANONYMOUS_CALLS_V2 = BASE_API_URL + "/v2/anonymous_calls"
    NOTIFICATION_SETTINGS_V2 = BASE_API_URL + "/v2/notification_settings"
    STICKER_PACKS_V2 = BASE_API_URL + "/v2/sticker_packs"
    SNS_THUMBNAIL_V2 = BASE_API_URL + "/v2/sns_thumbnail"
    EMAIL_VERIFICATION_URL_V2 = BASE_API_URL + "/v2/email_verification_urls"
    BUCKETS_V2 = BASE_API_URL + "/v2/buckets"
    ID_CHECK_V2 = BASE_API_URL + "/v2/id_check"
    PROMOTIONS_V2 = BASE_API_URL + "/v2/promotions"
    SKYFALL_V2 = BASE_API_URL + "/v2/skyfall"
    GENUINE_DEVICES_V2 = BASE_API_URL + "/v2/genuine_devices"
    LOBI_FRIENDS_V2 = BASE_API_URL + "/v2/lobi_friends"
    WEB_V2 = BASE_API_URL + "/v2/web"

    # api v3 endpoints
    USERS_V3 = BASE_API_URL + "/v3/users"
    PAYMENTS_V3 = BASE_API_URL + "/v3/payments"
    THREADS_V3 = BASE_API_URL + "/v3/threads"
    PINNED_V3 = BASE_API_URL + "/v3/pinned"
    POSTS_V3 = BASE_API_URL + "/v3/posts"
    CONVERSATIONS_V3 = BASE_API_URL + "/v3/conversations"
    HIDDEN_V3 = BASE_API_URL + "/v3/hidden"
    GROUPS_V3 = BASE_API_URL + "/v3/groups"
    CHAT_ROOMS_V3 = BASE_API_URL + "/v3/chat_rooms"
    CALLS_V3 = BASE_API_URL + "/v3/calls"
    SURVEYS_V3 = BASE_API_URL + "/v3/surveys"
    FRIENDS_V3 = BASE_API_URL + "/v3/friends"
    GAMES_V3 = BASE_API_URL + "/v3/games"
    GENRES_V3 = BASE_API_URL + "/v3/genres"
    ANONYMOUS_CALLS_V3 = BASE_API_URL + "/v3/anonymous_calls"
    NOTIFICATION_SETTINGS_V3 = BASE_API_URL + "/v3/notification_settings"
    STICKER_PACKS_V3 = BASE_API_URL + "/v3/sticker_packs"
    SNS_THUMBNAIL_V3 = BASE_API_URL + "/v3/sns_thumbnail"
    EMAIL_VERIFICATION_URL_V3 = BASE_API_URL + "/v3/email_verification_urls"
    BUCKETS_V3 = BASE_API_URL + "/v3/buckets"
    ID_CHECK_V3 = BASE_API_URL + "/v3/id_check"
    PROMOTIONS_V3 = BASE_API_URL + "/v3/promotions"
    SKYFALL_V3 = BASE_API_URL + "/v3/skyfall"
    GENUINE_DEVICES_V3 = BASE_API_URL + "/v3/genuine_devices"
    LOBI_FRIENDS_V3 = BASE_API_URL + "/v3/lobi_friends"
    WEB_V3 = BASE_API_URL + "/v3/web"

    # misc
    GET_EMAIL_GRANT_TOKEN = (
        "https://"
        + Configs.ID_CARD_CHECK_HOST_PRODUCTION
        + "/apis/v1/apps/yay/email_grant_tokens"
    )
