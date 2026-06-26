// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct Web3NftCollectionItemsResponse {
    #[serde(rename = "data", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub data: Option<Option<Vec<models::Web3NftCollectionItemDto>>>,
}

impl Web3NftCollectionItemsResponse {
    pub fn new() -> Web3NftCollectionItemsResponse {
        Web3NftCollectionItemsResponse {
            data: None,
        }
    }
}

