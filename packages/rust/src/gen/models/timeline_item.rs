// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct TimelineItem {
    #[serde(rename = "categories", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub categories: Option<Option<Vec<models::GroupCategory>>>,
    #[serde(rename = "is_replied_by_next_post", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub is_replied_by_next_post: Option<Option<bool>>,
    #[serde(rename = "is_replying_to_prev_post", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub is_replying_to_prev_post: Option<Option<bool>>,
    #[serde(rename = "post", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub post: Option<Option<models::ModelPost>>,
    #[serde(rename = "should_show_conversation_line", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub should_show_conversation_line: Option<Option<bool>>,
    #[serde(rename = "should_show_mention_box", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub should_show_mention_box: Option<Option<bool>>,
    #[serde(rename = "should_show_repost_button", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub should_show_repost_button: Option<Option<bool>>,
    #[serde(rename = "should_show_translate_button", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub should_show_translate_button: Option<Option<bool>>,
    #[serde(rename = "should_show_view_reposts", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub should_show_view_reposts: Option<Option<bool>>,
    #[serde(rename = "translated_text", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub translated_text: Option<Option<String>>,
}

impl TimelineItem {
    pub fn new() -> TimelineItem {
        TimelineItem {
            categories: None,
            is_replied_by_next_post: None,
            is_replying_to_prev_post: None,
            post: None,
            should_show_conversation_line: None,
            should_show_mention_box: None,
            should_show_repost_button: None,
            should_show_translate_button: None,
            should_show_view_reposts: None,
            translated_text: None,
        }
    }
}

