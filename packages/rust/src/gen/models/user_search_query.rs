// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct UserSearchQuery {
    #[serde(rename = "keyword", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub keyword: Option<Option<String>>,
    #[serde(rename = "no_recent_penalty", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub no_recent_penalty: Option<Option<bool>>,
    #[serde(rename = "same_prefecture", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub same_prefecture: Option<Option<bool>>,
    #[serde(rename = "scope", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub scope: Option<Option<models::UserSearchScope>>,
    #[serde(rename = "similar_age", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub similar_age: Option<Option<bool>>,
}

impl UserSearchQuery {
    pub fn new() -> UserSearchQuery {
        UserSearchQuery {
            keyword: None,
            no_recent_penalty: None,
            same_prefecture: None,
            scope: None,
            similar_age: None,
        }
    }
}

