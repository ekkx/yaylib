// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct CreateMuteKeywordResponse {
    #[serde(rename = "hidden_word", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub hidden_word: Option<Option<models::MuteKeyword>>,
}

impl CreateMuteKeywordResponse {
    pub fn new() -> CreateMuteKeywordResponse {
        CreateMuteKeywordResponse {
            hidden_word: None,
        }
    }
}

