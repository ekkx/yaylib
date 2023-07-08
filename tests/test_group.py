import unittest

from config import (
    tape,
    user_id,
    opponent_id,
    YaylibTestCase,
)


class YaylibGroupTests(YaylibTestCase):
    path = "./group/"


if __name__ == "__main__":
    unittest.main()
