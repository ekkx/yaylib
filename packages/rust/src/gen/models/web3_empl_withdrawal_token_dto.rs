// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct Web3EmplWithdrawalTokenDto {
    #[serde(rename = "chain_id", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub chain_id: Option<Option<String>>,
    #[serde(rename = "id", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub id: Option<Option<String>>,
    #[serde(rename = "image_url", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub image_url: Option<Option<String>>,
    #[serde(rename = "token_address", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub token_address: Option<Option<String>>,
    #[serde(rename = "token_name", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub token_name: Option<Option<String>>,
    #[serde(rename = "token_symbol", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub token_symbol: Option<Option<String>>,
    #[serde(rename = "velo_pool_version", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub velo_pool_version: Option<Option<String>>,
}

impl Web3EmplWithdrawalTokenDto {
    pub fn new() -> Web3EmplWithdrawalTokenDto {
        Web3EmplWithdrawalTokenDto {
            chain_id: None,
            id: None,
            image_url: None,
            token_address: None,
            token_name: None,
            token_symbol: None,
            velo_pool_version: None,
        }
    }
}

