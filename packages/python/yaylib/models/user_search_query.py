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

from pydantic import BaseModel, ConfigDict, StrictBool, StrictStr
from typing import Any, ClassVar, Dict, List, Optional
from yaylib.models.user_search_scope import UserSearchScope
from typing import Optional, Set
from typing_extensions import Self

class UserSearchQuery(BaseModel):
    """
    UserSearchQuery
    """ # noqa: E501
    keyword: Optional[StrictStr] = None
    no_recent_penalty: Optional[StrictBool] = None
    same_prefecture: Optional[StrictBool] = None
    scope: Optional[UserSearchScope] = None
    similar_age: Optional[StrictBool] = None
    __properties: ClassVar[List[str]] = ["keyword", "no_recent_penalty", "same_prefecture", "scope", "similar_age"]

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
        """Create an instance of UserSearchQuery from a JSON string"""
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
        # override the default output from pydantic by calling `to_dict()` of scope
        if self.scope:
            _dict['scope'] = self.scope.to_dict()
        # set to None if keyword (nullable) is None
        # and model_fields_set contains the field
        if self.keyword is None and "keyword" in self.model_fields_set:
            _dict['keyword'] = None

        # set to None if no_recent_penalty (nullable) is None
        # and model_fields_set contains the field
        if self.no_recent_penalty is None and "no_recent_penalty" in self.model_fields_set:
            _dict['no_recent_penalty'] = None

        # set to None if same_prefecture (nullable) is None
        # and model_fields_set contains the field
        if self.same_prefecture is None and "same_prefecture" in self.model_fields_set:
            _dict['same_prefecture'] = None

        # set to None if scope (nullable) is None
        # and model_fields_set contains the field
        if self.scope is None and "scope" in self.model_fields_set:
            _dict['scope'] = None

        # set to None if similar_age (nullable) is None
        # and model_fields_set contains the field
        if self.similar_age is None and "similar_age" in self.model_fields_set:
            _dict['similar_age'] = None

        return _dict

    @classmethod
    def from_dict(cls, obj: Optional[Dict[str, Any]]) -> Optional[Self]:
        """Create an instance of UserSearchQuery from a dict"""
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return cls.model_validate(obj)

        _obj = cls.model_validate({
            "keyword": obj.get("keyword"),
            "no_recent_penalty": obj.get("no_recent_penalty"),
            "same_prefecture": obj.get("same_prefecture"),
            "scope": UserSearchScope.from_dict(obj["scope"]) if obj.get("scope") is not None else None,
            "similar_age": obj.get("similar_age")
        })
        return _obj


