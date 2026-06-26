// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct GiftingAbilitiesResponse {
    #[serde(rename = "gifting_abilities", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub gifting_abilities: Option<Option<Vec<models::RealmGiftingAbility>>>,
}

impl GiftingAbilitiesResponse {
    pub fn new() -> GiftingAbilitiesResponse {
        GiftingAbilitiesResponse {
            gifting_abilities: None,
        }
    }
}

