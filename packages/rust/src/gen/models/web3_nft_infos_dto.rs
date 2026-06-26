// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct Web3NftInfosDto {
    #[serde(rename = "nft_infos", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub nft_infos: Option<Option<Vec<models::Web3NftInfoDto>>>,
}

impl Web3NftInfosDto {
    pub fn new() -> Web3NftInfosDto {
        Web3NftInfosDto {
            nft_infos: None,
        }
    }
}

