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
from yaylib.models.gift import Gift
from yaylib.models.model_group import ModelGroup
from yaylib.models.model_post import ModelPost
from yaylib.models.user import User
from typing import Optional, Set
from typing_extensions import Self

class ModelActivity(BaseModel):
    """
    ModelActivity
    """ # noqa: E501
    birthday_users: Optional[List[User]] = None
    birthday_users_count: Optional[StrictInt] = None
    body: Optional[StrictStr] = None
    content_preview: Optional[StrictStr] = None
    created_at_millis: Optional[StrictInt] = None
    empl_amount: Optional[StrictInt] = None
    followers: Optional[List[User]] = None
    followers_count: Optional[StrictInt] = None
    from_post: Optional[ModelPost] = None
    from_post_ids: Optional[List[StrictInt]] = None
    gift_item: Optional[Gift] = None
    group: Optional[ModelGroup] = None
    group_topic: Optional[StrictStr] = None
    id: Optional[StrictInt] = None
    is_anonymous: Optional[StrictBool] = None
    is_bulk_invitation: Optional[StrictBool] = None
    title: Optional[StrictStr] = None
    to_post: Optional[ModelPost] = None
    type: Optional[Dict[str, Any]] = None
    url: Optional[StrictStr] = None
    user: Optional[User] = None
    vip_reward: Optional[StrictInt] = None
    __properties: ClassVar[List[str]] = ["birthday_users", "birthday_users_count", "body", "content_preview", "created_at_millis", "empl_amount", "followers", "followers_count", "from_post", "from_post_ids", "gift_item", "group", "group_topic", "id", "is_anonymous", "is_bulk_invitation", "title", "to_post", "type", "url", "user", "vip_reward"]

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
        """Create an instance of ModelActivity from a JSON string"""
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
        # override the default output from pydantic by calling `to_dict()` of each item in birthday_users (list)
        _items = []
        if self.birthday_users:
            for _item_birthday_users in self.birthday_users:
                if _item_birthday_users:
                    _items.append(_item_birthday_users.to_dict())
            _dict['birthday_users'] = _items
        # override the default output from pydantic by calling `to_dict()` of each item in followers (list)
        _items = []
        if self.followers:
            for _item_followers in self.followers:
                if _item_followers:
                    _items.append(_item_followers.to_dict())
            _dict['followers'] = _items
        # override the default output from pydantic by calling `to_dict()` of from_post
        if self.from_post:
            _dict['from_post'] = self.from_post.to_dict()
        # override the default output from pydantic by calling `to_dict()` of gift_item
        if self.gift_item:
            _dict['gift_item'] = self.gift_item.to_dict()
        # override the default output from pydantic by calling `to_dict()` of group
        if self.group:
            _dict['group'] = self.group.to_dict()
        # override the default output from pydantic by calling `to_dict()` of to_post
        if self.to_post:
            _dict['to_post'] = self.to_post.to_dict()
        # override the default output from pydantic by calling `to_dict()` of user
        if self.user:
            _dict['user'] = self.user.to_dict()
        # set to None if birthday_users (nullable) is None
        # and model_fields_set contains the field
        if self.birthday_users is None and "birthday_users" in self.model_fields_set:
            _dict['birthday_users'] = None

        # set to None if birthday_users_count (nullable) is None
        # and model_fields_set contains the field
        if self.birthday_users_count is None and "birthday_users_count" in self.model_fields_set:
            _dict['birthday_users_count'] = None

        # set to None if body (nullable) is None
        # and model_fields_set contains the field
        if self.body is None and "body" in self.model_fields_set:
            _dict['body'] = None

        # set to None if content_preview (nullable) is None
        # and model_fields_set contains the field
        if self.content_preview is None and "content_preview" in self.model_fields_set:
            _dict['content_preview'] = None

        # set to None if created_at_millis (nullable) is None
        # and model_fields_set contains the field
        if self.created_at_millis is None and "created_at_millis" in self.model_fields_set:
            _dict['created_at_millis'] = None

        # set to None if empl_amount (nullable) is None
        # and model_fields_set contains the field
        if self.empl_amount is None and "empl_amount" in self.model_fields_set:
            _dict['empl_amount'] = None

        # set to None if followers (nullable) is None
        # and model_fields_set contains the field
        if self.followers is None and "followers" in self.model_fields_set:
            _dict['followers'] = None

        # set to None if followers_count (nullable) is None
        # and model_fields_set contains the field
        if self.followers_count is None and "followers_count" in self.model_fields_set:
            _dict['followers_count'] = None

        # set to None if from_post (nullable) is None
        # and model_fields_set contains the field
        if self.from_post is None and "from_post" in self.model_fields_set:
            _dict['from_post'] = None

        # set to None if from_post_ids (nullable) is None
        # and model_fields_set contains the field
        if self.from_post_ids is None and "from_post_ids" in self.model_fields_set:
            _dict['from_post_ids'] = None

        # set to None if gift_item (nullable) is None
        # and model_fields_set contains the field
        if self.gift_item is None and "gift_item" in self.model_fields_set:
            _dict['gift_item'] = None

        # set to None if group (nullable) is None
        # and model_fields_set contains the field
        if self.group is None and "group" in self.model_fields_set:
            _dict['group'] = None

        # set to None if group_topic (nullable) is None
        # and model_fields_set contains the field
        if self.group_topic is None and "group_topic" in self.model_fields_set:
            _dict['group_topic'] = None

        # set to None if id (nullable) is None
        # and model_fields_set contains the field
        if self.id is None and "id" in self.model_fields_set:
            _dict['id'] = None

        # set to None if is_anonymous (nullable) is None
        # and model_fields_set contains the field
        if self.is_anonymous is None and "is_anonymous" in self.model_fields_set:
            _dict['is_anonymous'] = None

        # set to None if is_bulk_invitation (nullable) is None
        # and model_fields_set contains the field
        if self.is_bulk_invitation is None and "is_bulk_invitation" in self.model_fields_set:
            _dict['is_bulk_invitation'] = None

        # set to None if title (nullable) is None
        # and model_fields_set contains the field
        if self.title is None and "title" in self.model_fields_set:
            _dict['title'] = None

        # set to None if to_post (nullable) is None
        # and model_fields_set contains the field
        if self.to_post is None and "to_post" in self.model_fields_set:
            _dict['to_post'] = None

        # set to None if type (nullable) is None
        # and model_fields_set contains the field
        if self.type is None and "type" in self.model_fields_set:
            _dict['type'] = None

        # set to None if url (nullable) is None
        # and model_fields_set contains the field
        if self.url is None and "url" in self.model_fields_set:
            _dict['url'] = None

        # set to None if user (nullable) is None
        # and model_fields_set contains the field
        if self.user is None and "user" in self.model_fields_set:
            _dict['user'] = None

        # set to None if vip_reward (nullable) is None
        # and model_fields_set contains the field
        if self.vip_reward is None and "vip_reward" in self.model_fields_set:
            _dict['vip_reward'] = None

        return _dict

    @classmethod
    def from_dict(cls, obj: Optional[Dict[str, Any]]) -> Optional[Self]:
        """Create an instance of ModelActivity from a dict"""
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return cls.model_validate(obj)

        _obj = cls.model_validate({
            "birthday_users": [User.from_dict(_item) for _item in obj["birthday_users"]] if obj.get("birthday_users") is not None else None,
            "birthday_users_count": obj.get("birthday_users_count"),
            "body": obj.get("body"),
            "content_preview": obj.get("content_preview"),
            "created_at_millis": obj.get("created_at_millis"),
            "empl_amount": obj.get("empl_amount"),
            "followers": [User.from_dict(_item) for _item in obj["followers"]] if obj.get("followers") is not None else None,
            "followers_count": obj.get("followers_count"),
            "from_post": ModelPost.from_dict(obj["from_post"]) if obj.get("from_post") is not None else None,
            "from_post_ids": obj.get("from_post_ids"),
            "gift_item": Gift.from_dict(obj["gift_item"]) if obj.get("gift_item") is not None else None,
            "group": ModelGroup.from_dict(obj["group"]) if obj.get("group") is not None else None,
            "group_topic": obj.get("group_topic"),
            "id": obj.get("id"),
            "is_anonymous": obj.get("is_anonymous"),
            "is_bulk_invitation": obj.get("is_bulk_invitation"),
            "title": obj.get("title"),
            "to_post": ModelPost.from_dict(obj["to_post"]) if obj.get("to_post") is not None else None,
            "type": obj.get("type"),
            "url": obj.get("url"),
            "user": User.from_dict(obj["user"]) if obj.get("user") is not None else None,
            "vip_reward": obj.get("vip_reward")
        })
        return _obj


