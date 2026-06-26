// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct InvitedUserDto {
    #[serde(rename = "referred", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub referred: Option<Option<models::InvitedUserDetailsDto>>,
    #[serde(rename = "total_earned_points", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub total_earned_points: Option<Option<i32>>,
}

impl InvitedUserDto {
    pub fn new() -> InvitedUserDto {
        InvitedUserDto {
            referred: None,
            total_earned_points: None,
        }
    }
}

