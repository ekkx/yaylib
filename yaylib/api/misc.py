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

import os
from datetime import datetime
from io import BytesIO
from typing import List
from urllib import parse

from PIL import Image

from .. import config
from ..constants import ImageType
from ..models import Attachment
from ..responses import (
    ApplicationConfigResponse,
    BanWordsResponse,
    EmailGrantTokenResponse,
    EmailVerificationPresignedUrlResponse,
    IdCheckerPresignedUrlResponse,
    PolicyAgreementsResponse,
    PopularWordsResponse,
    PresignedUrlResponse,
    PresignedUrlsResponse,
    Response,
    WebSocketTokenResponse,
)
from ..utils import (
    generate_uuid,
    get_hashed_filename,
    is_valid_image_format,
    is_valid_video_format,
)


class MiscApi:
    """未分類 API"""

    def __init__(self, client) -> None:
        # pylint: disable=import-outside-toplevel
        from ..client import Client

        self.__client: Client = client

    async def accept_policy_agreement(self, agreement_type: str) -> Response:
        """利用規約、ポリシー同意書に同意する

        Args:
            agreement_type (str):

        Returns:
            Response:
        """
        return await self.__client.request(
            "POST",
            config.API_HOST + f"/v1/users/policy_agreements/{agreement_type}",
            return_type=Response,
        )

    async def send_verification_code(
        self, email: str, intent: str, locale: str = "ja"
    ) -> Response:
        """メールアドレス認証コードを送信する

        Args:
            email (str):
            intent (str):
            locale (str):

        Returns:
            Response:
        """
        return await self.__client.request(
            "POST",
            self.get_email_verification_presigned_url(
                email=email, locale=locale, intent=intent
            ).url,
            json={"locale": locale, "email": email},
            return_type=Response,
        )

    async def get_email_grant_token(self, **params) -> EmailGrantTokenResponse:
        """メールアドレス認証トークンを取得する

        Args:
            code (int):
            email (str):

        Returns:
            EmailGrantTokenResponse:
        """
        return await self.__client.request(
            "POST",
            config.ID_CARD_CHECK_BASE_HOST + "/apis/v1/apps/yay/email_grant_tokens",
            json=params,
            return_type=EmailGrantTokenResponse,
        )

    async def get_email_verification_presigned_url(
        self, email: str, locale: str = "ja", intent: str = None
    ) -> EmailVerificationPresignedUrlResponse:
        """メールアドレス確認用の署名付きURLを取得する

        Args:
            email (str):
            locale (str):
            intent (str, optional):

        Returns:
            EmailVerificationPresignedUrlResponse:
        """
        return await self.__client.request(
            "POST",
            config.API_HOST + "/v1/email_verification_urls",
            json={
                "device_uuid": self.__client.device_uuid,
                "email": email,
                "locale": locale,
                "intent": intent,
            },
            return_type=EmailVerificationPresignedUrlResponse,
        )

    async def get_file_upload_presigned_urls(
        self, file_names: List[str]
    ) -> PresignedUrlsResponse:
        """ファイルアップロード用の署名付きURLを取得する

        Args:
            file_names (List[str]):

        Returns:
            PresignedUrlsResponse:
        """
        return await self.__client.request(
            "GET",
            config.API_HOST + "/v1/buckets/presigned_urls",
            params={"file_names[]": file_names},
            return_type=PresignedUrlsResponse,
        )

    async def get_id_checker_presigned_url(
        self, model: str, action: str, **params
    ) -> IdCheckerPresignedUrlResponse:
        # @QueryMap @NotNull Map<String, String> map
        """身分証明用の署名付きURLを取得する

        Args:
            model (str):
            action (str):

        Returns:
            IdCheckerPresignedUrlResponse:
        """
        return await self.__client.request(
            "GET",
            config.API_HOST + f"/v1/id_check/{model}/{action}",
            params=params,
            return_type=IdCheckerPresignedUrlResponse,
        )

    async def get_old_file_upload_presigned_url(
        self, video_file_name: str
    ) -> PresignedUrlResponse:
        """旧版ファイルアップロード用の署名付きURLを取得する

        Args:
            video_file_name (str):

        Returns:
            PresignedUrlResponse:
        """
        return await self.__client.request(
            "GET",
            config.API_HOST + "/v1/users/presigned_url",
            params={"video_file_name": video_file_name},
            return_type=PresignedUrlResponse,
        )

    async def get_policy_agreed(self) -> PolicyAgreementsResponse:
        """利用規約、ポリシー同意書に同意しているかどうかを取得する

        Returns:
            PolicyAgreementsResponse:
        """
        return await self.__client.request(
            "GET",
            config.API_HOST + "/v1/users/policy_agreements",
            return_type=PolicyAgreementsResponse,
        )

    async def get_web_socket_token(self) -> WebSocketTokenResponse:
        """WebSocket トークンを取得する

        Returns:
            WebSocketTokenResponse:
        """
        return await self.__client.request(
            "GET",
            config.API_HOST + "/v1/users/ws_token",
            return_type=WebSocketTokenResponse,
        )

    async def upload_image(
        self, image_paths: List[str], image_type: str
    ) -> List[Attachment]:
        """画像をアップロードして、サーバー上のファイルのリストを取得する

        Examples:
            投稿に画像を付与する場合

            >>> # サーバー上にアップロード
            >>> attachments = client.upload_image(
            >>>     image_type=yaylib.ImageType.POST,
            >>>     image_paths=["./example.jpg"],
            >>> )
            >>> # サーバー上のファイル名を指定
            >>> client.create_post(
            >>>     "Hello with yaylib!",
            >>>     attachment_filename=attachments[0].filename
            >>> )

        Args:
            image_paths (List[str]): 画像ファイルのパスのリスト
            image_type (str): 画像の種類

        Raises:
            ValueError: 画像タイプやフォーマットが不正な場合

        Returns:
            List[Attachment]: サーバー上のファイル情報
        """
        if image_type not in [
            "post",
            "chat_message",
            "chat_background",
            "report",
            "user_avatar",
            "user_cover",
            "group_cover",
            "group_thread_icon",
            "group_icon",
        ]:
            raise ValueError(f"Invalid image type. [{image_type}]")

        _files = []

        for key, image_path in enumerate(image_paths):
            filename, extension = os.path.splitext(image_path)

            if not is_valid_image_format(extension):
                raise ValueError(f"Invalid image format. [{filename + extension}]")

            image = Image.open(image_path)
            natural_width, natural_height = image.size

            resized_image = Image.open(image_path)

            if extension != ".gif" and image_type == ImageType.USER_AVATAR:
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
        res_presigned_url: PresignedUrlsResponse = (
            await self.get_file_upload_presigned_urls(file_names)
        )
        presigned_urls = res_presigned_url.presigned_urls

        res_upload: List[Attachment] = []

        x: Attachment
        for x in _files:
            p_url = next(
                (p.url for p in presigned_urls if x.filename in p.filename), None
            )
            if not p_url:
                continue

            image_data = BytesIO()
            if x.file.format == "GIF" and x.file.is_animated:
                x.file.save(image_data, format=x.file.format, save_all=True, disposal=2)
            else:
                x.file.save(image_data, format=x.file.format)
            image_data.seek(0)

            await self.__client.base_request("PUT", p_url, data=image_data.read())

            x.filename = parse.urlsplit(p_url).path.replace("/uploads/", "")

            res_upload.append(x)

        return res_upload

    async def upload_video(self, video_path: str) -> str:
        """動画をアップロードして、サーバー上のファイル名を取得する

        Examples:
            投稿に動画を付与する場合

            >>> # サーバー上にアップロード
            >>> filename = client.upload_video("./example.mp4")
            >>> # サーバー上のファイル名を指定
            >>> client.create_post(
            >>>     "Hello with yaylib!",
            >>>     video_file_name=filename
            >>> )

        Args:
            video_path (str): 動画ファイルのパス

        Raises:
            ValueError: 動画フォーマットが不正な場合

        Returns:
            str: サーバー上のファイル名
        """
        filename, extension = os.path.splitext(video_path)

        if not is_valid_video_format(extension):
            raise ValueError(f"Invalid video format. [{filename + extension}]")

        uuid = generate_uuid(False)[:16]
        filename = f"{uuid}_{int(datetime.now().timestamp())}{extension}"

        res_presigned_url: PresignedUrlResponse = (
            await self.get_old_file_upload_presigned_url(filename)
        )
        presigned_url = res_presigned_url.presigned_url

        with open(video_path, "br") as f:
            video = f.read()

        self.__client.logger.debug(f"Uploading video: {video_path}")

        await self.__client.base_request("PUT", presigned_url, data=video)

        return filename

    # config

    async def get_app_config(self) -> ApplicationConfigResponse:
        """アプリケーションのメタデータを取得する

        Returns:
            ApplicationConfigResponse:
        """
        return await self.__client.request(
            "GET",
            config.CONFIG_HOST + "/api/apps/yay",
            return_type=ApplicationConfigResponse,
        )

    async def get_banned_words(self, country_code: str = "jp") -> BanWordsResponse:
        """禁止ワードの一覧を取得する

        Args:
            country_code (str, optional):

        Returns:
            BanWordsResponse:
        """
        return await self.__client.request(
            "GET",
            config.CONFIG_HOST + f"/{country_code}/api/v2/banned_words",
            return_type=BanWordsResponse,
        )

    async def get_popular_words(self, country_code: str = "jp") -> PopularWordsResponse:
        """人気ワードの一覧を取得する

        Args:
            country_code (str, optional):

        Returns:
            PopularWordsResponse:
        """
        return await self.__client.request(
            "GET",
            config.CONFIG_HOST + f"/{country_code}/api/apps/yay/popular_words",
            return_type=PopularWordsResponse,
        )
