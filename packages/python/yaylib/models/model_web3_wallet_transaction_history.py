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

from pydantic import BaseModel, ConfigDict, StrictBool, StrictFloat, StrictInt, StrictStr
from typing import Any, ClassVar, Dict, List, Optional, Union
from yaylib.models.asset_type import AssetType
from yaylib.models.empl_details import EmplDetails
from yaylib.models.gift_transaction_detail import GiftTransactionDetail
from yaylib.models.web3_wallet_transaction_history_asset_info import Web3WalletTransactionHistoryAssetInfo
from typing import Optional, Set
from typing_extensions import Self

class ModelWeb3WalletTransactionHistory(BaseModel):
    """
    ModelWeb3WalletTransactionHistory
    """ # noqa: E501
    asset_info: Optional[Web3WalletTransactionHistoryAssetInfo] = None
    asset_type: Optional[AssetType] = None
    created_at_millis: Optional[StrictInt] = None
    currency: Optional[StrictStr] = None
    dot_money_amount: Optional[Union[StrictFloat, StrictInt]] = None
    empl_details: Optional[EmplDetails] = None
    from_address: Optional[StrictStr] = None
    gift_exchange_uuid: Optional[StrictStr] = None
    gift_transaction_detail: Optional[GiftTransactionDetail] = None
    gift_transaction_uuid: Optional[StrictStr] = None
    is_pending_transaction: Optional[StrictBool] = None
    nft_token_id: Optional[StrictStr] = None
    pal_transaction_detail: Optional[Dict[str, Any]] = None
    price: Optional[Union[StrictFloat, StrictInt]] = None
    race_id: Optional[StrictStr] = None
    race_ranks: Optional[List[StrictInt]] = None
    status: Optional[Dict[str, Any]] = None
    to_address: Optional[StrictStr] = None
    token_symbol: Optional[StrictStr] = None
    transaction_code: Optional[StrictStr] = None
    transaction_hash_hex: Optional[StrictStr] = None
    transaction_history_type: Optional[Dict[str, Any]] = None
    transaction_type: Optional[Dict[str, Any]] = None
    value: Optional[Union[StrictFloat, StrictInt]] = None
    value_symbol: Optional[StrictStr] = None
    value_text: Optional[StrictStr] = None
    __properties: ClassVar[List[str]] = ["asset_info", "asset_type", "created_at_millis", "currency", "dot_money_amount", "empl_details", "from_address", "gift_exchange_uuid", "gift_transaction_detail", "gift_transaction_uuid", "is_pending_transaction", "nft_token_id", "pal_transaction_detail", "price", "race_id", "race_ranks", "status", "to_address", "token_symbol", "transaction_code", "transaction_hash_hex", "transaction_history_type", "transaction_type", "value", "value_symbol", "value_text"]

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
        """Create an instance of ModelWeb3WalletTransactionHistory from a JSON string"""
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
        # override the default output from pydantic by calling `to_dict()` of asset_info
        if self.asset_info:
            _dict['asset_info'] = self.asset_info.to_dict()
        # override the default output from pydantic by calling `to_dict()` of asset_type
        if self.asset_type:
            _dict['asset_type'] = self.asset_type.to_dict()
        # override the default output from pydantic by calling `to_dict()` of empl_details
        if self.empl_details:
            _dict['empl_details'] = self.empl_details.to_dict()
        # override the default output from pydantic by calling `to_dict()` of gift_transaction_detail
        if self.gift_transaction_detail:
            _dict['gift_transaction_detail'] = self.gift_transaction_detail.to_dict()
        # set to None if asset_info (nullable) is None
        # and model_fields_set contains the field
        if self.asset_info is None and "asset_info" in self.model_fields_set:
            _dict['asset_info'] = None

        # set to None if asset_type (nullable) is None
        # and model_fields_set contains the field
        if self.asset_type is None and "asset_type" in self.model_fields_set:
            _dict['asset_type'] = None

        # set to None if created_at_millis (nullable) is None
        # and model_fields_set contains the field
        if self.created_at_millis is None and "created_at_millis" in self.model_fields_set:
            _dict['created_at_millis'] = None

        # set to None if currency (nullable) is None
        # and model_fields_set contains the field
        if self.currency is None and "currency" in self.model_fields_set:
            _dict['currency'] = None

        # set to None if dot_money_amount (nullable) is None
        # and model_fields_set contains the field
        if self.dot_money_amount is None and "dot_money_amount" in self.model_fields_set:
            _dict['dot_money_amount'] = None

        # set to None if empl_details (nullable) is None
        # and model_fields_set contains the field
        if self.empl_details is None and "empl_details" in self.model_fields_set:
            _dict['empl_details'] = None

        # set to None if from_address (nullable) is None
        # and model_fields_set contains the field
        if self.from_address is None and "from_address" in self.model_fields_set:
            _dict['from_address'] = None

        # set to None if gift_exchange_uuid (nullable) is None
        # and model_fields_set contains the field
        if self.gift_exchange_uuid is None and "gift_exchange_uuid" in self.model_fields_set:
            _dict['gift_exchange_uuid'] = None

        # set to None if gift_transaction_detail (nullable) is None
        # and model_fields_set contains the field
        if self.gift_transaction_detail is None and "gift_transaction_detail" in self.model_fields_set:
            _dict['gift_transaction_detail'] = None

        # set to None if gift_transaction_uuid (nullable) is None
        # and model_fields_set contains the field
        if self.gift_transaction_uuid is None and "gift_transaction_uuid" in self.model_fields_set:
            _dict['gift_transaction_uuid'] = None

        # set to None if is_pending_transaction (nullable) is None
        # and model_fields_set contains the field
        if self.is_pending_transaction is None and "is_pending_transaction" in self.model_fields_set:
            _dict['is_pending_transaction'] = None

        # set to None if nft_token_id (nullable) is None
        # and model_fields_set contains the field
        if self.nft_token_id is None and "nft_token_id" in self.model_fields_set:
            _dict['nft_token_id'] = None

        # set to None if pal_transaction_detail (nullable) is None
        # and model_fields_set contains the field
        if self.pal_transaction_detail is None and "pal_transaction_detail" in self.model_fields_set:
            _dict['pal_transaction_detail'] = None

        # set to None if price (nullable) is None
        # and model_fields_set contains the field
        if self.price is None and "price" in self.model_fields_set:
            _dict['price'] = None

        # set to None if race_id (nullable) is None
        # and model_fields_set contains the field
        if self.race_id is None and "race_id" in self.model_fields_set:
            _dict['race_id'] = None

        # set to None if race_ranks (nullable) is None
        # and model_fields_set contains the field
        if self.race_ranks is None and "race_ranks" in self.model_fields_set:
            _dict['race_ranks'] = None

        # set to None if status (nullable) is None
        # and model_fields_set contains the field
        if self.status is None and "status" in self.model_fields_set:
            _dict['status'] = None

        # set to None if to_address (nullable) is None
        # and model_fields_set contains the field
        if self.to_address is None and "to_address" in self.model_fields_set:
            _dict['to_address'] = None

        # set to None if token_symbol (nullable) is None
        # and model_fields_set contains the field
        if self.token_symbol is None and "token_symbol" in self.model_fields_set:
            _dict['token_symbol'] = None

        # set to None if transaction_code (nullable) is None
        # and model_fields_set contains the field
        if self.transaction_code is None and "transaction_code" in self.model_fields_set:
            _dict['transaction_code'] = None

        # set to None if transaction_hash_hex (nullable) is None
        # and model_fields_set contains the field
        if self.transaction_hash_hex is None and "transaction_hash_hex" in self.model_fields_set:
            _dict['transaction_hash_hex'] = None

        # set to None if transaction_history_type (nullable) is None
        # and model_fields_set contains the field
        if self.transaction_history_type is None and "transaction_history_type" in self.model_fields_set:
            _dict['transaction_history_type'] = None

        # set to None if transaction_type (nullable) is None
        # and model_fields_set contains the field
        if self.transaction_type is None and "transaction_type" in self.model_fields_set:
            _dict['transaction_type'] = None

        # set to None if value (nullable) is None
        # and model_fields_set contains the field
        if self.value is None and "value" in self.model_fields_set:
            _dict['value'] = None

        # set to None if value_symbol (nullable) is None
        # and model_fields_set contains the field
        if self.value_symbol is None and "value_symbol" in self.model_fields_set:
            _dict['value_symbol'] = None

        # set to None if value_text (nullable) is None
        # and model_fields_set contains the field
        if self.value_text is None and "value_text" in self.model_fields_set:
            _dict['value_text'] = None

        return _dict

    @classmethod
    def from_dict(cls, obj: Optional[Dict[str, Any]]) -> Optional[Self]:
        """Create an instance of ModelWeb3WalletTransactionHistory from a dict"""
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return cls.model_validate(obj)

        _obj = cls.model_validate({
            "asset_info": Web3WalletTransactionHistoryAssetInfo.from_dict(obj["asset_info"]) if obj.get("asset_info") is not None else None,
            "asset_type": AssetType.from_dict(obj["asset_type"]) if obj.get("asset_type") is not None else None,
            "created_at_millis": obj.get("created_at_millis"),
            "currency": obj.get("currency"),
            "dot_money_amount": obj.get("dot_money_amount"),
            "empl_details": EmplDetails.from_dict(obj["empl_details"]) if obj.get("empl_details") is not None else None,
            "from_address": obj.get("from_address"),
            "gift_exchange_uuid": obj.get("gift_exchange_uuid"),
            "gift_transaction_detail": GiftTransactionDetail.from_dict(obj["gift_transaction_detail"]) if obj.get("gift_transaction_detail") is not None else None,
            "gift_transaction_uuid": obj.get("gift_transaction_uuid"),
            "is_pending_transaction": obj.get("is_pending_transaction"),
            "nft_token_id": obj.get("nft_token_id"),
            "pal_transaction_detail": obj.get("pal_transaction_detail"),
            "price": obj.get("price"),
            "race_id": obj.get("race_id"),
            "race_ranks": obj.get("race_ranks"),
            "status": obj.get("status"),
            "to_address": obj.get("to_address"),
            "token_symbol": obj.get("token_symbol"),
            "transaction_code": obj.get("transaction_code"),
            "transaction_hash_hex": obj.get("transaction_hash_hex"),
            "transaction_history_type": obj.get("transaction_history_type"),
            "transaction_type": obj.get("transaction_type"),
            "value": obj.get("value"),
            "value_symbol": obj.get("value_symbol"),
            "value_text": obj.get("value_text")
        })
        return _obj


