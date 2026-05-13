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
from yaylib.models.transaction_gift_received_item import TransactionGiftReceivedItem
from typing import Optional, Set
from typing_extensions import Self

class TransactionGiftReceived(BaseModel):
    """
    TransactionGiftReceived
    """ # noqa: E501
    icon: Optional[StrictStr] = None
    item: Optional[TransactionGiftReceivedItem] = None
    item_id: Optional[StrictInt] = None
    quantity: Optional[StrictInt] = None
    __properties: ClassVar[List[str]] = ["icon", "item", "item_id", "quantity"]

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
        """Create an instance of TransactionGiftReceived from a JSON string"""
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
        # override the default output from pydantic by calling `to_dict()` of item
        if self.item:
            _dict['item'] = self.item.to_dict()
        # set to None if icon (nullable) is None
        # and model_fields_set contains the field
        if self.icon is None and "icon" in self.model_fields_set:
            _dict['icon'] = None

        # set to None if item (nullable) is None
        # and model_fields_set contains the field
        if self.item is None and "item" in self.model_fields_set:
            _dict['item'] = None

        # set to None if item_id (nullable) is None
        # and model_fields_set contains the field
        if self.item_id is None and "item_id" in self.model_fields_set:
            _dict['item_id'] = None

        # set to None if quantity (nullable) is None
        # and model_fields_set contains the field
        if self.quantity is None and "quantity" in self.model_fields_set:
            _dict['quantity'] = None

        return _dict

    @classmethod
    def from_dict(cls, obj: Optional[Dict[str, Any]]) -> Optional[Self]:
        """Create an instance of TransactionGiftReceived from a dict"""
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return cls.model_validate(obj)

        _obj = cls.model_validate({
            "icon": obj.get("icon"),
            "item": TransactionGiftReceivedItem.from_dict(obj["item"]) if obj.get("item") is not None else None,
            "item_id": obj.get("item_id"),
            "quantity": obj.get("quantity")
        })
        return _obj


