// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct TwoStepAuthRequestInfoResponse {
    #[serde(rename = "qr_code", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub qr_code: Option<Option<String>>,
    #[serde(rename = "secret", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub secret: Option<Option<String>>,
}

impl TwoStepAuthRequestInfoResponse {
    pub fn new() -> TwoStepAuthRequestInfoResponse {
        TwoStepAuthRequestInfoResponse {
            qr_code: None,
            secret: None,
        }
    }
}

