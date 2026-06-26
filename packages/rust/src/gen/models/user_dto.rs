// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct UserDto {
    #[serde(rename = "id", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub id: Option<Option<i64>>,
    #[serde(rename = "nickname", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub nickname: Option<Option<String>>,
    #[serde(rename = "profile_icon", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub profile_icon: Option<Option<String>>,
}

impl UserDto {
    pub fn new() -> UserDto {
        UserDto {
            id: None,
            nickname: None,
            profile_icon: None,
        }
    }
}

