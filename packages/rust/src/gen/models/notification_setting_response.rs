// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct NotificationSettingResponse {
    #[serde(rename = "setting", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub setting: Option<Option<models::UserSetting>>,
}

impl NotificationSettingResponse {
    pub fn new() -> NotificationSettingResponse {
        NotificationSettingResponse {
            setting: None,
        }
    }
}

