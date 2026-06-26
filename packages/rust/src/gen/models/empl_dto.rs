// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct EmplDto {
    #[serde(rename = "campaign", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub campaign: Option<Option<String>>,
    #[serde(rename = "purchased", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub purchased: Option<Option<String>>,
    #[serde(rename = "regular", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub regular: Option<Option<String>>,
    #[serde(rename = "rewarded", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub rewarded: Option<Option<String>>,
    #[serde(rename = "total", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub total: Option<Option<String>>,
}

impl EmplDto {
    pub fn new() -> EmplDto {
        EmplDto {
            campaign: None,
            purchased: None,
            regular: None,
            rewarded: None,
            total: None,
        }
    }
}

