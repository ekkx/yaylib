// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

/// 
#[derive(Clone, Copy, Debug, Eq, PartialEq, Ord, PartialOrd, Hash, Serialize, Deserialize)]
pub enum MissionType {
    #[serde(rename = "daily")]
    Daily,
    #[serde(rename = "special")]
    Special,
    #[serde(rename = "unlimited")]
    Endless,
    #[serde(rename = "once")]
    Normal,

    #[serde(other)]
    Unknown,
}

impl std::fmt::Display for MissionType {
    fn fmt(&self, f: &mut std::fmt::Formatter) -> std::fmt::Result {
        match self {
            Self::Daily => write!(f, "daily"),
            Self::Special => write!(f, "special"),
            Self::Endless => write!(f, "unlimited"),
            Self::Normal => write!(f, "once"),
            Self::Unknown => write!(f, "unknown"),
        }
    }
}

impl Default for MissionType {
    fn default() -> MissionType {
        Self::Unknown
    }
}

