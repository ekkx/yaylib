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

from pydantic import BaseModel, ConfigDict, StrictBool
from typing import Any, ClassVar, Dict, List, Optional
from typing import Optional, Set
from typing_extensions import Self

class Settings(BaseModel):
    """
    Settings
    """ # noqa: E501
    age_restricted_on_review: Optional[StrictBool] = None
    allow_reposts: Optional[StrictBool] = None
    caution_user_chat: Optional[StrictBool] = None
    following_only_call_invite: Optional[StrictBool] = None
    following_only_group_invite: Optional[StrictBool] = None
    following_restricted_on_review: Optional[StrictBool] = None
    hide_active_call: Optional[StrictBool] = None
    hide_gifts_received: Optional[StrictBool] = None
    hide_hot_post: Optional[StrictBool] = None
    hide_on_invitable: Optional[StrictBool] = None
    hide_online_status: Optional[StrictBool] = None
    hide_reply_following_timeline: Optional[StrictBool] = None
    hide_reply_public_timeline: Optional[StrictBool] = None
    hide_vip: Optional[StrictBool] = None
    invisible_on_user_search: Optional[StrictBool] = None
    no_reply_group_timeline: Optional[StrictBool] = None
    notification_birthday_to_followers: Optional[StrictBool] = None
    notification_bulk_call_invite: Optional[StrictBool] = None
    notification_call_invite: Optional[StrictBool] = None
    notification_chat: Optional[StrictBool] = None
    notification_chat_delete: Optional[StrictBool] = None
    notification_contact_friend: Optional[StrictBool] = None
    notification_daily_quest: Optional[StrictBool] = None
    notification_daily_summary: Optional[StrictBool] = None
    notification_follow: Optional[StrictBool] = None
    notification_follow_accept: Optional[StrictBool] = None
    notification_follow_request: Optional[StrictBool] = None
    notification_follower_conference_call: Optional[StrictBool] = None
    notification_follower_create_group: Optional[StrictBool] = None
    notification_following_birthdate_on: Optional[StrictBool] = None
    notification_following_post_after_break: Optional[StrictBool] = None
    notification_followings_in_call: Optional[StrictBool] = None
    notification_footprint: Optional[StrictBool] = None
    notification_gift_receive: Optional[StrictBool] = None
    notification_group_accept: Optional[StrictBool] = None
    notification_group_conference_call: Optional[StrictBool] = None
    notification_group_invite: Optional[StrictBool] = None
    notification_group_join: Optional[StrictBool] = None
    notification_group_message_tag_all: Optional[StrictBool] = None
    notification_group_moderator: Optional[StrictBool] = None
    notification_group_post: Optional[StrictBool] = None
    notification_group_request: Optional[StrictBool] = None
    notification_hima_now: Optional[StrictBool] = None
    notification_invitee_create_account: Optional[StrictBool] = None
    notification_latest_news: Optional[StrictBool] = None
    notification_like: Optional[StrictBool] = None
    notification_message_tag: Optional[StrictBool] = None
    notification_popular_post: Optional[StrictBool] = None
    notification_profile_screenshot: Optional[StrictBool] = None
    notification_reply: Optional[StrictBool] = None
    notification_repost: Optional[StrictBool] = None
    notification_review: Optional[StrictBool] = None
    notification_security_warning: Optional[StrictBool] = None
    notification_twitter_friend: Optional[StrictBool] = None
    privacy_mode: Optional[StrictBool] = None
    private_post: Optional[StrictBool] = None
    private_user_timeline: Optional[StrictBool] = None
    show_total_gifts_received_count: Optional[StrictBool] = None
    vip_invisible_footprint_mode: Optional[StrictBool] = None
    visible_on_sns_friend_recommendation: Optional[StrictBool] = None
    __properties: ClassVar[List[str]] = ["age_restricted_on_review", "allow_reposts", "caution_user_chat", "following_only_call_invite", "following_only_group_invite", "following_restricted_on_review", "hide_active_call", "hide_gifts_received", "hide_hot_post", "hide_on_invitable", "hide_online_status", "hide_reply_following_timeline", "hide_reply_public_timeline", "hide_vip", "invisible_on_user_search", "no_reply_group_timeline", "notification_birthday_to_followers", "notification_bulk_call_invite", "notification_call_invite", "notification_chat", "notification_chat_delete", "notification_contact_friend", "notification_daily_quest", "notification_daily_summary", "notification_follow", "notification_follow_accept", "notification_follow_request", "notification_follower_conference_call", "notification_follower_create_group", "notification_following_birthdate_on", "notification_following_post_after_break", "notification_followings_in_call", "notification_footprint", "notification_gift_receive", "notification_group_accept", "notification_group_conference_call", "notification_group_invite", "notification_group_join", "notification_group_message_tag_all", "notification_group_moderator", "notification_group_post", "notification_group_request", "notification_hima_now", "notification_invitee_create_account", "notification_latest_news", "notification_like", "notification_message_tag", "notification_popular_post", "notification_profile_screenshot", "notification_reply", "notification_repost", "notification_review", "notification_security_warning", "notification_twitter_friend", "privacy_mode", "private_post", "private_user_timeline", "show_total_gifts_received_count", "vip_invisible_footprint_mode", "visible_on_sns_friend_recommendation"]

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
        """Create an instance of Settings from a JSON string"""
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
        # set to None if age_restricted_on_review (nullable) is None
        # and model_fields_set contains the field
        if self.age_restricted_on_review is None and "age_restricted_on_review" in self.model_fields_set:
            _dict['age_restricted_on_review'] = None

        # set to None if allow_reposts (nullable) is None
        # and model_fields_set contains the field
        if self.allow_reposts is None and "allow_reposts" in self.model_fields_set:
            _dict['allow_reposts'] = None

        # set to None if caution_user_chat (nullable) is None
        # and model_fields_set contains the field
        if self.caution_user_chat is None and "caution_user_chat" in self.model_fields_set:
            _dict['caution_user_chat'] = None

        # set to None if following_only_call_invite (nullable) is None
        # and model_fields_set contains the field
        if self.following_only_call_invite is None and "following_only_call_invite" in self.model_fields_set:
            _dict['following_only_call_invite'] = None

        # set to None if following_only_group_invite (nullable) is None
        # and model_fields_set contains the field
        if self.following_only_group_invite is None and "following_only_group_invite" in self.model_fields_set:
            _dict['following_only_group_invite'] = None

        # set to None if following_restricted_on_review (nullable) is None
        # and model_fields_set contains the field
        if self.following_restricted_on_review is None and "following_restricted_on_review" in self.model_fields_set:
            _dict['following_restricted_on_review'] = None

        # set to None if hide_active_call (nullable) is None
        # and model_fields_set contains the field
        if self.hide_active_call is None and "hide_active_call" in self.model_fields_set:
            _dict['hide_active_call'] = None

        # set to None if hide_gifts_received (nullable) is None
        # and model_fields_set contains the field
        if self.hide_gifts_received is None and "hide_gifts_received" in self.model_fields_set:
            _dict['hide_gifts_received'] = None

        # set to None if hide_hot_post (nullable) is None
        # and model_fields_set contains the field
        if self.hide_hot_post is None and "hide_hot_post" in self.model_fields_set:
            _dict['hide_hot_post'] = None

        # set to None if hide_on_invitable (nullable) is None
        # and model_fields_set contains the field
        if self.hide_on_invitable is None and "hide_on_invitable" in self.model_fields_set:
            _dict['hide_on_invitable'] = None

        # set to None if hide_online_status (nullable) is None
        # and model_fields_set contains the field
        if self.hide_online_status is None and "hide_online_status" in self.model_fields_set:
            _dict['hide_online_status'] = None

        # set to None if hide_reply_following_timeline (nullable) is None
        # and model_fields_set contains the field
        if self.hide_reply_following_timeline is None and "hide_reply_following_timeline" in self.model_fields_set:
            _dict['hide_reply_following_timeline'] = None

        # set to None if hide_reply_public_timeline (nullable) is None
        # and model_fields_set contains the field
        if self.hide_reply_public_timeline is None and "hide_reply_public_timeline" in self.model_fields_set:
            _dict['hide_reply_public_timeline'] = None

        # set to None if hide_vip (nullable) is None
        # and model_fields_set contains the field
        if self.hide_vip is None and "hide_vip" in self.model_fields_set:
            _dict['hide_vip'] = None

        # set to None if invisible_on_user_search (nullable) is None
        # and model_fields_set contains the field
        if self.invisible_on_user_search is None and "invisible_on_user_search" in self.model_fields_set:
            _dict['invisible_on_user_search'] = None

        # set to None if no_reply_group_timeline (nullable) is None
        # and model_fields_set contains the field
        if self.no_reply_group_timeline is None and "no_reply_group_timeline" in self.model_fields_set:
            _dict['no_reply_group_timeline'] = None

        # set to None if notification_birthday_to_followers (nullable) is None
        # and model_fields_set contains the field
        if self.notification_birthday_to_followers is None and "notification_birthday_to_followers" in self.model_fields_set:
            _dict['notification_birthday_to_followers'] = None

        # set to None if notification_bulk_call_invite (nullable) is None
        # and model_fields_set contains the field
        if self.notification_bulk_call_invite is None and "notification_bulk_call_invite" in self.model_fields_set:
            _dict['notification_bulk_call_invite'] = None

        # set to None if notification_call_invite (nullable) is None
        # and model_fields_set contains the field
        if self.notification_call_invite is None and "notification_call_invite" in self.model_fields_set:
            _dict['notification_call_invite'] = None

        # set to None if notification_chat (nullable) is None
        # and model_fields_set contains the field
        if self.notification_chat is None and "notification_chat" in self.model_fields_set:
            _dict['notification_chat'] = None

        # set to None if notification_chat_delete (nullable) is None
        # and model_fields_set contains the field
        if self.notification_chat_delete is None and "notification_chat_delete" in self.model_fields_set:
            _dict['notification_chat_delete'] = None

        # set to None if notification_contact_friend (nullable) is None
        # and model_fields_set contains the field
        if self.notification_contact_friend is None and "notification_contact_friend" in self.model_fields_set:
            _dict['notification_contact_friend'] = None

        # set to None if notification_daily_quest (nullable) is None
        # and model_fields_set contains the field
        if self.notification_daily_quest is None and "notification_daily_quest" in self.model_fields_set:
            _dict['notification_daily_quest'] = None

        # set to None if notification_daily_summary (nullable) is None
        # and model_fields_set contains the field
        if self.notification_daily_summary is None and "notification_daily_summary" in self.model_fields_set:
            _dict['notification_daily_summary'] = None

        # set to None if notification_follow (nullable) is None
        # and model_fields_set contains the field
        if self.notification_follow is None and "notification_follow" in self.model_fields_set:
            _dict['notification_follow'] = None

        # set to None if notification_follow_accept (nullable) is None
        # and model_fields_set contains the field
        if self.notification_follow_accept is None and "notification_follow_accept" in self.model_fields_set:
            _dict['notification_follow_accept'] = None

        # set to None if notification_follow_request (nullable) is None
        # and model_fields_set contains the field
        if self.notification_follow_request is None and "notification_follow_request" in self.model_fields_set:
            _dict['notification_follow_request'] = None

        # set to None if notification_follower_conference_call (nullable) is None
        # and model_fields_set contains the field
        if self.notification_follower_conference_call is None and "notification_follower_conference_call" in self.model_fields_set:
            _dict['notification_follower_conference_call'] = None

        # set to None if notification_follower_create_group (nullable) is None
        # and model_fields_set contains the field
        if self.notification_follower_create_group is None and "notification_follower_create_group" in self.model_fields_set:
            _dict['notification_follower_create_group'] = None

        # set to None if notification_following_birthdate_on (nullable) is None
        # and model_fields_set contains the field
        if self.notification_following_birthdate_on is None and "notification_following_birthdate_on" in self.model_fields_set:
            _dict['notification_following_birthdate_on'] = None

        # set to None if notification_following_post_after_break (nullable) is None
        # and model_fields_set contains the field
        if self.notification_following_post_after_break is None and "notification_following_post_after_break" in self.model_fields_set:
            _dict['notification_following_post_after_break'] = None

        # set to None if notification_followings_in_call (nullable) is None
        # and model_fields_set contains the field
        if self.notification_followings_in_call is None and "notification_followings_in_call" in self.model_fields_set:
            _dict['notification_followings_in_call'] = None

        # set to None if notification_footprint (nullable) is None
        # and model_fields_set contains the field
        if self.notification_footprint is None and "notification_footprint" in self.model_fields_set:
            _dict['notification_footprint'] = None

        # set to None if notification_gift_receive (nullable) is None
        # and model_fields_set contains the field
        if self.notification_gift_receive is None and "notification_gift_receive" in self.model_fields_set:
            _dict['notification_gift_receive'] = None

        # set to None if notification_group_accept (nullable) is None
        # and model_fields_set contains the field
        if self.notification_group_accept is None and "notification_group_accept" in self.model_fields_set:
            _dict['notification_group_accept'] = None

        # set to None if notification_group_conference_call (nullable) is None
        # and model_fields_set contains the field
        if self.notification_group_conference_call is None and "notification_group_conference_call" in self.model_fields_set:
            _dict['notification_group_conference_call'] = None

        # set to None if notification_group_invite (nullable) is None
        # and model_fields_set contains the field
        if self.notification_group_invite is None and "notification_group_invite" in self.model_fields_set:
            _dict['notification_group_invite'] = None

        # set to None if notification_group_join (nullable) is None
        # and model_fields_set contains the field
        if self.notification_group_join is None and "notification_group_join" in self.model_fields_set:
            _dict['notification_group_join'] = None

        # set to None if notification_group_message_tag_all (nullable) is None
        # and model_fields_set contains the field
        if self.notification_group_message_tag_all is None and "notification_group_message_tag_all" in self.model_fields_set:
            _dict['notification_group_message_tag_all'] = None

        # set to None if notification_group_moderator (nullable) is None
        # and model_fields_set contains the field
        if self.notification_group_moderator is None and "notification_group_moderator" in self.model_fields_set:
            _dict['notification_group_moderator'] = None

        # set to None if notification_group_post (nullable) is None
        # and model_fields_set contains the field
        if self.notification_group_post is None and "notification_group_post" in self.model_fields_set:
            _dict['notification_group_post'] = None

        # set to None if notification_group_request (nullable) is None
        # and model_fields_set contains the field
        if self.notification_group_request is None and "notification_group_request" in self.model_fields_set:
            _dict['notification_group_request'] = None

        # set to None if notification_hima_now (nullable) is None
        # and model_fields_set contains the field
        if self.notification_hima_now is None and "notification_hima_now" in self.model_fields_set:
            _dict['notification_hima_now'] = None

        # set to None if notification_invitee_create_account (nullable) is None
        # and model_fields_set contains the field
        if self.notification_invitee_create_account is None and "notification_invitee_create_account" in self.model_fields_set:
            _dict['notification_invitee_create_account'] = None

        # set to None if notification_latest_news (nullable) is None
        # and model_fields_set contains the field
        if self.notification_latest_news is None and "notification_latest_news" in self.model_fields_set:
            _dict['notification_latest_news'] = None

        # set to None if notification_like (nullable) is None
        # and model_fields_set contains the field
        if self.notification_like is None and "notification_like" in self.model_fields_set:
            _dict['notification_like'] = None

        # set to None if notification_message_tag (nullable) is None
        # and model_fields_set contains the field
        if self.notification_message_tag is None and "notification_message_tag" in self.model_fields_set:
            _dict['notification_message_tag'] = None

        # set to None if notification_popular_post (nullable) is None
        # and model_fields_set contains the field
        if self.notification_popular_post is None and "notification_popular_post" in self.model_fields_set:
            _dict['notification_popular_post'] = None

        # set to None if notification_profile_screenshot (nullable) is None
        # and model_fields_set contains the field
        if self.notification_profile_screenshot is None and "notification_profile_screenshot" in self.model_fields_set:
            _dict['notification_profile_screenshot'] = None

        # set to None if notification_reply (nullable) is None
        # and model_fields_set contains the field
        if self.notification_reply is None and "notification_reply" in self.model_fields_set:
            _dict['notification_reply'] = None

        # set to None if notification_repost (nullable) is None
        # and model_fields_set contains the field
        if self.notification_repost is None and "notification_repost" in self.model_fields_set:
            _dict['notification_repost'] = None

        # set to None if notification_review (nullable) is None
        # and model_fields_set contains the field
        if self.notification_review is None and "notification_review" in self.model_fields_set:
            _dict['notification_review'] = None

        # set to None if notification_security_warning (nullable) is None
        # and model_fields_set contains the field
        if self.notification_security_warning is None and "notification_security_warning" in self.model_fields_set:
            _dict['notification_security_warning'] = None

        # set to None if notification_twitter_friend (nullable) is None
        # and model_fields_set contains the field
        if self.notification_twitter_friend is None and "notification_twitter_friend" in self.model_fields_set:
            _dict['notification_twitter_friend'] = None

        # set to None if privacy_mode (nullable) is None
        # and model_fields_set contains the field
        if self.privacy_mode is None and "privacy_mode" in self.model_fields_set:
            _dict['privacy_mode'] = None

        # set to None if private_post (nullable) is None
        # and model_fields_set contains the field
        if self.private_post is None and "private_post" in self.model_fields_set:
            _dict['private_post'] = None

        # set to None if private_user_timeline (nullable) is None
        # and model_fields_set contains the field
        if self.private_user_timeline is None and "private_user_timeline" in self.model_fields_set:
            _dict['private_user_timeline'] = None

        # set to None if show_total_gifts_received_count (nullable) is None
        # and model_fields_set contains the field
        if self.show_total_gifts_received_count is None and "show_total_gifts_received_count" in self.model_fields_set:
            _dict['show_total_gifts_received_count'] = None

        # set to None if vip_invisible_footprint_mode (nullable) is None
        # and model_fields_set contains the field
        if self.vip_invisible_footprint_mode is None and "vip_invisible_footprint_mode" in self.model_fields_set:
            _dict['vip_invisible_footprint_mode'] = None

        # set to None if visible_on_sns_friend_recommendation (nullable) is None
        # and model_fields_set contains the field
        if self.visible_on_sns_friend_recommendation is None and "visible_on_sns_friend_recommendation" in self.model_fields_set:
            _dict['visible_on_sns_friend_recommendation'] = None

        return _dict

    @classmethod
    def from_dict(cls, obj: Optional[Dict[str, Any]]) -> Optional[Self]:
        """Create an instance of Settings from a dict"""
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return cls.model_validate(obj)

        _obj = cls.model_validate({
            "age_restricted_on_review": obj.get("age_restricted_on_review"),
            "allow_reposts": obj.get("allow_reposts"),
            "caution_user_chat": obj.get("caution_user_chat"),
            "following_only_call_invite": obj.get("following_only_call_invite"),
            "following_only_group_invite": obj.get("following_only_group_invite"),
            "following_restricted_on_review": obj.get("following_restricted_on_review"),
            "hide_active_call": obj.get("hide_active_call"),
            "hide_gifts_received": obj.get("hide_gifts_received"),
            "hide_hot_post": obj.get("hide_hot_post"),
            "hide_on_invitable": obj.get("hide_on_invitable"),
            "hide_online_status": obj.get("hide_online_status"),
            "hide_reply_following_timeline": obj.get("hide_reply_following_timeline"),
            "hide_reply_public_timeline": obj.get("hide_reply_public_timeline"),
            "hide_vip": obj.get("hide_vip"),
            "invisible_on_user_search": obj.get("invisible_on_user_search"),
            "no_reply_group_timeline": obj.get("no_reply_group_timeline"),
            "notification_birthday_to_followers": obj.get("notification_birthday_to_followers"),
            "notification_bulk_call_invite": obj.get("notification_bulk_call_invite"),
            "notification_call_invite": obj.get("notification_call_invite"),
            "notification_chat": obj.get("notification_chat"),
            "notification_chat_delete": obj.get("notification_chat_delete"),
            "notification_contact_friend": obj.get("notification_contact_friend"),
            "notification_daily_quest": obj.get("notification_daily_quest"),
            "notification_daily_summary": obj.get("notification_daily_summary"),
            "notification_follow": obj.get("notification_follow"),
            "notification_follow_accept": obj.get("notification_follow_accept"),
            "notification_follow_request": obj.get("notification_follow_request"),
            "notification_follower_conference_call": obj.get("notification_follower_conference_call"),
            "notification_follower_create_group": obj.get("notification_follower_create_group"),
            "notification_following_birthdate_on": obj.get("notification_following_birthdate_on"),
            "notification_following_post_after_break": obj.get("notification_following_post_after_break"),
            "notification_followings_in_call": obj.get("notification_followings_in_call"),
            "notification_footprint": obj.get("notification_footprint"),
            "notification_gift_receive": obj.get("notification_gift_receive"),
            "notification_group_accept": obj.get("notification_group_accept"),
            "notification_group_conference_call": obj.get("notification_group_conference_call"),
            "notification_group_invite": obj.get("notification_group_invite"),
            "notification_group_join": obj.get("notification_group_join"),
            "notification_group_message_tag_all": obj.get("notification_group_message_tag_all"),
            "notification_group_moderator": obj.get("notification_group_moderator"),
            "notification_group_post": obj.get("notification_group_post"),
            "notification_group_request": obj.get("notification_group_request"),
            "notification_hima_now": obj.get("notification_hima_now"),
            "notification_invitee_create_account": obj.get("notification_invitee_create_account"),
            "notification_latest_news": obj.get("notification_latest_news"),
            "notification_like": obj.get("notification_like"),
            "notification_message_tag": obj.get("notification_message_tag"),
            "notification_popular_post": obj.get("notification_popular_post"),
            "notification_profile_screenshot": obj.get("notification_profile_screenshot"),
            "notification_reply": obj.get("notification_reply"),
            "notification_repost": obj.get("notification_repost"),
            "notification_review": obj.get("notification_review"),
            "notification_security_warning": obj.get("notification_security_warning"),
            "notification_twitter_friend": obj.get("notification_twitter_friend"),
            "privacy_mode": obj.get("privacy_mode"),
            "private_post": obj.get("private_post"),
            "private_user_timeline": obj.get("private_user_timeline"),
            "show_total_gifts_received_count": obj.get("show_total_gifts_received_count"),
            "vip_invisible_footprint_mode": obj.get("vip_invisible_footprint_mode"),
            "visible_on_sns_friend_recommendation": obj.get("visible_on_sns_friend_recommendation")
        })
        return _obj


