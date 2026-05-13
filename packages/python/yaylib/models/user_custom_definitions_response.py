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

class UserCustomDefinitionsResponse(BaseModel):
    """
    UserCustomDefinitionsResponse
    """ # noqa: E501
    age: Optional[StrictInt] = None
    created_at: Optional[StrictInt] = None
    followers_count: Optional[StrictInt] = None
    followings_count: Optional[StrictInt] = None
    last_loggedin_at: Optional[StrictInt] = None
    reported_count: Optional[StrictInt] = None
    status: Optional[StrictStr] = None
    __properties: ClassVar[List[str]] = ["age", "created_at", "followers_count", "followings_count", "last_loggedin_at", "reported_count", "status"]

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
        """Create an instance of UserCustomDefinitionsResponse from a JSON string"""
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
        # set to None if age (nullable) is None
        # and model_fields_set contains the field
        if self.age is None and "age" in self.model_fields_set:
            _dict['age'] = None

        # set to None if created_at (nullable) is None
        # and model_fields_set contains the field
        if self.created_at is None and "created_at" in self.model_fields_set:
            _dict['created_at'] = None

        # set to None if followers_count (nullable) is None
        # and model_fields_set contains the field
        if self.followers_count is None and "followers_count" in self.model_fields_set:
            _dict['followers_count'] = None

        # set to None if followings_count (nullable) is None
        # and model_fields_set contains the field
        if self.followings_count is None and "followings_count" in self.model_fields_set:
            _dict['followings_count'] = None

        # set to None if last_loggedin_at (nullable) is None
        # and model_fields_set contains the field
        if self.last_loggedin_at is None and "last_loggedin_at" in self.model_fields_set:
            _dict['last_loggedin_at'] = None

        # set to None if reported_count (nullable) is None
        # and model_fields_set contains the field
        if self.reported_count is None and "reported_count" in self.model_fields_set:
            _dict['reported_count'] = None

        # set to None if status (nullable) is None
        # and model_fields_set contains the field
        if self.status is None and "status" in self.model_fields_set:
            _dict['status'] = None

        return _dict

    @classmethod
    def from_dict(cls, obj: Optional[Dict[str, Any]]) -> Optional[Self]:
        """Create an instance of UserCustomDefinitionsResponse from a dict"""
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return cls.model_validate(obj)

        _obj = cls.model_validate({
            "age": obj.get("age"),
            "created_at": obj.get("created_at"),
            "followers_count": obj.get("followers_count"),
            "followings_count": obj.get("followings_count"),
            "last_loggedin_at": obj.get("last_loggedin_at"),
            "reported_count": obj.get("reported_count"),
            "status": obj.get("status")
        })
        return _obj


