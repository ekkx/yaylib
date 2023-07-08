import unittest

from config import (
    tape,
    user_id,
    opponent_id,
    YaylibTestCase,
)


class YaylibThreadTests(YaylibTestCase):
    path = "./thread/"


if __name__ == "__main__":
    unittest.main()
