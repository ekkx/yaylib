import unittest

from config import (
    tape,
    user_id,
    opponent_id,
    YaylibTestCase,
)


class YaylibCallTests(YaylibTestCase):
    path = "./call/"


if __name__ == "__main__":
    unittest.main()
