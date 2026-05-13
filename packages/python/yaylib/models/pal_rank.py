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
from yaylib.models.pal import Pal
from typing import Optional, Set
from typing_extensions import Self

class PalRank(BaseModel):
    """
    PalRank
    """ # noqa: E501
    own: Optional[StrictBool] = None
    pal: Optional[Pal] = None
    place: Optional[StrictInt] = None
    yay_user_name: Optional[StrictStr] = None
    __properties: ClassVar[List[str]] = ["own", "pal", "place", "yay_user_name"]

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
        """Create an instance of PalRank from a JSON string"""
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
        # override the default output from pydantic by calling `to_dict()` of pal
        if self.pal:
            _dict['pal'] = self.pal.to_dict()
        # set to None if own (nullable) is None
        # and model_fields_set contains the field
        if self.own is None and "own" in self.model_fields_set:
            _dict['own'] = None

        # set to None if pal (nullable) is None
        # and model_fields_set contains the field
        if self.pal is None and "pal" in self.model_fields_set:
            _dict['pal'] = None

        # set to None if place (nullable) is None
        # and model_fields_set contains the field
        if self.place is None and "place" in self.model_fields_set:
            _dict['place'] = None

        # set to None if yay_user_name (nullable) is None
        # and model_fields_set contains the field
        if self.yay_user_name is None and "yay_user_name" in self.model_fields_set:
            _dict['yay_user_name'] = None

        return _dict

    @classmethod
    def from_dict(cls, obj: Optional[Dict[str, Any]]) -> Optional[Self]:
        """Create an instance of PalRank from a dict"""
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return cls.model_validate(obj)

        _obj = cls.model_validate({
            "own": obj.get("own"),
            "pal": Pal.from_dict(obj["pal"]) if obj.get("pal") is not None else None,
            "place": obj.get("place"),
            "yay_user_name": obj.get("yay_user_name")
        })
        return _obj


