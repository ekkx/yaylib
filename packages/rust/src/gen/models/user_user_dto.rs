// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct UserUserDto {
    #[serde(rename = "age_verified", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub age_verified: Option<Option<bool>>,
    #[serde(rename = "biography", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub biography: Option<Option<String>>,
    #[serde(rename = "cover_image", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub cover_image: Option<Option<String>>,
    #[serde(rename = "cover_image_thumbnail", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub cover_image_thumbnail: Option<Option<String>>,
    #[serde(rename = "created_at", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub created_at: Option<Option<i64>>,
    #[serde(rename = "dangerous_user", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub dangerous_user: Option<Option<bool>>,
    #[serde(rename = "follow_pending", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub follow_pending: Option<Option<bool>>,
    #[serde(rename = "followed_by", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub followed_by: Option<Option<bool>>,
    #[serde(rename = "followers_count", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub followers_count: Option<Option<i32>>,
    #[serde(rename = "following", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub following: Option<Option<bool>>,
    #[serde(rename = "followings_count", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub followings_count: Option<Option<i32>>,
    #[serde(rename = "gender", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub gender: Option<Option<i32>>,
    #[serde(rename = "groups_users_count", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub groups_users_count: Option<Option<i32>>,
    #[serde(rename = "hidden", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub hidden: Option<Option<bool>>,
    #[serde(rename = "hide_vip", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub hide_vip: Option<Option<bool>>,
    #[serde(rename = "id", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub id: Option<Option<i64>>,
    #[serde(rename = "is_private", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub is_private: Option<Option<bool>>,
    #[serde(rename = "new_user", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub new_user: Option<Option<bool>>,
    #[serde(rename = "nickname", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub nickname: Option<Option<String>>,
    #[serde(rename = "posts_count", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub posts_count: Option<Option<i32>>,
    #[serde(rename = "prefecture", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub prefecture: Option<Option<String>>,
    #[serde(rename = "profile_icon", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub profile_icon: Option<Option<String>>,
    #[serde(rename = "profile_icon_thumbnail", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub profile_icon_thumbnail: Option<Option<String>>,
    #[serde(rename = "recently_kenta", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub recently_kenta: Option<Option<bool>>,
    #[serde(rename = "reviews_count", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub reviews_count: Option<Option<i32>>,
    #[serde(rename = "title", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub title: Option<Option<String>>,
    #[serde(rename = "vip", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub vip: Option<Option<bool>>,
}

impl UserUserDto {
    pub fn new() -> UserUserDto {
        UserUserDto {
            age_verified: None,
            biography: None,
            cover_image: None,
            cover_image_thumbnail: None,
            created_at: None,
            dangerous_user: None,
            follow_pending: None,
            followed_by: None,
            followers_count: None,
            following: None,
            followings_count: None,
            gender: None,
            groups_users_count: None,
            hidden: None,
            hide_vip: None,
            id: None,
            is_private: None,
            new_user: None,
            nickname: None,
            posts_count: None,
            prefecture: None,
            profile_icon: None,
            profile_icon_thumbnail: None,
            recently_kenta: None,
            reviews_count: None,
            title: None,
            vip: None,
        }
    }
}

