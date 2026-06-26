// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

/// 
#[derive(Clone, Copy, Debug, Eq, PartialEq, Ord, PartialOrd, Hash, Serialize, Deserialize)]
pub enum J0A {
    #[serde(rename = "PublicTimeline")]
    PublicTimeline,
    #[serde(rename = "FollowingTimeline")]
    FollowingTimeline,
    #[serde(rename = "GroupTimeline")]
    GroupTimeline,

    #[serde(other)]
    Unknown,
}

impl std::fmt::Display for J0A {
    fn fmt(&self, f: &mut std::fmt::Formatter) -> std::fmt::Result {
        match self {
            Self::PublicTimeline => write!(f, "PublicTimeline"),
            Self::FollowingTimeline => write!(f, "FollowingTimeline"),
            Self::GroupTimeline => write!(f, "GroupTimeline"),
            Self::Unknown => write!(f, "unknown"),
        }
    }
}

impl Default for J0A {
    fn default() -> J0A {
        Self::Unknown
    }
}

