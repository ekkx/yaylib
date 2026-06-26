// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct UserResponse {
    #[serde(rename = "apple_connected", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub apple_connected: Option<Option<bool>>,
    #[serde(rename = "birth_date", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub birth_date: Option<Option<String>>,
    #[serde(rename = "blocking_limit", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub blocking_limit: Option<Option<i32>>,
    #[serde(rename = "email_confirmed", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub email_confirmed: Option<Option<bool>>,
    #[serde(rename = "facebook_connected", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub facebook_connected: Option<Option<bool>>,
    #[serde(rename = "gifting_ability", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub gifting_ability: Option<Option<models::RealmGiftingAbility>>,
    #[serde(rename = "google_connected", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub google_connected: Option<Option<bool>>,
    #[serde(rename = "group_phone_on", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub group_phone_on: Option<Option<bool>>,
    #[serde(rename = "group_video_on", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub group_video_on: Option<Option<bool>>,
    #[serde(rename = "interests_count", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub interests_count: Option<Option<i32>>,
    #[serde(rename = "line_connected", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub line_connected: Option<Option<bool>>,
    #[serde(rename = "masked_email", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub masked_email: Option<Option<String>>,
    #[serde(rename = "phone_on", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub phone_on: Option<Option<bool>>,
    #[serde(rename = "push_notification", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub push_notification: Option<Option<bool>>,
    #[serde(rename = "twitter_id", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub twitter_id: Option<Option<String>>,
    #[serde(rename = "user", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub user: Option<Option<models::RealmUser>>,
    #[serde(rename = "uuid", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub uuid: Option<Option<String>>,
    #[serde(rename = "video_on", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub video_on: Option<Option<bool>>,
    #[serde(rename = "vip_until_seconds", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub vip_until_seconds: Option<Option<i64>>,
    #[serde(rename = "world_id_connected", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub world_id_connected: Option<Option<bool>>,
}

impl UserResponse {
    pub fn new() -> UserResponse {
        UserResponse {
            apple_connected: None,
            birth_date: None,
            blocking_limit: None,
            email_confirmed: None,
            facebook_connected: None,
            gifting_ability: None,
            google_connected: None,
            group_phone_on: None,
            group_video_on: None,
            interests_count: None,
            line_connected: None,
            masked_email: None,
            phone_on: None,
            push_notification: None,
            twitter_id: None,
            user: None,
            uuid: None,
            video_on: None,
            vip_until_seconds: None,
            world_id_connected: None,
        }
    }
}

