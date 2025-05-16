import unittest
from unittest.mock import MagicMock, patch

from yaylib.client import HeaderManager
from yaylib.device import Device
from yaylib.state import State


class TestHeaderManager(unittest.TestCase):
    def setUp(self):
        # モックステートを作成
        self.state = MagicMock(spec=State)
        self.state.device_uuid = "test-device-uuid"
        self.state.access_token = "test-access-token"

        # モックデバイスを作成
        self.device = MagicMock(spec=Device)
        self.device.get_user_agent.return_value = (
            "android 11 (3.5x 1440x2960 Galaxy S9)"
        )
        self.device.get_device_info.return_value = (
            "yay 1.0.0 android 11 (3.5x 1440x2960 Galaxy S9)"
        )

        # ヘッダーマネージャーを作成
        self.header_manager = HeaderManager(self.device, self.state, locale="en")

    def test_properties(self):
        """ヘッダーマネージャーのプロパティをテスト"""
        self.assertEqual(self.header_manager.locale, "en")
        self.assertEqual(
            self.header_manager.user_agent, "android 11 (3.5x 1440x2960 Galaxy S9)"
        )
        self.assertEqual(
            self.header_manager.device_info,
            "yay 1.0.0 android 11 (3.5x 1440x2960 Galaxy S9)",
        )
        self.assertEqual(self.header_manager.client_ip, "")
        self.assertEqual(self.header_manager.connection_speed, "0 kbps")
        self.assertEqual(self.header_manager.connection_type, "wifi")
        self.assertEqual(
            self.header_manager.content_type, "application/json;charset=UTF-8"
        )

    def test_client_ip_setter(self):
        """クライアントIPの設定をテスト"""
        self.header_manager.client_ip = "192.168.1.1"
        self.assertEqual(self.header_manager.client_ip, "192.168.1.1")

    @patch("yaylib.client.datetime")
    def test_generate_headers_without_jwt(self, mock_datetime):
        """JWT無しでのヘッダー生成をテスト"""
        # datetime.now().timestamp()をモック
        mock_now = MagicMock()
        mock_now.timestamp.return_value = 1620000000
        mock_datetime.now.return_value = mock_now

        # ヘッダーを生成
        headers = self.header_manager.generate(jwt_required=False)

        # ヘッダーを検証
        self.assertEqual(headers["Host"], "api.yay.space")
        self.assertEqual(headers["User-Agent"], "android 11 (3.5x 1440x2960 Galaxy S9)")
        self.assertEqual(headers["X-Timestamp"], "1620000000")
        self.assertEqual(headers["X-Device-UUID"], "test-device-uuid")
        self.assertEqual(headers["X-Connection-Type"], "wifi")
        self.assertEqual(headers["X-Connection-Speed"], "0 kbps")
        self.assertEqual(headers["Accept-Language"], "en")
        self.assertEqual(headers["Content-Type"], "application/json;charset=UTF-8")
        self.assertEqual(headers["Authorization"], "Bearer test-access-token")

        # JWTが含まれていないことを検証
        self.assertNotIn("X-Jwt", headers)

    @patch("yaylib.client.datetime")
    @patch("yaylib.client.generate_jwt")
    def test_generate_headers_with_jwt(self, mock_generate_jwt, mock_datetime):
        """JWTありでのヘッダー生成をテスト"""
        # datetime.now().timestamp()をモック
        mock_now = MagicMock()
        mock_now.timestamp.return_value = 1620000000
        mock_datetime.now.return_value = mock_now

        # generate_jwtをモック
        mock_generate_jwt.return_value = "test-jwt-token"

        # ヘッダーを生成
        headers = self.header_manager.generate(jwt_required=True)

        # ヘッダーを検証
        self.assertEqual(headers["Host"], "api.yay.space")
        self.assertEqual(headers["User-Agent"], "android 11 (3.5x 1440x2960 Galaxy S9)")
        self.assertEqual(headers["X-Timestamp"], "1620000000")
        self.assertEqual(headers["X-Device-UUID"], "test-device-uuid")
        self.assertEqual(headers["X-Connection-Type"], "wifi")
        self.assertEqual(headers["X-Connection-Speed"], "0 kbps")
        self.assertEqual(headers["Accept-Language"], "en")
        self.assertEqual(headers["Content-Type"], "application/json;charset=UTF-8")
        self.assertEqual(headers["Authorization"], "Bearer test-access-token")

        # JWTが含まれていることを検証
        self.assertEqual(headers["X-Jwt"], "test-jwt-token")

    @patch("yaylib.client.datetime")
    def test_generate_headers_with_client_ip(self, mock_datetime):
        """クライアントIPを含むヘッダー生成をテスト"""
        # datetime.now().timestamp()をモック
        mock_now = MagicMock()
        mock_now.timestamp.return_value = 1620000000
        mock_datetime.now.return_value = mock_now

        # クライアントIPを設定
        self.header_manager.client_ip = "192.168.1.1"

        # ヘッダーを生成
        headers = self.header_manager.generate(jwt_required=False)

        # クライアントIPが含まれていることを検証
        self.assertEqual(headers["X-Client-IP"], "192.168.1.1")

    @patch("yaylib.client.datetime")
    def test_generate_headers_without_access_token(self, mock_datetime):
        """アクセストークン無しでのヘッダー生成をテスト"""
        # datetime.now().timestamp()をモック
        mock_now = MagicMock()
        mock_now.timestamp.return_value = 1620000000
        mock_datetime.now.return_value = mock_now

        # アクセストークンを空文字列に設定
        self.state.access_token = ""

        # ヘッダーを生成
        headers = self.header_manager.generate(jwt_required=False)

        # Authorizationヘッダーが含まれていないことを検証
        self.assertNotIn("Authorization", headers)


if __name__ == "__main__":
    unittest.main()
