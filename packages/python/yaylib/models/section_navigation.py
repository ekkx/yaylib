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

class SectionNavigation(BaseModel):
    """
    SectionNavigation
    """ # noqa: E501
    icon_id: Optional[StrictInt] = None
    is_enabled: Optional[StrictBool] = None
    section_index: Optional[StrictInt] = None
    title_id: Optional[StrictInt] = None
    __properties: ClassVar[List[str]] = ["icon_id", "is_enabled", "section_index", "title_id"]

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
        """Create an instance of SectionNavigation from a JSON string"""
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
        # set to None if icon_id (nullable) is None
        # and model_fields_set contains the field
        if self.icon_id is None and "icon_id" in self.model_fields_set:
            _dict['icon_id'] = None

        # set to None if is_enabled (nullable) is None
        # and model_fields_set contains the field
        if self.is_enabled is None and "is_enabled" in self.model_fields_set:
            _dict['is_enabled'] = None

        # set to None if section_index (nullable) is None
        # and model_fields_set contains the field
        if self.section_index is None and "section_index" in self.model_fields_set:
            _dict['section_index'] = None

        # set to None if title_id (nullable) is None
        # and model_fields_set contains the field
        if self.title_id is None and "title_id" in self.model_fields_set:
            _dict['title_id'] = None

        return _dict

    @classmethod
    def from_dict(cls, obj: Optional[Dict[str, Any]]) -> Optional[Self]:
        """Create an instance of SectionNavigation from a dict"""
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return cls.model_validate(obj)

        _obj = cls.model_validate({
            "icon_id": obj.get("icon_id"),
            "is_enabled": obj.get("is_enabled"),
            "section_index": obj.get("section_index"),
            "title_id": obj.get("title_id")
        })
        return _obj


