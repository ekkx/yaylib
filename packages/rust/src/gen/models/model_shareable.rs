// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct ModelShareable {
    #[serde(rename = "group", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub group: Option<Option<models::ModelGroup>>,
    #[serde(rename = "post", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub post: Option<Option<Box<models::ModelPost>>>,
    #[serde(rename = "thread", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub thread: Option<Option<models::ModelThreadInfo>>,
}

impl ModelShareable {
    pub fn new() -> ModelShareable {
        ModelShareable {
            group: None,
            post: None,
            thread: None,
        }
    }
}

