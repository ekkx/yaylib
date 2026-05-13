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

from pydantic import BaseModel, ConfigDict, StrictBool
from typing import Any, ClassVar, Dict, List, Optional
from typing import Optional, Set
from typing_extensions import Self

class TimelineSettings(BaseModel):
    """
    TimelineSettings
    """ # noqa: E501
    faves_filter: Optional[StrictBool] = None
    hide_hot_post: Optional[StrictBool] = None
    hide_reply_following_timeline: Optional[StrictBool] = None
    hide_reply_public_timeline: Optional[StrictBool] = None
    __properties: ClassVar[List[str]] = ["faves_filter", "hide_hot_post", "hide_reply_following_timeline", "hide_reply_public_timeline"]

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
        """Create an instance of TimelineSettings from a JSON string"""
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
        # set to None if faves_filter (nullable) is None
        # and model_fields_set contains the field
        if self.faves_filter is None and "faves_filter" in self.model_fields_set:
            _dict['faves_filter'] = None

        # set to None if hide_hot_post (nullable) is None
        # and model_fields_set contains the field
        if self.hide_hot_post is None and "hide_hot_post" in self.model_fields_set:
            _dict['hide_hot_post'] = None

        # set to None if hide_reply_following_timeline (nullable) is None
        # and model_fields_set contains the field
        if self.hide_reply_following_timeline is None and "hide_reply_following_timeline" in self.model_fields_set:
            _dict['hide_reply_following_timeline'] = None

        # set to None if hide_reply_public_timeline (nullable) is None
        # and model_fields_set contains the field
        if self.hide_reply_public_timeline is None and "hide_reply_public_timeline" in self.model_fields_set:
            _dict['hide_reply_public_timeline'] = None

        return _dict

    @classmethod
    def from_dict(cls, obj: Optional[Dict[str, Any]]) -> Optional[Self]:
        """Create an instance of TimelineSettings from a dict"""
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return cls.model_validate(obj)

        _obj = cls.model_validate({
            "faves_filter": obj.get("faves_filter"),
            "hide_hot_post": obj.get("hide_hot_post"),
            "hide_reply_following_timeline": obj.get("hide_reply_following_timeline"),
            "hide_reply_public_timeline": obj.get("hide_reply_public_timeline")
        })
        return _obj


