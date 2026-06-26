// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct InvitedUserDetailsDto {
    #[serde(rename = "id", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub id: Option<Option<i64>>,
    #[serde(rename = "nickname", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub nickname: Option<Option<String>>,
    #[serde(rename = "profile_icon", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub profile_icon: Option<Option<String>>,
}

impl InvitedUserDetailsDto {
    pub fn new() -> InvitedUserDetailsDto {
        InvitedUserDetailsDto {
            id: None,
            nickname: None,
            profile_icon: None,
        }
    }
}

