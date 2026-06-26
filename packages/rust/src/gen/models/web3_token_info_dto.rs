// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct Web3TokenInfoDto {
    #[serde(rename = "chain_id", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub chain_id: Option<Option<i64>>,
    #[serde(rename = "localized_urls", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub localized_urls: Option<Option<models::LocalizedStringDto>>,
    #[serde(rename = "symbol", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub symbol: Option<Option<String>>,
}

impl Web3TokenInfoDto {
    pub fn new() -> Web3TokenInfoDto {
        Web3TokenInfoDto {
            chain_id: None,
            localized_urls: None,
            symbol: None,
        }
    }
}

