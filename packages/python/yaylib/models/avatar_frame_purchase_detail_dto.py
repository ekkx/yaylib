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

class AvatarFramePurchaseDetailDTO(BaseModel):
    """
    AvatarFramePurchaseDetailDTO
    """ # noqa: E501
    frame: Optional[StrictStr] = None
    frame_thumbnail: Optional[StrictStr] = None
    frame_type: Optional[StrictStr] = None
    id: Optional[StrictInt] = None
    vip_only: Optional[StrictBool] = None
    __properties: ClassVar[List[str]] = ["frame", "frame_thumbnail", "frame_type", "id", "vip_only"]

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
        """Create an instance of AvatarFramePurchaseDetailDTO from a JSON string"""
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
        # set to None if frame (nullable) is None
        # and model_fields_set contains the field
        if self.frame is None and "frame" in self.model_fields_set:
            _dict['frame'] = None

        # set to None if frame_thumbnail (nullable) is None
        # and model_fields_set contains the field
        if self.frame_thumbnail is None and "frame_thumbnail" in self.model_fields_set:
            _dict['frame_thumbnail'] = None

        # set to None if frame_type (nullable) is None
        # and model_fields_set contains the field
        if self.frame_type is None and "frame_type" in self.model_fields_set:
            _dict['frame_type'] = None

        # set to None if id (nullable) is None
        # and model_fields_set contains the field
        if self.id is None and "id" in self.model_fields_set:
            _dict['id'] = None

        # set to None if vip_only (nullable) is None
        # and model_fields_set contains the field
        if self.vip_only is None and "vip_only" in self.model_fields_set:
            _dict['vip_only'] = None

        return _dict

    @classmethod
    def from_dict(cls, obj: Optional[Dict[str, Any]]) -> Optional[Self]:
        """Create an instance of AvatarFramePurchaseDetailDTO from a dict"""
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return cls.model_validate(obj)

        _obj = cls.model_validate({
            "frame": obj.get("frame"),
            "frame_thumbnail": obj.get("frame_thumbnail"),
            "frame_type": obj.get("frame_type"),
            "id": obj.get("id"),
            "vip_only": obj.get("vip_only")
        })
        return _obj


