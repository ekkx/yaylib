// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct ModelActivity {
    #[serde(rename = "birthday_users", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub birthday_users: Option<Option<Vec<models::User>>>,
    #[serde(rename = "birthday_users_count", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub birthday_users_count: Option<Option<i32>>,
    #[serde(rename = "body", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub body: Option<Option<String>>,
    #[serde(rename = "content_preview", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub content_preview: Option<Option<String>>,
    #[serde(rename = "created_at_millis", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub created_at_millis: Option<Option<i64>>,
    #[serde(rename = "empl_amount", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub empl_amount: Option<Option<i32>>,
    #[serde(rename = "followers", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub followers: Option<Option<Vec<models::User>>>,
    #[serde(rename = "followers_count", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub followers_count: Option<Option<i32>>,
    #[serde(rename = "from_post", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub from_post: Option<Option<models::ModelPost>>,
    #[serde(rename = "from_post_ids", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub from_post_ids: Option<Option<Vec<i64>>>,
    #[serde(rename = "gift_item", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub gift_item: Option<Option<models::Gift>>,
    #[serde(rename = "group", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub group: Option<Option<models::ModelGroup>>,
    #[serde(rename = "group_topic", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub group_topic: Option<Option<String>>,
    #[serde(rename = "id", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub id: Option<Option<i64>>,
    #[serde(rename = "is_anonymous", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub is_anonymous: Option<Option<bool>>,
    #[serde(rename = "is_bulk_invitation", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub is_bulk_invitation: Option<Option<bool>>,
    #[serde(rename = "title", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub title: Option<Option<String>>,
    #[serde(rename = "to_post", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub to_post: Option<Option<models::ModelPost>>,
    #[serde(rename = "type", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub r#type: Option<Option<serde_json::Value>>,
    #[serde(rename = "url", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub url: Option<Option<String>>,
    #[serde(rename = "user", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub user: Option<Option<models::User>>,
    #[serde(rename = "vip_reward", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub vip_reward: Option<Option<i32>>,
}

impl ModelActivity {
    pub fn new() -> ModelActivity {
        ModelActivity {
            birthday_users: None,
            birthday_users_count: None,
            body: None,
            content_preview: None,
            created_at_millis: None,
            empl_amount: None,
            followers: None,
            followers_count: None,
            from_post: None,
            from_post_ids: None,
            gift_item: None,
            group: None,
            group_topic: None,
            id: None,
            is_anonymous: None,
            is_bulk_invitation: None,
            title: None,
            to_post: None,
            r#type: None,
            url: None,
            user: None,
            vip_reward: None,
        }
    }
}

