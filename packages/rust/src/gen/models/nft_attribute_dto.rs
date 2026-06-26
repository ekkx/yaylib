// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct NftAttributeDto {
    #[serde(rename = "trait_type", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub trait_type: Option<Option<String>>,
    #[serde(rename = "value", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub value: Option<Option<String>>,
}

impl NftAttributeDto {
    pub fn new() -> NftAttributeDto {
        NftAttributeDto {
            trait_type: None,
            value: None,
        }
    }
}

