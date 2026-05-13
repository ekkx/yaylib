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

class PalGacha(BaseModel):
    """
    PalGacha
    """ # noqa: E501
    banner_collection_url: Optional[StrictStr] = None
    banner_url: Optional[StrictStr] = None
    detail_url: Optional[StrictStr] = None
    gacha_type: Optional[StrictStr] = None
    id: Optional[StrictStr] = None
    price: Optional[StrictInt] = None
    __properties: ClassVar[List[str]] = ["banner_collection_url", "banner_url", "detail_url", "gacha_type", "id", "price"]

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
        """Create an instance of PalGacha from a JSON string"""
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
        # set to None if banner_collection_url (nullable) is None
        # and model_fields_set contains the field
        if self.banner_collection_url is None and "banner_collection_url" in self.model_fields_set:
            _dict['banner_collection_url'] = None

        # set to None if banner_url (nullable) is None
        # and model_fields_set contains the field
        if self.banner_url is None and "banner_url" in self.model_fields_set:
            _dict['banner_url'] = None

        # set to None if detail_url (nullable) is None
        # and model_fields_set contains the field
        if self.detail_url is None and "detail_url" in self.model_fields_set:
            _dict['detail_url'] = None

        # set to None if gacha_type (nullable) is None
        # and model_fields_set contains the field
        if self.gacha_type is None and "gacha_type" in self.model_fields_set:
            _dict['gacha_type'] = None

        # set to None if id (nullable) is None
        # and model_fields_set contains the field
        if self.id is None and "id" in self.model_fields_set:
            _dict['id'] = None

        # set to None if price (nullable) is None
        # and model_fields_set contains the field
        if self.price is None and "price" in self.model_fields_set:
            _dict['price'] = None

        return _dict

    @classmethod
    def from_dict(cls, obj: Optional[Dict[str, Any]]) -> Optional[Self]:
        """Create an instance of PalGacha from a dict"""
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return cls.model_validate(obj)

        _obj = cls.model_validate({
            "banner_collection_url": obj.get("banner_collection_url"),
            "banner_url": obj.get("banner_url"),
            "detail_url": obj.get("detail_url"),
            "gacha_type": obj.get("gacha_type"),
            "id": obj.get("id"),
            "price": obj.get("price")
        })
        return _obj


