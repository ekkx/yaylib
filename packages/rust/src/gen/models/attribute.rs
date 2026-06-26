// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct Attribute {
    #[serde(rename = "birthday", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub birthday: Option<Option<String>>,
    #[serde(rename = "body", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub body: Option<Option<String>>,
    #[serde(rename = "climbing", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub climbing: Option<Option<String>>,
    #[serde(rename = "eligible_league", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub eligible_league: Option<Option<i32>>,
    #[serde(rename = "fatigue", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub fatigue: Option<Option<f64>>,
    #[serde(rename = "generation", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub generation: Option<Option<i32>>,
    #[serde(rename = "grade", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub grade: Option<Option<String>>,
    #[serde(rename = "head", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub head: Option<Option<String>>,
    #[serde(rename = "item", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub item: Option<Option<String>>,
    #[serde(rename = "level", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub level: Option<Option<i32>>,
    #[serde(rename = "luck", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub luck: Option<Option<i32>>,
    #[serde(rename = "momentum", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub momentum: Option<Option<i32>>,
    #[serde(rename = "name", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub name: Option<Option<String>>,
    #[serde(rename = "power", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub power: Option<Option<i32>>,
    #[serde(rename = "racingstyle", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub racingstyle: Option<Option<String>>,
    #[serde(rename = "rank", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub rank: Option<Option<String>>,
    #[serde(rename = "running", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub running: Option<Option<String>>,
    #[serde(rename = "skin", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub skin: Option<Option<String>>,
    #[serde(rename = "speed", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub speed: Option<Option<i32>>,
    #[serde(rename = "stamina", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub stamina: Option<Option<i32>>,
    #[serde(rename = "status", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub status: Option<Option<String>>,
    #[serde(rename = "swimming", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub swimming: Option<Option<String>>,
    #[serde(rename = "technique", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub technique: Option<Option<i32>>,
    #[serde(rename = "tribe", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub tribe: Option<Option<String>>,
    #[serde(rename = "type", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub r#type: Option<Option<String>>,
}

impl Attribute {
    pub fn new() -> Attribute {
        Attribute {
            birthday: None,
            body: None,
            climbing: None,
            eligible_league: None,
            fatigue: None,
            generation: None,
            grade: None,
            head: None,
            item: None,
            level: None,
            luck: None,
            momentum: None,
            name: None,
            power: None,
            racingstyle: None,
            rank: None,
            running: None,
            skin: None,
            speed: None,
            stamina: None,
            status: None,
            swimming: None,
            technique: None,
            tribe: None,
            r#type: None,
        }
    }
}

