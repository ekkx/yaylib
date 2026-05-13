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

class CommonErrorResponse(BaseModel):
    """
    CommonErrorResponse
    """ # noqa: E501
    ban_until: Optional[StrictInt] = None
    error_code: Optional[StrictInt] = None
    message: Optional[StrictStr] = None
    remaining_quota: Optional[StrictInt] = None
    retry_in: Optional[StrictInt] = None
    __properties: ClassVar[List[str]] = ["ban_until", "error_code", "message", "remaining_quota", "retry_in"]

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
        """Create an instance of CommonErrorResponse from a JSON string"""
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
        # set to None if ban_until (nullable) is None
        # and model_fields_set contains the field
        if self.ban_until is None and "ban_until" in self.model_fields_set:
            _dict['ban_until'] = None

        # set to None if error_code (nullable) is None
        # and model_fields_set contains the field
        if self.error_code is None and "error_code" in self.model_fields_set:
            _dict['error_code'] = None

        # set to None if message (nullable) is None
        # and model_fields_set contains the field
        if self.message is None and "message" in self.model_fields_set:
            _dict['message'] = None

        # set to None if remaining_quota (nullable) is None
        # and model_fields_set contains the field
        if self.remaining_quota is None and "remaining_quota" in self.model_fields_set:
            _dict['remaining_quota'] = None

        # set to None if retry_in (nullable) is None
        # and model_fields_set contains the field
        if self.retry_in is None and "retry_in" in self.model_fields_set:
            _dict['retry_in'] = None

        return _dict

    @classmethod
    def from_dict(cls, obj: Optional[Dict[str, Any]]) -> Optional[Self]:
        """Create an instance of CommonErrorResponse from a dict"""
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return cls.model_validate(obj)

        _obj = cls.model_validate({
            "ban_until": obj.get("ban_until"),
            "error_code": obj.get("error_code"),
            "message": obj.get("message"),
            "remaining_quota": obj.get("remaining_quota"),
            "retry_in": obj.get("retry_in")
        })
        return _obj


