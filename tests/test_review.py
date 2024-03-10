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


class YaylibReviewTests(YaylibTestCase):
    path = "./review/"

    @tape.use_cassette(path + "reviews.yaml")
    def test_bookmarks(self):
        comment = "testing"
        number = 1

        self.api.create_review(user_id=opponent_id, comment=comment)

        # TODO: INVALID SIGNED INFO

        self.api.create_reviews(user_ids=[opponent_id], comment=comment)

        reviews = self.api.get_reviews(user_id, number=number).reviews
        if reviews:
            review_id = reviews[0].id
            self.api.pin_review(review_id)
            self.api.unpin_review(review_id)

        reviews = self.api.get_my_reviews(number=number).reviews
        review_id = reviews[0].id
        self.api.delete_reviews([review_id])


if __name__ == "__main__":
    unittest.main()
