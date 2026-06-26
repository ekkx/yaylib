// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

/// 
#[derive(Clone, Copy, Debug, Eq, PartialEq, Ord, PartialOrd, Hash, Serialize, Deserialize)]
pub enum PostType {
    #[serde(rename = "call")]
    Call,
    #[serde(rename = "image")]
    Image,
    #[serde(rename = "video")]
    Video,
    #[serde(rename = "survey")]
    Survey,
    #[serde(rename = "repost")]
    Repost,
    #[serde(rename = "thread")]
    Thread,
    #[serde(rename = "shareable_group")]
    ShareableGroup,
    #[serde(rename = "shareable_url")]
    ShareableUrl,
    #[serde(rename = "youtube")]
    Youtube,
    #[serde(rename = "shareable_thread")]
    ShareableThread,
    #[serde(rename = "sharepal")]
    SharePal,
    #[serde(rename = "undefined")]
    Undefined,
    #[serde(rename = "text")]
    Text,

    #[serde(other)]
    Unknown,
}

impl std::fmt::Display for PostType {
    fn fmt(&self, f: &mut std::fmt::Formatter) -> std::fmt::Result {
        match self {
            Self::Call => write!(f, "call"),
            Self::Image => write!(f, "image"),
            Self::Video => write!(f, "video"),
            Self::Survey => write!(f, "survey"),
            Self::Repost => write!(f, "repost"),
            Self::Thread => write!(f, "thread"),
            Self::ShareableGroup => write!(f, "shareable_group"),
            Self::ShareableUrl => write!(f, "shareable_url"),
            Self::Youtube => write!(f, "youtube"),
            Self::ShareableThread => write!(f, "shareable_thread"),
            Self::SharePal => write!(f, "sharepal"),
            Self::Undefined => write!(f, "undefined"),
            Self::Text => write!(f, "text"),
            Self::Unknown => write!(f, "unknown"),
        }
    }
}

impl Default for PostType {
    fn default() -> PostType {
        Self::Unknown
    }
}

