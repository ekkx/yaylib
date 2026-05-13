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
from yaylib.models.gender import Gender
from yaylib.models.gifting_ability import GiftingAbility
from yaylib.models.online_status import OnlineStatus
from yaylib.models.review_restriction import ReviewRestriction
from yaylib.models.title import Title
from typing import Optional, Set
from typing_extensions import Self

class User(BaseModel):
    """
    User
    """ # noqa: E501
    biography: Optional[StrictStr] = None
    birth_date: Optional[StrictStr] = None
    blocking_limit: Optional[StrictInt] = None
    connected_by: Optional[Dict[str, Any]] = None
    contact_phones: Optional[List[StrictStr]] = None
    country: Optional[Dict[str, Any]] = None
    cover_image: Optional[StrictStr] = None
    cover_image_thumbnail: Optional[StrictStr] = None
    created_at_millis: Optional[StrictInt] = None
    followers_count: Optional[StrictInt] = None
    followings_count: Optional[StrictInt] = None
    gender: Optional[Gender] = None
    gift_received_counter: Optional[StrictInt] = None
    gifting_ability: Optional[GiftingAbility] = None
    groups_users_count: Optional[StrictInt] = None
    id: Optional[StrictInt] = None
    interests_count: Optional[StrictInt] = None
    is_age_verified: Optional[StrictBool] = None
    is_apple_connected: Optional[StrictBool] = None
    is_chat_request_on: Optional[StrictBool] = None
    is_dangerous: Optional[StrictBool] = None
    is_email_confirmed: Optional[StrictBool] = None
    is_facebook_connected: Optional[StrictBool] = None
    is_follow_pending: Optional[StrictBool] = None
    is_followed_by: Optional[StrictBool] = None
    is_followed_by_pending: Optional[StrictBool] = None
    is_following: Optional[StrictBool] = None
    is_from_different_generation_and_trusted: Optional[StrictBool] = None
    is_google_connected: Optional[StrictBool] = None
    is_group_phone_on: Optional[StrictBool] = None
    is_group_video_on: Optional[StrictBool] = None
    is_hidden: Optional[StrictBool] = None
    is_hide_vip: Optional[StrictBool] = None
    is_interests_selected: Optional[StrictBool] = None
    is_line_connected: Optional[StrictBool] = None
    is_muted: Optional[StrictBool] = None
    is_new: Optional[StrictBool] = None
    is_phone_on: Optional[StrictBool] = None
    is_private: Optional[StrictBool] = None
    is_recently_penalized: Optional[StrictBool] = None
    is_registered: Optional[StrictBool] = None
    is_two_fa_enabled: Optional[StrictBool] = None
    is_video_on: Optional[StrictBool] = None
    is_vip: Optional[StrictBool] = None
    is_world_id_connected: Optional[StrictBool] = None
    last_logged_in_at_millis: Optional[StrictInt] = None
    login_streak_count: Optional[StrictInt] = None
    masked_email: Optional[StrictStr] = None
    matching_id: Optional[StrictInt] = None
    media_count: Optional[StrictInt] = None
    mobile_number: Optional[StrictStr] = None
    nickname: Optional[StrictStr] = None
    online_status: Optional[OnlineStatus] = None
    posts_count: Optional[StrictInt] = None
    prefecture: Optional[StrictStr] = None
    profile_icon: Optional[StrictStr] = None
    profile_icon_frame: Optional[StrictStr] = None
    profile_icon_frame_thumbnail: Optional[StrictStr] = None
    profile_icon_thumbnail: Optional[StrictStr] = None
    review_restriction: Optional[ReviewRestriction] = None
    reviews_count: Optional[StrictInt] = None
    title: Optional[Title] = None
    twitter_id: Optional[StrictStr] = None
    username: Optional[StrictStr] = None
    uuid: Optional[StrictStr] = None
    vip_until_millis: Optional[StrictInt] = None
    __properties: ClassVar[List[str]] = ["biography", "birth_date", "blocking_limit", "connected_by", "contact_phones", "country", "cover_image", "cover_image_thumbnail", "created_at_millis", "followers_count", "followings_count", "gender", "gift_received_counter", "gifting_ability", "groups_users_count", "id", "interests_count", "is_age_verified", "is_apple_connected", "is_chat_request_on", "is_dangerous", "is_email_confirmed", "is_facebook_connected", "is_follow_pending", "is_followed_by", "is_followed_by_pending", "is_following", "is_from_different_generation_and_trusted", "is_google_connected", "is_group_phone_on", "is_group_video_on", "is_hidden", "is_hide_vip", "is_interests_selected", "is_line_connected", "is_muted", "is_new", "is_phone_on", "is_private", "is_recently_penalized", "is_registered", "is_two_fa_enabled", "is_video_on", "is_vip", "is_world_id_connected", "last_logged_in_at_millis", "login_streak_count", "masked_email", "matching_id", "media_count", "mobile_number", "nickname", "online_status", "posts_count", "prefecture", "profile_icon", "profile_icon_frame", "profile_icon_frame_thumbnail", "profile_icon_thumbnail", "review_restriction", "reviews_count", "title", "twitter_id", "username", "uuid", "vip_until_millis"]

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
        """Create an instance of User from a JSON string"""
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
        # override the default output from pydantic by calling `to_dict()` of gender
        if self.gender:
            _dict['gender'] = self.gender.to_dict()
        # override the default output from pydantic by calling `to_dict()` of gifting_ability
        if self.gifting_ability:
            _dict['gifting_ability'] = self.gifting_ability.to_dict()
        # override the default output from pydantic by calling `to_dict()` of online_status
        if self.online_status:
            _dict['online_status'] = self.online_status.to_dict()
        # override the default output from pydantic by calling `to_dict()` of review_restriction
        if self.review_restriction:
            _dict['review_restriction'] = self.review_restriction.to_dict()
        # override the default output from pydantic by calling `to_dict()` of title
        if self.title:
            _dict['title'] = self.title.to_dict()
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

        # set to None if connected_by (nullable) is None
        # and model_fields_set contains the field
        if self.connected_by is None and "connected_by" in self.model_fields_set:
            _dict['connected_by'] = None

        # set to None if contact_phones (nullable) is None
        # and model_fields_set contains the field
        if self.contact_phones is None and "contact_phones" in self.model_fields_set:
            _dict['contact_phones'] = None

        # set to None if country (nullable) is None
        # and model_fields_set contains the field
        if self.country is None and "country" in self.model_fields_set:
            _dict['country'] = None

        # set to None if cover_image (nullable) is None
        # and model_fields_set contains the field
        if self.cover_image is None and "cover_image" in self.model_fields_set:
            _dict['cover_image'] = None

        # set to None if cover_image_thumbnail (nullable) is None
        # and model_fields_set contains the field
        if self.cover_image_thumbnail is None and "cover_image_thumbnail" in self.model_fields_set:
            _dict['cover_image_thumbnail'] = None

        # set to None if created_at_millis (nullable) is None
        # and model_fields_set contains the field
        if self.created_at_millis is None and "created_at_millis" in self.model_fields_set:
            _dict['created_at_millis'] = None

        # set to None if followers_count (nullable) is None
        # and model_fields_set contains the field
        if self.followers_count is None and "followers_count" in self.model_fields_set:
            _dict['followers_count'] = None

        # set to None if followings_count (nullable) is None
        # and model_fields_set contains the field
        if self.followings_count is None and "followings_count" in self.model_fields_set:
            _dict['followings_count'] = None

        # set to None if gender (nullable) is None
        # and model_fields_set contains the field
        if self.gender is None and "gender" in self.model_fields_set:
            _dict['gender'] = None

        # set to None if gift_received_counter (nullable) is None
        # and model_fields_set contains the field
        if self.gift_received_counter is None and "gift_received_counter" in self.model_fields_set:
            _dict['gift_received_counter'] = None

        # set to None if gifting_ability (nullable) is None
        # and model_fields_set contains the field
        if self.gifting_ability is None and "gifting_ability" in self.model_fields_set:
            _dict['gifting_ability'] = None

        # set to None if groups_users_count (nullable) is None
        # and model_fields_set contains the field
        if self.groups_users_count is None and "groups_users_count" in self.model_fields_set:
            _dict['groups_users_count'] = None

        # set to None if id (nullable) is None
        # and model_fields_set contains the field
        if self.id is None and "id" in self.model_fields_set:
            _dict['id'] = None

        # set to None if interests_count (nullable) is None
        # and model_fields_set contains the field
        if self.interests_count is None and "interests_count" in self.model_fields_set:
            _dict['interests_count'] = None

        # set to None if is_age_verified (nullable) is None
        # and model_fields_set contains the field
        if self.is_age_verified is None and "is_age_verified" in self.model_fields_set:
            _dict['is_age_verified'] = None

        # set to None if is_apple_connected (nullable) is None
        # and model_fields_set contains the field
        if self.is_apple_connected is None and "is_apple_connected" in self.model_fields_set:
            _dict['is_apple_connected'] = None

        # set to None if is_chat_request_on (nullable) is None
        # and model_fields_set contains the field
        if self.is_chat_request_on is None and "is_chat_request_on" in self.model_fields_set:
            _dict['is_chat_request_on'] = None

        # set to None if is_dangerous (nullable) is None
        # and model_fields_set contains the field
        if self.is_dangerous is None and "is_dangerous" in self.model_fields_set:
            _dict['is_dangerous'] = None

        # set to None if is_email_confirmed (nullable) is None
        # and model_fields_set contains the field
        if self.is_email_confirmed is None and "is_email_confirmed" in self.model_fields_set:
            _dict['is_email_confirmed'] = None

        # set to None if is_facebook_connected (nullable) is None
        # and model_fields_set contains the field
        if self.is_facebook_connected is None and "is_facebook_connected" in self.model_fields_set:
            _dict['is_facebook_connected'] = None

        # set to None if is_follow_pending (nullable) is None
        # and model_fields_set contains the field
        if self.is_follow_pending is None and "is_follow_pending" in self.model_fields_set:
            _dict['is_follow_pending'] = None

        # set to None if is_followed_by (nullable) is None
        # and model_fields_set contains the field
        if self.is_followed_by is None and "is_followed_by" in self.model_fields_set:
            _dict['is_followed_by'] = None

        # set to None if is_followed_by_pending (nullable) is None
        # and model_fields_set contains the field
        if self.is_followed_by_pending is None and "is_followed_by_pending" in self.model_fields_set:
            _dict['is_followed_by_pending'] = None

        # set to None if is_following (nullable) is None
        # and model_fields_set contains the field
        if self.is_following is None and "is_following" in self.model_fields_set:
            _dict['is_following'] = None

        # set to None if is_from_different_generation_and_trusted (nullable) is None
        # and model_fields_set contains the field
        if self.is_from_different_generation_and_trusted is None and "is_from_different_generation_and_trusted" in self.model_fields_set:
            _dict['is_from_different_generation_and_trusted'] = None

        # set to None if is_google_connected (nullable) is None
        # and model_fields_set contains the field
        if self.is_google_connected is None and "is_google_connected" in self.model_fields_set:
            _dict['is_google_connected'] = None

        # set to None if is_group_phone_on (nullable) is None
        # and model_fields_set contains the field
        if self.is_group_phone_on is None and "is_group_phone_on" in self.model_fields_set:
            _dict['is_group_phone_on'] = None

        # set to None if is_group_video_on (nullable) is None
        # and model_fields_set contains the field
        if self.is_group_video_on is None and "is_group_video_on" in self.model_fields_set:
            _dict['is_group_video_on'] = None

        # set to None if is_hidden (nullable) is None
        # and model_fields_set contains the field
        if self.is_hidden is None and "is_hidden" in self.model_fields_set:
            _dict['is_hidden'] = None

        # set to None if is_hide_vip (nullable) is None
        # and model_fields_set contains the field
        if self.is_hide_vip is None and "is_hide_vip" in self.model_fields_set:
            _dict['is_hide_vip'] = None

        # set to None if is_interests_selected (nullable) is None
        # and model_fields_set contains the field
        if self.is_interests_selected is None and "is_interests_selected" in self.model_fields_set:
            _dict['is_interests_selected'] = None

        # set to None if is_line_connected (nullable) is None
        # and model_fields_set contains the field
        if self.is_line_connected is None and "is_line_connected" in self.model_fields_set:
            _dict['is_line_connected'] = None

        # set to None if is_muted (nullable) is None
        # and model_fields_set contains the field
        if self.is_muted is None and "is_muted" in self.model_fields_set:
            _dict['is_muted'] = None

        # set to None if is_new (nullable) is None
        # and model_fields_set contains the field
        if self.is_new is None and "is_new" in self.model_fields_set:
            _dict['is_new'] = None

        # set to None if is_phone_on (nullable) is None
        # and model_fields_set contains the field
        if self.is_phone_on is None and "is_phone_on" in self.model_fields_set:
            _dict['is_phone_on'] = None

        # set to None if is_private (nullable) is None
        # and model_fields_set contains the field
        if self.is_private is None and "is_private" in self.model_fields_set:
            _dict['is_private'] = None

        # set to None if is_recently_penalized (nullable) is None
        # and model_fields_set contains the field
        if self.is_recently_penalized is None and "is_recently_penalized" in self.model_fields_set:
            _dict['is_recently_penalized'] = None

        # set to None if is_registered (nullable) is None
        # and model_fields_set contains the field
        if self.is_registered is None and "is_registered" in self.model_fields_set:
            _dict['is_registered'] = None

        # set to None if is_two_fa_enabled (nullable) is None
        # and model_fields_set contains the field
        if self.is_two_fa_enabled is None and "is_two_fa_enabled" in self.model_fields_set:
            _dict['is_two_fa_enabled'] = None

        # set to None if is_video_on (nullable) is None
        # and model_fields_set contains the field
        if self.is_video_on is None and "is_video_on" in self.model_fields_set:
            _dict['is_video_on'] = None

        # set to None if is_vip (nullable) is None
        # and model_fields_set contains the field
        if self.is_vip is None and "is_vip" in self.model_fields_set:
            _dict['is_vip'] = None

        # set to None if is_world_id_connected (nullable) is None
        # and model_fields_set contains the field
        if self.is_world_id_connected is None and "is_world_id_connected" in self.model_fields_set:
            _dict['is_world_id_connected'] = None

        # set to None if last_logged_in_at_millis (nullable) is None
        # and model_fields_set contains the field
        if self.last_logged_in_at_millis is None and "last_logged_in_at_millis" in self.model_fields_set:
            _dict['last_logged_in_at_millis'] = None

        # set to None if login_streak_count (nullable) is None
        # and model_fields_set contains the field
        if self.login_streak_count is None and "login_streak_count" in self.model_fields_set:
            _dict['login_streak_count'] = None

        # set to None if masked_email (nullable) is None
        # and model_fields_set contains the field
        if self.masked_email is None and "masked_email" in self.model_fields_set:
            _dict['masked_email'] = None

        # set to None if matching_id (nullable) is None
        # and model_fields_set contains the field
        if self.matching_id is None and "matching_id" in self.model_fields_set:
            _dict['matching_id'] = None

        # set to None if media_count (nullable) is None
        # and model_fields_set contains the field
        if self.media_count is None and "media_count" in self.model_fields_set:
            _dict['media_count'] = None

        # set to None if mobile_number (nullable) is None
        # and model_fields_set contains the field
        if self.mobile_number is None and "mobile_number" in self.model_fields_set:
            _dict['mobile_number'] = None

        # set to None if nickname (nullable) is None
        # and model_fields_set contains the field
        if self.nickname is None and "nickname" in self.model_fields_set:
            _dict['nickname'] = None

        # set to None if online_status (nullable) is None
        # and model_fields_set contains the field
        if self.online_status is None and "online_status" in self.model_fields_set:
            _dict['online_status'] = None

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

        # set to None if profile_icon_frame (nullable) is None
        # and model_fields_set contains the field
        if self.profile_icon_frame is None and "profile_icon_frame" in self.model_fields_set:
            _dict['profile_icon_frame'] = None

        # set to None if profile_icon_frame_thumbnail (nullable) is None
        # and model_fields_set contains the field
        if self.profile_icon_frame_thumbnail is None and "profile_icon_frame_thumbnail" in self.model_fields_set:
            _dict['profile_icon_frame_thumbnail'] = None

        # set to None if profile_icon_thumbnail (nullable) is None
        # and model_fields_set contains the field
        if self.profile_icon_thumbnail is None and "profile_icon_thumbnail" in self.model_fields_set:
            _dict['profile_icon_thumbnail'] = None

        # set to None if review_restriction (nullable) is None
        # and model_fields_set contains the field
        if self.review_restriction is None and "review_restriction" in self.model_fields_set:
            _dict['review_restriction'] = None

        # set to None if reviews_count (nullable) is None
        # and model_fields_set contains the field
        if self.reviews_count is None and "reviews_count" in self.model_fields_set:
            _dict['reviews_count'] = None

        # set to None if title (nullable) is None
        # and model_fields_set contains the field
        if self.title is None and "title" in self.model_fields_set:
            _dict['title'] = None

        # set to None if twitter_id (nullable) is None
        # and model_fields_set contains the field
        if self.twitter_id is None and "twitter_id" in self.model_fields_set:
            _dict['twitter_id'] = None

        # set to None if username (nullable) is None
        # and model_fields_set contains the field
        if self.username is None and "username" in self.model_fields_set:
            _dict['username'] = None

        # set to None if uuid (nullable) is None
        # and model_fields_set contains the field
        if self.uuid is None and "uuid" in self.model_fields_set:
            _dict['uuid'] = None

        # set to None if vip_until_millis (nullable) is None
        # and model_fields_set contains the field
        if self.vip_until_millis is None and "vip_until_millis" in self.model_fields_set:
            _dict['vip_until_millis'] = None

        return _dict

    @classmethod
    def from_dict(cls, obj: Optional[Dict[str, Any]]) -> Optional[Self]:
        """Create an instance of User from a dict"""
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return cls.model_validate(obj)

        _obj = cls.model_validate({
            "biography": obj.get("biography"),
            "birth_date": obj.get("birth_date"),
            "blocking_limit": obj.get("blocking_limit"),
            "connected_by": obj.get("connected_by"),
            "contact_phones": obj.get("contact_phones"),
            "country": obj.get("country"),
            "cover_image": obj.get("cover_image"),
            "cover_image_thumbnail": obj.get("cover_image_thumbnail"),
            "created_at_millis": obj.get("created_at_millis"),
            "followers_count": obj.get("followers_count"),
            "followings_count": obj.get("followings_count"),
            "gender": Gender.from_dict(obj["gender"]) if obj.get("gender") is not None else None,
            "gift_received_counter": obj.get("gift_received_counter"),
            "gifting_ability": GiftingAbility.from_dict(obj["gifting_ability"]) if obj.get("gifting_ability") is not None else None,
            "groups_users_count": obj.get("groups_users_count"),
            "id": obj.get("id"),
            "interests_count": obj.get("interests_count"),
            "is_age_verified": obj.get("is_age_verified"),
            "is_apple_connected": obj.get("is_apple_connected"),
            "is_chat_request_on": obj.get("is_chat_request_on"),
            "is_dangerous": obj.get("is_dangerous"),
            "is_email_confirmed": obj.get("is_email_confirmed"),
            "is_facebook_connected": obj.get("is_facebook_connected"),
            "is_follow_pending": obj.get("is_follow_pending"),
            "is_followed_by": obj.get("is_followed_by"),
            "is_followed_by_pending": obj.get("is_followed_by_pending"),
            "is_following": obj.get("is_following"),
            "is_from_different_generation_and_trusted": obj.get("is_from_different_generation_and_trusted"),
            "is_google_connected": obj.get("is_google_connected"),
            "is_group_phone_on": obj.get("is_group_phone_on"),
            "is_group_video_on": obj.get("is_group_video_on"),
            "is_hidden": obj.get("is_hidden"),
            "is_hide_vip": obj.get("is_hide_vip"),
            "is_interests_selected": obj.get("is_interests_selected"),
            "is_line_connected": obj.get("is_line_connected"),
            "is_muted": obj.get("is_muted"),
            "is_new": obj.get("is_new"),
            "is_phone_on": obj.get("is_phone_on"),
            "is_private": obj.get("is_private"),
            "is_recently_penalized": obj.get("is_recently_penalized"),
            "is_registered": obj.get("is_registered"),
            "is_two_fa_enabled": obj.get("is_two_fa_enabled"),
            "is_video_on": obj.get("is_video_on"),
            "is_vip": obj.get("is_vip"),
            "is_world_id_connected": obj.get("is_world_id_connected"),
            "last_logged_in_at_millis": obj.get("last_logged_in_at_millis"),
            "login_streak_count": obj.get("login_streak_count"),
            "masked_email": obj.get("masked_email"),
            "matching_id": obj.get("matching_id"),
            "media_count": obj.get("media_count"),
            "mobile_number": obj.get("mobile_number"),
            "nickname": obj.get("nickname"),
            "online_status": OnlineStatus.from_dict(obj["online_status"]) if obj.get("online_status") is not None else None,
            "posts_count": obj.get("posts_count"),
            "prefecture": obj.get("prefecture"),
            "profile_icon": obj.get("profile_icon"),
            "profile_icon_frame": obj.get("profile_icon_frame"),
            "profile_icon_frame_thumbnail": obj.get("profile_icon_frame_thumbnail"),
            "profile_icon_thumbnail": obj.get("profile_icon_thumbnail"),
            "review_restriction": ReviewRestriction.from_dict(obj["review_restriction"]) if obj.get("review_restriction") is not None else None,
            "reviews_count": obj.get("reviews_count"),
            "title": Title.from_dict(obj["title"]) if obj.get("title") is not None else None,
            "twitter_id": obj.get("twitter_id"),
            "username": obj.get("username"),
            "uuid": obj.get("uuid"),
            "vip_until_millis": obj.get("vip_until_millis")
        })
        return _obj


