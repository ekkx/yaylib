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

from pydantic import BaseModel, ConfigDict, StrictFloat, StrictInt
from typing import Any, ClassVar, Dict, List, Optional, Union
from typing import Optional, Set
from typing_extensions import Self

class MultiplierBreakdown(BaseModel):
    """
    MultiplierBreakdown
    """ # noqa: E501
    activity: Optional[Union[StrictFloat, StrictInt]] = None
    special_mission: Optional[Union[StrictFloat, StrictInt]] = None
    vip_multiplier: Optional[Union[StrictFloat, StrictInt]] = None
    __properties: ClassVar[List[str]] = ["activity", "special_mission", "vip_multiplier"]

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
        """Create an instance of MultiplierBreakdown from a JSON string"""
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

        # set to None if special_mission (nullable) is None
        # and model_fields_set contains the field
        if self.special_mission is None and "special_mission" in self.model_fields_set:
            _dict['special_mission'] = None

        # set to None if vip_multiplier (nullable) is None
        # and model_fields_set contains the field
        if self.vip_multiplier is None and "vip_multiplier" in self.model_fields_set:
            _dict['vip_multiplier'] = None

        return _dict

    @classmethod
    def from_dict(cls, obj: Optional[Dict[str, Any]]) -> Optional[Self]:
        """Create an instance of MultiplierBreakdown from a dict"""
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return cls.model_validate(obj)

        _obj = cls.model_validate({
            "activity": obj.get("activity"),
            "special_mission": obj.get("special_mission"),
            "vip_multiplier": obj.get("vip_multiplier")
        })
        return _obj


