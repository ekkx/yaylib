// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct ConferenceCallResponse {
    #[serde(rename = "conference_call", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub conference_call: Option<Option<models::RealmConferenceCall>>,
    #[serde(rename = "conference_call_user_uuid", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub conference_call_user_uuid: Option<Option<String>>,
}

impl ConferenceCallResponse {
    pub fn new() -> ConferenceCallResponse {
        ConferenceCallResponse {
            conference_call: None,
            conference_call_user_uuid: None,
        }
    }
}

