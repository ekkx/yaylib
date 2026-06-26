// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct BlockedUserIdsResponse {
    #[serde(rename = "block_ids", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub block_ids: Option<Option<Vec<i64>>>,
}

impl BlockedUserIdsResponse {
    pub fn new() -> BlockedUserIdsResponse {
        BlockedUserIdsResponse {
            block_ids: None,
        }
    }
}

