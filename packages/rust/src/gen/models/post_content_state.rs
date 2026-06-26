// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct PostContentState {
    #[serde(rename = "place_holder_id", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub place_holder_id: Option<Option<i32>>,
}

impl PostContentState {
    pub fn new() -> PostContentState {
        PostContentState {
            place_holder_id: None,
        }
    }
}

