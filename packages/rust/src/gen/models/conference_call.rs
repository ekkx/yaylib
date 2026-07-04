// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct ConferenceCall {
    #[serde(rename = "agora_channel", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub agora_channel: Option<Option<String>>,
    #[serde(rename = "agora_token", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub agora_token: Option<Option<String>>,
    #[serde(rename = "anonymous_call_users_count", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub anonymous_call_users_count: Option<Option<i32>>,
    #[serde(rename = "bump_params", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub bump_params: Option<Option<models::BumpParams>>,
    #[serde(rename = "call_type", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub call_type: Option<Option<models::CallType>>,
    #[serde(rename = "conference_call_user_uuid", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub conference_call_user_uuid: Option<Option<String>>,
    #[serde(rename = "conference_call_users", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub conference_call_users: Option<Option<Vec<models::User>>>,
    #[serde(rename = "conference_call_users_role", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub conference_call_users_role: Option<Option<Vec<models::ModelConferenceCallUserRole>>>,
    #[serde(rename = "duration_seconds", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub duration_seconds: Option<Option<i64>>,
    #[serde(rename = "game", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub game: Option<Option<models::Game>>,
    #[serde(rename = "genre", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub genre: Option<Option<models::Genre>>,
    #[serde(rename = "group_id", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub group_id: Option<Option<i64>>,
    #[serde(rename = "id", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub id: Option<Option<i64>>,
    #[serde(rename = "is_active", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub is_active: Option<Option<bool>>,
    #[serde(rename = "joinable_by", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub joinable_by: Option<Option<models::JoinableBy>>,
    #[serde(rename = "max_participants", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub max_participants: Option<Option<i32>>,
    #[serde(rename = "post_id", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub post_id: Option<Option<i64>>,
}

impl ConferenceCall {
    pub fn new() -> ConferenceCall {
        ConferenceCall {
            agora_channel: None,
            agora_token: None,
            anonymous_call_users_count: None,
            bump_params: None,
            call_type: None,
            conference_call_user_uuid: None,
            conference_call_users: None,
            conference_call_users_role: None,
            duration_seconds: None,
            game: None,
            genre: None,
            group_id: None,
            id: None,
            is_active: None,
            joinable_by: None,
            max_participants: None,
            post_id: None,
        }
    }
}

