// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct User {
    #[serde(rename = "biography", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub biography: Option<Option<String>>,
    #[serde(rename = "birth_date", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub birth_date: Option<Option<String>>,
    #[serde(rename = "blocking_limit", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub blocking_limit: Option<Option<i32>>,
    #[serde(rename = "connected_by", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub connected_by: Option<Option<serde_json::Value>>,
    #[serde(rename = "contact_phones", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub contact_phones: Option<Option<Vec<String>>>,
    #[serde(rename = "country", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub country: Option<Option<serde_json::Value>>,
    #[serde(rename = "cover_image", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub cover_image: Option<Option<String>>,
    #[serde(rename = "cover_image_thumbnail", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub cover_image_thumbnail: Option<Option<String>>,
    #[serde(rename = "created_at_millis", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub created_at_millis: Option<Option<i64>>,
    #[serde(rename = "followers_count", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub followers_count: Option<Option<i32>>,
    #[serde(rename = "followings_count", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub followings_count: Option<Option<i32>>,
    #[serde(rename = "gender", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub gender: Option<Option<models::Gender>>,
    #[serde(rename = "gift_received_counter", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub gift_received_counter: Option<Option<i32>>,
    #[serde(rename = "gifting_ability", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub gifting_ability: Option<Option<models::GiftingAbility>>,
    #[serde(rename = "groups_users_count", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub groups_users_count: Option<Option<i32>>,
    #[serde(rename = "id", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub id: Option<Option<i64>>,
    #[serde(rename = "interests_count", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub interests_count: Option<Option<i32>>,
    #[serde(rename = "is_age_verified", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub is_age_verified: Option<Option<bool>>,
    #[serde(rename = "is_apple_connected", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub is_apple_connected: Option<Option<bool>>,
    #[serde(rename = "is_chat_request_on", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub is_chat_request_on: Option<Option<bool>>,
    #[serde(rename = "is_dangerous", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub is_dangerous: Option<Option<bool>>,
    #[serde(rename = "is_email_confirmed", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub is_email_confirmed: Option<Option<bool>>,
    #[serde(rename = "is_facebook_connected", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub is_facebook_connected: Option<Option<bool>>,
    #[serde(rename = "is_follow_pending", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub is_follow_pending: Option<Option<bool>>,
    #[serde(rename = "is_followed_by", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub is_followed_by: Option<Option<bool>>,
    #[serde(rename = "is_followed_by_pending", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub is_followed_by_pending: Option<Option<bool>>,
    #[serde(rename = "is_following", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub is_following: Option<Option<bool>>,
    #[serde(rename = "is_from_different_generation_and_trusted", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub is_from_different_generation_and_trusted: Option<Option<bool>>,
    #[serde(rename = "is_google_connected", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub is_google_connected: Option<Option<bool>>,
    #[serde(rename = "is_group_phone_on", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub is_group_phone_on: Option<Option<bool>>,
    #[serde(rename = "is_group_video_on", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub is_group_video_on: Option<Option<bool>>,
    #[serde(rename = "is_hidden", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub is_hidden: Option<Option<bool>>,
    #[serde(rename = "is_hide_vip", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub is_hide_vip: Option<Option<bool>>,
    #[serde(rename = "is_interests_selected", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub is_interests_selected: Option<Option<bool>>,
    #[serde(rename = "is_line_connected", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub is_line_connected: Option<Option<bool>>,
    #[serde(rename = "is_muted", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub is_muted: Option<Option<bool>>,
    #[serde(rename = "is_new", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub is_new: Option<Option<bool>>,
    #[serde(rename = "is_phone_on", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub is_phone_on: Option<Option<bool>>,
    #[serde(rename = "is_private", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub is_private: Option<Option<bool>>,
    #[serde(rename = "is_recently_penalized", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub is_recently_penalized: Option<Option<bool>>,
    #[serde(rename = "is_registered", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub is_registered: Option<Option<bool>>,
    #[serde(rename = "is_two_fa_enabled", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub is_two_fa_enabled: Option<Option<bool>>,
    #[serde(rename = "is_video_on", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub is_video_on: Option<Option<bool>>,
    #[serde(rename = "is_vip", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub is_vip: Option<Option<bool>>,
    #[serde(rename = "is_world_id_connected", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub is_world_id_connected: Option<Option<bool>>,
    #[serde(rename = "last_logged_in_at_millis", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub last_logged_in_at_millis: Option<Option<i64>>,
    #[serde(rename = "login_streak_count", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub login_streak_count: Option<Option<i32>>,
    #[serde(rename = "masked_email", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub masked_email: Option<Option<String>>,
    #[serde(rename = "matching_id", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub matching_id: Option<Option<i64>>,
    #[serde(rename = "media_count", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub media_count: Option<Option<i32>>,
    #[serde(rename = "mobile_number", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub mobile_number: Option<Option<String>>,
    #[serde(rename = "nickname", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub nickname: Option<Option<String>>,
    #[serde(rename = "online_status", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub online_status: Option<Option<models::OnlineStatus>>,
    #[serde(rename = "posts_count", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub posts_count: Option<Option<i32>>,
    #[serde(rename = "prefecture", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub prefecture: Option<Option<String>>,
    #[serde(rename = "profile_icon", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub profile_icon: Option<Option<String>>,
    #[serde(rename = "profile_icon_frame", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub profile_icon_frame: Option<Option<String>>,
    #[serde(rename = "profile_icon_frame_thumbnail", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub profile_icon_frame_thumbnail: Option<Option<String>>,
    #[serde(rename = "profile_icon_thumbnail", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub profile_icon_thumbnail: Option<Option<String>>,
    #[serde(rename = "review_restriction", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub review_restriction: Option<Option<models::ReviewRestriction>>,
    #[serde(rename = "reviews_count", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub reviews_count: Option<Option<i32>>,
    #[serde(rename = "title", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub title: Option<Option<models::Title>>,
    #[serde(rename = "twitter_id", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub twitter_id: Option<Option<String>>,
    #[serde(rename = "username", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub username: Option<Option<String>>,
    #[serde(rename = "uuid", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub uuid: Option<Option<String>>,
    #[serde(rename = "vip_until_millis", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub vip_until_millis: Option<Option<i64>>,
}

