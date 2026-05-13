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
from yaylib.models.online_status_enum import OnlineStatusEnum
from yaylib.models.realm_gifting_ability import RealmGiftingAbility
from typing import Optional, Set
from typing_extensions import Self

class RealmUser(BaseModel):
    """
    RealmUser
    """ # noqa: E501
    age_verified: Optional[StrictBool] = None
    apple_connected: Optional[StrictBool] = None
    biography: Optional[StrictStr] = None
    birth_date: Optional[StrictStr] = None
    blocking_limit: Optional[StrictInt] = None
    chat_request: Optional[StrictBool] = None
    connected_by: Optional[List[StrictStr]] = None
    contact_phones: Optional[List[StrictStr]] = None
    country_code: Optional[StrictStr] = None
    cover_image: Optional[StrictStr] = None
    cover_image_thumbnail: Optional[StrictStr] = None
    created_at_seconds: Optional[StrictInt] = None
    dangerous_user: Optional[StrictBool] = None
    email_confirmed: Optional[StrictBool] = None
    facebook_connected: Optional[StrictBool] = None
    follow_pending: Optional[StrictBool] = None
    followed_by: Optional[StrictBool] = None
    followed_by_pending: Optional[StrictBool] = None
    followers_count: Optional[StrictInt] = None
    following: Optional[StrictBool] = None
    followings_count: Optional[StrictInt] = None
    frame: Optional[StrictStr] = None
    frame_thumbnail: Optional[StrictStr] = None
    from_different_generation_and_trusted: Optional[StrictBool] = None
    gender: Optional[StrictInt] = None
    generation: Optional[StrictInt] = None
    gifting_ability: Optional[RealmGiftingAbility] = None
    google_connected: Optional[StrictBool] = None
    group_phone_on: Optional[StrictBool] = None
    group_user: Optional[GroupUser] = None
    group_video_on: Optional[StrictBool] = None
    groups_users_count: Optional[StrictInt] = None
    hidden: Optional[StrictBool] = None
    hide_vip: Optional[StrictBool] = None
    id: Optional[StrictInt] = None
    interests_count: Optional[StrictInt] = None
    interests_selected: Optional[StrictBool] = None
    is_private: Optional[StrictBool] = None
    last_logged_in_at_seconds: Optional[StrictInt] = None
    line_connected: Optional[StrictBool] = None
    login_streak_count: Optional[StrictInt] = None
    masked_email: Optional[StrictStr] = None
    media_count: Optional[StrictInt] = None
    mobile_number: Optional[StrictStr] = None
    new_user: Optional[StrictBool] = None
    nickname: Optional[StrictStr] = None
    online_status: Optional[OnlineStatusEnum] = None
    phone_on: Optional[StrictBool] = None
    posts_count: Optional[StrictInt] = None
    prefecture: Optional[StrictStr] = None
    profile_icon: Optional[StrictStr] = None
    profile_icon_thumbnail: Optional[StrictStr] = None
    push_notification: Optional[StrictBool] = None
    recently_kenta: Optional[StrictBool] = None
    restricted_review_by: Optional[StrictStr] = None
    reviews_count: Optional[StrictInt] = None
    title: Optional[StrictStr] = None
    total_gifts_received_count: Optional[StrictInt] = None
    twitter_id: Optional[StrictStr] = None
    two_fa_enabled: Optional[StrictBool] = None
    updated_time_millis: Optional[StrictInt] = None
    username: Optional[StrictStr] = None
    uuid: Optional[StrictStr] = None
    video_on: Optional[StrictBool] = None
    vip: Optional[StrictBool] = None
    vip_until_seconds: Optional[StrictInt] = None
    world_id_connected: Optional[StrictBool] = None
    __properties: ClassVar[List[str]] = ["age_verified", "apple_connected", "biography", "birth_date", "blocking_limit", "chat_request", "connected_by", "contact_phones", "country_code", "cover_image", "cover_image_thumbnail", "created_at_seconds", "dangerous_user", "email_confirmed", "facebook_connected", "follow_pending", "followed_by", "followed_by_pending", "followers_count", "following", "followings_count", "frame", "frame_thumbnail", "from_different_generation_and_trusted", "gender", "generation", "gifting_ability", "google_connected", "group_phone_on", "group_user", "group_video_on", "groups_users_count", "hidden", "hide_vip", "id", "interests_count", "interests_selected", "is_private", "last_logged_in_at_seconds", "line_connected", "login_streak_count", "masked_email", "media_count", "mobile_number", "new_user", "nickname", "online_status", "phone_on", "posts_count", "prefecture", "profile_icon", "profile_icon_thumbnail", "push_notification", "recently_kenta", "restricted_review_by", "reviews_count", "title", "total_gifts_received_count", "twitter_id", "two_fa_enabled", "updated_time_millis", "username", "uuid", "video_on", "vip", "vip_until_seconds", "world_id_connected"]

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
        """Create an instance of RealmUser from a JSON string"""
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
        # override the default output from pydantic by calling `to_dict()` of group_user
        if self.group_user:
            _dict['group_user'] = self.group_user.to_dict()
        # set to None if age_verified (nullable) is None
        # and model_fields_set contains the field
        if self.age_verified is None and "age_verified" in self.model_fields_set:
            _dict['age_verified'] = None

        # set to None if apple_connected (nullable) is None
        # and model_fields_set contains the field
        if self.apple_connected is None and "apple_connected" in self.model_fields_set:
            _dict['apple_connected'] = None

        # set to None if biography (nullable) is None
        # and model_fields_set contains the field
        if self.biography is None and "biography" in self.model_fields_set:
            _dict['biography'] = None

        # set to None if birth_date (nullable) is None
        # and model_fields_set contains the field
        if self.birth_date is None and "birth_date" in self.model_fields_set:
            _dict['birth_date'] = None

        # set to None if blocking_limit (nullable) is None
        # and model_fields_set contains the field
        if self.blocking_limit is None and "blocking_limit" in self.model_fields_set:
            _dict['blocking_limit'] = None

        # set to None if chat_request (nullable) is None
        # and model_fields_set contains the field
        if self.chat_request is None and "chat_request" in self.model_fields_set:
            _dict['chat_request'] = None

        # set to None if connected_by (nullable) is None
        # and model_fields_set contains the field
        if self.connected_by is None and "connected_by" in self.model_fields_set:
            _dict['connected_by'] = None

        # set to None if contact_phones (nullable) is None
        # and model_fields_set contains the field
        if self.contact_phones is None and "contact_phones" in self.model_fields_set:
            _dict['contact_phones'] = None

        # set to None if country_code (nullable) is None
        # and model_fields_set contains the field
        if self.country_code is None and "country_code" in self.model_fields_set:
            _dict['country_code'] = None

        # set to None if cover_image (nullable) is None
        # and model_fields_set contains the field
        if self.cover_image is None and "cover_image" in self.model_fields_set:
            _dict['cover_image'] = None

        # set to None if cover_image_thumbnail (nullable) is None
        # and model_fields_set contains the field
        if self.cover_image_thumbnail is None and "cover_image_thumbnail" in self.model_fields_set:
            _dict['cover_image_thumbnail'] = None

        # set to None if created_at_seconds (nullable) is None
        # and model_fields_set contains the field
        if self.created_at_seconds is None and "created_at_seconds" in self.model_fields_set:
            _dict['created_at_seconds'] = None

        # set to None if dangerous_user (nullable) is None
        # and model_fields_set contains the field
        if self.dangerous_user is None and "dangerous_user" in self.model_fields_set:
            _dict['dangerous_user'] = None

        # set to None if email_confirmed (nullable) is None
        # and model_fields_set contains the field
        if self.email_confirmed is None and "email_confirmed" in self.model_fields_set:
            _dict['email_confirmed'] = None

        # set to None if facebook_connected (nullable) is None
        # and model_fields_set contains the field
        if self.facebook_connected is None and "facebook_connected" in self.model_fields_set:
            _dict['facebook_connected'] = None

        # set to None if follow_pending (nullable) is None
        # and model_fields_set contains the field
        if self.follow_pending is None and "follow_pending" in self.model_fields_set:
            _dict['follow_pending'] = None

        # set to None if followed_by (nullable) is None
        # and model_fields_set contains the field
        if self.followed_by is None and "followed_by" in self.model_fields_set:
            _dict['followed_by'] = None

        # set to None if followed_by_pending (nullable) is None
        # and model_fields_set contains the field
        if self.followed_by_pending is None and "followed_by_pending" in self.model_fields_set:
            _dict['followed_by_pending'] = None

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

        # set to None if frame (nullable) is None
        # and model_fields_set contains the field
        if self.frame is None and "frame" in self.model_fields_set:
            _dict['frame'] = None

        # set to None if frame_thumbnail (nullable) is None
        # and model_fields_set contains the field
        if self.frame_thumbnail is None and "frame_thumbnail" in self.model_fields_set:
            _dict['frame_thumbnail'] = None

        # set to None if from_different_generation_and_trusted (nullable) is None
        # and model_fields_set contains the field
        if self.from_different_generation_and_trusted is None and "from_different_generation_and_trusted" in self.model_fields_set:
            _dict['from_different_generation_and_trusted'] = None

        # set to None if gender (nullable) is None
        # and model_fields_set contains the field
        if self.gender is None and "gender" in self.model_fields_set:
            _dict['gender'] = None

        # set to None if generation (nullable) is None
        # and model_fields_set contains the field
        if self.generation is None and "generation" in self.model_fields_set:
            _dict['generation'] = None

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

        # set to None if group_user (nullable) is None
        # and model_fields_set contains the field
        if self.group_user is None and "group_user" in self.model_fields_set:
            _dict['group_user'] = None

        # set to None if group_video_on (nullable) is None
        # and model_fields_set contains the field
        if self.group_video_on is None and "group_video_on" in self.model_fields_set:
            _dict['group_video_on'] = None

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

        # set to None if interests_count (nullable) is None
        # and model_fields_set contains the field
        if self.interests_count is None and "interests_count" in self.model_fields_set:
            _dict['interests_count'] = None

        # set to None if interests_selected (nullable) is None
        # and model_fields_set contains the field
        if self.interests_selected is None and "interests_selected" in self.model_fields_set:
            _dict['interests_selected'] = None

        # set to None if is_private (nullable) is None
        # and model_fields_set contains the field
        if self.is_private is None and "is_private" in self.model_fields_set:
            _dict['is_private'] = None

        # set to None if last_logged_in_at_seconds (nullable) is None
        # and model_fields_set contains the field
        if self.last_logged_in_at_seconds is None and "last_logged_in_at_seconds" in self.model_fields_set:
            _dict['last_logged_in_at_seconds'] = None

        # set to None if line_connected (nullable) is None
        # and model_fields_set contains the field
        if self.line_connected is None and "line_connected" in self.model_fields_set:
            _dict['line_connected'] = None

        # set to None if login_streak_count (nullable) is None
        # and model_fields_set contains the field
        if self.login_streak_count is None and "login_streak_count" in self.model_fields_set:
            _dict['login_streak_count'] = None

        # set to None if masked_email (nullable) is None
        # and model_fields_set contains the field
        if self.masked_email is None and "masked_email" in self.model_fields_set:
            _dict['masked_email'] = None

        # set to None if media_count (nullable) is None
        # and model_fields_set contains the field
        if self.media_count is None and "media_count" in self.model_fields_set:
            _dict['media_count'] = None

        # set to None if mobile_number (nullable) is None
        # and model_fields_set contains the field
        if self.mobile_number is None and "mobile_number" in self.model_fields_set:
            _dict['mobile_number'] = None

        # set to None if new_user (nullable) is None
        # and model_fields_set contains the field
        if self.new_user is None and "new_user" in self.model_fields_set:
            _dict['new_user'] = None

        # set to None if nickname (nullable) is None
        # and model_fields_set contains the field
        if self.nickname is None and "nickname" in self.model_fields_set:
            _dict['nickname'] = None

        # set to None if online_status (nullable) is None
        # and model_fields_set contains the field
        if self.online_status is None and "online_status" in self.model_fields_set:
            _dict['online_status'] = None

        # set to None if phone_on (nullable) is None
        # and model_fields_set contains the field
        if self.phone_on is None and "phone_on" in self.model_fields_set:
            _dict['phone_on'] = None

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

        # set to None if push_notification (nullable) is None
        # and model_fields_set contains the field
        if self.push_notification is None and "push_notification" in self.model_fields_set:
            _dict['push_notification'] = None

        # set to None if recently_kenta (nullable) is None
        # and model_fields_set contains the field
        if self.recently_kenta is None and "recently_kenta" in self.model_fields_set:
            _dict['recently_kenta'] = None

        # set to None if restricted_review_by (nullable) is None
        # and model_fields_set contains the field
        if self.restricted_review_by is None and "restricted_review_by" in self.model_fields_set:
            _dict['restricted_review_by'] = None

        # set to None if reviews_count (nullable) is None
        # and model_fields_set contains the field
        if self.reviews_count is None and "reviews_count" in self.model_fields_set:
            _dict['reviews_count'] = None

        # set to None if title (nullable) is None
        # and model_fields_set contains the field
        if self.title is None and "title" in self.model_fields_set:
            _dict['title'] = None

        # set to None if total_gifts_received_count (nullable) is None
        # and model_fields_set contains the field
        if self.total_gifts_received_count is None and "total_gifts_received_count" in self.model_fields_set:
            _dict['total_gifts_received_count'] = None

        # set to None if twitter_id (nullable) is None
        # and model_fields_set contains the field
        if self.twitter_id is None and "twitter_id" in self.model_fields_set:
            _dict['twitter_id'] = None

        # set to None if two_fa_enabled (nullable) is None
        # and model_fields_set contains the field
        if self.two_fa_enabled is None and "two_fa_enabled" in self.model_fields_set:
            _dict['two_fa_enabled'] = None

        # set to None if updated_time_millis (nullable) is None
        # and model_fields_set contains the field
        if self.updated_time_millis is None and "updated_time_millis" in self.model_fields_set:
            _dict['updated_time_millis'] = None

        # set to None if username (nullable) is None
        # and model_fields_set contains the field
        if self.username is None and "username" in self.model_fields_set:
            _dict['username'] = None

        # set to None if uuid (nullable) is None
        # and model_fields_set contains the field
        if self.uuid is None and "uuid" in self.model_fields_set:
            _dict['uuid'] = None

        # set to None if video_on (nullable) is None
        # and model_fields_set contains the field
        if self.video_on is None and "video_on" in self.model_fields_set:
            _dict['video_on'] = None

        # set to None if vip (nullable) is None
        # and model_fields_set contains the field
        if self.vip is None and "vip" in self.model_fields_set:
            _dict['vip'] = None

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
        """Create an instance of RealmUser from a dict"""
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return cls.model_validate(obj)

        _obj = cls.model_validate({
            "age_verified": obj.get("age_verified"),
            "apple_connected": obj.get("apple_connected"),
            "biography": obj.get("biography"),
            "birth_date": obj.get("birth_date"),
            "blocking_limit": obj.get("blocking_limit"),
            "chat_request": obj.get("chat_request"),
            "connected_by": obj.get("connected_by"),
            "contact_phones": obj.get("contact_phones"),
            "country_code": obj.get("country_code"),
            "cover_image": obj.get("cover_image"),
            "cover_image_thumbnail": obj.get("cover_image_thumbnail"),
            "created_at_seconds": obj.get("created_at_seconds"),
            "dangerous_user": obj.get("dangerous_user"),
            "email_confirmed": obj.get("email_confirmed"),
            "facebook_connected": obj.get("facebook_connected"),
            "follow_pending": obj.get("follow_pending"),
            "followed_by": obj.get("followed_by"),
            "followed_by_pending": obj.get("followed_by_pending"),
            "followers_count": obj.get("followers_count"),
            "following": obj.get("following"),
            "followings_count": obj.get("followings_count"),
            "frame": obj.get("frame"),
            "frame_thumbnail": obj.get("frame_thumbnail"),
            "from_different_generation_and_trusted": obj.get("from_different_generation_and_trusted"),
            "gender": obj.get("gender"),
            "generation": obj.get("generation"),
            "gifting_ability": RealmGiftingAbility.from_dict(obj["gifting_ability"]) if obj.get("gifting_ability") is not None else None,
            "google_connected": obj.get("google_connected"),
            "group_phone_on": obj.get("group_phone_on"),
            "group_user": GroupUser.from_dict(obj["group_user"]) if obj.get("group_user") is not None else None,
            "group_video_on": obj.get("group_video_on"),
            "groups_users_count": obj.get("groups_users_count"),
            "hidden": obj.get("hidden"),
            "hide_vip": obj.get("hide_vip"),
            "id": obj.get("id"),
            "interests_count": obj.get("interests_count"),
            "interests_selected": obj.get("interests_selected"),
            "is_private": obj.get("is_private"),
            "last_logged_in_at_seconds": obj.get("last_logged_in_at_seconds"),
            "line_connected": obj.get("line_connected"),
            "login_streak_count": obj.get("login_streak_count"),
            "masked_email": obj.get("masked_email"),
            "media_count": obj.get("media_count"),
            "mobile_number": obj.get("mobile_number"),
            "new_user": obj.get("new_user"),
            "nickname": obj.get("nickname"),
            "online_status": obj.get("online_status"),
            "phone_on": obj.get("phone_on"),
            "posts_count": obj.get("posts_count"),
            "prefecture": obj.get("prefecture"),
            "profile_icon": obj.get("profile_icon"),
            "profile_icon_thumbnail": obj.get("profile_icon_thumbnail"),
            "push_notification": obj.get("push_notification"),
            "recently_kenta": obj.get("recently_kenta"),
            "restricted_review_by": obj.get("restricted_review_by"),
            "reviews_count": obj.get("reviews_count"),
            "title": obj.get("title"),
            "total_gifts_received_count": obj.get("total_gifts_received_count"),
            "twitter_id": obj.get("twitter_id"),
            "two_fa_enabled": obj.get("two_fa_enabled"),
            "updated_time_millis": obj.get("updated_time_millis"),
            "username": obj.get("username"),
            "uuid": obj.get("uuid"),
            "video_on": obj.get("video_on"),
            "vip": obj.get("vip"),
            "vip_until_seconds": obj.get("vip_until_seconds"),
            "world_id_connected": obj.get("world_id_connected")
        })
        return _obj

from yaylib.models.group_user import GroupUser
# TODO: Rewrite to not use raise_errors
RealmUser.model_rebuild(raise_errors=False)

