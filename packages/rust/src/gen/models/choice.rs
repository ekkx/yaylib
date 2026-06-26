// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct Choice {
    #[serde(rename = "id", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub id: Option<Option<i64>>,
    #[serde(rename = "label", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub label: Option<Option<String>>,
    #[serde(rename = "votes_count", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub votes_count: Option<Option<i32>>,
}

impl Choice {
    pub fn new() -> Choice {
        Choice {
            id: None,
            label: None,
            votes_count: None,
        }
    }
}

