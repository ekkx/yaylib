// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct StateChanges {
    #[serde(rename = "attack", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub attack: Option<Option<i32>>,
    #[serde(rename = "defense", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub defense: Option<Option<i32>>,
    #[serde(rename = "energy", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub energy: Option<Option<i32>>,
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

impl StateChanges {
    pub fn new() -> StateChanges {
        StateChanges {
            attack: None,
            defense: None,
            energy: None,
            luck: None,
            power: None,
            speed: None,
            stamina: None,
            technique: None,
        }
    }
}

