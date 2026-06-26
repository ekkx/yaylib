// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct VerifyDeviceResponse {
    #[serde(rename = "verified", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub verified: Option<Option<bool>>,
    #[serde(rename = "verified_at", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub verified_at: Option<Option<String>>,
}

impl VerifyDeviceResponse {
    pub fn new() -> VerifyDeviceResponse {
        VerifyDeviceResponse {
            verified: None,
            verified_at: None,
        }
    }
}

