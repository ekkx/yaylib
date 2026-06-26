// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct SearchCriteria {
    #[serde(rename = "biography", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub biography: Option<Option<String>>,
    #[serde(rename = "gender", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub gender: Option<Option<i32>>,
    #[serde(rename = "nickname", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub nickname: Option<Option<String>>,
    #[serde(rename = "prefecture", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub prefecture: Option<Option<String>>,
    #[serde(rename = "username", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub username: Option<Option<String>>,
}

impl SearchCriteria {
    pub fn new() -> SearchCriteria {
        SearchCriteria {
            biography: None,
            gender: None,
            nickname: None,
            prefecture: None,
            username: None,
        }
    }
}

