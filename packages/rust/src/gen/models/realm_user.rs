// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct RealmUser {
    #[serde(rename = "age_verified", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub age_verified: Option<Option<bool>>,
    #[serde(rename = "apple_connected", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub apple_connected: Option<Option<bool>>,
    #[serde(rename = "biography", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub biography: Option<Option<String>>,
    #[serde(rename = "birth_date", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub birth_date: Option<Option<String>>,
    #[serde(rename = "blocking_limit", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub blocking_limit: Option<Option<i32>>,
    #[serde(rename = "chat_request", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub chat_request: Option<Option<bool>>,
    #[serde(rename = "connected_by", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub connected_by: Option<Option<Vec<String>>>,
    #[serde(rename = "contact_phones", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub contact_phones: Option<Option<Vec<String>>>,
    #[serde(rename = "country_code", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub country_code: Option<Option<String>>,
    #[serde(rename = "cover_image", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub cover_image: Option<Option<String>>,
    #[serde(rename = "cover_image_thumbnail", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub cover_image_thumbnail: Option<Option<String>>,
    #[serde(rename = "created_at_seconds", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub created_at_seconds: Option<Option<i64>>,
    #[serde(rename = "dangerous_user", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub dangerous_user: Option<Option<bool>>,
    #[serde(rename = "email_confirmed", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub email_confirmed: Option<Option<bool>>,
    #[serde(rename = "facebook_connected", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub facebook_connected: Option<Option<bool>>,
    #[serde(rename = "follow_pending", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub follow_pending: Option<Option<bool>>,
    #[serde(rename = "followed_by", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub followed_by: Option<Option<bool>>,
    #[serde(rename = "followed_by_pending", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub followed_by_pending: Option<Option<bool>>,
    #[serde(rename = "followers_count", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub followers_count: Option<Option<i32>>,
    #[serde(rename = "following", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub following: Option<Option<bool>>,
    #[serde(rename = "followings_count", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub followings_count: Option<Option<i32>>,
    #[serde(rename = "frame", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub frame: Option<Option<String>>,
    #[serde(rename = "frame_thumbnail", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub frame_thumbnail: Option<Option<String>>,
    #[serde(rename = "from_different_generation_and_trusted", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub from_different_generation_and_trusted: Option<Option<bool>>,
    #[serde(rename = "gender", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub gender: Option<Option<i32>>,
    #[serde(rename = "generation", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub generation: Option<Option<i32>>,
    #[serde(rename = "gifting_ability", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub gifting_ability: Option<Option<models::RealmGiftingAbility>>,
    #[serde(rename = "google_connected", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub google_connected: Option<Option<bool>>,
    #[serde(rename = "group_phone_on", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub group_phone_on: Option<Option<bool>>,
    #[serde(rename = "group_user", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub group_user: Option<Option<Box<models::GroupUser>>>,
    #[serde(rename = "group_video_on", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub group_video_on: Option<Option<bool>>,
    #[serde(rename = "groups_users_count", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub groups_users_count: Option<Option<i32>>,
    #[serde(rename = "hidden", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub hidden: Option<Option<bool>>,
    #[serde(rename = "hide_vip", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub hide_vip: Option<Option<bool>>,
    #[serde(rename = "id", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub id: Option<Option<i64>>,
    #[serde(rename = "interests_count", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub interests_count: Option<Option<i32>>,
    #[serde(rename = "interests_selected", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub interests_selected: Option<Option<bool>>,
    #[serde(rename = "is_private", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub is_private: Option<Option<bool>>,
    #[serde(rename = "last_logged_in_at_seconds", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub last_logged_in_at_seconds: Option<Option<i64>>,
    #[serde(rename = "line_connected", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub line_connected: Option<Option<bool>>,
    #[serde(rename = "login_streak_count", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub login_streak_count: Option<Option<i32>>,
    #[serde(rename = "masked_email", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub masked_email: Option<Option<String>>,
    #[serde(rename = "media_count", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub media_count: Option<Option<i32>>,
    #[serde(rename = "mobile_number", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub mobile_number: Option<Option<String>>,
    #[serde(rename = "new_user", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub new_user: Option<Option<bool>>,
    #[serde(rename = "nickname", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub nickname: Option<Option<String>>,
    #[serde(rename = "online_status", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub online_status: Option<Option<models::OnlineStatusEnum>>,
    #[serde(rename = "phone_on", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub phone_on: Option<Option<bool>>,
    #[serde(rename = "posts_count", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub posts_count: Option<Option<i32>>,
    #[serde(rename = "prefecture", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub prefecture: Option<Option<String>>,
    #[serde(rename = "profile_icon", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub profile_icon: Option<Option<String>>,
    #[serde(rename = "profile_icon_thumbnail", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub profile_icon_thumbnail: Option<Option<String>>,
    #[serde(rename = "push_notification", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub push_notification: Option<Option<bool>>,
    #[serde(rename = "recently_kenta", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub recently_kenta: Option<Option<bool>>,
    #[serde(rename = "restricted_review_by", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub restricted_review_by: Option<Option<String>>,
    #[serde(rename = "reviews_count", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub reviews_count: Option<Option<i32>>,
    #[serde(rename = "title", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub title: Option<Option<String>>,
    #[serde(rename = "total_gifts_received_count", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub total_gifts_received_count: Option<Option<i32>>,
    #[serde(rename = "twitter_id", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub twitter_id: Option<Option<String>>,
    #[serde(rename = "two_fa_enabled", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub two_fa_enabled: Option<Option<bool>>,
    #[serde(rename = "updated_time_millis", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub updated_time_millis: Option<Option<i64>>,
    #[serde(rename = "username", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub username: Option<Option<String>>,
    #[serde(rename = "uuid", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub uuid: Option<Option<String>>,
    #[serde(rename = "video_on", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub video_on: Option<Option<bool>>,
    #[serde(rename = "vip", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub vip: Option<Option<bool>>,
    #[serde(rename = "vip_until_seconds", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub vip_until_seconds: Option<Option<i64>>,
    #[serde(rename = "world_id_connected", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub world_id_connected: Option<Option<bool>>,
}

