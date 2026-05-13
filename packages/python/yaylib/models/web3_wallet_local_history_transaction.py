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

from pydantic import BaseModel, ConfigDict, StrictFloat, StrictInt, StrictStr
from typing import Any, ClassVar, Dict, List, Optional, Union
from typing import Optional, Set
from typing_extensions import Self

class Web3WalletLocalHistoryTransaction(BaseModel):
    """
    Web3WalletLocalHistoryTransaction
    """ # noqa: E501
    amount: Optional[Union[StrictFloat, StrictInt]] = None
    asset_type_name: Optional[StrictStr] = None
    chain: Optional[StrictStr] = None
    chain_id: Optional[StrictInt] = None
    chain_url: Optional[StrictStr] = None
    contract_address: Optional[StrictStr] = None
    gas_limit: Optional[StrictInt] = None
    gas_percentage: Optional[StrictInt] = None
    gas_price: Optional[Union[StrictFloat, StrictInt]] = None
    inputs_hex: Optional[StrictStr] = None
    nonce: Optional[StrictInt] = None
    pal_id: Optional[StrictStr] = None
    timestamp: Optional[StrictInt] = None
    transaction_hash_hex: Optional[StrictStr] = None
    transaction_history_status_name: Optional[StrictStr] = None
    transfer_mode_key: Optional[StrictStr] = None
    tx_unique_id: Optional[StrictStr] = None
    __properties: ClassVar[List[str]] = ["amount", "asset_type_name", "chain", "chain_id", "chain_url", "contract_address", "gas_limit", "gas_percentage", "gas_price", "inputs_hex", "nonce", "pal_id", "timestamp", "transaction_hash_hex", "transaction_history_status_name", "transfer_mode_key", "tx_unique_id"]

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
        """Create an instance of Web3WalletLocalHistoryTransaction from a JSON string"""
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
        # set to None if amount (nullable) is None
        # and model_fields_set contains the field
        if self.amount is None and "amount" in self.model_fields_set:
            _dict['amount'] = None

        # set to None if asset_type_name (nullable) is None
        # and model_fields_set contains the field
        if self.asset_type_name is None and "asset_type_name" in self.model_fields_set:
            _dict['asset_type_name'] = None

        # set to None if chain (nullable) is None
        # and model_fields_set contains the field
        if self.chain is None and "chain" in self.model_fields_set:
            _dict['chain'] = None

        # set to None if chain_id (nullable) is None
        # and model_fields_set contains the field
        if self.chain_id is None and "chain_id" in self.model_fields_set:
            _dict['chain_id'] = None

        # set to None if chain_url (nullable) is None
        # and model_fields_set contains the field
        if self.chain_url is None and "chain_url" in self.model_fields_set:
            _dict['chain_url'] = None

        # set to None if contract_address (nullable) is None
        # and model_fields_set contains the field
        if self.contract_address is None and "contract_address" in self.model_fields_set:
            _dict['contract_address'] = None

        # set to None if gas_limit (nullable) is None
        # and model_fields_set contains the field
        if self.gas_limit is None and "gas_limit" in self.model_fields_set:
            _dict['gas_limit'] = None

        # set to None if gas_percentage (nullable) is None
        # and model_fields_set contains the field
        if self.gas_percentage is None and "gas_percentage" in self.model_fields_set:
            _dict['gas_percentage'] = None

        # set to None if gas_price (nullable) is None
        # and model_fields_set contains the field
        if self.gas_price is None and "gas_price" in self.model_fields_set:
            _dict['gas_price'] = None

        # set to None if inputs_hex (nullable) is None
        # and model_fields_set contains the field
        if self.inputs_hex is None and "inputs_hex" in self.model_fields_set:
            _dict['inputs_hex'] = None

        # set to None if nonce (nullable) is None
        # and model_fields_set contains the field
        if self.nonce is None and "nonce" in self.model_fields_set:
            _dict['nonce'] = None

        # set to None if pal_id (nullable) is None
        # and model_fields_set contains the field
        if self.pal_id is None and "pal_id" in self.model_fields_set:
            _dict['pal_id'] = None

        # set to None if timestamp (nullable) is None
        # and model_fields_set contains the field
        if self.timestamp is None and "timestamp" in self.model_fields_set:
            _dict['timestamp'] = None

        # set to None if transaction_hash_hex (nullable) is None
        # and model_fields_set contains the field
        if self.transaction_hash_hex is None and "transaction_hash_hex" in self.model_fields_set:
            _dict['transaction_hash_hex'] = None

        # set to None if transaction_history_status_name (nullable) is None
        # and model_fields_set contains the field
        if self.transaction_history_status_name is None and "transaction_history_status_name" in self.model_fields_set:
            _dict['transaction_history_status_name'] = None

        # set to None if transfer_mode_key (nullable) is None
        # and model_fields_set contains the field
        if self.transfer_mode_key is None and "transfer_mode_key" in self.model_fields_set:
            _dict['transfer_mode_key'] = None

        # set to None if tx_unique_id (nullable) is None
        # and model_fields_set contains the field
        if self.tx_unique_id is None and "tx_unique_id" in self.model_fields_set:
            _dict['tx_unique_id'] = None

        return _dict

    @classmethod
    def from_dict(cls, obj: Optional[Dict[str, Any]]) -> Optional[Self]:
        """Create an instance of Web3WalletLocalHistoryTransaction from a dict"""
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return cls.model_validate(obj)

        _obj = cls.model_validate({
            "amount": obj.get("amount"),
            "asset_type_name": obj.get("asset_type_name"),
            "chain": obj.get("chain"),
            "chain_id": obj.get("chain_id"),
            "chain_url": obj.get("chain_url"),
            "contract_address": obj.get("contract_address"),
            "gas_limit": obj.get("gas_limit"),
            "gas_percentage": obj.get("gas_percentage"),
            "gas_price": obj.get("gas_price"),
            "inputs_hex": obj.get("inputs_hex"),
            "nonce": obj.get("nonce"),
            "pal_id": obj.get("pal_id"),
            "timestamp": obj.get("timestamp"),
            "transaction_hash_hex": obj.get("transaction_hash_hex"),
            "transaction_history_status_name": obj.get("transaction_history_status_name"),
            "transfer_mode_key": obj.get("transfer_mode_key"),
            "tx_unique_id": obj.get("tx_unique_id")
        })
        return _obj


