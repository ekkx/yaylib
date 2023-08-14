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
import httpx
from io import BytesIO
from PIL import Image
from typing import List
from urllib import parse

from ..config import Endpoints, Configs
from ..models import ApplicationConfig, Attachment, BanWord, PresignedUrl, PopularWord
from ..responses import (
    EmailGrantTokenResponse,
    EmailVerificationPresignedUrlResponse,
    PresignedUrlResponse,
    PresignedUrlsResponse,
    IdCheckerPresignedUrlResponse,
    VerifyDeviceResponse,
    WebSocketTokenResponse,
    ApplicationConfigResponse,
    BanWordsResponse,
    PopularWordsResponse,
)
from ..utils import is_valid_image_format, get_hashed_filename


upload_item_types = [
    "post",
    "chat_message",
    "chat_background",
    "report",
    "user_avatar",
    "user_cover",
    "group_cover",
    "group_thread_icon",
    "group_icon",
]


def accept_policy_agreement(self, type: str, access_token: str = None):
    response = self._make_request(
        "POST",
        endpoint=f"{Endpoints.USERS_V1}/policy_agreements/{type}",
        access_token=access_token,
    )
    self.logger.info("Accepted to policy agreement.")
    return response


def generate_sns_thumbnail(self, access_token: str = None, **params):
    """

    Parameters:
    ----------

        - resource_type: str - (Required)
        - resource_id: int - (Required)

    """
    response = self._make_request(
        "GET",
        endpoint=f"{Endpoints.SNS_THUMBNAIL_V1}/generate",
        params=params,
        access_token=access_token,
    )
    self.logger.info("SNS thumbnail generated.")
    return response


def send_verification_code(self, email: str):
    response = self._make_request(
        "POST",
        endpoint=self.get_email_verification_presigned_url(
            email=email, locale="ja"
        ).url,
        payload={"locale": "ja", "email": email},
    )
    self.logger.info(f"Verification code successfully sent to '{email}'.")
    return response


def get_email_grant_token(self, code: int, email: str) -> EmailGrantTokenResponse:
    return self._make_request(
        "POST",
        endpoint=f"{Endpoints.GET_EMAIL_GRANT_TOKEN}",
        payload={"code": code, "email": email},
        data_type=EmailGrantTokenResponse,
    ).email_grant_token


def get_email_verification_presigned_url(
    self, email: str, locale: str, intent: str = None, access_token: str = None
) -> str:
    return self._make_request(
        "POST",
        endpoint=f"{Endpoints.EMAIL_VERIFICATION_URL_V1}",
        payload={
            "device_uuid": self.device_uuid,
            "email": email,
            "locale": locale,
            "intent": intent,
        },
        data_type=EmailVerificationPresignedUrlResponse,
        access_token=access_token,
    )


def get_file_upload_presigned_urls(
    self, file_names: List[str], access_token: str = None
) -> List[PresignedUrl]:
    return self._make_request(
        "GET",
        endpoint=f"{Endpoints.BUCKETS_V1}/presigned_urls",
        params={"file_names[]": file_names},
        data_type=PresignedUrlsResponse,
        access_token=access_token,
    ).presigned_urls


def get_id_checker_presigned_url(
    self, model: str, action: str, access_token: str = None, **params
) -> str:
    # TODO: @QueryMap @NotNull Map<String, String> map
    """
    Meow..
    """
    return self._make_request(
        "GET",
        endpoint=f"{Endpoints.ID_CHECK_V1}/{model}/{action}",
        params=params,
        data_type=IdCheckerPresignedUrlResponse,
        access_token=access_token,
    ).presigned_url


def get_old_file_upload_presigned_url(
    self, video_file_name: str, access_token: str = None
) -> str:
    return self._make_request(
        "GET",
        endpoint=f"{Endpoints.USERS_V1}/presigned_url",
        params={"video_file_name": video_file_name},
        data_type=PresignedUrlResponse,
        access_token=access_token,
    ).presigned_url


def get_web_socket_token(self, headers: dict = None, access_token: str = None) -> str:
    self._check_authorization(access_token)
    return self._make_request(
        "GET",
        endpoint=f"{Endpoints.USERS_V1}/ws_token",
        data_type=WebSocketTokenResponse,
        headers=headers,
        access_token=access_token,
    ).token


