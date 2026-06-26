// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct MessageResponse {
    #[serde(rename = "conference_call", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub conference_call: Option<Option<models::RealmConferenceCall>>,
    #[serde(rename = "id", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub id: Option<Option<i64>>,
}

impl MessageResponse {
    pub fn new() -> MessageResponse {
        MessageResponse {
            conference_call: None,
            id: None,
        }
    }
}

