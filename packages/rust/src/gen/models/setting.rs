// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct Setting {
    #[serde(rename = "notification_group_join", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub notification_group_join: Option<Option<bool>>,
    #[serde(rename = "notification_group_message_tag_all", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub notification_group_message_tag_all: Option<Option<bool>>,
    #[serde(rename = "notification_group_post", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub notification_group_post: Option<Option<bool>>,
    #[serde(rename = "notification_group_request", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub notification_group_request: Option<Option<bool>>,
}

impl Setting {
    pub fn new() -> Setting {
        Setting {
            notification_group_join: None,
            notification_group_message_tag_all: None,
            notification_group_post: None,
            notification_group_request: None,
        }
    }
}

