// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct CreateGroupQuota {
    #[serde(rename = "remaining_quota", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub remaining_quota: Option<Option<i32>>,
    #[serde(rename = "used_quota", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub used_quota: Option<Option<i32>>,
}

impl CreateGroupQuota {
    pub fn new() -> CreateGroupQuota {
        CreateGroupQuota {
            remaining_quota: None,
            used_quota: None,
        }
    }
}

