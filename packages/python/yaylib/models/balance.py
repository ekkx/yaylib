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

class Balance(BaseModel):
    """
    Balance
    """ # noqa: E501
    egg: Optional[StrictStr] = None
    empl: Optional[StrictStr] = None
    l1_native: Optional[StrictStr] = None
    l1_yay: Optional[StrictStr] = None
    native: Optional[StrictStr] = None
    pal: Optional[StrictStr] = None
    yay: Optional[StrictStr] = None
    __properties: ClassVar[List[str]] = ["egg", "empl", "l1_native", "l1_yay", "native", "pal", "yay"]

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
        """Create an instance of Balance from a JSON string"""
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
        # set to None if egg (nullable) is None
        # and model_fields_set contains the field
        if self.egg is None and "egg" in self.model_fields_set:
            _dict['egg'] = None

        # set to None if empl (nullable) is None
        # and model_fields_set contains the field
        if self.empl is None and "empl" in self.model_fields_set:
            _dict['empl'] = None

        # set to None if l1_native (nullable) is None
        # and model_fields_set contains the field
        if self.l1_native is None and "l1_native" in self.model_fields_set:
            _dict['l1_native'] = None

        # set to None if l1_yay (nullable) is None
        # and model_fields_set contains the field
        if self.l1_yay is None and "l1_yay" in self.model_fields_set:
            _dict['l1_yay'] = None

        # set to None if native (nullable) is None
        # and model_fields_set contains the field
        if self.native is None and "native" in self.model_fields_set:
            _dict['native'] = None

        # set to None if pal (nullable) is None
        # and model_fields_set contains the field
        if self.pal is None and "pal" in self.model_fields_set:
            _dict['pal'] = None

        # set to None if yay (nullable) is None
        # and model_fields_set contains the field
        if self.yay is None and "yay" in self.model_fields_set:
            _dict['yay'] = None

        return _dict

    @classmethod
    def from_dict(cls, obj: Optional[Dict[str, Any]]) -> Optional[Self]:
        """Create an instance of Balance from a dict"""
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return cls.model_validate(obj)

        _obj = cls.model_validate({
            "egg": obj.get("egg"),
            "empl": obj.get("empl"),
            "l1_native": obj.get("l1_native"),
            "l1_yay": obj.get("l1_yay"),
            "native": obj.get("native"),
            "pal": obj.get("pal"),
            "yay": obj.get("yay")
        })
        return _obj


