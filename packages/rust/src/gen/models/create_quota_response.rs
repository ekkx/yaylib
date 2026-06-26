// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct CreateQuotaResponse {
    #[serde(rename = "create", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub create: Option<Option<models::CreateGroupQuota>>,
}

impl CreateQuotaResponse {
    pub fn new() -> CreateQuotaResponse {
        CreateQuotaResponse {
            create: None,
        }
    }
}

