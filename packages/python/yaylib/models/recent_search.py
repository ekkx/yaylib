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
from yaylib.models.post_tag import PostTag
from yaylib.models.realm_user import RealmUser
from typing import Optional, Set
from typing_extensions import Self

class RecentSearch(BaseModel):
    """
    RecentSearch
    """ # noqa: E501
    hashtag: Optional[PostTag] = None
    id: Optional[StrictInt] = None
    keyword: Optional[StrictStr] = None
    type: Optional[StrictStr] = None
    user: Optional[RealmUser] = None
    __properties: ClassVar[List[str]] = ["hashtag", "id", "keyword", "type", "user"]

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
        """Create an instance of RecentSearch from a JSON string"""
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
        # override the default output from pydantic by calling `to_dict()` of hashtag
        if self.hashtag:
            _dict['hashtag'] = self.hashtag.to_dict()
        # override the default output from pydantic by calling `to_dict()` of user
        if self.user:
            _dict['user'] = self.user.to_dict()
        # set to None if hashtag (nullable) is None
        # and model_fields_set contains the field
        if self.hashtag is None and "hashtag" in self.model_fields_set:
            _dict['hashtag'] = None

        # set to None if id (nullable) is None
        # and model_fields_set contains the field
        if self.id is None and "id" in self.model_fields_set:
            _dict['id'] = None

        # set to None if keyword (nullable) is None
        # and model_fields_set contains the field
        if self.keyword is None and "keyword" in self.model_fields_set:
            _dict['keyword'] = None

        # set to None if type (nullable) is None
        # and model_fields_set contains the field
        if self.type is None and "type" in self.model_fields_set:
            _dict['type'] = None

        # set to None if user (nullable) is None
        # and model_fields_set contains the field
        if self.user is None and "user" in self.model_fields_set:
            _dict['user'] = None

        return _dict

    @classmethod
    def from_dict(cls, obj: Optional[Dict[str, Any]]) -> Optional[Self]:
        """Create an instance of RecentSearch from a dict"""
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return cls.model_validate(obj)

        _obj = cls.model_validate({
            "hashtag": PostTag.from_dict(obj["hashtag"]) if obj.get("hashtag") is not None else None,
            "id": obj.get("id"),
            "keyword": obj.get("keyword"),
            "type": obj.get("type"),
            "user": RealmUser.from_dict(obj["user"]) if obj.get("user") is not None else None
        })
        return _obj


