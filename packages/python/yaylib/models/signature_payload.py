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

class SignaturePayload(BaseModel):
    """
    SignaturePayload
    """ # noqa: E501
    action: Optional[StrictStr] = None
    call_id: Optional[StrictInt] = None
    receiver_uuid: Optional[StrictStr] = None
    sender_uuid: Optional[StrictStr] = None
    signature: Optional[StrictStr] = None
    timestamp: Optional[StrictInt] = None
    __properties: ClassVar[List[str]] = ["action", "call_id", "receiver_uuid", "sender_uuid", "signature", "timestamp"]

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
        """Create an instance of SignaturePayload from a JSON string"""
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
        # set to None if action (nullable) is None
        # and model_fields_set contains the field
        if self.action is None and "action" in self.model_fields_set:
            _dict['action'] = None

        # set to None if call_id (nullable) is None
        # and model_fields_set contains the field
        if self.call_id is None and "call_id" in self.model_fields_set:
            _dict['call_id'] = None

        # set to None if receiver_uuid (nullable) is None
        # and model_fields_set contains the field
        if self.receiver_uuid is None and "receiver_uuid" in self.model_fields_set:
            _dict['receiver_uuid'] = None

        # set to None if sender_uuid (nullable) is None
        # and model_fields_set contains the field
        if self.sender_uuid is None and "sender_uuid" in self.model_fields_set:
            _dict['sender_uuid'] = None

        # set to None if signature (nullable) is None
        # and model_fields_set contains the field
        if self.signature is None and "signature" in self.model_fields_set:
            _dict['signature'] = None

        # set to None if timestamp (nullable) is None
        # and model_fields_set contains the field
        if self.timestamp is None and "timestamp" in self.model_fields_set:
            _dict['timestamp'] = None

        return _dict

    @classmethod
    def from_dict(cls, obj: Optional[Dict[str, Any]]) -> Optional[Self]:
        """Create an instance of SignaturePayload from a dict"""
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return cls.model_validate(obj)

        _obj = cls.model_validate({
            "action": obj.get("action"),
            "call_id": obj.get("call_id"),
            "receiver_uuid": obj.get("receiver_uuid"),
            "sender_uuid": obj.get("sender_uuid"),
            "signature": obj.get("signature"),
            "timestamp": obj.get("timestamp")
        })
        return _obj


