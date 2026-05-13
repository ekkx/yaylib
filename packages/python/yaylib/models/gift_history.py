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
from yaylib.models.gift_count import GiftCount
from yaylib.models.gift_slug_item import GiftSlugItem
from yaylib.models.realm_user import RealmUser
from typing import Optional, Set
from typing_extensions import Self

class GiftHistory(BaseModel):
    """
    GiftHistory
    """ # noqa: E501
    anonymous: Optional[StrictBool] = None
    gift_item: Optional[GiftSlugItem] = None
    gifts_count: Optional[List[GiftCount]] = None
    receiver: Optional[RealmUser] = None
    sender: Optional[RealmUser] = None
    sent_at: Optional[StrictInt] = None
    __properties: ClassVar[List[str]] = ["anonymous", "gift_item", "gifts_count", "receiver", "sender", "sent_at"]

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
        """Create an instance of GiftHistory from a JSON string"""
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
        # override the default output from pydantic by calling `to_dict()` of each item in gifts_count (list)
        _items = []
        if self.gifts_count:
            for _item_gifts_count in self.gifts_count:
                if _item_gifts_count:
                    _items.append(_item_gifts_count.to_dict())
            _dict['gifts_count'] = _items
        # override the default output from pydantic by calling `to_dict()` of receiver
        if self.receiver:
            _dict['receiver'] = self.receiver.to_dict()
        # override the default output from pydantic by calling `to_dict()` of sender
        if self.sender:
            _dict['sender'] = self.sender.to_dict()
        # set to None if anonymous (nullable) is None
        # and model_fields_set contains the field
        if self.anonymous is None and "anonymous" in self.model_fields_set:
            _dict['anonymous'] = None

        # set to None if gift_item (nullable) is None
        # and model_fields_set contains the field
        if self.gift_item is None and "gift_item" in self.model_fields_set:
            _dict['gift_item'] = None

        # set to None if gifts_count (nullable) is None
        # and model_fields_set contains the field
        if self.gifts_count is None and "gifts_count" in self.model_fields_set:
            _dict['gifts_count'] = None

        # set to None if receiver (nullable) is None
        # and model_fields_set contains the field
        if self.receiver is None and "receiver" in self.model_fields_set:
            _dict['receiver'] = None

        # set to None if sender (nullable) is None
        # and model_fields_set contains the field
        if self.sender is None and "sender" in self.model_fields_set:
            _dict['sender'] = None

        # set to None if sent_at (nullable) is None
        # and model_fields_set contains the field
        if self.sent_at is None and "sent_at" in self.model_fields_set:
            _dict['sent_at'] = None

        return _dict

    @classmethod
    def from_dict(cls, obj: Optional[Dict[str, Any]]) -> Optional[Self]:
        """Create an instance of GiftHistory from a dict"""
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return cls.model_validate(obj)

        _obj = cls.model_validate({
            "anonymous": obj.get("anonymous"),
            "gift_item": GiftSlugItem.from_dict(obj["gift_item"]) if obj.get("gift_item") is not None else None,
            "gifts_count": [GiftCount.from_dict(_item) for _item in obj["gifts_count"]] if obj.get("gifts_count") is not None else None,
            "receiver": RealmUser.from_dict(obj["receiver"]) if obj.get("receiver") is not None else None,
            "sender": RealmUser.from_dict(obj["sender"]) if obj.get("sender") is not None else None,
            "sent_at": obj.get("sent_at")
        })
        return _obj


