// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct RealmMessage {
    #[serde(rename = "attachment", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub attachment: Option<Option<String>>,
    #[serde(rename = "attachment_android", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub attachment_android: Option<Option<String>>,
    #[serde(rename = "attachment_read_count", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub attachment_read_count: Option<Option<i32>>,
    #[serde(rename = "attachment_thumbnail", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub attachment_thumbnail: Option<Option<String>>,
    #[serde(rename = "conference_call", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub conference_call: Option<Option<models::RealmConferenceCall>>,
    #[serde(rename = "created_at", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub created_at: Option<Option<i64>>,
    #[serde(rename = "font_size", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub font_size: Option<Option<i32>>,
    #[serde(rename = "gif", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub gif: Option<Option<models::GifImage>>,
    #[serde(rename = "id", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub id: Option<Option<i64>>,
    #[serde(rename = "invitation", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub invitation: Option<Option<models::ChatInvitation>>,
    #[serde(rename = "is_error", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub is_error: Option<Option<bool>>,
    #[serde(rename = "is_sent", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub is_sent: Option<Option<bool>>,
    #[serde(rename = "message_type", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub message_type: Option<Option<models::MessageType>>,
    #[serde(rename = "parent", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub parent: Option<Option<models::ParentMessage>>,
    #[serde(rename = "refresh_retry_count", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub refresh_retry_count: Option<Option<i32>>,
    #[serde(rename = "room_id", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub room_id: Option<Option<i64>>,
    #[serde(rename = "sticker", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub sticker: Option<Option<models::Sticker>>,
    #[serde(rename = "text", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub text: Option<Option<String>>,
    #[serde(rename = "user_id", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub user_id: Option<Option<i64>>,
    #[serde(rename = "video_processed", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub video_processed: Option<Option<bool>>,
    #[serde(rename = "video_thumbnail_url", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub video_thumbnail_url: Option<Option<String>>,
    #[serde(rename = "video_url", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub video_url: Option<Option<String>>,
}

impl RealmMessage {
    pub fn new() -> RealmMessage {
        RealmMessage {
            attachment: None,
            attachment_android: None,
            attachment_read_count: None,
            attachment_thumbnail: None,
            conference_call: None,
            created_at: None,
            font_size: None,
            gif: None,
            id: None,
            invitation: None,
            is_error: None,
            is_sent: None,
            message_type: None,
            parent: None,
            refresh_retry_count: None,
            room_id: None,
            sticker: None,
            text: None,
            user_id: None,
            video_processed: None,
            video_thumbnail_url: None,
            video_url: None,
        }
    }
}

