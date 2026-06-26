// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct GifImageCategory {
    #[serde(rename = "gifs", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub gifs: Option<Option<Vec<models::GifImage>>>,
    #[serde(rename = "id", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub id: Option<Option<i64>>,
    #[serde(rename = "language", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub language: Option<Option<String>>,
    #[serde(rename = "name", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub name: Option<Option<String>>,
}

impl GifImageCategory {
    pub fn new() -> GifImageCategory {
        GifImageCategory {
            gifs: None,
            id: None,
            language: None,
            name: None,
        }
    }
}

