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
from yaylib.models.mission_type import MissionType
from typing import Optional, Set
from typing_extensions import Self

class Mission(BaseModel):
    """
    Mission
    """ # noqa: E501
    action: Optional[StrictStr] = None
    detail: Optional[StrictStr] = None
    is_multiplier: Optional[StrictBool] = None
    mission_type: Optional[MissionType] = None
    required_actions_count: Optional[StrictInt] = None
    __properties: ClassVar[List[str]] = ["action", "detail", "is_multiplier", "mission_type", "required_actions_count"]

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
        """Create an instance of Mission from a JSON string"""
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
        # set to None if action (nullable) is None
        # and model_fields_set contains the field
        if self.action is None and "action" in self.model_fields_set:
            _dict['action'] = None

        # set to None if detail (nullable) is None
        # and model_fields_set contains the field
        if self.detail is None and "detail" in self.model_fields_set:
            _dict['detail'] = None

        # set to None if is_multiplier (nullable) is None
        # and model_fields_set contains the field
        if self.is_multiplier is None and "is_multiplier" in self.model_fields_set:
            _dict['is_multiplier'] = None

        # set to None if mission_type (nullable) is None
        # and model_fields_set contains the field
        if self.mission_type is None and "mission_type" in self.model_fields_set:
            _dict['mission_type'] = None

        # set to None if required_actions_count (nullable) is None
        # and model_fields_set contains the field
        if self.required_actions_count is None and "required_actions_count" in self.model_fields_set:
            _dict['required_actions_count'] = None

        return _dict

    @classmethod
    def from_dict(cls, obj: Optional[Dict[str, Any]]) -> Optional[Self]:
        """Create an instance of Mission from a dict"""
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return cls.model_validate(obj)

        _obj = cls.model_validate({
            "action": obj.get("action"),
            "detail": obj.get("detail"),
            "is_multiplier": obj.get("is_multiplier"),
            "mission_type": obj.get("mission_type"),
            "required_actions_count": obj.get("required_actions_count")
        })
        return _obj


