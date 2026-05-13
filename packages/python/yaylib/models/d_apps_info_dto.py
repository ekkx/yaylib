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
from yaylib.models.d_app_dto import DAppDTO
from typing import Optional, Set
from typing_extensions import Self

class DAppsInfoDTO(BaseModel):
    """
    DAppsInfoDTO
    """ # noqa: E501
    dapps: Optional[List[DAppDTO]] = None
    educational_links: Optional[List[DAppDTO]] = None
    __properties: ClassVar[List[str]] = ["dapps", "educational_links"]

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
        """Create an instance of DAppsInfoDTO from a JSON string"""
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
        # override the default output from pydantic by calling `to_dict()` of each item in dapps (list)
        _items = []
        if self.dapps:
            for _item_dapps in self.dapps:
                if _item_dapps:
                    _items.append(_item_dapps.to_dict())
            _dict['dapps'] = _items
        # override the default output from pydantic by calling `to_dict()` of each item in educational_links (list)
        _items = []
        if self.educational_links:
            for _item_educational_links in self.educational_links:
                if _item_educational_links:
                    _items.append(_item_educational_links.to_dict())
            _dict['educational_links'] = _items
        # set to None if dapps (nullable) is None
        # and model_fields_set contains the field
        if self.dapps is None and "dapps" in self.model_fields_set:
            _dict['dapps'] = None

        # set to None if educational_links (nullable) is None
        # and model_fields_set contains the field
        if self.educational_links is None and "educational_links" in self.model_fields_set:
            _dict['educational_links'] = None

        return _dict

    @classmethod
    def from_dict(cls, obj: Optional[Dict[str, Any]]) -> Optional[Self]:
        """Create an instance of DAppsInfoDTO from a dict"""
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return cls.model_validate(obj)

        _obj = cls.model_validate({
            "dapps": [DAppDTO.from_dict(_item) for _item in obj["dapps"]] if obj.get("dapps") is not None else None,
            "educational_links": [DAppDTO.from_dict(_item) for _item in obj["educational_links"]] if obj.get("educational_links") is not None else None
        })
        return _obj


