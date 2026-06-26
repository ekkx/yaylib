// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct MuteKeywordResponse {
    #[serde(rename = "hidden_words", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub hidden_words: Option<Option<Vec<models::MuteKeyword>>>,
}

impl MuteKeywordResponse {
    pub fn new() -> MuteKeywordResponse {
        MuteKeywordResponse {
            hidden_words: None,
        }
    }
}

