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

class Web3EmplWithdrawalTokenDTO(BaseModel):
    """
    Web3EmplWithdrawalTokenDTO
    """ # noqa: E501
    chain_id: Optional[StrictStr] = None
    id: Optional[StrictStr] = None
    image_url: Optional[StrictStr] = None
    token_address: Optional[StrictStr] = None
    token_name: Optional[StrictStr] = None
    token_symbol: Optional[StrictStr] = None
    velo_pool_version: Optional[StrictStr] = None
    __properties: ClassVar[List[str]] = ["chain_id", "id", "image_url", "token_address", "token_name", "token_symbol", "velo_pool_version"]

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
        """Create an instance of Web3EmplWithdrawalTokenDTO from a JSON string"""
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

        # set to None if id (nullable) is None
        # and model_fields_set contains the field
        if self.id is None and "id" in self.model_fields_set:
            _dict['id'] = None

        # set to None if image_url (nullable) is None
        # and model_fields_set contains the field
        if self.image_url is None and "image_url" in self.model_fields_set:
            _dict['image_url'] = None

        # set to None if token_address (nullable) is None
        # and model_fields_set contains the field
        if self.token_address is None and "token_address" in self.model_fields_set:
            _dict['token_address'] = None

        # set to None if token_name (nullable) is None
        # and model_fields_set contains the field
        if self.token_name is None and "token_name" in self.model_fields_set:
            _dict['token_name'] = None

        # set to None if token_symbol (nullable) is None
        # and model_fields_set contains the field
        if self.token_symbol is None and "token_symbol" in self.model_fields_set:
            _dict['token_symbol'] = None

        # set to None if velo_pool_version (nullable) is None
        # and model_fields_set contains the field
        if self.velo_pool_version is None and "velo_pool_version" in self.model_fields_set:
            _dict['velo_pool_version'] = None

        return _dict

    @classmethod
    def from_dict(cls, obj: Optional[Dict[str, Any]]) -> Optional[Self]:
        """Create an instance of Web3EmplWithdrawalTokenDTO from a dict"""
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return cls.model_validate(obj)

        _obj = cls.model_validate({
            "chain_id": obj.get("chain_id"),
            "id": obj.get("id"),
            "image_url": obj.get("image_url"),
            "token_address": obj.get("token_address"),
            "token_name": obj.get("token_name"),
            "token_symbol": obj.get("token_symbol"),
            "velo_pool_version": obj.get("velo_pool_version")
        })
        return _obj


