// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct RealmGame {
    #[serde(rename = "icon_url", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub icon_url: Option<Option<String>>,
    #[serde(rename = "id", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub id: Option<Option<i64>>,
    #[serde(rename = "platform_details", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub platform_details: Option<Option<models::RealmPlatformDetails>>,
    #[serde(rename = "title", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub title: Option<Option<String>>,
}

impl RealmGame {
    pub fn new() -> RealmGame {
        RealmGame {
            icon_url: None,
            id: None,
            platform_details: None,
            title: None,
        }
    }
}

