import unittest

from . import test_storage


def suite():
    suite = unittest.TestSuite()
    suite.addTest(test_storage.TestStorage)
    return suite


if __name__ == "__main__":
    runner = unittest.TextTestRunner()
    runner.run(suite())
