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
from yaylib.models.conference_call_bump_params import ConferenceCallBumpParams
from yaylib.models.conference_call_user_role import ConferenceCallUserRole
from yaylib.models.realm_game import RealmGame
from yaylib.models.realm_genre import RealmGenre
from yaylib.models.realm_user import RealmUser
from typing import Optional, Set
from typing_extensions import Self

class RealmConferenceCall(BaseModel):
    """
    RealmConferenceCall
    """ # noqa: E501
    active: Optional[StrictBool] = None
    agora_channel: Optional[StrictStr] = None
    agora_token: Optional[StrictStr] = None
    anonymous_call_users_count: Optional[StrictInt] = None
    bump_params: Optional[ConferenceCallBumpParams] = None
    call_type: Optional[StrictStr] = None
    conference_call_user_roles: Optional[List[ConferenceCallUserRole]] = None
    conference_call_users: Optional[List[RealmUser]] = None
    conference_call_users_count: Optional[StrictInt] = None
    duration_seconds: Optional[StrictInt] = None
    game: Optional[RealmGame] = None
    genre: Optional[RealmGenre] = None
    group_id: Optional[StrictInt] = None
    id: Optional[StrictInt] = None
    joinable_by: Optional[StrictStr] = None
    max_participants: Optional[StrictInt] = None
    post_id: Optional[StrictInt] = None
    server: Optional[StrictStr] = None
    __properties: ClassVar[List[str]] = ["active", "agora_channel", "agora_token", "anonymous_call_users_count", "bump_params", "call_type", "conference_call_user_roles", "conference_call_users", "conference_call_users_count", "duration_seconds", "game", "genre", "group_id", "id", "joinable_by", "max_participants", "post_id", "server"]

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
        """Create an instance of RealmConferenceCall from a JSON string"""
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
        # override the default output from pydantic by calling `to_dict()` of bump_params
        if self.bump_params:
            _dict['bump_params'] = self.bump_params.to_dict()
        # override the default output from pydantic by calling `to_dict()` of each item in conference_call_user_roles (list)
        _items = []
        if self.conference_call_user_roles:
            for _item_conference_call_user_roles in self.conference_call_user_roles:
                if _item_conference_call_user_roles:
                    _items.append(_item_conference_call_user_roles.to_dict())
            _dict['conference_call_user_roles'] = _items
        # override the default output from pydantic by calling `to_dict()` of each item in conference_call_users (list)
        _items = []
        if self.conference_call_users:
            for _item_conference_call_users in self.conference_call_users:
                if _item_conference_call_users:
                    _items.append(_item_conference_call_users.to_dict())
            _dict['conference_call_users'] = _items
        # override the default output from pydantic by calling `to_dict()` of game
        if self.game:
            _dict['game'] = self.game.to_dict()
        # override the default output from pydantic by calling `to_dict()` of genre
        if self.genre:
            _dict['genre'] = self.genre.to_dict()
        # set to None if active (nullable) is None
        # and model_fields_set contains the field
        if self.active is None and "active" in self.model_fields_set:
            _dict['active'] = None

        # set to None if agora_channel (nullable) is None
        # and model_fields_set contains the field
        if self.agora_channel is None and "agora_channel" in self.model_fields_set:
            _dict['agora_channel'] = None

        # set to None if agora_token (nullable) is None
        # and model_fields_set contains the field
        if self.agora_token is None and "agora_token" in self.model_fields_set:
            _dict['agora_token'] = None

        # set to None if anonymous_call_users_count (nullable) is None
        # and model_fields_set contains the field
        if self.anonymous_call_users_count is None and "anonymous_call_users_count" in self.model_fields_set:
            _dict['anonymous_call_users_count'] = None

        # set to None if bump_params (nullable) is None
        # and model_fields_set contains the field
        if self.bump_params is None and "bump_params" in self.model_fields_set:
            _dict['bump_params'] = None

        # set to None if call_type (nullable) is None
        # and model_fields_set contains the field
        if self.call_type is None and "call_type" in self.model_fields_set:
            _dict['call_type'] = None

        # set to None if conference_call_user_roles (nullable) is None
        # and model_fields_set contains the field
        if self.conference_call_user_roles is None and "conference_call_user_roles" in self.model_fields_set:
            _dict['conference_call_user_roles'] = None

        # set to None if conference_call_users (nullable) is None
        # and model_fields_set contains the field
        if self.conference_call_users is None and "conference_call_users" in self.model_fields_set:
            _dict['conference_call_users'] = None

        # set to None if conference_call_users_count (nullable) is None
        # and model_fields_set contains the field
        if self.conference_call_users_count is None and "conference_call_users_count" in self.model_fields_set:
            _dict['conference_call_users_count'] = None

        # set to None if duration_seconds (nullable) is None
        # and model_fields_set contains the field
        if self.duration_seconds is None and "duration_seconds" in self.model_fields_set:
            _dict['duration_seconds'] = None

        # set to None if game (nullable) is None
        # and model_fields_set contains the field
        if self.game is None and "game" in self.model_fields_set:
            _dict['game'] = None

        # set to None if genre (nullable) is None
        # and model_fields_set contains the field
        if self.genre is None and "genre" in self.model_fields_set:
            _dict['genre'] = None

        # set to None if group_id (nullable) is None
        # and model_fields_set contains the field
        if self.group_id is None and "group_id" in self.model_fields_set:
            _dict['group_id'] = None

        # set to None if id (nullable) is None
        # and model_fields_set contains the field
        if self.id is None and "id" in self.model_fields_set:
            _dict['id'] = None

        # set to None if joinable_by (nullable) is None
        # and model_fields_set contains the field
        if self.joinable_by is None and "joinable_by" in self.model_fields_set:
            _dict['joinable_by'] = None

        # set to None if max_participants (nullable) is None
        # and model_fields_set contains the field
        if self.max_participants is None and "max_participants" in self.model_fields_set:
            _dict['max_participants'] = None

        # set to None if post_id (nullable) is None
        # and model_fields_set contains the field
        if self.post_id is None and "post_id" in self.model_fields_set:
            _dict['post_id'] = None

        # set to None if server (nullable) is None
        # and model_fields_set contains the field
        if self.server is None and "server" in self.model_fields_set:
            _dict['server'] = None

        return _dict

    @classmethod
    def from_dict(cls, obj: Optional[Dict[str, Any]]) -> Optional[Self]:
        """Create an instance of RealmConferenceCall from a dict"""
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return cls.model_validate(obj)

        _obj = cls.model_validate({
            "active": obj.get("active"),
            "agora_channel": obj.get("agora_channel"),
            "agora_token": obj.get("agora_token"),
            "anonymous_call_users_count": obj.get("anonymous_call_users_count"),
            "bump_params": ConferenceCallBumpParams.from_dict(obj["bump_params"]) if obj.get("bump_params") is not None else None,
            "call_type": obj.get("call_type"),
            "conference_call_user_roles": [ConferenceCallUserRole.from_dict(_item) for _item in obj["conference_call_user_roles"]] if obj.get("conference_call_user_roles") is not None else None,
            "conference_call_users": [RealmUser.from_dict(_item) for _item in obj["conference_call_users"]] if obj.get("conference_call_users") is not None else None,
            "conference_call_users_count": obj.get("conference_call_users_count"),
            "duration_seconds": obj.get("duration_seconds"),
            "game": RealmGame.from_dict(obj["game"]) if obj.get("game") is not None else None,
            "genre": RealmGenre.from_dict(obj["genre"]) if obj.get("genre") is not None else None,
            "group_id": obj.get("group_id"),
            "id": obj.get("id"),
            "joinable_by": obj.get("joinable_by"),
            "max_participants": obj.get("max_participants"),
            "post_id": obj.get("post_id"),
            "server": obj.get("server")
        })
        return _obj


