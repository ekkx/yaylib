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
from typing import Optional, Set
from typing_extensions import Self

class ModelVideo(BaseModel):
    """
    ModelVideo
    """ # noqa: E501
    bitrate: Optional[StrictInt] = None
    height: Optional[StrictInt] = None
    id: Optional[StrictInt] = None
    is_completed: Optional[StrictBool] = None
    thumbnail_big_url: Optional[StrictStr] = None
    thumbnail_url: Optional[StrictStr] = None
    video_url: Optional[StrictStr] = None
    views_count: Optional[StrictInt] = None
    width: Optional[StrictInt] = None
    __properties: ClassVar[List[str]] = ["bitrate", "height", "id", "is_completed", "thumbnail_big_url", "thumbnail_url", "video_url", "views_count", "width"]

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
        """Create an instance of ModelVideo from a JSON string"""
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
        # set to None if bitrate (nullable) is None
        # and model_fields_set contains the field
        if self.bitrate is None and "bitrate" in self.model_fields_set:
            _dict['bitrate'] = None

        # set to None if height (nullable) is None
        # and model_fields_set contains the field
        if self.height is None and "height" in self.model_fields_set:
            _dict['height'] = None

        # set to None if id (nullable) is None
        # and model_fields_set contains the field
        if self.id is None and "id" in self.model_fields_set:
            _dict['id'] = None

        # set to None if is_completed (nullable) is None
        # and model_fields_set contains the field
        if self.is_completed is None and "is_completed" in self.model_fields_set:
            _dict['is_completed'] = None

        # set to None if thumbnail_big_url (nullable) is None
        # and model_fields_set contains the field
        if self.thumbnail_big_url is None and "thumbnail_big_url" in self.model_fields_set:
            _dict['thumbnail_big_url'] = None

        # set to None if thumbnail_url (nullable) is None
        # and model_fields_set contains the field
        if self.thumbnail_url is None and "thumbnail_url" in self.model_fields_set:
            _dict['thumbnail_url'] = None

        # set to None if video_url (nullable) is None
        # and model_fields_set contains the field
        if self.video_url is None and "video_url" in self.model_fields_set:
            _dict['video_url'] = None

        # set to None if views_count (nullable) is None
        # and model_fields_set contains the field
        if self.views_count is None and "views_count" in self.model_fields_set:
            _dict['views_count'] = None

        # set to None if width (nullable) is None
        # and model_fields_set contains the field
        if self.width is None and "width" in self.model_fields_set:
            _dict['width'] = None

        return _dict

    @classmethod
    def from_dict(cls, obj: Optional[Dict[str, Any]]) -> Optional[Self]:
        """Create an instance of ModelVideo from a dict"""
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return cls.model_validate(obj)

        _obj = cls.model_validate({
            "bitrate": obj.get("bitrate"),
            "height": obj.get("height"),
            "id": obj.get("id"),
            "is_completed": obj.get("is_completed"),
            "thumbnail_big_url": obj.get("thumbnail_big_url"),
            "thumbnail_url": obj.get("thumbnail_url"),
            "video_url": obj.get("video_url"),
            "views_count": obj.get("views_count"),
            "width": obj.get("width")
        })
        return _obj


