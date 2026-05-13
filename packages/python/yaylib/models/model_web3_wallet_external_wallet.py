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
from typing import Optional, Set
from typing_extensions import Self

class ModelWeb3WalletExternalWallet(BaseModel):
    """
    ModelWeb3WalletExternalWallet
    """ # noqa: E501
    eggs_count: Optional[StrictInt] = None
    pals_count: Optional[StrictInt] = None
    wallet_address: Optional[StrictStr] = None
    __properties: ClassVar[List[str]] = ["eggs_count", "pals_count", "wallet_address"]

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
        """Create an instance of ModelWeb3WalletExternalWallet from a JSON string"""
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
        # set to None if eggs_count (nullable) is None
        # and model_fields_set contains the field
        if self.eggs_count is None and "eggs_count" in self.model_fields_set:
            _dict['eggs_count'] = None

        # set to None if pals_count (nullable) is None
        # and model_fields_set contains the field
        if self.pals_count is None and "pals_count" in self.model_fields_set:
            _dict['pals_count'] = None

        # set to None if wallet_address (nullable) is None
        # and model_fields_set contains the field
        if self.wallet_address is None and "wallet_address" in self.model_fields_set:
            _dict['wallet_address'] = None

        return _dict

    @classmethod
    def from_dict(cls, obj: Optional[Dict[str, Any]]) -> Optional[Self]:
        """Create an instance of ModelWeb3WalletExternalWallet from a dict"""
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return cls.model_validate(obj)

        _obj = cls.model_validate({
            "eggs_count": obj.get("eggs_count"),
            "pals_count": obj.get("pals_count"),
            "wallet_address": obj.get("wallet_address")
        })
        return _obj


