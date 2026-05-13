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

class VideoProcessedData(BaseModel):
    """
    VideoProcessedData
    """ # noqa: E501
    id: Optional[StrictInt] = None
    video_processed: Optional[StrictBool] = None
    video_thumbnail_big_url: Optional[StrictStr] = None
    video_thumbnail_url: Optional[StrictStr] = None
    video_url: Optional[StrictStr] = None
    __properties: ClassVar[List[str]] = ["id", "video_processed", "video_thumbnail_big_url", "video_thumbnail_url", "video_url"]

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
        """Create an instance of VideoProcessedData from a JSON string"""
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
        # set to None if id (nullable) is None
        # and model_fields_set contains the field
        if self.id is None and "id" in self.model_fields_set:
            _dict['id'] = None

        # set to None if video_processed (nullable) is None
        # and model_fields_set contains the field
        if self.video_processed is None and "video_processed" in self.model_fields_set:
            _dict['video_processed'] = None

        # set to None if video_thumbnail_big_url (nullable) is None
        # and model_fields_set contains the field
        if self.video_thumbnail_big_url is None and "video_thumbnail_big_url" in self.model_fields_set:
            _dict['video_thumbnail_big_url'] = None

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
        """Create an instance of VideoProcessedData from a dict"""
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return cls.model_validate(obj)

        _obj = cls.model_validate({
            "id": obj.get("id"),
            "video_processed": obj.get("video_processed"),
            "video_thumbnail_big_url": obj.get("video_thumbnail_big_url"),
            "video_thumbnail_url": obj.get("video_thumbnail_url"),
            "video_url": obj.get("video_url")
        })
        return _obj


