// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct LoginUserResponse {
    #[serde(rename = "access_token", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub access_token: Option<Option<String>>,
    #[serde(rename = "expires_in", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub expires_in: Option<Option<i64>>,
    #[serde(rename = "is_new", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub is_new: Option<Option<bool>>,
    #[serde(rename = "refresh_token", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub refresh_token: Option<Option<String>>,
    #[serde(rename = "sns_info", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub sns_info: Option<Option<models::SnsInfo>>,
    #[serde(rename = "user_id", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub user_id: Option<Option<i64>>,
    #[serde(rename = "username", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub username: Option<Option<String>>,
}

impl LoginUserResponse {
    pub fn new() -> LoginUserResponse {
        LoginUserResponse {
            access_token: None,
            expires_in: None,
            is_new: None,
            refresh_token: None,
            sns_info: None,
            user_id: None,
            username: None,
        }
    }
}

