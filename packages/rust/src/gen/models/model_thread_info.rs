// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct ModelThreadInfo {
    #[serde(rename = "created_at_seconds", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub created_at_seconds: Option<Option<i64>>,
    #[serde(rename = "has_new_updates", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub has_new_updates: Option<Option<bool>>,
    #[serde(rename = "id", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub id: Option<Option<i64>>,
    #[serde(rename = "is_joined", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub is_joined: Option<Option<bool>>,
    #[serde(rename = "last_post", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub last_post: Option<Option<Box<models::ModelPost>>>,
    #[serde(rename = "owner", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub owner: Option<Option<models::User>>,
    #[serde(rename = "posts_count", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub posts_count: Option<Option<i32>>,
    #[serde(rename = "thread_icon", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub thread_icon: Option<Option<String>>,
    #[serde(rename = "title", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub title: Option<Option<String>>,
    #[serde(rename = "unread_count", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub unread_count: Option<Option<i32>>,
    #[serde(rename = "updated_at_seconds", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub updated_at_seconds: Option<Option<i64>>,
}

impl ModelThreadInfo {
    pub fn new() -> ModelThreadInfo {
        ModelThreadInfo {
            created_at_seconds: None,
            has_new_updates: None,
            id: None,
            is_joined: None,
            last_post: None,
            owner: None,
            posts_count: None,
            thread_icon: None,
            title: None,
            unread_count: None,
            updated_at_seconds: None,
        }
    }
}

