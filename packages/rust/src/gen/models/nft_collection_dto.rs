// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct NftCollectionDto {
    #[serde(rename = "contract_address", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub contract_address: Option<Option<String>>,
    #[serde(rename = "image", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub image: Option<Option<String>>,
    #[serde(rename = "media_type", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub media_type: Option<Option<String>>,
    #[serde(rename = "name", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub name: Option<Option<String>>,
    #[serde(rename = "priority", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub priority: Option<Option<i32>>,
    #[serde(rename = "status", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub status: Option<Option<String>>,
}

impl NftCollectionDto {
    pub fn new() -> NftCollectionDto {
        NftCollectionDto {
            contract_address: None,
            image: None,
            media_type: None,
            name: None,
            priority: None,
            status: None,
        }
    }
}

