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
from yaylib.models.realm_user import RealmUser
from typing import Optional, Set
from typing_extensions import Self

class ThreadInfo(BaseModel):
    """
    ThreadInfo
    """ # noqa: E501
    created_at: Optional[StrictInt] = None
    id: Optional[StrictInt] = None
    is_joined: Optional[StrictBool] = None
    last_post: Optional[Post] = None
    new_updates: Optional[StrictBool] = None
    owner: Optional[RealmUser] = None
    posts_count: Optional[StrictInt] = None
    thread_icon: Optional[StrictStr] = None
    title: Optional[StrictStr] = None
    unread_count: Optional[StrictInt] = None
    updated_at: Optional[StrictInt] = None
    __properties: ClassVar[List[str]] = ["created_at", "id", "is_joined", "last_post", "new_updates", "owner", "posts_count", "thread_icon", "title", "unread_count", "updated_at"]

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
        """Create an instance of ThreadInfo from a JSON string"""
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
        # override the default output from pydantic by calling `to_dict()` of last_post
        if self.last_post:
            _dict['last_post'] = self.last_post.to_dict()
        # override the default output from pydantic by calling `to_dict()` of owner
        if self.owner:
            _dict['owner'] = self.owner.to_dict()
        # set to None if created_at (nullable) is None
        # and model_fields_set contains the field
        if self.created_at is None and "created_at" in self.model_fields_set:
            _dict['created_at'] = None

        # set to None if id (nullable) is None
        # and model_fields_set contains the field
        if self.id is None and "id" in self.model_fields_set:
            _dict['id'] = None

        # set to None if is_joined (nullable) is None
        # and model_fields_set contains the field
        if self.is_joined is None and "is_joined" in self.model_fields_set:
            _dict['is_joined'] = None

        # set to None if last_post (nullable) is None
        # and model_fields_set contains the field
        if self.last_post is None and "last_post" in self.model_fields_set:
            _dict['last_post'] = None

        # set to None if new_updates (nullable) is None
        # and model_fields_set contains the field
        if self.new_updates is None and "new_updates" in self.model_fields_set:
            _dict['new_updates'] = None

        # set to None if owner (nullable) is None
        # and model_fields_set contains the field
        if self.owner is None and "owner" in self.model_fields_set:
            _dict['owner'] = None

        # set to None if posts_count (nullable) is None
        # and model_fields_set contains the field
        if self.posts_count is None and "posts_count" in self.model_fields_set:
            _dict['posts_count'] = None

        # set to None if thread_icon (nullable) is None
        # and model_fields_set contains the field
        if self.thread_icon is None and "thread_icon" in self.model_fields_set:
            _dict['thread_icon'] = None

        # set to None if title (nullable) is None
        # and model_fields_set contains the field
        if self.title is None and "title" in self.model_fields_set:
            _dict['title'] = None

        # set to None if unread_count (nullable) is None
        # and model_fields_set contains the field
        if self.unread_count is None and "unread_count" in self.model_fields_set:
            _dict['unread_count'] = None

        # set to None if updated_at (nullable) is None
        # and model_fields_set contains the field
        if self.updated_at is None and "updated_at" in self.model_fields_set:
            _dict['updated_at'] = None

        return _dict

    @classmethod
    def from_dict(cls, obj: Optional[Dict[str, Any]]) -> Optional[Self]:
        """Create an instance of ThreadInfo from a dict"""
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return cls.model_validate(obj)

        _obj = cls.model_validate({
            "created_at": obj.get("created_at"),
            "id": obj.get("id"),
            "is_joined": obj.get("is_joined"),
            "last_post": Post.from_dict(obj["last_post"]) if obj.get("last_post") is not None else None,
            "new_updates": obj.get("new_updates"),
            "owner": RealmUser.from_dict(obj["owner"]) if obj.get("owner") is not None else None,
            "posts_count": obj.get("posts_count"),
            "thread_icon": obj.get("thread_icon"),
            "title": obj.get("title"),
            "unread_count": obj.get("unread_count"),
            "updated_at": obj.get("updated_at")
        })
        return _obj

from yaylib.models.post import Post
# TODO: Rewrite to not use raise_errors
ThreadInfo.model_rebuild(raise_errors=False)

