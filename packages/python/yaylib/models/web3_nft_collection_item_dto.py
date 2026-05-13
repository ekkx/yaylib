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
from yaylib.models.nft_metadata_dto import NFTMetadataDTO
from typing import Optional, Set
from typing_extensions import Self

class Web3NftCollectionItemDTO(BaseModel):
    """
    Web3NftCollectionItemDTO
    """ # noqa: E501
    balance: Optional[StrictStr] = None
    chain_id: Optional[StrictInt] = None
    contract_address: Optional[StrictStr] = None
    extra_metadata: Optional[NFTMetadataDTO] = None
    image_url: Optional[StrictStr] = None
    name: Optional[StrictStr] = None
    token_id: Optional[StrictStr] = None
    token_type: Optional[StrictStr] = None
    __properties: ClassVar[List[str]] = ["balance", "chain_id", "contract_address", "extra_metadata", "image_url", "name", "token_id", "token_type"]

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
        """Create an instance of Web3NftCollectionItemDTO from a JSON string"""
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
        # override the default output from pydantic by calling `to_dict()` of extra_metadata
        if self.extra_metadata:
            _dict['extra_metadata'] = self.extra_metadata.to_dict()
        # set to None if balance (nullable) is None
        # and model_fields_set contains the field
        if self.balance is None and "balance" in self.model_fields_set:
            _dict['balance'] = None

        # set to None if chain_id (nullable) is None
        # and model_fields_set contains the field
        if self.chain_id is None and "chain_id" in self.model_fields_set:
            _dict['chain_id'] = None

        # set to None if contract_address (nullable) is None
        # and model_fields_set contains the field
        if self.contract_address is None and "contract_address" in self.model_fields_set:
            _dict['contract_address'] = None

        # set to None if extra_metadata (nullable) is None
        # and model_fields_set contains the field
        if self.extra_metadata is None and "extra_metadata" in self.model_fields_set:
            _dict['extra_metadata'] = None

        # set to None if image_url (nullable) is None
        # and model_fields_set contains the field
        if self.image_url is None and "image_url" in self.model_fields_set:
            _dict['image_url'] = None

        # set to None if name (nullable) is None
        # and model_fields_set contains the field
        if self.name is None and "name" in self.model_fields_set:
            _dict['name'] = None

        # set to None if token_id (nullable) is None
        # and model_fields_set contains the field
        if self.token_id is None and "token_id" in self.model_fields_set:
            _dict['token_id'] = None

        # set to None if token_type (nullable) is None
        # and model_fields_set contains the field
        if self.token_type is None and "token_type" in self.model_fields_set:
            _dict['token_type'] = None

        return _dict

    @classmethod
    def from_dict(cls, obj: Optional[Dict[str, Any]]) -> Optional[Self]:
        """Create an instance of Web3NftCollectionItemDTO from a dict"""
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return cls.model_validate(obj)

        _obj = cls.model_validate({
            "balance": obj.get("balance"),
            "chain_id": obj.get("chain_id"),
            "contract_address": obj.get("contract_address"),
            "extra_metadata": NFTMetadataDTO.from_dict(obj["extra_metadata"]) if obj.get("extra_metadata") is not None else None,
            "image_url": obj.get("image_url"),
            "name": obj.get("name"),
            "token_id": obj.get("token_id"),
            "token_type": obj.get("token_type")
        })
        return _obj


