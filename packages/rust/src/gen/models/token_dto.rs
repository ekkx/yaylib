// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct TokenDto {
    #[serde(rename = "balance", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub balance: Option<Option<String>>,
    #[serde(rename = "bridge", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub bridge: Option<Option<models::BridgeDto>>,
    #[serde(rename = "contract_address", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub contract_address: Option<Option<String>>,
    #[serde(rename = "decimals", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub decimals: Option<Option<i32>>,
    #[serde(rename = "image_url", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub image_url: Option<Option<String>>,
    #[serde(rename = "price", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub price: Option<Option<models::PriceDto>>,
    #[serde(rename = "priority", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub priority: Option<Option<i32>>,
    #[serde(rename = "status", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub status: Option<Option<String>>,
    #[serde(rename = "swap", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub swap: Option<Option<models::SwapDto>>,
    #[serde(rename = "symbol", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub symbol: Option<Option<String>>,
    #[serde(rename = "token_type", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub token_type: Option<Option<String>>,
}

impl TokenDto {
    pub fn new() -> TokenDto {
        TokenDto {
            balance: None,
            bridge: None,
            contract_address: None,
            decimals: None,
            image_url: None,
            price: None,
            priority: None,
            status: None,
            swap: None,
            symbol: None,
            token_type: None,
        }
    }
}

