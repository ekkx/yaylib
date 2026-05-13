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

from pydantic import BaseModel, ConfigDict
from typing import Any, ClassVar, Dict, List, Optional
from yaylib.models.empl_expiring_response import EmplExpiringResponse
from typing import Optional, Set
from typing_extensions import Self

class EmplExpiringListResponse(BaseModel):
    """
    EmplExpiringListResponse
    """ # noqa: E501
    regular_expiring: Optional[List[EmplExpiringResponse]] = None
    rewarded_expiring: Optional[List[EmplExpiringResponse]] = None
    __properties: ClassVar[List[str]] = ["regular_expiring", "rewarded_expiring"]

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
        """Create an instance of EmplExpiringListResponse from a JSON string"""
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
        # override the default output from pydantic by calling `to_dict()` of each item in regular_expiring (list)
        _items = []
        if self.regular_expiring:
            for _item_regular_expiring in self.regular_expiring:
                if _item_regular_expiring:
                    _items.append(_item_regular_expiring.to_dict())
            _dict['regular_expiring'] = _items
        # override the default output from pydantic by calling `to_dict()` of each item in rewarded_expiring (list)
        _items = []
        if self.rewarded_expiring:
            for _item_rewarded_expiring in self.rewarded_expiring:
                if _item_rewarded_expiring:
                    _items.append(_item_rewarded_expiring.to_dict())
            _dict['rewarded_expiring'] = _items
        # set to None if regular_expiring (nullable) is None
        # and model_fields_set contains the field
        if self.regular_expiring is None and "regular_expiring" in self.model_fields_set:
            _dict['regular_expiring'] = None

        # set to None if rewarded_expiring (nullable) is None
        # and model_fields_set contains the field
        if self.rewarded_expiring is None and "rewarded_expiring" in self.model_fields_set:
            _dict['rewarded_expiring'] = None

        return _dict

    @classmethod
    def from_dict(cls, obj: Optional[Dict[str, Any]]) -> Optional[Self]:
        """Create an instance of EmplExpiringListResponse from a dict"""
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return cls.model_validate(obj)

        _obj = cls.model_validate({
            "regular_expiring": [EmplExpiringResponse.from_dict(_item) for _item in obj["regular_expiring"]] if obj.get("regular_expiring") is not None else None,
            "rewarded_expiring": [EmplExpiringResponse.from_dict(_item) for _item in obj["rewarded_expiring"]] if obj.get("rewarded_expiring") is not None else None
        })
        return _obj


