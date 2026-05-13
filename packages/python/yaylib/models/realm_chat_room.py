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

from pydantic import BaseModel, ConfigDict, StrictBool, StrictInt, StrictStr
from typing import Any, ClassVar, Dict, List, Optional
from yaylib.models.chat_room_last_message import ChatRoomLastMessage
from yaylib.models.realm_user import RealmUser
from yaylib.models.user_setting import UserSetting
from typing import Optional, Set
from typing_extensions import Self

class RealmChatRoom(BaseModel):
    """
    RealmChatRoom
    """ # noqa: E501
    background: Optional[StrictStr] = None
    id: Optional[StrictInt] = None
    is_group: Optional[StrictBool] = None
    is_request: Optional[StrictBool] = None
    last_message: Optional[ChatRoomLastMessage] = None
    members: Optional[Dict[str, Any]] = None
    name: Optional[StrictStr] = None
    owner: Optional[RealmUser] = None
    unread_count: Optional[StrictInt] = None
    updated_at: Optional[StrictInt] = None
    user_setting: Optional[UserSetting] = None
    __properties: ClassVar[List[str]] = ["background", "id", "is_group", "is_request", "last_message", "members", "name", "owner", "unread_count", "updated_at", "user_setting"]

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
        """Create an instance of RealmChatRoom from a JSON string"""
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
        # override the default output from pydantic by calling `to_dict()` of last_message
        if self.last_message:
            _dict['last_message'] = self.last_message.to_dict()
        # override the default output from pydantic by calling `to_dict()` of owner
        if self.owner:
            _dict['owner'] = self.owner.to_dict()
        # override the default output from pydantic by calling `to_dict()` of user_setting
        if self.user_setting:
            _dict['user_setting'] = self.user_setting.to_dict()
        # set to None if background (nullable) is None
        # and model_fields_set contains the field
        if self.background is None and "background" in self.model_fields_set:
            _dict['background'] = None

        # set to None if id (nullable) is None
        # and model_fields_set contains the field
        if self.id is None and "id" in self.model_fields_set:
            _dict['id'] = None

        # set to None if is_group (nullable) is None
        # and model_fields_set contains the field
        if self.is_group is None and "is_group" in self.model_fields_set:
            _dict['is_group'] = None

        # set to None if is_request (nullable) is None
        # and model_fields_set contains the field
        if self.is_request is None and "is_request" in self.model_fields_set:
            _dict['is_request'] = None

        # set to None if last_message (nullable) is None
        # and model_fields_set contains the field
        if self.last_message is None and "last_message" in self.model_fields_set:
            _dict['last_message'] = None

        # set to None if members (nullable) is None
        # and model_fields_set contains the field
        if self.members is None and "members" in self.model_fields_set:
            _dict['members'] = None

        # set to None if name (nullable) is None
        # and model_fields_set contains the field
        if self.name is None and "name" in self.model_fields_set:
            _dict['name'] = None

        # set to None if owner (nullable) is None
        # and model_fields_set contains the field
        if self.owner is None and "owner" in self.model_fields_set:
            _dict['owner'] = None

        # set to None if unread_count (nullable) is None
        # and model_fields_set contains the field
        if self.unread_count is None and "unread_count" in self.model_fields_set:
            _dict['unread_count'] = None

        # set to None if updated_at (nullable) is None
        # and model_fields_set contains the field
        if self.updated_at is None and "updated_at" in self.model_fields_set:
            _dict['updated_at'] = None

        # set to None if user_setting (nullable) is None
        # and model_fields_set contains the field
        if self.user_setting is None and "user_setting" in self.model_fields_set:
            _dict['user_setting'] = None

        return _dict

    @classmethod
    def from_dict(cls, obj: Optional[Dict[str, Any]]) -> Optional[Self]:
        """Create an instance of RealmChatRoom from a dict"""
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return cls.model_validate(obj)

        _obj = cls.model_validate({
            "background": obj.get("background"),
            "id": obj.get("id"),
            "is_group": obj.get("is_group"),
            "is_request": obj.get("is_request"),
            "last_message": ChatRoomLastMessage.from_dict(obj["last_message"]) if obj.get("last_message") is not None else None,
            "members": obj.get("members"),
            "name": obj.get("name"),
            "owner": RealmUser.from_dict(obj["owner"]) if obj.get("owner") is not None else None,
            "unread_count": obj.get("unread_count"),
            "updated_at": obj.get("updated_at"),
            "user_setting": UserSetting.from_dict(obj["user_setting"]) if obj.get("user_setting") is not None else None
        })
        return _obj


