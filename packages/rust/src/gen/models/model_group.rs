// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct ModelGroup {
    #[serde(rename = "allow_members_to_post_media", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub allow_members_to_post_media: Option<Option<bool>>,
    #[serde(rename = "allow_members_to_post_url", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub allow_members_to_post_url: Option<Option<bool>>,
    #[serde(rename = "allow_ownership_transfer", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub allow_ownership_transfer: Option<Option<bool>>,
    #[serde(rename = "allow_thread_creation_by", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub allow_thread_creation_by: Option<Option<models::GroupRole>>,
    #[serde(rename = "cover_image", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub cover_image: Option<Option<String>>,
    #[serde(rename = "cover_image_thumbnail", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub cover_image_thumbnail: Option<Option<String>>,
    #[serde(rename = "description", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub description: Option<Option<String>>,
    #[serde(rename = "display_icon_thumbnail", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub display_icon_thumbnail: Option<Option<String>>,
    #[serde(rename = "gender", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub gender: Option<Option<models::Gender>>,
    #[serde(rename = "generation", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub generation: Option<Option<models::Generation>>,
    #[serde(rename = "group_category_id", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub group_category_id: Option<Option<i64>>,
    #[serde(rename = "groups_users_count", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub groups_users_count: Option<Option<i64>>,
    #[serde(rename = "guidelines", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub guidelines: Option<Option<String>>,
    #[serde(rename = "highlighted_count", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub highlighted_count: Option<Option<i64>>,
    #[serde(rename = "icon_image", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub icon_image: Option<Option<String>>,
    #[serde(rename = "icon_image_thumbnail", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub icon_image_thumbnail: Option<Option<String>>,
    #[serde(rename = "id", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub id: Option<Option<i64>>,
    #[serde(rename = "invited_to_join", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub invited_to_join: Option<Option<bool>>,
    #[serde(rename = "is_call_timeline_display", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub is_call_timeline_display: Option<Option<bool>>,
    #[serde(rename = "is_hide_conference_call", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub is_hide_conference_call: Option<Option<bool>>,
    #[serde(rename = "is_hide_from_game_eight", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub is_hide_from_game_eight: Option<Option<bool>>,
    #[serde(rename = "is_hide_reported_posts", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub is_hide_reported_posts: Option<Option<bool>>,
    #[serde(rename = "is_joined", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub is_joined: Option<Option<bool>>,
    #[serde(rename = "is_only_verified_age", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub is_only_verified_age: Option<Option<bool>>,
    #[serde(rename = "is_only_verified_phone", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub is_only_verified_phone: Option<Option<bool>>,
    #[serde(rename = "is_pending", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub is_pending: Option<Option<bool>>,
    #[serde(rename = "is_pinned", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub is_pinned: Option<Option<bool>>,
    #[serde(rename = "is_private", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub is_private: Option<Option<bool>>,
    #[serde(rename = "is_safe_mode", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub is_safe_mode: Option<Option<bool>>,
    #[serde(rename = "is_secret", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub is_secret: Option<Option<bool>>,
    #[serde(rename = "is_seizable", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub is_seizable: Option<Option<bool>>,
    #[serde(rename = "is_walkthrough_requested", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub is_walkthrough_requested: Option<Option<bool>>,
    #[serde(rename = "join_status", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub join_status: Option<Option<serde_json::Value>>,
    #[serde(rename = "joined_community_campaign", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub joined_community_campaign: Option<Option<bool>>,
    #[serde(rename = "moderator_ids", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub moderator_ids: Option<Option<Vec<i64>>>,
    #[serde(rename = "owner", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub owner: Option<Option<models::User>>,
    #[serde(rename = "pending_count", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub pending_count: Option<Option<i64>>,
    #[serde(rename = "pending_deputize_ids", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub pending_deputize_ids: Option<Option<Vec<i64>>>,
    #[serde(rename = "pending_transfer_id", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub pending_transfer_id: Option<Option<i64>>,
    #[serde(rename = "posts_count", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub posts_count: Option<Option<i64>>,
    #[serde(rename = "related_count", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub related_count: Option<Option<i64>>,
    #[serde(rename = "seizable_before_millis", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub seizable_before_millis: Option<Option<i64>>,
    #[serde(rename = "sub_category_id", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub sub_category_id: Option<Option<i64>>,
    #[serde(rename = "threads_count", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub threads_count: Option<Option<i64>>,
    #[serde(rename = "title", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub title: Option<Option<serde_json::Value>>,
    #[serde(rename = "topic", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub topic: Option<Option<String>>,
    #[serde(rename = "unread_counts", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub unread_counts: Option<Option<i64>>,
    #[serde(rename = "unread_threads_count", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub unread_threads_count: Option<Option<i64>>,
    #[serde(rename = "updated_at_seconds", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub updated_at_seconds: Option<Option<i64>>,
    #[serde(rename = "views_count", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub views_count: Option<Option<i64>>,
}

impl ModelGroup {
    pub fn new() -> ModelGroup {
        ModelGroup {
            allow_members_to_post_media: None,
            allow_members_to_post_url: None,
            allow_ownership_transfer: None,
            allow_thread_creation_by: None,
            cover_image: None,
            cover_image_thumbnail: None,
            description: None,
            display_icon_thumbnail: None,
            gender: None,
            generation: None,
            group_category_id: None,
            groups_users_count: None,
            guidelines: None,
            highlighted_count: None,
            icon_image: None,
            icon_image_thumbnail: None,
            id: None,
            invited_to_join: None,
            is_call_timeline_display: None,
            is_hide_conference_call: None,
            is_hide_from_game_eight: None,
            is_hide_reported_posts: None,
            is_joined: None,
            is_only_verified_age: None,
            is_only_verified_phone: None,
            is_pending: None,
            is_pinned: None,
            is_private: None,
            is_safe_mode: None,
            is_secret: None,
            is_seizable: None,
            is_walkthrough_requested: None,
            join_status: None,
            joined_community_campaign: None,
            moderator_ids: None,
            owner: None,
            pending_count: None,
            pending_deputize_ids: None,
            pending_transfer_id: None,
            posts_count: None,
            related_count: None,
            seizable_before_millis: None,
            sub_category_id: None,
            threads_count: None,
            title: None,
            topic: None,
            unread_counts: None,
            unread_threads_count: None,
            updated_at_seconds: None,
            views_count: None,
        }
    }
}

