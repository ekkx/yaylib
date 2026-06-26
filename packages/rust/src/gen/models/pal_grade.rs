// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

/// 
#[derive(Clone, Copy, Debug, Eq, PartialEq, Ord, PartialOrd, Hash, Serialize, Deserialize)]
pub enum PalGrade {
    #[serde(rename = "Egg")]
    Egg,
    #[serde(rename = "Pal")]
    Normal,
    #[serde(rename = "Super Pal")]
    Super,
    #[serde(rename = "Ultimate Pal")]
    UltimatePal,

    #[serde(other)]
    Unknown,
}

impl std::fmt::Display for PalGrade {
    fn fmt(&self, f: &mut std::fmt::Formatter) -> std::fmt::Result {
        match self {
            Self::Egg => write!(f, "Egg"),
            Self::Normal => write!(f, "Pal"),
            Self::Super => write!(f, "Super Pal"),
            Self::UltimatePal => write!(f, "Ultimate Pal"),
            Self::Unknown => write!(f, "unknown"),
        }
    }
}

impl Default for PalGrade {
    fn default() -> PalGrade {
        Self::Unknown
    }
}

