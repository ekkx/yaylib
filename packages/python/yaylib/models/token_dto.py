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
from yaylib.models.bridge_dto import BridgeDTO
from yaylib.models.price_dto import PriceDTO
from yaylib.models.swap_dto import SwapDTO
from typing import Optional, Set
from typing_extensions import Self

class TokenDTO(BaseModel):
    """
    TokenDTO
    """ # noqa: E501
    balance: Optional[StrictStr] = None
    bridge: Optional[BridgeDTO] = None
    contract_address: Optional[StrictStr] = None
    decimals: Optional[StrictInt] = None
    image_url: Optional[StrictStr] = None
    price: Optional[PriceDTO] = None
    priority: Optional[StrictInt] = None
    status: Optional[StrictStr] = None
    swap: Optional[SwapDTO] = None
    symbol: Optional[StrictStr] = None
    token_type: Optional[StrictStr] = None
    __properties: ClassVar[List[str]] = ["balance", "bridge", "contract_address", "decimals", "image_url", "price", "priority", "status", "swap", "symbol", "token_type"]

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
        """Create an instance of TokenDTO from a JSON string"""
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
        # override the default output from pydantic by calling `to_dict()` of bridge
        if self.bridge:
            _dict['bridge'] = self.bridge.to_dict()
        # override the default output from pydantic by calling `to_dict()` of price
        if self.price:
            _dict['price'] = self.price.to_dict()
        # override the default output from pydantic by calling `to_dict()` of swap
        if self.swap:
            _dict['swap'] = self.swap.to_dict()
        # set to None if balance (nullable) is None
        # and model_fields_set contains the field
        if self.balance is None and "balance" in self.model_fields_set:
            _dict['balance'] = None

        # set to None if bridge (nullable) is None
        # and model_fields_set contains the field
        if self.bridge is None and "bridge" in self.model_fields_set:
            _dict['bridge'] = None

        # set to None if contract_address (nullable) is None
        # and model_fields_set contains the field
        if self.contract_address is None and "contract_address" in self.model_fields_set:
            _dict['contract_address'] = None

        # set to None if decimals (nullable) is None
        # and model_fields_set contains the field
        if self.decimals is None and "decimals" in self.model_fields_set:
            _dict['decimals'] = None

        # set to None if image_url (nullable) is None
        # and model_fields_set contains the field
        if self.image_url is None and "image_url" in self.model_fields_set:
            _dict['image_url'] = None

        # set to None if price (nullable) is None
        # and model_fields_set contains the field
        if self.price is None and "price" in self.model_fields_set:
            _dict['price'] = None

        # set to None if priority (nullable) is None
        # and model_fields_set contains the field
        if self.priority is None and "priority" in self.model_fields_set:
            _dict['priority'] = None

        # set to None if status (nullable) is None
        # and model_fields_set contains the field
        if self.status is None and "status" in self.model_fields_set:
            _dict['status'] = None

        # set to None if swap (nullable) is None
        # and model_fields_set contains the field
        if self.swap is None and "swap" in self.model_fields_set:
            _dict['swap'] = None

        # set to None if symbol (nullable) is None
        # and model_fields_set contains the field
        if self.symbol is None and "symbol" in self.model_fields_set:
            _dict['symbol'] = None

        # set to None if token_type (nullable) is None
        # and model_fields_set contains the field
        if self.token_type is None and "token_type" in self.model_fields_set:
            _dict['token_type'] = None

        return _dict

    @classmethod
    def from_dict(cls, obj: Optional[Dict[str, Any]]) -> Optional[Self]:
        """Create an instance of TokenDTO from a dict"""
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return cls.model_validate(obj)

        _obj = cls.model_validate({
            "balance": obj.get("balance"),
            "bridge": BridgeDTO.from_dict(obj["bridge"]) if obj.get("bridge") is not None else None,
            "contract_address": obj.get("contract_address"),
            "decimals": obj.get("decimals"),
            "image_url": obj.get("image_url"),
            "price": PriceDTO.from_dict(obj["price"]) if obj.get("price") is not None else None,
            "priority": obj.get("priority"),
            "status": obj.get("status"),
            "swap": SwapDTO.from_dict(obj["swap"]) if obj.get("swap") is not None else None,
            "symbol": obj.get("symbol"),
            "token_type": obj.get("token_type")
        })
        return _obj


