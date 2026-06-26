// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct RealmGift {
    #[serde(rename = "currency", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub currency: Option<Option<String>>,
    #[serde(rename = "icon", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub icon: Option<Option<String>>,
    #[serde(rename = "icon_thumbnail", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub icon_thumbnail: Option<Option<String>>,
    #[serde(rename = "id", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub id: Option<Option<i64>>,
    #[serde(rename = "price", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub price: Option<Option<f64>>,
    #[serde(rename = "slug", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub slug: Option<Option<String>>,
    #[serde(rename = "title", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub title: Option<Option<String>>,
}

impl RealmGift {
    pub fn new() -> RealmGift {
        RealmGift {
            currency: None,
            icon: None,
            icon_thumbnail: None,
            id: None,
            price: None,
            slug: None,
            title: None,
        }
    }
}

