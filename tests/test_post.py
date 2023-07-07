import unittest

from config import (
    tape,
    user_id,
    opponent_id,
    YaylibTestCase,
)


class YaylibPostTests(YaylibTestCase):

    path = "./post/"

    @tape.use_cassette(path + "posts.yaml")
    def test_posts(self):
        text = "testing"
        font_size = 2
        color = 3
        number = 1

        # create post
        post_id = self.api.create_post(
            text, font_size, color,
        ).id
        self.api.get_post(post_id=post_id)

        # create reply
        conversation_id = self.api.create_post(
            text, font_size, color, in_reply_to=post_id,
        ).conversation_id
        self.api.get_conversation(conversation_id, number=number)
        self.api.get_conversation_root_posts(conversation_id)
        self.api.create_repost(
            post_id, text, font_size, color, in_reply_to=post_id,
        )
        self.api.get_post_reposts(post_id=post_id, number=number)

    @tape.use_cassette(path + "timeline.yaml")
    def test_timeline(self):
        number = 1
        user_id = 93
        group_id = 179925
        keyword = "こんにちは"
        hashtag = "いいねでレター"
        experiment_num = 1
        variant_num = 1

        self.api.get_timeline(number=number)
        self.api.get_timeline(number=number, noreply_mode=True)
        self.api.get_timeline_by_keyword(keyword=keyword, number=number)
        self.api.get_timeline_by_hashtag(hashtag=hashtag, number=number)
        self.api.get_timeline_calls(number=number)
        self.api.get_timeline_calls(number=number, group_id=group_id)
        self.api.get_group_timeline(group_id=group_id, number=number)
        self.api.get_group_timeline_by_keyword(
            group_id=group_id,
            keyword=keyword,
            number=number
        )
        self.api.get_following_timeline(number=number)
        self.api.get_following_call_timeline(number=number)
        self.api.get_recommended_posts(
            number=number,
            experiment_num=experiment_num,
            variant_num=variant_num
        )
        self.api.get_user_timeline(user_id=user_id, number=number)
        self.api.get_my_posts(number=number)

    @tape.use_cassette(path + "bookmarks.yaml")
    def test_bookmarks(self):
        post_id = 371574235

        self.api.add_bookmark(user_id, post_id)
        self.api.get_bookmark(user_id)
        self.api.remove_bookmark(user_id, post_id)


if __name__ == '__main__':
    unittest.main()
