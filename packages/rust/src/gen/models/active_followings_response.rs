// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct ActiveFollowingsResponse {
    #[serde(rename = "last_loggedin_at", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub last_loggedin_at: Option<Option<i64>>,
    #[serde(rename = "users", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub users: Option<Option<Vec<models::RealmUser>>>,
}

impl ActiveFollowingsResponse {
    pub fn new() -> ActiveFollowingsResponse {
        ActiveFollowingsResponse {
            last_loggedin_at: None,
            users: None,
        }
    }
}

