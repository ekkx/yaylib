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
from typing import Optional, Set
from typing_extensions import Self

class PalRaceLeagueDTO(BaseModel):
    """
    PalRaceLeagueDTO
    """ # noqa: E501
    accessible: Optional[StrictBool] = None
    entry_fee: Optional[StrictInt] = None
    game_end_at: Optional[StrictInt] = None
    name: Optional[StrictStr] = None
    price_reward: Optional[StrictInt] = None
    registration_close_at: Optional[StrictInt] = None
    __properties: ClassVar[List[str]] = ["accessible", "entry_fee", "game_end_at", "name", "price_reward", "registration_close_at"]

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
        """Create an instance of PalRaceLeagueDTO from a JSON string"""
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
        # set to None if accessible (nullable) is None
        # and model_fields_set contains the field
        if self.accessible is None and "accessible" in self.model_fields_set:
            _dict['accessible'] = None

        # set to None if entry_fee (nullable) is None
        # and model_fields_set contains the field
        if self.entry_fee is None and "entry_fee" in self.model_fields_set:
            _dict['entry_fee'] = None

        # set to None if game_end_at (nullable) is None
        # and model_fields_set contains the field
        if self.game_end_at is None and "game_end_at" in self.model_fields_set:
            _dict['game_end_at'] = None

        # set to None if name (nullable) is None
        # and model_fields_set contains the field
        if self.name is None and "name" in self.model_fields_set:
            _dict['name'] = None

        # set to None if price_reward (nullable) is None
        # and model_fields_set contains the field
        if self.price_reward is None and "price_reward" in self.model_fields_set:
            _dict['price_reward'] = None

        # set to None if registration_close_at (nullable) is None
        # and model_fields_set contains the field
        if self.registration_close_at is None and "registration_close_at" in self.model_fields_set:
            _dict['registration_close_at'] = None

        return _dict

    @classmethod
    def from_dict(cls, obj: Optional[Dict[str, Any]]) -> Optional[Self]:
        """Create an instance of PalRaceLeagueDTO from a dict"""
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return cls.model_validate(obj)

        _obj = cls.model_validate({
            "accessible": obj.get("accessible"),
            "entry_fee": obj.get("entry_fee"),
            "game_end_at": obj.get("game_end_at"),
            "name": obj.get("name"),
            "price_reward": obj.get("price_reward"),
            "registration_close_at": obj.get("registration_close_at")
        })
        return _obj


