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
from yaylib.models.empl_dto import EmplDTO
from yaylib.models.pal_details_dto import PalDetailsDTO
from typing import Optional, Set
from typing_extensions import Self

class EmplActivityDTO(BaseModel):
    """
    EmplActivityDTO
    """ # noqa: E501
    amount: Optional[StrictStr] = None
    avatar_frame_user_uuid: Optional[StrictStr] = None
    battle_history_id: Optional[StrictInt] = None
    chain_id: Optional[StrictInt] = None
    details: Optional[PalDetailsDTO] = None
    empl: Optional[EmplDTO] = None
    gift_exchange_uuid: Optional[StrictStr] = None
    gift_transaction_uuid: Optional[StrictStr] = None
    id: Optional[StrictInt] = None
    money_amount: Optional[Union[StrictFloat, StrictInt]] = None
    occurred_at: Optional[StrictInt] = None
    price: Optional[Union[StrictFloat, StrictInt]] = None
    race_id: Optional[StrictStr] = None
    status: Optional[StrictStr] = None
    token_address: Optional[StrictStr] = None
    token_symbol: Optional[StrictStr] = None
    transaction_code: Optional[StrictStr] = None
    type: Optional[StrictStr] = None
    __properties: ClassVar[List[str]] = ["amount", "avatar_frame_user_uuid", "battle_history_id", "chain_id", "details", "empl", "gift_exchange_uuid", "gift_transaction_uuid", "id", "money_amount", "occurred_at", "price", "race_id", "status", "token_address", "token_symbol", "transaction_code", "type"]

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
        """Create an instance of EmplActivityDTO from a JSON string"""
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
        # override the default output from pydantic by calling `to_dict()` of details
        if self.details:
            _dict['details'] = self.details.to_dict()
        # override the default output from pydantic by calling `to_dict()` of empl
        if self.empl:
            _dict['empl'] = self.empl.to_dict()
        # set to None if amount (nullable) is None
        # and model_fields_set contains the field
        if self.amount is None and "amount" in self.model_fields_set:
            _dict['amount'] = None

        # set to None if avatar_frame_user_uuid (nullable) is None
        # and model_fields_set contains the field
        if self.avatar_frame_user_uuid is None and "avatar_frame_user_uuid" in self.model_fields_set:
            _dict['avatar_frame_user_uuid'] = None

        # set to None if battle_history_id (nullable) is None
        # and model_fields_set contains the field
        if self.battle_history_id is None and "battle_history_id" in self.model_fields_set:
            _dict['battle_history_id'] = None

        # set to None if chain_id (nullable) is None
        # and model_fields_set contains the field
        if self.chain_id is None and "chain_id" in self.model_fields_set:
            _dict['chain_id'] = None

        # set to None if details (nullable) is None
        # and model_fields_set contains the field
        if self.details is None and "details" in self.model_fields_set:
            _dict['details'] = None

        # set to None if empl (nullable) is None
        # and model_fields_set contains the field
        if self.empl is None and "empl" in self.model_fields_set:
            _dict['empl'] = None

        # set to None if gift_exchange_uuid (nullable) is None
        # and model_fields_set contains the field
        if self.gift_exchange_uuid is None and "gift_exchange_uuid" in self.model_fields_set:
            _dict['gift_exchange_uuid'] = None

        # set to None if gift_transaction_uuid (nullable) is None
        # and model_fields_set contains the field
        if self.gift_transaction_uuid is None and "gift_transaction_uuid" in self.model_fields_set:
            _dict['gift_transaction_uuid'] = None

        # set to None if id (nullable) is None
        # and model_fields_set contains the field
        if self.id is None and "id" in self.model_fields_set:
            _dict['id'] = None

        # set to None if money_amount (nullable) is None
        # and model_fields_set contains the field
        if self.money_amount is None and "money_amount" in self.model_fields_set:
            _dict['money_amount'] = None

        # set to None if occurred_at (nullable) is None
        # and model_fields_set contains the field
        if self.occurred_at is None and "occurred_at" in self.model_fields_set:
            _dict['occurred_at'] = None

        # set to None if price (nullable) is None
        # and model_fields_set contains the field
        if self.price is None and "price" in self.model_fields_set:
            _dict['price'] = None

        # set to None if race_id (nullable) is None
        # and model_fields_set contains the field
        if self.race_id is None and "race_id" in self.model_fields_set:
            _dict['race_id'] = None

        # set to None if status (nullable) is None
        # and model_fields_set contains the field
        if self.status is None and "status" in self.model_fields_set:
            _dict['status'] = None

        # set to None if token_address (nullable) is None
        # and model_fields_set contains the field
        if self.token_address is None and "token_address" in self.model_fields_set:
            _dict['token_address'] = None

        # set to None if token_symbol (nullable) is None
        # and model_fields_set contains the field
        if self.token_symbol is None and "token_symbol" in self.model_fields_set:
            _dict['token_symbol'] = None

        # set to None if transaction_code (nullable) is None
        # and model_fields_set contains the field
        if self.transaction_code is None and "transaction_code" in self.model_fields_set:
            _dict['transaction_code'] = None

        # set to None if type (nullable) is None
        # and model_fields_set contains the field
        if self.type is None and "type" in self.model_fields_set:
            _dict['type'] = None

        return _dict

    @classmethod
    def from_dict(cls, obj: Optional[Dict[str, Any]]) -> Optional[Self]:
        """Create an instance of EmplActivityDTO from a dict"""
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return cls.model_validate(obj)

        _obj = cls.model_validate({
            "amount": obj.get("amount"),
            "avatar_frame_user_uuid": obj.get("avatar_frame_user_uuid"),
            "battle_history_id": obj.get("battle_history_id"),
            "chain_id": obj.get("chain_id"),
            "details": PalDetailsDTO.from_dict(obj["details"]) if obj.get("details") is not None else None,
            "empl": EmplDTO.from_dict(obj["empl"]) if obj.get("empl") is not None else None,
            "gift_exchange_uuid": obj.get("gift_exchange_uuid"),
            "gift_transaction_uuid": obj.get("gift_transaction_uuid"),
            "id": obj.get("id"),
            "money_amount": obj.get("money_amount"),
            "occurred_at": obj.get("occurred_at"),
            "price": obj.get("price"),
            "race_id": obj.get("race_id"),
            "status": obj.get("status"),
            "token_address": obj.get("token_address"),
            "token_symbol": obj.get("token_symbol"),
            "transaction_code": obj.get("transaction_code"),
            "type": obj.get("type")
        })
        return _obj


