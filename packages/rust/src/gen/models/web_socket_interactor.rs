// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct WebSocketInteractor {
    #[serde(rename = "d", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub d: Option<Option<serde_json::Value>>,
    #[serde(rename = "f", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub f: Option<Option<serde_json::Value>>,
    #[serde(rename = "h", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub h: Option<Option<serde_json::Value>>,
    #[serde(rename = "i", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub i: Option<Option<serde_json::Value>>,
}

impl WebSocketInteractor {
    pub fn new() -> WebSocketInteractor {
        WebSocketInteractor {
            d: None,
            f: None,
            h: None,
            i: None,
        }
    }
}

