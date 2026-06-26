// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct InvitationCodeResponse {
    #[serde(rename = "invitation_code", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub invitation_code: Option<Option<models::CodeResponse>>,
}

impl InvitationCodeResponse {
    pub fn new() -> InvitationCodeResponse {
        InvitationCodeResponse {
            invitation_code: None,
        }
    }
}

