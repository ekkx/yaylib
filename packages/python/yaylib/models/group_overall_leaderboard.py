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
from yaylib.models.group_ranking import GroupRanking
from typing import Optional, Set
from typing_extensions import Self

class GroupOverallLeaderboard(BaseModel):
    """
    GroupOverallLeaderboard
    """ # noqa: E501
    group_rankings: Optional[List[GroupRanking]] = None
    updated_at_millis: Optional[StrictInt] = None
    __properties: ClassVar[List[str]] = ["group_rankings", "updated_at_millis"]

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
        """Create an instance of GroupOverallLeaderboard from a JSON string"""
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
        # override the default output from pydantic by calling `to_dict()` of each item in group_rankings (list)
        _items = []
        if self.group_rankings:
            for _item_group_rankings in self.group_rankings:
                if _item_group_rankings:
                    _items.append(_item_group_rankings.to_dict())
            _dict['group_rankings'] = _items
        # set to None if group_rankings (nullable) is None
        # and model_fields_set contains the field
        if self.group_rankings is None and "group_rankings" in self.model_fields_set:
            _dict['group_rankings'] = None

        # set to None if updated_at_millis (nullable) is None
        # and model_fields_set contains the field
        if self.updated_at_millis is None and "updated_at_millis" in self.model_fields_set:
            _dict['updated_at_millis'] = None

        return _dict

    @classmethod
    def from_dict(cls, obj: Optional[Dict[str, Any]]) -> Optional[Self]:
        """Create an instance of GroupOverallLeaderboard from a dict"""
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return cls.model_validate(obj)

        _obj = cls.model_validate({
            "group_rankings": [GroupRanking.from_dict(_item) for _item in obj["group_rankings"]] if obj.get("group_rankings") is not None else None,
            "updated_at_millis": obj.get("updated_at_millis")
        })
        return _obj


