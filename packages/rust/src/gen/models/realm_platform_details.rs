// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct RealmPlatformDetails {
    #[serde(rename = "affiliate_url", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub affiliate_url: Option<Option<String>>,
    #[serde(rename = "package_id", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub package_id: Option<Option<String>>,
}

impl RealmPlatformDetails {
    pub fn new() -> RealmPlatformDetails {
        RealmPlatformDetails {
            affiliate_url: None,
            package_id: None,
        }
    }
}

