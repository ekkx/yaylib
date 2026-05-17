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
from typing import Optional, Set
from typing_extensions import Self

class TimelinePage(BaseModel):
    """
    TimelinePage
    """ # noqa: E501
    id: Optional[StrictInt] = None
    items: Optional[List[Dict[str, Any]]] = None
    next_page_value: Optional[StrictStr] = None
    pinned_items: Optional[List[Dict[str, Any]]] = None
    show_more_hot_posts_button: Optional[StrictBool] = None
    total_item_count: Optional[StrictInt] = None
    total_item_limit: Optional[StrictInt] = None
    __properties: ClassVar[List[str]] = ["id", "items", "next_page_value", "pinned_items", "show_more_hot_posts_button", "total_item_count", "total_item_limit"]

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
        """Create an instance of TimelinePage from a JSON string"""
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
        # set to None if id (nullable) is None
        # and model_fields_set contains the field
        if self.id is None and "id" in self.model_fields_set:
            _dict['id'] = None

        # set to None if items (nullable) is None
        # and model_fields_set contains the field
        if self.items is None and "items" in self.model_fields_set:
            _dict['items'] = None

        # set to None if next_page_value (nullable) is None
        # and model_fields_set contains the field
        if self.next_page_value is None and "next_page_value" in self.model_fields_set:
            _dict['next_page_value'] = None

        # set to None if pinned_items (nullable) is None
        # and model_fields_set contains the field
        if self.pinned_items is None and "pinned_items" in self.model_fields_set:
            _dict['pinned_items'] = None

        # set to None if show_more_hot_posts_button (nullable) is None
        # and model_fields_set contains the field
        if self.show_more_hot_posts_button is None and "show_more_hot_posts_button" in self.model_fields_set:
            _dict['show_more_hot_posts_button'] = None

        # set to None if total_item_count (nullable) is None
        # and model_fields_set contains the field
        if self.total_item_count is None and "total_item_count" in self.model_fields_set:
            _dict['total_item_count'] = None

        # set to None if total_item_limit (nullable) is None
        # and model_fields_set contains the field
        if self.total_item_limit is None and "total_item_limit" in self.model_fields_set:
            _dict['total_item_limit'] = None

        return _dict

    @classmethod
    def from_dict(cls, obj: Optional[Dict[str, Any]]) -> Optional[Self]:
        """Create an instance of TimelinePage from a dict"""
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return cls.model_validate(obj)

        _obj = cls.model_validate({
            "id": obj.get("id"),
            "items": obj.get("items"),
            "next_page_value": obj.get("next_page_value"),
            "pinned_items": obj.get("pinned_items"),
            "show_more_hot_posts_button": obj.get("show_more_hot_posts_button"),
            "total_item_count": obj.get("total_item_count"),
            "total_item_limit": obj.get("total_item_limit")
        })
        return _obj


