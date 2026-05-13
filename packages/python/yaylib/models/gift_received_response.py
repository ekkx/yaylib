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
from yaylib.models.received_gift import ReceivedGift
from typing import Optional, Set
from typing_extensions import Self

class GiftReceivedResponse(BaseModel):
    """
    GiftReceivedResponse
    """ # noqa: E501
    next_page_value: Optional[StrictInt] = None
    received_gifts: Optional[List[ReceivedGift]] = None
    total_count: Optional[StrictInt] = None
    __properties: ClassVar[List[str]] = ["next_page_value", "received_gifts", "total_count"]

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
        """Create an instance of GiftReceivedResponse from a JSON string"""
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
        # override the default output from pydantic by calling `to_dict()` of each item in received_gifts (list)
        _items = []
        if self.received_gifts:
            for _item_received_gifts in self.received_gifts:
                if _item_received_gifts:
                    _items.append(_item_received_gifts.to_dict())
            _dict['received_gifts'] = _items
        # set to None if next_page_value (nullable) is None
        # and model_fields_set contains the field
        if self.next_page_value is None and "next_page_value" in self.model_fields_set:
            _dict['next_page_value'] = None

        # set to None if received_gifts (nullable) is None
        # and model_fields_set contains the field
        if self.received_gifts is None and "received_gifts" in self.model_fields_set:
            _dict['received_gifts'] = None

        # set to None if total_count (nullable) is None
        # and model_fields_set contains the field
        if self.total_count is None and "total_count" in self.model_fields_set:
            _dict['total_count'] = None

        return _dict

    @classmethod
    def from_dict(cls, obj: Optional[Dict[str, Any]]) -> Optional[Self]:
        """Create an instance of GiftReceivedResponse from a dict"""
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return cls.model_validate(obj)

        _obj = cls.model_validate({
            "next_page_value": obj.get("next_page_value"),
            "received_gifts": [ReceivedGift.from_dict(_item) for _item in obj["received_gifts"]] if obj.get("received_gifts") is not None else None,
            "total_count": obj.get("total_count")
        })
        return _obj


