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

from pydantic import BaseModel, ConfigDict, StrictFloat, StrictInt, StrictStr
from typing import Any, ClassVar, Dict, List, Optional, Union
from yaylib.models.breakdown import Breakdown
from yaylib.models.mission import Mission
from typing import Optional, Set
from typing_extensions import Self

class CampaignPointHistoryDTO(BaseModel):
    """
    CampaignPointHistoryDTO
    """ # noqa: E501
    breakdown: Optional[Breakdown] = None
    created_at: Optional[StrictInt] = None
    id: Optional[StrictInt] = None
    mission: Optional[Mission] = None
    multiplier: Optional[Union[StrictFloat, StrictInt]] = None
    point_type: Optional[StrictStr] = None
    total_points: Optional[StrictInt] = None
    __properties: ClassVar[List[str]] = ["breakdown", "created_at", "id", "mission", "multiplier", "point_type", "total_points"]

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
        """Create an instance of CampaignPointHistoryDTO from a JSON string"""
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
        # override the default output from pydantic by calling `to_dict()` of breakdown
        if self.breakdown:
            _dict['breakdown'] = self.breakdown.to_dict()
        # override the default output from pydantic by calling `to_dict()` of mission
        if self.mission:
            _dict['mission'] = self.mission.to_dict()
        # set to None if breakdown (nullable) is None
        # and model_fields_set contains the field
        if self.breakdown is None and "breakdown" in self.model_fields_set:
            _dict['breakdown'] = None

        # set to None if created_at (nullable) is None
        # and model_fields_set contains the field
        if self.created_at is None and "created_at" in self.model_fields_set:
            _dict['created_at'] = None

        # set to None if id (nullable) is None
        # and model_fields_set contains the field
        if self.id is None and "id" in self.model_fields_set:
            _dict['id'] = None

        # set to None if mission (nullable) is None
        # and model_fields_set contains the field
        if self.mission is None and "mission" in self.model_fields_set:
            _dict['mission'] = None

        # set to None if multiplier (nullable) is None
        # and model_fields_set contains the field
        if self.multiplier is None and "multiplier" in self.model_fields_set:
            _dict['multiplier'] = None

        # set to None if point_type (nullable) is None
        # and model_fields_set contains the field
        if self.point_type is None and "point_type" in self.model_fields_set:
            _dict['point_type'] = None

        # set to None if total_points (nullable) is None
        # and model_fields_set contains the field
        if self.total_points is None and "total_points" in self.model_fields_set:
            _dict['total_points'] = None

        return _dict

    @classmethod
    def from_dict(cls, obj: Optional[Dict[str, Any]]) -> Optional[Self]:
        """Create an instance of CampaignPointHistoryDTO from a dict"""
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return cls.model_validate(obj)

        _obj = cls.model_validate({
            "breakdown": Breakdown.from_dict(obj["breakdown"]) if obj.get("breakdown") is not None else None,
            "created_at": obj.get("created_at"),
            "id": obj.get("id"),
            "mission": Mission.from_dict(obj["mission"]) if obj.get("mission") is not None else None,
            "multiplier": obj.get("multiplier"),
            "point_type": obj.get("point_type"),
            "total_points": obj.get("total_points")
        })
        return _obj


