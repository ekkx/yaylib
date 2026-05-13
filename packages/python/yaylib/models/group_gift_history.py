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

from pydantic import BaseModel, ConfigDict, StrictInt
from typing import Any, ClassVar, Dict, List, Optional
from yaylib.models.gift_count import GiftCount
from yaylib.models.realm_user import RealmUser
from typing import Optional, Set
from typing_extensions import Self

class GroupGiftHistory(BaseModel):
    """
    GroupGiftHistory
    """ # noqa: E501
    gifts_count: Optional[List[GiftCount]] = None
    received_date: Optional[StrictInt] = None
    user: Optional[RealmUser] = None
    __properties: ClassVar[List[str]] = ["gifts_count", "received_date", "user"]

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
        """Create an instance of GroupGiftHistory from a JSON string"""
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
        # override the default output from pydantic by calling `to_dict()` of each item in gifts_count (list)
        _items = []
        if self.gifts_count:
            for _item_gifts_count in self.gifts_count:
                if _item_gifts_count:
                    _items.append(_item_gifts_count.to_dict())
            _dict['gifts_count'] = _items
        # override the default output from pydantic by calling `to_dict()` of user
        if self.user:
            _dict['user'] = self.user.to_dict()
        # set to None if gifts_count (nullable) is None
        # and model_fields_set contains the field
        if self.gifts_count is None and "gifts_count" in self.model_fields_set:
            _dict['gifts_count'] = None

        # set to None if received_date (nullable) is None
        # and model_fields_set contains the field
        if self.received_date is None and "received_date" in self.model_fields_set:
            _dict['received_date'] = None

        # set to None if user (nullable) is None
        # and model_fields_set contains the field
        if self.user is None and "user" in self.model_fields_set:
            _dict['user'] = None

        return _dict

    @classmethod
    def from_dict(cls, obj: Optional[Dict[str, Any]]) -> Optional[Self]:
        """Create an instance of GroupGiftHistory from a dict"""
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return cls.model_validate(obj)

        _obj = cls.model_validate({
            "gifts_count": [GiftCount.from_dict(_item) for _item in obj["gifts_count"]] if obj.get("gifts_count") is not None else None,
            "received_date": obj.get("received_date"),
            "user": RealmUser.from_dict(obj["user"]) if obj.get("user") is not None else None
        })
        return _obj


