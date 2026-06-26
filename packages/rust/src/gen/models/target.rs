// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct Target {
    #[serde(rename = "followed_by", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub followed_by: Option<Option<bool>>,
}

impl Target {
    pub fn new() -> Target {
        Target {
            followed_by: None,
        }
    }
}

