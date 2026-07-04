// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

/// 
#[derive(Clone, Copy, Debug, Eq, PartialEq, Ord, PartialOrd, Hash, Serialize, Deserialize)]
pub enum CallType {
    #[serde(rename = "vdo")]
    Video,
    #[serde(rename = "voice")]
    Voice,

    #[serde(other)]
    Unknown,
}

impl std::fmt::Display for CallType {
    fn fmt(&self, f: &mut std::fmt::Formatter) -> std::fmt::Result {
        match self {
            Self::Video => write!(f, "vdo"),
            Self::Voice => write!(f, "voice"),
            Self::Unknown => write!(f, "unknown"),
        }
    }
}

impl Default for CallType {
    fn default() -> CallType {
        Self::Unknown
    }
}

