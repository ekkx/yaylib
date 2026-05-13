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
from yaylib.models.addition_gas_percent_dto import AdditionGasPercentDTO
from yaylib.models.nft_collection_dto import NftCollectionDTO
from yaylib.models.token_dto import TokenDTO
from typing import Optional, Set
from typing_extensions import Self

class BlockchainDTO(BaseModel):
    """
    BlockchainDTO
    """ # noqa: E501
    addition_gas_percent: Optional[AdditionGasPercentDTO] = None
    block_explorer: Optional[StrictStr] = None
    chain_id: Optional[StrictInt] = None
    chain_name: Optional[StrictStr] = None
    image_url: Optional[StrictStr] = None
    nft_collections: Optional[List[NftCollectionDTO]] = None
    priority: Optional[StrictInt] = None
    rpc_url: Optional[StrictStr] = None
    tokens: Optional[List[TokenDTO]] = None
    __properties: ClassVar[List[str]] = ["addition_gas_percent", "block_explorer", "chain_id", "chain_name", "image_url", "nft_collections", "priority", "rpc_url", "tokens"]

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
        """Create an instance of BlockchainDTO from a JSON string"""
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
        # override the default output from pydantic by calling `to_dict()` of addition_gas_percent
        if self.addition_gas_percent:
            _dict['addition_gas_percent'] = self.addition_gas_percent.to_dict()
        # override the default output from pydantic by calling `to_dict()` of each item in nft_collections (list)
        _items = []
        if self.nft_collections:
            for _item_nft_collections in self.nft_collections:
                if _item_nft_collections:
                    _items.append(_item_nft_collections.to_dict())
            _dict['nft_collections'] = _items
        # override the default output from pydantic by calling `to_dict()` of each item in tokens (list)
        _items = []
        if self.tokens:
            for _item_tokens in self.tokens:
                if _item_tokens:
                    _items.append(_item_tokens.to_dict())
            _dict['tokens'] = _items
        # set to None if addition_gas_percent (nullable) is None
        # and model_fields_set contains the field
        if self.addition_gas_percent is None and "addition_gas_percent" in self.model_fields_set:
            _dict['addition_gas_percent'] = None

        # set to None if block_explorer (nullable) is None
        # and model_fields_set contains the field
        if self.block_explorer is None and "block_explorer" in self.model_fields_set:
            _dict['block_explorer'] = None

        # set to None if chain_id (nullable) is None
        # and model_fields_set contains the field
        if self.chain_id is None and "chain_id" in self.model_fields_set:
            _dict['chain_id'] = None

        # set to None if chain_name (nullable) is None
        # and model_fields_set contains the field
        if self.chain_name is None and "chain_name" in self.model_fields_set:
            _dict['chain_name'] = None

        # set to None if image_url (nullable) is None
        # and model_fields_set contains the field
        if self.image_url is None and "image_url" in self.model_fields_set:
            _dict['image_url'] = None

        # set to None if nft_collections (nullable) is None
        # and model_fields_set contains the field
        if self.nft_collections is None and "nft_collections" in self.model_fields_set:
            _dict['nft_collections'] = None

        # set to None if priority (nullable) is None
        # and model_fields_set contains the field
        if self.priority is None and "priority" in self.model_fields_set:
            _dict['priority'] = None

        # set to None if rpc_url (nullable) is None
        # and model_fields_set contains the field
        if self.rpc_url is None and "rpc_url" in self.model_fields_set:
            _dict['rpc_url'] = None

        # set to None if tokens (nullable) is None
        # and model_fields_set contains the field
        if self.tokens is None and "tokens" in self.model_fields_set:
            _dict['tokens'] = None

        return _dict

    @classmethod
    def from_dict(cls, obj: Optional[Dict[str, Any]]) -> Optional[Self]:
        """Create an instance of BlockchainDTO from a dict"""
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return cls.model_validate(obj)

        _obj = cls.model_validate({
            "addition_gas_percent": AdditionGasPercentDTO.from_dict(obj["addition_gas_percent"]) if obj.get("addition_gas_percent") is not None else None,
            "block_explorer": obj.get("block_explorer"),
            "chain_id": obj.get("chain_id"),
            "chain_name": obj.get("chain_name"),
            "image_url": obj.get("image_url"),
            "nft_collections": [NftCollectionDTO.from_dict(_item) for _item in obj["nft_collections"]] if obj.get("nft_collections") is not None else None,
            "priority": obj.get("priority"),
            "rpc_url": obj.get("rpc_url"),
            "tokens": [TokenDTO.from_dict(_item) for _item in obj["tokens"]] if obj.get("tokens") is not None else None
        })
        return _obj