impl RealmUser {
    pub fn new() -> RealmUser {
        RealmUser {
            age_verified: None,
            apple_connected: None,
            biography: None,
            birth_date: None,
            blocking_limit: None,
            chat_request: None,
            connected_by: None,
            contact_phones: None,
            country_code: None,
            cover_image: None,
            cover_image_thumbnail: None,
            created_at_seconds: None,
            dangerous_user: None,
            email_confirmed: None,
            facebook_connected: None,
            follow_pending: None,
            followed_by: None,
            followed_by_pending: None,
            followers_count: None,
            following: None,
            followings_count: None,
            frame: None,
            frame_thumbnail: None,
            from_different_generation_and_trusted: None,
            gender: None,
            generation: None,
            gifting_ability: None,
            google_connected: None,
            group_phone_on: None,
            group_user: None,
            group_video_on: None,
            groups_users_count: None,
            hidden: None,
            hide_vip: None,
            id: None,
            interests_count: None,
            interests_selected: None,
            is_private: None,
            last_logged_in_at_seconds: None,
            line_connected: None,
            login_streak_count: None,
            masked_email: None,
            media_count: None,
            mobile_number: None,
            new_user: None,
            nickname: None,
            online_status: None,
            phone_on: None,
            posts_count: None,
            prefecture: None,
            profile_icon: None,
            profile_icon_thumbnail: None,
            push_notification: None,
            recently_kenta: None,
            restricted_review_by: None,
            reviews_count: None,
            title: None,
            total_gifts_received_count: None,
            twitter_id: None,
            two_fa_enabled: None,
            updated_time_millis: None,
            username: None,
            uuid: None,
            video_on: None,
            vip: None,
            vip_until_seconds: None,
            world_id_connected: None,
        }
    }
}

