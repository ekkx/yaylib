// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct LoginSnsUserRequest {
    #[serde(rename = "access_token", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub access_token: Option<Option<String>>,
    #[serde(rename = "access_token_secret", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub access_token_secret: Option<Option<String>>,
    #[serde(rename = "api_key", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub api_key: Option<Option<String>>,
    #[serde(rename = "authorization_code", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub authorization_code: Option<Option<String>>,
    #[serde(rename = "consumer_key", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub consumer_key: Option<Option<String>>,
    #[serde(rename = "consumer_secret", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub consumer_secret: Option<Option<String>>,
    #[serde(rename = "email", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub email: Option<Option<String>>,
    #[serde(rename = "provider", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub provider: Option<Option<String>>,
    #[serde(rename = "two_f_a_code", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub two_f_a_code: Option<Option<String>>,
}

impl LoginSnsUserRequest {
    pub fn new() -> LoginSnsUserRequest {
        LoginSnsUserRequest {
            access_token: None,
            access_token_secret: None,
            api_key: None,
            authorization_code: None,
            consumer_key: None,
            consumer_secret: None,
            email: None,
            provider: None,
            two_f_a_code: None,
        }
    }
}

