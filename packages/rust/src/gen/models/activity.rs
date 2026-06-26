// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct Activity {
    #[serde(rename = "birthday_users", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub birthday_users: Option<Option<Vec<models::RealmUser>>>,
    #[serde(rename = "birthday_users_count", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub birthday_users_count: Option<Option<i32>>,
    #[serde(rename = "created_at", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub created_at: Option<Option<i64>>,
    #[serde(rename = "empl_amount", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub empl_amount: Option<Option<i32>>,
    #[serde(rename = "followers", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub followers: Option<Option<Vec<models::RealmUser>>>,
    #[serde(rename = "followers_count", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub followers_count: Option<Option<i32>>,
    #[serde(rename = "from_post", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub from_post: Option<Option<models::Post>>,
    #[serde(rename = "from_post_ids", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub from_post_ids: Option<Option<Vec<i64>>>,
    #[serde(rename = "gift_item", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub gift_item: Option<Option<models::RealmGift>>,
    #[serde(rename = "group", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub group: Option<Option<models::Group>>,
    #[serde(rename = "id", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub id: Option<Option<i64>>,
    #[serde(rename = "metadata", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub metadata: Option<Option<models::Metadata>>,
    #[serde(rename = "to_post", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub to_post: Option<Option<models::Post>>,
    #[serde(rename = "type", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub r#type: Option<Option<String>>,
    #[serde(rename = "user", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub user: Option<Option<models::RealmUser>>,
    #[serde(rename = "vip_reward", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub vip_reward: Option<Option<i32>>,
}

impl Activity {
    pub fn new() -> Activity {
        Activity {
            birthday_users: None,
            birthday_users_count: None,
            created_at: None,
            empl_amount: None,
            followers: None,
            followers_count: None,
            from_post: None,
            from_post_ids: None,
            gift_item: None,
            group: None,
            id: None,
            metadata: None,
            to_post: None,
            r#type: None,
            user: None,
            vip_reward: None,
        }
    }
}

