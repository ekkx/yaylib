// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct ModelWeb3WalletTransactionHistory {
    #[serde(rename = "asset_info", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub asset_info: Option<Option<models::Web3WalletTransactionHistoryAssetInfo>>,
    #[serde(rename = "asset_type", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub asset_type: Option<Option<models::AssetType>>,
    #[serde(rename = "created_at_millis", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub created_at_millis: Option<Option<i64>>,
    #[serde(rename = "currency", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub currency: Option<Option<String>>,
    #[serde(rename = "dot_money_amount", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub dot_money_amount: Option<Option<f64>>,
    #[serde(rename = "empl_details", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub empl_details: Option<Option<models::EmplDetails>>,
    #[serde(rename = "from_address", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub from_address: Option<Option<String>>,
    #[serde(rename = "gift_exchange_uuid", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub gift_exchange_uuid: Option<Option<String>>,
    #[serde(rename = "gift_transaction_detail", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub gift_transaction_detail: Option<Option<models::GiftTransactionDetail>>,
    #[serde(rename = "gift_transaction_uuid", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub gift_transaction_uuid: Option<Option<String>>,
    #[serde(rename = "is_pending_transaction", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub is_pending_transaction: Option<Option<bool>>,
    #[serde(rename = "nft_token_id", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub nft_token_id: Option<Option<String>>,
    #[serde(rename = "pal_transaction_detail", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub pal_transaction_detail: Option<Option<serde_json::Value>>,
    #[serde(rename = "price", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub price: Option<Option<f64>>,
    #[serde(rename = "race_id", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub race_id: Option<Option<String>>,
    #[serde(rename = "race_ranks", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub race_ranks: Option<Option<Vec<i32>>>,
    #[serde(rename = "status", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub status: Option<Option<serde_json::Value>>,
    #[serde(rename = "to_address", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub to_address: Option<Option<String>>,
    #[serde(rename = "token_symbol", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub token_symbol: Option<Option<String>>,
    #[serde(rename = "transaction_code", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub transaction_code: Option<Option<String>>,
    #[serde(rename = "transaction_hash_hex", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub transaction_hash_hex: Option<Option<String>>,
    #[serde(rename = "transaction_history_type", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub transaction_history_type: Option<Option<serde_json::Value>>,
    #[serde(rename = "transaction_type", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub transaction_type: Option<Option<serde_json::Value>>,
    #[serde(rename = "value", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub value: Option<Option<f64>>,
    #[serde(rename = "value_symbol", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub value_symbol: Option<Option<String>>,
    #[serde(rename = "value_text", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub value_text: Option<Option<String>>,
}

impl ModelWeb3WalletTransactionHistory {
    pub fn new() -> ModelWeb3WalletTransactionHistory {
        ModelWeb3WalletTransactionHistory {
            asset_info: None,
            asset_type: None,
            created_at_millis: None,
            currency: None,
            dot_money_amount: None,
            empl_details: None,
            from_address: None,
            gift_exchange_uuid: None,
            gift_transaction_detail: None,
            gift_transaction_uuid: None,
            is_pending_transaction: None,
            nft_token_id: None,
            pal_transaction_detail: None,
            price: None,
            race_id: None,
            race_ranks: None,
            status: None,
            to_address: None,
            token_symbol: None,
            transaction_code: None,
            transaction_hash_hex: None,
            transaction_history_type: None,
            transaction_type: None,
            value: None,
            value_symbol: None,
            value_text: None,
        }
    }
}

