"""
MIT License

Copyright (c) 2023 ekkx

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
"""

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
            text,
            font_size,
            color,
        ).id
        self.api.get_post(post_id=post_id)

        # create reply
        conversation_id = self.api.create_post(
            text,
            font_size,
            color,
            in_reply_to=post_id,
        ).conversation_id
        self.api.get_conversation(conversation_id, number=number)
        self.api.get_conversation_root_posts(conversation_id)
        self.api.create_repost(
            post_id,
            text,
            font_size,
            color,
            in_reply_to=post_id,
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
            group_id=group_id, keyword=keyword, number=number
        )
        self.api.get_following_timeline(number=number)
        self.api.get_following_call_timeline(number=number)
        self.api.get_recommended_posts(
            number=number, experiment_num=experiment_num, variant_num=variant_num
        )
        self.api.get_user_timeline(user_id=user_id, number=number)
        self.api.get_my_posts(number=number)

    @tape.use_cassette(path + "bookmarks.yaml")
    def test_bookmarks(self):
        post_id = 371574235

        self.api.add_bookmark(user_id, post_id)
        self.api.get_bookmark(user_id)
        self.api.remove_bookmark(user_id, post_id)


if __name__ == "__main__":
    unittest.main()
