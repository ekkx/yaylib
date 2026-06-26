// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct Web3NftInfoDto {
    #[serde(rename = "chain_id", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub chain_id: Option<Option<i64>>,
    #[serde(rename = "contract_address", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub contract_address: Option<Option<String>>,
    #[serde(rename = "localized_urls", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub localized_urls: Option<Option<models::LocalizedStringDto>>,
    #[serde(rename = "name", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub name: Option<Option<String>>,
}

impl Web3NftInfoDto {
    pub fn new() -> Web3NftInfoDto {
        Web3NftInfoDto {
            chain_id: None,
            contract_address: None,
            localized_urls: None,
            name: None,
        }
    }
}

