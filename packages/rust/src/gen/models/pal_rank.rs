// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct PalRank {
    #[serde(rename = "own", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub own: Option<Option<bool>>,
    #[serde(rename = "pal", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub pal: Option<Option<models::Pal>>,
    #[serde(rename = "place", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub place: Option<Option<i32>>,
    #[serde(rename = "yay_user_name", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub yay_user_name: Option<Option<String>>,
}

impl PalRank {
    pub fn new() -> PalRank {
        PalRank {
            own: None,
            pal: None,
            place: None,
            yay_user_name: None,
        }
    }
}

