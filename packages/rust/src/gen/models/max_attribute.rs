// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct MaxAttribute {
    #[serde(rename = "luck", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub luck: Option<Option<i32>>,
    #[serde(rename = "power", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub power: Option<Option<i32>>,
    #[serde(rename = "speed", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub speed: Option<Option<i32>>,
    #[serde(rename = "stamina", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub stamina: Option<Option<i32>>,
    #[serde(rename = "technique", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub technique: Option<Option<i32>>,
}

impl MaxAttribute {
    pub fn new() -> MaxAttribute {
        MaxAttribute {
            luck: None,
            power: None,
            speed: None,
            stamina: None,
            technique: None,
        }
    }
}

