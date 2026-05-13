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
from yaylib.models.model_user_rank import ModelUserRank
from typing import Optional, Set
from typing_extensions import Self

class Web3WalletUserStatusDetails(BaseModel):
    """
    Web3WalletUserStatusDetails
    """ # noqa: E501
    activity_rank: Optional[ModelUserRank] = None
    is_rank_calculated: Optional[StrictBool] = None
    is_score_calculated: Optional[StrictBool] = None
    score: Optional[StrictInt] = None
    score_max_format: Optional[StrictStr] = None
    user_rank: Optional[ModelUserRank] = None
    __properties: ClassVar[List[str]] = ["activity_rank", "is_rank_calculated", "is_score_calculated", "score", "score_max_format", "user_rank"]

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
        """Create an instance of Web3WalletUserStatusDetails from a JSON string"""
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
        # override the default output from pydantic by calling `to_dict()` of activity_rank
        if self.activity_rank:
            _dict['activity_rank'] = self.activity_rank.to_dict()
        # override the default output from pydantic by calling `to_dict()` of user_rank
        if self.user_rank:
            _dict['user_rank'] = self.user_rank.to_dict()
        # set to None if activity_rank (nullable) is None
        # and model_fields_set contains the field
        if self.activity_rank is None and "activity_rank" in self.model_fields_set:
            _dict['activity_rank'] = None

        # set to None if is_rank_calculated (nullable) is None
        # and model_fields_set contains the field
        if self.is_rank_calculated is None and "is_rank_calculated" in self.model_fields_set:
            _dict['is_rank_calculated'] = None

        # set to None if is_score_calculated (nullable) is None
        # and model_fields_set contains the field
        if self.is_score_calculated is None and "is_score_calculated" in self.model_fields_set:
            _dict['is_score_calculated'] = None

        # set to None if score (nullable) is None
        # and model_fields_set contains the field
        if self.score is None and "score" in self.model_fields_set:
            _dict['score'] = None

        # set to None if score_max_format (nullable) is None
        # and model_fields_set contains the field
        if self.score_max_format is None and "score_max_format" in self.model_fields_set:
            _dict['score_max_format'] = None

        # set to None if user_rank (nullable) is None
        # and model_fields_set contains the field
        if self.user_rank is None and "user_rank" in self.model_fields_set:
            _dict['user_rank'] = None

        return _dict

    @classmethod
    def from_dict(cls, obj: Optional[Dict[str, Any]]) -> Optional[Self]:
        """Create an instance of Web3WalletUserStatusDetails from a dict"""
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return cls.model_validate(obj)

        _obj = cls.model_validate({
            "activity_rank": ModelUserRank.from_dict(obj["activity_rank"]) if obj.get("activity_rank") is not None else None,
            "is_rank_calculated": obj.get("is_rank_calculated"),
            "is_score_calculated": obj.get("is_score_calculated"),
            "score": obj.get("score"),
            "score_max_format": obj.get("score_max_format"),
            "user_rank": ModelUserRank.from_dict(obj["user_rank"]) if obj.get("user_rank") is not None else None
        })
        return _obj


