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
from yaylib.models.gift_count import GiftCount
from yaylib.models.group import Group
from yaylib.models.message_tag import MessageTag
from yaylib.models.post_type import PostType
from yaylib.models.realm_conference_call import RealmConferenceCall
from yaylib.models.realm_user import RealmUser
from yaylib.models.shared_url import SharedUrl
from yaylib.models.survey import Survey
from yaylib.models.video import Video
from typing import Optional, Set
from typing_extensions import Self

class Post(BaseModel):
    """
    Post
    """ # noqa: E501
    attachment: Optional[StrictStr] = None
    attachment_2: Optional[StrictStr] = None
    attachment_2_thumbnail: Optional[StrictStr] = None
    attachment_3: Optional[StrictStr] = None
    attachment_3_thumbnail: Optional[StrictStr] = None
    attachment_4: Optional[StrictStr] = None
    attachment_4_thumbnail: Optional[StrictStr] = None
    attachment_5: Optional[StrictStr] = None
    attachment_5_thumbnail: Optional[StrictStr] = None
    attachment_6: Optional[StrictStr] = None
    attachment_6_thumbnail: Optional[StrictStr] = None
    attachment_7: Optional[StrictStr] = None
    attachment_7_thumbnail: Optional[StrictStr] = None
    attachment_8: Optional[StrictStr] = None
    attachment_8_thumbnail: Optional[StrictStr] = None
    attachment_9: Optional[StrictStr] = None
    attachment_9_thumbnail: Optional[StrictStr] = None
    attachment_thumbnail: Optional[StrictStr] = None
    color: Optional[StrictInt] = None
    conference_call: Optional[RealmConferenceCall] = None
    conversation_id: Optional[StrictInt] = None
    created_at: Optional[StrictInt] = None
    edited_at: Optional[StrictInt] = None
    font_size: Optional[StrictInt] = None
    gifts_count: Optional[List[GiftCount]] = None
    group: Optional[Group] = None
    group_id: Optional[StrictInt] = None
    highlighted: Optional[StrictBool] = None
    id: Optional[StrictInt] = None
    in_reply_to: Optional[StrictInt] = None
    in_reply_to_post: Optional[Post] = None
    in_reply_to_post_count: Optional[StrictInt] = None
    is_fail_to_send: Optional[StrictBool] = None
    is_muted: Optional[StrictBool] = None
    liked: Optional[StrictBool] = None
    likers: Optional[List[RealmUser]] = None
    likers_count: Optional[StrictInt] = None
    likes_count: Optional[StrictInt] = None
    mentions: Optional[List[RealmUser]] = None
    message_tags: Optional[List[MessageTag]] = None
    post_type: Optional[PostType] = None
    reported_count: Optional[StrictInt] = None
    repostable: Optional[StrictBool] = None
    reposted: Optional[StrictBool] = None
    reposts_count: Optional[StrictInt] = None
    shareable: Optional[Shareable] = None
    shared_thread: Optional[ThreadInfo] = None
    shared_url: Optional[SharedUrl] = None
    survey: Optional[Survey] = None
    tag: Optional[StrictStr] = None
    text: Optional[StrictStr] = None
    thread: Optional[ThreadInfo] = None
    thread_id: Optional[StrictInt] = None
    total_gifts_count: Optional[StrictInt] = None
    updated_at: Optional[StrictInt] = None
    user: Optional[RealmUser] = None
    videos: Optional[List[Video]] = None
    __properties: ClassVar[List[str]] = ["attachment", "attachment_2", "attachment_2_thumbnail", "attachment_3", "attachment_3_thumbnail", "attachment_4", "attachment_4_thumbnail", "attachment_5", "attachment_5_thumbnail", "attachment_6", "attachment_6_thumbnail", "attachment_7", "attachment_7_thumbnail", "attachment_8", "attachment_8_thumbnail", "attachment_9", "attachment_9_thumbnail", "attachment_thumbnail", "color", "conference_call", "conversation_id", "created_at", "edited_at", "font_size", "gifts_count", "group", "group_id", "highlighted", "id", "in_reply_to", "in_reply_to_post", "in_reply_to_post_count", "is_fail_to_send", "is_muted", "liked", "likers", "likers_count", "likes_count", "mentions", "message_tags", "post_type", "reported_count", "repostable", "reposted", "reposts_count", "shareable", "shared_thread", "shared_url", "survey", "tag", "text", "thread", "thread_id", "total_gifts_count", "updated_at", "user", "videos"]

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
        """Create an instance of Post from a JSON string"""
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
        # override the default output from pydantic by calling `to_dict()` of conference_call
        if self.conference_call:
            _dict['conference_call'] = self.conference_call.to_dict()
        # override the default output from pydantic by calling `to_dict()` of each item in gifts_count (list)
        _items = []
        if self.gifts_count:
            for _item_gifts_count in self.gifts_count:
                if _item_gifts_count:
                    _items.append(_item_gifts_count.to_dict())
            _dict['gifts_count'] = _items
        # override the default output from pydantic by calling `to_dict()` of group
        if self.group:
            _dict['group'] = self.group.to_dict()
        # override the default output from pydantic by calling `to_dict()` of in_reply_to_post
        if self.in_reply_to_post:
            _dict['in_reply_to_post'] = self.in_reply_to_post.to_dict()
        # override the default output from pydantic by calling `to_dict()` of each item in likers (list)
        _items = []
        if self.likers:
            for _item_likers in self.likers:
                if _item_likers:
                    _items.append(_item_likers.to_dict())
            _dict['likers'] = _items
        # override the default output from pydantic by calling `to_dict()` of each item in mentions (list)
        _items = []
        if self.mentions:
            for _item_mentions in self.mentions:
                if _item_mentions:
                    _items.append(_item_mentions.to_dict())
            _dict['mentions'] = _items
        # override the default output from pydantic by calling `to_dict()` of each item in message_tags (list)
        _items = []
        if self.message_tags:
            for _item_message_tags in self.message_tags:
                if _item_message_tags:
                    _items.append(_item_message_tags.to_dict())
            _dict['message_tags'] = _items
        # override the default output from pydantic by calling `to_dict()` of shareable
        if self.shareable:
            _dict['shareable'] = self.shareable.to_dict()
        # override the default output from pydantic by calling `to_dict()` of shared_thread
        if self.shared_thread:
            _dict['shared_thread'] = self.shared_thread.to_dict()
        # override the default output from pydantic by calling `to_dict()` of shared_url
        if self.shared_url:
            _dict['shared_url'] = self.shared_url.to_dict()
        # override the default output from pydantic by calling `to_dict()` of survey
        if self.survey:
            _dict['survey'] = self.survey.to_dict()
        # override the default output from pydantic by calling `to_dict()` of thread
        if self.thread:
            _dict['thread'] = self.thread.to_dict()
        # override the default output from pydantic by calling `to_dict()` of user
        if self.user:
            _dict['user'] = self.user.to_dict()
        # override the default output from pydantic by calling `to_dict()` of each item in videos (list)
        _items = []
        if self.videos:
            for _item_videos in self.videos:
                if _item_videos:
                    _items.append(_item_videos.to_dict())
            _dict['videos'] = _items
        # set to None if attachment (nullable) is None
        # and model_fields_set contains the field
        if self.attachment is None and "attachment" in self.model_fields_set:
            _dict['attachment'] = None

        # set to None if attachment_2 (nullable) is None
        # and model_fields_set contains the field
        if self.attachment_2 is None and "attachment_2" in self.model_fields_set:
            _dict['attachment_2'] = None

        # set to None if attachment_2_thumbnail (nullable) is None
        # and model_fields_set contains the field
        if self.attachment_2_thumbnail is None and "attachment_2_thumbnail" in self.model_fields_set:
            _dict['attachment_2_thumbnail'] = None

        # set to None if attachment_3 (nullable) is None
        # and model_fields_set contains the field
        if self.attachment_3 is None and "attachment_3" in self.model_fields_set:
            _dict['attachment_3'] = None

        # set to None if attachment_3_thumbnail (nullable) is None
        # and model_fields_set contains the field
        if self.attachment_3_thumbnail is None and "attachment_3_thumbnail" in self.model_fields_set:
            _dict['attachment_3_thumbnail'] = None

        # set to None if attachment_4 (nullable) is None
        # and model_fields_set contains the field
        if self.attachment_4 is None and "attachment_4" in self.model_fields_set:
            _dict['attachment_4'] = None

        # set to None if attachment_4_thumbnail (nullable) is None
        # and model_fields_set contains the field
        if self.attachment_4_thumbnail is None and "attachment_4_thumbnail" in self.model_fields_set:
            _dict['attachment_4_thumbnail'] = None

        # set to None if attachment_5 (nullable) is None
        # and model_fields_set contains the field
        if self.attachment_5 is None and "attachment_5" in self.model_fields_set:
            _dict['attachment_5'] = None

        # set to None if attachment_5_thumbnail (nullable) is None
        # and model_fields_set contains the field
        if self.attachment_5_thumbnail is None and "attachment_5_thumbnail" in self.model_fields_set:
            _dict['attachment_5_thumbnail'] = None

        # set to None if attachment_6 (nullable) is None
        # and model_fields_set contains the field
        if self.attachment_6 is None and "attachment_6" in self.model_fields_set:
            _dict['attachment_6'] = None

        # set to None if attachment_6_thumbnail (nullable) is None
        # and model_fields_set contains the field
        if self.attachment_6_thumbnail is None and "attachment_6_thumbnail" in self.model_fields_set:
            _dict['attachment_6_thumbnail'] = None

        # set to None if attachment_7 (nullable) is None
        # and model_fields_set contains the field
        if self.attachment_7 is None and "attachment_7" in self.model_fields_set:
            _dict['attachment_7'] = None

        # set to None if attachment_7_thumbnail (nullable) is None
        # and model_fields_set contains the field
        if self.attachment_7_thumbnail is None and "attachment_7_thumbnail" in self.model_fields_set:
            _dict['attachment_7_thumbnail'] = None

        # set to None if attachment_8 (nullable) is None
        # and model_fields_set contains the field
        if self.attachment_8 is None and "attachment_8" in self.model_fields_set:
            _dict['attachment_8'] = None

        # set to None if attachment_8_thumbnail (nullable) is None
        # and model_fields_set contains the field
        if self.attachment_8_thumbnail is None and "attachment_8_thumbnail" in self.model_fields_set:
            _dict['attachment_8_thumbnail'] = None

        # set to None if attachment_9 (nullable) is None
        # and model_fields_set contains the field
        if self.attachment_9 is None and "attachment_9" in self.model_fields_set:
            _dict['attachment_9'] = None

        # set to None if attachment_9_thumbnail (nullable) is None
        # and model_fields_set contains the field
        if self.attachment_9_thumbnail is None and "attachment_9_thumbnail" in self.model_fields_set:
            _dict['attachment_9_thumbnail'] = None

        # set to None if attachment_thumbnail (nullable) is None
        # and model_fields_set contains the field
        if self.attachment_thumbnail is None and "attachment_thumbnail" in self.model_fields_set:
            _dict['attachment_thumbnail'] = None

        # set to None if color (nullable) is None
        # and model_fields_set contains the field
        if self.color is None and "color" in self.model_fields_set:
            _dict['color'] = None

        # set to None if conference_call (nullable) is None
        # and model_fields_set contains the field
        if self.conference_call is None and "conference_call" in self.model_fields_set:
            _dict['conference_call'] = None

        # set to None if conversation_id (nullable) is None
        # and model_fields_set contains the field
        if self.conversation_id is None and "conversation_id" in self.model_fields_set:
            _dict['conversation_id'] = None

        # set to None if created_at (nullable) is None
        # and model_fields_set contains the field
        if self.created_at is None and "created_at" in self.model_fields_set:
            _dict['created_at'] = None

        # set to None if edited_at (nullable) is None
        # and model_fields_set contains the field
        if self.edited_at is None and "edited_at" in self.model_fields_set:
            _dict['edited_at'] = None

        # set to None if font_size (nullable) is None
        # and model_fields_set contains the field
        if self.font_size is None and "font_size" in self.model_fields_set:
            _dict['font_size'] = None

        # set to None if gifts_count (nullable) is None
        # and model_fields_set contains the field
        if self.gifts_count is None and "gifts_count" in self.model_fields_set:
            _dict['gifts_count'] = None

        # set to None if group (nullable) is None
        # and model_fields_set contains the field
        if self.group is None and "group" in self.model_fields_set:
            _dict['group'] = None

        # set to None if group_id (nullable) is None
        # and model_fields_set contains the field
        if self.group_id is None and "group_id" in self.model_fields_set:
            _dict['group_id'] = None

        # set to None if highlighted (nullable) is None
        # and model_fields_set contains the field
        if self.highlighted is None and "highlighted" in self.model_fields_set:
            _dict['highlighted'] = None

        # set to None if id (nullable) is None
        # and model_fields_set contains the field
        if self.id is None and "id" in self.model_fields_set:
            _dict['id'] = None

        # set to None if in_reply_to (nullable) is None
        # and model_fields_set contains the field
        if self.in_reply_to is None and "in_reply_to" in self.model_fields_set:
            _dict['in_reply_to'] = None

        # set to None if in_reply_to_post (nullable) is None
        # and model_fields_set contains the field
        if self.in_reply_to_post is None and "in_reply_to_post" in self.model_fields_set:
            _dict['in_reply_to_post'] = None

        # set to None if in_reply_to_post_count (nullable) is None
        # and model_fields_set contains the field
        if self.in_reply_to_post_count is None and "in_reply_to_post_count" in self.model_fields_set:
            _dict['in_reply_to_post_count'] = None

        # set to None if is_fail_to_send (nullable) is None
        # and model_fields_set contains the field
        if self.is_fail_to_send is None and "is_fail_to_send" in self.model_fields_set:
            _dict['is_fail_to_send'] = None

        # set to None if is_muted (nullable) is None
        # and model_fields_set contains the field
        if self.is_muted is None and "is_muted" in self.model_fields_set:
            _dict['is_muted'] = None

        # set to None if liked (nullable) is None
        # and model_fields_set contains the field
        if self.liked is None and "liked" in self.model_fields_set:
            _dict['liked'] = None

        # set to None if likers (nullable) is None
        # and model_fields_set contains the field
        if self.likers is None and "likers" in self.model_fields_set:
            _dict['likers'] = None

        # set to None if likers_count (nullable) is None
        # and model_fields_set contains the field
        if self.likers_count is None and "likers_count" in self.model_fields_set:
            _dict['likers_count'] = None

        # set to None if likes_count (nullable) is None
        # and model_fields_set contains the field
        if self.likes_count is None and "likes_count" in self.model_fields_set:
            _dict['likes_count'] = None

        # set to None if mentions (nullable) is None
        # and model_fields_set contains the field
        if self.mentions is None and "mentions" in self.model_fields_set:
            _dict['mentions'] = None

        # set to None if message_tags (nullable) is None
        # and model_fields_set contains the field
        if self.message_tags is None and "message_tags" in self.model_fields_set:
            _dict['message_tags'] = None

        # set to None if post_type (nullable) is None
        # and model_fields_set contains the field
        if self.post_type is None and "post_type" in self.model_fields_set:
            _dict['post_type'] = None

        # set to None if reported_count (nullable) is None
        # and model_fields_set contains the field
        if self.reported_count is None and "reported_count" in self.model_fields_set:
            _dict['reported_count'] = None

        # set to None if repostable (nullable) is None
        # and model_fields_set contains the field
        if self.repostable is None and "repostable" in self.model_fields_set:
            _dict['repostable'] = None

        # set to None if reposted (nullable) is None
        # and model_fields_set contains the field
        if self.reposted is None and "reposted" in self.model_fields_set:
            _dict['reposted'] = None

        # set to None if reposts_count (nullable) is None
        # and model_fields_set contains the field
        if self.reposts_count is None and "reposts_count" in self.model_fields_set:
            _dict['reposts_count'] = None

        # set to None if shareable (nullable) is None
        # and model_fields_set contains the field
        if self.shareable is None and "shareable" in self.model_fields_set:
            _dict['shareable'] = None

        # set to None if shared_thread (nullable) is None
        # and model_fields_set contains the field
        if self.shared_thread is None and "shared_thread" in self.model_fields_set:
            _dict['shared_thread'] = None

        # set to None if shared_url (nullable) is None
        # and model_fields_set contains the field
        if self.shared_url is None and "shared_url" in self.model_fields_set:
            _dict['shared_url'] = None

        # set to None if survey (nullable) is None
        # and model_fields_set contains the field
        if self.survey is None and "survey" in self.model_fields_set:
            _dict['survey'] = None

        # set to None if tag (nullable) is None
        # and model_fields_set contains the field
        if self.tag is None and "tag" in self.model_fields_set:
            _dict['tag'] = None

        # set to None if text (nullable) is None
        # and model_fields_set contains the field
        if self.text is None and "text" in self.model_fields_set:
            _dict['text'] = None

        # set to None if thread (nullable) is None
        # and model_fields_set contains the field
        if self.thread is None and "thread" in self.model_fields_set:
            _dict['thread'] = None

        # set to None if thread_id (nullable) is None
        # and model_fields_set contains the field
        if self.thread_id is None and "thread_id" in self.model_fields_set:
            _dict['thread_id'] = None

        # set to None if total_gifts_count (nullable) is None
        # and model_fields_set contains the field
        if self.total_gifts_count is None and "total_gifts_count" in self.model_fields_set:
            _dict['total_gifts_count'] = None

        # set to None if updated_at (nullable) is None
        # and model_fields_set contains the field
        if self.updated_at is None and "updated_at" in self.model_fields_set:
            _dict['updated_at'] = None

        # set to None if user (nullable) is None
        # and model_fields_set contains the field
        if self.user is None and "user" in self.model_fields_set:
            _dict['user'] = None

        # set to None if videos (nullable) is None
        # and model_fields_set contains the field
        if self.videos is None and "videos" in self.model_fields_set:
            _dict['videos'] = None

        return _dict

    @classmethod
    def from_dict(cls, obj: Optional[Dict[str, Any]]) -> Optional[Self]:
        """Create an instance of Post from a dict"""
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return cls.model_validate(obj)

        _obj = cls.model_validate({
            "attachment": obj.get("attachment"),
            "attachment_2": obj.get("attachment_2"),
            "attachment_2_thumbnail": obj.get("attachment_2_thumbnail"),
            "attachment_3": obj.get("attachment_3"),
            "attachment_3_thumbnail": obj.get("attachment_3_thumbnail"),
            "attachment_4": obj.get("attachment_4"),
            "attachment_4_thumbnail": obj.get("attachment_4_thumbnail"),
            "attachment_5": obj.get("attachment_5"),
            "attachment_5_thumbnail": obj.get("attachment_5_thumbnail"),
            "attachment_6": obj.get("attachment_6"),
            "attachment_6_thumbnail": obj.get("attachment_6_thumbnail"),
            "attachment_7": obj.get("attachment_7"),
            "attachment_7_thumbnail": obj.get("attachment_7_thumbnail"),
            "attachment_8": obj.get("attachment_8"),
            "attachment_8_thumbnail": obj.get("attachment_8_thumbnail"),
            "attachment_9": obj.get("attachment_9"),
            "attachment_9_thumbnail": obj.get("attachment_9_thumbnail"),
            "attachment_thumbnail": obj.get("attachment_thumbnail"),
            "color": obj.get("color"),
            "conference_call": RealmConferenceCall.from_dict(obj["conference_call"]) if obj.get("conference_call") is not None else None,
            "conversation_id": obj.get("conversation_id"),
            "created_at": obj.get("created_at"),
            "edited_at": obj.get("edited_at"),
            "font_size": obj.get("font_size"),
            "gifts_count": [GiftCount.from_dict(_item) for _item in obj["gifts_count"]] if obj.get("gifts_count") is not None else None,
            "group": Group.from_dict(obj["group"]) if obj.get("group") is not None else None,
            "group_id": obj.get("group_id"),
            "highlighted": obj.get("highlighted"),
            "id": obj.get("id"),
            "in_reply_to": obj.get("in_reply_to"),
            "in_reply_to_post": Post.from_dict(obj["in_reply_to_post"]) if obj.get("in_reply_to_post") is not None else None,
            "in_reply_to_post_count": obj.get("in_reply_to_post_count"),
            "is_fail_to_send": obj.get("is_fail_to_send"),
            "is_muted": obj.get("is_muted"),
            "liked": obj.get("liked"),
            "likers": [RealmUser.from_dict(_item) for _item in obj["likers"]] if obj.get("likers") is not None else None,
            "likers_count": obj.get("likers_count"),
            "likes_count": obj.get("likes_count"),
            "mentions": [RealmUser.from_dict(_item) for _item in obj["mentions"]] if obj.get("mentions") is not None else None,
            "message_tags": [MessageTag.from_dict(_item) for _item in obj["message_tags"]] if obj.get("message_tags") is not None else None,
            "post_type": obj.get("post_type"),
            "reported_count": obj.get("reported_count"),
            "repostable": obj.get("repostable"),
            "reposted": obj.get("reposted"),
            "reposts_count": obj.get("reposts_count"),
            "shareable": Shareable.from_dict(obj["shareable"]) if obj.get("shareable") is not None else None,
            "shared_thread": ThreadInfo.from_dict(obj["shared_thread"]) if obj.get("shared_thread") is not None else None,
            "shared_url": SharedUrl.from_dict(obj["shared_url"]) if obj.get("shared_url") is not None else None,
            "survey": Survey.from_dict(obj["survey"]) if obj.get("survey") is not None else None,
            "tag": obj.get("tag"),
            "text": obj.get("text"),
            "thread": ThreadInfo.from_dict(obj["thread"]) if obj.get("thread") is not None else None,
            "thread_id": obj.get("thread_id"),
            "total_gifts_count": obj.get("total_gifts_count"),
            "updated_at": obj.get("updated_at"),
            "user": RealmUser.from_dict(obj["user"]) if obj.get("user") is not None else None,
            "videos": [Video.from_dict(_item) for _item in obj["videos"]] if obj.get("videos") is not None else None
        })
        return _obj

from yaylib.models.shareable import Shareable
from yaylib.models.thread_info import ThreadInfo
# TODO: Rewrite to not use raise_errors
Post.model_rebuild(raise_errors=False)

