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
from yaylib.models.log import Log
from typing import Optional, Set
from typing_extensions import Self

class Web3WalletReceipt(BaseModel):
    """
    Web3WalletReceipt
    """ # noqa: E501
    block_hash: Optional[StrictStr] = None
    block_number: Optional[StrictStr] = None
    contract_address: Optional[StrictStr] = None
    cumulative_gas_used: Optional[StrictStr] = None
    gas_used: Optional[StrictStr] = None
    logs: Optional[List[Log]] = None
    logs_bloom: Optional[StrictStr] = None
    root: Optional[StrictStr] = None
    status: Optional[StrictStr] = None
    transaction_hash: Optional[StrictStr] = None
    transaction_index: Optional[StrictStr] = None
    __properties: ClassVar[List[str]] = ["block_hash", "block_number", "contract_address", "cumulative_gas_used", "gas_used", "logs", "logs_bloom", "root", "status", "transaction_hash", "transaction_index"]

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
        """Create an instance of Web3WalletReceipt from a JSON string"""
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
        # override the default output from pydantic by calling `to_dict()` of each item in logs (list)
        _items = []
        if self.logs:
            for _item_logs in self.logs:
                if _item_logs:
                    _items.append(_item_logs.to_dict())
            _dict['logs'] = _items
        # set to None if block_hash (nullable) is None
        # and model_fields_set contains the field
        if self.block_hash is None and "block_hash" in self.model_fields_set:
            _dict['block_hash'] = None

        # set to None if block_number (nullable) is None
        # and model_fields_set contains the field
        if self.block_number is None and "block_number" in self.model_fields_set:
            _dict['block_number'] = None

        # set to None if contract_address (nullable) is None
        # and model_fields_set contains the field
        if self.contract_address is None and "contract_address" in self.model_fields_set:
            _dict['contract_address'] = None

        # set to None if cumulative_gas_used (nullable) is None
        # and model_fields_set contains the field
        if self.cumulative_gas_used is None and "cumulative_gas_used" in self.model_fields_set:
            _dict['cumulative_gas_used'] = None

        # set to None if gas_used (nullable) is None
        # and model_fields_set contains the field
        if self.gas_used is None and "gas_used" in self.model_fields_set:
            _dict['gas_used'] = None

        # set to None if logs (nullable) is None
        # and model_fields_set contains the field
        if self.logs is None and "logs" in self.model_fields_set:
            _dict['logs'] = None

        # set to None if logs_bloom (nullable) is None
        # and model_fields_set contains the field
        if self.logs_bloom is None and "logs_bloom" in self.model_fields_set:
            _dict['logs_bloom'] = None

        # set to None if root (nullable) is None
        # and model_fields_set contains the field
        if self.root is None and "root" in self.model_fields_set:
            _dict['root'] = None

        # set to None if status (nullable) is None
        # and model_fields_set contains the field
        if self.status is None and "status" in self.model_fields_set:
            _dict['status'] = None

        # set to None if transaction_hash (nullable) is None
        # and model_fields_set contains the field
        if self.transaction_hash is None and "transaction_hash" in self.model_fields_set:
            _dict['transaction_hash'] = None

        # set to None if transaction_index (nullable) is None
        # and model_fields_set contains the field
        if self.transaction_index is None and "transaction_index" in self.model_fields_set:
            _dict['transaction_index'] = None

        return _dict

    @classmethod
    def from_dict(cls, obj: Optional[Dict[str, Any]]) -> Optional[Self]:
        """Create an instance of Web3WalletReceipt from a dict"""
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return cls.model_validate(obj)

        _obj = cls.model_validate({
            "block_hash": obj.get("block_hash"),
            "block_number": obj.get("block_number"),
            "contract_address": obj.get("contract_address"),
            "cumulative_gas_used": obj.get("cumulative_gas_used"),
            "gas_used": obj.get("gas_used"),
            "logs": [Log.from_dict(_item) for _item in obj["logs"]] if obj.get("logs") is not None else None,
            "logs_bloom": obj.get("logs_bloom"),
            "root": obj.get("root"),
            "status": obj.get("status"),
            "transaction_hash": obj.get("transaction_hash"),
            "transaction_index": obj.get("transaction_index")
        })
        return _obj


