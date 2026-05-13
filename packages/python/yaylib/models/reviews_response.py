# coding: utf-8

"""
    Yay! API

    No description provided

    Schema version: 4.26.1
    Code generated; DO NOT EDIT

    Do not edit the class manually.
"""  # noqa: E501


from __future__ import annotations
import pprint
import re  # noqa: F401
import json

from pydantic import BaseModel, ConfigDict
from typing import Any, ClassVar, Dict, List, Optional
from yaylib.models.review import Review
from typing import Optional, Set
from typing_extensions import Self

class ReviewsResponse(BaseModel):
    """
    ReviewsResponse
    """ # noqa: E501
    pinned_reviews: Optional[List[Review]] = None
    reviews: Optional[List[Review]] = None
    __properties: ClassVar[List[str]] = ["pinned_reviews", "reviews"]

    model_config = ConfigDict(
        populate_by_name=True,
        validate_assignment=True,
        protected_namespaces=(),
    )


    def to_str(self) -> str:
        """Returns the string representation of the model using alias"""
        return pprint.pformat(self.model_dump(by_alias=True))

    def to_json(self) -> str:
        """Returns the JSON representation of the model using alias"""
        # TODO: pydantic v2: use .model_dump_json(by_alias=True, exclude_unset=True) instead
        return json.dumps(self.to_dict())

    @classmethod
    def from_json(cls, json_str: str) -> Optional[Self]:
        """Create an instance of ReviewsResponse from a JSON string"""
        return cls.from_dict(json.loads(json_str))

    def to_dict(self) -> Dict[str, Any]:
        """Return the dictionary representation of the model using alias.

        This has the following differences from calling pydantic's
        `self.model_dump(by_alias=True)`:

        * `None` is only added to the output dict for nullable fields that
          were set at model initialization. Other fields with value `None`
          are ignored.
        """
        excluded_fields: Set[str] = set([
        ])

        _dict = self.model_dump(
            by_alias=True,
            exclude=excluded_fields,
            exclude_none=True,
        )
        # override the default output from pydantic by calling `to_dict()` of each item in pinned_reviews (list)
        _items = []
        if self.pinned_reviews:
            for _item_pinned_reviews in self.pinned_reviews:
                if _item_pinned_reviews:
                    _items.append(_item_pinned_reviews.to_dict())
            _dict['pinned_reviews'] = _items
        # override the default output from pydantic by calling `to_dict()` of each item in reviews (list)
        _items = []
        if self.reviews:
            for _item_reviews in self.reviews:
                if _item_reviews:
                    _items.append(_item_reviews.to_dict())
            _dict['reviews'] = _items
        # set to None if pinned_reviews (nullable) is None
        # and model_fields_set contains the field
        if self.pinned_reviews is None and "pinned_reviews" in self.model_fields_set:
            _dict['pinned_reviews'] = None

        # set to None if reviews (nullable) is None
        # and model_fields_set contains the field
        if self.reviews is None and "reviews" in self.model_fields_set:
            _dict['reviews'] = None

        return _dict

    @classmethod
    def from_dict(cls, obj: Optional[Dict[str, Any]]) -> Optional[Self]:
        """Create an instance of ReviewsResponse from a dict"""
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return cls.model_validate(obj)

        _obj = cls.model_validate({
            "pinned_reviews": [Review.from_dict(_item) for _item in obj["pinned_reviews"]] if obj.get("pinned_reviews") is not None else None,
            "reviews": [Review.from_dict(_item) for _item in obj["reviews"]] if obj.get("reviews") is not None else None
        })
        return _obj


