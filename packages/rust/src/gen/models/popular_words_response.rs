// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct PopularWordsResponse {
    #[serde(rename = "popular_words", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub popular_words: Option<Option<Vec<models::PopularWord>>>,
}

impl PopularWordsResponse {
    pub fn new() -> PopularWordsResponse {
        PopularWordsResponse {
            popular_words: None,
        }
    }
}

