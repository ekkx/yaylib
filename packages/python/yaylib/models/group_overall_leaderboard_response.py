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

from pydantic import BaseModel, ConfigDict, StrictInt
from typing import Any, ClassVar, Dict, List, Optional
from yaylib.models.group_leaderboard import GroupLeaderboard
from typing import Optional, Set
from typing_extensions import Self

class GroupOverallLeaderboardResponse(BaseModel):
    """
    GroupOverallLeaderboardResponse
    """ # noqa: E501
    group_leaderboard: Optional[List[GroupLeaderboard]] = None
    updated_at: Optional[StrictInt] = None
    __properties: ClassVar[List[str]] = ["group_leaderboard", "updated_at"]

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
        """Create an instance of GroupOverallLeaderboardResponse from a JSON string"""
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
        # override the default output from pydantic by calling `to_dict()` of each item in group_leaderboard (list)
        _items = []
        if self.group_leaderboard:
            for _item_group_leaderboard in self.group_leaderboard:
                if _item_group_leaderboard:
                    _items.append(_item_group_leaderboard.to_dict())
            _dict['group_leaderboard'] = _items
        # set to None if group_leaderboard (nullable) is None
        # and model_fields_set contains the field
        if self.group_leaderboard is None and "group_leaderboard" in self.model_fields_set:
            _dict['group_leaderboard'] = None

        # set to None if updated_at (nullable) is None
        # and model_fields_set contains the field
        if self.updated_at is None and "updated_at" in self.model_fields_set:
            _dict['updated_at'] = None

        return _dict

    @classmethod
    def from_dict(cls, obj: Optional[Dict[str, Any]]) -> Optional[Self]:
        """Create an instance of GroupOverallLeaderboardResponse from a dict"""
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return cls.model_validate(obj)

        _obj = cls.model_validate({
            "group_leaderboard": [GroupLeaderboard.from_dict(_item) for _item in obj["group_leaderboard"]] if obj.get("group_leaderboard") is not None else None,
            "updated_at": obj.get("updated_at")
        })
        return _obj


