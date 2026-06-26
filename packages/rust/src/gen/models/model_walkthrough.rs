// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct ModelWalkthrough {
    #[serde(rename = "article_id", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub article_id: Option<Option<i64>>,
    #[serde(rename = "title", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub title: Option<Option<String>>,
    #[serde(rename = "url", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub url: Option<Option<String>>,
}

impl ModelWalkthrough {
    pub fn new() -> ModelWalkthrough {
        ModelWalkthrough {
            article_id: None,
            title: None,
            url: None,
        }
    }
}

