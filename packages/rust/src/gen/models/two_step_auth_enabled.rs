// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct TwoStepAuthEnabled {
    #[serde(rename = "recovery_code_list", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub recovery_code_list: Option<Option<Vec<String>>>,
}

impl TwoStepAuthEnabled {
    pub fn new() -> TwoStepAuthEnabled {
        TwoStepAuthEnabled {
            recovery_code_list: None,
        }
    }
}

