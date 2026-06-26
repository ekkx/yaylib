// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct TimelineSettings {
    #[serde(rename = "faves_filter", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub faves_filter: Option<Option<bool>>,
    #[serde(rename = "hide_hot_post", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub hide_hot_post: Option<Option<bool>>,
    #[serde(rename = "hide_reply_following_timeline", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub hide_reply_following_timeline: Option<Option<bool>>,
    #[serde(rename = "hide_reply_public_timeline", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub hide_reply_public_timeline: Option<Option<bool>>,
}

impl TimelineSettings {
    pub fn new() -> TimelineSettings {
        TimelineSettings {
            faves_filter: None,
            hide_hot_post: None,
            hide_reply_following_timeline: None,
            hide_reply_public_timeline: None,
        }
    }
}

