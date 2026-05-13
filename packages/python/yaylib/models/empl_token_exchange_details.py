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

class EmplTokenExchangeDetails(BaseModel):
    """
    EmplTokenExchangeDetails
    """ # noqa: E501
    chain_id: Optional[StrictInt] = None
    max_token_amount: Optional[StrictStr] = None
    min_token_amount: Optional[StrictStr] = None
    token_out_address: Optional[StrictStr] = None
    token_symbol: Optional[StrictStr] = None
    transaction_fee: Optional[Union[StrictFloat, StrictInt]] = None
    __properties: ClassVar[List[str]] = ["chain_id", "max_token_amount", "min_token_amount", "token_out_address", "token_symbol", "transaction_fee"]

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
        """Create an instance of EmplTokenExchangeDetails from a JSON string"""
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
        # set to None if chain_id (nullable) is None
        # and model_fields_set contains the field
        if self.chain_id is None and "chain_id" in self.model_fields_set:
            _dict['chain_id'] = None

        # set to None if max_token_amount (nullable) is None
        # and model_fields_set contains the field
        if self.max_token_amount is None and "max_token_amount" in self.model_fields_set:
            _dict['max_token_amount'] = None

        # set to None if min_token_amount (nullable) is None
        # and model_fields_set contains the field
        if self.min_token_amount is None and "min_token_amount" in self.model_fields_set:
            _dict['min_token_amount'] = None

        # set to None if token_out_address (nullable) is None
        # and model_fields_set contains the field
        if self.token_out_address is None and "token_out_address" in self.model_fields_set:
            _dict['token_out_address'] = None

        # set to None if token_symbol (nullable) is None
        # and model_fields_set contains the field
        if self.token_symbol is None and "token_symbol" in self.model_fields_set:
            _dict['token_symbol'] = None

        # set to None if transaction_fee (nullable) is None
        # and model_fields_set contains the field
        if self.transaction_fee is None and "transaction_fee" in self.model_fields_set:
            _dict['transaction_fee'] = None

        return _dict

    @classmethod
    def from_dict(cls, obj: Optional[Dict[str, Any]]) -> Optional[Self]:
        """Create an instance of EmplTokenExchangeDetails from a dict"""
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return cls.model_validate(obj)

        _obj = cls.model_validate({
            "chain_id": obj.get("chain_id"),
            "max_token_amount": obj.get("max_token_amount"),
            "min_token_amount": obj.get("min_token_amount"),
            "token_out_address": obj.get("token_out_address"),
            "token_symbol": obj.get("token_symbol"),
            "transaction_fee": obj.get("transaction_fee")
        })
        return _obj


