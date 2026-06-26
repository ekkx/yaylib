// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct PalDto {
    #[serde(rename = "age_in_days", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub age_in_days: Option<Option<i32>>,
    #[serde(rename = "attributes", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub attributes: Option<Option<models::Attribute>>,
    #[serde(rename = "current_empl_earned", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub current_empl_earned: Option<Option<String>>,
    #[serde(rename = "current_level", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub current_level: Option<Option<i32>>,
    #[serde(rename = "description", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub description: Option<Option<String>>,
    #[serde(rename = "empl_earned_limit", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub empl_earned_limit: Option<Option<String>>,
    #[serde(rename = "from_liquidity_pool", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub from_liquidity_pool: Option<Option<bool>>,
    #[serde(rename = "hatching_animation_url", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub hatching_animation_url: Option<Option<String>>,
    #[serde(rename = "image", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub image: Option<Option<String>>,
    #[serde(rename = "in_locked_party", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub in_locked_party: Option<Option<bool>>,
    #[serde(rename = "is_alive", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub is_alive: Option<Option<bool>>,
    #[serde(rename = "is_pending", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub is_pending: Option<Option<bool>>,
    #[serde(rename = "is_valid_to_evolve", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub is_valid_to_evolve: Option<Option<bool>>,
    #[serde(rename = "is_valid_to_level_up", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub is_valid_to_level_up: Option<Option<bool>>,
    #[serde(rename = "max_age_in_days", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub max_age_in_days: Option<Option<i32>>,
    #[serde(rename = "max_attributes", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub max_attributes: Option<Option<models::MaxAttribute>>,
    #[serde(rename = "name", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub name: Option<Option<String>>,
    #[serde(rename = "original_wallet_address", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub original_wallet_address: Option<Option<String>>,
    #[serde(rename = "race_status", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub race_status: Option<Option<String>>,
    #[serde(rename = "token_id", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub token_id: Option<Option<String>>,
}

impl PalDto {
    pub fn new() -> PalDto {
        PalDto {
            age_in_days: None,
            attributes: None,
            current_empl_earned: None,
            current_level: None,
            description: None,
            empl_earned_limit: None,
            from_liquidity_pool: None,
            hatching_animation_url: None,
            image: None,
            in_locked_party: None,
            is_alive: None,
            is_pending: None,
            is_valid_to_evolve: None,
            is_valid_to_level_up: None,
            max_age_in_days: None,
            max_attributes: None,
            name: None,
            original_wallet_address: None,
            race_status: None,
            token_id: None,
        }
    }
}

