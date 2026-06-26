// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct RtmTokenResponse {
    #[serde(rename = "token", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub token: Option<Option<String>>,
}

impl RtmTokenResponse {
    pub fn new() -> RtmTokenResponse {
        RtmTokenResponse {
            token: None,
        }
    }
}

