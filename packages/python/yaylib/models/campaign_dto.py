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
from yaylib.models.user_campaign import UserCampaign
from typing import Optional, Set
from typing_extensions import Self

class CampaignDTO(BaseModel):
    """
    CampaignDTO
    """ # noqa: E501
    banner_image: Optional[StrictStr] = None
    cover_image: Optional[StrictStr] = None
    description: Optional[StrictStr] = None
    description_url: Optional[StrictStr] = None
    end_at: Optional[StrictInt] = None
    id: Optional[StrictInt] = None
    multiplier_detail_url: Optional[StrictStr] = None
    name: Optional[StrictStr] = None
    ranking: Optional[StrictInt] = None
    start_at: Optional[StrictInt] = None
    total_points: Optional[StrictInt] = None
    user_campaign: Optional[UserCampaign] = None
    __properties: ClassVar[List[str]] = ["banner_image", "cover_image", "description", "description_url", "end_at", "id", "multiplier_detail_url", "name", "ranking", "start_at", "total_points", "user_campaign"]

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
        """Create an instance of CampaignDTO from a JSON string"""
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
        # override the default output from pydantic by calling `to_dict()` of user_campaign
        if self.user_campaign:
            _dict['user_campaign'] = self.user_campaign.to_dict()
        # set to None if banner_image (nullable) is None
        # and model_fields_set contains the field
        if self.banner_image is None and "banner_image" in self.model_fields_set:
            _dict['banner_image'] = None

        # set to None if cover_image (nullable) is None
        # and model_fields_set contains the field
        if self.cover_image is None and "cover_image" in self.model_fields_set:
            _dict['cover_image'] = None

        # set to None if description (nullable) is None
        # and model_fields_set contains the field
        if self.description is None and "description" in self.model_fields_set:
            _dict['description'] = None

        # set to None if description_url (nullable) is None
        # and model_fields_set contains the field
        if self.description_url is None and "description_url" in self.model_fields_set:
            _dict['description_url'] = None

        # set to None if end_at (nullable) is None
        # and model_fields_set contains the field
        if self.end_at is None and "end_at" in self.model_fields_set:
            _dict['end_at'] = None

        # set to None if id (nullable) is None
        # and model_fields_set contains the field
        if self.id is None and "id" in self.model_fields_set:
            _dict['id'] = None

        # set to None if multiplier_detail_url (nullable) is None
        # and model_fields_set contains the field
        if self.multiplier_detail_url is None and "multiplier_detail_url" in self.model_fields_set:
            _dict['multiplier_detail_url'] = None

        # set to None if name (nullable) is None
        # and model_fields_set contains the field
        if self.name is None and "name" in self.model_fields_set:
            _dict['name'] = None

        # set to None if ranking (nullable) is None
        # and model_fields_set contains the field
        if self.ranking is None and "ranking" in self.model_fields_set:
            _dict['ranking'] = None

        # set to None if start_at (nullable) is None
        # and model_fields_set contains the field
        if self.start_at is None and "start_at" in self.model_fields_set:
            _dict['start_at'] = None

        # set to None if total_points (nullable) is None
        # and model_fields_set contains the field
        if self.total_points is None and "total_points" in self.model_fields_set:
            _dict['total_points'] = None

        # set to None if user_campaign (nullable) is None
        # and model_fields_set contains the field
        if self.user_campaign is None and "user_campaign" in self.model_fields_set:
            _dict['user_campaign'] = None

        return _dict

    @classmethod
    def from_dict(cls, obj: Optional[Dict[str, Any]]) -> Optional[Self]:
        """Create an instance of CampaignDTO from a dict"""
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return cls.model_validate(obj)

        _obj = cls.model_validate({
            "banner_image": obj.get("banner_image"),
            "cover_image": obj.get("cover_image"),
            "description": obj.get("description"),
            "description_url": obj.get("description_url"),
            "end_at": obj.get("end_at"),
            "id": obj.get("id"),
            "multiplier_detail_url": obj.get("multiplier_detail_url"),
            "name": obj.get("name"),
            "ranking": obj.get("ranking"),
            "start_at": obj.get("start_at"),
            "total_points": obj.get("total_points"),
            "user_campaign": UserCampaign.from_dict(obj["user_campaign"]) if obj.get("user_campaign") is not None else None
        })
        return _obj


