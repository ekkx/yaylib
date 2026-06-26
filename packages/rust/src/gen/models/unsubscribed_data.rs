// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct UnsubscribedData {
    #[serde(rename = "user_ids", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub user_ids: Option<Option<Vec<i64>>>,
}

impl UnsubscribedData {
    pub fn new() -> UnsubscribedData {
        UnsubscribedData {
            user_ids: None,
        }
    }
}

