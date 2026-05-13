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
from yaylib.models.realm_user import RealmUser
from typing import Optional, Set
from typing_extensions import Self

class UserRank(BaseModel):
    """
    UserRank
    """ # noqa: E501
    ranking: Optional[StrictInt] = None
    top_gifts: Optional[List[StrictStr]] = None
    total_scores: Optional[StrictInt] = None
    updated_at: Optional[StrictInt] = None
    user: Optional[RealmUser] = None
    __properties: ClassVar[List[str]] = ["ranking", "top_gifts", "total_scores", "updated_at", "user"]

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
        """Create an instance of UserRank from a JSON string"""
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
        # override the default output from pydantic by calling `to_dict()` of user
        if self.user:
            _dict['user'] = self.user.to_dict()
        # set to None if ranking (nullable) is None
        # and model_fields_set contains the field
        if self.ranking is None and "ranking" in self.model_fields_set:
            _dict['ranking'] = None

        # set to None if top_gifts (nullable) is None
        # and model_fields_set contains the field
        if self.top_gifts is None and "top_gifts" in self.model_fields_set:
            _dict['top_gifts'] = None

        # set to None if total_scores (nullable) is None
        # and model_fields_set contains the field
        if self.total_scores is None and "total_scores" in self.model_fields_set:
            _dict['total_scores'] = None

        # set to None if updated_at (nullable) is None
        # and model_fields_set contains the field
        if self.updated_at is None and "updated_at" in self.model_fields_set:
            _dict['updated_at'] = None

        # set to None if user (nullable) is None
        # and model_fields_set contains the field
        if self.user is None and "user" in self.model_fields_set:
            _dict['user'] = None

        return _dict

    @classmethod
    def from_dict(cls, obj: Optional[Dict[str, Any]]) -> Optional[Self]:
        """Create an instance of UserRank from a dict"""
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return cls.model_validate(obj)

        _obj = cls.model_validate({
            "ranking": obj.get("ranking"),
            "top_gifts": obj.get("top_gifts"),
            "total_scores": obj.get("total_scores"),
            "updated_at": obj.get("updated_at"),
            "user": RealmUser.from_dict(obj["user"]) if obj.get("user") is not None else None
        })
        return _obj


