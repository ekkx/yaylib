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

from pydantic import BaseModel, ConfigDict, StrictBool, StrictStr
from typing import Any, ClassVar, Dict, List, Optional
from yaylib.models.post import Post
from typing import Optional, Set
from typing_extensions import Self

class PostsResponse(BaseModel):
    """
    PostsResponse
    """ # noqa: E501
    has_more_hot_posts: Optional[StrictBool] = None
    next_page_value: Optional[StrictStr] = None
    pinned_posts: Optional[List[Post]] = None
    posts: Optional[List[Post]] = None
    __properties: ClassVar[List[str]] = ["has_more_hot_posts", "next_page_value", "pinned_posts", "posts"]

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
        """Create an instance of PostsResponse from a JSON string"""
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
        # override the default output from pydantic by calling `to_dict()` of each item in pinned_posts (list)
        _items = []
        if self.pinned_posts:
            for _item_pinned_posts in self.pinned_posts:
                if _item_pinned_posts:
                    _items.append(_item_pinned_posts.to_dict())
            _dict['pinned_posts'] = _items
        # override the default output from pydantic by calling `to_dict()` of each item in posts (list)
        _items = []
        if self.posts:
            for _item_posts in self.posts:
                if _item_posts:
                    _items.append(_item_posts.to_dict())
            _dict['posts'] = _items
        # set to None if has_more_hot_posts (nullable) is None
        # and model_fields_set contains the field
        if self.has_more_hot_posts is None and "has_more_hot_posts" in self.model_fields_set:
            _dict['has_more_hot_posts'] = None

        # set to None if next_page_value (nullable) is None
        # and model_fields_set contains the field
        if self.next_page_value is None and "next_page_value" in self.model_fields_set:
            _dict['next_page_value'] = None

        # set to None if pinned_posts (nullable) is None
        # and model_fields_set contains the field
        if self.pinned_posts is None and "pinned_posts" in self.model_fields_set:
            _dict['pinned_posts'] = None

        # set to None if posts (nullable) is None
        # and model_fields_set contains the field
        if self.posts is None and "posts" in self.model_fields_set:
            _dict['posts'] = None

        return _dict

    @classmethod
    def from_dict(cls, obj: Optional[Dict[str, Any]]) -> Optional[Self]:
        """Create an instance of PostsResponse from a dict"""
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return cls.model_validate(obj)

        _obj = cls.model_validate({
            "has_more_hot_posts": obj.get("has_more_hot_posts"),
            "next_page_value": obj.get("next_page_value"),
            "pinned_posts": [Post.from_dict(_item) for _item in obj["pinned_posts"]] if obj.get("pinned_posts") is not None else None,
            "posts": [Post.from_dict(_item) for _item in obj["posts"]] if obj.get("posts") is not None else None
        })
        return _obj


