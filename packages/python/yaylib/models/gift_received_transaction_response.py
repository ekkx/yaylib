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

from pydantic import BaseModel, ConfigDict, StrictInt
from typing import Any, ClassVar, Dict, List, Optional
from yaylib.models.realm_user import RealmUser
from yaylib.models.transaction_gift_received import TransactionGiftReceived
from typing import Optional, Set
from typing_extensions import Self

class GiftReceivedTransactionResponse(BaseModel):
    """
    GiftReceivedTransactionResponse
    """ # noqa: E501
    created_at: Optional[StrictInt] = None
    gifts: Optional[List[TransactionGiftReceived]] = None
    receiver: Optional[RealmUser] = None
    __properties: ClassVar[List[str]] = ["created_at", "gifts", "receiver"]

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
        """Create an instance of GiftReceivedTransactionResponse from a JSON string"""
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
        # override the default output from pydantic by calling `to_dict()` of each item in gifts (list)
        _items = []
        if self.gifts:
            for _item_gifts in self.gifts:
                if _item_gifts:
                    _items.append(_item_gifts.to_dict())
            _dict['gifts'] = _items
        # override the default output from pydantic by calling `to_dict()` of receiver
        if self.receiver:
            _dict['receiver'] = self.receiver.to_dict()
        # set to None if created_at (nullable) is None
        # and model_fields_set contains the field
        if self.created_at is None and "created_at" in self.model_fields_set:
            _dict['created_at'] = None

        # set to None if gifts (nullable) is None
        # and model_fields_set contains the field
        if self.gifts is None and "gifts" in self.model_fields_set:
            _dict['gifts'] = None

        # set to None if receiver (nullable) is None
        # and model_fields_set contains the field
        if self.receiver is None and "receiver" in self.model_fields_set:
            _dict['receiver'] = None

        return _dict

    @classmethod
    def from_dict(cls, obj: Optional[Dict[str, Any]]) -> Optional[Self]:
        """Create an instance of GiftReceivedTransactionResponse from a dict"""
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return cls.model_validate(obj)

        _obj = cls.model_validate({
            "created_at": obj.get("created_at"),
            "gifts": [TransactionGiftReceived.from_dict(_item) for _item in obj["gifts"]] if obj.get("gifts") is not None else None,
            "receiver": RealmUser.from_dict(obj["receiver"]) if obj.get("receiver") is not None else None
        })
        return _obj


