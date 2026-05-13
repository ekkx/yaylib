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

from pydantic import BaseModel, ConfigDict, StrictFloat, StrictInt, StrictStr
from typing import Any, ClassVar, Dict, List, Optional, Union
from typing import Optional, Set
from typing_extensions import Self

class HatchGacha(BaseModel):
    """
    HatchGacha
    """ # noqa: E501
    empl_cost: Optional[Union[StrictFloat, StrictInt]] = None
    gacha_type: Optional[Dict[str, Any]] = None
    pal_image_url: Optional[StrictStr] = None
    pal_name: Optional[StrictStr] = None
    pal_rank: Optional[Dict[str, Any]] = None
    token_id: Optional[StrictStr] = None
    __properties: ClassVar[List[str]] = ["empl_cost", "gacha_type", "pal_image_url", "pal_name", "pal_rank", "token_id"]

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
        """Create an instance of HatchGacha from a JSON string"""
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
        # set to None if empl_cost (nullable) is None
        # and model_fields_set contains the field
        if self.empl_cost is None and "empl_cost" in self.model_fields_set:
            _dict['empl_cost'] = None

        # set to None if gacha_type (nullable) is None
        # and model_fields_set contains the field
        if self.gacha_type is None and "gacha_type" in self.model_fields_set:
            _dict['gacha_type'] = None

        # set to None if pal_image_url (nullable) is None
        # and model_fields_set contains the field
        if self.pal_image_url is None and "pal_image_url" in self.model_fields_set:
            _dict['pal_image_url'] = None

        # set to None if pal_name (nullable) is None
        # and model_fields_set contains the field
        if self.pal_name is None and "pal_name" in self.model_fields_set:
            _dict['pal_name'] = None

        # set to None if pal_rank (nullable) is None
        # and model_fields_set contains the field
        if self.pal_rank is None and "pal_rank" in self.model_fields_set:
            _dict['pal_rank'] = None

        # set to None if token_id (nullable) is None
        # and model_fields_set contains the field
        if self.token_id is None and "token_id" in self.model_fields_set:
            _dict['token_id'] = None

        return _dict

    @classmethod
    def from_dict(cls, obj: Optional[Dict[str, Any]]) -> Optional[Self]:
        """Create an instance of HatchGacha from a dict"""
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return cls.model_validate(obj)

        _obj = cls.model_validate({
            "empl_cost": obj.get("empl_cost"),
            "gacha_type": obj.get("gacha_type"),
            "pal_image_url": obj.get("pal_image_url"),
            "pal_name": obj.get("pal_name"),
            "pal_rank": obj.get("pal_rank"),
            "token_id": obj.get("token_id")
        })
        return _obj


