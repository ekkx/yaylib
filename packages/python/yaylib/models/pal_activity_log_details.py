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
from yaylib.models.pal_grade import PalGrade
from typing import Optional, Set
from typing_extensions import Self

class PalActivityLogDetails(BaseModel):
    """
    PalActivityLogDetails
    """ # noqa: E501
    attacking: Optional[StrictInt] = None
    cost: Optional[Union[StrictFloat, StrictInt]] = None
    gacha_type: Optional[StrictStr] = None
    item_status: Optional[StrictStr] = None
    item_subtype: Optional[StrictStr] = None
    item_type: Optional[StrictStr] = None
    level: Optional[StrictInt] = None
    nft_type: Optional[StrictStr] = None
    outcome: Optional[StrictInt] = None
    pal_grade: Optional[PalGrade] = None
    pal_image: Optional[StrictStr] = None
    pal_name: Optional[StrictStr] = None
    places: Optional[List[StrictInt]] = None
    race_id: Optional[StrictStr] = None
    rank: Optional[StrictStr] = None
    result: Optional[StrictStr] = None
    token_id: Optional[StrictStr] = None
    transaction_code: Optional[StrictStr] = None
    __properties: ClassVar[List[str]] = ["attacking", "cost", "gacha_type", "item_status", "item_subtype", "item_type", "level", "nft_type", "outcome", "pal_grade", "pal_image", "pal_name", "places", "race_id", "rank", "result", "token_id", "transaction_code"]

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
        """Create an instance of PalActivityLogDetails from a JSON string"""
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
        # set to None if attacking (nullable) is None
        # and model_fields_set contains the field
        if self.attacking is None and "attacking" in self.model_fields_set:
            _dict['attacking'] = None

        # set to None if cost (nullable) is None
        # and model_fields_set contains the field
        if self.cost is None and "cost" in self.model_fields_set:
            _dict['cost'] = None

        # set to None if gacha_type (nullable) is None
        # and model_fields_set contains the field
        if self.gacha_type is None and "gacha_type" in self.model_fields_set:
            _dict['gacha_type'] = None

        # set to None if item_status (nullable) is None
        # and model_fields_set contains the field
        if self.item_status is None and "item_status" in self.model_fields_set:
            _dict['item_status'] = None

        # set to None if item_subtype (nullable) is None
        # and model_fields_set contains the field
        if self.item_subtype is None and "item_subtype" in self.model_fields_set:
            _dict['item_subtype'] = None

        # set to None if item_type (nullable) is None
        # and model_fields_set contains the field
        if self.item_type is None and "item_type" in self.model_fields_set:
            _dict['item_type'] = None

        # set to None if level (nullable) is None
        # and model_fields_set contains the field
        if self.level is None and "level" in self.model_fields_set:
            _dict['level'] = None

        # set to None if nft_type (nullable) is None
        # and model_fields_set contains the field
        if self.nft_type is None and "nft_type" in self.model_fields_set:
            _dict['nft_type'] = None

        # set to None if outcome (nullable) is None
        # and model_fields_set contains the field
        if self.outcome is None and "outcome" in self.model_fields_set:
            _dict['outcome'] = None

        # set to None if pal_grade (nullable) is None
        # and model_fields_set contains the field
        if self.pal_grade is None and "pal_grade" in self.model_fields_set:
            _dict['pal_grade'] = None

        # set to None if pal_image (nullable) is None
        # and model_fields_set contains the field
        if self.pal_image is None and "pal_image" in self.model_fields_set:
            _dict['pal_image'] = None

        # set to None if pal_name (nullable) is None
        # and model_fields_set contains the field
        if self.pal_name is None and "pal_name" in self.model_fields_set:
            _dict['pal_name'] = None

        # set to None if places (nullable) is None
        # and model_fields_set contains the field
        if self.places is None and "places" in self.model_fields_set:
            _dict['places'] = None

        # set to None if race_id (nullable) is None
        # and model_fields_set contains the field
        if self.race_id is None and "race_id" in self.model_fields_set:
            _dict['race_id'] = None

        # set to None if rank (nullable) is None
        # and model_fields_set contains the field
        if self.rank is None and "rank" in self.model_fields_set:
            _dict['rank'] = None

        # set to None if result (nullable) is None
        # and model_fields_set contains the field
        if self.result is None and "result" in self.model_fields_set:
            _dict['result'] = None

        # set to None if token_id (nullable) is None
        # and model_fields_set contains the field
        if self.token_id is None and "token_id" in self.model_fields_set:
            _dict['token_id'] = None

        # set to None if transaction_code (nullable) is None
        # and model_fields_set contains the field
        if self.transaction_code is None and "transaction_code" in self.model_fields_set:
            _dict['transaction_code'] = None

        return _dict

    @classmethod
    def from_dict(cls, obj: Optional[Dict[str, Any]]) -> Optional[Self]:
        """Create an instance of PalActivityLogDetails from a dict"""
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return cls.model_validate(obj)

        _obj = cls.model_validate({
            "attacking": obj.get("attacking"),
            "cost": obj.get("cost"),
            "gacha_type": obj.get("gacha_type"),
            "item_status": obj.get("item_status"),
            "item_subtype": obj.get("item_subtype"),
            "item_type": obj.get("item_type"),
            "level": obj.get("level"),
            "nft_type": obj.get("nft_type"),
            "outcome": obj.get("outcome"),
            "pal_grade": obj.get("pal_grade"),
            "pal_image": obj.get("pal_image"),
            "pal_name": obj.get("pal_name"),
            "places": obj.get("places"),
            "race_id": obj.get("race_id"),
            "rank": obj.get("rank"),
            "result": obj.get("result"),
            "token_id": obj.get("token_id"),
            "transaction_code": obj.get("transaction_code")
        })
        return _obj


