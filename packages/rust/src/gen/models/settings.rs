// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct Settings {
    #[serde(rename = "age_restricted_on_review", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub age_restricted_on_review: Option<Option<bool>>,
    #[serde(rename = "allow_reposts", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub allow_reposts: Option<Option<bool>>,
    #[serde(rename = "caution_user_chat", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub caution_user_chat: Option<Option<bool>>,
    #[serde(rename = "following_only_call_invite", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub following_only_call_invite: Option<Option<bool>>,
    #[serde(rename = "following_only_group_invite", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub following_only_group_invite: Option<Option<bool>>,
    #[serde(rename = "following_restricted_on_review", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub following_restricted_on_review: Option<Option<bool>>,
    #[serde(rename = "hide_active_call", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub hide_active_call: Option<Option<bool>>,
    #[serde(rename = "hide_gifts_received", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub hide_gifts_received: Option<Option<bool>>,
    #[serde(rename = "hide_hot_post", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub hide_hot_post: Option<Option<bool>>,
    #[serde(rename = "hide_on_invitable", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub hide_on_invitable: Option<Option<bool>>,
    #[serde(rename = "hide_online_status", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub hide_online_status: Option<Option<bool>>,
    #[serde(rename = "hide_reply_following_timeline", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub hide_reply_following_timeline: Option<Option<bool>>,
    #[serde(rename = "hide_reply_public_timeline", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub hide_reply_public_timeline: Option<Option<bool>>,
    #[serde(rename = "hide_vip", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub hide_vip: Option<Option<bool>>,
    #[serde(rename = "invisible_on_user_search", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub invisible_on_user_search: Option<Option<bool>>,
    #[serde(rename = "no_reply_group_timeline", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub no_reply_group_timeline: Option<Option<bool>>,
    #[serde(rename = "notification_birthday_to_followers", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub notification_birthday_to_followers: Option<Option<bool>>,
    #[serde(rename = "notification_bulk_call_invite", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub notification_bulk_call_invite: Option<Option<bool>>,
    #[serde(rename = "notification_call_invite", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub notification_call_invite: Option<Option<bool>>,
    #[serde(rename = "notification_chat", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub notification_chat: Option<Option<bool>>,
    #[serde(rename = "notification_chat_delete", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub notification_chat_delete: Option<Option<bool>>,
    #[serde(rename = "notification_contact_friend", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub notification_contact_friend: Option<Option<bool>>,
    #[serde(rename = "notification_daily_quest", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub notification_daily_quest: Option<Option<bool>>,
    #[serde(rename = "notification_daily_summary", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub notification_daily_summary: Option<Option<bool>>,
    #[serde(rename = "notification_follow", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub notification_follow: Option<Option<bool>>,
    #[serde(rename = "notification_follow_accept", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub notification_follow_accept: Option<Option<bool>>,
    #[serde(rename = "notification_follow_request", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub notification_follow_request: Option<Option<bool>>,
    #[serde(rename = "notification_follower_conference_call", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub notification_follower_conference_call: Option<Option<bool>>,
    #[serde(rename = "notification_follower_create_group", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub notification_follower_create_group: Option<Option<bool>>,
    #[serde(rename = "notification_following_birthdate_on", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub notification_following_birthdate_on: Option<Option<bool>>,
    #[serde(rename = "notification_following_post_after_break", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub notification_following_post_after_break: Option<Option<bool>>,
    #[serde(rename = "notification_followings_in_call", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub notification_followings_in_call: Option<Option<bool>>,
    #[serde(rename = "notification_footprint", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub notification_footprint: Option<Option<bool>>,
    #[serde(rename = "notification_gift_receive", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub notification_gift_receive: Option<Option<bool>>,
    #[serde(rename = "notification_group_accept", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub notification_group_accept: Option<Option<bool>>,
    #[serde(rename = "notification_group_conference_call", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub notification_group_conference_call: Option<Option<bool>>,
    #[serde(rename = "notification_group_invite", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub notification_group_invite: Option<Option<bool>>,
    #[serde(rename = "notification_group_join", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub notification_group_join: Option<Option<bool>>,
    #[serde(rename = "notification_group_message_tag_all", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub notification_group_message_tag_all: Option<Option<bool>>,
    #[serde(rename = "notification_group_moderator", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub notification_group_moderator: Option<Option<bool>>,
    #[serde(rename = "notification_group_post", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub notification_group_post: Option<Option<bool>>,
    #[serde(rename = "notification_group_request", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub notification_group_request: Option<Option<bool>>,
    #[serde(rename = "notification_hima_now", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub notification_hima_now: Option<Option<bool>>,
    #[serde(rename = "notification_invitee_create_account", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub notification_invitee_create_account: Option<Option<bool>>,
    #[serde(rename = "notification_latest_news", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub notification_latest_news: Option<Option<bool>>,
    #[serde(rename = "notification_like", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub notification_like: Option<Option<bool>>,
    #[serde(rename = "notification_message_tag", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub notification_message_tag: Option<Option<bool>>,
    #[serde(rename = "notification_popular_post", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub notification_popular_post: Option<Option<bool>>,
    #[serde(rename = "notification_profile_screenshot", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub notification_profile_screenshot: Option<Option<bool>>,
    #[serde(rename = "notification_reply", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub notification_reply: Option<Option<bool>>,
    #[serde(rename = "notification_repost", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub notification_repost: Option<Option<bool>>,
    #[serde(rename = "notification_review", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub notification_review: Option<Option<bool>>,
    #[serde(rename = "notification_security_warning", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub notification_security_warning: Option<Option<bool>>,
    #[serde(rename = "notification_twitter_friend", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub notification_twitter_friend: Option<Option<bool>>,
    #[serde(rename = "privacy_mode", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub privacy_mode: Option<Option<bool>>,
    #[serde(rename = "private_post", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub private_post: Option<Option<bool>>,
    #[serde(rename = "private_user_timeline", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub private_user_timeline: Option<Option<bool>>,
    #[serde(rename = "show_total_gifts_received_count", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub show_total_gifts_received_count: Option<Option<bool>>,
    #[serde(rename = "vip_invisible_footprint_mode", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub vip_invisible_footprint_mode: Option<Option<bool>>,
    #[serde(rename = "visible_on_sns_friend_recommendation", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub visible_on_sns_friend_recommendation: Option<Option<bool>>,
}

