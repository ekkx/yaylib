// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct Group {
    #[serde(rename = "allow_members_to_post_image_and_video", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub allow_members_to_post_image_and_video: Option<Option<bool>>,
    #[serde(rename = "allow_members_to_post_url", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub allow_members_to_post_url: Option<Option<bool>>,
    #[serde(rename = "allow_ownership_transfer", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub allow_ownership_transfer: Option<Option<bool>>,
    #[serde(rename = "allow_thread_creation_by", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub allow_thread_creation_by: Option<Option<String>>,
    #[serde(rename = "call_timeline_display", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub call_timeline_display: Option<Option<bool>>,
    #[serde(rename = "cover_image", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub cover_image: Option<Option<String>>,
    #[serde(rename = "cover_image_thumbnail", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub cover_image_thumbnail: Option<Option<String>>,
    #[serde(rename = "description", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub description: Option<Option<String>>,
    #[serde(rename = "gender", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub gender: Option<Option<i32>>,
    #[serde(rename = "generation_groups_limit", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub generation_groups_limit: Option<Option<i32>>,
    #[serde(rename = "group_category_id", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub group_category_id: Option<Option<i64>>,
    #[serde(rename = "group_icon", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub group_icon: Option<Option<String>>,
    #[serde(rename = "group_icon_thumbnail", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub group_icon_thumbnail: Option<Option<String>>,
    #[serde(rename = "groups_users_count", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub groups_users_count: Option<Option<i64>>,
    #[serde(rename = "guidelines", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub guidelines: Option<Option<String>>,
    #[serde(rename = "hide_conference_call", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub hide_conference_call: Option<Option<bool>>,
    #[serde(rename = "hide_from_game_eight", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub hide_from_game_eight: Option<Option<bool>>,
    #[serde(rename = "hide_reported_posts", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub hide_reported_posts: Option<Option<bool>>,
    #[serde(rename = "highlighted_count", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub highlighted_count: Option<Option<i64>>,
    #[serde(rename = "id", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub id: Option<Option<i64>>,
    #[serde(rename = "invited_to_join", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub invited_to_join: Option<Option<bool>>,
    #[serde(rename = "is_joined", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub is_joined: Option<Option<bool>>,
    #[serde(rename = "is_pending", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub is_pending: Option<Option<bool>>,
    #[serde(rename = "is_private", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub is_private: Option<Option<bool>>,
    #[serde(rename = "is_related", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub is_related: Option<Option<bool>>,
    #[serde(rename = "joined_community_campaign", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub joined_community_campaign: Option<Option<bool>>,
    #[serde(rename = "moderator_ids", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub moderator_ids: Option<Option<Vec<i64>>>,
    #[serde(rename = "only_mobile_verified", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub only_mobile_verified: Option<Option<bool>>,
    #[serde(rename = "only_verified_age", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub only_verified_age: Option<Option<bool>>,
    #[serde(rename = "owner", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub owner: Option<Option<models::RealmUser>>,
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
    #[serde(rename = "safe_mode", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub safe_mode: Option<Option<bool>>,
    #[serde(rename = "secret", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub secret: Option<Option<bool>>,
    #[serde(rename = "seizable", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub seizable: Option<Option<bool>>,
    #[serde(rename = "seizable_before", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub seizable_before: Option<Option<i64>>,
    #[serde(rename = "sub_category_id", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub sub_category_id: Option<Option<i64>>,
    #[serde(rename = "threads_count", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub threads_count: Option<Option<i64>>,
    #[serde(rename = "title", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub title: Option<Option<String>>,
    #[serde(rename = "topic", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub topic: Option<Option<String>>,
    #[serde(rename = "unread_counts", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub unread_counts: Option<Option<i64>>,
    #[serde(rename = "unread_threads_count", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub unread_threads_count: Option<Option<i64>>,
    #[serde(rename = "updated_at", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub updated_at: Option<Option<i64>>,
    #[serde(rename = "user_id", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub user_id: Option<Option<i64>>,
    #[serde(rename = "views_count", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub views_count: Option<Option<i64>>,
    #[serde(rename = "walkthrough_requested", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub walkthrough_requested: Option<Option<bool>>,
}

impl Group {
    pub fn new() -> Group {
        Group {
            allow_members_to_post_image_and_video: None,
            allow_members_to_post_url: None,
            allow_ownership_transfer: None,
            allow_thread_creation_by: None,
            call_timeline_display: None,
            cover_image: None,
            cover_image_thumbnail: None,
            description: None,
            gender: None,
            generation_groups_limit: None,
            group_category_id: None,
            group_icon: None,
            group_icon_thumbnail: None,
            groups_users_count: None,
            guidelines: None,
            hide_conference_call: None,
            hide_from_game_eight: None,
            hide_reported_posts: None,
            highlighted_count: None,
            id: None,
            invited_to_join: None,
            is_joined: None,
            is_pending: None,
            is_private: None,
            is_related: None,
            joined_community_campaign: None,
            moderator_ids: None,
            only_mobile_verified: None,
            only_verified_age: None,
            owner: None,
            pending_count: None,
            pending_deputize_ids: None,
            pending_transfer_id: None,
            posts_count: None,
            related_count: None,
            safe_mode: None,
            secret: None,
            seizable: None,
            seizable_before: None,
            sub_category_id: None,
            threads_count: None,
            title: None,
            topic: None,
            unread_counts: None,
            unread_threads_count: None,
            updated_at: None,
            user_id: None,
            views_count: None,
            walkthrough_requested: None,
        }
    }
}

