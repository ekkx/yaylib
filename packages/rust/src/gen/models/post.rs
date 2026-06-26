// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct Post {
    #[serde(rename = "attachment", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub attachment: Option<Option<String>>,
    #[serde(rename = "attachment_2", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub attachment_2: Option<Option<String>>,
    #[serde(rename = "attachment_2_thumbnail", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub attachment_2_thumbnail: Option<Option<String>>,
    #[serde(rename = "attachment_3", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub attachment_3: Option<Option<String>>,
    #[serde(rename = "attachment_3_thumbnail", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub attachment_3_thumbnail: Option<Option<String>>,
    #[serde(rename = "attachment_4", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub attachment_4: Option<Option<String>>,
    #[serde(rename = "attachment_4_thumbnail", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub attachment_4_thumbnail: Option<Option<String>>,
    #[serde(rename = "attachment_5", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub attachment_5: Option<Option<String>>,
    #[serde(rename = "attachment_5_thumbnail", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub attachment_5_thumbnail: Option<Option<String>>,
    #[serde(rename = "attachment_6", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub attachment_6: Option<Option<String>>,
    #[serde(rename = "attachment_6_thumbnail", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub attachment_6_thumbnail: Option<Option<String>>,
    #[serde(rename = "attachment_7", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub attachment_7: Option<Option<String>>,
    #[serde(rename = "attachment_7_thumbnail", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub attachment_7_thumbnail: Option<Option<String>>,
    #[serde(rename = "attachment_8", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub attachment_8: Option<Option<String>>,
    #[serde(rename = "attachment_8_thumbnail", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub attachment_8_thumbnail: Option<Option<String>>,
    #[serde(rename = "attachment_9", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub attachment_9: Option<Option<String>>,
    #[serde(rename = "attachment_9_thumbnail", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub attachment_9_thumbnail: Option<Option<String>>,
    #[serde(rename = "attachment_thumbnail", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub attachment_thumbnail: Option<Option<String>>,
    #[serde(rename = "color", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub color: Option<Option<i32>>,
    #[serde(rename = "conference_call", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub conference_call: Option<Option<models::RealmConferenceCall>>,
    #[serde(rename = "conversation_id", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub conversation_id: Option<Option<i64>>,
    #[serde(rename = "created_at", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub created_at: Option<Option<i64>>,
    #[serde(rename = "edited_at", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub edited_at: Option<Option<i64>>,
    #[serde(rename = "font_size", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub font_size: Option<Option<i32>>,
    #[serde(rename = "gifts_count", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub gifts_count: Option<Option<Vec<models::GiftCount>>>,
    #[serde(rename = "group", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub group: Option<Option<models::Group>>,
    #[serde(rename = "group_id", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub group_id: Option<Option<i64>>,
    #[serde(rename = "highlighted", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub highlighted: Option<Option<bool>>,
    #[serde(rename = "id", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub id: Option<Option<i64>>,
    #[serde(rename = "in_reply_to", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub in_reply_to: Option<Option<i64>>,
    #[serde(rename = "in_reply_to_post", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub in_reply_to_post: Option<Option<Box<models::Post>>>,
    #[serde(rename = "in_reply_to_post_count", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub in_reply_to_post_count: Option<Option<i32>>,
    #[serde(rename = "is_fail_to_send", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub is_fail_to_send: Option<Option<bool>>,
    #[serde(rename = "is_muted", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub is_muted: Option<Option<bool>>,
    #[serde(rename = "liked", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub liked: Option<Option<bool>>,
    #[serde(rename = "likers", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub likers: Option<Option<Vec<models::RealmUser>>>,
    #[serde(rename = "likers_count", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub likers_count: Option<Option<i32>>,
    #[serde(rename = "likes_count", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub likes_count: Option<Option<i32>>,
    #[serde(rename = "mentions", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub mentions: Option<Option<Vec<models::RealmUser>>>,
    #[serde(rename = "message_tags", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub message_tags: Option<Option<Vec<models::MessageTag>>>,
    #[serde(rename = "post_type", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub post_type: Option<Option<models::PostType>>,
    #[serde(rename = "reported_count", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub reported_count: Option<Option<i32>>,
    #[serde(rename = "repostable", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub repostable: Option<Option<bool>>,
    #[serde(rename = "reposted", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub reposted: Option<Option<bool>>,
    #[serde(rename = "reposts_count", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub reposts_count: Option<Option<i32>>,
    #[serde(rename = "shareable", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub shareable: Option<Option<Box<models::Shareable>>>,
    #[serde(rename = "shared_thread", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub shared_thread: Option<Option<Box<models::ThreadInfo>>>,
    #[serde(rename = "shared_url", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub shared_url: Option<Option<models::SharedUrl>>,
    #[serde(rename = "survey", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub survey: Option<Option<models::Survey>>,
    #[serde(rename = "tag", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub tag: Option<Option<String>>,
    #[serde(rename = "text", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub text: Option<Option<String>>,
    #[serde(rename = "thread", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub thread: Option<Option<Box<models::ThreadInfo>>>,
    #[serde(rename = "thread_id", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub thread_id: Option<Option<i64>>,
    #[serde(rename = "total_gifts_count", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub total_gifts_count: Option<Option<i32>>,
    #[serde(rename = "updated_at", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub updated_at: Option<Option<i64>>,
    #[serde(rename = "user", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub user: Option<Option<models::RealmUser>>,
    #[serde(rename = "videos", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub videos: Option<Option<Vec<models::Video>>>,
}

impl Post {
    pub fn new() -> Post {
        Post {
            attachment: None,
            attachment_2: None,
            attachment_2_thumbnail: None,
            attachment_3: None,
            attachment_3_thumbnail: None,
            attachment_4: None,
            attachment_4_thumbnail: None,
            attachment_5: None,
            attachment_5_thumbnail: None,
            attachment_6: None,
            attachment_6_thumbnail: None,
            attachment_7: None,
            attachment_7_thumbnail: None,
            attachment_8: None,
            attachment_8_thumbnail: None,
            attachment_9: None,
            attachment_9_thumbnail: None,
            attachment_thumbnail: None,
            color: None,
            conference_call: None,
            conversation_id: None,
            created_at: None,
            edited_at: None,
            font_size: None,
            gifts_count: None,
            group: None,
            group_id: None,
            highlighted: None,
            id: None,
            in_reply_to: None,
            in_reply_to_post: None,
            in_reply_to_post_count: None,
            is_fail_to_send: None,
            is_muted: None,
            liked: None,
            likers: None,
            likers_count: None,
            likes_count: None,
            mentions: None,
            message_tags: None,
            post_type: None,
            reported_count: None,
            repostable: None,
            reposted: None,
            reposts_count: None,
            shareable: None,
            shared_thread: None,
            shared_url: None,
            survey: None,
            tag: None,
            text: None,
            thread: None,
            thread_id: None,
            total_gifts_count: None,
            updated_at: None,
            user: None,
            videos: None,
        }
    }
}

