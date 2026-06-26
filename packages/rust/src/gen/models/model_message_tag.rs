// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct ModelMessageTag {
    #[serde(rename = "end_at", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub end_at: Option<Option<i32>>,
    #[serde(rename = "length", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub length: Option<Option<i32>>,
    #[serde(rename = "offset", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub offset: Option<Option<i32>>,
    #[serde(rename = "type", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub r#type: Option<Option<models::MessageTagType>>,
    #[serde(rename = "user_id", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub user_id: Option<Option<i64>>,
}

impl ModelMessageTag {
    pub fn new() -> ModelMessageTag {
        ModelMessageTag {
            end_at: None,
            length: None,
            offset: None,
            r#type: None,
            user_id: None,
        }
    }
}

