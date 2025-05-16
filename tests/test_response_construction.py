import unittest
from unittest.mock import MagicMock

from yaylib.client import Client
from yaylib.models import Model
from yaylib.responses import Response


class TestModel(Model):
    """レスポンスパース用のテストモデル"""

    def __init__(self, data):
        self.data = data
        self.test_field = data.get("test_field")
        self.nested = data.get("nested")


class TestResponse(Response):
    """テストレスポンスクラス"""

    def __init__(self, data):
        self.data = data
        self.result = data.get("result")
        self.test_field = data.get("test_field")
        self.items = data.get("items", [])


class TestResponseConstruction(unittest.TestCase):
    def setUp(self):
        # モック依存関係を持つクライアントを作成
        self.client = MagicMock(spec=Client)

    def test_construct_response_with_none(self):
        """None入力でのレスポンス構築をテスト"""
        # プライベートメソッドなので直接呼び出す
        result = Client._Client__construct_response(self.client, None, TestModel)
        self.assertIsNone(result)

    def test_construct_response_with_dict(self):
        """辞書入力でのレスポンス構築をテスト"""
        test_data = {"test_field": "test_value", "nested": {"key": "value"}}

        # プライベートメソッドなので直接呼び出す
        result = Client._Client__construct_response(self.client, test_data, TestModel)

        # 結果を検証
        self.assertIsInstance(result, TestModel)
        self.assertEqual(result.test_field, "test_value")
        self.assertEqual(result.nested, {"key": "value"})

    def test_construct_response_with_list(self):
        """リスト入力でのレスポンス構築をテスト"""
        test_data = [
            {"test_field": "value1", "nested": {"key1": "value1"}},
            {"test_field": "value2", "nested": {"key2": "value2"}},
        ]

        # プライベートメソッドなので直接呼び出す
        result = Client._Client__construct_response(self.client, test_data, TestModel)

        # 結果を検証
        self.assertIsInstance(result, list)
        self.assertEqual(len(result), 2)

        # 最初の項目を確認
        self.assertIsInstance(result[0], TestModel)
        self.assertEqual(result[0].test_field, "value1")
        self.assertEqual(result[0].nested, {"key1": "value1"})

        # 2番目の項目を確認
        self.assertIsInstance(result[1], TestModel)
        self.assertEqual(result[1].test_field, "value2")
        self.assertEqual(result[1].nested, {"key2": "value2"})

    def test_construct_response_without_model(self):
        """モデルを指定せずにレスポンスを構築するテスト"""
        test_data = {"test_field": "test_value", "nested": {"key": "value"}}

        # プライベートメソッドなので直接呼び出す
        result = Client._Client__construct_response(self.client, test_data, None)

        # 結果がそのまま返されることを検証
        self.assertEqual(result, test_data)

    def test_construct_response_with_empty_list(self):
        """空リストでのレスポンス構築をテスト"""
        test_data = []

        # プライベートメソッドなので直接呼び出す
        result = Client._Client__construct_response(self.client, test_data, TestModel)

        # 結果が空リストであることを検証
        self.assertEqual(result, [])

    def test_construct_response_with_complex_structure(self):
        """より複雑な構造でのレスポンス構築をテスト"""
        test_data = {
            "test_field": "parent_value",
            "items": [{"test_field": "child1"}, {"test_field": "child2"}],
        }

        # プライベートメソッドなので直接呼び出す
        result = Client._Client__construct_response(
            self.client, test_data, TestResponse
        )

        # 結果を検証
        self.assertIsInstance(result, TestResponse)
        self.assertEqual(result.test_field, "parent_value")
        self.assertEqual(len(result.items), 2)
        self.assertEqual(result.items[0]["test_field"], "child1")
        self.assertEqual(result.items[1]["test_field"], "child2")


if __name__ == "__main__":
    unittest.main()
