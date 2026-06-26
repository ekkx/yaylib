// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct CommonIdsRequest {
    #[serde(rename = "ids", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub ids: Option<Option<Vec<i64>>>,
}

impl CommonIdsRequest {
    pub fn new() -> CommonIdsRequest {
        CommonIdsRequest {
            ids: None,
        }
    }
}

