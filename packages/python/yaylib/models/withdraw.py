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
from yaylib.models.asset_type import AssetType
from typing import Optional, Set
from typing_extensions import Self

class Withdraw(BaseModel):
    """
    Withdraw
    """ # noqa: E501
    asset_type: Optional[AssetType] = None
    pal_image_url: Optional[StrictStr] = None
    pal_name: Optional[StrictStr] = None
    __properties: ClassVar[List[str]] = ["asset_type", "pal_image_url", "pal_name"]

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
        """Create an instance of Withdraw from a JSON string"""
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
        # override the default output from pydantic by calling `to_dict()` of asset_type
        if self.asset_type:
            _dict['asset_type'] = self.asset_type.to_dict()
        # set to None if asset_type (nullable) is None
        # and model_fields_set contains the field
        if self.asset_type is None and "asset_type" in self.model_fields_set:
            _dict['asset_type'] = None

        # set to None if pal_image_url (nullable) is None
        # and model_fields_set contains the field
        if self.pal_image_url is None and "pal_image_url" in self.model_fields_set:
            _dict['pal_image_url'] = None

        # set to None if pal_name (nullable) is None
        # and model_fields_set contains the field
        if self.pal_name is None and "pal_name" in self.model_fields_set:
            _dict['pal_name'] = None

        return _dict

    @classmethod
    def from_dict(cls, obj: Optional[Dict[str, Any]]) -> Optional[Self]:
        """Create an instance of Withdraw from a dict"""
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return cls.model_validate(obj)

        _obj = cls.model_validate({
            "asset_type": AssetType.from_dict(obj["asset_type"]) if obj.get("asset_type") is not None else None,
            "pal_image_url": obj.get("pal_image_url"),
            "pal_name": obj.get("pal_name")
        })
        return _obj


