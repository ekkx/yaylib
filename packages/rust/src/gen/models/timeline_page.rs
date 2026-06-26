// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct TimelinePage {
    #[serde(rename = "id", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub id: Option<Option<i64>>,
    #[serde(rename = "items", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub items: Option<Option<Vec<serde_json::Value>>>,
    #[serde(rename = "next_page_value", default, skip_serializing_if = "Option::is_none", deserialize_with = "crate::gen::cursor::deserialize_option_string_or_number")]
    pub next_page_value: Option<String>,
    #[serde(rename = "pinned_items", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub pinned_items: Option<Option<Vec<serde_json::Value>>>,
    #[serde(rename = "show_more_hot_posts_button", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub show_more_hot_posts_button: Option<Option<bool>>,
    #[serde(rename = "total_item_count", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub total_item_count: Option<Option<i32>>,
    #[serde(rename = "total_item_limit", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub total_item_limit: Option<Option<i32>>,
}

impl TimelinePage {
    pub fn new() -> TimelinePage {
        TimelinePage {
            id: None,
            items: None,
            next_page_value: None,
            pinned_items: None,
            show_more_hot_posts_button: None,
            total_item_count: None,
            total_item_limit: None,
        }
    }
}

