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

from pydantic import BaseModel, ConfigDict, StrictStr
from typing import Any, ClassVar, Dict, List, Optional
from typing import Optional, Set
from typing_extensions import Self

class LoginSnsUserRequest(BaseModel):
    """
    LoginSnsUserRequest
    """ # noqa: E501
    access_token: Optional[StrictStr] = None
    access_token_secret: Optional[StrictStr] = None
    api_key: Optional[StrictStr] = None
    authorization_code: Optional[StrictStr] = None
    consumer_key: Optional[StrictStr] = None
    consumer_secret: Optional[StrictStr] = None
    email: Optional[StrictStr] = None
    provider: Optional[StrictStr] = None
    two_f_a_code: Optional[StrictStr] = None
    __properties: ClassVar[List[str]] = ["access_token", "access_token_secret", "api_key", "authorization_code", "consumer_key", "consumer_secret", "email", "provider", "two_f_a_code"]

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
        """Create an instance of LoginSnsUserRequest from a JSON string"""
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

        # set to None if access_token_secret (nullable) is None
        # and model_fields_set contains the field
        if self.access_token_secret is None and "access_token_secret" in self.model_fields_set:
            _dict['access_token_secret'] = None

        # set to None if api_key (nullable) is None
        # and model_fields_set contains the field
        if self.api_key is None and "api_key" in self.model_fields_set:
            _dict['api_key'] = None

        # set to None if authorization_code (nullable) is None
        # and model_fields_set contains the field
        if self.authorization_code is None and "authorization_code" in self.model_fields_set:
            _dict['authorization_code'] = None

        # set to None if consumer_key (nullable) is None
        # and model_fields_set contains the field
        if self.consumer_key is None and "consumer_key" in self.model_fields_set:
            _dict['consumer_key'] = None

        # set to None if consumer_secret (nullable) is None
        # and model_fields_set contains the field
        if self.consumer_secret is None and "consumer_secret" in self.model_fields_set:
            _dict['consumer_secret'] = None

        # set to None if email (nullable) is None
        # and model_fields_set contains the field
        if self.email is None and "email" in self.model_fields_set:
            _dict['email'] = None

        # set to None if provider (nullable) is None
        # and model_fields_set contains the field
        if self.provider is None and "provider" in self.model_fields_set:
            _dict['provider'] = None

        # set to None if two_f_a_code (nullable) is None
        # and model_fields_set contains the field
        if self.two_f_a_code is None and "two_f_a_code" in self.model_fields_set:
            _dict['two_f_a_code'] = None

        return _dict

    @classmethod
    def from_dict(cls, obj: Optional[Dict[str, Any]]) -> Optional[Self]:
        """Create an instance of LoginSnsUserRequest from a dict"""
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return cls.model_validate(obj)

        _obj = cls.model_validate({
            "access_token": obj.get("access_token"),
            "access_token_secret": obj.get("access_token_secret"),
            "api_key": obj.get("api_key"),
            "authorization_code": obj.get("authorization_code"),
            "consumer_key": obj.get("consumer_key"),
            "consumer_secret": obj.get("consumer_secret"),
            "email": obj.get("email"),
            "provider": obj.get("provider"),
            "two_f_a_code": obj.get("two_f_a_code")
        })
        return _obj


