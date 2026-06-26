// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct HatchGacha {
    #[serde(rename = "empl_cost", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub empl_cost: Option<Option<f64>>,
    #[serde(rename = "gacha_type", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub gacha_type: Option<Option<serde_json::Value>>,
    #[serde(rename = "pal_image_url", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub pal_image_url: Option<Option<String>>,
    #[serde(rename = "pal_name", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub pal_name: Option<Option<String>>,
    #[serde(rename = "pal_rank", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub pal_rank: Option<Option<serde_json::Value>>,
    #[serde(rename = "token_id", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub token_id: Option<Option<String>>,
}

impl HatchGacha {
    pub fn new() -> HatchGacha {
        HatchGacha {
            empl_cost: None,
            gacha_type: None,
            pal_image_url: None,
            pal_name: None,
            pal_rank: None,
            token_id: None,
        }
    }
}

