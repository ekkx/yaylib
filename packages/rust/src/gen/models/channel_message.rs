// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct ChannelMessage {
    #[serde(rename = "identifier", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub identifier: Option<Option<String>>,
    #[serde(rename = "message", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub message: Option<Option<models::EventMessage>>,
    #[serde(rename = "type", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub r#type: Option<Option<String>>,
}

impl ChannelMessage {
    pub fn new() -> ChannelMessage {
        ChannelMessage {
            identifier: None,
            message: None,
            r#type: None,
        }
    }
}

