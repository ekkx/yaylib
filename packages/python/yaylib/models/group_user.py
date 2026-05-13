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

class GroupUser(BaseModel):
    """
    GroupUser
    """ # noqa: E501
    is_moderator: Optional[StrictBool] = None
    is_mute: Optional[StrictBool] = None
    pending_deputize: Optional[StrictBool] = None
    pending_transfer: Optional[StrictBool] = None
    title: Optional[StrictStr] = None
    user: Optional[RealmUser] = None
    __properties: ClassVar[List[str]] = ["is_moderator", "is_mute", "pending_deputize", "pending_transfer", "title", "user"]

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
        """Create an instance of GroupUser from a JSON string"""
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
        # override the default output from pydantic by calling `to_dict()` of user
        if self.user:
            _dict['user'] = self.user.to_dict()
        # set to None if is_moderator (nullable) is None
        # and model_fields_set contains the field
        if self.is_moderator is None and "is_moderator" in self.model_fields_set:
            _dict['is_moderator'] = None

        # set to None if is_mute (nullable) is None
        # and model_fields_set contains the field
        if self.is_mute is None and "is_mute" in self.model_fields_set:
            _dict['is_mute'] = None

        # set to None if pending_deputize (nullable) is None
        # and model_fields_set contains the field
        if self.pending_deputize is None and "pending_deputize" in self.model_fields_set:
            _dict['pending_deputize'] = None

        # set to None if pending_transfer (nullable) is None
        # and model_fields_set contains the field
        if self.pending_transfer is None and "pending_transfer" in self.model_fields_set:
            _dict['pending_transfer'] = None

        # set to None if title (nullable) is None
        # and model_fields_set contains the field
        if self.title is None and "title" in self.model_fields_set:
            _dict['title'] = None

        # set to None if user (nullable) is None
        # and model_fields_set contains the field
        if self.user is None and "user" in self.model_fields_set:
            _dict['user'] = None

        return _dict

    @classmethod
    def from_dict(cls, obj: Optional[Dict[str, Any]]) -> Optional[Self]:
        """Create an instance of GroupUser from a dict"""
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return cls.model_validate(obj)

        _obj = cls.model_validate({
            "is_moderator": obj.get("is_moderator"),
            "is_mute": obj.get("is_mute"),
            "pending_deputize": obj.get("pending_deputize"),
            "pending_transfer": obj.get("pending_transfer"),
            "title": obj.get("title"),
            "user": RealmUser.from_dict(obj["user"]) if obj.get("user") is not None else None
        })
        return _obj

from yaylib.models.realm_user import RealmUser
# TODO: Rewrite to not use raise_errors
GroupUser.model_rebuild(raise_errors=False)

