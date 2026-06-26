// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct BanWordsResponse {
    #[serde(rename = "ban_words", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub ban_words: Option<Option<Vec<models::BanWord>>>,
}

impl BanWordsResponse {
    pub fn new() -> BanWordsResponse {
        BanWordsResponse {
            ban_words: None,
        }
    }
}

