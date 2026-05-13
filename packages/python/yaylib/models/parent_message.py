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
from yaylib.models.gif_image import GifImage
from yaylib.models.message_type import MessageType
from yaylib.models.sticker import Sticker
from typing import Optional, Set
from typing_extensions import Self

class ParentMessage(BaseModel):
    """
    ParentMessage
    """ # noqa: E501
    attachment: Optional[StrictStr] = None
    attachment_thumbnail: Optional[StrictStr] = None
    created_at: Optional[StrictInt] = None
    font_size: Optional[StrictInt] = None
    gif: Optional[GifImage] = None
    id: Optional[StrictInt] = None
    message_type: Optional[MessageType] = None
    room_id: Optional[StrictInt] = None
    sticker: Optional[Sticker] = None
    text: Optional[StrictStr] = None
    user_id: Optional[StrictInt] = None
    video_thumbnail_url: Optional[StrictStr] = None
    video_url: Optional[StrictStr] = None
    __properties: ClassVar[List[str]] = ["attachment", "attachment_thumbnail", "created_at", "font_size", "gif", "id", "message_type", "room_id", "sticker", "text", "user_id", "video_thumbnail_url", "video_url"]

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
        """Create an instance of ParentMessage from a JSON string"""
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
        # override the default output from pydantic by calling `to_dict()` of gif
        if self.gif:
            _dict['gif'] = self.gif.to_dict()
        # override the default output from pydantic by calling `to_dict()` of sticker
        if self.sticker:
            _dict['sticker'] = self.sticker.to_dict()
        # set to None if attachment (nullable) is None
        # and model_fields_set contains the field
        if self.attachment is None and "attachment" in self.model_fields_set:
            _dict['attachment'] = None

        # set to None if attachment_thumbnail (nullable) is None
        # and model_fields_set contains the field
        if self.attachment_thumbnail is None and "attachment_thumbnail" in self.model_fields_set:
            _dict['attachment_thumbnail'] = None

        # set to None if created_at (nullable) is None
        # and model_fields_set contains the field
        if self.created_at is None and "created_at" in self.model_fields_set:
            _dict['created_at'] = None

        # set to None if font_size (nullable) is None
        # and model_fields_set contains the field
        if self.font_size is None and "font_size" in self.model_fields_set:
            _dict['font_size'] = None

        # set to None if gif (nullable) is None
        # and model_fields_set contains the field
        if self.gif is None and "gif" in self.model_fields_set:
            _dict['gif'] = None

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

        # set to None if sticker (nullable) is None
        # and model_fields_set contains the field
        if self.sticker is None and "sticker" in self.model_fields_set:
            _dict['sticker'] = None

        # set to None if text (nullable) is None
        # and model_fields_set contains the field
        if self.text is None and "text" in self.model_fields_set:
            _dict['text'] = None

        # set to None if user_id (nullable) is None
        # and model_fields_set contains the field
        if self.user_id is None and "user_id" in self.model_fields_set:
            _dict['user_id'] = None

        # set to None if video_thumbnail_url (nullable) is None
        # and model_fields_set contains the field
        if self.video_thumbnail_url is None and "video_thumbnail_url" in self.model_fields_set:
            _dict['video_thumbnail_url'] = None

        # set to None if video_url (nullable) is None
        # and model_fields_set contains the field
        if self.video_url is None and "video_url" in self.model_fields_set:
            _dict['video_url'] = None

        return _dict

    @classmethod
    def from_dict(cls, obj: Optional[Dict[str, Any]]) -> Optional[Self]:
        """Create an instance of ParentMessage from a dict"""
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return cls.model_validate(obj)

        _obj = cls.model_validate({
            "attachment": obj.get("attachment"),
            "attachment_thumbnail": obj.get("attachment_thumbnail"),
            "created_at": obj.get("created_at"),
            "font_size": obj.get("font_size"),
            "gif": GifImage.from_dict(obj["gif"]) if obj.get("gif") is not None else None,
            "id": obj.get("id"),
            "message_type": obj.get("message_type"),
            "room_id": obj.get("room_id"),
            "sticker": Sticker.from_dict(obj["sticker"]) if obj.get("sticker") is not None else None,
            "text": obj.get("text"),
            "user_id": obj.get("user_id"),
            "video_thumbnail_url": obj.get("video_thumbnail_url"),
            "video_url": obj.get("video_url")
        })
        return _obj


