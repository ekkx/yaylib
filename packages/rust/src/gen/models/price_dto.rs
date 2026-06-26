// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct PriceDto {
    #[serde(rename = "jpy", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub jpy: Option<Option<String>>,
    #[serde(rename = "usd", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub usd: Option<Option<String>>,
}

impl PriceDto {
    pub fn new() -> PriceDto {
        PriceDto {
            jpy: None,
            usd: None,
        }
    }
}

