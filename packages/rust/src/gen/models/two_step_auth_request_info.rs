// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct TwoStepAuthRequestInfo {
    #[serde(rename = "qr_code_url", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub qr_code_url: Option<Option<String>>,
    #[serde(rename = "setup_key", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub setup_key: Option<Option<String>>,
}

impl TwoStepAuthRequestInfo {
    pub fn new() -> TwoStepAuthRequestInfo {
        TwoStepAuthRequestInfo {
            qr_code_url: None,
            setup_key: None,
        }
    }
}

