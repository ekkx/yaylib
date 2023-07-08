import unittest

from config import (
    tape,
    user_id,
    opponent_id,
    YaylibTestCase,
)


class YaylibUserTests(YaylibTestCase):
    path = "./user/"


if __name__ == "__main__":
    unittest.main()
