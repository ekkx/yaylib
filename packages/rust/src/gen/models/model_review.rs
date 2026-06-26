// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct ModelReview {
    #[serde(rename = "comment", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub comment: Option<Option<String>>,
    #[serde(rename = "created_at_millis", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub created_at_millis: Option<Option<i64>>,
    #[serde(rename = "id", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub id: Option<Option<i64>>,
    #[serde(rename = "is_mutual", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub is_mutual: Option<Option<bool>>,
    #[serde(rename = "is_pinned", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub is_pinned: Option<Option<bool>>,
    #[serde(rename = "reported_count", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub reported_count: Option<Option<i32>>,
    #[serde(rename = "reviewer", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub reviewer: Option<Option<models::User>>,
    #[serde(rename = "user", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub user: Option<Option<models::User>>,
}

impl ModelReview {
    pub fn new() -> ModelReview {
        ModelReview {
            comment: None,
            created_at_millis: None,
            id: None,
            is_mutual: None,
            is_pinned: None,
            reported_count: None,
            reviewer: None,
            user: None,
        }
    }
}

