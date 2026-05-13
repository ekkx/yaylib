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
from yaylib.models.user import User
from typing import Optional, Set
from typing_extensions import Self

class ModelReview(BaseModel):
    """
    ModelReview
    """ # noqa: E501
    comment: Optional[StrictStr] = None
    created_at_millis: Optional[StrictInt] = None
    id: Optional[StrictInt] = None
    is_mutual: Optional[StrictBool] = None
    is_pinned: Optional[StrictBool] = None
    reported_count: Optional[StrictInt] = None
    reviewer: Optional[User] = None
    user: Optional[User] = None
    __properties: ClassVar[List[str]] = ["comment", "created_at_millis", "id", "is_mutual", "is_pinned", "reported_count", "reviewer", "user"]

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
        """Create an instance of ModelReview from a JSON string"""
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
        # override the default output from pydantic by calling `to_dict()` of reviewer
        if self.reviewer:
            _dict['reviewer'] = self.reviewer.to_dict()
        # override the default output from pydantic by calling `to_dict()` of user
        if self.user:
            _dict['user'] = self.user.to_dict()
        # set to None if comment (nullable) is None
        # and model_fields_set contains the field
        if self.comment is None and "comment" in self.model_fields_set:
            _dict['comment'] = None

        # set to None if created_at_millis (nullable) is None
        # and model_fields_set contains the field
        if self.created_at_millis is None and "created_at_millis" in self.model_fields_set:
            _dict['created_at_millis'] = None

        # set to None if id (nullable) is None
        # and model_fields_set contains the field
        if self.id is None and "id" in self.model_fields_set:
            _dict['id'] = None

        # set to None if is_mutual (nullable) is None
        # and model_fields_set contains the field
        if self.is_mutual is None and "is_mutual" in self.model_fields_set:
            _dict['is_mutual'] = None

        # set to None if is_pinned (nullable) is None
        # and model_fields_set contains the field
        if self.is_pinned is None and "is_pinned" in self.model_fields_set:
            _dict['is_pinned'] = None

        # set to None if reported_count (nullable) is None
        # and model_fields_set contains the field
        if self.reported_count is None and "reported_count" in self.model_fields_set:
            _dict['reported_count'] = None

        # set to None if reviewer (nullable) is None
        # and model_fields_set contains the field
        if self.reviewer is None and "reviewer" in self.model_fields_set:
            _dict['reviewer'] = None

        # set to None if user (nullable) is None
        # and model_fields_set contains the field
        if self.user is None and "user" in self.model_fields_set:
            _dict['user'] = None

        return _dict

    @classmethod
    def from_dict(cls, obj: Optional[Dict[str, Any]]) -> Optional[Self]:
        """Create an instance of ModelReview from a dict"""
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return cls.model_validate(obj)

        _obj = cls.model_validate({
            "comment": obj.get("comment"),
            "created_at_millis": obj.get("created_at_millis"),
            "id": obj.get("id"),
            "is_mutual": obj.get("is_mutual"),
            "is_pinned": obj.get("is_pinned"),
            "reported_count": obj.get("reported_count"),
            "reviewer": User.from_dict(obj["reviewer"]) if obj.get("reviewer") is not None else None,
            "user": User.from_dict(obj["user"]) if obj.get("user") is not None else None
        })
        return _obj


