// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct BlockchainDto {
    #[serde(rename = "addition_gas_percent", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub addition_gas_percent: Option<Option<models::AdditionGasPercentDto>>,
    #[serde(rename = "block_explorer", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub block_explorer: Option<Option<String>>,
    #[serde(rename = "chain_id", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub chain_id: Option<Option<i64>>,
    #[serde(rename = "chain_name", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub chain_name: Option<Option<String>>,
    #[serde(rename = "image_url", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub image_url: Option<Option<String>>,
    #[serde(rename = "nft_collections", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub nft_collections: Option<Option<Vec<models::NftCollectionDto>>>,
    #[serde(rename = "priority", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub priority: Option<Option<i32>>,
    #[serde(rename = "rpc_url", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub rpc_url: Option<Option<String>>,
    #[serde(rename = "tokens", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub tokens: Option<Option<Vec<models::TokenDto>>>,
}

impl BlockchainDto {
    pub fn new() -> BlockchainDto {
        BlockchainDto {
            addition_gas_percent: None,
            block_explorer: None,
            chain_id: None,
            chain_name: None,
            image_url: None,
            nft_collections: None,
            priority: None,
            rpc_url: None,
            tokens: None,
        }
    }
}

