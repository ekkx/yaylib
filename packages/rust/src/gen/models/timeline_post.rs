// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct TimelinePost {
    #[serde(rename = "posts", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub posts: Option<Option<Vec<models::ModelPost>>>,
    #[serde(rename = "root_posts", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub root_posts: Option<Option<Vec<models::ModelPost>>>,
}

impl TimelinePost {
    pub fn new() -> TimelinePost {
        TimelinePost {
            posts: None,
            root_posts: None,
        }
    }
}

