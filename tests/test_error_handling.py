import json
import unittest
from json.decoder import JSONDecodeError
from unittest.mock import AsyncMock, MagicMock

from yaylib.errors import (
    AccessTokenExpiredError,
    AccessTokenInvalidError,
    HTTPAuthenticationError,
    HTTPBadRequestError,
    HTTPForbiddenError,
    HTTPInternalServerError,
    HTTPNotFoundError,
    HTTPRateLimitError,
    QuotaLimitExceededError,
    TooManyRequestsError,
    UnauthorizedError,
    raise_for_code,
    raise_for_status,
)


class MockResponse:
    """aiohttp.ClientResponseのモック"""

    def __init__(self, status=200, data=None, headers=None):
        self.status = status
        self._data = data or {}
        self.headers = headers or {}
        self.content = MagicMock()
        self.content.read = AsyncMock(return_value=b"")

    async def json(self, content_type=None):
        return self._data

    async def text(self):
        return json.dumps(self._data)

    async def read(self):
        return b""


class TestErrorHandling(unittest.IsolatedAsyncioTestCase):
    async def test_raise_for_status_success(self):
        """成功レスポンスでのraise_for_statusをテスト"""
        response = MockResponse(status=200)
        # 例外が発生しないはず
        await raise_for_status(response)

    async def test_raise_for_status_server_error(self):
        """サーバーエラーでのraise_for_statusをテスト"""
        response = MockResponse(status=500)
        with self.assertRaises(HTTPInternalServerError):
            await raise_for_status(response)

    async def test_raise_for_code_success(self):
        """成功レスポンスでのraise_for_codeをテスト"""
        response = MockResponse(status=200, data={"code": 0, "message": "Success"})
        # 例外が発生しないはず
        await raise_for_code(response)

    async def test_raise_for_code_unauthorized(self):
        """認証エラーでのraise_for_codeをテスト"""
        response = MockResponse(
            status=401,  # なんでもいい
            data={"result": "error", "error_code": 401, "message": "Unauthorized"},
        )
        with self.assertRaises(UnauthorizedError):
            await raise_for_code(response)

    async def test_raise_for_code_access_token_expired(self):
        """アクセストークン期限切れエラーでのraise_for_codeをテスト"""
        response = MockResponse(
            status=401,  # なんでもいい
            data={
                "result": "error",
                "error_code": -3,
                "message": "Access token expired",
            },
        )
        with self.assertRaises(AccessTokenExpiredError):
            await raise_for_code(response)

    async def test_raise_for_code_access_token_invalid(self):
        """アクセストークン無効エラーでのraise_for_codeをテスト"""
        response = MockResponse(
            status=401,  # なんでもいい
            data={
                "result": "error",
                "error_code": -18,
                "message": "Access token invalid",
            },
        )
        with self.assertRaises(AccessTokenInvalidError):
            await raise_for_code(response)

    async def test_raise_for_code_too_many_requests(self):
        """リクエスト過多エラーでのraise_for_codeをテスト"""
        response = MockResponse(
            status=401,  # なんでもいい
            data={"result": "error", "error_code": 429, "message": "Too many requests"},
        )
        with self.assertRaises(TooManyRequestsError):
            await raise_for_code(response)

    async def test_raise_for_code_quota_limit_exceeded(self):
        """クォータ制限超過エラーでのraise_for_codeをテスト"""
        response = MockResponse(
            status=401,  # なんでもいい
            data={
                "result": "error",
                "error_code": -343,
                "message": "Quota limit exceeded",
            },
        )
        with self.assertRaises(QuotaLimitExceededError):
            await raise_for_code(response)

    async def test_raise_for_code_json_error(self):
        """JSONパースエラーでのraise_for_codeをテスト"""
        # json()が呼び出されたときに例外を発生させるレスポンスを作成
        response = MockResponse(status=200)
        # ContentTypeErrorの代わりにJSONDecodeErrorを使用
        response.json = AsyncMock(side_effect=JSONDecodeError("Invalid JSON", "", 0))

        # 例外が発生しないはず（JSONエラーは無視される）
        await raise_for_code(response)

    async def test_raise_for_code_missing_code(self):
        """レスポンスにコードが欠けている場合のraise_for_codeをテスト"""
        response = MockResponse(status=200, data={"message": "No code field"})

        # 例外が発生しないはず（コードの欠落は無視される）
        await raise_for_code(response)

    async def test_http_bad_request_error(self):
        """HTTPBadRequestErrorの発生をテスト"""
        response = MockResponse(status=400)
        with self.assertRaises(HTTPBadRequestError):
            await raise_for_status(response)

    async def test_http_authentication_error(self):
        """HTTPAuthenticationErrorの発生をテスト"""
        response = MockResponse(status=401)
        with self.assertRaises(HTTPAuthenticationError):
            await raise_for_status(response)

    async def test_http_forbidden_error(self):
        """HTTPForbiddenErrorの発生をテスト"""
        response = MockResponse(status=403)
        with self.assertRaises(HTTPForbiddenError):
            await raise_for_status(response)

    async def test_http_not_found_error(self):
        """HTTPNotFoundErrorの発生をテスト"""
        response = MockResponse(status=404)
        with self.assertRaises(HTTPNotFoundError):
            await raise_for_status(response)

    async def test_http_rate_limit_error(self):
        """HTTPRateLimitErrorの発生をテスト"""
        response = MockResponse(status=429)
        with self.assertRaises(HTTPRateLimitError):
            await raise_for_status(response)


if __name__ == "__main__":
    unittest.main()
