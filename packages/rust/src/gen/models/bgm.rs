// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct Bgm {
    #[serde(rename = "id", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub id: Option<Option<i64>>,
    #[serde(rename = "music_url", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub music_url: Option<Option<String>>,
    #[serde(rename = "order", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub order: Option<Option<i32>>,
    #[serde(rename = "title", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub title: Option<Option<String>>,
}

impl Bgm {
    pub fn new() -> Bgm {
        Bgm {
            id: None,
            music_url: None,
            order: None,
            title: None,
        }
    }
}

