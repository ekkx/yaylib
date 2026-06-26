// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct TwoStepAuthEnabledResponse {
    #[serde(rename = "recovery_codes", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub recovery_codes: Option<Option<Vec<String>>>,
}

impl TwoStepAuthEnabledResponse {
    pub fn new() -> TwoStepAuthEnabledResponse {
        TwoStepAuthEnabledResponse {
            recovery_codes: None,
        }
    }
}

