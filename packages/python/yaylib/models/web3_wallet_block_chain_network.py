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
from yaylib.models.model_web3_wallet_gas_percent import ModelWeb3WalletGasPercent
from typing import Optional, Set
from typing_extensions import Self

class Web3WalletBlockChainNetwork(BaseModel):
    """
    Web3WalletBlockChainNetwork
    """ # noqa: E501
    l1_amm_address: Optional[StrictStr] = None
    l1_block_chain_network_addition_gas_percent: Optional[ModelWeb3WalletGasPercent] = None
    l1_block_chain_network_id: Optional[StrictInt] = None
    l1_block_chain_network_name: Optional[StrictStr] = None
    l1_block_chain_network_url: Optional[StrictStr] = None
    l1_quoter_address: Optional[StrictStr] = None
    l1_weth_address: Optional[StrictStr] = None
    l2_amm_address: Optional[StrictStr] = None
    l2_block_chain_network_addition_gas_percent: Optional[ModelWeb3WalletGasPercent] = None
    l2_block_chain_network_id: Optional[StrictInt] = None
    l2_block_chain_network_name: Optional[StrictStr] = None
    l2_block_chain_network_url: Optional[StrictStr] = None
    l2_quoter_address: Optional[StrictStr] = None
    l2_weth_address: Optional[StrictStr] = None
    misc_pal_base_image_uri: Optional[StrictStr] = None
    misc_pal_base_json_uri: Optional[StrictStr] = None
    smart_contract_empl_address: Optional[StrictStr] = None
    smart_contract_empl_deposit_address: Optional[StrictStr] = None
    smart_contract_empl_withdraw_address: Optional[StrictStr] = None
    smart_contract_l1_yay_address: Optional[StrictStr] = None
    smart_contract_l2_yay_address: Optional[StrictStr] = None
    smart_contract_pal_address: Optional[StrictStr] = None
    smart_contract_pal_deposit_withdraw_address: Optional[StrictStr] = None
    smart_contract_x2y2delegateerc721_address: Optional[StrictStr] = None
    smart_contract_x2y2market_address: Optional[StrictStr] = None
    __properties: ClassVar[List[str]] = ["l1_amm_address", "l1_block_chain_network_addition_gas_percent", "l1_block_chain_network_id", "l1_block_chain_network_name", "l1_block_chain_network_url", "l1_quoter_address", "l1_weth_address", "l2_amm_address", "l2_block_chain_network_addition_gas_percent", "l2_block_chain_network_id", "l2_block_chain_network_name", "l2_block_chain_network_url", "l2_quoter_address", "l2_weth_address", "misc_pal_base_image_uri", "misc_pal_base_json_uri", "smart_contract_empl_address", "smart_contract_empl_deposit_address", "smart_contract_empl_withdraw_address", "smart_contract_l1_yay_address", "smart_contract_l2_yay_address", "smart_contract_pal_address", "smart_contract_pal_deposit_withdraw_address", "smart_contract_x2y2delegateerc721_address", "smart_contract_x2y2market_address"]

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
        """Create an instance of Web3WalletBlockChainNetwork from a JSON string"""
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
        # override the default output from pydantic by calling `to_dict()` of l1_block_chain_network_addition_gas_percent
        if self.l1_block_chain_network_addition_gas_percent:
            _dict['l1_block_chain_network_addition_gas_percent'] = self.l1_block_chain_network_addition_gas_percent.to_dict()
        # override the default output from pydantic by calling `to_dict()` of l2_block_chain_network_addition_gas_percent
        if self.l2_block_chain_network_addition_gas_percent:
            _dict['l2_block_chain_network_addition_gas_percent'] = self.l2_block_chain_network_addition_gas_percent.to_dict()
        # set to None if l1_amm_address (nullable) is None
        # and model_fields_set contains the field
        if self.l1_amm_address is None and "l1_amm_address" in self.model_fields_set:
            _dict['l1_amm_address'] = None

        # set to None if l1_block_chain_network_addition_gas_percent (nullable) is None
        # and model_fields_set contains the field
        if self.l1_block_chain_network_addition_gas_percent is None and "l1_block_chain_network_addition_gas_percent" in self.model_fields_set:
            _dict['l1_block_chain_network_addition_gas_percent'] = None

        # set to None if l1_block_chain_network_id (nullable) is None
        # and model_fields_set contains the field
        if self.l1_block_chain_network_id is None and "l1_block_chain_network_id" in self.model_fields_set:
            _dict['l1_block_chain_network_id'] = None

        # set to None if l1_block_chain_network_name (nullable) is None
        # and model_fields_set contains the field
        if self.l1_block_chain_network_name is None and "l1_block_chain_network_name" in self.model_fields_set:
            _dict['l1_block_chain_network_name'] = None

        # set to None if l1_block_chain_network_url (nullable) is None
        # and model_fields_set contains the field
        if self.l1_block_chain_network_url is None and "l1_block_chain_network_url" in self.model_fields_set:
            _dict['l1_block_chain_network_url'] = None

        # set to None if l1_quoter_address (nullable) is None
        # and model_fields_set contains the field
        if self.l1_quoter_address is None and "l1_quoter_address" in self.model_fields_set:
            _dict['l1_quoter_address'] = None

        # set to None if l1_weth_address (nullable) is None
        # and model_fields_set contains the field
        if self.l1_weth_address is None and "l1_weth_address" in self.model_fields_set:
            _dict['l1_weth_address'] = None

        # set to None if l2_amm_address (nullable) is None
        # and model_fields_set contains the field
        if self.l2_amm_address is None and "l2_amm_address" in self.model_fields_set:
            _dict['l2_amm_address'] = None

        # set to None if l2_block_chain_network_addition_gas_percent (nullable) is None
        # and model_fields_set contains the field
        if self.l2_block_chain_network_addition_gas_percent is None and "l2_block_chain_network_addition_gas_percent" in self.model_fields_set:
            _dict['l2_block_chain_network_addition_gas_percent'] = None

        # set to None if l2_block_chain_network_id (nullable) is None
        # and model_fields_set contains the field
        if self.l2_block_chain_network_id is None and "l2_block_chain_network_id" in self.model_fields_set:
            _dict['l2_block_chain_network_id'] = None

        # set to None if l2_block_chain_network_name (nullable) is None
        # and model_fields_set contains the field
        if self.l2_block_chain_network_name is None and "l2_block_chain_network_name" in self.model_fields_set:
            _dict['l2_block_chain_network_name'] = None

        # set to None if l2_block_chain_network_url (nullable) is None
        # and model_fields_set contains the field
        if self.l2_block_chain_network_url is None and "l2_block_chain_network_url" in self.model_fields_set:
            _dict['l2_block_chain_network_url'] = None

        # set to None if l2_quoter_address (nullable) is None
        # and model_fields_set contains the field
        if self.l2_quoter_address is None and "l2_quoter_address" in self.model_fields_set:
            _dict['l2_quoter_address'] = None

        # set to None if l2_weth_address (nullable) is None
        # and model_fields_set contains the field
        if self.l2_weth_address is None and "l2_weth_address" in self.model_fields_set:
            _dict['l2_weth_address'] = None

        # set to None if misc_pal_base_image_uri (nullable) is None
        # and model_fields_set contains the field
        if self.misc_pal_base_image_uri is None and "misc_pal_base_image_uri" in self.model_fields_set:
            _dict['misc_pal_base_image_uri'] = None

        # set to None if misc_pal_base_json_uri (nullable) is None
        # and model_fields_set contains the field
        if self.misc_pal_base_json_uri is None and "misc_pal_base_json_uri" in self.model_fields_set:
            _dict['misc_pal_base_json_uri'] = None

        # set to None if smart_contract_empl_address (nullable) is None
        # and model_fields_set contains the field
        if self.smart_contract_empl_address is None and "smart_contract_empl_address" in self.model_fields_set:
            _dict['smart_contract_empl_address'] = None

        # set to None if smart_contract_empl_deposit_address (nullable) is None
        # and model_fields_set contains the field
        if self.smart_contract_empl_deposit_address is None and "smart_contract_empl_deposit_address" in self.model_fields_set:
            _dict['smart_contract_empl_deposit_address'] = None

        # set to None if smart_contract_empl_withdraw_address (nullable) is None
        # and model_fields_set contains the field
        if self.smart_contract_empl_withdraw_address is None and "smart_contract_empl_withdraw_address" in self.model_fields_set:
            _dict['smart_contract_empl_withdraw_address'] = None

        # set to None if smart_contract_l1_yay_address (nullable) is None
        # and model_fields_set contains the field
        if self.smart_contract_l1_yay_address is None and "smart_contract_l1_yay_address" in self.model_fields_set:
            _dict['smart_contract_l1_yay_address'] = None

        # set to None if smart_contract_l2_yay_address (nullable) is None
        # and model_fields_set contains the field
        if self.smart_contract_l2_yay_address is None and "smart_contract_l2_yay_address" in self.model_fields_set:
            _dict['smart_contract_l2_yay_address'] = None

        # set to None if smart_contract_pal_address (nullable) is None
        # and model_fields_set contains the field
        if self.smart_contract_pal_address is None and "smart_contract_pal_address" in self.model_fields_set:
            _dict['smart_contract_pal_address'] = None

        # set to None if smart_contract_pal_deposit_withdraw_address (nullable) is None
        # and model_fields_set contains the field
        if self.smart_contract_pal_deposit_withdraw_address is None and "smart_contract_pal_deposit_withdraw_address" in self.model_fields_set:
            _dict['smart_contract_pal_deposit_withdraw_address'] = None

        # set to None if smart_contract_x2y2delegateerc721_address (nullable) is None
        # and model_fields_set contains the field
        if self.smart_contract_x2y2delegateerc721_address is None and "smart_contract_x2y2delegateerc721_address" in self.model_fields_set:
            _dict['smart_contract_x2y2delegateerc721_address'] = None

        # set to None if smart_contract_x2y2market_address (nullable) is None
        # and model_fields_set contains the field
        if self.smart_contract_x2y2market_address is None and "smart_contract_x2y2market_address" in self.model_fields_set:
            _dict['smart_contract_x2y2market_address'] = None

        return _dict

    @classmethod
    def from_dict(cls, obj: Optional[Dict[str, Any]]) -> Optional[Self]:
        """Create an instance of Web3WalletBlockChainNetwork from a dict"""
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return cls.model_validate(obj)

        _obj = cls.model_validate({
            "l1_amm_address": obj.get("l1_amm_address"),
            "l1_block_chain_network_addition_gas_percent": ModelWeb3WalletGasPercent.from_dict(obj["l1_block_chain_network_addition_gas_percent"]) if obj.get("l1_block_chain_network_addition_gas_percent") is not None else None,
            "l1_block_chain_network_id": obj.get("l1_block_chain_network_id"),
            "l1_block_chain_network_name": obj.get("l1_block_chain_network_name"),
            "l1_block_chain_network_url": obj.get("l1_block_chain_network_url"),
            "l1_quoter_address": obj.get("l1_quoter_address"),
            "l1_weth_address": obj.get("l1_weth_address"),
            "l2_amm_address": obj.get("l2_amm_address"),
            "l2_block_chain_network_addition_gas_percent": ModelWeb3WalletGasPercent.from_dict(obj["l2_block_chain_network_addition_gas_percent"]) if obj.get("l2_block_chain_network_addition_gas_percent") is not None else None,
            "l2_block_chain_network_id": obj.get("l2_block_chain_network_id"),
            "l2_block_chain_network_name": obj.get("l2_block_chain_network_name"),
            "l2_block_chain_network_url": obj.get("l2_block_chain_network_url"),
            "l2_quoter_address": obj.get("l2_quoter_address"),
            "l2_weth_address": obj.get("l2_weth_address"),
            "misc_pal_base_image_uri": obj.get("misc_pal_base_image_uri"),
            "misc_pal_base_json_uri": obj.get("misc_pal_base_json_uri"),
            "smart_contract_empl_address": obj.get("smart_contract_empl_address"),
            "smart_contract_empl_deposit_address": obj.get("smart_contract_empl_deposit_address"),
            "smart_contract_empl_withdraw_address": obj.get("smart_contract_empl_withdraw_address"),
            "smart_contract_l1_yay_address": obj.get("smart_contract_l1_yay_address"),
            "smart_contract_l2_yay_address": obj.get("smart_contract_l2_yay_address"),
            "smart_contract_pal_address": obj.get("smart_contract_pal_address"),
            "smart_contract_pal_deposit_withdraw_address": obj.get("smart_contract_pal_deposit_withdraw_address"),
            "smart_contract_x2y2delegateerc721_address": obj.get("smart_contract_x2y2delegateerc721_address"),
            "smart_contract_x2y2market_address": obj.get("smart_contract_x2y2market_address")
        })
        return _obj


