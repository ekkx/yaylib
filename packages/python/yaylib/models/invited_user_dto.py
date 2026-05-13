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
from yaylib.models.invited_user_details_dto import InvitedUserDetailsDTO
from typing import Optional, Set
from typing_extensions import Self

class InvitedUserDTO(BaseModel):
    """
    InvitedUserDTO
    """ # noqa: E501
    referred: Optional[InvitedUserDetailsDTO] = None
    total_earned_points: Optional[StrictInt] = None
    __properties: ClassVar[List[str]] = ["referred", "total_earned_points"]

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
        """Create an instance of InvitedUserDTO from a JSON string"""
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
        # override the default output from pydantic by calling `to_dict()` of referred
        if self.referred:
            _dict['referred'] = self.referred.to_dict()
        # set to None if referred (nullable) is None
        # and model_fields_set contains the field
        if self.referred is None and "referred" in self.model_fields_set:
            _dict['referred'] = None

        # set to None if total_earned_points (nullable) is None
        # and model_fields_set contains the field
        if self.total_earned_points is None and "total_earned_points" in self.model_fields_set:
            _dict['total_earned_points'] = None

        return _dict

    @classmethod
    def from_dict(cls, obj: Optional[Dict[str, Any]]) -> Optional[Self]:
        """Create an instance of InvitedUserDTO from a dict"""
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return cls.model_validate(obj)

        _obj = cls.model_validate({
            "referred": InvitedUserDetailsDTO.from_dict(obj["referred"]) if obj.get("referred") is not None else None,
            "total_earned_points": obj.get("total_earned_points")
        })
        return _obj


