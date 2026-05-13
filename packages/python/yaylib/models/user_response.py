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
from yaylib.models.realm_gifting_ability import RealmGiftingAbility
from yaylib.models.realm_user import RealmUser
from typing import Optional, Set
from typing_extensions import Self

class UserResponse(BaseModel):
    """
    UserResponse
    """ # noqa: E501
    apple_connected: Optional[StrictBool] = None
    birth_date: Optional[StrictStr] = None
    blocking_limit: Optional[StrictInt] = None
    email_confirmed: Optional[StrictBool] = None
    facebook_connected: Optional[StrictBool] = None
    gifting_ability: Optional[RealmGiftingAbility] = None
    google_connected: Optional[StrictBool] = None
    group_phone_on: Optional[StrictBool] = None
    group_video_on: Optional[StrictBool] = None
    interests_count: Optional[StrictInt] = None
    line_connected: Optional[StrictBool] = None
    masked_email: Optional[StrictStr] = None
    phone_on: Optional[StrictBool] = None
    push_notification: Optional[StrictBool] = None
    twitter_id: Optional[StrictStr] = None
    user: Optional[RealmUser] = None
    uuid: Optional[StrictStr] = None
    video_on: Optional[StrictBool] = None
    vip_until_seconds: Optional[StrictInt] = None
    world_id_connected: Optional[StrictBool] = None
    __properties: ClassVar[List[str]] = ["apple_connected", "birth_date", "blocking_limit", "email_confirmed", "facebook_connected", "gifting_ability", "google_connected", "group_phone_on", "group_video_on", "interests_count", "line_connected", "masked_email", "phone_on", "push_notification", "twitter_id", "user", "uuid", "video_on", "vip_until_seconds", "world_id_connected"]

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
        """Create an instance of UserResponse from a JSON string"""
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
        # override the default output from pydantic by calling `to_dict()` of gifting_ability
        if self.gifting_ability:
            _dict['gifting_ability'] = self.gifting_ability.to_dict()
        # override the default output from pydantic by calling `to_dict()` of user
        if self.user:
            _dict['user'] = self.user.to_dict()
        # set to None if apple_connected (nullable) is None
        # and model_fields_set contains the field
        if self.apple_connected is None and "apple_connected" in self.model_fields_set:
            _dict['apple_connected'] = None

        # set to None if birth_date (nullable) is None
        # and model_fields_set contains the field
        if self.birth_date is None and "birth_date" in self.model_fields_set:
            _dict['birth_date'] = None

        # set to None if blocking_limit (nullable) is None
        # and model_fields_set contains the field
        if self.blocking_limit is None and "blocking_limit" in self.model_fields_set:
            _dict['blocking_limit'] = None

        # set to None if email_confirmed (nullable) is None
        # and model_fields_set contains the field
        if self.email_confirmed is None and "email_confirmed" in self.model_fields_set:
            _dict['email_confirmed'] = None

        # set to None if facebook_connected (nullable) is None
        # and model_fields_set contains the field
        if self.facebook_connected is None and "facebook_connected" in self.model_fields_set:
            _dict['facebook_connected'] = None

        # set to None if gifting_ability (nullable) is None
        # and model_fields_set contains the field
        if self.gifting_ability is None and "gifting_ability" in self.model_fields_set:
            _dict['gifting_ability'] = None

        # set to None if google_connected (nullable) is None
        # and model_fields_set contains the field
        if self.google_connected is None and "google_connected" in self.model_fields_set:
            _dict['google_connected'] = None

        # set to None if group_phone_on (nullable) is None
        # and model_fields_set contains the field
        if self.group_phone_on is None and "group_phone_on" in self.model_fields_set:
            _dict['group_phone_on'] = None

        # set to None if group_video_on (nullable) is None
        # and model_fields_set contains the field
        if self.group_video_on is None and "group_video_on" in self.model_fields_set:
            _dict['group_video_on'] = None

        # set to None if interests_count (nullable) is None
        # and model_fields_set contains the field
        if self.interests_count is None and "interests_count" in self.model_fields_set:
            _dict['interests_count'] = None

        # set to None if line_connected (nullable) is None
        # and model_fields_set contains the field
        if self.line_connected is None and "line_connected" in self.model_fields_set:
            _dict['line_connected'] = None

        # set to None if masked_email (nullable) is None
        # and model_fields_set contains the field
        if self.masked_email is None and "masked_email" in self.model_fields_set:
            _dict['masked_email'] = None

        # set to None if phone_on (nullable) is None
        # and model_fields_set contains the field
        if self.phone_on is None and "phone_on" in self.model_fields_set:
            _dict['phone_on'] = None

        # set to None if push_notification (nullable) is None
        # and model_fields_set contains the field
        if self.push_notification is None and "push_notification" in self.model_fields_set:
            _dict['push_notification'] = None

        # set to None if twitter_id (nullable) is None
        # and model_fields_set contains the field
        if self.twitter_id is None and "twitter_id" in self.model_fields_set:
            _dict['twitter_id'] = None

        # set to None if user (nullable) is None
        # and model_fields_set contains the field
        if self.user is None and "user" in self.model_fields_set:
            _dict['user'] = None

        # set to None if uuid (nullable) is None
        # and model_fields_set contains the field
        if self.uuid is None and "uuid" in self.model_fields_set:
            _dict['uuid'] = None

        # set to None if video_on (nullable) is None
        # and model_fields_set contains the field
        if self.video_on is None and "video_on" in self.model_fields_set:
            _dict['video_on'] = None

        # set to None if vip_until_seconds (nullable) is None
        # and model_fields_set contains the field
        if self.vip_until_seconds is None and "vip_until_seconds" in self.model_fields_set:
            _dict['vip_until_seconds'] = None

        # set to None if world_id_connected (nullable) is None
        # and model_fields_set contains the field
        if self.world_id_connected is None and "world_id_connected" in self.model_fields_set:
            _dict['world_id_connected'] = None

        return _dict

    @classmethod
    def from_dict(cls, obj: Optional[Dict[str, Any]]) -> Optional[Self]:
        """Create an instance of UserResponse from a dict"""
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return cls.model_validate(obj)

        _obj = cls.model_validate({
            "apple_connected": obj.get("apple_connected"),
            "birth_date": obj.get("birth_date"),
            "blocking_limit": obj.get("blocking_limit"),
            "email_confirmed": obj.get("email_confirmed"),
            "facebook_connected": obj.get("facebook_connected"),
            "gifting_ability": RealmGiftingAbility.from_dict(obj["gifting_ability"]) if obj.get("gifting_ability") is not None else None,
            "google_connected": obj.get("google_connected"),
            "group_phone_on": obj.get("group_phone_on"),
            "group_video_on": obj.get("group_video_on"),
            "interests_count": obj.get("interests_count"),
            "line_connected": obj.get("line_connected"),
            "masked_email": obj.get("masked_email"),
            "phone_on": obj.get("phone_on"),
            "push_notification": obj.get("push_notification"),
            "twitter_id": obj.get("twitter_id"),
            "user": RealmUser.from_dict(obj["user"]) if obj.get("user") is not None else None,
            "uuid": obj.get("uuid"),
            "video_on": obj.get("video_on"),
            "vip_until_seconds": obj.get("vip_until_seconds"),
            "world_id_connected": obj.get("world_id_connected")
        })
        return _obj


