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

class Web3WalletDailyQuest(BaseModel):
    """
    Web3WalletDailyQuest
    """ # noqa: E501
    is_calculated: Optional[StrictBool] = None
    pal_count: Optional[StrictInt] = None
    spent_duration_seconds: Optional[StrictInt] = None
    total_duration_needed_seconds: Optional[StrictInt] = None
    updated_at_millis: Optional[StrictInt] = None
    __properties: ClassVar[List[str]] = ["is_calculated", "pal_count", "spent_duration_seconds", "total_duration_needed_seconds", "updated_at_millis"]

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
        """Create an instance of Web3WalletDailyQuest from a JSON string"""
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
        # set to None if is_calculated (nullable) is None
        # and model_fields_set contains the field
        if self.is_calculated is None and "is_calculated" in self.model_fields_set:
            _dict['is_calculated'] = None

        # set to None if pal_count (nullable) is None
        # and model_fields_set contains the field
        if self.pal_count is None and "pal_count" in self.model_fields_set:
            _dict['pal_count'] = None

        # set to None if spent_duration_seconds (nullable) is None
        # and model_fields_set contains the field
        if self.spent_duration_seconds is None and "spent_duration_seconds" in self.model_fields_set:
            _dict['spent_duration_seconds'] = None

        # set to None if total_duration_needed_seconds (nullable) is None
        # and model_fields_set contains the field
        if self.total_duration_needed_seconds is None and "total_duration_needed_seconds" in self.model_fields_set:
            _dict['total_duration_needed_seconds'] = None

        # set to None if updated_at_millis (nullable) is None
        # and model_fields_set contains the field
        if self.updated_at_millis is None and "updated_at_millis" in self.model_fields_set:
            _dict['updated_at_millis'] = None

        return _dict

    @classmethod
    def from_dict(cls, obj: Optional[Dict[str, Any]]) -> Optional[Self]:
        """Create an instance of Web3WalletDailyQuest from a dict"""
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return cls.model_validate(obj)

        _obj = cls.model_validate({
            "is_calculated": obj.get("is_calculated"),
            "pal_count": obj.get("pal_count"),
            "spent_duration_seconds": obj.get("spent_duration_seconds"),
            "total_duration_needed_seconds": obj.get("total_duration_needed_seconds"),
            "updated_at_millis": obj.get("updated_at_millis")
        })
        return _obj


