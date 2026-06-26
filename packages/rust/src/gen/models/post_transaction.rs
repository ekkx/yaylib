// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct PostTransaction {
    #[serde(rename = "receipt_json_str", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub receipt_json_str: Option<Option<String>>,
}

impl PostTransaction {
    pub fn new() -> PostTransaction {
        PostTransaction {
            receipt_json_str: None,
        }
    }
}

