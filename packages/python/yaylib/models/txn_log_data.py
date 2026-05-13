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

class TXNLogData(BaseModel):
    """
    TXNLogData
    """ # noqa: E501
    contract_address: Optional[StrictStr] = None
    inputs_hex: Optional[StrictStr] = None
    tx_unique_id: Optional[StrictStr] = None
    __properties: ClassVar[List[str]] = ["contract_address", "inputs_hex", "tx_unique_id"]

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
        """Create an instance of TXNLogData from a JSON string"""
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
        # set to None if contract_address (nullable) is None
        # and model_fields_set contains the field
        if self.contract_address is None and "contract_address" in self.model_fields_set:
            _dict['contract_address'] = None

        # set to None if inputs_hex (nullable) is None
        # and model_fields_set contains the field
        if self.inputs_hex is None and "inputs_hex" in self.model_fields_set:
            _dict['inputs_hex'] = None

        # set to None if tx_unique_id (nullable) is None
        # and model_fields_set contains the field
        if self.tx_unique_id is None and "tx_unique_id" in self.model_fields_set:
            _dict['tx_unique_id'] = None

        return _dict

    @classmethod
    def from_dict(cls, obj: Optional[Dict[str, Any]]) -> Optional[Self]:
        """Create an instance of TXNLogData from a dict"""
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return cls.model_validate(obj)

        _obj = cls.model_validate({
            "contract_address": obj.get("contract_address"),
            "inputs_hex": obj.get("inputs_hex"),
            "tx_unique_id": obj.get("tx_unique_id")
        })
        return _obj


