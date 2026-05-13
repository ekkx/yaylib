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
from yaylib.models.gift_reward_gift import GiftRewardGift
from typing import Optional, Set
from typing_extensions import Self

class GiftReward(BaseModel):
    """
    GiftReward
    """ # noqa: E501
    empl_reward: Optional[StrictInt] = None
    gifts: Optional[List[GiftRewardGift]] = None
    id: Optional[StrictInt] = None
    __properties: ClassVar[List[str]] = ["empl_reward", "gifts", "id"]

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
        """Create an instance of GiftReward from a JSON string"""
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
        # set to None if empl_reward (nullable) is None
        # and model_fields_set contains the field
        if self.empl_reward is None and "empl_reward" in self.model_fields_set:
            _dict['empl_reward'] = None

        # set to None if gifts (nullable) is None
        # and model_fields_set contains the field
        if self.gifts is None and "gifts" in self.model_fields_set:
            _dict['gifts'] = None

        # set to None if id (nullable) is None
        # and model_fields_set contains the field
        if self.id is None and "id" in self.model_fields_set:
            _dict['id'] = None

        return _dict

    @classmethod
    def from_dict(cls, obj: Optional[Dict[str, Any]]) -> Optional[Self]:
        """Create an instance of GiftReward from a dict"""
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return cls.model_validate(obj)

        _obj = cls.model_validate({
            "empl_reward": obj.get("empl_reward"),
            "gifts": [GiftRewardGift.from_dict(_item) for _item in obj["gifts"]] if obj.get("gifts") is not None else None,
            "id": obj.get("id")
        })
        return _obj


