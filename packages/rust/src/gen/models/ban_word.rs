// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct BanWord {
    #[serde(rename = "id", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub id: Option<Option<i64>>,
    #[serde(rename = "type", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub r#type: Option<Option<String>>,
    #[serde(rename = "word", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub word: Option<Option<String>>,
}

impl BanWord {
    pub fn new() -> BanWord {
        BanWord {
            id: None,
            r#type: None,
            word: None,
        }
    }
}

