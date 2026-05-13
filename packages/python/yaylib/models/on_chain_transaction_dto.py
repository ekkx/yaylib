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
from yaylib.models.asset_info_dto import AssetInfoDTO
from typing import Optional, Set
from typing_extensions import Self

class OnChainTransactionDTO(BaseModel):
    """
    OnChainTransactionDTO
    """ # noqa: E501
    amount: Optional[StrictStr] = None
    asset_info: Optional[AssetInfoDTO] = None
    created_at: Optional[StrictInt] = None
    currency: Optional[StrictStr] = None
    description: Optional[StrictStr] = None
    from_address: Optional[StrictStr] = None
    hash: Optional[StrictStr] = None
    status: Optional[StrictStr] = None
    to_address: Optional[StrictStr] = None
    token_contract: Optional[StrictStr] = None
    token_type: Optional[StrictStr] = None
    __properties: ClassVar[List[str]] = ["amount", "asset_info", "created_at", "currency", "description", "from_address", "hash", "status", "to_address", "token_contract", "token_type"]

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
        """Create an instance of OnChainTransactionDTO from a JSON string"""
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
        # override the default output from pydantic by calling `to_dict()` of asset_info
        if self.asset_info:
            _dict['asset_info'] = self.asset_info.to_dict()
        # set to None if amount (nullable) is None
        # and model_fields_set contains the field
        if self.amount is None and "amount" in self.model_fields_set:
            _dict['amount'] = None

        # set to None if asset_info (nullable) is None
        # and model_fields_set contains the field
        if self.asset_info is None and "asset_info" in self.model_fields_set:
            _dict['asset_info'] = None

        # set to None if created_at (nullable) is None
        # and model_fields_set contains the field
        if self.created_at is None and "created_at" in self.model_fields_set:
            _dict['created_at'] = None

        # set to None if currency (nullable) is None
        # and model_fields_set contains the field
        if self.currency is None and "currency" in self.model_fields_set:
            _dict['currency'] = None

        # set to None if description (nullable) is None
        # and model_fields_set contains the field
        if self.description is None and "description" in self.model_fields_set:
            _dict['description'] = None

        # set to None if from_address (nullable) is None
        # and model_fields_set contains the field
        if self.from_address is None and "from_address" in self.model_fields_set:
            _dict['from_address'] = None

        # set to None if hash (nullable) is None
        # and model_fields_set contains the field
        if self.hash is None and "hash" in self.model_fields_set:
            _dict['hash'] = None

        # set to None if status (nullable) is None
        # and model_fields_set contains the field
        if self.status is None and "status" in self.model_fields_set:
            _dict['status'] = None

        # set to None if to_address (nullable) is None
        # and model_fields_set contains the field
        if self.to_address is None and "to_address" in self.model_fields_set:
            _dict['to_address'] = None

        # set to None if token_contract (nullable) is None
        # and model_fields_set contains the field
        if self.token_contract is None and "token_contract" in self.model_fields_set:
            _dict['token_contract'] = None

        # set to None if token_type (nullable) is None
        # and model_fields_set contains the field
        if self.token_type is None and "token_type" in self.model_fields_set:
            _dict['token_type'] = None

        return _dict

    @classmethod
    def from_dict(cls, obj: Optional[Dict[str, Any]]) -> Optional[Self]:
        """Create an instance of OnChainTransactionDTO from a dict"""
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return cls.model_validate(obj)

        _obj = cls.model_validate({
            "amount": obj.get("amount"),
            "asset_info": AssetInfoDTO.from_dict(obj["asset_info"]) if obj.get("asset_info") is not None else None,
            "created_at": obj.get("created_at"),
            "currency": obj.get("currency"),
            "description": obj.get("description"),
            "from_address": obj.get("from_address"),
            "hash": obj.get("hash"),
            "status": obj.get("status"),
            "to_address": obj.get("to_address"),
            "token_contract": obj.get("token_contract"),
            "token_type": obj.get("token_type")
        })
        return _obj


