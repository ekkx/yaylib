// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct GroupThreadListResponse {
    #[serde(rename = "next_page_value", default, skip_serializing_if = "Option::is_none", deserialize_with = "crate::gen::cursor::deserialize_option_string_or_number")]
    pub next_page_value: Option<String>,
    #[serde(rename = "threads", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub threads: Option<Option<Vec<models::ThreadInfo>>>,
}

impl GroupThreadListResponse {
    pub fn new() -> GroupThreadListResponse {
        GroupThreadListResponse {
            next_page_value: None,
            threads: None,
        }
    }
}

