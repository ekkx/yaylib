// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct EmplActivityDto {
    #[serde(rename = "amount", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub amount: Option<Option<String>>,
    #[serde(rename = "avatar_frame_user_uuid", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub avatar_frame_user_uuid: Option<Option<String>>,
    #[serde(rename = "battle_history_id", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub battle_history_id: Option<Option<i64>>,
    #[serde(rename = "chain_id", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub chain_id: Option<Option<i64>>,
    #[serde(rename = "details", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub details: Option<Option<models::PalDetailsDto>>,
    #[serde(rename = "empl", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub empl: Option<Option<models::EmplDto>>,
    #[serde(rename = "gift_exchange_uuid", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub gift_exchange_uuid: Option<Option<String>>,
    #[serde(rename = "gift_transaction_uuid", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub gift_transaction_uuid: Option<Option<String>>,
    #[serde(rename = "id", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub id: Option<Option<i64>>,
    #[serde(rename = "money_amount", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub money_amount: Option<Option<f64>>,
    #[serde(rename = "occurred_at", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub occurred_at: Option<Option<i64>>,
    #[serde(rename = "price", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub price: Option<Option<f64>>,
    #[serde(rename = "race_id", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub race_id: Option<Option<String>>,
    #[serde(rename = "status", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub status: Option<Option<String>>,
    #[serde(rename = "token_address", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub token_address: Option<Option<String>>,
    #[serde(rename = "token_symbol", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub token_symbol: Option<Option<String>>,
    #[serde(rename = "transaction_code", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub transaction_code: Option<Option<String>>,
    #[serde(rename = "type", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub r#type: Option<Option<String>>,
}

impl EmplActivityDto {
    pub fn new() -> EmplActivityDto {
        EmplActivityDto {
            amount: None,
            avatar_frame_user_uuid: None,
            battle_history_id: None,
            chain_id: None,
            details: None,
            empl: None,
            gift_exchange_uuid: None,
            gift_transaction_uuid: None,
            id: None,
            money_amount: None,
            occurred_at: None,
            price: None,
            race_id: None,
            status: None,
            token_address: None,
            token_symbol: None,
            transaction_code: None,
            r#type: None,
        }
    }
}

