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

from pydantic import BaseModel, ConfigDict, StrictInt, StrictStr
from typing import Any, ClassVar, Dict, List, Optional
from yaylib.models.top_campaign_user_dto import TopCampaignUserDTO
from typing import Optional, Set
from typing_extensions import Self

class CampaignRankingResponse(BaseModel):
    """
    CampaignRankingResponse
    """ # noqa: E501
    my_ranking: Optional[StrictInt] = None
    next_cursor: Optional[StrictStr] = None
    top_campaign_users: Optional[List[TopCampaignUserDTO]] = None
    updated_at: Optional[StrictInt] = None
    __properties: ClassVar[List[str]] = ["my_ranking", "next_cursor", "top_campaign_users", "updated_at"]

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
        """Create an instance of CampaignRankingResponse from a JSON string"""
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
        # override the default output from pydantic by calling `to_dict()` of each item in top_campaign_users (list)
        _items = []
        if self.top_campaign_users:
            for _item_top_campaign_users in self.top_campaign_users:
                if _item_top_campaign_users:
                    _items.append(_item_top_campaign_users.to_dict())
            _dict['top_campaign_users'] = _items
        # set to None if my_ranking (nullable) is None
        # and model_fields_set contains the field
        if self.my_ranking is None and "my_ranking" in self.model_fields_set:
            _dict['my_ranking'] = None

        # set to None if next_cursor (nullable) is None
        # and model_fields_set contains the field
        if self.next_cursor is None and "next_cursor" in self.model_fields_set:
            _dict['next_cursor'] = None

        # set to None if top_campaign_users (nullable) is None
        # and model_fields_set contains the field
        if self.top_campaign_users is None and "top_campaign_users" in self.model_fields_set:
            _dict['top_campaign_users'] = None

        # set to None if updated_at (nullable) is None
        # and model_fields_set contains the field
        if self.updated_at is None and "updated_at" in self.model_fields_set:
            _dict['updated_at'] = None

        return _dict

    @classmethod
    def from_dict(cls, obj: Optional[Dict[str, Any]]) -> Optional[Self]:
        """Create an instance of CampaignRankingResponse from a dict"""
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return cls.model_validate(obj)

        _obj = cls.model_validate({
            "my_ranking": obj.get("my_ranking"),
            "next_cursor": obj.get("next_cursor"),
            "top_campaign_users": [TopCampaignUserDTO.from_dict(_item) for _item in obj["top_campaign_users"]] if obj.get("top_campaign_users") is not None else None,
            "updated_at": obj.get("updated_at")
        })
        return _obj


