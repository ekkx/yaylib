// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct RealmChatRoom {
    #[serde(rename = "background", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub background: Option<Option<String>>,
    #[serde(rename = "id", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub id: Option<Option<i64>>,
    #[serde(rename = "is_group", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub is_group: Option<Option<bool>>,
    #[serde(rename = "is_request", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub is_request: Option<Option<bool>>,
    #[serde(rename = "last_message", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub last_message: Option<Option<models::ChatRoomLastMessage>>,
    #[serde(rename = "members", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub members: Option<Option<Vec<models::RealmUser>>>,
    #[serde(rename = "name", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub name: Option<Option<String>>,
    #[serde(rename = "owner", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub owner: Option<Option<models::RealmUser>>,
    #[serde(rename = "unread_count", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub unread_count: Option<Option<i32>>,
    #[serde(rename = "updated_at", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub updated_at: Option<Option<i64>>,
    #[serde(rename = "user_setting", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub user_setting: Option<Option<models::UserSetting>>,
}

impl RealmChatRoom {
    pub fn new() -> RealmChatRoom {
        RealmChatRoom {
            background: None,
            id: None,
            is_group: None,
            is_request: None,
            last_message: None,
            members: None,
            name: None,
            owner: None,
            unread_count: None,
            updated_at: None,
            user_setting: None,
        }
    }
}

