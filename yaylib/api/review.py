from datetime import datetime
from typing import Dict, List

from ..config import *
from ..errors import *
from ..models import *
from ..responses import *
from ..utils import *


def create_review(self, user_id: int, comment: str):
    pass


def create_reviews(self, user_ids: List[int], comment: str):
    pass


def delete_reviews(self, review_ids: List[int]):
    pass


def get_my_reviews(self, from_id: int = None) -> ReviewsResponse:
    pass


def get_reviews(self, user_id: int, from_id: int = None) -> ReviewsResponse:
    pass


def pin_review(self, review_id: int):
    pass


def unpin_review(self, review_id: int):
    pass