impl User {
    pub fn new() -> User {
        User {
            biography: None,
            birth_date: None,
            blocking_limit: None,
            connected_by: None,
            contact_phones: None,
            country: None,
            cover_image: None,
            cover_image_thumbnail: None,
            created_at_millis: None,
            followers_count: None,
            followings_count: None,
            gender: None,
            gift_received_counter: None,
            gifting_ability: None,
            groups_users_count: None,
            id: None,
            interests_count: None,
            is_age_verified: None,
            is_apple_connected: None,
            is_chat_request_on: None,
            is_dangerous: None,
            is_email_confirmed: None,
            is_facebook_connected: None,
            is_follow_pending: None,
            is_followed_by: None,
            is_followed_by_pending: None,
            is_following: None,
            is_from_different_generation_and_trusted: None,
            is_google_connected: None,
            is_group_phone_on: None,
            is_group_video_on: None,
            is_hidden: None,
            is_hide_vip: None,
            is_interests_selected: None,
            is_line_connected: None,
            is_muted: None,
            is_new: None,
            is_phone_on: None,
            is_private: None,
            is_recently_penalized: None,
            is_registered: None,
            is_two_fa_enabled: None,
            is_video_on: None,
            is_vip: None,
            is_world_id_connected: None,
            last_logged_in_at_millis: None,
            login_streak_count: None,
            masked_email: None,
            matching_id: None,
            media_count: None,
            mobile_number: None,
            nickname: None,
            online_status: None,
            posts_count: None,
            prefecture: None,
            profile_icon: None,
            profile_icon_frame: None,
            profile_icon_frame_thumbnail: None,
            profile_icon_thumbnail: None,
            review_restriction: None,
            reviews_count: None,
            title: None,
            twitter_id: None,
            username: None,
            uuid: None,
            vip_until_millis: None,
        }
    }
}

