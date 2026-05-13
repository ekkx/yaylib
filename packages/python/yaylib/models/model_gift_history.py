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

from pydantic import BaseModel, ConfigDict, StrictBool, StrictInt
from typing import Any, ClassVar, Dict, List, Optional
from yaylib.models.gift import Gift
from yaylib.models.received_gift import ReceivedGift
from yaylib.models.user import User
from typing import Optional, Set
from typing_extensions import Self

class ModelGiftHistory(BaseModel):
    """
    ModelGiftHistory
    """ # noqa: E501
    gift_item: Optional[Gift] = None
    gifts: Optional[List[ReceivedGift]] = None
    is_sender_anonymous: Optional[StrictBool] = None
    receiver: Optional[User] = None
    sender: Optional[User] = None
    transaction_at_seconds: Optional[StrictInt] = None
    __properties: ClassVar[List[str]] = ["gift_item", "gifts", "is_sender_anonymous", "receiver", "sender", "transaction_at_seconds"]

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
        """Create an instance of ModelGiftHistory from a JSON string"""
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
        # override the default output from pydantic by calling `to_dict()` of gift_item
        if self.gift_item:
            _dict['gift_item'] = self.gift_item.to_dict()
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
        # override the default output from pydantic by calling `to_dict()` of sender
        if self.sender:
            _dict['sender'] = self.sender.to_dict()
        # set to None if gift_item (nullable) is None
        # and model_fields_set contains the field
        if self.gift_item is None and "gift_item" in self.model_fields_set:
            _dict['gift_item'] = None

        # set to None if gifts (nullable) is None
        # and model_fields_set contains the field
        if self.gifts is None and "gifts" in self.model_fields_set:
            _dict['gifts'] = None

        # set to None if is_sender_anonymous (nullable) is None
        # and model_fields_set contains the field
        if self.is_sender_anonymous is None and "is_sender_anonymous" in self.model_fields_set:
            _dict['is_sender_anonymous'] = None

        # set to None if receiver (nullable) is None
        # and model_fields_set contains the field
        if self.receiver is None and "receiver" in self.model_fields_set:
            _dict['receiver'] = None

        # set to None if sender (nullable) is None
        # and model_fields_set contains the field
        if self.sender is None and "sender" in self.model_fields_set:
            _dict['sender'] = None

        # set to None if transaction_at_seconds (nullable) is None
        # and model_fields_set contains the field
        if self.transaction_at_seconds is None and "transaction_at_seconds" in self.model_fields_set:
            _dict['transaction_at_seconds'] = None

        return _dict

    @classmethod
    def from_dict(cls, obj: Optional[Dict[str, Any]]) -> Optional[Self]:
        """Create an instance of ModelGiftHistory from a dict"""
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return cls.model_validate(obj)

        _obj = cls.model_validate({
            "gift_item": Gift.from_dict(obj["gift_item"]) if obj.get("gift_item") is not None else None,
            "gifts": [ReceivedGift.from_dict(_item) for _item in obj["gifts"]] if obj.get("gifts") is not None else None,
            "is_sender_anonymous": obj.get("is_sender_anonymous"),
            "receiver": User.from_dict(obj["receiver"]) if obj.get("receiver") is not None else None,
            "sender": User.from_dict(obj["sender"]) if obj.get("sender") is not None else None,
            "transaction_at_seconds": obj.get("transaction_at_seconds")
        })
        return _obj


