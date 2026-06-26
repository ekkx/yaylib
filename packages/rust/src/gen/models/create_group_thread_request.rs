// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct CreateGroupThreadRequest {
    #[serde(rename = "group_id", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub group_id: Option<Option<i64>>,
    #[serde(rename = "thread_icon_filename", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub thread_icon_filename: Option<Option<String>>,
    #[serde(rename = "title", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub title: Option<Option<String>>,
}

impl CreateGroupThreadRequest {
    pub fn new() -> CreateGroupThreadRequest {
        CreateGroupThreadRequest {
            group_id: None,
            thread_icon_filename: None,
            title: None,
        }
    }
}

