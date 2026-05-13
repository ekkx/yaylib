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

class PalDetailsDTO(BaseModel):
    """
    PalDetailsDTO
    """ # noqa: E501
    cost: Optional[Union[StrictFloat, StrictInt]] = None
    level: Optional[StrictInt] = None
    nft_type: Optional[StrictStr] = None
    pal_image: Optional[StrictStr] = None
    pal_name: Optional[StrictStr] = None
    rank: Optional[StrictStr] = None
    token_id: Optional[StrictStr] = None
    __properties: ClassVar[List[str]] = ["cost", "level", "nft_type", "pal_image", "pal_name", "rank", "token_id"]

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
        """Create an instance of PalDetailsDTO from a JSON string"""
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
        # set to None if cost (nullable) is None
        # and model_fields_set contains the field
        if self.cost is None and "cost" in self.model_fields_set:
            _dict['cost'] = None

        # set to None if level (nullable) is None
        # and model_fields_set contains the field
        if self.level is None and "level" in self.model_fields_set:
            _dict['level'] = None

        # set to None if nft_type (nullable) is None
        # and model_fields_set contains the field
        if self.nft_type is None and "nft_type" in self.model_fields_set:
            _dict['nft_type'] = None

        # set to None if pal_image (nullable) is None
        # and model_fields_set contains the field
        if self.pal_image is None and "pal_image" in self.model_fields_set:
            _dict['pal_image'] = None

        # set to None if pal_name (nullable) is None
        # and model_fields_set contains the field
        if self.pal_name is None and "pal_name" in self.model_fields_set:
            _dict['pal_name'] = None

        # set to None if rank (nullable) is None
        # and model_fields_set contains the field
        if self.rank is None and "rank" in self.model_fields_set:
            _dict['rank'] = None

        # set to None if token_id (nullable) is None
        # and model_fields_set contains the field
        if self.token_id is None and "token_id" in self.model_fields_set:
            _dict['token_id'] = None

        return _dict

    @classmethod
    def from_dict(cls, obj: Optional[Dict[str, Any]]) -> Optional[Self]:
        """Create an instance of PalDetailsDTO from a dict"""
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return cls.model_validate(obj)

        _obj = cls.model_validate({
            "cost": obj.get("cost"),
            "level": obj.get("level"),
            "nft_type": obj.get("nft_type"),
            "pal_image": obj.get("pal_image"),
            "pal_name": obj.get("pal_name"),
            "rank": obj.get("rank"),
            "token_id": obj.get("token_id")
        })
        return _obj


