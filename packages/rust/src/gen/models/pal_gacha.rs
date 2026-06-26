// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct PalGacha {
    #[serde(rename = "banner_collection_url", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub banner_collection_url: Option<Option<String>>,
    #[serde(rename = "banner_url", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub banner_url: Option<Option<String>>,
    #[serde(rename = "detail_url", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub detail_url: Option<Option<String>>,
    #[serde(rename = "gacha_type", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub gacha_type: Option<Option<String>>,
    #[serde(rename = "id", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub id: Option<Option<String>>,
    #[serde(rename = "price", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub price: Option<Option<i32>>,
}

impl PalGacha {
    pub fn new() -> PalGacha {
        PalGacha {
            banner_collection_url: None,
            banner_url: None,
            detail_url: None,
            gacha_type: None,
            id: None,
            price: None,
        }
    }
}

