// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct MuteKeywordRequest {
    #[serde(rename = "context", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub context: Option<Option<Vec<String>>>,
    #[serde(rename = "word", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub word: Option<Option<String>>,
}

impl MuteKeywordRequest {
    pub fn new() -> MuteKeywordRequest {
        MuteKeywordRequest {
            context: None,
            word: None,
        }
    }
}

