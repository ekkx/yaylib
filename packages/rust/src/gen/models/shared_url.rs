// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct SharedUrl {
    #[serde(rename = "description", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub description: Option<Option<String>>,
    #[serde(rename = "image_url", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub image_url: Option<Option<String>>,
    #[serde(rename = "title", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub title: Option<Option<String>>,
    #[serde(rename = "url", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub url: Option<Option<String>>,
}

impl SharedUrl {
    pub fn new() -> SharedUrl {
        SharedUrl {
            description: None,
            image_url: None,
            title: None,
            url: None,
        }
    }
}

