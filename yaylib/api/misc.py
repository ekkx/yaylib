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

from __future__ import annotations

import os
import httpx
from datetime import datetime
from io import BytesIO
from PIL import Image
from urllib import parse

from .. import client
from ..config import Configs
from ..models import Attachment
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
    PolicyAgreementsResponse,
)
from ..types import ImageType
from ..utils import (
    is_valid_image_format,
    is_valid_video_format,
    get_hashed_filename,
    generate_uuid,
)


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


class MiscAPI(object):
    def __init__(self, base: client.BaseClient) -> None:
        self.__base = base

    def accept_policy_agreement(self, type: str):
        return self.__base._request(
            "POST",
            route=f"/v1/users/policy_agreements/{type}",
            bypass_delay=True,
        )

    def generate_sns_thumbnail(self, **params):
        """

        Parameters:
        -----------

            - resource_type: str - (Required)
            - resource_id: int - (Required)

        """
        return self.__base._request(
            "GET", route=f"/v1/sns_thumbnail/generate", params=params
        )

    def send_verification_code(self, email: str, intent: str, locale: str):
        return self.__base._request(
            "POST",
            route=self.get_email_verification_presigned_url(
                email=email, locale=locale, intent=intent
            ).url,
            payload={"locale": "ja", "email": email},
            base_url=None,
        )

    def get_email_grant_token(self, code: int, email: str) -> EmailGrantTokenResponse:
        return self.__base._request(
            "POST",
            base_url=Configs.ID_CARD_CHECK_HOST_PRODUCTION,
            route="/apis/v1/apps/yay/email_grant_tokens",
            payload={"code": code, "email": email},
            data_type=EmailGrantTokenResponse,
        )

    def get_email_verification_presigned_url(
        self, email: str, locale: str, intent: str = None
    ) -> str:
        return self.__base._request(
            "POST",
            route="/v1/email_verification_urls",
            payload={
                "device_uuid": self.__base.device_uuid,
                "email": email,
                "locale": locale,
                "intent": intent,
            },
            data_type=EmailVerificationPresignedUrlResponse,
            bypass_delay=True,
        )

    def get_file_upload_presigned_urls(
        self, file_names: list[str]
    ) -> PresignedUrlsResponse:
        return self.__base._request(
            "GET",
            route="/v1/buckets/presigned_urls",
            params={"file_names[]": file_names},
            data_type=PresignedUrlsResponse,
            bypass_delay=True,
        )

    def get_id_checker_presigned_url(
        self, model: str, action: str, **params
    ) -> IdCheckerPresignedUrlResponse:
        # TODO: @QueryMap @NotNull Map<String, String> map
        """
        Meow..
        """
        return self.__base._request(
            "GET",
            route=f"/v1/id_check/{model}/{action}",
            params=params,
            data_type=IdCheckerPresignedUrlResponse,
            bypass_delay=True,
        )

    def get_old_file_upload_presigned_url(
        self, video_file_name: str
    ) -> PresignedUrlResponse:
        return self.__base._request(
            "GET",
            route=f"/v1/users/presigned_url",
            params={"video_file_name": video_file_name},
            data_type=PresignedUrlResponse,
            bypass_delay=True,
        )

    def get_policy_agreements(self) -> PolicyAgreementsResponse:
        return self.__base._request(
            "GET",
            route=f"/v1/users/policy_agreements",
            data_type=PolicyAgreementsResponse,
            bypass_delay=True,
        )

    def get_web_socket_token(self, headers: dict = None) -> WebSocketTokenResponse:
        return self.__base._request(
            "GET",
            route=f"/v1/users/ws_token",
            data_type=WebSocketTokenResponse,
            headers=headers,
            bypass_delay=True,
        )

    def verify_device(
        self,
        app_version: str,
        device_uuid: str,
        platform: str,
        verification_string: str,
    ) -> VerifyDeviceResponse:
        # TODO: check platform, verification_string
        return self.__base._request(
            "POST",
            route="/v1/genuine_devices/verify",
            payload={
                "app_version": app_version,
                "device_uuid": device_uuid,
                "platform": platform,
                "verification_string": verification_string,
            },
            data_type=VerifyDeviceResponse,
        )

    def upload_image(self, image_paths: list[str], image_type: str) -> list[Attachment]:
        """

        画像をアップロードして、サーバー上のファイル名のリストを返します。

        Parameteres:
        ------------

            - image_path: list[str] - (required): 画像パスのリスト
            - image_type: str - (required): 画像の種類

        Examples
        --------

        投稿に画像を付与する場合

        >>> # サーバー上にアップロード
        >>> attachments = api.upload_image(
        >>>     image_type=yaylib.IMAGE_TYPE_POST,
        >>>     image_paths=["./example.jpg"],
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

            resized_image = Image.open(image_path)

            if extension != ".gif" and image_type == ImageType.user_avatar:
                resized_image.thumbnail((200, 200))

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

            uuid = generate_uuid(False)[:16]

            original_attachment.filename = get_hashed_filename(
                original_attachment, image_type, key, uuid
            )
            thumbnail_attachment.filename = get_hashed_filename(
                thumbnail_attachment, image_type, key, uuid
            )

            _files.append(original_attachment)
            _files.append(thumbnail_attachment)

        file_names = [x.filename for x in _files]
        res_presigned_url = self.get_file_upload_presigned_urls(
            file_names
        ).presigned_urls

        res_upload = []

        x: Attachment
        for x in _files:
            p_url = next(
                (p.url for p in res_presigned_url if x.filename in p.filename), None
            )
            if not p_url:
                continue

            image_data = BytesIO()
            if x.file.format == "GIF" and x.file.is_animated:
                x.file.save(image_data, format=x.file.format, save_all=True)
            else:
                x.file.save(image_data, format=x.file.format)
            image_data.seek(0)

            response = httpx.put(p_url, data=image_data.read())
            response.raise_for_status()

            x.filename = parse.urlsplit(p_url).path.replace("/uploads/", "")

            res_upload.append(x)

        return res_upload

    def upload_video(self, video_path: str) -> str:
        filename, extension = os.path.splitext(video_path)

        if not is_valid_video_format(extension):
            raise ValueError(f"Invalid video format. [{filename + extension}]")

        uuid = generate_uuid(False)[:16]
        filename = f"{uuid}_{int(datetime.now().timestamp())}{extension}"

        res_presigned_url = self.get_old_file_upload_presigned_url(
            filename
        ).presigned_url

        with open(video_path, "br") as f:
            video = f.read()

        self.__base.logger.debug(f"Uploading video: {video_path}")

        response = httpx.put(res_presigned_url, data=video)
        response.raise_for_status()

        return filename

    # config

    def get_app_config(self) -> ApplicationConfigResponse:
        return self.__base._request(
            "GET",
            route=f"https://{Configs.CONFIG_HOST}/api/apps/yay",
            data_type=ApplicationConfigResponse,
        )

    def get_banned_words(self, country_code: str = "jp") -> BanWordsResponse:
        return self.__base._request(
            "GET",
            route=f"https://{Configs.CONFIG_HOST}/{country_code}/api/v2/banned_words",
            data_type=BanWordsResponse,
        )

    def get_popular_words(self, country_code: str = "jp") -> PopularWordsResponse:
        return self.__base._request(
            "GET",
            route=f"https://{Configs.CONFIG_HOST}/{country_code}/api/apps/yay/popular_words",
            data_type=PopularWordsResponse,
        )
