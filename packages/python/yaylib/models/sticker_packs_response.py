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

from pydantic import BaseModel, ConfigDict
from typing import Any, ClassVar, Dict, List, Optional
from yaylib.models.sticker_pack import StickerPack
from typing import Optional, Set
from typing_extensions import Self

class StickerPacksResponse(BaseModel):
    """
    StickerPacksResponse
    """ # noqa: E501
    sticker_packs: Optional[List[StickerPack]] = None
    __properties: ClassVar[List[str]] = ["sticker_packs"]

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
        """Create an instance of StickerPacksResponse from a JSON string"""
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
        # override the default output from pydantic by calling `to_dict()` of each item in sticker_packs (list)
        _items = []
        if self.sticker_packs:
            for _item_sticker_packs in self.sticker_packs:
                if _item_sticker_packs:
                    _items.append(_item_sticker_packs.to_dict())
            _dict['sticker_packs'] = _items
        # set to None if sticker_packs (nullable) is None
        # and model_fields_set contains the field
        if self.sticker_packs is None and "sticker_packs" in self.model_fields_set:
            _dict['sticker_packs'] = None

        return _dict

    @classmethod
    def from_dict(cls, obj: Optional[Dict[str, Any]]) -> Optional[Self]:
        """Create an instance of StickerPacksResponse from a dict"""
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return cls.model_validate(obj)

        _obj = cls.model_validate({
            "sticker_packs": [StickerPack.from_dict(_item) for _item in obj["sticker_packs"]] if obj.get("sticker_packs") is not None else None
        })
        return _obj


