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
from yaylib.models.pal import Pal
from yaylib.models.pal_race_league_dto import PalRaceLeagueDTO
from yaylib.models.pal_rank import PalRank
from yaylib.models.race_reward import RaceReward
from yaylib.models.reward import Reward
from typing import Optional, Set
from typing_extensions import Self

class RaceHistory(BaseModel):
    """
    RaceHistory
    """ # noqa: E501
    animation_url: Optional[StrictStr] = None
    league: Optional[PalRaceLeagueDTO] = None
    league_name: Optional[StrictStr] = None
    pals: Optional[List[Pal]] = None
    race_date: Optional[StrictInt] = None
    race_entry_fee: Optional[StrictInt] = None
    race_id: Optional[StrictStr] = None
    race_reward: Optional[RaceReward] = None
    race_start_at: Optional[StrictInt] = None
    race_status: Optional[StrictStr] = None
    results: Optional[List[PalRank]] = None
    reward: Optional[Reward] = None
    __properties: ClassVar[List[str]] = ["animation_url", "league", "league_name", "pals", "race_date", "race_entry_fee", "race_id", "race_reward", "race_start_at", "race_status", "results", "reward"]

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
        """Create an instance of RaceHistory from a JSON string"""
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
        # override the default output from pydantic by calling `to_dict()` of league
        if self.league:
            _dict['league'] = self.league.to_dict()
        # override the default output from pydantic by calling `to_dict()` of each item in pals (list)
        _items = []
        if self.pals:
            for _item_pals in self.pals:
                if _item_pals:
                    _items.append(_item_pals.to_dict())
            _dict['pals'] = _items
        # override the default output from pydantic by calling `to_dict()` of race_reward
        if self.race_reward:
            _dict['race_reward'] = self.race_reward.to_dict()
        # override the default output from pydantic by calling `to_dict()` of each item in results (list)
        _items = []
        if self.results:
            for _item_results in self.results:
                if _item_results:
                    _items.append(_item_results.to_dict())
            _dict['results'] = _items
        # override the default output from pydantic by calling `to_dict()` of reward
        if self.reward:
            _dict['reward'] = self.reward.to_dict()
        # set to None if animation_url (nullable) is None
        # and model_fields_set contains the field
        if self.animation_url is None and "animation_url" in self.model_fields_set:
            _dict['animation_url'] = None

        # set to None if league (nullable) is None
        # and model_fields_set contains the field
        if self.league is None and "league" in self.model_fields_set:
            _dict['league'] = None

        # set to None if league_name (nullable) is None
        # and model_fields_set contains the field
        if self.league_name is None and "league_name" in self.model_fields_set:
            _dict['league_name'] = None

        # set to None if pals (nullable) is None
        # and model_fields_set contains the field
        if self.pals is None and "pals" in self.model_fields_set:
            _dict['pals'] = None

        # set to None if race_date (nullable) is None
        # and model_fields_set contains the field
        if self.race_date is None and "race_date" in self.model_fields_set:
            _dict['race_date'] = None

        # set to None if race_entry_fee (nullable) is None
        # and model_fields_set contains the field
        if self.race_entry_fee is None and "race_entry_fee" in self.model_fields_set:
            _dict['race_entry_fee'] = None

        # set to None if race_id (nullable) is None
        # and model_fields_set contains the field
        if self.race_id is None and "race_id" in self.model_fields_set:
            _dict['race_id'] = None

        # set to None if race_reward (nullable) is None
        # and model_fields_set contains the field
        if self.race_reward is None and "race_reward" in self.model_fields_set:
            _dict['race_reward'] = None

        # set to None if race_start_at (nullable) is None
        # and model_fields_set contains the field
        if self.race_start_at is None and "race_start_at" in self.model_fields_set:
            _dict['race_start_at'] = None

        # set to None if race_status (nullable) is None
        # and model_fields_set contains the field
        if self.race_status is None and "race_status" in self.model_fields_set:
            _dict['race_status'] = None

        # set to None if results (nullable) is None
        # and model_fields_set contains the field
        if self.results is None and "results" in self.model_fields_set:
            _dict['results'] = None

        # set to None if reward (nullable) is None
        # and model_fields_set contains the field
        if self.reward is None and "reward" in self.model_fields_set:
            _dict['reward'] = None

        return _dict

    @classmethod
    def from_dict(cls, obj: Optional[Dict[str, Any]]) -> Optional[Self]:
        """Create an instance of RaceHistory from a dict"""
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return cls.model_validate(obj)

        _obj = cls.model_validate({
            "animation_url": obj.get("animation_url"),
            "league": PalRaceLeagueDTO.from_dict(obj["league"]) if obj.get("league") is not None else None,
            "league_name": obj.get("league_name"),
            "pals": [Pal.from_dict(_item) for _item in obj["pals"]] if obj.get("pals") is not None else None,
            "race_date": obj.get("race_date"),
            "race_entry_fee": obj.get("race_entry_fee"),
            "race_id": obj.get("race_id"),
            "race_reward": RaceReward.from_dict(obj["race_reward"]) if obj.get("race_reward") is not None else None,
            "race_start_at": obj.get("race_start_at"),
            "race_status": obj.get("race_status"),
            "results": [PalRank.from_dict(_item) for _item in obj["results"]] if obj.get("results") is not None else None,
            "reward": Reward.from_dict(obj["reward"]) if obj.get("reward") is not None else None
        })
        return _obj


