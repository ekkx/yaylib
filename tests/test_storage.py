import os
import unittest

from yaylib.state import Storage, User


db_path = os.path.dirname(os.path.dirname(__file__)) + "/.config/test.db"


test_user = User(
    user_id=2342713,
    email="test@email.com",
    device_uuid="0d060774-2d1b-4ce1-avb3-e5a2d0523124",
    access_token="983ea785322e8d3efe105a40b1aed4c9721509a11f3c32lv63ad351c7cqe1a9c",
    refresh_token="55ee3b11c3b8022890vab1cde720282c1c8dc6cef0baceeeibe26wc40724e018",
)


class TestStorage(unittest.TestCase):
    def setUp(self):
        self.clean()
        self.storage = Storage(db_path)

    def tearDown(self):
        self.clean()

    @staticmethod
    def clean():
        if os.path.isfile(db_path):
            os.remove(db_path)

    def test_get_user(self):
        result = self.storage.create_user(test_user)
        self.assertTrue(result)

        user_by_id = self.storage.get_user(user_id=test_user.user_id)

        self.assertIsNotNone(user_by_id)
        self.assertEqual(user_by_id.user_id, test_user.user_id)
        self.assertEqual(user_by_id.email, test_user.email)
        self.assertEqual(user_by_id.device_uuid, test_user.device_uuid)
        self.assertEqual(user_by_id.access_token, test_user.access_token)
        self.assertEqual(user_by_id.refresh_token, test_user.refresh_token)

        user_by_email = self.storage.get_user(email=test_user.email)

        self.assertIsNotNone(user_by_email)
        self.assertEqual(user_by_email.user_id, test_user.user_id)
        self.assertEqual(user_by_email.email, test_user.email)
        self.assertEqual(user_by_email.device_uuid, test_user.device_uuid)
        self.assertEqual(user_by_email.access_token, test_user.access_token)
        self.assertEqual(user_by_email.refresh_token, test_user.refresh_token)

        no_user_by_id = self.storage.get_user(user_id=0)
        self.assertIsNone(no_user_by_id)

        no_user_by_email = self.storage.get_user(email="email_that_does_not@exist.com")
        self.assertIsNone(no_user_by_email)

        result = self.storage.delete_user(test_user.user_id)
        self.assertTrue(result)

    def test_create_user(self):
        result = self.storage.create_user(test_user)
        self.assertTrue(result)

        user = self.storage.get_user(test_user.user_id)

        self.assertIsNotNone(user)
        self.assertEqual(user.user_id, test_user.user_id)
        self.assertEqual(user.email, test_user.email)
        self.assertEqual(user.device_uuid, test_user.device_uuid)
        self.assertEqual(user.access_token, test_user.access_token)
        self.assertEqual(user.refresh_token, test_user.refresh_token)

        result = self.storage.delete_user(user.user_id)
        self.assertTrue(result)

    def test_update_user(self):
        result = self.storage.create_user(test_user)
        self.assertTrue(result)

        updated_user = User(
            user_id=test_user.user_id,
            email="updated@email.com",
            device_uuid="00000000-0000-0000-0000-00000000000",
            access_token=test_user.refresh_token,
            refresh_token=test_user.access_token,
        )

        result = self.storage.update_user(
            test_user.user_id,
            email=updated_user.email,
            device_uuid=updated_user.device_uuid,
            access_token=updated_user.access_token,
            refresh_token=updated_user.refresh_token,
        )
        self.assertTrue(result)

        user = self.storage.get_user(test_user.user_id)

        self.assertIsNotNone(user)
        self.assertEqual(user.user_id, updated_user.user_id)
        self.assertEqual(user.email, updated_user.email)
        self.assertEqual(user.device_uuid, updated_user.device_uuid)
        self.assertEqual(user.access_token, updated_user.access_token)
        self.assertEqual(user.refresh_token, updated_user.refresh_token)

        result = self.storage.delete_user(user.user_id)
        self.assertTrue(result)

    def test_delete_user(self):
        result = self.storage.create_user(test_user)
        self.assertTrue(result)

        result = self.storage.delete_user(test_user.user_id)
        self.assertTrue(result)

        user = self.storage.get_user(test_user.user_id)
        self.assertIsNone(user)
