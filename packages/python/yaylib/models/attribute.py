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

from pydantic import BaseModel, ConfigDict, StrictFloat, StrictInt, StrictStr
from typing import Any, ClassVar, Dict, List, Optional, Union
from typing import Optional, Set
from typing_extensions import Self

class Attribute(BaseModel):
    """
    Attribute
    """ # noqa: E501
    birthday: Optional[StrictStr] = None
    body: Optional[StrictStr] = None
    climbing: Optional[StrictStr] = None
    eligible_league: Optional[StrictInt] = None
    fatigue: Optional[Union[StrictFloat, StrictInt]] = None
    generation: Optional[StrictInt] = None
    grade: Optional[StrictStr] = None
    head: Optional[StrictStr] = None
    item: Optional[StrictStr] = None
    level: Optional[StrictInt] = None
    luck: Optional[StrictInt] = None
    momentum: Optional[StrictInt] = None
    name: Optional[StrictStr] = None
    power: Optional[StrictInt] = None
    racingstyle: Optional[StrictStr] = None
    rank: Optional[StrictStr] = None
    running: Optional[StrictStr] = None
    skin: Optional[StrictStr] = None
    speed: Optional[StrictInt] = None
    stamina: Optional[StrictInt] = None
    status: Optional[StrictStr] = None
    swimming: Optional[StrictStr] = None
    technique: Optional[StrictInt] = None
    tribe: Optional[StrictStr] = None
    type: Optional[StrictStr] = None
    __properties: ClassVar[List[str]] = ["birthday", "body", "climbing", "eligible_league", "fatigue", "generation", "grade", "head", "item", "level", "luck", "momentum", "name", "power", "racingstyle", "rank", "running", "skin", "speed", "stamina", "status", "swimming", "technique", "tribe", "type"]

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
        """Create an instance of Attribute from a JSON string"""
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
        # set to None if birthday (nullable) is None
        # and model_fields_set contains the field
        if self.birthday is None and "birthday" in self.model_fields_set:
            _dict['birthday'] = None

        # set to None if body (nullable) is None
        # and model_fields_set contains the field
        if self.body is None and "body" in self.model_fields_set:
            _dict['body'] = None

        # set to None if climbing (nullable) is None
        # and model_fields_set contains the field
        if self.climbing is None and "climbing" in self.model_fields_set:
            _dict['climbing'] = None

        # set to None if eligible_league (nullable) is None
        # and model_fields_set contains the field
        if self.eligible_league is None and "eligible_league" in self.model_fields_set:
            _dict['eligible_league'] = None

        # set to None if fatigue (nullable) is None
        # and model_fields_set contains the field
        if self.fatigue is None and "fatigue" in self.model_fields_set:
            _dict['fatigue'] = None

        # set to None if generation (nullable) is None
        # and model_fields_set contains the field
        if self.generation is None and "generation" in self.model_fields_set:
            _dict['generation'] = None

        # set to None if grade (nullable) is None
        # and model_fields_set contains the field
        if self.grade is None and "grade" in self.model_fields_set:
            _dict['grade'] = None

        # set to None if head (nullable) is None
        # and model_fields_set contains the field
        if self.head is None and "head" in self.model_fields_set:
            _dict['head'] = None

        # set to None if item (nullable) is None
        # and model_fields_set contains the field
        if self.item is None and "item" in self.model_fields_set:
            _dict['item'] = None

        # set to None if level (nullable) is None
        # and model_fields_set contains the field
        if self.level is None and "level" in self.model_fields_set:
            _dict['level'] = None

        # set to None if luck (nullable) is None
        # and model_fields_set contains the field
        if self.luck is None and "luck" in self.model_fields_set:
            _dict['luck'] = None

        # set to None if momentum (nullable) is None
        # and model_fields_set contains the field
        if self.momentum is None and "momentum" in self.model_fields_set:
            _dict['momentum'] = None

        # set to None if name (nullable) is None
        # and model_fields_set contains the field
        if self.name is None and "name" in self.model_fields_set:
            _dict['name'] = None

        # set to None if power (nullable) is None
        # and model_fields_set contains the field
        if self.power is None and "power" in self.model_fields_set:
            _dict['power'] = None

        # set to None if racingstyle (nullable) is None
        # and model_fields_set contains the field
        if self.racingstyle is None and "racingstyle" in self.model_fields_set:
            _dict['racingstyle'] = None

        # set to None if rank (nullable) is None
        # and model_fields_set contains the field
        if self.rank is None and "rank" in self.model_fields_set:
            _dict['rank'] = None

        # set to None if running (nullable) is None
        # and model_fields_set contains the field
        if self.running is None and "running" in self.model_fields_set:
            _dict['running'] = None

        # set to None if skin (nullable) is None
        # and model_fields_set contains the field
        if self.skin is None and "skin" in self.model_fields_set:
            _dict['skin'] = None

        # set to None if speed (nullable) is None
        # and model_fields_set contains the field
        if self.speed is None and "speed" in self.model_fields_set:
            _dict['speed'] = None

        # set to None if stamina (nullable) is None
        # and model_fields_set contains the field
        if self.stamina is None and "stamina" in self.model_fields_set:
            _dict['stamina'] = None

        # set to None if status (nullable) is None
        # and model_fields_set contains the field
        if self.status is None and "status" in self.model_fields_set:
            _dict['status'] = None

        # set to None if swimming (nullable) is None
        # and model_fields_set contains the field
        if self.swimming is None and "swimming" in self.model_fields_set:
            _dict['swimming'] = None

        # set to None if technique (nullable) is None
        # and model_fields_set contains the field
        if self.technique is None and "technique" in self.model_fields_set:
            _dict['technique'] = None

        # set to None if tribe (nullable) is None
        # and model_fields_set contains the field
        if self.tribe is None and "tribe" in self.model_fields_set:
            _dict['tribe'] = None

        # set to None if type (nullable) is None
        # and model_fields_set contains the field
        if self.type is None and "type" in self.model_fields_set:
            _dict['type'] = None

        return _dict

    @classmethod
    def from_dict(cls, obj: Optional[Dict[str, Any]]) -> Optional[Self]:
        """Create an instance of Attribute from a dict"""
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return cls.model_validate(obj)

        _obj = cls.model_validate({
            "birthday": obj.get("birthday"),
            "body": obj.get("body"),
            "climbing": obj.get("climbing"),
            "eligible_league": obj.get("eligible_league"),
            "fatigue": obj.get("fatigue"),
            "generation": obj.get("generation"),
            "grade": obj.get("grade"),
            "head": obj.get("head"),
            "item": obj.get("item"),
            "level": obj.get("level"),
            "luck": obj.get("luck"),
            "momentum": obj.get("momentum"),
            "name": obj.get("name"),
            "power": obj.get("power"),
            "racingstyle": obj.get("racingstyle"),
            "rank": obj.get("rank"),
            "running": obj.get("running"),
            "skin": obj.get("skin"),
            "speed": obj.get("speed"),
            "stamina": obj.get("stamina"),
            "status": obj.get("status"),
            "swimming": obj.get("swimming"),
            "technique": obj.get("technique"),
            "tribe": obj.get("tribe"),
            "type": obj.get("type")
        })
        return _obj


