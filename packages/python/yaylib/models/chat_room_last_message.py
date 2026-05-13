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

from pydantic import BaseModel, ConfigDict, StrictInt, StrictStr
from typing import Any, ClassVar, Dict, List, Optional
from yaylib.models.message_type import MessageType
from yaylib.models.realm_conference_call import RealmConferenceCall
from typing import Optional, Set
from typing_extensions import Self

class ChatRoomLastMessage(BaseModel):
    """
    ChatRoomLastMessage
    """ # noqa: E501
    conference_call: Optional[RealmConferenceCall] = None
    created_at: Optional[StrictInt] = None
    id: Optional[StrictInt] = None
    message_type: Optional[MessageType] = None
    room_id: Optional[StrictInt] = None
    text: Optional[StrictStr] = None
    user_id: Optional[StrictInt] = None
    __properties: ClassVar[List[str]] = ["conference_call", "created_at", "id", "message_type", "room_id", "text", "user_id"]

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
        """Create an instance of ChatRoomLastMessage from a JSON string"""
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
        # override the default output from pydantic by calling `to_dict()` of conference_call
        if self.conference_call:
            _dict['conference_call'] = self.conference_call.to_dict()
        # set to None if conference_call (nullable) is None
        # and model_fields_set contains the field
        if self.conference_call is None and "conference_call" in self.model_fields_set:
            _dict['conference_call'] = None

        # set to None if created_at (nullable) is None
        # and model_fields_set contains the field
        if self.created_at is None and "created_at" in self.model_fields_set:
            _dict['created_at'] = None

        # set to None if id (nullable) is None
        # and model_fields_set contains the field
        if self.id is None and "id" in self.model_fields_set:
            _dict['id'] = None

        # set to None if message_type (nullable) is None
        # and model_fields_set contains the field
        if self.message_type is None and "message_type" in self.model_fields_set:
            _dict['message_type'] = None

        # set to None if room_id (nullable) is None
        # and model_fields_set contains the field
        if self.room_id is None and "room_id" in self.model_fields_set:
            _dict['room_id'] = None

        # set to None if text (nullable) is None
        # and model_fields_set contains the field
        if self.text is None and "text" in self.model_fields_set:
            _dict['text'] = None

        # set to None if user_id (nullable) is None
        # and model_fields_set contains the field
        if self.user_id is None and "user_id" in self.model_fields_set:
            _dict['user_id'] = None

        return _dict

    @classmethod
    def from_dict(cls, obj: Optional[Dict[str, Any]]) -> Optional[Self]:
        """Create an instance of ChatRoomLastMessage from a dict"""
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return cls.model_validate(obj)

        _obj = cls.model_validate({
            "conference_call": RealmConferenceCall.from_dict(obj["conference_call"]) if obj.get("conference_call") is not None else None,
            "created_at": obj.get("created_at"),
            "id": obj.get("id"),
            "message_type": obj.get("message_type"),
            "room_id": obj.get("room_id"),
            "text": obj.get("text"),
            "user_id": obj.get("user_id")
        })
        return _obj


