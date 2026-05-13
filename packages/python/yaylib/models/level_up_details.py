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
from yaylib.models.level_up_details_pal import LevelUpDetailsPal
from yaylib.models.state_changes import StateChanges
from typing import Optional, Set
from typing_extensions import Self

class LevelUpDetails(BaseModel):
    """
    LevelUpDetails
    """ # noqa: E501
    available: Optional[StrictBool] = None
    cost: Optional[StrictStr] = None
    empl_earn_limit: Optional[StrictStr] = None
    level: Optional[StrictInt] = None
    pal: Optional[LevelUpDetailsPal] = None
    stat_changes: Optional[StateChanges] = None
    __properties: ClassVar[List[str]] = ["available", "cost", "empl_earn_limit", "level", "pal", "stat_changes"]

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
        """Create an instance of LevelUpDetails from a JSON string"""
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
        # override the default output from pydantic by calling `to_dict()` of stat_changes
        if self.stat_changes:
            _dict['stat_changes'] = self.stat_changes.to_dict()
        # set to None if available (nullable) is None
        # and model_fields_set contains the field
        if self.available is None and "available" in self.model_fields_set:
            _dict['available'] = None

        # set to None if cost (nullable) is None
        # and model_fields_set contains the field
        if self.cost is None and "cost" in self.model_fields_set:
            _dict['cost'] = None

        # set to None if empl_earn_limit (nullable) is None
        # and model_fields_set contains the field
        if self.empl_earn_limit is None and "empl_earn_limit" in self.model_fields_set:
            _dict['empl_earn_limit'] = None

        # set to None if level (nullable) is None
        # and model_fields_set contains the field
        if self.level is None and "level" in self.model_fields_set:
            _dict['level'] = None

        # set to None if pal (nullable) is None
        # and model_fields_set contains the field
        if self.pal is None and "pal" in self.model_fields_set:
            _dict['pal'] = None

        # set to None if stat_changes (nullable) is None
        # and model_fields_set contains the field
        if self.stat_changes is None and "stat_changes" in self.model_fields_set:
            _dict['stat_changes'] = None

        return _dict

    @classmethod
    def from_dict(cls, obj: Optional[Dict[str, Any]]) -> Optional[Self]:
        """Create an instance of LevelUpDetails from a dict"""
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return cls.model_validate(obj)

        _obj = cls.model_validate({
            "available": obj.get("available"),
            "cost": obj.get("cost"),
            "empl_earn_limit": obj.get("empl_earn_limit"),
            "level": obj.get("level"),
            "pal": LevelUpDetailsPal.from_dict(obj["pal"]) if obj.get("pal") is not None else None,
            "stat_changes": StateChanges.from_dict(obj["stat_changes"]) if obj.get("stat_changes") is not None else None
        })
        return _obj


