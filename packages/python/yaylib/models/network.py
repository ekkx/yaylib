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
from yaylib.models.model_web3_wallet_gas_percent import ModelWeb3WalletGasPercent
from typing import Optional, Set
from typing_extensions import Self

class Network(BaseModel):
    """
    Network
    """ # noqa: E501
    addition_gas_percent: Optional[ModelWeb3WalletGasPercent] = None
    amm_address: Optional[StrictStr] = None
    id: Optional[StrictInt] = None
    name: Optional[StrictStr] = None
    quoter_address: Optional[StrictStr] = None
    url: Optional[StrictStr] = None
    weth_address: Optional[StrictStr] = None
    yay_address: Optional[StrictStr] = None
    __properties: ClassVar[List[str]] = ["addition_gas_percent", "amm_address", "id", "name", "quoter_address", "url", "weth_address", "yay_address"]

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
        """Create an instance of Network from a JSON string"""
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
        # override the default output from pydantic by calling `to_dict()` of addition_gas_percent
        if self.addition_gas_percent:
            _dict['addition_gas_percent'] = self.addition_gas_percent.to_dict()
        # set to None if addition_gas_percent (nullable) is None
        # and model_fields_set contains the field
        if self.addition_gas_percent is None and "addition_gas_percent" in self.model_fields_set:
            _dict['addition_gas_percent'] = None

        # set to None if amm_address (nullable) is None
        # and model_fields_set contains the field
        if self.amm_address is None and "amm_address" in self.model_fields_set:
            _dict['amm_address'] = None

        # set to None if id (nullable) is None
        # and model_fields_set contains the field
        if self.id is None and "id" in self.model_fields_set:
            _dict['id'] = None

        # set to None if name (nullable) is None
        # and model_fields_set contains the field
        if self.name is None and "name" in self.model_fields_set:
            _dict['name'] = None

        # set to None if quoter_address (nullable) is None
        # and model_fields_set contains the field
        if self.quoter_address is None and "quoter_address" in self.model_fields_set:
            _dict['quoter_address'] = None

        # set to None if url (nullable) is None
        # and model_fields_set contains the field
        if self.url is None and "url" in self.model_fields_set:
            _dict['url'] = None

        # set to None if weth_address (nullable) is None
        # and model_fields_set contains the field
        if self.weth_address is None and "weth_address" in self.model_fields_set:
            _dict['weth_address'] = None

        # set to None if yay_address (nullable) is None
        # and model_fields_set contains the field
        if self.yay_address is None and "yay_address" in self.model_fields_set:
            _dict['yay_address'] = None

        return _dict

    @classmethod
    def from_dict(cls, obj: Optional[Dict[str, Any]]) -> Optional[Self]:
        """Create an instance of Network from a dict"""
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return cls.model_validate(obj)

        _obj = cls.model_validate({
            "addition_gas_percent": ModelWeb3WalletGasPercent.from_dict(obj["addition_gas_percent"]) if obj.get("addition_gas_percent") is not None else None,
            "amm_address": obj.get("amm_address"),
            "id": obj.get("id"),
            "name": obj.get("name"),
            "quoter_address": obj.get("quoter_address"),
            "url": obj.get("url"),
            "weth_address": obj.get("weth_address"),
            "yay_address": obj.get("yay_address")
        })
        return _obj


