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

class PalGachaHistoryResponseItem(BaseModel):
    """
    PalGachaHistoryResponseItem
    """ # noqa: E501
    created_at: Optional[StrictInt] = None
    gacha_type: Optional[StrictStr] = None
    id: Optional[StrictStr] = None
    pal_image_url: Optional[StrictStr] = None
    pal_name: Optional[StrictStr] = None
    pal_state: Optional[StrictStr] = None
    pal_token_id: Optional[StrictStr] = None
    price: Optional[StrictInt] = None
    rarity: Optional[StrictStr] = None
    __properties: ClassVar[List[str]] = ["created_at", "gacha_type", "id", "pal_image_url", "pal_name", "pal_state", "pal_token_id", "price", "rarity"]

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
        """Create an instance of PalGachaHistoryResponseItem from a JSON string"""
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
        # set to None if created_at (nullable) is None
        # and model_fields_set contains the field
        if self.created_at is None and "created_at" in self.model_fields_set:
            _dict['created_at'] = None

        # set to None if gacha_type (nullable) is None
        # and model_fields_set contains the field
        if self.gacha_type is None and "gacha_type" in self.model_fields_set:
            _dict['gacha_type'] = None

        # set to None if id (nullable) is None
        # and model_fields_set contains the field
        if self.id is None and "id" in self.model_fields_set:
            _dict['id'] = None

        # set to None if pal_image_url (nullable) is None
        # and model_fields_set contains the field
        if self.pal_image_url is None and "pal_image_url" in self.model_fields_set:
            _dict['pal_image_url'] = None

        # set to None if pal_name (nullable) is None
        # and model_fields_set contains the field
        if self.pal_name is None and "pal_name" in self.model_fields_set:
            _dict['pal_name'] = None

        # set to None if pal_state (nullable) is None
        # and model_fields_set contains the field
        if self.pal_state is None and "pal_state" in self.model_fields_set:
            _dict['pal_state'] = None

        # set to None if pal_token_id (nullable) is None
        # and model_fields_set contains the field
        if self.pal_token_id is None and "pal_token_id" in self.model_fields_set:
            _dict['pal_token_id'] = None

        # set to None if price (nullable) is None
        # and model_fields_set contains the field
        if self.price is None and "price" in self.model_fields_set:
            _dict['price'] = None

        # set to None if rarity (nullable) is None
        # and model_fields_set contains the field
        if self.rarity is None and "rarity" in self.model_fields_set:
            _dict['rarity'] = None

        return _dict

    @classmethod
    def from_dict(cls, obj: Optional[Dict[str, Any]]) -> Optional[Self]:
        """Create an instance of PalGachaHistoryResponseItem from a dict"""
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return cls.model_validate(obj)

        _obj = cls.model_validate({
            "created_at": obj.get("created_at"),
            "gacha_type": obj.get("gacha_type"),
            "id": obj.get("id"),
            "pal_image_url": obj.get("pal_image_url"),
            "pal_name": obj.get("pal_name"),
            "pal_state": obj.get("pal_state"),
            "pal_token_id": obj.get("pal_token_id"),
            "price": obj.get("price"),
            "rarity": obj.get("rarity")
        })
        return _obj


