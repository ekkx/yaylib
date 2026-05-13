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

class Web3TransactionDetail(BaseModel):
    """
    Web3TransactionDetail
    """ # noqa: E501
    amount: Optional[StrictStr] = None
    chain_id: Optional[StrictInt] = None
    currency: Optional[StrictStr] = None
    empl_amount: Optional[Union[StrictFloat, StrictInt]] = None
    error: Optional[Dict[str, Any]] = None
    final_status_updated_at_millis: Optional[StrictInt] = None
    id: Optional[StrictInt] = None
    max_token_amount: Optional[StrictStr] = None
    min_token_amount: Optional[StrictStr] = None
    status: Optional[Dict[str, Any]] = None
    token_out_address: Optional[StrictStr] = None
    token_out_amount: Optional[StrictStr] = None
    token_symbol: Optional[StrictStr] = None
    transaction_fee: Optional[Union[StrictFloat, StrictInt]] = None
    wallet_address: Optional[StrictStr] = None
    __properties: ClassVar[List[str]] = ["amount", "chain_id", "currency", "empl_amount", "error", "final_status_updated_at_millis", "id", "max_token_amount", "min_token_amount", "status", "token_out_address", "token_out_amount", "token_symbol", "transaction_fee", "wallet_address"]

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
        """Create an instance of Web3TransactionDetail from a JSON string"""
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

        # set to None if chain_id (nullable) is None
        # and model_fields_set contains the field
        if self.chain_id is None and "chain_id" in self.model_fields_set:
            _dict['chain_id'] = None

        # set to None if currency (nullable) is None
        # and model_fields_set contains the field
        if self.currency is None and "currency" in self.model_fields_set:
            _dict['currency'] = None

        # set to None if empl_amount (nullable) is None
        # and model_fields_set contains the field
        if self.empl_amount is None and "empl_amount" in self.model_fields_set:
            _dict['empl_amount'] = None

        # set to None if error (nullable) is None
        # and model_fields_set contains the field
        if self.error is None and "error" in self.model_fields_set:
            _dict['error'] = None

        # set to None if final_status_updated_at_millis (nullable) is None
        # and model_fields_set contains the field
        if self.final_status_updated_at_millis is None and "final_status_updated_at_millis" in self.model_fields_set:
            _dict['final_status_updated_at_millis'] = None

        # set to None if id (nullable) is None
        # and model_fields_set contains the field
        if self.id is None and "id" in self.model_fields_set:
            _dict['id'] = None

        # set to None if max_token_amount (nullable) is None
        # and model_fields_set contains the field
        if self.max_token_amount is None and "max_token_amount" in self.model_fields_set:
            _dict['max_token_amount'] = None

        # set to None if min_token_amount (nullable) is None
        # and model_fields_set contains the field
        if self.min_token_amount is None and "min_token_amount" in self.model_fields_set:
            _dict['min_token_amount'] = None

        # set to None if status (nullable) is None
        # and model_fields_set contains the field
        if self.status is None and "status" in self.model_fields_set:
            _dict['status'] = None

        # set to None if token_out_address (nullable) is None
        # and model_fields_set contains the field
        if self.token_out_address is None and "token_out_address" in self.model_fields_set:
            _dict['token_out_address'] = None

        # set to None if token_out_amount (nullable) is None
        # and model_fields_set contains the field
        if self.token_out_amount is None and "token_out_amount" in self.model_fields_set:
            _dict['token_out_amount'] = None

        # set to None if token_symbol (nullable) is None
        # and model_fields_set contains the field
        if self.token_symbol is None and "token_symbol" in self.model_fields_set:
            _dict['token_symbol'] = None

        # set to None if transaction_fee (nullable) is None
        # and model_fields_set contains the field
        if self.transaction_fee is None and "transaction_fee" in self.model_fields_set:
            _dict['transaction_fee'] = None

        # set to None if wallet_address (nullable) is None
        # and model_fields_set contains the field
        if self.wallet_address is None and "wallet_address" in self.model_fields_set:
            _dict['wallet_address'] = None

        return _dict

    @classmethod
    def from_dict(cls, obj: Optional[Dict[str, Any]]) -> Optional[Self]:
        """Create an instance of Web3TransactionDetail from a dict"""
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return cls.model_validate(obj)

        _obj = cls.model_validate({
            "amount": obj.get("amount"),
            "chain_id": obj.get("chain_id"),
            "currency": obj.get("currency"),
            "empl_amount": obj.get("empl_amount"),
            "error": obj.get("error"),
            "final_status_updated_at_millis": obj.get("final_status_updated_at_millis"),
            "id": obj.get("id"),
            "max_token_amount": obj.get("max_token_amount"),
            "min_token_amount": obj.get("min_token_amount"),
            "status": obj.get("status"),
            "token_out_address": obj.get("token_out_address"),
            "token_out_amount": obj.get("token_out_amount"),
            "token_symbol": obj.get("token_symbol"),
            "transaction_fee": obj.get("transaction_fee"),
            "wallet_address": obj.get("wallet_address")
        })
        return _obj


