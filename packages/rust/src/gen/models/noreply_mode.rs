// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

/// 
#[derive(Clone, Copy, Debug, Eq, PartialEq, Ord, PartialOrd, Hash, Serialize, Deserialize)]
pub enum NoreplyMode {
    #[serde(rename = "")]
    False,
    #[serde(rename = "noreply_")]
    True,

    #[serde(other)]
    Unknown,
}

impl std::fmt::Display for NoreplyMode {
    fn fmt(&self, f: &mut std::fmt::Formatter) -> std::fmt::Result {
        match self {
            Self::False => write!(f, ""),
            Self::True => write!(f, "noreply_"),
            Self::Unknown => write!(f, "unknown"),
        }
    }
}

impl Default for NoreplyMode {
    fn default() -> NoreplyMode {
        Self::Unknown
    }
}

