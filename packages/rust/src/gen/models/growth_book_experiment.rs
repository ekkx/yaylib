// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct GrowthBookExperiment {
    #[serde(rename = "experiment_number", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub experiment_number: Option<Option<i32>>,
    #[serde(rename = "number", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub number: Option<Option<i32>>,
    #[serde(rename = "variant_number", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub variant_number: Option<Option<i32>>,
    #[serde(rename = "visible", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub visible: Option<Option<bool>>,
}

impl GrowthBookExperiment {
    pub fn new() -> GrowthBookExperiment {
        GrowthBookExperiment {
            experiment_number: None,
            number: None,
            variant_number: None,
            visible: None,
        }
    }
}

