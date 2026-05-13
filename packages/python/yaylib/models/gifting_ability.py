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

from pydantic import BaseModel, ConfigDict, StrictBool, StrictInt
from typing import Any, ClassVar, Dict, List, Optional
from typing import Optional, Set
from typing_extensions import Self

class GiftingAbility(BaseModel):
    """
    GiftingAbility
    """ # noqa: E501
    can_receive: Optional[StrictBool] = None
    can_send: Optional[StrictBool] = None
    enabled: Optional[StrictBool] = None
    user_id: Optional[StrictInt] = None
    __properties: ClassVar[List[str]] = ["can_receive", "can_send", "enabled", "user_id"]

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
        """Create an instance of GiftingAbility from a JSON string"""
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
        # set to None if can_receive (nullable) is None
        # and model_fields_set contains the field
        if self.can_receive is None and "can_receive" in self.model_fields_set:
            _dict['can_receive'] = None

        # set to None if can_send (nullable) is None
        # and model_fields_set contains the field
        if self.can_send is None and "can_send" in self.model_fields_set:
            _dict['can_send'] = None

        # set to None if enabled (nullable) is None
        # and model_fields_set contains the field
        if self.enabled is None and "enabled" in self.model_fields_set:
            _dict['enabled'] = None

        # set to None if user_id (nullable) is None
        # and model_fields_set contains the field
        if self.user_id is None and "user_id" in self.model_fields_set:
            _dict['user_id'] = None

        return _dict

    @classmethod
    def from_dict(cls, obj: Optional[Dict[str, Any]]) -> Optional[Self]:
        """Create an instance of GiftingAbility from a dict"""
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return cls.model_validate(obj)

        _obj = cls.model_validate({
            "can_receive": obj.get("can_receive"),
            "can_send": obj.get("can_send"),
            "enabled": obj.get("enabled"),
            "user_id": obj.get("user_id")
        })
        return _obj


