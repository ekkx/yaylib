// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct ConferenceCallBumpParams {
    #[serde(rename = "participant_limit", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub participant_limit: Option<Option<i32>>,
}

impl ConferenceCallBumpParams {
    pub fn new() -> ConferenceCallBumpParams {
        ConferenceCallBumpParams {
            participant_limit: None,
        }
    }
}

