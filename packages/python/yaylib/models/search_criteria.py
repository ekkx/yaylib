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
from typing import Optional, Set
from typing_extensions import Self

class SearchCriteria(BaseModel):
    """
    SearchCriteria
    """ # noqa: E501
    biography: Optional[StrictStr] = None
    gender: Optional[StrictInt] = None
    nickname: Optional[StrictStr] = None
    prefecture: Optional[StrictStr] = None
    username: Optional[StrictStr] = None
    __properties: ClassVar[List[str]] = ["biography", "gender", "nickname", "prefecture", "username"]

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
        """Create an instance of SearchCriteria from a JSON string"""
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
        # set to None if biography (nullable) is None
        # and model_fields_set contains the field
        if self.biography is None and "biography" in self.model_fields_set:
            _dict['biography'] = None

        # set to None if gender (nullable) is None
        # and model_fields_set contains the field
        if self.gender is None and "gender" in self.model_fields_set:
            _dict['gender'] = None

        # set to None if nickname (nullable) is None
        # and model_fields_set contains the field
        if self.nickname is None and "nickname" in self.model_fields_set:
            _dict['nickname'] = None

        # set to None if prefecture (nullable) is None
        # and model_fields_set contains the field
        if self.prefecture is None and "prefecture" in self.model_fields_set:
            _dict['prefecture'] = None

        # set to None if username (nullable) is None
        # and model_fields_set contains the field
        if self.username is None and "username" in self.model_fields_set:
            _dict['username'] = None

        return _dict

    @classmethod
    def from_dict(cls, obj: Optional[Dict[str, Any]]) -> Optional[Self]:
        """Create an instance of SearchCriteria from a dict"""
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return cls.model_validate(obj)

        _obj = cls.model_validate({
            "biography": obj.get("biography"),
            "gender": obj.get("gender"),
            "nickname": obj.get("nickname"),
            "prefecture": obj.get("prefecture"),
            "username": obj.get("username")
        })
        return _obj


