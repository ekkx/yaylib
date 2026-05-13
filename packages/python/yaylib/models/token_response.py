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

class TokenResponse(BaseModel):
    """
    TokenResponse
    """ # noqa: E501
    access_token: Optional[StrictStr] = None
    created_at: Optional[StrictInt] = None
    expires_in: Optional[StrictInt] = None
    id: Optional[StrictInt] = None
    refresh_token: Optional[StrictStr] = None
    __properties: ClassVar[List[str]] = ["access_token", "created_at", "expires_in", "id", "refresh_token"]

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
        """Create an instance of TokenResponse from a JSON string"""
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
        # set to None if access_token (nullable) is None
        # and model_fields_set contains the field
        if self.access_token is None and "access_token" in self.model_fields_set:
            _dict['access_token'] = None

        # set to None if created_at (nullable) is None
        # and model_fields_set contains the field
        if self.created_at is None and "created_at" in self.model_fields_set:
            _dict['created_at'] = None

        # set to None if expires_in (nullable) is None
        # and model_fields_set contains the field
        if self.expires_in is None and "expires_in" in self.model_fields_set:
            _dict['expires_in'] = None

        # set to None if id (nullable) is None
        # and model_fields_set contains the field
        if self.id is None and "id" in self.model_fields_set:
            _dict['id'] = None

        # set to None if refresh_token (nullable) is None
        # and model_fields_set contains the field
        if self.refresh_token is None and "refresh_token" in self.model_fields_set:
            _dict['refresh_token'] = None

        return _dict

    @classmethod
    def from_dict(cls, obj: Optional[Dict[str, Any]]) -> Optional[Self]:
        """Create an instance of TokenResponse from a dict"""
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return cls.model_validate(obj)

        _obj = cls.model_validate({
            "access_token": obj.get("access_token"),
            "created_at": obj.get("created_at"),
            "expires_in": obj.get("expires_in"),
            "id": obj.get("id"),
            "refresh_token": obj.get("refresh_token")
        })
        return _obj


