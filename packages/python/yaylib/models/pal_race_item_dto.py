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

from pydantic import BaseModel, ConfigDict, StrictInt, StrictStr
from typing import Any, ClassVar, Dict, List, Optional
from yaylib.models.profile_image_dto import ProfileImageDTO
from typing import Optional, Set
from typing_extensions import Self

class PalRaceItemDTO(BaseModel):
    """
    PalRaceItemDTO
    """ # noqa: E501
    current_level: Optional[StrictInt] = None
    fatigue_level: Optional[StrictStr] = None
    nickname: Optional[StrictStr] = None
    overall_strength: Optional[StrictInt] = None
    profile_image: Optional[ProfileImageDTO] = None
    rank: Optional[StrictStr] = None
    token_id: Optional[StrictStr] = None
    type: Optional[StrictStr] = None
    win_rate: Optional[StrictStr] = None
    winning_streak: Optional[StrictInt] = None
    __properties: ClassVar[List[str]] = ["current_level", "fatigue_level", "nickname", "overall_strength", "profile_image", "rank", "token_id", "type", "win_rate", "winning_streak"]

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
        """Create an instance of PalRaceItemDTO from a JSON string"""
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
        # override the default output from pydantic by calling `to_dict()` of profile_image
        if self.profile_image:
            _dict['profile_image'] = self.profile_image.to_dict()
        # set to None if current_level (nullable) is None
        # and model_fields_set contains the field
        if self.current_level is None and "current_level" in self.model_fields_set:
            _dict['current_level'] = None

        # set to None if fatigue_level (nullable) is None
        # and model_fields_set contains the field
        if self.fatigue_level is None and "fatigue_level" in self.model_fields_set:
            _dict['fatigue_level'] = None

        # set to None if nickname (nullable) is None
        # and model_fields_set contains the field
        if self.nickname is None and "nickname" in self.model_fields_set:
            _dict['nickname'] = None

        # set to None if overall_strength (nullable) is None
        # and model_fields_set contains the field
        if self.overall_strength is None and "overall_strength" in self.model_fields_set:
            _dict['overall_strength'] = None

        # set to None if profile_image (nullable) is None
        # and model_fields_set contains the field
        if self.profile_image is None and "profile_image" in self.model_fields_set:
            _dict['profile_image'] = None

        # set to None if rank (nullable) is None
        # and model_fields_set contains the field
        if self.rank is None and "rank" in self.model_fields_set:
            _dict['rank'] = None

        # set to None if token_id (nullable) is None
        # and model_fields_set contains the field
        if self.token_id is None and "token_id" in self.model_fields_set:
            _dict['token_id'] = None

        # set to None if type (nullable) is None
        # and model_fields_set contains the field
        if self.type is None and "type" in self.model_fields_set:
            _dict['type'] = None

        # set to None if win_rate (nullable) is None
        # and model_fields_set contains the field
        if self.win_rate is None and "win_rate" in self.model_fields_set:
            _dict['win_rate'] = None

        # set to None if winning_streak (nullable) is None
        # and model_fields_set contains the field
        if self.winning_streak is None and "winning_streak" in self.model_fields_set:
            _dict['winning_streak'] = None

        return _dict

    @classmethod
    def from_dict(cls, obj: Optional[Dict[str, Any]]) -> Optional[Self]:
        """Create an instance of PalRaceItemDTO from a dict"""
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return cls.model_validate(obj)

        _obj = cls.model_validate({
            "current_level": obj.get("current_level"),
            "fatigue_level": obj.get("fatigue_level"),
            "nickname": obj.get("nickname"),
            "overall_strength": obj.get("overall_strength"),
            "profile_image": ProfileImageDTO.from_dict(obj["profile_image"]) if obj.get("profile_image") is not None else None,
            "rank": obj.get("rank"),
            "token_id": obj.get("token_id"),
            "type": obj.get("type"),
            "win_rate": obj.get("win_rate"),
            "winning_streak": obj.get("winning_streak")
        })
        return _obj


