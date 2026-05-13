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
from typing import Optional, Set
from typing_extensions import Self

class CallStatusResponse(BaseModel):
    """
    CallStatusResponse
    """ # noqa: E501
    phone_status: Optional[StrictBool] = None
    room_url: Optional[StrictStr] = None
    video_status: Optional[StrictBool] = None
    __properties: ClassVar[List[str]] = ["phone_status", "room_url", "video_status"]

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
        """Create an instance of CallStatusResponse from a JSON string"""
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
        # set to None if phone_status (nullable) is None
        # and model_fields_set contains the field
        if self.phone_status is None and "phone_status" in self.model_fields_set:
            _dict['phone_status'] = None

        # set to None if room_url (nullable) is None
        # and model_fields_set contains the field
        if self.room_url is None and "room_url" in self.model_fields_set:
            _dict['room_url'] = None

        # set to None if video_status (nullable) is None
        # and model_fields_set contains the field
        if self.video_status is None and "video_status" in self.model_fields_set:
            _dict['video_status'] = None

        return _dict

    @classmethod
    def from_dict(cls, obj: Optional[Dict[str, Any]]) -> Optional[Self]:
        """Create an instance of CallStatusResponse from a dict"""
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return cls.model_validate(obj)

        _obj = cls.model_validate({
            "phone_status": obj.get("phone_status"),
            "room_url": obj.get("room_url"),
            "video_status": obj.get("video_status")
        })
        return _obj


