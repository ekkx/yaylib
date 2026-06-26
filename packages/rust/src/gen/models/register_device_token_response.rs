// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct RegisterDeviceTokenResponse {
    #[serde(rename = "created_at", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub created_at: Option<Option<i64>>,
    #[serde(rename = "id", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub id: Option<Option<i64>>,
    #[serde(rename = "server_device_id", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub server_device_id: Option<Option<String>>,
    #[serde(rename = "updated_at", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub updated_at: Option<Option<i64>>,
    #[serde(rename = "uuid", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub uuid: Option<Option<String>>,
}

impl RegisterDeviceTokenResponse {
    pub fn new() -> RegisterDeviceTokenResponse {
        RegisterDeviceTokenResponse {
            created_at: None,
            id: None,
            server_device_id: None,
            updated_at: None,
            uuid: None,
        }
    }
}

