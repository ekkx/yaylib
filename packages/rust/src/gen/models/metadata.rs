// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct Metadata {
    #[serde(rename = "anonymous", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub anonymous: Option<Option<bool>>,
    #[serde(rename = "body", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub body: Option<Option<String>>,
    #[serde(rename = "bulk_invitation", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub bulk_invitation: Option<Option<bool>>,
    #[serde(rename = "content_preview", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub content_preview: Option<Option<String>>,
    #[serde(rename = "group_topic", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub group_topic: Option<Option<String>>,
    #[serde(rename = "title", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub title: Option<Option<String>>,
    #[serde(rename = "url", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub url: Option<Option<String>>,
}

impl Metadata {
    pub fn new() -> Metadata {
        Metadata {
            anonymous: None,
            body: None,
            bulk_invitation: None,
            content_preview: None,
            group_topic: None,
            title: None,
            url: None,
        }
    }
}

