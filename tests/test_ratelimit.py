import unittest
from unittest.mock import patch

from yaylib.client import RateLimit


class TestException(Exception):
    pass


class TestRateLimit(unittest.IsolatedAsyncioTestCase):
    def test_initial_retries_performed(self):
        ratelimit = RateLimit(wait_on_ratelimit=True, max_retries=3, retry_after=1)
        self.assertEqual(ratelimit.retries_performed, 0)

    def test_initial_max_retries(self):
        ratelimit = RateLimit(wait_on_ratelimit=True, max_retries=3, retry_after=1)
        self.assertEqual(ratelimit.max_retries, 3)

    def test_max_retries_reached(self):
        ratelimit = RateLimit(wait_on_ratelimit=True, max_retries=3, retry_after=1)
        ratelimit.retries_performed = 3
        self.assertTrue(ratelimit.max_retries_reached)

    def test_reset(self):
        ratelimit = RateLimit(wait_on_ratelimit=True, max_retries=3, retry_after=1)
        ratelimit.retries_performed = 2
        self.assertEqual(ratelimit.retries_performed, 2)
        ratelimit.reset()
        self.assertEqual(ratelimit.retries_performed, 0)

    @patch("asyncio.sleep", return_value=None)
    async def test_wait_on_ratelimit_enabled(self, mock_sleep):
        ratelimit = RateLimit(wait_on_ratelimit=True, max_retries=1, retry_after=1)

        await ratelimit.wait(TestException())
        self.assertEqual(ratelimit.retries_performed, 1)

        with self.assertRaises(TestException):
            await ratelimit.wait(TestException())

    @patch("asyncio.sleep", return_value=None)
    async def test_wait_on_ratelimit_disabled(self, mock_sleep):
        ratelimit = RateLimit(wait_on_ratelimit=False, max_retries=2, retry_after=1)

        self.assertTrue(ratelimit.max_retries_reached)

        with self.assertRaises(TestException):
            await ratelimit.wait(TestException())
