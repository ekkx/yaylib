import unittest

from config import (
    tape,
    user_id,
    opponent_id,
    YaylibTestCase,
)


class YaylibCassandraTests(YaylibTestCase):
    path = "./cassandra/"

    @tape.use_cassette(path + "test_get_user_activities.yaml")
    def test_get_user_activities(self):
        number = 1
        self.api.get_user_activities(number=number)

    @tape.use_cassette(path + "test_get_user_merged_activities.yaml")
    def test_get_user_merged_activities(self):
        number = 1
        self.api.get_user_merged_activities(number=number)

    # @tape.use_cassette(path + "test_received_notification.yaml")
    # def test_received_notification(self):
    #     pid = ""
    #     type = ""
    #     self.api.received_notification(pid, type)


if __name__ == "__main__":
    unittest.main()
