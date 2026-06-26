// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct FootprintsResponse {
    #[serde(rename = "footprints", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub footprints: Option<Option<Vec<models::FootprintDto>>>,
    #[serde(rename = "next_page_value", default, skip_serializing_if = "Option::is_none", deserialize_with = "crate::gen::cursor::deserialize_option_string_or_number")]
    pub next_page_value: Option<String>,
}

impl FootprintsResponse {
    pub fn new() -> FootprintsResponse {
        FootprintsResponse {
            footprints: None,
            next_page_value: None,
        }
    }
}

