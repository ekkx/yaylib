// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct UserCustomDefinitionsResponse {
    #[serde(rename = "age", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub age: Option<Option<i32>>,
    #[serde(rename = "created_at", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub created_at: Option<Option<i64>>,
    #[serde(rename = "followers_count", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub followers_count: Option<Option<i32>>,
    #[serde(rename = "followings_count", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub followings_count: Option<Option<i32>>,
    #[serde(rename = "last_loggedin_at", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub last_loggedin_at: Option<Option<i64>>,
    #[serde(rename = "reported_count", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub reported_count: Option<Option<i32>>,
    #[serde(rename = "status", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub status: Option<Option<String>>,
}

impl UserCustomDefinitionsResponse {
    pub fn new() -> UserCustomDefinitionsResponse {
        UserCustomDefinitionsResponse {
            age: None,
            created_at: None,
            followers_count: None,
            followings_count: None,
            last_loggedin_at: None,
            reported_count: None,
            status: None,
        }
    }
}

