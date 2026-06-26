// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct AssetInfoDto {
    #[serde(rename = "image", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub image: Option<Option<String>>,
    #[serde(rename = "name", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub name: Option<Option<String>>,
    #[serde(rename = "token_id", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub token_id: Option<Option<String>>,
}

impl AssetInfoDto {
    pub fn new() -> AssetInfoDto {
        AssetInfoDto {
            image: None,
            name: None,
            token_id: None,
        }
    }
}

