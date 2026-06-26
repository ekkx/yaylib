// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct MissionActionX {
    #[serde(rename = "description_id", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub description_id: Option<Option<i32>>,
    #[serde(rename = "redirect_button_id", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub redirect_button_id: Option<Option<i32>>,
    #[serde(rename = "reward_details_id", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub reward_details_id: Option<Option<i32>>,
    #[serde(rename = "title_id", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub title_id: Option<Option<i32>>,
}

impl MissionActionX {
    pub fn new() -> MissionActionX {
        MissionActionX {
            description_id: None,
            redirect_button_id: None,
            reward_details_id: None,
            title_id: None,
        }
    }
}

