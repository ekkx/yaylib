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
from yaylib.models.group_category import GroupCategory
from yaylib.models.model_post import ModelPost
from typing import Optional, Set
from typing_extensions import Self

class TimelineItem(BaseModel):
    """
    TimelineItem
    """ # noqa: E501
    categories: Optional[List[GroupCategory]] = None
    is_replied_by_next_post: Optional[StrictBool] = None
    is_replying_to_prev_post: Optional[StrictBool] = None
    post: Optional[ModelPost] = None
    should_show_conversation_line: Optional[StrictBool] = None
    should_show_mention_box: Optional[StrictBool] = None
    should_show_repost_button: Optional[StrictBool] = None
    should_show_translate_button: Optional[StrictBool] = None
    should_show_view_reposts: Optional[StrictBool] = None
    translated_text: Optional[StrictStr] = None
    __properties: ClassVar[List[str]] = ["categories", "is_replied_by_next_post", "is_replying_to_prev_post", "post", "should_show_conversation_line", "should_show_mention_box", "should_show_repost_button", "should_show_translate_button", "should_show_view_reposts", "translated_text"]

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
        """Create an instance of TimelineItem from a JSON string"""
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
        # override the default output from pydantic by calling `to_dict()` of each item in categories (list)
        _items = []
        if self.categories:
            for _item_categories in self.categories:
                if _item_categories:
                    _items.append(_item_categories.to_dict())
            _dict['categories'] = _items
        # override the default output from pydantic by calling `to_dict()` of post
        if self.post:
            _dict['post'] = self.post.to_dict()
        # set to None if categories (nullable) is None
        # and model_fields_set contains the field
        if self.categories is None and "categories" in self.model_fields_set:
            _dict['categories'] = None

        # set to None if is_replied_by_next_post (nullable) is None
        # and model_fields_set contains the field
        if self.is_replied_by_next_post is None and "is_replied_by_next_post" in self.model_fields_set:
            _dict['is_replied_by_next_post'] = None

        # set to None if is_replying_to_prev_post (nullable) is None
        # and model_fields_set contains the field
        if self.is_replying_to_prev_post is None and "is_replying_to_prev_post" in self.model_fields_set:
            _dict['is_replying_to_prev_post'] = None

        # set to None if post (nullable) is None
        # and model_fields_set contains the field
        if self.post is None and "post" in self.model_fields_set:
            _dict['post'] = None

        # set to None if should_show_conversation_line (nullable) is None
        # and model_fields_set contains the field
        if self.should_show_conversation_line is None and "should_show_conversation_line" in self.model_fields_set:
            _dict['should_show_conversation_line'] = None

        # set to None if should_show_mention_box (nullable) is None
        # and model_fields_set contains the field
        if self.should_show_mention_box is None and "should_show_mention_box" in self.model_fields_set:
            _dict['should_show_mention_box'] = None

        # set to None if should_show_repost_button (nullable) is None
        # and model_fields_set contains the field
        if self.should_show_repost_button is None and "should_show_repost_button" in self.model_fields_set:
            _dict['should_show_repost_button'] = None

        # set to None if should_show_translate_button (nullable) is None
        # and model_fields_set contains the field
        if self.should_show_translate_button is None and "should_show_translate_button" in self.model_fields_set:
            _dict['should_show_translate_button'] = None

        # set to None if should_show_view_reposts (nullable) is None
        # and model_fields_set contains the field
        if self.should_show_view_reposts is None and "should_show_view_reposts" in self.model_fields_set:
            _dict['should_show_view_reposts'] = None

        # set to None if translated_text (nullable) is None
        # and model_fields_set contains the field
        if self.translated_text is None and "translated_text" in self.model_fields_set:
            _dict['translated_text'] = None

        return _dict

    @classmethod
    def from_dict(cls, obj: Optional[Dict[str, Any]]) -> Optional[Self]:
        """Create an instance of TimelineItem from a dict"""
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return cls.model_validate(obj)

        _obj = cls.model_validate({
            "categories": [GroupCategory.from_dict(_item) for _item in obj["categories"]] if obj.get("categories") is not None else None,
            "is_replied_by_next_post": obj.get("is_replied_by_next_post"),
            "is_replying_to_prev_post": obj.get("is_replying_to_prev_post"),
            "post": ModelPost.from_dict(obj["post"]) if obj.get("post") is not None else None,
            "should_show_conversation_line": obj.get("should_show_conversation_line"),
            "should_show_mention_box": obj.get("should_show_mention_box"),
            "should_show_repost_button": obj.get("should_show_repost_button"),
            "should_show_translate_button": obj.get("should_show_translate_button"),
            "should_show_view_reposts": obj.get("should_show_view_reposts"),
            "translated_text": obj.get("translated_text")
        })
        return _obj


