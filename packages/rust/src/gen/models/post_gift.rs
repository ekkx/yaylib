// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct PostGift {
    #[serde(rename = "count", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub count: Option<Option<i32>>,
    #[serde(rename = "gift", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub gift: Option<Option<models::Gift>>,
}

impl PostGift {
    pub fn new() -> PostGift {
        PostGift {
            count: None,
            gift: None,
        }
    }
}

