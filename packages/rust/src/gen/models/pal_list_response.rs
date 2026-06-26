// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct PalListResponse {
    #[serde(rename = "cursor", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub cursor: Option<Option<String>>,
    #[serde(rename = "result", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub result: Option<Option<Vec<models::PalDto>>>,
    #[serde(rename = "total_count", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub total_count: Option<Option<i32>>,
}

impl PalListResponse {
    pub fn new() -> PalListResponse {
        PalListResponse {
            cursor: None,
            result: None,
            total_count: None,
        }
    }
}

