// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

/// 
#[derive(Clone, Copy, Debug, Eq, PartialEq, Ord, PartialOrd, Hash, Serialize, Deserialize)]
pub enum JoinableBy {
    #[serde(rename = "anyone")]
    Anyone,
    #[serde(rename = "users")]
    Users,
    #[serde(rename = "followers_only")]
    FollowersOnly,
    #[serde(rename = "mutual_follows_only")]
    MutualFollowsOnly,
    #[serde(rename = "invite_only")]
    InviteOnly,
    #[serde(rename = "noone")]
    NoOne,

    #[serde(other)]
    Unknown,
}

impl std::fmt::Display for JoinableBy {
    fn fmt(&self, f: &mut std::fmt::Formatter) -> std::fmt::Result {
        match self {
            Self::Anyone => write!(f, "anyone"),
            Self::Users => write!(f, "users"),
            Self::FollowersOnly => write!(f, "followers_only"),
            Self::MutualFollowsOnly => write!(f, "mutual_follows_only"),
            Self::InviteOnly => write!(f, "invite_only"),
            Self::NoOne => write!(f, "noone"),
            Self::Unknown => write!(f, "unknown"),
        }
    }
}

impl Default for JoinableBy {
    fn default() -> JoinableBy {
        Self::Unknown
    }
}

