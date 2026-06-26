// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct Hatch {
    #[serde(rename = "pal_image_url", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub pal_image_url: Option<Option<String>>,
    #[serde(rename = "pal_name", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub pal_name: Option<Option<String>>,
}

impl Hatch {
    pub fn new() -> Hatch {
        Hatch {
            pal_image_url: None,
            pal_name: None,
        }
    }
}

