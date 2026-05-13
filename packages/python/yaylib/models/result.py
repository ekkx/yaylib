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
from yaylib.models.details import Details
from yaylib.models.empl_transaction import EmplTransaction
from typing import Optional, Set
from typing_extensions import Self

class Result(BaseModel):
    """
    Result
    """ # noqa: E501
    amount: Optional[StrictStr] = None
    currency: Optional[StrictStr] = None
    details: Optional[Details] = None
    empl_transaction: Optional[EmplTransaction] = None
    final_status_updated_at: Optional[StrictInt] = None
    id: Optional[StrictInt] = None
    status: Optional[StrictStr] = None
    transaction_code: Optional[StrictStr] = None
    wallet_address: Optional[StrictStr] = None
    __properties: ClassVar[List[str]] = ["amount", "currency", "details", "empl_transaction", "final_status_updated_at", "id", "status", "transaction_code", "wallet_address"]

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
        """Create an instance of Result from a JSON string"""
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
        # override the default output from pydantic by calling `to_dict()` of details
        if self.details:
            _dict['details'] = self.details.to_dict()
        # override the default output from pydantic by calling `to_dict()` of empl_transaction
        if self.empl_transaction:
            _dict['empl_transaction'] = self.empl_transaction.to_dict()
        # set to None if amount (nullable) is None
        # and model_fields_set contains the field
        if self.amount is None and "amount" in self.model_fields_set:
            _dict['amount'] = None

        # set to None if currency (nullable) is None
        # and model_fields_set contains the field
        if self.currency is None and "currency" in self.model_fields_set:
            _dict['currency'] = None

        # set to None if details (nullable) is None
        # and model_fields_set contains the field
        if self.details is None and "details" in self.model_fields_set:
            _dict['details'] = None

        # set to None if empl_transaction (nullable) is None
        # and model_fields_set contains the field
        if self.empl_transaction is None and "empl_transaction" in self.model_fields_set:
            _dict['empl_transaction'] = None

        # set to None if final_status_updated_at (nullable) is None
        # and model_fields_set contains the field
        if self.final_status_updated_at is None and "final_status_updated_at" in self.model_fields_set:
            _dict['final_status_updated_at'] = None

        # set to None if id (nullable) is None
        # and model_fields_set contains the field
        if self.id is None and "id" in self.model_fields_set:
            _dict['id'] = None

        # set to None if status (nullable) is None
        # and model_fields_set contains the field
        if self.status is None and "status" in self.model_fields_set:
            _dict['status'] = None

        # set to None if transaction_code (nullable) is None
        # and model_fields_set contains the field
        if self.transaction_code is None and "transaction_code" in self.model_fields_set:
            _dict['transaction_code'] = None

        # set to None if wallet_address (nullable) is None
        # and model_fields_set contains the field
        if self.wallet_address is None and "wallet_address" in self.model_fields_set:
            _dict['wallet_address'] = None

        return _dict

    @classmethod
    def from_dict(cls, obj: Optional[Dict[str, Any]]) -> Optional[Self]:
        """Create an instance of Result from a dict"""
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return cls.model_validate(obj)

        _obj = cls.model_validate({
            "amount": obj.get("amount"),
            "currency": obj.get("currency"),
            "details": Details.from_dict(obj["details"]) if obj.get("details") is not None else None,
            "empl_transaction": EmplTransaction.from_dict(obj["empl_transaction"]) if obj.get("empl_transaction") is not None else None,
            "final_status_updated_at": obj.get("final_status_updated_at"),
            "id": obj.get("id"),
            "status": obj.get("status"),
            "transaction_code": obj.get("transaction_code"),
            "wallet_address": obj.get("wallet_address")
        })
        return _obj


