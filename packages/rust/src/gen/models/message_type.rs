// Code generated from the Yay! API schema; DO NOT EDIT.

use crate::gen::models;
use serde::{Deserialize, Serialize};

/// 
#[derive(Clone, Copy, Debug, Eq, PartialEq, Ord, PartialOrd, Hash, Serialize, Deserialize)]
pub enum MessageType {
    #[serde(rename = "text")]
    Text,
    #[serde(rename = "image")]
    Image,
    #[serde(rename = "eternal_image")]
    EternalImage,
    #[serde(rename = "video")]
    Video,
    #[serde(rename = "eternal_video")]
    EternalVideo,
    #[serde(rename = "screenshot_warning")]
    Screenshot,
    #[serde(rename = "gif")]
    Gif,
    #[serde(rename = "sticker")]
    Sticker,
    #[serde(rename = "individual_call")]
    PrivateCall,
    #[serde(rename = "deleted")]
    Deleted,
    #[serde(rename = "user_joined")]
    UserJoined,
    #[serde(rename = "user_leave")]
    UserLeave,
    #[serde(rename = "owner_kick")]
    OwnerKick,

    #[serde(other)]
    Unknown,
}

impl std::fmt::Display for MessageType {
    fn fmt(&self, f: &mut std::fmt::Formatter) -> std::fmt::Result {
        match self {
            Self::Text => write!(f, "text"),
            Self::Image => write!(f, "image"),
            Self::EternalImage => write!(f, "eternal_image"),
            Self::Video => write!(f, "video"),
            Self::EternalVideo => write!(f, "eternal_video"),
            Self::Screenshot => write!(f, "screenshot_warning"),
            Self::Gif => write!(f, "gif"),
            Self::Sticker => write!(f, "sticker"),
            Self::PrivateCall => write!(f, "individual_call"),
            Self::Deleted => write!(f, "deleted"),
            Self::UserJoined => write!(f, "user_joined"),
            Self::UserLeave => write!(f, "user_leave"),
            Self::OwnerKick => write!(f, "owner_kick"),
            Self::Unknown => write!(f, "unknown"),
        }
    }
}

impl Default for MessageType {
    fn default() -> MessageType {
        Self::Unknown
    }
}

