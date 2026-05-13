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

from pydantic import BaseModel, ConfigDict, StrictInt
from typing import Any, ClassVar, Dict, List, Optional
from yaylib.models.realm_game import RealmGame
from typing import Optional, Set
from typing_extensions import Self

class GamesResponse(BaseModel):
    """
    GamesResponse
    """ # noqa: E501
    from_id: Optional[StrictInt] = None
    games: Optional[List[RealmGame]] = None
    __properties: ClassVar[List[str]] = ["from_id", "games"]

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
        """Create an instance of GamesResponse from a JSON string"""
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
        # override the default output from pydantic by calling `to_dict()` of each item in games (list)
        _items = []
        if self.games:
            for _item_games in self.games:
                if _item_games:
                    _items.append(_item_games.to_dict())
            _dict['games'] = _items
        # set to None if from_id (nullable) is None
        # and model_fields_set contains the field
        if self.from_id is None and "from_id" in self.model_fields_set:
            _dict['from_id'] = None

        # set to None if games (nullable) is None
        # and model_fields_set contains the field
        if self.games is None and "games" in self.model_fields_set:
            _dict['games'] = None

        return _dict

    @classmethod
    def from_dict(cls, obj: Optional[Dict[str, Any]]) -> Optional[Self]:
        """Create an instance of GamesResponse from a dict"""
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return cls.model_validate(obj)

        _obj = cls.model_validate({
            "from_id": obj.get("from_id"),
            "games": [RealmGame.from_dict(_item) for _item in obj["games"]] if obj.get("games") is not None else None
        })
        return _obj


