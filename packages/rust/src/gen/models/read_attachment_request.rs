// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct ReadAttachmentRequest {
    #[serde(rename = "attachment_msg_ids", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub attachment_msg_ids: Option<Option<Vec<i64>>>,
}

impl ReadAttachmentRequest {
    pub fn new() -> ReadAttachmentRequest {
        ReadAttachmentRequest {
            attachment_msg_ids: None,
        }
    }
}

