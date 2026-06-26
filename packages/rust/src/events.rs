use serde::Deserialize;

#[derive(Debug, Clone)]
pub enum Event {
    NewMessage(Box<NewMessageEvent>),
    VideoProcessed(VideoProcessedEvent),
    ChatDeleted(ChatDeletedEvent),
    TotalChatRequest(TotalChatRequestEvent),
    Unsubscribed(UnsubscribedEvent),
    GroupUpdated(GroupUpdatedEvent),
    CallFinished(CallFinishedEvent),
    Raw(RawEvent),
}

/// Carries the full chat message DTO. Fields are inherited from the
/// generated [`RealmMessage`](crate::gen::models::RealmMessage) model
/// via `#[serde(flatten)]`; `video_thumbnail_big_url` is a WebSocket-only
/// field that does NOT appear on the REST Message DTO.
#[derive(Debug, Clone, Deserialize)]
pub struct NewMessageEvent {
    #[serde(flatten)]
    pub message: crate::gen::models::RealmMessage,
    /// WebSocket-only field; does NOT appear on the REST Message DTO.
    #[serde(default)]
    pub video_thumbnail_big_url: Option<String>,
}

#[derive(Debug, Clone, Deserialize)]
pub struct VideoProcessedEvent {
    pub id: i64,
    #[serde(default)]
    pub video_processed: bool,
    #[serde(default)]
    pub video_url: Option<String>,
    #[serde(default)]
    pub video_thumbnail_url: Option<String>,
    #[serde(default)]
    pub video_thumbnail_big_url: Option<String>,
}

#[derive(Debug, Clone, Deserialize)]
pub struct ChatDeletedEvent {
    pub room_id: i64,
}

#[derive(Debug, Clone, Deserialize)]
pub struct TotalChatRequestEvent {
    pub total_count: i32,
}

#[derive(Debug, Clone, Deserialize)]
pub struct UnsubscribedEvent {
    #[serde(default)]
    pub user_ids: Vec<i64>,
}

#[derive(Debug, Clone, Deserialize)]
pub struct GroupUpdatedEvent {
    pub group_id: i64,
}

#[derive(Debug, Clone, Deserialize)]
pub struct CallFinishedEvent {
    pub id: i64,
}

#[derive(Debug, Clone)]
pub struct RawEvent {
    pub name: String,
    pub data: serde_json::Value,
}

pub(crate) fn decode_event(name: &str, data: &[u8]) -> Option<Event> {
    match name {
        "new_message" => serde_json::from_slice::<NewMessageEvent>(data)
            .ok()
            .map(|e| Event::NewMessage(Box::new(e))),
        "video_processed" => serde_json::from_slice::<VideoProcessedEvent>(data)
            .ok()
            .map(Event::VideoProcessed),
        "chat_deleted" => serde_json::from_slice::<ChatDeletedEvent>(data)
            .ok()
            .map(Event::ChatDeleted),
        "total_chat_request" => serde_json::from_slice::<TotalChatRequestEvent>(data)
            .ok()
            .map(Event::TotalChatRequest),
        "unsubscribed" => serde_json::from_slice::<UnsubscribedEvent>(data)
            .ok()
            .map(Event::Unsubscribed),
        "new_post" => serde_json::from_slice::<GroupUpdatedEvent>(data)
            .ok()
            .map(Event::GroupUpdated),
        "conference_call_finished" => serde_json::from_slice::<CallFinishedEvent>(data)
            .ok()
            .map(Event::CallFinished),
        _ => serde_json::from_slice::<serde_json::Value>(data)
            .ok()
            .map(|v| {
                Event::Raw(RawEvent {
                    name: name.to_string(),
                    data: v,
                })
            }),
    }
}
