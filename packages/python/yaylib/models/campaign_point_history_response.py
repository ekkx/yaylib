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
from yaylib.models.campaign_point_history_dto import CampaignPointHistoryDTO
from typing import Optional, Set
from typing_extensions import Self

class CampaignPointHistoryResponse(BaseModel):
    """
    CampaignPointHistoryResponse
    """ # noqa: E501
    point_histories: Optional[List[CampaignPointHistoryDTO]] = None
    __properties: ClassVar[List[str]] = ["point_histories"]

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
        """Create an instance of CampaignPointHistoryResponse from a JSON string"""
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
        # override the default output from pydantic by calling `to_dict()` of each item in point_histories (list)
        _items = []
        if self.point_histories:
            for _item_point_histories in self.point_histories:
                if _item_point_histories:
                    _items.append(_item_point_histories.to_dict())
            _dict['point_histories'] = _items
        # set to None if point_histories (nullable) is None
        # and model_fields_set contains the field
        if self.point_histories is None and "point_histories" in self.model_fields_set:
            _dict['point_histories'] = None

        return _dict

    @classmethod
    def from_dict(cls, obj: Optional[Dict[str, Any]]) -> Optional[Self]:
        """Create an instance of CampaignPointHistoryResponse from a dict"""
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return cls.model_validate(obj)

        _obj = cls.model_validate({
            "point_histories": [CampaignPointHistoryDTO.from_dict(_item) for _item in obj["point_histories"]] if obj.get("point_histories") is not None else None
        })
        return _obj


