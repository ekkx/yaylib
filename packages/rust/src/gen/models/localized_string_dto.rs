// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct LocalizedStringDto {
    #[serde(rename = "en", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub en: Option<Option<String>>,
    #[serde(rename = "id", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub id: Option<Option<String>>,
    #[serde(rename = "ja", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub ja: Option<Option<String>>,
    #[serde(rename = "ko", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub ko: Option<Option<String>>,
    #[serde(rename = "th", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub th: Option<Option<String>>,
    #[serde(rename = "zh_cn", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub zh_cn: Option<Option<String>>,
    #[serde(rename = "zh_hk", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub zh_hk: Option<Option<String>>,
}

impl LocalizedStringDto {
    pub fn new() -> LocalizedStringDto {
        LocalizedStringDto {
            en: None,
            id: None,
            ja: None,
            ko: None,
            th: None,
            zh_cn: None,
            zh_hk: None,
        }
    }
}

