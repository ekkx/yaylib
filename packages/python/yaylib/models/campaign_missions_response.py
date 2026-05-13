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

from pydantic import BaseModel, ConfigDict
from typing import Any, ClassVar, Dict, List, Optional
from yaylib.models.mission_dto import MissionDTO
from typing import Optional, Set
from typing_extensions import Self

class CampaignMissionsResponse(BaseModel):
    """
    CampaignMissionsResponse
    """ # noqa: E501
    daily_missions: Optional[List[MissionDTO]] = None
    multiplier_missions: Optional[List[MissionDTO]] = None
    normal_missions: Optional[List[MissionDTO]] = None
    unlimited_missions: Optional[List[MissionDTO]] = None
    __properties: ClassVar[List[str]] = ["daily_missions", "multiplier_missions", "normal_missions", "unlimited_missions"]

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
        """Create an instance of CampaignMissionsResponse from a JSON string"""
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
        # override the default output from pydantic by calling `to_dict()` of each item in daily_missions (list)
        _items = []
        if self.daily_missions:
            for _item_daily_missions in self.daily_missions:
                if _item_daily_missions:
                    _items.append(_item_daily_missions.to_dict())
            _dict['daily_missions'] = _items
        # override the default output from pydantic by calling `to_dict()` of each item in multiplier_missions (list)
        _items = []
        if self.multiplier_missions:
            for _item_multiplier_missions in self.multiplier_missions:
                if _item_multiplier_missions:
                    _items.append(_item_multiplier_missions.to_dict())
            _dict['multiplier_missions'] = _items
        # override the default output from pydantic by calling `to_dict()` of each item in normal_missions (list)
        _items = []
        if self.normal_missions:
            for _item_normal_missions in self.normal_missions:
                if _item_normal_missions:
                    _items.append(_item_normal_missions.to_dict())
            _dict['normal_missions'] = _items
        # override the default output from pydantic by calling `to_dict()` of each item in unlimited_missions (list)
        _items = []
        if self.unlimited_missions:
            for _item_unlimited_missions in self.unlimited_missions:
                if _item_unlimited_missions:
                    _items.append(_item_unlimited_missions.to_dict())
            _dict['unlimited_missions'] = _items
        # set to None if daily_missions (nullable) is None
        # and model_fields_set contains the field
        if self.daily_missions is None and "daily_missions" in self.model_fields_set:
            _dict['daily_missions'] = None

        # set to None if multiplier_missions (nullable) is None
        # and model_fields_set contains the field
        if self.multiplier_missions is None and "multiplier_missions" in self.model_fields_set:
            _dict['multiplier_missions'] = None

        # set to None if normal_missions (nullable) is None
        # and model_fields_set contains the field
        if self.normal_missions is None and "normal_missions" in self.model_fields_set:
            _dict['normal_missions'] = None

        # set to None if unlimited_missions (nullable) is None
        # and model_fields_set contains the field
        if self.unlimited_missions is None and "unlimited_missions" in self.model_fields_set:
            _dict['unlimited_missions'] = None

        return _dict

    @classmethod
    def from_dict(cls, obj: Optional[Dict[str, Any]]) -> Optional[Self]:
        """Create an instance of CampaignMissionsResponse from a dict"""
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return cls.model_validate(obj)

        _obj = cls.model_validate({
            "daily_missions": [MissionDTO.from_dict(_item) for _item in obj["daily_missions"]] if obj.get("daily_missions") is not None else None,
            "multiplier_missions": [MissionDTO.from_dict(_item) for _item in obj["multiplier_missions"]] if obj.get("multiplier_missions") is not None else None,
            "normal_missions": [MissionDTO.from_dict(_item) for _item in obj["normal_missions"]] if obj.get("normal_missions") is not None else None,
            "unlimited_missions": [MissionDTO.from_dict(_item) for _item in obj["unlimited_missions"]] if obj.get("unlimited_missions") is not None else None
        })
        return _obj


