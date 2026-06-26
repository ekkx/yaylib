// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct DAppDto {
    #[serde(rename = "description_localized", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub description_localized: Option<Option<models::LocalizedStringDto>>,
    #[serde(rename = "enabled", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub enabled: Option<Option<bool>>,
    #[serde(rename = "icon_url", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub icon_url: Option<Option<String>>,
    #[serde(rename = "id", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub id: Option<Option<String>>,
    #[serde(rename = "order", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub order: Option<Option<i32>>,
    #[serde(rename = "search_keywords", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub search_keywords: Option<Option<Vec<String>>>,
    #[serde(rename = "show_in_recommend_d_apps", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub show_in_recommend_d_apps: Option<Option<bool>>,
    #[serde(rename = "title_localized", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub title_localized: Option<Option<models::LocalizedStringDto>>,
    #[serde(rename = "url_localized", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub url_localized: Option<Option<models::LocalizedStringDto>>,
}

impl DAppDto {
    pub fn new() -> DAppDto {
        DAppDto {
            description_localized: None,
            enabled: None,
            icon_url: None,
            id: None,
            order: None,
            search_keywords: None,
            show_in_recommend_d_apps: None,
            title_localized: None,
            url_localized: None,
        }
    }
}

