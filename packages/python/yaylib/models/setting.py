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

class Setting(BaseModel):
    """
    Setting
    """ # noqa: E501
    notification_group_join: Optional[StrictBool] = None
    notification_group_message_tag_all: Optional[StrictBool] = None
    notification_group_post: Optional[StrictBool] = None
    notification_group_request: Optional[StrictBool] = None
    __properties: ClassVar[List[str]] = ["notification_group_join", "notification_group_message_tag_all", "notification_group_post", "notification_group_request"]

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
        """Create an instance of Setting from a JSON string"""
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
        # set to None if notification_group_join (nullable) is None
        # and model_fields_set contains the field
        if self.notification_group_join is None and "notification_group_join" in self.model_fields_set:
            _dict['notification_group_join'] = None

        # set to None if notification_group_message_tag_all (nullable) is None
        # and model_fields_set contains the field
        if self.notification_group_message_tag_all is None and "notification_group_message_tag_all" in self.model_fields_set:
            _dict['notification_group_message_tag_all'] = None

        # set to None if notification_group_post (nullable) is None
        # and model_fields_set contains the field
        if self.notification_group_post is None and "notification_group_post" in self.model_fields_set:
            _dict['notification_group_post'] = None

        # set to None if notification_group_request (nullable) is None
        # and model_fields_set contains the field
        if self.notification_group_request is None and "notification_group_request" in self.model_fields_set:
            _dict['notification_group_request'] = None

        return _dict

    @classmethod
    def from_dict(cls, obj: Optional[Dict[str, Any]]) -> Optional[Self]:
        """Create an instance of Setting from a dict"""
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return cls.model_validate(obj)

        _obj = cls.model_validate({
            "notification_group_join": obj.get("notification_group_join"),
            "notification_group_message_tag_all": obj.get("notification_group_message_tag_all"),
            "notification_group_post": obj.get("notification_group_post"),
            "notification_group_request": obj.get("notification_group_request")
        })
        return _obj


