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

class LocalizedStringDTO(BaseModel):
    """
    LocalizedStringDTO
    """ # noqa: E501
    en: Optional[StrictStr] = None
    id: Optional[StrictStr] = None
    ja: Optional[StrictStr] = None
    ko: Optional[StrictStr] = None
    th: Optional[StrictStr] = None
    zh_cn: Optional[StrictStr] = None
    zh_hk: Optional[StrictStr] = None
    __properties: ClassVar[List[str]] = ["en", "id", "ja", "ko", "th", "zh_cn", "zh_hk"]

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
        """Create an instance of LocalizedStringDTO from a JSON string"""
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
        # set to None if en (nullable) is None
        # and model_fields_set contains the field
        if self.en is None and "en" in self.model_fields_set:
            _dict['en'] = None

        # set to None if id (nullable) is None
        # and model_fields_set contains the field
        if self.id is None and "id" in self.model_fields_set:
            _dict['id'] = None

        # set to None if ja (nullable) is None
        # and model_fields_set contains the field
        if self.ja is None and "ja" in self.model_fields_set:
            _dict['ja'] = None

        # set to None if ko (nullable) is None
        # and model_fields_set contains the field
        if self.ko is None and "ko" in self.model_fields_set:
            _dict['ko'] = None

        # set to None if th (nullable) is None
        # and model_fields_set contains the field
        if self.th is None and "th" in self.model_fields_set:
            _dict['th'] = None

        # set to None if zh_cn (nullable) is None
        # and model_fields_set contains the field
        if self.zh_cn is None and "zh_cn" in self.model_fields_set:
            _dict['zh_cn'] = None

        # set to None if zh_hk (nullable) is None
        # and model_fields_set contains the field
        if self.zh_hk is None and "zh_hk" in self.model_fields_set:
            _dict['zh_hk'] = None

        return _dict

    @classmethod
    def from_dict(cls, obj: Optional[Dict[str, Any]]) -> Optional[Self]:
        """Create an instance of LocalizedStringDTO from a dict"""
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return cls.model_validate(obj)

        _obj = cls.model_validate({
            "en": obj.get("en"),
            "id": obj.get("id"),
            "ja": obj.get("ja"),
            "ko": obj.get("ko"),
            "th": obj.get("th"),
            "zh_cn": obj.get("zh_cn"),
            "zh_hk": obj.get("zh_hk")
        })
        return _obj


