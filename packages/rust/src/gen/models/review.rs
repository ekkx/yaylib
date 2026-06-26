// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct Review {
    #[serde(rename = "comment", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub comment: Option<Option<String>>,
    #[serde(rename = "created_at", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub created_at: Option<Option<i64>>,
    #[serde(rename = "id", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub id: Option<Option<i64>>,
    #[serde(rename = "mutual_review", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub mutual_review: Option<Option<bool>>,
    #[serde(rename = "reported_count", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub reported_count: Option<Option<i32>>,
    #[serde(rename = "reviewer", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub reviewer: Option<Option<models::RealmUser>>,
    #[serde(rename = "user", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub user: Option<Option<models::RealmUser>>,
}

impl Review {
    pub fn new() -> Review {
        Review {
            comment: None,
            created_at: None,
            id: None,
            mutual_review: None,
            reported_count: None,
            reviewer: None,
            user: None,
        }
    }
}

