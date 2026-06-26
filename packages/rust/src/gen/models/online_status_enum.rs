// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

/// 
#[derive(Clone, Copy, Debug, Eq, PartialEq, Ord, PartialOrd, Hash, Serialize, Deserialize)]
pub enum OnlineStatusEnum {
    #[serde(rename = "offline")]
    Offline,
    #[serde(rename = "hidden")]
    Hidden,

    #[serde(other)]
    Unknown,
}

impl std::fmt::Display for OnlineStatusEnum {
    fn fmt(&self, f: &mut std::fmt::Formatter) -> std::fmt::Result {
        match self {
            Self::Offline => write!(f, "offline"),
            Self::Hidden => write!(f, "hidden"),
            Self::Unknown => write!(f, "unknown"),
        }
    }
}

impl Default for OnlineStatusEnum {
    fn default() -> OnlineStatusEnum {
        Self::Unknown
    }
}

