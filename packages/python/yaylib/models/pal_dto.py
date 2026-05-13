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
from yaylib.models.attribute import Attribute
from yaylib.models.max_attribute import MaxAttribute
from typing import Optional, Set
from typing_extensions import Self

class PalDTO(BaseModel):
    """
    PalDTO
    """ # noqa: E501
    age_in_days: Optional[StrictInt] = None
    attributes: Optional[Attribute] = None
    current_empl_earned: Optional[StrictStr] = None
    current_level: Optional[StrictInt] = None
    description: Optional[StrictStr] = None
    empl_earned_limit: Optional[StrictStr] = None
    from_liquidity_pool: Optional[StrictBool] = None
    hatching_animation_url: Optional[StrictStr] = None
    image: Optional[StrictStr] = None
    in_locked_party: Optional[StrictBool] = None
    is_alive: Optional[StrictBool] = None
    is_pending: Optional[StrictBool] = None
    is_valid_to_evolve: Optional[StrictBool] = None
    is_valid_to_level_up: Optional[StrictBool] = None
    max_age_in_days: Optional[StrictInt] = None
    max_attributes: Optional[MaxAttribute] = None
    name: Optional[StrictStr] = None
    original_wallet_address: Optional[StrictStr] = None
    race_status: Optional[StrictStr] = None
    token_id: Optional[StrictStr] = None
    __properties: ClassVar[List[str]] = ["age_in_days", "attributes", "current_empl_earned", "current_level", "description", "empl_earned_limit", "from_liquidity_pool", "hatching_animation_url", "image", "in_locked_party", "is_alive", "is_pending", "is_valid_to_evolve", "is_valid_to_level_up", "max_age_in_days", "max_attributes", "name", "original_wallet_address", "race_status", "token_id"]

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
        """Create an instance of PalDTO from a JSON string"""
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
        # override the default output from pydantic by calling `to_dict()` of attributes
        if self.attributes:
            _dict['attributes'] = self.attributes.to_dict()
        # override the default output from pydantic by calling `to_dict()` of max_attributes
        if self.max_attributes:
            _dict['max_attributes'] = self.max_attributes.to_dict()
        # set to None if age_in_days (nullable) is None
        # and model_fields_set contains the field
        if self.age_in_days is None and "age_in_days" in self.model_fields_set:
            _dict['age_in_days'] = None

        # set to None if attributes (nullable) is None
        # and model_fields_set contains the field
        if self.attributes is None and "attributes" in self.model_fields_set:
            _dict['attributes'] = None

        # set to None if current_empl_earned (nullable) is None
        # and model_fields_set contains the field
        if self.current_empl_earned is None and "current_empl_earned" in self.model_fields_set:
            _dict['current_empl_earned'] = None

        # set to None if current_level (nullable) is None
        # and model_fields_set contains the field
        if self.current_level is None and "current_level" in self.model_fields_set:
            _dict['current_level'] = None

        # set to None if description (nullable) is None
        # and model_fields_set contains the field
        if self.description is None and "description" in self.model_fields_set:
            _dict['description'] = None

        # set to None if empl_earned_limit (nullable) is None
        # and model_fields_set contains the field
        if self.empl_earned_limit is None and "empl_earned_limit" in self.model_fields_set:
            _dict['empl_earned_limit'] = None

        # set to None if from_liquidity_pool (nullable) is None
        # and model_fields_set contains the field
        if self.from_liquidity_pool is None and "from_liquidity_pool" in self.model_fields_set:
            _dict['from_liquidity_pool'] = None

        # set to None if hatching_animation_url (nullable) is None
        # and model_fields_set contains the field
        if self.hatching_animation_url is None and "hatching_animation_url" in self.model_fields_set:
            _dict['hatching_animation_url'] = None

        # set to None if image (nullable) is None
        # and model_fields_set contains the field
        if self.image is None and "image" in self.model_fields_set:
            _dict['image'] = None

        # set to None if in_locked_party (nullable) is None
        # and model_fields_set contains the field
        if self.in_locked_party is None and "in_locked_party" in self.model_fields_set:
            _dict['in_locked_party'] = None

        # set to None if is_alive (nullable) is None
        # and model_fields_set contains the field
        if self.is_alive is None and "is_alive" in self.model_fields_set:
            _dict['is_alive'] = None

        # set to None if is_pending (nullable) is None
        # and model_fields_set contains the field
        if self.is_pending is None and "is_pending" in self.model_fields_set:
            _dict['is_pending'] = None

        # set to None if is_valid_to_evolve (nullable) is None
        # and model_fields_set contains the field
        if self.is_valid_to_evolve is None and "is_valid_to_evolve" in self.model_fields_set:
            _dict['is_valid_to_evolve'] = None

        # set to None if is_valid_to_level_up (nullable) is None
        # and model_fields_set contains the field
        if self.is_valid_to_level_up is None and "is_valid_to_level_up" in self.model_fields_set:
            _dict['is_valid_to_level_up'] = None

        # set to None if max_age_in_days (nullable) is None
        # and model_fields_set contains the field
        if self.max_age_in_days is None and "max_age_in_days" in self.model_fields_set:
            _dict['max_age_in_days'] = None

        # set to None if max_attributes (nullable) is None
        # and model_fields_set contains the field
        if self.max_attributes is None and "max_attributes" in self.model_fields_set:
            _dict['max_attributes'] = None

        # set to None if name (nullable) is None
        # and model_fields_set contains the field
        if self.name is None and "name" in self.model_fields_set:
            _dict['name'] = None

        # set to None if original_wallet_address (nullable) is None
        # and model_fields_set contains the field
        if self.original_wallet_address is None and "original_wallet_address" in self.model_fields_set:
            _dict['original_wallet_address'] = None

        # set to None if race_status (nullable) is None
        # and model_fields_set contains the field
        if self.race_status is None and "race_status" in self.model_fields_set:
            _dict['race_status'] = None

        # set to None if token_id (nullable) is None
        # and model_fields_set contains the field
        if self.token_id is None and "token_id" in self.model_fields_set:
            _dict['token_id'] = None

        return _dict

    @classmethod
    def from_dict(cls, obj: Optional[Dict[str, Any]]) -> Optional[Self]:
        """Create an instance of PalDTO from a dict"""
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return cls.model_validate(obj)

        _obj = cls.model_validate({
            "age_in_days": obj.get("age_in_days"),
            "attributes": Attribute.from_dict(obj["attributes"]) if obj.get("attributes") is not None else None,
            "current_empl_earned": obj.get("current_empl_earned"),
            "current_level": obj.get("current_level"),
            "description": obj.get("description"),
            "empl_earned_limit": obj.get("empl_earned_limit"),
            "from_liquidity_pool": obj.get("from_liquidity_pool"),
            "hatching_animation_url": obj.get("hatching_animation_url"),
            "image": obj.get("image"),
            "in_locked_party": obj.get("in_locked_party"),
            "is_alive": obj.get("is_alive"),
            "is_pending": obj.get("is_pending"),
            "is_valid_to_evolve": obj.get("is_valid_to_evolve"),
            "is_valid_to_level_up": obj.get("is_valid_to_level_up"),
            "max_age_in_days": obj.get("max_age_in_days"),
            "max_attributes": MaxAttribute.from_dict(obj["max_attributes"]) if obj.get("max_attributes") is not None else None,
            "name": obj.get("name"),
            "original_wallet_address": obj.get("original_wallet_address"),
            "race_status": obj.get("race_status"),
            "token_id": obj.get("token_id")
        })
        return _obj


