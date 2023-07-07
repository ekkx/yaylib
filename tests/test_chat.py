import unittest

from config import (
    tape,
    user_id,
    opponent_id,
    YaylibTestCase,
)


class YaylibChatTests(YaylibTestCase):

    path = "./chat/"


if __name__ == '__main__':
    unittest.main()
