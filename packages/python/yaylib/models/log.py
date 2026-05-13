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

from pydantic import BaseModel, ConfigDict, StrictBool, StrictStr
from typing import Any, ClassVar, Dict, List, Optional
from typing import Optional, Set
from typing_extensions import Self

class Log(BaseModel):
    """
    Log
    """ # noqa: E501
    address: Optional[StrictStr] = None
    block_hash: Optional[StrictStr] = None
    block_number: Optional[StrictStr] = None
    data: Optional[StrictStr] = None
    log_index: Optional[StrictStr] = None
    removed: Optional[StrictBool] = None
    topics: Optional[List[StrictStr]] = None
    transaction_hash: Optional[StrictStr] = None
    transaction_index: Optional[StrictStr] = None
    __properties: ClassVar[List[str]] = ["address", "block_hash", "block_number", "data", "log_index", "removed", "topics", "transaction_hash", "transaction_index"]

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
        """Create an instance of Log from a JSON string"""
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
        # set to None if address (nullable) is None
        # and model_fields_set contains the field
        if self.address is None and "address" in self.model_fields_set:
            _dict['address'] = None

        # set to None if block_hash (nullable) is None
        # and model_fields_set contains the field
        if self.block_hash is None and "block_hash" in self.model_fields_set:
            _dict['block_hash'] = None

        # set to None if block_number (nullable) is None
        # and model_fields_set contains the field
        if self.block_number is None and "block_number" in self.model_fields_set:
            _dict['block_number'] = None

        # set to None if data (nullable) is None
        # and model_fields_set contains the field
        if self.data is None and "data" in self.model_fields_set:
            _dict['data'] = None

        # set to None if log_index (nullable) is None
        # and model_fields_set contains the field
        if self.log_index is None and "log_index" in self.model_fields_set:
            _dict['log_index'] = None

        # set to None if removed (nullable) is None
        # and model_fields_set contains the field
        if self.removed is None and "removed" in self.model_fields_set:
            _dict['removed'] = None

        # set to None if topics (nullable) is None
        # and model_fields_set contains the field
        if self.topics is None and "topics" in self.model_fields_set:
            _dict['topics'] = None

        # set to None if transaction_hash (nullable) is None
        # and model_fields_set contains the field
        if self.transaction_hash is None and "transaction_hash" in self.model_fields_set:
            _dict['transaction_hash'] = None

        # set to None if transaction_index (nullable) is None
        # and model_fields_set contains the field
        if self.transaction_index is None and "transaction_index" in self.model_fields_set:
            _dict['transaction_index'] = None

        return _dict

    @classmethod
    def from_dict(cls, obj: Optional[Dict[str, Any]]) -> Optional[Self]:
        """Create an instance of Log from a dict"""
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return cls.model_validate(obj)

        _obj = cls.model_validate({
            "address": obj.get("address"),
            "block_hash": obj.get("block_hash"),
            "block_number": obj.get("block_number"),
            "data": obj.get("data"),
            "log_index": obj.get("log_index"),
            "removed": obj.get("removed"),
            "topics": obj.get("topics"),
            "transaction_hash": obj.get("transaction_hash"),
            "transaction_index": obj.get("transaction_index")
        })
        return _obj


