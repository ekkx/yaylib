// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct Web3WalletTransactionHistory {
    #[serde(rename = "amount", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub amount: Option<Option<String>>,
    #[serde(rename = "asset_info", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub asset_info: Option<Option<models::AssetInfo>>,
    #[serde(rename = "created_at", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub created_at: Option<Option<i64>>,
    #[serde(rename = "currency", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub currency: Option<Option<String>>,
    #[serde(rename = "description", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub description: Option<Option<String>>,
    #[serde(rename = "details", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub details: Option<Option<models::EmplTokenExchangeDetails>>,
    #[serde(rename = "final_status_updated_at", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub final_status_updated_at: Option<Option<i64>>,
    #[serde(rename = "from_address", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub from_address: Option<Option<String>>,
    #[serde(rename = "gift_transaction_uuid", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub gift_transaction_uuid: Option<Option<String>>,
    #[serde(rename = "hash", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub hash: Option<Option<String>>,
    #[serde(rename = "id", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub id: Option<Option<String>>,
    #[serde(rename = "status", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub status: Option<Option<String>>,
    #[serde(rename = "to_address", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub to_address: Option<Option<String>>,
    #[serde(rename = "token_contract", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub token_contract: Option<Option<String>>,
    #[serde(rename = "token_id", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub token_id: Option<Option<String>>,
    #[serde(rename = "transaction_code", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub transaction_code: Option<Option<String>>,
    #[serde(rename = "type", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub r#type: Option<Option<serde_json::Value>>,
    #[serde(rename = "updated_at", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub updated_at: Option<Option<i64>>,
}

impl Web3WalletTransactionHistory {
    pub fn new() -> Web3WalletTransactionHistory {
        Web3WalletTransactionHistory {
            amount: None,
            asset_info: None,
            created_at: None,
            currency: None,
            description: None,
            details: None,
            final_status_updated_at: None,
            from_address: None,
            gift_transaction_uuid: None,
            hash: None,
            id: None,
            status: None,
            to_address: None,
            token_contract: None,
            token_id: None,
            transaction_code: None,
            r#type: None,
            updated_at: None,
        }
    }
}

