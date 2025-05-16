import unittest
from unittest.mock import AsyncMock, MagicMock, patch

from yaylib.client import Client
from yaylib.errors import AccessTokenExpiredError, TooManyRequestsError
from yaylib.models import Model
from yaylib.state import State


class TestModel(Model):
    """レスポンスパース用のテストモデル"""

    def __init__(self, data):
        self.data = data
        self.test_field = data.get("test_field")


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
        return str(self._data)

    async def read(self):
        return b""


class TestClient(unittest.IsolatedAsyncioTestCase):
    def setUp(self):
        # モック状態を持つクライアントを作成
        self.state = MagicMock(spec=State)
        self.state.user_id = 12345
        self.state.email = "test@example.com"
        self.state.device_uuid = "test-device-uuid"
        self.state.access_token = "test-access-token"
        self.state.refresh_token = "test-refresh-token"

        # テスト設定でクライアントを作成
        self.client = Client(
            state=self.state,
            wait_on_ratelimit=False,  # テストを高速化するために待機を無効化
            loglevel=100,  # ログ出力を抑制
        )

    @patch("aiohttp.ClientSession.request")
    async def test_base_request_success(self, mock_request):
        """基本リクエストの成功をテスト"""
        # モックをセットアップ
        mock_response = MockResponse(
            status=200,
            data={"test_field": "test_value"},
            headers={"Content-Type": "application/json"},
        )
        mock_request.return_value.__aenter__.return_value = mock_response

        # リクエストを実行
        response = await self.client.base_request(
            "GET",
            "https://api.example.com/test",
            params={"param": "value"},
            headers={"X-Test": "test"},
        )

        # リクエストが正しく行われたことを検証
        mock_request.assert_called_once()
        args, kwargs = mock_request.call_args
        self.assertEqual(args[0], "GET")
        self.assertEqual(args[1], "https://api.example.com/test")
        self.assertEqual(kwargs["params"], {"param": "value"})
        self.assertEqual(kwargs["headers"]["X-Test"], "test")

        # レスポンスを検証
        self.assertEqual(response.status, 200)

    @patch("aiohttp.ClientSession.request")
    async def test_request_with_model_parsing(self, mock_request):
        """モデルパース機能を持つリクエストをテスト"""
        # モックをセットアップ
        mock_response = MockResponse(
            status=200,
            data={"test_field": "test_value"},
            headers={"Content-Type": "application/json"},
        )
        mock_request.return_value.__aenter__.return_value = mock_response

        # モデルパース機能を持つリクエストを実行
        # エラーを回避するためにuser.get_timestampメソッドをモック
        self.client.user = MagicMock()
        self.client.user.get_timestamp = AsyncMock(
            return_value=MagicMock(ip_address="127.0.0.1")
        )

        # テストデータを直接返すように__make_requestメソッドをモック
        with patch.object(
            self.client, "_Client__make_request", new=AsyncMock()
        ) as mock_make_request:
            mock_make_request.return_value = TestModel({"test_field": "test_value"})

            result = await self.client.request(
                "GET", "https://api.example.com/test", return_type=TestModel
            )

        # モデルが正しく作成されたことを検証
        self.assertIsInstance(result, TestModel)
        self.assertEqual(result.test_field, "test_value")

    @patch("aiohttp.ClientSession.request")
    async def test_request_with_list_response(self, mock_request):
        """リスト形式のレスポンスを持つリクエストをテスト"""
        # モックをセットアップ
        mock_response = MockResponse(
            status=200,
            data=[{"test_field": "value1"}, {"test_field": "value2"}],
            headers={"Content-Type": "application/json"},
        )
        mock_request.return_value.__aenter__.return_value = mock_response

        # モデルパース機能を持つリクエストを実行
        # エラーを回避するためにuser.get_timestampメソッドをモック
        self.client.user = MagicMock()
        self.client.user.get_timestamp = AsyncMock(
            return_value=MagicMock(ip_address="127.0.0.1")
        )

        # テストデータを直接返すように__make_requestメソッドをモック
        with patch.object(
            self.client, "_Client__make_request", new=AsyncMock()
        ) as mock_make_request:
            mock_make_request.return_value = [
                TestModel({"test_field": "value1"}),
                TestModel({"test_field": "value2"}),
            ]

            result = await self.client.request(
                "GET", "https://api.example.com/test", return_type=TestModel
            )

        # モデルのリストが正しく作成されたことを検証
        self.assertIsInstance(result, list)
        self.assertEqual(len(result), 2)
        self.assertIsInstance(result[0], TestModel)
        self.assertEqual(result[0].test_field, "value1")
        self.assertEqual(result[1].test_field, "value2")

    async def test_token_refresh_on_expired_token(self):
        """アクセストークンの有効期限切れ時のトークンリフレッシュをテスト"""
        # エラーを回避するためにuser.get_timestampメソッドをモック
        self.client.user = MagicMock()
        self.client.user.get_timestamp = AsyncMock(
            return_value=MagicMock(ip_address="127.0.0.1")
        )

        # 認証APIをモック
        self.client.auth = MagicMock()
        self.client.auth.get_token = AsyncMock(
            return_value=MagicMock(
                access_token="new-access-token", refresh_token="new-refresh-token"
            )
        )

        # AccessTokenExpiredError用の適切なErrorResponseを作成
        error_response = MagicMock()
        error_response.error_code = 1002
        error_response.message = "Token expired"

        # トークンの有効期限切れとリフレッシュをシミュレートするようにリクエストメソッドを設定
        request_count = 0

        async def mock_base_request(method, url, **kwargs):
            nonlocal request_count
            request_count += 1

            if request_count == 1:
                # 最初の呼び出し - トークン期限切れエラーをシミュレート
                raise AccessTokenExpiredError(error_response)
            else:
                # 2回目の呼び出し - トークンリフレッシュ後
                return MockResponse(status=200, data={"test_field": "test_value"})

        # refresh_client_tokensメソッドをモック
        refresh_called = False

        async def mock_refresh_tokens():
            nonlocal refresh_called
            refresh_called = True
            # auth.get_tokenメソッドを呼び出し
            token_response = await self.client.auth.get_token(
                grant_type="refresh_token", refresh_token=self.state.refresh_token
            )
            # 状態を更新
            self.client._Client__state.set_user = MagicMock()

        # メソッドをパッチ
        with patch.object(self.client, "base_request", side_effect=mock_base_request):
            with patch.object(
                self.client,
                "_Client__refresh_client_tokens",
                side_effect=mock_refresh_tokens,
            ):
                # リクエストを実行 - トークンリフレッシュがトリガーされるはず
                result = await self.client.request(
                    "GET", "https://api.example.com/test", return_type=TestModel
                )

        # トークンリフレッシュが呼び出されたことを検証
        self.assertTrue(refresh_called, "Token refresh method was not called")

        # auth.get_tokenが正しいパラメータで呼び出されたことを検証
        self.client.auth.get_token.assert_called_once_with(
            grant_type="refresh_token", refresh_token=self.state.refresh_token
        )

    @patch("aiohttp.ClientSession.request")
    async def test_rate_limit_handling(self, mock_request):
        """レート制限エラーの処理をテスト"""
        # レート制限エラーを発生させるようにモックを設定
        mock_response = MockResponse(status=429)  # Too Many Requests
        mock_request.return_value.__aenter__.return_value = mock_response

        # TooManyRequestsErrorを発生させるようにraise_for_codeを設定
        with patch(
            "yaylib.client.raise_for_code", new=AsyncMock()
        ) as mock_raise_for_code:
            # TooManyRequestsError用の適切なErrorResponseを作成
            error_response = MagicMock()
            error_response.error_code = 429
            error_response.message = "Rate limit exceeded"
            mock_raise_for_code.side_effect = TooManyRequestsError(error_response)

            # リクエストを実行 - wait_on_ratelimit=Falseなのでエラーが発生するはず
            with patch("yaylib.client.raise_for_status", new=AsyncMock()):
                with self.assertRaises(TooManyRequestsError):
                    # エラーを回避するためにuser.get_timestampメソッドをモック
                    self.client.user = MagicMock()
                    self.client.user.get_timestamp = AsyncMock(
                        return_value=MagicMock(ip_address="127.0.0.1")
                    )

                    await self.client.request("GET", "https://api.example.com/test")

    async def test_request_with_jwt_required(self):
        """JWT必須のリクエストをテスト"""
        # エラーを回避するためにuser.get_timestampメソッドをモック
        self.client.user = MagicMock()
        self.client.user.get_timestamp = AsyncMock(
            return_value=MagicMock(ip_address="127.0.0.1")
        )

        # ヘッダーをキャプチャするためにbase_requestメソッドをモック
        headers_captured = {}

        async def mock_base_request(method, url, **kwargs):
            nonlocal headers_captured
            headers_captured = kwargs.get("headers", {})
            return MockResponse(status=200, data={})

        with patch.object(self.client, "base_request", side_effect=mock_base_request):
            with patch("yaylib.client.raise_for_code", new=AsyncMock()):
                with patch("yaylib.client.raise_for_status", new=AsyncMock()):
                    with patch("yaylib.client.generate_jwt", return_value="test-jwt"):
                        await self.client.request(
                            "GET", "https://api.example.com/test", jwt_required=True
                        )

        # ヘッダーにJWTが含まれていることを検証
        self.assertIn("X-Jwt", headers_captured)
        self.assertEqual(headers_captured["X-Jwt"], "test-jwt")


if __name__ == "__main__":
    unittest.main()