def verify_device(
    self,
    app_version: str,
    device_uuid: str,
    platform: str,
    verification_string: str,
    access_token: str = None,
) -> VerifyDeviceResponse:
    # TODO: check platform, verification_string
    response = self._make_request(
        "POST",
        endpoint=f"{Endpoints.GENUINE_DEVICES_V1}/verify",
        payload={
            "app_version": app_version,
            "device_uuid": device_uuid,
            "platform": platform,
            "verification_string": verification_string,
        },
        data_type=VerifyDeviceResponse,
        access_token=access_token,
    )
    self.logger.info("Device has been verified.")
    return response


def upload_image(self, image_paths: List[str], image_type: str) -> List[Attachment]:
    """

    画像をアップロードして、サーバー上のファイル名のリストを返します。

    Parameteres
    -----------

        - image_path: List[str] - (required): 画像パスのリスト
        - image_type: str - (required): 画像の種類

    Examples
    --------

    投稿に画像を付与する場合

    >>> # サーバー上にアップロード
    >>> attachments = api.upload_image(
    >>>     image_type=yaylib.IMAGE_TYPE_POST,
    >>>     image_paths=["./test.jpg"],
    >>> )
    >>> # サーバー上のファイル名を指定
    >>> api.create_post(
    >>>     "Hello with yaylib!",
    >>>     attachment_filename=attachments[0].filename
    >>> )

    """
    if image_type not in upload_item_types:
        raise TypeError(f"Invalid image type. [{image_type}]")

    _files = []

    for key, image_path in enumerate(image_paths):
        filename, extension = os.path.splitext(image_path)

        if not is_valid_image_format(extension):
            raise ValueError(f"Invalid image format. [{filename + extension}]")

        image = Image.open(image_path)
        natural_width, natural_height = image.size
        resized_image = image

        if extension != ".gif":
            resized_image.thumbnail((450, 450))

        original_attachment = Attachment(
            file=image,
            filename="",
            original_file_name=filename,
            original_file_extension=extension,
            natural_width=natural_width,
            natural_height=natural_height,
            is_thumb=False,
        )

        thumbnail_attachment = Attachment(
            file=resized_image,
            filename="",
            original_file_name=filename,
            original_file_extension=extension,
            natural_width=natural_width,
            natural_height=natural_height,
            is_thumb=True,
        )

        uuid = self.generate_uuid()[:8]

        original_attachment.filename = get_hashed_filename(
            original_attachment, image_type, key, uuid
        )
        thumbnail_attachment.filename = get_hashed_filename(
            thumbnail_attachment, image_type, key, uuid
        )

        _files.append(original_attachment)
        _files.append(thumbnail_attachment)

    file_names = [x.filename for x in _files]
    res_presigned_url = get_file_upload_presigned_urls(self, file_names)

    res_upload = []

    for x in _files:
        p_url = next(
            (p.url for p in res_presigned_url if x.filename in p.filename), None
        )
        if not p_url:
            continue

        image_data = BytesIO()
        image.save(image_data, format=x.file.format)
        image_data.seek(0)

        response = httpx.put(p_url, data=image_data.read())
        response.raise_for_status()

        x.filename = parse.urlsplit(p_url).path.replace("/uploads/", "")

        res_upload.append(x)

    return res_upload


def upload_video(self, video_path: str, access_token: str = None):
    pass


# config


def get_app_config(self) -> ApplicationConfig:
    response = self._make_request(
        "GET",
        endpoint=f"https://{Configs.YAY_CONFIG_HOST}/api/apps/yay",
        data_type=ApplicationConfigResponse,
    ).app
    return response


def get_banned_words(self, country_code: str = "jp") -> List[BanWord]:
    response = self._make_request(
        "GET",
        endpoint=f"https://{Configs.YAY_CONFIG_HOST}/{country_code}/api/v2/banned_words",
        data_type=BanWordsResponse,
    ).ban_words
    return response


def get_popular_words(self, country_code: str = "jp") -> List[PopularWord]:
    response = self._make_request(
        "GET",
        endpoint=f"https://{Configs.YAY_CONFIG_HOST}/{country_code}/api/apps/yay/popular_words",
        data_type=PopularWordsResponse,
    ).popular_words
    return response
