// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct OnChainTransactionDto {
    #[serde(rename = "amount", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub amount: Option<Option<String>>,
    #[serde(rename = "asset_info", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub asset_info: Option<Option<models::AssetInfoDto>>,
    #[serde(rename = "created_at", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub created_at: Option<Option<i64>>,
    #[serde(rename = "currency", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub currency: Option<Option<String>>,
    #[serde(rename = "description", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub description: Option<Option<String>>,
    #[serde(rename = "from_address", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub from_address: Option<Option<String>>,
    #[serde(rename = "hash", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub hash: Option<Option<String>>,
    #[serde(rename = "status", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub status: Option<Option<String>>,
    #[serde(rename = "to_address", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub to_address: Option<Option<String>>,
    #[serde(rename = "token_contract", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub token_contract: Option<Option<String>>,
    #[serde(rename = "token_type", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub token_type: Option<Option<String>>,
}

impl OnChainTransactionDto {
    pub fn new() -> OnChainTransactionDto {
        OnChainTransactionDto {
            amount: None,
            asset_info: None,
            created_at: None,
            currency: None,
            description: None,
            from_address: None,
            hash: None,
            status: None,
            to_address: None,
            token_contract: None,
            token_type: None,
        }
    }
}

