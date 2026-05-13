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

from pydantic import BaseModel, ConfigDict, StrictInt, StrictStr
from typing import Any, ClassVar, Dict, List, Optional
from yaylib.models.group import Group
from yaylib.models.metadata import Metadata
from yaylib.models.post import Post
from yaylib.models.realm_gift import RealmGift
from yaylib.models.realm_user import RealmUser
from typing import Optional, Set
from typing_extensions import Self

class Activity(BaseModel):
    """
    Activity
    """ # noqa: E501
    birthday_users: Optional[List[RealmUser]] = None
    birthday_users_count: Optional[StrictInt] = None
    created_at: Optional[StrictInt] = None
    empl_amount: Optional[StrictInt] = None
    followers: Optional[List[RealmUser]] = None
    followers_count: Optional[StrictInt] = None
    from_post: Optional[Post] = None
    from_post_ids: Optional[List[StrictInt]] = None
    gift_item: Optional[RealmGift] = None
    group: Optional[Group] = None
    id: Optional[StrictInt] = None
    metadata: Optional[Metadata] = None
    to_post: Optional[Post] = None
    type: Optional[StrictStr] = None
    user: Optional[RealmUser] = None
    vip_reward: Optional[StrictInt] = None
    __properties: ClassVar[List[str]] = ["birthday_users", "birthday_users_count", "created_at", "empl_amount", "followers", "followers_count", "from_post", "from_post_ids", "gift_item", "group", "id", "metadata", "to_post", "type", "user", "vip_reward"]

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
        """Create an instance of Activity from a JSON string"""
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
        # override the default output from pydantic by calling `to_dict()` of metadata
        if self.metadata:
            _dict['metadata'] = self.metadata.to_dict()
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

        # set to None if created_at (nullable) is None
        # and model_fields_set contains the field
        if self.created_at is None and "created_at" in self.model_fields_set:
            _dict['created_at'] = None

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

        # set to None if id (nullable) is None
        # and model_fields_set contains the field
        if self.id is None and "id" in self.model_fields_set:
            _dict['id'] = None

        # set to None if metadata (nullable) is None
        # and model_fields_set contains the field
        if self.metadata is None and "metadata" in self.model_fields_set:
            _dict['metadata'] = None

        # set to None if to_post (nullable) is None
        # and model_fields_set contains the field
        if self.to_post is None and "to_post" in self.model_fields_set:
            _dict['to_post'] = None

        # set to None if type (nullable) is None
        # and model_fields_set contains the field
        if self.type is None and "type" in self.model_fields_set:
            _dict['type'] = None

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
        """Create an instance of Activity from a dict"""
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return cls.model_validate(obj)

        _obj = cls.model_validate({
            "birthday_users": [RealmUser.from_dict(_item) for _item in obj["birthday_users"]] if obj.get("birthday_users") is not None else None,
            "birthday_users_count": obj.get("birthday_users_count"),
            "created_at": obj.get("created_at"),
            "empl_amount": obj.get("empl_amount"),
            "followers": [RealmUser.from_dict(_item) for _item in obj["followers"]] if obj.get("followers") is not None else None,
            "followers_count": obj.get("followers_count"),
            "from_post": Post.from_dict(obj["from_post"]) if obj.get("from_post") is not None else None,
            "from_post_ids": obj.get("from_post_ids"),
            "gift_item": RealmGift.from_dict(obj["gift_item"]) if obj.get("gift_item") is not None else None,
            "group": Group.from_dict(obj["group"]) if obj.get("group") is not None else None,
            "id": obj.get("id"),
            "metadata": Metadata.from_dict(obj["metadata"]) if obj.get("metadata") is not None else None,
            "to_post": Post.from_dict(obj["to_post"]) if obj.get("to_post") is not None else None,
            "type": obj.get("type"),
            "user": RealmUser.from_dict(obj["user"]) if obj.get("user") is not None else None,
            "vip_reward": obj.get("vip_reward")
        })
        return _obj


