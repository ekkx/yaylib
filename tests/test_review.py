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


if __name__ == '__main__':
    unittest.main()
