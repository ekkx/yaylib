import unittest
import yaylib

from config import test_email, test_password


class BaseClientTests(unittest.TestCase):
    def setUp(self):
        self.base = yaylib.BaseClient()
        self.base.login(test_email, test_password)

    def test_make_request(self):
        pass
