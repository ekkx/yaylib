// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct PolicyAgreementsResponse {
    #[serde(rename = "latest_dotmoney_agreed", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub latest_dotmoney_agreed: Option<Option<bool>>,
    #[serde(rename = "latest_empl2_agreed", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub latest_empl2_agreed: Option<Option<bool>>,
    #[serde(rename = "latest_privacy_policy_agreed", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub latest_privacy_policy_agreed: Option<Option<bool>>,
    #[serde(rename = "latest_terms_of_use_agreed", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub latest_terms_of_use_agreed: Option<Option<bool>>,
}

impl PolicyAgreementsResponse {
    pub fn new() -> PolicyAgreementsResponse {
        PolicyAgreementsResponse {
            latest_dotmoney_agreed: None,
            latest_empl2_agreed: None,
            latest_privacy_policy_agreed: None,
            latest_terms_of_use_agreed: None,
        }
    }
}

