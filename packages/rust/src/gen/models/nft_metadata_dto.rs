// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct NftMetadataDto {
    #[serde(rename = "attributes", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub attributes: Option<Option<Vec<models::NftAttributeDto>>>,
}

impl NftMetadataDto {
    pub fn new() -> NftMetadataDto {
        NftMetadataDto {
            attributes: None,
        }
    }
}

