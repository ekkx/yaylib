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
from yaylib.models.reward import Reward
from typing import Optional, Set
from typing_extensions import Self

class RaceReward(BaseModel):
    """
    RaceReward
    """ # noqa: E501
    first_prize: Optional[Reward] = None
    second_prize: Optional[Reward] = None
    third_prize: Optional[Reward] = None
    __properties: ClassVar[List[str]] = ["first_prize", "second_prize", "third_prize"]

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
        """Create an instance of RaceReward from a JSON string"""
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
        # override the default output from pydantic by calling `to_dict()` of first_prize
        if self.first_prize:
            _dict['first_prize'] = self.first_prize.to_dict()
        # override the default output from pydantic by calling `to_dict()` of second_prize
        if self.second_prize:
            _dict['second_prize'] = self.second_prize.to_dict()
        # override the default output from pydantic by calling `to_dict()` of third_prize
        if self.third_prize:
            _dict['third_prize'] = self.third_prize.to_dict()
        # set to None if first_prize (nullable) is None
        # and model_fields_set contains the field
        if self.first_prize is None and "first_prize" in self.model_fields_set:
            _dict['first_prize'] = None

        # set to None if second_prize (nullable) is None
        # and model_fields_set contains the field
        if self.second_prize is None and "second_prize" in self.model_fields_set:
            _dict['second_prize'] = None

        # set to None if third_prize (nullable) is None
        # and model_fields_set contains the field
        if self.third_prize is None and "third_prize" in self.model_fields_set:
            _dict['third_prize'] = None

        return _dict

    @classmethod
    def from_dict(cls, obj: Optional[Dict[str, Any]]) -> Optional[Self]:
        """Create an instance of RaceReward from a dict"""
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return cls.model_validate(obj)

        _obj = cls.model_validate({
            "first_prize": Reward.from_dict(obj["first_prize"]) if obj.get("first_prize") is not None else None,
            "second_prize": Reward.from_dict(obj["second_prize"]) if obj.get("second_prize") is not None else None,
            "third_prize": Reward.from_dict(obj["third_prize"]) if obj.get("third_prize") is not None else None
        })
        return _obj


