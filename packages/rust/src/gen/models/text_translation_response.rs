// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct TextTranslationResponse {
    #[serde(rename = "translated_text", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub translated_text: Option<Option<String>>,
}

impl TextTranslationResponse {
    pub fn new() -> TextTranslationResponse {
        TextTranslationResponse {
            translated_text: None,
        }
    }
}

