// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct FootprintDto {
    #[serde(rename = "id", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub id: Option<Option<i64>>,
    #[serde(rename = "user", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub user: Option<Option<models::UserUserDto>>,
    #[serde(rename = "visited_at", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub visited_at: Option<Option<i64>>,
}

impl FootprintDto {
    pub fn new() -> FootprintDto {
        FootprintDto {
            id: None,
            user: None,
            visited_at: None,
        }
    }
}

