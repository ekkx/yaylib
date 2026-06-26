// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct FeaturesPalDto {
    #[serde(rename = "chain_id", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub chain_id: Option<Option<i64>>,
    #[serde(rename = "contract_address", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub contract_address: Option<Option<String>>,
    #[serde(rename = "depositwithdraw_address", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub depositwithdraw_address: Option<Option<String>>,
}

impl FeaturesPalDto {
    pub fn new() -> FeaturesPalDto {
        FeaturesPalDto {
            chain_id: None,
            contract_address: None,
            depositwithdraw_address: None,
        }
    }
}

