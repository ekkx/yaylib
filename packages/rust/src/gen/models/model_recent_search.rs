// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct ModelRecentSearch {
    #[serde(rename = "hashtag", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub hashtag: Option<Option<models::ModelPostTag>>,
    #[serde(rename = "id", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub id: Option<Option<i64>>,
    #[serde(rename = "keyword", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub keyword: Option<Option<String>>,
    #[serde(rename = "type", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub r#type: Option<Option<models::RecentSearchType>>,
    #[serde(rename = "user", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub user: Option<Option<models::User>>,
}

impl ModelRecentSearch {
    pub fn new() -> ModelRecentSearch {
        ModelRecentSearch {
            hashtag: None,
            id: None,
            keyword: None,
            r#type: None,
            user: None,
        }
    }
}

