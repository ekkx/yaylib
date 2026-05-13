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

from pydantic import BaseModel, ConfigDict
from typing import Any, ClassVar, Dict, List, Optional
from yaylib.models.popular_word import PopularWord
from typing import Optional, Set
from typing_extensions import Self

class PopularWordsResponse(BaseModel):
    """
    PopularWordsResponse
    """ # noqa: E501
    popular_words: Optional[List[PopularWord]] = None
    __properties: ClassVar[List[str]] = ["popular_words"]

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
        """Create an instance of PopularWordsResponse from a JSON string"""
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
        # override the default output from pydantic by calling `to_dict()` of each item in popular_words (list)
        _items = []
        if self.popular_words:
            for _item_popular_words in self.popular_words:
                if _item_popular_words:
                    _items.append(_item_popular_words.to_dict())
            _dict['popular_words'] = _items
        # set to None if popular_words (nullable) is None
        # and model_fields_set contains the field
        if self.popular_words is None and "popular_words" in self.model_fields_set:
            _dict['popular_words'] = None

        return _dict

    @classmethod
    def from_dict(cls, obj: Optional[Dict[str, Any]]) -> Optional[Self]:
        """Create an instance of PopularWordsResponse from a dict"""
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return cls.model_validate(obj)

        _obj = cls.model_validate({
            "popular_words": [PopularWord.from_dict(_item) for _item in obj["popular_words"]] if obj.get("popular_words") is not None else None
        })
        return _obj


