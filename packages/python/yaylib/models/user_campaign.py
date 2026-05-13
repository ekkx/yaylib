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
from yaylib.models.multiplier_breakdown import MultiplierBreakdown
from typing import Optional, Set
from typing_extensions import Self

class UserCampaign(BaseModel):
    """
    UserCampaign
    """ # noqa: E501
    multiplier: Optional[Union[StrictFloat, StrictInt]] = None
    multiplier_breakdown: Optional[MultiplierBreakdown] = None
    points: Optional[StrictInt] = None
    ranking: Optional[StrictInt] = None
    __properties: ClassVar[List[str]] = ["multiplier", "multiplier_breakdown", "points", "ranking"]

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
        """Create an instance of UserCampaign from a JSON string"""
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
        # override the default output from pydantic by calling `to_dict()` of multiplier_breakdown
        if self.multiplier_breakdown:
            _dict['multiplier_breakdown'] = self.multiplier_breakdown.to_dict()
        # set to None if multiplier (nullable) is None
        # and model_fields_set contains the field
        if self.multiplier is None and "multiplier" in self.model_fields_set:
            _dict['multiplier'] = None

        # set to None if multiplier_breakdown (nullable) is None
        # and model_fields_set contains the field
        if self.multiplier_breakdown is None and "multiplier_breakdown" in self.model_fields_set:
            _dict['multiplier_breakdown'] = None

        # set to None if points (nullable) is None
        # and model_fields_set contains the field
        if self.points is None and "points" in self.model_fields_set:
            _dict['points'] = None

        # set to None if ranking (nullable) is None
        # and model_fields_set contains the field
        if self.ranking is None and "ranking" in self.model_fields_set:
            _dict['ranking'] = None

        return _dict

    @classmethod
    def from_dict(cls, obj: Optional[Dict[str, Any]]) -> Optional[Self]:
        """Create an instance of UserCampaign from a dict"""
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return cls.model_validate(obj)

        _obj = cls.model_validate({
            "multiplier": obj.get("multiplier"),
            "multiplier_breakdown": MultiplierBreakdown.from_dict(obj["multiplier_breakdown"]) if obj.get("multiplier_breakdown") is not None else None,
            "points": obj.get("points"),
            "ranking": obj.get("ranking")
        })
        return _obj


