// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct FriendShipsResponse {
    #[serde(rename = "relationship", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub relationship: Option<Option<models::Relationship>>,
}

impl FriendShipsResponse {
    pub fn new() -> FriendShipsResponse {
        FriendShipsResponse {
            relationship: None,
        }
    }
}

