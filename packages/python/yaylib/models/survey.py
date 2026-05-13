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

from pydantic import BaseModel, ConfigDict, StrictBool, StrictInt
from typing import Any, ClassVar, Dict, List, Optional
from yaylib.models.choice import Choice
from typing import Optional, Set
from typing_extensions import Self

class Survey(BaseModel):
    """
    Survey
    """ # noqa: E501
    choices: Optional[List[Choice]] = None
    id: Optional[StrictInt] = None
    voted: Optional[StrictBool] = None
    votes_count: Optional[StrictInt] = None
    __properties: ClassVar[List[str]] = ["choices", "id", "voted", "votes_count"]

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
        """Create an instance of Survey from a JSON string"""
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
        # override the default output from pydantic by calling `to_dict()` of each item in choices (list)
        _items = []
        if self.choices:
            for _item_choices in self.choices:
                if _item_choices:
                    _items.append(_item_choices.to_dict())
            _dict['choices'] = _items
        # set to None if choices (nullable) is None
        # and model_fields_set contains the field
        if self.choices is None and "choices" in self.model_fields_set:
            _dict['choices'] = None

        # set to None if id (nullable) is None
        # and model_fields_set contains the field
        if self.id is None and "id" in self.model_fields_set:
            _dict['id'] = None

        # set to None if voted (nullable) is None
        # and model_fields_set contains the field
        if self.voted is None and "voted" in self.model_fields_set:
            _dict['voted'] = None

        # set to None if votes_count (nullable) is None
        # and model_fields_set contains the field
        if self.votes_count is None and "votes_count" in self.model_fields_set:
            _dict['votes_count'] = None

        return _dict

    @classmethod
    def from_dict(cls, obj: Optional[Dict[str, Any]]) -> Optional[Self]:
        """Create an instance of Survey from a dict"""
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return cls.model_validate(obj)

        _obj = cls.model_validate({
            "choices": [Choice.from_dict(_item) for _item in obj["choices"]] if obj.get("choices") is not None else None,
            "id": obj.get("id"),
            "voted": obj.get("voted"),
            "votes_count": obj.get("votes_count")
        })
        return _obj


