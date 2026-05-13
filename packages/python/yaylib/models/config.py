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

from pydantic import BaseModel, ConfigDict, StrictStr
from typing import Any, ClassVar, Dict, List, Optional
from typing import Optional, Set
from typing_extensions import Self

class Config(BaseModel):
    """
    Config
    """ # noqa: E501
    ad_tester_user_ids: Optional[StrictStr] = None
    android_news_version: Optional[StrictStr] = None
    dapps: Optional[StrictStr] = None
    is_chat_websocket_enabled: Optional[StrictStr] = None
    is_direct_vip_purchase_enabled: Optional[StrictStr] = None
    is_gift_features_enabled: Optional[StrictStr] = None
    is_id_card_and_phone_verification_check_for_review_enabled: Optional[StrictStr] = None
    is_id_card_check_required_blocker_view_enabled: Optional[StrictStr] = None
    is_maintenanced: Optional[StrictStr] = None
    is_phone_verification_required_blocker_view_enabled: Optional[StrictStr] = None
    is_speed_test_enabled: Optional[StrictStr] = None
    is_ssl_pinning_enabled: Optional[StrictStr] = None
    is_star_enabled: Optional[StrictStr] = None
    latest_android_app_version: Optional[StrictStr] = None
    line_official_account_id: Optional[StrictStr] = None
    localized_android_news_title: Optional[StrictStr] = None
    localized_android_news_url: Optional[StrictStr] = None
    localized_maintenance_url: Optional[StrictStr] = None
    localized_news_title: Optional[StrictStr] = None
    localized_news_url: Optional[StrictStr] = None
    maintenance_features: Optional[StrictStr] = None
    max_image_frame_count: Optional[StrictStr] = None
    minimum_android_app_version_required: Optional[StrictStr] = None
    minimum_android_version_supported: Optional[StrictStr] = None
    news_version: Optional[StrictStr] = None
    nft_infos: Optional[StrictStr] = None
    promotion_sticker_pack_ids: Optional[StrictStr] = None
    should_append_user_id_to_news_url: Optional[StrictStr] = None
    support_email_address: Optional[StrictStr] = None
    token_infos: Optional[StrictStr] = None
    use_random_message_refresh_rate: Optional[StrictStr] = None
    web3_is_maintenanced: Optional[StrictStr] = None
    web3_localized_maintenance_url: Optional[StrictStr] = None
    web3_maintenance_features: Optional[StrictStr] = None
    web3_maintenance_url: Optional[StrictStr] = None
    x_yay_global_id: Optional[StrictStr] = None
    x_yay_jp_id: Optional[StrictStr] = None
    x_yay_update_id: Optional[StrictStr] = None
    __properties: ClassVar[List[str]] = ["ad_tester_user_ids", "android_news_version", "dapps", "is_chat_websocket_enabled", "is_direct_vip_purchase_enabled", "is_gift_features_enabled", "is_id_card_and_phone_verification_check_for_review_enabled", "is_id_card_check_required_blocker_view_enabled", "is_maintenanced", "is_phone_verification_required_blocker_view_enabled", "is_speed_test_enabled", "is_ssl_pinning_enabled", "is_star_enabled", "latest_android_app_version", "line_official_account_id", "localized_android_news_title", "localized_android_news_url", "localized_maintenance_url", "localized_news_title", "localized_news_url", "maintenance_features", "max_image_frame_count", "minimum_android_app_version_required", "minimum_android_version_supported", "news_version", "nft_infos", "promotion_sticker_pack_ids", "should_append_user_id_to_news_url", "support_email_address", "token_infos", "use_random_message_refresh_rate", "web3_is_maintenanced", "web3_localized_maintenance_url", "web3_maintenance_features", "web3_maintenance_url", "x_yay_global_id", "x_yay_jp_id", "x_yay_update_id"]

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
        """Create an instance of Config from a JSON string"""
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
        # set to None if ad_tester_user_ids (nullable) is None
        # and model_fields_set contains the field
        if self.ad_tester_user_ids is None and "ad_tester_user_ids" in self.model_fields_set:
            _dict['ad_tester_user_ids'] = None

        # set to None if android_news_version (nullable) is None
        # and model_fields_set contains the field
        if self.android_news_version is None and "android_news_version" in self.model_fields_set:
            _dict['android_news_version'] = None

        # set to None if dapps (nullable) is None
        # and model_fields_set contains the field
        if self.dapps is None and "dapps" in self.model_fields_set:
            _dict['dapps'] = None

        # set to None if is_chat_websocket_enabled (nullable) is None
        # and model_fields_set contains the field
        if self.is_chat_websocket_enabled is None and "is_chat_websocket_enabled" in self.model_fields_set:
            _dict['is_chat_websocket_enabled'] = None

        # set to None if is_direct_vip_purchase_enabled (nullable) is None
        # and model_fields_set contains the field
        if self.is_direct_vip_purchase_enabled is None and "is_direct_vip_purchase_enabled" in self.model_fields_set:
            _dict['is_direct_vip_purchase_enabled'] = None

        # set to None if is_gift_features_enabled (nullable) is None
        # and model_fields_set contains the field
        if self.is_gift_features_enabled is None and "is_gift_features_enabled" in self.model_fields_set:
            _dict['is_gift_features_enabled'] = None

        # set to None if is_id_card_and_phone_verification_check_for_review_enabled (nullable) is None
        # and model_fields_set contains the field
        if self.is_id_card_and_phone_verification_check_for_review_enabled is None and "is_id_card_and_phone_verification_check_for_review_enabled" in self.model_fields_set:
            _dict['is_id_card_and_phone_verification_check_for_review_enabled'] = None

        # set to None if is_id_card_check_required_blocker_view_enabled (nullable) is None
        # and model_fields_set contains the field
        if self.is_id_card_check_required_blocker_view_enabled is None and "is_id_card_check_required_blocker_view_enabled" in self.model_fields_set:
            _dict['is_id_card_check_required_blocker_view_enabled'] = None

        # set to None if is_maintenanced (nullable) is None
        # and model_fields_set contains the field
        if self.is_maintenanced is None and "is_maintenanced" in self.model_fields_set:
            _dict['is_maintenanced'] = None

        # set to None if is_phone_verification_required_blocker_view_enabled (nullable) is None
        # and model_fields_set contains the field
        if self.is_phone_verification_required_blocker_view_enabled is None and "is_phone_verification_required_blocker_view_enabled" in self.model_fields_set:
            _dict['is_phone_verification_required_blocker_view_enabled'] = None

        # set to None if is_speed_test_enabled (nullable) is None
        # and model_fields_set contains the field
        if self.is_speed_test_enabled is None and "is_speed_test_enabled" in self.model_fields_set:
            _dict['is_speed_test_enabled'] = None

        # set to None if is_ssl_pinning_enabled (nullable) is None
        # and model_fields_set contains the field
        if self.is_ssl_pinning_enabled is None and "is_ssl_pinning_enabled" in self.model_fields_set:
            _dict['is_ssl_pinning_enabled'] = None

        # set to None if is_star_enabled (nullable) is None
        # and model_fields_set contains the field
        if self.is_star_enabled is None and "is_star_enabled" in self.model_fields_set:
            _dict['is_star_enabled'] = None

        # set to None if latest_android_app_version (nullable) is None
        # and model_fields_set contains the field
        if self.latest_android_app_version is None and "latest_android_app_version" in self.model_fields_set:
            _dict['latest_android_app_version'] = None

        # set to None if line_official_account_id (nullable) is None
        # and model_fields_set contains the field
        if self.line_official_account_id is None and "line_official_account_id" in self.model_fields_set:
            _dict['line_official_account_id'] = None

        # set to None if localized_android_news_title (nullable) is None
        # and model_fields_set contains the field
        if self.localized_android_news_title is None and "localized_android_news_title" in self.model_fields_set:
            _dict['localized_android_news_title'] = None

        # set to None if localized_android_news_url (nullable) is None
        # and model_fields_set contains the field
        if self.localized_android_news_url is None and "localized_android_news_url" in self.model_fields_set:
            _dict['localized_android_news_url'] = None

        # set to None if localized_maintenance_url (nullable) is None
        # and model_fields_set contains the field
        if self.localized_maintenance_url is None and "localized_maintenance_url" in self.model_fields_set:
            _dict['localized_maintenance_url'] = None

        # set to None if localized_news_title (nullable) is None
        # and model_fields_set contains the field
        if self.localized_news_title is None and "localized_news_title" in self.model_fields_set:
            _dict['localized_news_title'] = None

        # set to None if localized_news_url (nullable) is None
        # and model_fields_set contains the field
        if self.localized_news_url is None and "localized_news_url" in self.model_fields_set:
            _dict['localized_news_url'] = None

        # set to None if maintenance_features (nullable) is None
        # and model_fields_set contains the field
        if self.maintenance_features is None and "maintenance_features" in self.model_fields_set:
            _dict['maintenance_features'] = None

        # set to None if max_image_frame_count (nullable) is None
        # and model_fields_set contains the field
        if self.max_image_frame_count is None and "max_image_frame_count" in self.model_fields_set:
            _dict['max_image_frame_count'] = None

        # set to None if minimum_android_app_version_required (nullable) is None
        # and model_fields_set contains the field
        if self.minimum_android_app_version_required is None and "minimum_android_app_version_required" in self.model_fields_set:
            _dict['minimum_android_app_version_required'] = None

        # set to None if minimum_android_version_supported (nullable) is None
        # and model_fields_set contains the field
        if self.minimum_android_version_supported is None and "minimum_android_version_supported" in self.model_fields_set:
            _dict['minimum_android_version_supported'] = None

        # set to None if news_version (nullable) is None
        # and model_fields_set contains the field
        if self.news_version is None and "news_version" in self.model_fields_set:
            _dict['news_version'] = None

        # set to None if nft_infos (nullable) is None
        # and model_fields_set contains the field
        if self.nft_infos is None and "nft_infos" in self.model_fields_set:
            _dict['nft_infos'] = None

        # set to None if promotion_sticker_pack_ids (nullable) is None
        # and model_fields_set contains the field
        if self.promotion_sticker_pack_ids is None and "promotion_sticker_pack_ids" in self.model_fields_set:
            _dict['promotion_sticker_pack_ids'] = None

        # set to None if should_append_user_id_to_news_url (nullable) is None
        # and model_fields_set contains the field
        if self.should_append_user_id_to_news_url is None and "should_append_user_id_to_news_url" in self.model_fields_set:
            _dict['should_append_user_id_to_news_url'] = None

        # set to None if support_email_address (nullable) is None
        # and model_fields_set contains the field
        if self.support_email_address is None and "support_email_address" in self.model_fields_set:
            _dict['support_email_address'] = None

        # set to None if token_infos (nullable) is None
        # and model_fields_set contains the field
        if self.token_infos is None and "token_infos" in self.model_fields_set:
            _dict['token_infos'] = None

        # set to None if use_random_message_refresh_rate (nullable) is None
        # and model_fields_set contains the field
        if self.use_random_message_refresh_rate is None and "use_random_message_refresh_rate" in self.model_fields_set:
            _dict['use_random_message_refresh_rate'] = None

        # set to None if web3_is_maintenanced (nullable) is None
        # and model_fields_set contains the field
        if self.web3_is_maintenanced is None and "web3_is_maintenanced" in self.model_fields_set:
            _dict['web3_is_maintenanced'] = None

        # set to None if web3_localized_maintenance_url (nullable) is None
        # and model_fields_set contains the field
        if self.web3_localized_maintenance_url is None and "web3_localized_maintenance_url" in self.model_fields_set:
            _dict['web3_localized_maintenance_url'] = None

        # set to None if web3_maintenance_features (nullable) is None
        # and model_fields_set contains the field
        if self.web3_maintenance_features is None and "web3_maintenance_features" in self.model_fields_set:
            _dict['web3_maintenance_features'] = None

        # set to None if web3_maintenance_url (nullable) is None
        # and model_fields_set contains the field
        if self.web3_maintenance_url is None and "web3_maintenance_url" in self.model_fields_set:
            _dict['web3_maintenance_url'] = None

        # set to None if x_yay_global_id (nullable) is None
        # and model_fields_set contains the field
        if self.x_yay_global_id is None and "x_yay_global_id" in self.model_fields_set:
            _dict['x_yay_global_id'] = None

        # set to None if x_yay_jp_id (nullable) is None
        # and model_fields_set contains the field
        if self.x_yay_jp_id is None and "x_yay_jp_id" in self.model_fields_set:
            _dict['x_yay_jp_id'] = None

        # set to None if x_yay_update_id (nullable) is None
        # and model_fields_set contains the field
        if self.x_yay_update_id is None and "x_yay_update_id" in self.model_fields_set:
            _dict['x_yay_update_id'] = None

        return _dict

    @classmethod
    def from_dict(cls, obj: Optional[Dict[str, Any]]) -> Optional[Self]:
        """Create an instance of Config from a dict"""
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return cls.model_validate(obj)

        _obj = cls.model_validate({
            "ad_tester_user_ids": obj.get("ad_tester_user_ids"),
            "android_news_version": obj.get("android_news_version"),
            "dapps": obj.get("dapps"),
            "is_chat_websocket_enabled": obj.get("is_chat_websocket_enabled"),
            "is_direct_vip_purchase_enabled": obj.get("is_direct_vip_purchase_enabled"),
            "is_gift_features_enabled": obj.get("is_gift_features_enabled"),
            "is_id_card_and_phone_verification_check_for_review_enabled": obj.get("is_id_card_and_phone_verification_check_for_review_enabled"),
            "is_id_card_check_required_blocker_view_enabled": obj.get("is_id_card_check_required_blocker_view_enabled"),
            "is_maintenanced": obj.get("is_maintenanced"),
            "is_phone_verification_required_blocker_view_enabled": obj.get("is_phone_verification_required_blocker_view_enabled"),
            "is_speed_test_enabled": obj.get("is_speed_test_enabled"),
            "is_ssl_pinning_enabled": obj.get("is_ssl_pinning_enabled"),
            "is_star_enabled": obj.get("is_star_enabled"),
            "latest_android_app_version": obj.get("latest_android_app_version"),
            "line_official_account_id": obj.get("line_official_account_id"),
            "localized_android_news_title": obj.get("localized_android_news_title"),
            "localized_android_news_url": obj.get("localized_android_news_url"),
            "localized_maintenance_url": obj.get("localized_maintenance_url"),
            "localized_news_title": obj.get("localized_news_title"),
            "localized_news_url": obj.get("localized_news_url"),
            "maintenance_features": obj.get("maintenance_features"),
            "max_image_frame_count": obj.get("max_image_frame_count"),
            "minimum_android_app_version_required": obj.get("minimum_android_app_version_required"),
            "minimum_android_version_supported": obj.get("minimum_android_version_supported"),
            "news_version": obj.get("news_version"),
            "nft_infos": obj.get("nft_infos"),
            "promotion_sticker_pack_ids": obj.get("promotion_sticker_pack_ids"),
            "should_append_user_id_to_news_url": obj.get("should_append_user_id_to_news_url"),
            "support_email_address": obj.get("support_email_address"),
            "token_infos": obj.get("token_infos"),
            "use_random_message_refresh_rate": obj.get("use_random_message_refresh_rate"),
            "web3_is_maintenanced": obj.get("web3_is_maintenanced"),
            "web3_localized_maintenance_url": obj.get("web3_localized_maintenance_url"),
            "web3_maintenance_features": obj.get("web3_maintenance_features"),
            "web3_maintenance_url": obj.get("web3_maintenance_url"),
            "x_yay_global_id": obj.get("x_yay_global_id"),
            "x_yay_jp_id": obj.get("x_yay_jp_id"),
            "x_yay_update_id": obj.get("x_yay_update_id")
        })
        return _obj