impl Settings {
    pub fn new() -> Settings {
        Settings {
            age_restricted_on_review: None,
            allow_reposts: None,
            caution_user_chat: None,
            following_only_call_invite: None,
            following_only_group_invite: None,
            following_restricted_on_review: None,
            hide_active_call: None,
            hide_gifts_received: None,
            hide_hot_post: None,
            hide_on_invitable: None,
            hide_online_status: None,
            hide_reply_following_timeline: None,
            hide_reply_public_timeline: None,
            hide_vip: None,
            invisible_on_user_search: None,
            no_reply_group_timeline: None,
            notification_birthday_to_followers: None,
            notification_bulk_call_invite: None,
            notification_call_invite: None,
            notification_chat: None,
            notification_chat_delete: None,
            notification_contact_friend: None,
            notification_daily_quest: None,
            notification_daily_summary: None,
            notification_follow: None,
            notification_follow_accept: None,
            notification_follow_request: None,
            notification_follower_conference_call: None,
            notification_follower_create_group: None,
            notification_following_birthdate_on: None,
            notification_following_post_after_break: None,
            notification_followings_in_call: None,
            notification_footprint: None,
            notification_gift_receive: None,
            notification_group_accept: None,
            notification_group_conference_call: None,
            notification_group_invite: None,
            notification_group_join: None,
            notification_group_message_tag_all: None,
            notification_group_moderator: None,
            notification_group_post: None,
            notification_group_request: None,
            notification_hima_now: None,
            notification_invitee_create_account: None,
            notification_latest_news: None,
            notification_like: None,
            notification_message_tag: None,
            notification_popular_post: None,
            notification_profile_screenshot: None,
            notification_reply: None,
            notification_repost: None,
            notification_review: None,
            notification_security_warning: None,
            notification_twitter_friend: None,
            privacy_mode: None,
            private_post: None,
            private_user_timeline: None,
            show_total_gifts_received_count: None,
            vip_invisible_footprint_mode: None,
            visible_on_sns_friend_recommendation: None,
        }
    }
}

