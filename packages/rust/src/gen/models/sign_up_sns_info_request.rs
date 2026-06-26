// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct SignUpSnsInfoRequest {
    #[serde(rename = "type", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub r#type: Option<Option<String>>,
    #[serde(rename = "uid", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub uid: Option<Option<String>>,
}

impl SignUpSnsInfoRequest {
    pub fn new() -> SignUpSnsInfoRequest {
        SignUpSnsInfoRequest {
            r#type: None,
            uid: None,
        }
    }
}

