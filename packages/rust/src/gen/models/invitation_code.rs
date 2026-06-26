// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct InvitationCode {
    #[serde(rename = "possible_to_edit_at_millis", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub possible_to_edit_at_millis: Option<Option<i64>>,
    #[serde(rename = "value", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub value: Option<Option<String>>,
}

impl InvitationCode {
    pub fn new() -> InvitationCode {
        InvitationCode {
            possible_to_edit_at_millis: None,
            value: None,
        }
    }
}

