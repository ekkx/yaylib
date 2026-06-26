// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct UsersByTimestampResponse {
    #[serde(rename = "last_timestamp", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub last_timestamp: Option<Option<i64>>,
    #[serde(rename = "users", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub users: Option<Option<Vec<models::RealmUser>>>,
}

impl UsersByTimestampResponse {
    pub fn new() -> UsersByTimestampResponse {
        UsersByTimestampResponse {
            last_timestamp: None,
            users: None,
        }
    }
}

