// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct PalActivityLogDetails {
    #[serde(rename = "attacking", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub attacking: Option<Option<i32>>,
    #[serde(rename = "cost", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub cost: Option<Option<f64>>,
    #[serde(rename = "gacha_type", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub gacha_type: Option<Option<String>>,
    #[serde(rename = "item_status", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub item_status: Option<Option<String>>,
    #[serde(rename = "item_subtype", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub item_subtype: Option<Option<String>>,
    #[serde(rename = "item_type", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub item_type: Option<Option<String>>,
    #[serde(rename = "level", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub level: Option<Option<i32>>,
    #[serde(rename = "nft_type", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub nft_type: Option<Option<String>>,
    #[serde(rename = "outcome", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub outcome: Option<Option<i64>>,
    #[serde(rename = "pal_grade", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub pal_grade: Option<Option<models::PalGrade>>,
    #[serde(rename = "pal_image", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub pal_image: Option<Option<String>>,
    #[serde(rename = "pal_name", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub pal_name: Option<Option<String>>,
    #[serde(rename = "places", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub places: Option<Option<Vec<i32>>>,
    #[serde(rename = "race_id", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub race_id: Option<Option<String>>,
    #[serde(rename = "rank", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub rank: Option<Option<String>>,
    #[serde(rename = "result", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub result: Option<Option<String>>,
    #[serde(rename = "token_id", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub token_id: Option<Option<String>>,
    #[serde(rename = "transaction_code", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub transaction_code: Option<Option<String>>,
}

impl PalActivityLogDetails {
    pub fn new() -> PalActivityLogDetails {
        PalActivityLogDetails {
            attacking: None,
            cost: None,
            gacha_type: None,
            item_status: None,
            item_subtype: None,
            item_type: None,
            level: None,
            nft_type: None,
            outcome: None,
            pal_grade: None,
            pal_image: None,
            pal_name: None,
            places: None,
            race_id: None,
            rank: None,
            result: None,
            token_id: None,
            transaction_code: None,
        }
    }
}

