// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct Web3NftCollectionItemDto {
    #[serde(rename = "balance", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub balance: Option<Option<String>>,
    #[serde(rename = "chain_id", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub chain_id: Option<Option<i64>>,
    #[serde(rename = "contract_address", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub contract_address: Option<Option<String>>,
    #[serde(rename = "extra_metadata", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub extra_metadata: Option<Option<models::NftMetadataDto>>,
    #[serde(rename = "image_url", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub image_url: Option<Option<String>>,
    #[serde(rename = "name", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub name: Option<Option<String>>,
    #[serde(rename = "token_id", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub token_id: Option<Option<String>>,
    #[serde(rename = "token_type", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub token_type: Option<Option<String>>,
}

impl Web3NftCollectionItemDto {
    pub fn new() -> Web3NftCollectionItemDto {
        Web3NftCollectionItemDto {
            balance: None,
            chain_id: None,
            contract_address: None,
            extra_metadata: None,
            image_url: None,
            name: None,
            token_id: None,
            token_type: None,
        }
    }
}

