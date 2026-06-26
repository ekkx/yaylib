// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct LoginEmailUserRequest {
    #[serde(rename = "api_key", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub api_key: Option<Option<String>>,
    #[serde(rename = "email", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub email: Option<Option<String>>,
    #[serde(rename = "password", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub password: Option<Option<String>>,
    #[serde(rename = "two_f_a_code", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub two_f_a_code: Option<Option<String>>,
    #[serde(rename = "uuid", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub uuid: Option<Option<String>>,
}

impl LoginEmailUserRequest {
    pub fn new() -> LoginEmailUserRequest {
        LoginEmailUserRequest {
            api_key: None,
            email: None,
            password: None,
            two_f_a_code: None,
            uuid: None,
        }
    }
}

