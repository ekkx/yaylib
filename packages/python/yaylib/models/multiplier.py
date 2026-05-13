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

from pydantic import BaseModel, ConfigDict, Field, StrictStr
from typing import Any, ClassVar, Dict, List, Optional
from typing import Optional, Set
from typing_extensions import Self

class Multiplier(BaseModel):
    """
    Multiplier
    """ # noqa: E501
    activity: Optional[StrictStr] = None
    base: Optional[StrictStr] = None
    var_date: Optional[StrictStr] = Field(default=None, alias="date")
    mission: Optional[StrictStr] = None
    multiplier: Optional[StrictStr] = None
    vip: Optional[StrictStr] = None
    __properties: ClassVar[List[str]] = ["activity", "base", "date", "mission", "multiplier", "vip"]

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
        """Create an instance of Multiplier from a JSON string"""
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
        # set to None if activity (nullable) is None
        # and model_fields_set contains the field
        if self.activity is None and "activity" in self.model_fields_set:
            _dict['activity'] = None

        # set to None if base (nullable) is None
        # and model_fields_set contains the field
        if self.base is None and "base" in self.model_fields_set:
            _dict['base'] = None

        # set to None if var_date (nullable) is None
        # and model_fields_set contains the field
        if self.var_date is None and "var_date" in self.model_fields_set:
            _dict['date'] = None

        # set to None if mission (nullable) is None
        # and model_fields_set contains the field
        if self.mission is None and "mission" in self.model_fields_set:
            _dict['mission'] = None

        # set to None if multiplier (nullable) is None
        # and model_fields_set contains the field
        if self.multiplier is None and "multiplier" in self.model_fields_set:
            _dict['multiplier'] = None

        # set to None if vip (nullable) is None
        # and model_fields_set contains the field
        if self.vip is None and "vip" in self.model_fields_set:
            _dict['vip'] = None

        return _dict

    @classmethod
    def from_dict(cls, obj: Optional[Dict[str, Any]]) -> Optional[Self]:
        """Create an instance of Multiplier from a dict"""
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return cls.model_validate(obj)

        _obj = cls.model_validate({
            "activity": obj.get("activity"),
            "base": obj.get("base"),
            "date": obj.get("date"),
            "mission": obj.get("mission"),
            "multiplier": obj.get("multiplier"),
            "vip": obj.get("vip")
        })
        return _obj


