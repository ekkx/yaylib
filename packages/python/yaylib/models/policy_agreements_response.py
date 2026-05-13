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

from pydantic import BaseModel, ConfigDict, StrictBool
from typing import Any, ClassVar, Dict, List, Optional
from typing import Optional, Set
from typing_extensions import Self

class PolicyAgreementsResponse(BaseModel):
    """
    PolicyAgreementsResponse
    """ # noqa: E501
    latest_dotmoney_agreed: Optional[StrictBool] = None
    latest_empl2_agreed: Optional[StrictBool] = None
    latest_privacy_policy_agreed: Optional[StrictBool] = None
    latest_terms_of_use_agreed: Optional[StrictBool] = None
    __properties: ClassVar[List[str]] = ["latest_dotmoney_agreed", "latest_empl2_agreed", "latest_privacy_policy_agreed", "latest_terms_of_use_agreed"]

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
        """Create an instance of PolicyAgreementsResponse from a JSON string"""
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
        # set to None if latest_dotmoney_agreed (nullable) is None
        # and model_fields_set contains the field
        if self.latest_dotmoney_agreed is None and "latest_dotmoney_agreed" in self.model_fields_set:
            _dict['latest_dotmoney_agreed'] = None

        # set to None if latest_empl2_agreed (nullable) is None
        # and model_fields_set contains the field
        if self.latest_empl2_agreed is None and "latest_empl2_agreed" in self.model_fields_set:
            _dict['latest_empl2_agreed'] = None

        # set to None if latest_privacy_policy_agreed (nullable) is None
        # and model_fields_set contains the field
        if self.latest_privacy_policy_agreed is None and "latest_privacy_policy_agreed" in self.model_fields_set:
            _dict['latest_privacy_policy_agreed'] = None

        # set to None if latest_terms_of_use_agreed (nullable) is None
        # and model_fields_set contains the field
        if self.latest_terms_of_use_agreed is None and "latest_terms_of_use_agreed" in self.model_fields_set:
            _dict['latest_terms_of_use_agreed'] = None

        return _dict

    @classmethod
    def from_dict(cls, obj: Optional[Dict[str, Any]]) -> Optional[Self]:
        """Create an instance of PolicyAgreementsResponse from a dict"""
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return cls.model_validate(obj)

        _obj = cls.model_validate({
            "latest_dotmoney_agreed": obj.get("latest_dotmoney_agreed"),
            "latest_empl2_agreed": obj.get("latest_empl2_agreed"),
            "latest_privacy_policy_agreed": obj.get("latest_privacy_policy_agreed"),
            "latest_terms_of_use_agreed": obj.get("latest_terms_of_use_agreed")
        })
        return _obj


