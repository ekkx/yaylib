// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct Config {
    #[serde(rename = "ad_tester_user_ids", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub ad_tester_user_ids: Option<Option<String>>,
    #[serde(rename = "android_news_version", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub android_news_version: Option<Option<String>>,
    #[serde(rename = "dapps", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub dapps: Option<Option<String>>,
    #[serde(rename = "is_chat_websocket_enabled", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub is_chat_websocket_enabled: Option<Option<String>>,
    #[serde(rename = "is_direct_vip_purchase_enabled", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub is_direct_vip_purchase_enabled: Option<Option<String>>,
    #[serde(rename = "is_gift_features_enabled", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub is_gift_features_enabled: Option<Option<String>>,
    #[serde(rename = "is_id_card_and_phone_verification_check_for_review_enabled", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub is_id_card_and_phone_verification_check_for_review_enabled: Option<Option<String>>,
    #[serde(rename = "is_id_card_check_required_blocker_view_enabled", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub is_id_card_check_required_blocker_view_enabled: Option<Option<String>>,
    #[serde(rename = "is_maintenanced", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub is_maintenanced: Option<Option<String>>,
    #[serde(rename = "is_phone_verification_required_blocker_view_enabled", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub is_phone_verification_required_blocker_view_enabled: Option<Option<String>>,
    #[serde(rename = "is_speed_test_enabled", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub is_speed_test_enabled: Option<Option<String>>,
    #[serde(rename = "is_ssl_pinning_enabled", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub is_ssl_pinning_enabled: Option<Option<String>>,
    #[serde(rename = "is_star_enabled", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub is_star_enabled: Option<Option<String>>,
    #[serde(rename = "latest_android_app_version", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub latest_android_app_version: Option<Option<String>>,
    #[serde(rename = "line_official_account_id", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub line_official_account_id: Option<Option<String>>,
    #[serde(rename = "localized_android_news_title", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub localized_android_news_title: Option<Option<String>>,
    #[serde(rename = "localized_android_news_url", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub localized_android_news_url: Option<Option<String>>,
    #[serde(rename = "localized_maintenance_url", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub localized_maintenance_url: Option<Option<String>>,
    #[serde(rename = "localized_news_title", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub localized_news_title: Option<Option<String>>,
    #[serde(rename = "localized_news_url", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub localized_news_url: Option<Option<String>>,
    #[serde(rename = "maintenance_features", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub maintenance_features: Option<Option<String>>,
    #[serde(rename = "max_image_frame_count", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub max_image_frame_count: Option<Option<String>>,
    #[serde(rename = "minimum_android_app_version_required", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub minimum_android_app_version_required: Option<Option<String>>,
    #[serde(rename = "minimum_android_version_supported", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub minimum_android_version_supported: Option<Option<String>>,
    #[serde(rename = "news_version", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub news_version: Option<Option<String>>,
    #[serde(rename = "nft_infos", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub nft_infos: Option<Option<String>>,
    #[serde(rename = "promotion_sticker_pack_ids", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub promotion_sticker_pack_ids: Option<Option<String>>,
    #[serde(rename = "should_append_user_id_to_news_url", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub should_append_user_id_to_news_url: Option<Option<String>>,
    #[serde(rename = "support_email_address", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub support_email_address: Option<Option<String>>,
    #[serde(rename = "token_infos", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub token_infos: Option<Option<String>>,
    #[serde(rename = "use_random_message_refresh_rate", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub use_random_message_refresh_rate: Option<Option<String>>,
    #[serde(rename = "web3_is_maintenanced", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub web3_is_maintenanced: Option<Option<String>>,
    #[serde(rename = "web3_localized_maintenance_url", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub web3_localized_maintenance_url: Option<Option<String>>,
    #[serde(rename = "web3_maintenance_features", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub web3_maintenance_features: Option<Option<String>>,
    #[serde(rename = "web3_maintenance_url", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub web3_maintenance_url: Option<Option<String>>,
    #[serde(rename = "x_yay_global_id", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub x_yay_global_id: Option<Option<String>>,
    #[serde(rename = "x_yay_jp_id", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub x_yay_jp_id: Option<Option<String>>,
    #[serde(rename = "x_yay_update_id", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub x_yay_update_id: Option<Option<String>>,
}

impl Config {
    pub fn new() -> Config {
        Config {
            ad_tester_user_ids: None,
            android_news_version: None,
            dapps: None,
            is_chat_websocket_enabled: None,
            is_direct_vip_purchase_enabled: None,
            is_gift_features_enabled: None,
            is_id_card_and_phone_verification_check_for_review_enabled: None,
            is_id_card_check_required_blocker_view_enabled: None,
            is_maintenanced: None,
            is_phone_verification_required_blocker_view_enabled: None,
            is_speed_test_enabled: None,
            is_ssl_pinning_enabled: None,
            is_star_enabled: None,
            latest_android_app_version: None,
            line_official_account_id: None,
            localized_android_news_title: None,
            localized_android_news_url: None,
            localized_maintenance_url: None,
            localized_news_title: None,
            localized_news_url: None,
            maintenance_features: None,
            max_image_frame_count: None,
            minimum_android_app_version_required: None,
            minimum_android_version_supported: None,
            news_version: None,
            nft_infos: None,
            promotion_sticker_pack_ids: None,
            should_append_user_id_to_news_url: None,
            support_email_address: None,
            token_infos: None,
            use_random_message_refresh_rate: None,
            web3_is_maintenanced: None,
            web3_localized_maintenance_url: None,
            web3_maintenance_features: None,
            web3_maintenance_url: None,
            x_yay_global_id: None,
            x_yay_jp_id: None,
            x_yay_update_id: None,
        }
    }
}

