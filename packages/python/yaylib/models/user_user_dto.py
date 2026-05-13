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
from typing import Optional, Set
from typing_extensions import Self

class UserUserDTO(BaseModel):
    """
    UserUserDTO
    """ # noqa: E501
    age_verified: Optional[StrictBool] = None
    biography: Optional[StrictStr] = None
    cover_image: Optional[StrictStr] = None
    cover_image_thumbnail: Optional[StrictStr] = None
    created_at: Optional[StrictInt] = None
    dangerous_user: Optional[StrictBool] = None
    follow_pending: Optional[StrictBool] = None
    followed_by: Optional[StrictBool] = None
    followers_count: Optional[StrictInt] = None
    following: Optional[StrictBool] = None
    followings_count: Optional[StrictInt] = None
    gender: Optional[StrictInt] = None
    groups_users_count: Optional[StrictInt] = None
    hidden: Optional[StrictBool] = None
    hide_vip: Optional[StrictBool] = None
    id: Optional[StrictInt] = None
    is_private: Optional[StrictBool] = None
    new_user: Optional[StrictBool] = None
    nickname: Optional[StrictStr] = None
    posts_count: Optional[StrictInt] = None
    prefecture: Optional[StrictStr] = None
    profile_icon: Optional[StrictStr] = None
    profile_icon_thumbnail: Optional[StrictStr] = None
    recently_kenta: Optional[StrictBool] = None
    reviews_count: Optional[StrictInt] = None
    title: Optional[StrictStr] = None
    vip: Optional[StrictBool] = None
    __properties: ClassVar[List[str]] = ["age_verified", "biography", "cover_image", "cover_image_thumbnail", "created_at", "dangerous_user", "follow_pending", "followed_by", "followers_count", "following", "followings_count", "gender", "groups_users_count", "hidden", "hide_vip", "id", "is_private", "new_user", "nickname", "posts_count", "prefecture", "profile_icon", "profile_icon_thumbnail", "recently_kenta", "reviews_count", "title", "vip"]

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
        """Create an instance of UserUserDTO from a JSON string"""
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
        # set to None if age_verified (nullable) is None
        # and model_fields_set contains the field
        if self.age_verified is None and "age_verified" in self.model_fields_set:
            _dict['age_verified'] = None

        # set to None if biography (nullable) is None
        # and model_fields_set contains the field
        if self.biography is None and "biography" in self.model_fields_set:
            _dict['biography'] = None

        # set to None if cover_image (nullable) is None
        # and model_fields_set contains the field
        if self.cover_image is None and "cover_image" in self.model_fields_set:
            _dict['cover_image'] = None

        # set to None if cover_image_thumbnail (nullable) is None
        # and model_fields_set contains the field
        if self.cover_image_thumbnail is None and "cover_image_thumbnail" in self.model_fields_set:
            _dict['cover_image_thumbnail'] = None

        # set to None if created_at (nullable) is None
        # and model_fields_set contains the field
        if self.created_at is None and "created_at" in self.model_fields_set:
            _dict['created_at'] = None

        # set to None if dangerous_user (nullable) is None
        # and model_fields_set contains the field
        if self.dangerous_user is None and "dangerous_user" in self.model_fields_set:
            _dict['dangerous_user'] = None

        # set to None if follow_pending (nullable) is None
        # and model_fields_set contains the field
        if self.follow_pending is None and "follow_pending" in self.model_fields_set:
            _dict['follow_pending'] = None

        # set to None if followed_by (nullable) is None
        # and model_fields_set contains the field
        if self.followed_by is None and "followed_by" in self.model_fields_set:
            _dict['followed_by'] = None

        # set to None if followers_count (nullable) is None
        # and model_fields_set contains the field
        if self.followers_count is None and "followers_count" in self.model_fields_set:
            _dict['followers_count'] = None

        # set to None if following (nullable) is None
        # and model_fields_set contains the field
        if self.following is None and "following" in self.model_fields_set:
            _dict['following'] = None

        # set to None if followings_count (nullable) is None
        # and model_fields_set contains the field
        if self.followings_count is None and "followings_count" in self.model_fields_set:
            _dict['followings_count'] = None

        # set to None if gender (nullable) is None
        # and model_fields_set contains the field
        if self.gender is None and "gender" in self.model_fields_set:
            _dict['gender'] = None

        # set to None if groups_users_count (nullable) is None
        # and model_fields_set contains the field
        if self.groups_users_count is None and "groups_users_count" in self.model_fields_set:
            _dict['groups_users_count'] = None

        # set to None if hidden (nullable) is None
        # and model_fields_set contains the field
        if self.hidden is None and "hidden" in self.model_fields_set:
            _dict['hidden'] = None

        # set to None if hide_vip (nullable) is None
        # and model_fields_set contains the field
        if self.hide_vip is None and "hide_vip" in self.model_fields_set:
            _dict['hide_vip'] = None

        # set to None if id (nullable) is None
        # and model_fields_set contains the field
        if self.id is None and "id" in self.model_fields_set:
            _dict['id'] = None

        # set to None if is_private (nullable) is None
        # and model_fields_set contains the field
        if self.is_private is None and "is_private" in self.model_fields_set:
            _dict['is_private'] = None

        # set to None if new_user (nullable) is None
        # and model_fields_set contains the field
        if self.new_user is None and "new_user" in self.model_fields_set:
            _dict['new_user'] = None

        # set to None if nickname (nullable) is None
        # and model_fields_set contains the field
        if self.nickname is None and "nickname" in self.model_fields_set:
            _dict['nickname'] = None

        # set to None if posts_count (nullable) is None
        # and model_fields_set contains the field
        if self.posts_count is None and "posts_count" in self.model_fields_set:
            _dict['posts_count'] = None

        # set to None if prefecture (nullable) is None
        # and model_fields_set contains the field
        if self.prefecture is None and "prefecture" in self.model_fields_set:
            _dict['prefecture'] = None

        # set to None if profile_icon (nullable) is None
        # and model_fields_set contains the field
        if self.profile_icon is None and "profile_icon" in self.model_fields_set:
            _dict['profile_icon'] = None

        # set to None if profile_icon_thumbnail (nullable) is None
        # and model_fields_set contains the field
        if self.profile_icon_thumbnail is None and "profile_icon_thumbnail" in self.model_fields_set:
            _dict['profile_icon_thumbnail'] = None

        # set to None if recently_kenta (nullable) is None
        # and model_fields_set contains the field
        if self.recently_kenta is None and "recently_kenta" in self.model_fields_set:
            _dict['recently_kenta'] = None

        # set to None if reviews_count (nullable) is None
        # and model_fields_set contains the field
        if self.reviews_count is None and "reviews_count" in self.model_fields_set:
            _dict['reviews_count'] = None

        # set to None if title (nullable) is None
        # and model_fields_set contains the field
        if self.title is None and "title" in self.model_fields_set:
            _dict['title'] = None

        # set to None if vip (nullable) is None
        # and model_fields_set contains the field
        if self.vip is None and "vip" in self.model_fields_set:
            _dict['vip'] = None

        return _dict

    @classmethod
    def from_dict(cls, obj: Optional[Dict[str, Any]]) -> Optional[Self]:
        """Create an instance of UserUserDTO from a dict"""
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return cls.model_validate(obj)

        _obj = cls.model_validate({
            "age_verified": obj.get("age_verified"),
            "biography": obj.get("biography"),
            "cover_image": obj.get("cover_image"),
            "cover_image_thumbnail": obj.get("cover_image_thumbnail"),
            "created_at": obj.get("created_at"),
            "dangerous_user": obj.get("dangerous_user"),
            "follow_pending": obj.get("follow_pending"),
            "followed_by": obj.get("followed_by"),
            "followers_count": obj.get("followers_count"),
            "following": obj.get("following"),
            "followings_count": obj.get("followings_count"),
            "gender": obj.get("gender"),
            "groups_users_count": obj.get("groups_users_count"),
            "hidden": obj.get("hidden"),
            "hide_vip": obj.get("hide_vip"),
            "id": obj.get("id"),
            "is_private": obj.get("is_private"),
            "new_user": obj.get("new_user"),
            "nickname": obj.get("nickname"),
            "posts_count": obj.get("posts_count"),
            "prefecture": obj.get("prefecture"),
            "profile_icon": obj.get("profile_icon"),
            "profile_icon_thumbnail": obj.get("profile_icon_thumbnail"),
            "recently_kenta": obj.get("recently_kenta"),
            "reviews_count": obj.get("reviews_count"),
            "title": obj.get("title"),
            "vip": obj.get("vip")
        })
        return _obj


