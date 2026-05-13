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
from yaylib.models.generation import Generation
from yaylib.models.group_role import GroupRole
from yaylib.models.user import User
from typing import Optional, Set
from typing_extensions import Self

class ModelGroup(BaseModel):
    """
    ModelGroup
    """ # noqa: E501
    allow_members_to_post_media: Optional[StrictBool] = None
    allow_members_to_post_url: Optional[StrictBool] = None
    allow_ownership_transfer: Optional[StrictBool] = None
    allow_thread_creation_by: Optional[GroupRole] = None
    cover_image: Optional[StrictStr] = None
    cover_image_thumbnail: Optional[StrictStr] = None
    description: Optional[StrictStr] = None
    display_icon_thumbnail: Optional[StrictStr] = None
    gender: Optional[Gender] = None
    generation: Optional[Generation] = None
    group_category_id: Optional[StrictInt] = None
    groups_users_count: Optional[StrictInt] = None
    guidelines: Optional[StrictStr] = None
    highlighted_count: Optional[StrictInt] = None
    icon_image: Optional[StrictStr] = None
    icon_image_thumbnail: Optional[StrictStr] = None
    id: Optional[StrictInt] = None
    invited_to_join: Optional[StrictBool] = None
    is_call_timeline_display: Optional[StrictBool] = None
    is_hide_conference_call: Optional[StrictBool] = None
    is_hide_from_game_eight: Optional[StrictBool] = None
    is_hide_reported_posts: Optional[StrictBool] = None
    is_joined: Optional[StrictBool] = None
    is_only_verified_age: Optional[StrictBool] = None
    is_only_verified_phone: Optional[StrictBool] = None
    is_pending: Optional[StrictBool] = None
    is_pinned: Optional[StrictBool] = None
    is_private: Optional[StrictBool] = None
    is_safe_mode: Optional[StrictBool] = None
    is_secret: Optional[StrictBool] = None
    is_seizable: Optional[StrictBool] = None
    is_walkthrough_requested: Optional[StrictBool] = None
    join_status: Optional[Dict[str, Any]] = None
    joined_community_campaign: Optional[StrictBool] = None
    moderator_ids: Optional[List[StrictInt]] = None
    owner: Optional[User] = None
    pending_count: Optional[StrictInt] = None
    pending_deputize_ids: Optional[List[StrictInt]] = None
    pending_transfer_id: Optional[StrictInt] = None
    posts_count: Optional[StrictInt] = None
    related_count: Optional[StrictInt] = None
    seizable_before_millis: Optional[StrictInt] = None
    sub_category_id: Optional[StrictInt] = None
    threads_count: Optional[StrictInt] = None
    title: Optional[Dict[str, Any]] = None
    topic: Optional[StrictStr] = None
    unread_counts: Optional[StrictInt] = None
    unread_threads_count: Optional[StrictInt] = None
    updated_at_seconds: Optional[StrictInt] = None
    views_count: Optional[StrictInt] = None
    __properties: ClassVar[List[str]] = ["allow_members_to_post_media", "allow_members_to_post_url", "allow_ownership_transfer", "allow_thread_creation_by", "cover_image", "cover_image_thumbnail", "description", "display_icon_thumbnail", "gender", "generation", "group_category_id", "groups_users_count", "guidelines", "highlighted_count", "icon_image", "icon_image_thumbnail", "id", "invited_to_join", "is_call_timeline_display", "is_hide_conference_call", "is_hide_from_game_eight", "is_hide_reported_posts", "is_joined", "is_only_verified_age", "is_only_verified_phone", "is_pending", "is_pinned", "is_private", "is_safe_mode", "is_secret", "is_seizable", "is_walkthrough_requested", "join_status", "joined_community_campaign", "moderator_ids", "owner", "pending_count", "pending_deputize_ids", "pending_transfer_id", "posts_count", "related_count", "seizable_before_millis", "sub_category_id", "threads_count", "title", "topic", "unread_counts", "unread_threads_count", "updated_at_seconds", "views_count"]

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
        """Create an instance of ModelGroup from a JSON string"""
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
        # override the default output from pydantic by calling `to_dict()` of allow_thread_creation_by
        if self.allow_thread_creation_by:
            _dict['allow_thread_creation_by'] = self.allow_thread_creation_by.to_dict()
        # override the default output from pydantic by calling `to_dict()` of gender
        if self.gender:
            _dict['gender'] = self.gender.to_dict()
        # override the default output from pydantic by calling `to_dict()` of generation
        if self.generation:
            _dict['generation'] = self.generation.to_dict()
        # override the default output from pydantic by calling `to_dict()` of owner
        if self.owner:
            _dict['owner'] = self.owner.to_dict()
        # set to None if allow_members_to_post_media (nullable) is None
        # and model_fields_set contains the field
        if self.allow_members_to_post_media is None and "allow_members_to_post_media" in self.model_fields_set:
            _dict['allow_members_to_post_media'] = None

        # set to None if allow_members_to_post_url (nullable) is None
        # and model_fields_set contains the field
        if self.allow_members_to_post_url is None and "allow_members_to_post_url" in self.model_fields_set:
            _dict['allow_members_to_post_url'] = None

        # set to None if allow_ownership_transfer (nullable) is None
        # and model_fields_set contains the field
        if self.allow_ownership_transfer is None and "allow_ownership_transfer" in self.model_fields_set:
            _dict['allow_ownership_transfer'] = None

        # set to None if allow_thread_creation_by (nullable) is None
        # and model_fields_set contains the field
        if self.allow_thread_creation_by is None and "allow_thread_creation_by" in self.model_fields_set:
            _dict['allow_thread_creation_by'] = None

        # set to None if cover_image (nullable) is None
        # and model_fields_set contains the field
        if self.cover_image is None and "cover_image" in self.model_fields_set:
            _dict['cover_image'] = None

        # set to None if cover_image_thumbnail (nullable) is None
        # and model_fields_set contains the field
        if self.cover_image_thumbnail is None and "cover_image_thumbnail" in self.model_fields_set:
            _dict['cover_image_thumbnail'] = None

        # set to None if description (nullable) is None
        # and model_fields_set contains the field
        if self.description is None and "description" in self.model_fields_set:
            _dict['description'] = None

        # set to None if display_icon_thumbnail (nullable) is None
        # and model_fields_set contains the field
        if self.display_icon_thumbnail is None and "display_icon_thumbnail" in self.model_fields_set:
            _dict['display_icon_thumbnail'] = None

        # set to None if gender (nullable) is None
        # and model_fields_set contains the field
        if self.gender is None and "gender" in self.model_fields_set:
            _dict['gender'] = None

        # set to None if generation (nullable) is None
        # and model_fields_set contains the field
        if self.generation is None and "generation" in self.model_fields_set:
            _dict['generation'] = None

        # set to None if group_category_id (nullable) is None
        # and model_fields_set contains the field
        if self.group_category_id is None and "group_category_id" in self.model_fields_set:
            _dict['group_category_id'] = None

        # set to None if groups_users_count (nullable) is None
        # and model_fields_set contains the field
        if self.groups_users_count is None and "groups_users_count" in self.model_fields_set:
            _dict['groups_users_count'] = None

        # set to None if guidelines (nullable) is None
        # and model_fields_set contains the field
        if self.guidelines is None and "guidelines" in self.model_fields_set:
            _dict['guidelines'] = None

        # set to None if highlighted_count (nullable) is None
        # and model_fields_set contains the field
        if self.highlighted_count is None and "highlighted_count" in self.model_fields_set:
            _dict['highlighted_count'] = None

        # set to None if icon_image (nullable) is None
        # and model_fields_set contains the field
        if self.icon_image is None and "icon_image" in self.model_fields_set:
            _dict['icon_image'] = None

        # set to None if icon_image_thumbnail (nullable) is None
        # and model_fields_set contains the field
        if self.icon_image_thumbnail is None and "icon_image_thumbnail" in self.model_fields_set:
            _dict['icon_image_thumbnail'] = None

        # set to None if id (nullable) is None
        # and model_fields_set contains the field
        if self.id is None and "id" in self.model_fields_set:
            _dict['id'] = None

        # set to None if invited_to_join (nullable) is None
        # and model_fields_set contains the field
        if self.invited_to_join is None and "invited_to_join" in self.model_fields_set:
            _dict['invited_to_join'] = None

        # set to None if is_call_timeline_display (nullable) is None
        # and model_fields_set contains the field
        if self.is_call_timeline_display is None and "is_call_timeline_display" in self.model_fields_set:
            _dict['is_call_timeline_display'] = None

        # set to None if is_hide_conference_call (nullable) is None
        # and model_fields_set contains the field
        if self.is_hide_conference_call is None and "is_hide_conference_call" in self.model_fields_set:
            _dict['is_hide_conference_call'] = None

        # set to None if is_hide_from_game_eight (nullable) is None
        # and model_fields_set contains the field
        if self.is_hide_from_game_eight is None and "is_hide_from_game_eight" in self.model_fields_set:
            _dict['is_hide_from_game_eight'] = None

        # set to None if is_hide_reported_posts (nullable) is None
        # and model_fields_set contains the field
        if self.is_hide_reported_posts is None and "is_hide_reported_posts" in self.model_fields_set:
            _dict['is_hide_reported_posts'] = None

        # set to None if is_joined (nullable) is None
        # and model_fields_set contains the field
        if self.is_joined is None and "is_joined" in self.model_fields_set:
            _dict['is_joined'] = None

        # set to None if is_only_verified_age (nullable) is None
        # and model_fields_set contains the field
        if self.is_only_verified_age is None and "is_only_verified_age" in self.model_fields_set:
            _dict['is_only_verified_age'] = None

        # set to None if is_only_verified_phone (nullable) is None
        # and model_fields_set contains the field
        if self.is_only_verified_phone is None and "is_only_verified_phone" in self.model_fields_set:
            _dict['is_only_verified_phone'] = None

        # set to None if is_pending (nullable) is None
        # and model_fields_set contains the field
        if self.is_pending is None and "is_pending" in self.model_fields_set:
            _dict['is_pending'] = None

        # set to None if is_pinned (nullable) is None
        # and model_fields_set contains the field
        if self.is_pinned is None and "is_pinned" in self.model_fields_set:
            _dict['is_pinned'] = None

        # set to None if is_private (nullable) is None
        # and model_fields_set contains the field
        if self.is_private is None and "is_private" in self.model_fields_set:
            _dict['is_private'] = None

        # set to None if is_safe_mode (nullable) is None
        # and model_fields_set contains the field
        if self.is_safe_mode is None and "is_safe_mode" in self.model_fields_set:
            _dict['is_safe_mode'] = None

        # set to None if is_secret (nullable) is None
        # and model_fields_set contains the field
        if self.is_secret is None and "is_secret" in self.model_fields_set:
            _dict['is_secret'] = None

        # set to None if is_seizable (nullable) is None
        # and model_fields_set contains the field
        if self.is_seizable is None and "is_seizable" in self.model_fields_set:
            _dict['is_seizable'] = None

        # set to None if is_walkthrough_requested (nullable) is None
        # and model_fields_set contains the field
        if self.is_walkthrough_requested is None and "is_walkthrough_requested" in self.model_fields_set:
            _dict['is_walkthrough_requested'] = None

        # set to None if join_status (nullable) is None
        # and model_fields_set contains the field
        if self.join_status is None and "join_status" in self.model_fields_set:
            _dict['join_status'] = None

        # set to None if joined_community_campaign (nullable) is None
        # and model_fields_set contains the field
        if self.joined_community_campaign is None and "joined_community_campaign" in self.model_fields_set:
            _dict['joined_community_campaign'] = None

        # set to None if moderator_ids (nullable) is None
        # and model_fields_set contains the field
        if self.moderator_ids is None and "moderator_ids" in self.model_fields_set:
            _dict['moderator_ids'] = None

        # set to None if owner (nullable) is None
        # and model_fields_set contains the field
        if self.owner is None and "owner" in self.model_fields_set:
            _dict['owner'] = None

        # set to None if pending_count (nullable) is None
        # and model_fields_set contains the field
        if self.pending_count is None and "pending_count" in self.model_fields_set:
            _dict['pending_count'] = None

        # set to None if pending_deputize_ids (nullable) is None
        # and model_fields_set contains the field
        if self.pending_deputize_ids is None and "pending_deputize_ids" in self.model_fields_set:
            _dict['pending_deputize_ids'] = None

        # set to None if pending_transfer_id (nullable) is None
        # and model_fields_set contains the field
        if self.pending_transfer_id is None and "pending_transfer_id" in self.model_fields_set:
            _dict['pending_transfer_id'] = None

        # set to None if posts_count (nullable) is None
        # and model_fields_set contains the field
        if self.posts_count is None and "posts_count" in self.model_fields_set:
            _dict['posts_count'] = None

        # set to None if related_count (nullable) is None
        # and model_fields_set contains the field
        if self.related_count is None and "related_count" in self.model_fields_set:
            _dict['related_count'] = None

        # set to None if seizable_before_millis (nullable) is None
        # and model_fields_set contains the field
        if self.seizable_before_millis is None and "seizable_before_millis" in self.model_fields_set:
            _dict['seizable_before_millis'] = None

        # set to None if sub_category_id (nullable) is None
        # and model_fields_set contains the field
        if self.sub_category_id is None and "sub_category_id" in self.model_fields_set:
            _dict['sub_category_id'] = None

        # set to None if threads_count (nullable) is None
        # and model_fields_set contains the field
        if self.threads_count is None and "threads_count" in self.model_fields_set:
            _dict['threads_count'] = None

        # set to None if title (nullable) is None
        # and model_fields_set contains the field
        if self.title is None and "title" in self.model_fields_set:
            _dict['title'] = None

        # set to None if topic (nullable) is None
        # and model_fields_set contains the field
        if self.topic is None and "topic" in self.model_fields_set:
            _dict['topic'] = None

        # set to None if unread_counts (nullable) is None
        # and model_fields_set contains the field
        if self.unread_counts is None and "unread_counts" in self.model_fields_set:
            _dict['unread_counts'] = None

        # set to None if unread_threads_count (nullable) is None
        # and model_fields_set contains the field
        if self.unread_threads_count is None and "unread_threads_count" in self.model_fields_set:
            _dict['unread_threads_count'] = None

        # set to None if updated_at_seconds (nullable) is None
        # and model_fields_set contains the field
        if self.updated_at_seconds is None and "updated_at_seconds" in self.model_fields_set:
            _dict['updated_at_seconds'] = None

        # set to None if views_count (nullable) is None
        # and model_fields_set contains the field
        if self.views_count is None and "views_count" in self.model_fields_set:
            _dict['views_count'] = None

        return _dict

    @classmethod
    def from_dict(cls, obj: Optional[Dict[str, Any]]) -> Optional[Self]:
        """Create an instance of ModelGroup from a dict"""
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return cls.model_validate(obj)

        _obj = cls.model_validate({
            "allow_members_to_post_media": obj.get("allow_members_to_post_media"),
            "allow_members_to_post_url": obj.get("allow_members_to_post_url"),
            "allow_ownership_transfer": obj.get("allow_ownership_transfer"),
            "allow_thread_creation_by": GroupRole.from_dict(obj["allow_thread_creation_by"]) if obj.get("allow_thread_creation_by") is not None else None,
            "cover_image": obj.get("cover_image"),
            "cover_image_thumbnail": obj.get("cover_image_thumbnail"),
            "description": obj.get("description"),
            "display_icon_thumbnail": obj.get("display_icon_thumbnail"),
            "gender": Gender.from_dict(obj["gender"]) if obj.get("gender") is not None else None,
            "generation": Generation.from_dict(obj["generation"]) if obj.get("generation") is not None else None,
            "group_category_id": obj.get("group_category_id"),
            "groups_users_count": obj.get("groups_users_count"),
            "guidelines": obj.get("guidelines"),
            "highlighted_count": obj.get("highlighted_count"),
            "icon_image": obj.get("icon_image"),
            "icon_image_thumbnail": obj.get("icon_image_thumbnail"),
            "id": obj.get("id"),
            "invited_to_join": obj.get("invited_to_join"),
            "is_call_timeline_display": obj.get("is_call_timeline_display"),
            "is_hide_conference_call": obj.get("is_hide_conference_call"),
            "is_hide_from_game_eight": obj.get("is_hide_from_game_eight"),
            "is_hide_reported_posts": obj.get("is_hide_reported_posts"),
            "is_joined": obj.get("is_joined"),
            "is_only_verified_age": obj.get("is_only_verified_age"),
            "is_only_verified_phone": obj.get("is_only_verified_phone"),
            "is_pending": obj.get("is_pending"),
            "is_pinned": obj.get("is_pinned"),
            "is_private": obj.get("is_private"),
            "is_safe_mode": obj.get("is_safe_mode"),
            "is_secret": obj.get("is_secret"),
            "is_seizable": obj.get("is_seizable"),
            "is_walkthrough_requested": obj.get("is_walkthrough_requested"),
            "join_status": obj.get("join_status"),
            "joined_community_campaign": obj.get("joined_community_campaign"),
            "moderator_ids": obj.get("moderator_ids"),
            "owner": User.from_dict(obj["owner"]) if obj.get("owner") is not None else None,
            "pending_count": obj.get("pending_count"),
            "pending_deputize_ids": obj.get("pending_deputize_ids"),
            "pending_transfer_id": obj.get("pending_transfer_id"),
            "posts_count": obj.get("posts_count"),
            "related_count": obj.get("related_count"),
            "seizable_before_millis": obj.get("seizable_before_millis"),
            "sub_category_id": obj.get("sub_category_id"),
            "threads_count": obj.get("threads_count"),
            "title": obj.get("title"),
            "topic": obj.get("topic"),
            "unread_counts": obj.get("unread_counts"),
            "unread_threads_count": obj.get("unread_threads_count"),
            "updated_at_seconds": obj.get("updated_at_seconds"),
            "views_count": obj.get("views_count")
        })
        return _obj


