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

from pydantic import BaseModel, ConfigDict, StrictStr
from typing import Any, ClassVar, Dict, List, Optional
from yaylib.models.footprint_dto import FootprintDTO
from typing import Optional, Set
from typing_extensions import Self

class FootprintsResponse(BaseModel):
    """
    FootprintsResponse
    """ # noqa: E501
    footprints: Optional[List[FootprintDTO]] = None
    next_page_value: Optional[str] = None
    __properties: ClassVar[List[str]] = ["footprints", "next_page_value"]

    model_config = ConfigDict(coerce_numbers_to_str=True, 
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
        """Create an instance of FootprintsResponse from a JSON string"""
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
        # override the default output from pydantic by calling `to_dict()` of each item in footprints (list)
        _items = []
        if self.footprints:
            for _item_footprints in self.footprints:
                if _item_footprints:
                    _items.append(_item_footprints.to_dict())
            _dict['footprints'] = _items
        # set to None if footprints (nullable) is None
        # and model_fields_set contains the field
        if self.footprints is None and "footprints" in self.model_fields_set:
            _dict['footprints'] = None

        # set to None if next_page_value (nullable) is None
        # and model_fields_set contains the field
        if self.next_page_value is None and "next_page_value" in self.model_fields_set:
            _dict['next_page_value'] = None

        return _dict

    @classmethod
    def from_dict(cls, obj: Optional[Dict[str, Any]]) -> Optional[Self]:
        """Create an instance of FootprintsResponse from a dict"""
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return cls.model_validate(obj)

        _obj = cls.model_validate({
            "footprints": [FootprintDTO.from_dict(_item) for _item in obj["footprints"]] if obj.get("footprints") is not None else None,
            "next_page_value": obj.get("next_page_value")
        })
        return _obj


