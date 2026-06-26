// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct Shareable {
    #[serde(rename = "group", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub group: Option<Option<models::Group>>,
    #[serde(rename = "post", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub post: Option<Option<Box<models::Post>>>,
    #[serde(rename = "thread", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub thread: Option<Option<models::ThreadInfo>>,
}

impl Shareable {
    pub fn new() -> Shareable {
        Shareable {
            group: None,
            post: None,
            thread: None,
        }
    }
}

