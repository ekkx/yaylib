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

from pydantic import BaseModel, ConfigDict, StrictBool
from typing import Any, ClassVar, Dict, List, Optional
from yaylib.models.cooldown import Cooldown
from yaylib.models.qualification import Qualification
from typing import Optional, Set
from typing_extensions import Self

class PalFreePoolStatus(BaseModel):
    """
    PalFreePoolStatus
    """ # noqa: E501
    cooldown: Optional[Cooldown] = None
    pals_available: Optional[StrictBool] = None
    qualifications: Optional[List[Qualification]] = None
    qualified: Optional[StrictBool] = None
    __properties: ClassVar[List[str]] = ["cooldown", "pals_available", "qualifications", "qualified"]

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
        """Create an instance of PalFreePoolStatus from a JSON string"""
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
        # override the default output from pydantic by calling `to_dict()` of cooldown
        if self.cooldown:
            _dict['cooldown'] = self.cooldown.to_dict()
        # override the default output from pydantic by calling `to_dict()` of each item in qualifications (list)
        _items = []
        if self.qualifications:
            for _item_qualifications in self.qualifications:
                if _item_qualifications:
                    _items.append(_item_qualifications.to_dict())
            _dict['qualifications'] = _items
        # set to None if cooldown (nullable) is None
        # and model_fields_set contains the field
        if self.cooldown is None and "cooldown" in self.model_fields_set:
            _dict['cooldown'] = None

        # set to None if pals_available (nullable) is None
        # and model_fields_set contains the field
        if self.pals_available is None and "pals_available" in self.model_fields_set:
            _dict['pals_available'] = None

        # set to None if qualifications (nullable) is None
        # and model_fields_set contains the field
        if self.qualifications is None and "qualifications" in self.model_fields_set:
            _dict['qualifications'] = None

        # set to None if qualified (nullable) is None
        # and model_fields_set contains the field
        if self.qualified is None and "qualified" in self.model_fields_set:
            _dict['qualified'] = None

        return _dict

    @classmethod
    def from_dict(cls, obj: Optional[Dict[str, Any]]) -> Optional[Self]:
        """Create an instance of PalFreePoolStatus from a dict"""
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return cls.model_validate(obj)

        _obj = cls.model_validate({
            "cooldown": Cooldown.from_dict(obj["cooldown"]) if obj.get("cooldown") is not None else None,
            "pals_available": obj.get("pals_available"),
            "qualifications": [Qualification.from_dict(_item) for _item in obj["qualifications"]] if obj.get("qualifications") is not None else None,
            "qualified": obj.get("qualified")
        })
        return _obj


