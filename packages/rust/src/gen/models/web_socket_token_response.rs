// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct WebSocketTokenResponse {
    #[serde(rename = "token", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub token: Option<Option<String>>,
}

impl WebSocketTokenResponse {
    pub fn new() -> WebSocketTokenResponse {
        WebSocketTokenResponse {
            token: None,
        }
    }
}

