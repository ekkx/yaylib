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
from typing import Optional, Set
from typing_extensions import Self

class Web3EmplTokenExchangeConversionAndFeeDTO(BaseModel):
    """
    Web3EmplTokenExchangeConversionAndFeeDTO
    """ # noqa: E501
    empl_amount: Optional[StrictStr] = None
    empl_tx_fee: Optional[StrictStr] = None
    max_amount_out: Optional[StrictStr] = None
    min_amount_out: Optional[StrictStr] = None
    token_out_address: Optional[StrictStr] = None
    __properties: ClassVar[List[str]] = ["empl_amount", "empl_tx_fee", "max_amount_out", "min_amount_out", "token_out_address"]

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
        """Create an instance of Web3EmplTokenExchangeConversionAndFeeDTO from a JSON string"""
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
        # set to None if empl_amount (nullable) is None
        # and model_fields_set contains the field
        if self.empl_amount is None and "empl_amount" in self.model_fields_set:
            _dict['empl_amount'] = None

        # set to None if empl_tx_fee (nullable) is None
        # and model_fields_set contains the field
        if self.empl_tx_fee is None and "empl_tx_fee" in self.model_fields_set:
            _dict['empl_tx_fee'] = None

        # set to None if max_amount_out (nullable) is None
        # and model_fields_set contains the field
        if self.max_amount_out is None and "max_amount_out" in self.model_fields_set:
            _dict['max_amount_out'] = None

        # set to None if min_amount_out (nullable) is None
        # and model_fields_set contains the field
        if self.min_amount_out is None and "min_amount_out" in self.model_fields_set:
            _dict['min_amount_out'] = None

        # set to None if token_out_address (nullable) is None
        # and model_fields_set contains the field
        if self.token_out_address is None and "token_out_address" in self.model_fields_set:
            _dict['token_out_address'] = None

        return _dict

    @classmethod
    def from_dict(cls, obj: Optional[Dict[str, Any]]) -> Optional[Self]:
        """Create an instance of Web3EmplTokenExchangeConversionAndFeeDTO from a dict"""
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return cls.model_validate(obj)

        _obj = cls.model_validate({
            "empl_amount": obj.get("empl_amount"),
            "empl_tx_fee": obj.get("empl_tx_fee"),
            "max_amount_out": obj.get("max_amount_out"),
            "min_amount_out": obj.get("min_amount_out"),
            "token_out_address": obj.get("token_out_address")
        })
        return _obj


