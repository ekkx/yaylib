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
from yaylib.models.realm_user import RealmUser
from typing import Optional, Set
from typing_extensions import Self

class HiddenResponse(BaseModel):
    """
    HiddenResponse
    """ # noqa: E501
    hidden_users: Optional[List[RealmUser]] = None
    limit: Optional[StrictInt] = None
    next_page_value: Optional[StrictInt] = None
    total_count: Optional[StrictInt] = None
    __properties: ClassVar[List[str]] = ["hidden_users", "limit", "next_page_value", "total_count"]

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
        """Create an instance of HiddenResponse from a JSON string"""
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
        # override the default output from pydantic by calling `to_dict()` of each item in hidden_users (list)
        _items = []
        if self.hidden_users:
            for _item_hidden_users in self.hidden_users:
                if _item_hidden_users:
                    _items.append(_item_hidden_users.to_dict())
            _dict['hidden_users'] = _items
        # set to None if hidden_users (nullable) is None
        # and model_fields_set contains the field
        if self.hidden_users is None and "hidden_users" in self.model_fields_set:
            _dict['hidden_users'] = None

        # set to None if limit (nullable) is None
        # and model_fields_set contains the field
        if self.limit is None and "limit" in self.model_fields_set:
            _dict['limit'] = None

        # set to None if next_page_value (nullable) is None
        # and model_fields_set contains the field
        if self.next_page_value is None and "next_page_value" in self.model_fields_set:
            _dict['next_page_value'] = None

        # set to None if total_count (nullable) is None
        # and model_fields_set contains the field
        if self.total_count is None and "total_count" in self.model_fields_set:
            _dict['total_count'] = None

        return _dict

    @classmethod
    def from_dict(cls, obj: Optional[Dict[str, Any]]) -> Optional[Self]:
        """Create an instance of HiddenResponse from a dict"""
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return cls.model_validate(obj)

        _obj = cls.model_validate({
            "hidden_users": [RealmUser.from_dict(_item) for _item in obj["hidden_users"]] if obj.get("hidden_users") is not None else None,
            "limit": obj.get("limit"),
            "next_page_value": obj.get("next_page_value"),
            "total_count": obj.get("total_count")
        })
        return _obj


