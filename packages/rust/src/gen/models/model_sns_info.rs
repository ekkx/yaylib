// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct ModelSnsInfo {
    #[serde(rename = "biography", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub biography: Option<Option<String>>,
    #[serde(rename = "gender", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub gender: Option<Option<String>>,
    #[serde(rename = "nickname", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub nickname: Option<Option<String>>,
    #[serde(rename = "profile_image", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub profile_image: Option<Option<String>>,
    #[serde(rename = "type", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub r#type: Option<Option<String>>,
    #[serde(rename = "uid", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub uid: Option<Option<String>>,
}

impl ModelSnsInfo {
    pub fn new() -> ModelSnsInfo {
        ModelSnsInfo {
            biography: None,
            gender: None,
            nickname: None,
            profile_image: None,
            r#type: None,
            uid: None,
        }
    }
}

