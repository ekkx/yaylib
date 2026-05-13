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

from pydantic import BaseModel, ConfigDict, StrictBool, StrictInt, StrictStr
from typing import Any, ClassVar, Dict, List, Optional
from yaylib.models.details_dto import DetailsDTO
from typing import Optional, Set
from typing_extensions import Self

class CampaignVipEmplBonusDTO(BaseModel):
    """
    CampaignVipEmplBonusDTO
    """ # noqa: E501
    claimed: Optional[StrictBool] = None
    details: Optional[DetailsDTO] = None
    end_at: Optional[StrictInt] = None
    key: Optional[StrictStr] = None
    __properties: ClassVar[List[str]] = ["claimed", "details", "end_at", "key"]

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
        """Create an instance of CampaignVipEmplBonusDTO from a JSON string"""
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
        # override the default output from pydantic by calling `to_dict()` of details
        if self.details:
            _dict['details'] = self.details.to_dict()
        # set to None if claimed (nullable) is None
        # and model_fields_set contains the field
        if self.claimed is None and "claimed" in self.model_fields_set:
            _dict['claimed'] = None

        # set to None if details (nullable) is None
        # and model_fields_set contains the field
        if self.details is None and "details" in self.model_fields_set:
            _dict['details'] = None

        # set to None if end_at (nullable) is None
        # and model_fields_set contains the field
        if self.end_at is None and "end_at" in self.model_fields_set:
            _dict['end_at'] = None

        # set to None if key (nullable) is None
        # and model_fields_set contains the field
        if self.key is None and "key" in self.model_fields_set:
            _dict['key'] = None

        return _dict

    @classmethod
    def from_dict(cls, obj: Optional[Dict[str, Any]]) -> Optional[Self]:
        """Create an instance of CampaignVipEmplBonusDTO from a dict"""
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return cls.model_validate(obj)

        _obj = cls.model_validate({
            "claimed": obj.get("claimed"),
            "details": DetailsDTO.from_dict(obj["details"]) if obj.get("details") is not None else None,
            "end_at": obj.get("end_at"),
            "key": obj.get("key")
        })
        return _obj


