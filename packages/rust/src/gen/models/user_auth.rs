// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct UserAuth {
    #[serde(rename = "access_token", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub access_token: Option<Option<String>>,
    #[serde(rename = "expires_in", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub expires_in: Option<Option<i64>>,
    #[serde(rename = "refresh_token", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub refresh_token: Option<Option<String>>,
    #[serde(rename = "user_id", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub user_id: Option<Option<i64>>,
}

impl UserAuth {
    pub fn new() -> UserAuth {
        UserAuth {
            access_token: None,
            expires_in: None,
            refresh_token: None,
            user_id: None,
        }
    }
}

