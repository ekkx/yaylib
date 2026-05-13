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
from yaylib.models.localized_string_dto import LocalizedStringDTO
from typing import Optional, Set
from typing_extensions import Self

class DAppDTO(BaseModel):
    """
    DAppDTO
    """ # noqa: E501
    description_localized: Optional[LocalizedStringDTO] = None
    enabled: Optional[StrictBool] = None
    icon_url: Optional[StrictStr] = None
    id: Optional[StrictStr] = None
    order: Optional[StrictInt] = None
    search_keywords: Optional[List[StrictStr]] = None
    show_in_recommend_d_apps: Optional[StrictBool] = None
    title_localized: Optional[LocalizedStringDTO] = None
    url_localized: Optional[LocalizedStringDTO] = None
    __properties: ClassVar[List[str]] = ["description_localized", "enabled", "icon_url", "id", "order", "search_keywords", "show_in_recommend_d_apps", "title_localized", "url_localized"]

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
        """Create an instance of DAppDTO from a JSON string"""
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
        # override the default output from pydantic by calling `to_dict()` of description_localized
        if self.description_localized:
            _dict['description_localized'] = self.description_localized.to_dict()
        # override the default output from pydantic by calling `to_dict()` of title_localized
        if self.title_localized:
            _dict['title_localized'] = self.title_localized.to_dict()
        # override the default output from pydantic by calling `to_dict()` of url_localized
        if self.url_localized:
            _dict['url_localized'] = self.url_localized.to_dict()
        # set to None if description_localized (nullable) is None
        # and model_fields_set contains the field
        if self.description_localized is None and "description_localized" in self.model_fields_set:
            _dict['description_localized'] = None

        # set to None if enabled (nullable) is None
        # and model_fields_set contains the field
        if self.enabled is None and "enabled" in self.model_fields_set:
            _dict['enabled'] = None

        # set to None if icon_url (nullable) is None
        # and model_fields_set contains the field
        if self.icon_url is None and "icon_url" in self.model_fields_set:
            _dict['icon_url'] = None

        # set to None if id (nullable) is None
        # and model_fields_set contains the field
        if self.id is None and "id" in self.model_fields_set:
            _dict['id'] = None

        # set to None if order (nullable) is None
        # and model_fields_set contains the field
        if self.order is None and "order" in self.model_fields_set:
            _dict['order'] = None

        # set to None if search_keywords (nullable) is None
        # and model_fields_set contains the field
        if self.search_keywords is None and "search_keywords" in self.model_fields_set:
            _dict['search_keywords'] = None

        # set to None if show_in_recommend_d_apps (nullable) is None
        # and model_fields_set contains the field
        if self.show_in_recommend_d_apps is None and "show_in_recommend_d_apps" in self.model_fields_set:
            _dict['show_in_recommend_d_apps'] = None

        # set to None if title_localized (nullable) is None
        # and model_fields_set contains the field
        if self.title_localized is None and "title_localized" in self.model_fields_set:
            _dict['title_localized'] = None

        # set to None if url_localized (nullable) is None
        # and model_fields_set contains the field
        if self.url_localized is None and "url_localized" in self.model_fields_set:
            _dict['url_localized'] = None

        return _dict

    @classmethod
    def from_dict(cls, obj: Optional[Dict[str, Any]]) -> Optional[Self]:
        """Create an instance of DAppDTO from a dict"""
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return cls.model_validate(obj)

        _obj = cls.model_validate({
            "description_localized": LocalizedStringDTO.from_dict(obj["description_localized"]) if obj.get("description_localized") is not None else None,
            "enabled": obj.get("enabled"),
            "icon_url": obj.get("icon_url"),
            "id": obj.get("id"),
            "order": obj.get("order"),
            "search_keywords": obj.get("search_keywords"),
            "show_in_recommend_d_apps": obj.get("show_in_recommend_d_apps"),
            "title_localized": LocalizedStringDTO.from_dict(obj["title_localized"]) if obj.get("title_localized") is not None else None,
            "url_localized": LocalizedStringDTO.from_dict(obj["url_localized"]) if obj.get("url_localized") is not None else None
        })
        return _obj


