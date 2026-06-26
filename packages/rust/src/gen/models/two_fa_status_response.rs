// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct TwoFaStatusResponse {
    #[serde(rename = "two_fa_auth_required", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub two_fa_auth_required: Option<Option<models::TwoFaAuthRequiredDto>>,
    #[serde(rename = "two_fa_enabled", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub two_fa_enabled: Option<Option<bool>>,
}

impl TwoFaStatusResponse {
    pub fn new() -> TwoFaStatusResponse {
        TwoFaStatusResponse {
            two_fa_auth_required: None,
            two_fa_enabled: None,
        }
    }
}

