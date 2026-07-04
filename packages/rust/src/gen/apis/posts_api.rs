// Code generated from the Yay! API schema; DO NOT EDIT.


use reqwest;
use serde::{Deserialize, Serialize};
use crate::gen::{apis::ResponseContent, models};
use super::{Error, configuration};


/// struct for typed errors of method [`create_conference_call_post`]
#[derive(Debug, Clone, Serialize, Deserialize)]
#[serde(untagged)]
pub enum CreateConferenceCallPostError {
    UnknownValue(serde_json::Value),
}

/// struct for typed errors of method [`create_post`]
#[derive(Debug, Clone, Serialize, Deserialize)]
#[serde(untagged)]
pub enum CreatePostError {
    UnknownValue(serde_json::Value),
}

/// struct for typed errors of method [`create_share_post`]
#[derive(Debug, Clone, Serialize, Deserialize)]
#[serde(untagged)]
pub enum CreateSharePostError {
    UnknownValue(serde_json::Value),
}

/// struct for typed errors of method [`delete_all_posts`]
#[derive(Debug, Clone, Serialize, Deserialize)]
#[serde(untagged)]
pub enum DeleteAllPostsError {
    UnknownValue(serde_json::Value),
}

/// struct for typed errors of method [`delete_posts`]
#[derive(Debug, Clone, Serialize, Deserialize)]
#[serde(untagged)]
pub enum DeletePostsError {
    UnknownValue(serde_json::Value),
}

/// struct for typed errors of method [`get_active_call_post`]
#[derive(Debug, Clone, Serialize, Deserialize)]
#[serde(untagged)]
pub enum GetActiveCallPostError {
    UnknownValue(serde_json::Value),
}

/// struct for typed errors of method [`get_call_followers_timeline`]
#[derive(Debug, Clone, Serialize, Deserialize)]
#[serde(untagged)]
pub enum GetCallFollowersTimelineError {
    UnknownValue(serde_json::Value),
}

/// struct for typed errors of method [`get_call_timeline`]
#[derive(Debug, Clone, Serialize, Deserialize)]
#[serde(untagged)]
pub enum GetCallTimelineError {
    UnknownValue(serde_json::Value),
}

/// struct for typed errors of method [`get_following_timeline`]
#[derive(Debug, Clone, Serialize, Deserialize)]
#[serde(untagged)]
pub enum GetFollowingTimelineError {
    UnknownValue(serde_json::Value),
}

/// struct for typed errors of method [`get_group_calls`]
#[derive(Debug, Clone, Serialize, Deserialize)]
#[serde(untagged)]
pub enum GetGroupCallsError {
    UnknownValue(serde_json::Value),
}

/// struct for typed errors of method [`get_group_timeline`]
#[derive(Debug, Clone, Serialize, Deserialize)]
#[serde(untagged)]
pub enum GetGroupTimelineError {
    UnknownValue(serde_json::Value),
}

/// struct for typed errors of method [`get_my_posts`]
#[derive(Debug, Clone, Serialize, Deserialize)]
#[serde(untagged)]
pub enum GetMyPostsError {
    UnknownValue(serde_json::Value),
}

/// struct for typed errors of method [`get_post`]
#[derive(Debug, Clone, Serialize, Deserialize)]
#[serde(untagged)]
pub enum GetPostError {
    UnknownValue(serde_json::Value),
}

/// struct for typed errors of method [`get_post_gift_transactions`]
#[derive(Debug, Clone, Serialize, Deserialize)]
#[serde(untagged)]
pub enum GetPostGiftTransactionsError {
    UnknownValue(serde_json::Value),
}

/// struct for typed errors of method [`get_post_likers`]
#[derive(Debug, Clone, Serialize, Deserialize)]
#[serde(untagged)]
pub enum GetPostLikersError {
    UnknownValue(serde_json::Value),
}

/// struct for typed errors of method [`get_post_reposts`]
#[derive(Debug, Clone, Serialize, Deserialize)]
#[serde(untagged)]
pub enum GetPostRepostsError {
    UnknownValue(serde_json::Value),
}

/// struct for typed errors of method [`get_post_url_metadata`]
#[derive(Debug, Clone, Serialize, Deserialize)]
#[serde(untagged)]
pub enum GetPostUrlMetadataError {
    UnknownValue(serde_json::Value),
}

/// struct for typed errors of method [`get_posts_by_ids`]
#[derive(Debug, Clone, Serialize, Deserialize)]
#[serde(untagged)]
pub enum GetPostsByIdsError {
    UnknownValue(serde_json::Value),
}

/// struct for typed errors of method [`get_posts_by_tag`]
#[derive(Debug, Clone, Serialize, Deserialize)]
#[serde(untagged)]
pub enum GetPostsByTagError {
    UnknownValue(serde_json::Value),
}

/// struct for typed errors of method [`get_recent_engagement_posts`]
#[derive(Debug, Clone, Serialize, Deserialize)]
#[serde(untagged)]
pub enum GetRecentEngagementPostsError {
    UnknownValue(serde_json::Value),
}

/// struct for typed errors of method [`get_recommended_post_tags`]
#[derive(Debug, Clone, Serialize, Deserialize)]
#[serde(untagged)]
pub enum GetRecommendedPostTagsError {
    UnknownValue(serde_json::Value),
}

/// struct for typed errors of method [`get_recommended_timeline`]
#[derive(Debug, Clone, Serialize, Deserialize)]
#[serde(untagged)]
pub enum GetRecommendedTimelineError {
    UnknownValue(serde_json::Value),
}

/// struct for typed errors of method [`get_timeline`]
#[derive(Debug, Clone, Serialize, Deserialize)]
#[serde(untagged)]
pub enum GetTimelineError {
    UnknownValue(serde_json::Value),
}

/// struct for typed errors of method [`get_user_timeline`]
#[derive(Debug, Clone, Serialize, Deserialize)]
#[serde(untagged)]
pub enum GetUserTimelineError {
    UnknownValue(serde_json::Value),
}

/// struct for typed errors of method [`like_posts`]
#[derive(Debug, Clone, Serialize, Deserialize)]
#[serde(untagged)]
pub enum LikePostsError {
    UnknownValue(serde_json::Value),
}

/// struct for typed errors of method [`move_post_to_specific_thread`]
#[derive(Debug, Clone, Serialize, Deserialize)]
#[serde(untagged)]
pub enum MovePostToSpecificThreadError {
    UnknownValue(serde_json::Value),
}

/// struct for typed errors of method [`move_post_to_thread`]
#[derive(Debug, Clone, Serialize, Deserialize)]
#[serde(untagged)]
pub enum MovePostToThreadError {
    UnknownValue(serde_json::Value),
}

/// struct for typed errors of method [`pin_group_post`]
#[derive(Debug, Clone, Serialize, Deserialize)]
#[serde(untagged)]
pub enum PinGroupPostError {
    UnknownValue(serde_json::Value),
}

/// struct for typed errors of method [`report_post`]
#[derive(Debug, Clone, Serialize, Deserialize)]
#[serde(untagged)]
pub enum ReportPostError {
    UnknownValue(serde_json::Value),
}

/// struct for typed errors of method [`repost`]
#[derive(Debug, Clone, Serialize, Deserialize)]
#[serde(untagged)]
pub enum RepostError {
    UnknownValue(serde_json::Value),
}

/// struct for typed errors of method [`search_posts`]
#[derive(Debug, Clone, Serialize, Deserialize)]
#[serde(untagged)]
pub enum SearchPostsError {
    UnknownValue(serde_json::Value),
}

/// struct for typed errors of method [`unlike_post`]
#[derive(Debug, Clone, Serialize, Deserialize)]
#[serde(untagged)]
pub enum UnlikePostError {
    UnknownValue(serde_json::Value),
}

/// struct for typed errors of method [`unpin_group_post`]
#[derive(Debug, Clone, Serialize, Deserialize)]
#[serde(untagged)]
pub enum UnpinGroupPostError {
    UnknownValue(serde_json::Value),
}

/// struct for typed errors of method [`update_post`]
#[derive(Debug, Clone, Serialize, Deserialize)]
#[serde(untagged)]
pub enum UpdatePostError {
    UnknownValue(serde_json::Value),
}

/// struct for typed errors of method [`validate_post`]
#[derive(Debug, Clone, Serialize, Deserialize)]
#[serde(untagged)]
pub enum ValidatePostError {
    UnknownValue(serde_json::Value),
}

/// struct for typed errors of method [`view_post_video`]
#[derive(Debug, Clone, Serialize, Deserialize)]
#[serde(untagged)]
pub enum ViewPostVideoError {
    UnknownValue(serde_json::Value),
}


pub async fn create_conference_call_post(configuration: &configuration::Configuration, api_key: Option<&str>, attachment_2_filename: Option<&str>, attachment_3_filename: Option<&str>, attachment_4_filename: Option<&str>, attachment_5_filename: Option<&str>, attachment_6_filename: Option<&str>, attachment_7_filename: Option<&str>, attachment_8_filename: Option<&str>, attachment_9_filename: Option<&str>, attachment_filename: Option<&str>, call_type: Option<models::CallType>, category_id: Option<i64>, color: Option<i32>, font_size: Option<i32>, game_title: Option<&str>, group_id: Option<i64>, joinable_by: Option<models::JoinableBy>, language: Option<&str>, message_tags: Option<serde_json::Value>, signed_info: Option<&str>, text: Option<&str>, timestamp: Option<i64>, uuid: Option<&str>) -> Result<models::CreatePostResponse, Error<CreateConferenceCallPostError>> {
    let local_var_configuration = configuration;

    let local_var_client = &local_var_configuration.client;

    let local_var_uri_str = format!("{}/v2/posts/new_conference_call", local_var_configuration.base_path);
    let mut local_var_req_builder = local_var_client.request(reqwest::Method::POST, local_var_uri_str.as_str());

    if let Some(ref local_var_user_agent) = local_var_configuration.user_agent {
        local_var_req_builder = local_var_req_builder.header(reqwest::header::USER_AGENT, local_var_user_agent.clone());
    }
    let mut local_var_form = reqwest::multipart::Form::new();
    if let Some(local_var_param_value) = api_key {
        local_var_form = local_var_form.text("api_key", local_var_param_value.to_string());
    }
    if let Some(local_var_param_value) = attachment_2_filename {
        local_var_form = local_var_form.text("attachment_2_filename", local_var_param_value.to_string());
    }
    if let Some(local_var_param_value) = attachment_3_filename {
        local_var_form = local_var_form.text("attachment_3_filename", local_var_param_value.to_string());
    }
    if let Some(local_var_param_value) = attachment_4_filename {
        local_var_form = local_var_form.text("attachment_4_filename", local_var_param_value.to_string());
    }
    if let Some(local_var_param_value) = attachment_5_filename {
        local_var_form = local_var_form.text("attachment_5_filename", local_var_param_value.to_string());
    }
    if let Some(local_var_param_value) = attachment_6_filename {
        local_var_form = local_var_form.text("attachment_6_filename", local_var_param_value.to_string());
    }
    if let Some(local_var_param_value) = attachment_7_filename {
        local_var_form = local_var_form.text("attachment_7_filename", local_var_param_value.to_string());
    }
    if let Some(local_var_param_value) = attachment_8_filename {
        local_var_form = local_var_form.text("attachment_8_filename", local_var_param_value.to_string());
    }
    if let Some(local_var_param_value) = attachment_9_filename {
        local_var_form = local_var_form.text("attachment_9_filename", local_var_param_value.to_string());
    }
    if let Some(local_var_param_value) = attachment_filename {
        local_var_form = local_var_form.text("attachment_filename", local_var_param_value.to_string());
    }
    if let Some(local_var_param_value) = call_type {
        local_var_form = local_var_form.text("call_type", local_var_param_value.to_string());
    }
    if let Some(local_var_param_value) = category_id {
        local_var_form = local_var_form.text("category_id", local_var_param_value.to_string());
    }
    if let Some(local_var_param_value) = color {
        local_var_form = local_var_form.text("color", local_var_param_value.to_string());
    }
    if let Some(local_var_param_value) = font_size {
        local_var_form = local_var_form.text("font_size", local_var_param_value.to_string());
    }
    if let Some(local_var_param_value) = game_title {
        local_var_form = local_var_form.text("game_title", local_var_param_value.to_string());
    }
    if let Some(local_var_param_value) = group_id {
        local_var_form = local_var_form.text("group_id", local_var_param_value.to_string());
    }
    if let Some(local_var_param_value) = joinable_by {
        local_var_form = local_var_form.text("joinable_by", local_var_param_value.to_string());
    }
    if let Some(local_var_param_value) = language {
        local_var_form = local_var_form.text("language", local_var_param_value.to_string());
    }
    if let Some(local_var_param_value) = message_tags {
        local_var_form = local_var_form.text("message_tags", local_var_param_value.to_string());
    }
    if let Some(local_var_param_value) = signed_info {
        local_var_form = local_var_form.text("signed_info", local_var_param_value.to_string());
    }
    if let Some(local_var_param_value) = text {
        local_var_form = local_var_form.text("text", local_var_param_value.to_string());
    }
    if let Some(local_var_param_value) = timestamp {
        local_var_form = local_var_form.text("timestamp", local_var_param_value.to_string());
    }
    if let Some(local_var_param_value) = uuid {
        local_var_form = local_var_form.text("uuid", local_var_param_value.to_string());
    }
    local_var_req_builder = local_var_req_builder.multipart(local_var_form);

    let local_var_req = local_var_req_builder.build()?;
    let local_var_resp = local_var_client.execute(local_var_req).await?;

    let local_var_status = local_var_resp.status();
    let local_var_content = local_var_resp.text().await?;

    if !local_var_status.is_client_error() && !local_var_status.is_server_error() {
        serde_json::from_str(&local_var_content).map_err(Error::from)
    } else {
        let local_var_entity: Option<CreateConferenceCallPostError> = serde_json::from_str(&local_var_content).ok();
        let local_var_error = ResponseContent { status: local_var_status, content: local_var_content, entity: local_var_entity };
        Err(Error::ResponseError(local_var_error))
    }
}

pub async fn create_post(configuration: &configuration::Configuration, x_jwt: Option<&str>, attachment_2_filename: Option<&str>, attachment_3_filename: Option<&str>, attachment_4_filename: Option<&str>, attachment_5_filename: Option<&str>, attachment_6_filename: Option<&str>, attachment_7_filename: Option<&str>, attachment_8_filename: Option<&str>, attachment_9_filename: Option<&str>, attachment_filename: Option<&str>, choices: Option<Vec<String>>, color: Option<i32>, font_size: Option<i32>, group_id: Option<i64>, in_reply_to: Option<i64>, language: Option<&str>, mention_ids: Option<Vec<i64>>, message_tags: Option<serde_json::Value>, post_type: Option<models::PostType>, shared_url: Option<serde_json::Value>, text: Option<&str>, video_file_name: Option<&str>) -> Result<models::Post, Error<CreatePostError>> {
    let local_var_configuration = configuration;

    let local_var_client = &local_var_configuration.client;

    let local_var_uri_str = format!("{}/v3/posts/new", local_var_configuration.base_path);
    let mut local_var_req_builder = local_var_client.request(reqwest::Method::POST, local_var_uri_str.as_str());

    if let Some(ref local_var_user_agent) = local_var_configuration.user_agent {
        local_var_req_builder = local_var_req_builder.header(reqwest::header::USER_AGENT, local_var_user_agent.clone());
    }
    if let Some(local_var_param_value) = x_jwt {
        local_var_req_builder = local_var_req_builder.header("X-Jwt", local_var_param_value.to_string());
    }
    let mut local_var_form = reqwest::multipart::Form::new();
    if let Some(local_var_param_value) = attachment_2_filename {
        local_var_form = local_var_form.text("attachment_2_filename", local_var_param_value.to_string());
    }
    if let Some(local_var_param_value) = attachment_3_filename {
        local_var_form = local_var_form.text("attachment_3_filename", local_var_param_value.to_string());
    }
    if let Some(local_var_param_value) = attachment_4_filename {
        local_var_form = local_var_form.text("attachment_4_filename", local_var_param_value.to_string());
    }
    if let Some(local_var_param_value) = attachment_5_filename {
        local_var_form = local_var_form.text("attachment_5_filename", local_var_param_value.to_string());
    }
    if let Some(local_var_param_value) = attachment_6_filename {
        local_var_form = local_var_form.text("attachment_6_filename", local_var_param_value.to_string());
    }
    if let Some(local_var_param_value) = attachment_7_filename {
        local_var_form = local_var_form.text("attachment_7_filename", local_var_param_value.to_string());
    }
    if let Some(local_var_param_value) = attachment_8_filename {
        local_var_form = local_var_form.text("attachment_8_filename", local_var_param_value.to_string());
    }
    if let Some(local_var_param_value) = attachment_9_filename {
        local_var_form = local_var_form.text("attachment_9_filename", local_var_param_value.to_string());
    }
    if let Some(local_var_param_value) = attachment_filename {
        local_var_form = local_var_form.text("attachment_filename", local_var_param_value.to_string());
    }
    if let Some(local_var_param_value) = choices {
        local_var_form = local_var_form.text("choices[]", local_var_param_value.into_iter().map(|p| p.to_string()).collect::<Vec<String>>().join(",").to_string());
    }
    if let Some(local_var_param_value) = color {
        local_var_form = local_var_form.text("color", local_var_param_value.to_string());
    }
    if let Some(local_var_param_value) = font_size {
        local_var_form = local_var_form.text("font_size", local_var_param_value.to_string());
    }
    if let Some(local_var_param_value) = group_id {
        local_var_form = local_var_form.text("group_id", local_var_param_value.to_string());
    }
    if let Some(local_var_param_value) = in_reply_to {
        local_var_form = local_var_form.text("in_reply_to", local_var_param_value.to_string());
    }
    if let Some(local_var_param_value) = language {
        local_var_form = local_var_form.text("language", local_var_param_value.to_string());
    }
    if let Some(local_var_param_value) = mention_ids {
        local_var_form = local_var_form.text("mention_ids[]", local_var_param_value.into_iter().map(|p| p.to_string()).collect::<Vec<String>>().join(",").to_string());
    }
    if let Some(local_var_param_value) = message_tags {
        local_var_form = local_var_form.text("message_tags", local_var_param_value.to_string());
    }
    if let Some(local_var_param_value) = post_type {
        local_var_form = local_var_form.text("post_type", local_var_param_value.to_string());
    }
    if let Some(local_var_param_value) = shared_url {
        local_var_form = local_var_form.text("shared_url", local_var_param_value.to_string());
    }
    if let Some(local_var_param_value) = text {
        local_var_form = local_var_form.text("text", local_var_param_value.to_string());
    }
    if let Some(local_var_param_value) = video_file_name {
        local_var_form = local_var_form.text("video_file_name", local_var_param_value.to_string());
    }
    local_var_req_builder = local_var_req_builder.multipart(local_var_form);

    let local_var_req = local_var_req_builder.build()?;
    let local_var_resp = local_var_client.execute(local_var_req).await?;

    let local_var_status = local_var_resp.status();
    let local_var_content = local_var_resp.text().await?;

    if !local_var_status.is_client_error() && !local_var_status.is_server_error() {
        serde_json::from_str(&local_var_content).map_err(Error::from)
    } else {
        let local_var_entity: Option<CreatePostError> = serde_json::from_str(&local_var_content).ok();
        let local_var_error = ResponseContent { status: local_var_status, content: local_var_content, entity: local_var_entity };
        Err(Error::ResponseError(local_var_error))
    }
}

pub async fn create_share_post(configuration: &configuration::Configuration, api_key: Option<&str>, color: Option<i32>, font_size: Option<i32>, group_id: Option<i64>, language: Option<&str>, message_tags: Option<serde_json::Value>, shareable_id: Option<i64>, shareable_type: Option<&str>, signed_info: Option<&str>, text: Option<&str>, timestamp: Option<i64>, uuid: Option<&str>) -> Result<models::Post, Error<CreateSharePostError>> {
    let local_var_configuration = configuration;

    let local_var_client = &local_var_configuration.client;

    let local_var_uri_str = format!("{}/v2/posts/new_share_post", local_var_configuration.base_path);
    let mut local_var_req_builder = local_var_client.request(reqwest::Method::POST, local_var_uri_str.as_str());

    if let Some(ref local_var_user_agent) = local_var_configuration.user_agent {
        local_var_req_builder = local_var_req_builder.header(reqwest::header::USER_AGENT, local_var_user_agent.clone());
    }
    let mut local_var_form = reqwest::multipart::Form::new();
    if let Some(local_var_param_value) = api_key {
        local_var_form = local_var_form.text("api_key", local_var_param_value.to_string());
    }
    if let Some(local_var_param_value) = color {
        local_var_form = local_var_form.text("color", local_var_param_value.to_string());
    }
    if let Some(local_var_param_value) = font_size {
        local_var_form = local_var_form.text("font_size", local_var_param_value.to_string());
    }
    if let Some(local_var_param_value) = group_id {
        local_var_form = local_var_form.text("group_id", local_var_param_value.to_string());
    }
    if let Some(local_var_param_value) = language {
        local_var_form = local_var_form.text("language", local_var_param_value.to_string());
    }
    if let Some(local_var_param_value) = message_tags {
        local_var_form = local_var_form.text("message_tags", local_var_param_value.to_string());
    }
    if let Some(local_var_param_value) = shareable_id {
        local_var_form = local_var_form.text("shareable_id", local_var_param_value.to_string());
    }
    if let Some(local_var_param_value) = shareable_type {
        local_var_form = local_var_form.text("shareable_type", local_var_param_value.to_string());
    }
    if let Some(local_var_param_value) = signed_info {
        local_var_form = local_var_form.text("signed_info", local_var_param_value.to_string());
    }
    if let Some(local_var_param_value) = text {
        local_var_form = local_var_form.text("text", local_var_param_value.to_string());
    }
    if let Some(local_var_param_value) = timestamp {
        local_var_form = local_var_form.text("timestamp", local_var_param_value.to_string());
    }
    if let Some(local_var_param_value) = uuid {
        local_var_form = local_var_form.text("uuid", local_var_param_value.to_string());
    }
    local_var_req_builder = local_var_req_builder.multipart(local_var_form);

    let local_var_req = local_var_req_builder.build()?;
    let local_var_resp = local_var_client.execute(local_var_req).await?;

    let local_var_status = local_var_resp.status();
    let local_var_content = local_var_resp.text().await?;

    if !local_var_status.is_client_error() && !local_var_status.is_server_error() {
        serde_json::from_str(&local_var_content).map_err(Error::from)
    } else {
        let local_var_entity: Option<CreateSharePostError> = serde_json::from_str(&local_var_content).ok();
        let local_var_error = ResponseContent { status: local_var_status, content: local_var_content, entity: local_var_entity };
        Err(Error::ResponseError(local_var_error))
    }
}

pub async fn delete_all_posts(configuration: &configuration::Configuration, ) -> Result<(), Error<DeleteAllPostsError>> {
    let local_var_configuration = configuration;

    let local_var_client = &local_var_configuration.client;

    let local_var_uri_str = format!("{}/v1/posts/delete_all_post", local_var_configuration.base_path);
    let mut local_var_req_builder = local_var_client.request(reqwest::Method::POST, local_var_uri_str.as_str());

    if let Some(ref local_var_user_agent) = local_var_configuration.user_agent {
        local_var_req_builder = local_var_req_builder.header(reqwest::header::USER_AGENT, local_var_user_agent.clone());
    }

    let local_var_req = local_var_req_builder.build()?;
    let local_var_resp = local_var_client.execute(local_var_req).await?;

    let local_var_status = local_var_resp.status();
    let local_var_content = local_var_resp.text().await?;

    if !local_var_status.is_client_error() && !local_var_status.is_server_error() {
        Ok(())
    } else {
        let local_var_entity: Option<DeleteAllPostsError> = serde_json::from_str(&local_var_content).ok();
        let local_var_error = ResponseContent { status: local_var_status, content: local_var_content, entity: local_var_entity };
        Err(Error::ResponseError(local_var_error))
    }
}

pub async fn delete_posts(configuration: &configuration::Configuration, posts_ids: Option<Vec<i64>>) -> Result<(), Error<DeletePostsError>> {
    let local_var_configuration = configuration;

    let local_var_client = &local_var_configuration.client;

    let local_var_uri_str = format!("{}/v2/posts/mass_destroy", local_var_configuration.base_path);
    let mut local_var_req_builder = local_var_client.request(reqwest::Method::POST, local_var_uri_str.as_str());

    if let Some(ref local_var_user_agent) = local_var_configuration.user_agent {
        local_var_req_builder = local_var_req_builder.header(reqwest::header::USER_AGENT, local_var_user_agent.clone());
    }
    let mut local_var_form = reqwest::multipart::Form::new();
    if let Some(local_var_param_value) = posts_ids {
        local_var_form = local_var_form.text("posts_ids[]", local_var_param_value.into_iter().map(|p| p.to_string()).collect::<Vec<String>>().join(",").to_string());
    }
    local_var_req_builder = local_var_req_builder.multipart(local_var_form);

    let local_var_req = local_var_req_builder.build()?;
    let local_var_resp = local_var_client.execute(local_var_req).await?;

    let local_var_status = local_var_resp.status();
    let local_var_content = local_var_resp.text().await?;

    if !local_var_status.is_client_error() && !local_var_status.is_server_error() {
        Ok(())
    } else {
        let local_var_entity: Option<DeletePostsError> = serde_json::from_str(&local_var_content).ok();
        let local_var_error = ResponseContent { status: local_var_status, content: local_var_content, entity: local_var_entity };
        Err(Error::ResponseError(local_var_error))
    }
}

pub async fn get_active_call_post(configuration: &configuration::Configuration, user_id: Option<i64>) -> Result<models::PostResponse, Error<GetActiveCallPostError>> {
    let local_var_configuration = configuration;

    let local_var_client = &local_var_configuration.client;

    let local_var_uri_str = format!("{}/v1/posts/active_call", local_var_configuration.base_path);
    let mut local_var_req_builder = local_var_client.request(reqwest::Method::GET, local_var_uri_str.as_str());

    if let Some(ref local_var_str) = user_id {
        local_var_req_builder = local_var_req_builder.query(&[("user_id", &local_var_str.to_string())]);
    }
    if let Some(ref local_var_user_agent) = local_var_configuration.user_agent {
        local_var_req_builder = local_var_req_builder.header(reqwest::header::USER_AGENT, local_var_user_agent.clone());
    }

    let local_var_req = local_var_req_builder.build()?;
    let local_var_resp = local_var_client.execute(local_var_req).await?;

    let local_var_status = local_var_resp.status();
    let local_var_content = local_var_resp.text().await?;

    if !local_var_status.is_client_error() && !local_var_status.is_server_error() {
        serde_json::from_str(&local_var_content).map_err(Error::from)
    } else {
        let local_var_entity: Option<GetActiveCallPostError> = serde_json::from_str(&local_var_content).ok();
        let local_var_error = ResponseContent { status: local_var_status, content: local_var_content, entity: local_var_entity };
        Err(Error::ResponseError(local_var_error))
    }
}

pub async fn get_call_followers_timeline(configuration: &configuration::Configuration, from_timestamp: Option<i64>, number: Option<i32>, category_id: Option<i64>, call_type: Option<models::CallType>, include_circle_call: Option<bool>, exclude_recent_gomimushi: Option<bool>) -> Result<models::PostsResponse, Error<GetCallFollowersTimelineError>> {
    let local_var_configuration = configuration;

    let local_var_client = &local_var_configuration.client;

    let local_var_uri_str = format!("{}/v2/posts/call_followers_timeline", local_var_configuration.base_path);
    let mut local_var_req_builder = local_var_client.request(reqwest::Method::GET, local_var_uri_str.as_str());

    if let Some(ref local_var_str) = from_timestamp {
        local_var_req_builder = local_var_req_builder.query(&[("from_timestamp", &local_var_str.to_string())]);
    }
    if let Some(ref local_var_str) = number {
        local_var_req_builder = local_var_req_builder.query(&[("number", &local_var_str.to_string())]);
    }
    if let Some(ref local_var_str) = category_id {
        local_var_req_builder = local_var_req_builder.query(&[("category_id", &local_var_str.to_string())]);
    }
    if let Some(ref local_var_str) = call_type {
        local_var_req_builder = local_var_req_builder.query(&[("call_type", &local_var_str.to_string())]);
    }
    if let Some(ref local_var_str) = include_circle_call {
        local_var_req_builder = local_var_req_builder.query(&[("include_circle_call", &local_var_str.to_string())]);
    }
    if let Some(ref local_var_str) = exclude_recent_gomimushi {
        local_var_req_builder = local_var_req_builder.query(&[("exclude_recent_gomimushi", &local_var_str.to_string())]);
    }
    if let Some(ref local_var_user_agent) = local_var_configuration.user_agent {
        local_var_req_builder = local_var_req_builder.header(reqwest::header::USER_AGENT, local_var_user_agent.clone());
    }

    let local_var_req = local_var_req_builder.build()?;
    let local_var_resp = local_var_client.execute(local_var_req).await?;

    let local_var_status = local_var_resp.status();
    let local_var_content = local_var_resp.text().await?;

    if !local_var_status.is_client_error() && !local_var_status.is_server_error() {
        serde_json::from_str(&local_var_content).map_err(Error::from)
    } else {
        let local_var_entity: Option<GetCallFollowersTimelineError> = serde_json::from_str(&local_var_content).ok();
        let local_var_error = ResponseContent { status: local_var_status, content: local_var_content, entity: local_var_entity };
        Err(Error::ResponseError(local_var_error))
    }
}

pub async fn get_call_timeline(configuration: &configuration::Configuration, group_id: Option<i64>, from_timestamp: Option<i64>, number: Option<i32>, category_id: Option<i64>, call_type: Option<models::CallType>, include_circle_call: Option<bool>, cross_generation: Option<bool>, exclude_recent_gomimushi: Option<bool>) -> Result<models::PostsResponse, Error<GetCallTimelineError>> {
    let local_var_configuration = configuration;

    let local_var_client = &local_var_configuration.client;

    let local_var_uri_str = format!("{}/v2/posts/call_timeline", local_var_configuration.base_path);
    let mut local_var_req_builder = local_var_client.request(reqwest::Method::GET, local_var_uri_str.as_str());

    if let Some(ref local_var_str) = group_id {
        local_var_req_builder = local_var_req_builder.query(&[("group_id", &local_var_str.to_string())]);
    }
    if let Some(ref local_var_str) = from_timestamp {
        local_var_req_builder = local_var_req_builder.query(&[("from_timestamp", &local_var_str.to_string())]);
    }
    if let Some(ref local_var_str) = number {
        local_var_req_builder = local_var_req_builder.query(&[("number", &local_var_str.to_string())]);
    }
    if let Some(ref local_var_str) = category_id {
        local_var_req_builder = local_var_req_builder.query(&[("category_id", &local_var_str.to_string())]);
    }
    if let Some(ref local_var_str) = call_type {
        local_var_req_builder = local_var_req_builder.query(&[("call_type", &local_var_str.to_string())]);
    }
    if let Some(ref local_var_str) = include_circle_call {
        local_var_req_builder = local_var_req_builder.query(&[("include_circle_call", &local_var_str.to_string())]);
    }
    if let Some(ref local_var_str) = cross_generation {
        local_var_req_builder = local_var_req_builder.query(&[("cross_generation", &local_var_str.to_string())]);
    }
    if let Some(ref local_var_str) = exclude_recent_gomimushi {
        local_var_req_builder = local_var_req_builder.query(&[("exclude_recent_gomimushi", &local_var_str.to_string())]);
    }
    if let Some(ref local_var_user_agent) = local_var_configuration.user_agent {
        local_var_req_builder = local_var_req_builder.header(reqwest::header::USER_AGENT, local_var_user_agent.clone());
    }

    let local_var_req = local_var_req_builder.build()?;
    let local_var_resp = local_var_client.execute(local_var_req).await?;

    let local_var_status = local_var_resp.status();
    let local_var_content = local_var_resp.text().await?;

    if !local_var_status.is_client_error() && !local_var_status.is_server_error() {
        serde_json::from_str(&local_var_content).map_err(Error::from)
    } else {
        let local_var_entity: Option<GetCallTimelineError> = serde_json::from_str(&local_var_content).ok();
        let local_var_error = ResponseContent { status: local_var_status, content: local_var_content, entity: local_var_entity };
        Err(Error::ResponseError(local_var_error))
    }
}

pub async fn get_following_timeline(configuration: &configuration::Configuration, from: Option<&str>, from_post_id: Option<i64>, only_root: Option<bool>, order_by: Option<&str>, number: Option<i32>, mxn: Option<i32>, reduce_selfie: Option<bool>, custom_generation_range: Option<bool>) -> Result<models::PostsResponse, Error<GetFollowingTimelineError>> {
    let local_var_configuration = configuration;

    let local_var_client = &local_var_configuration.client;

    let local_var_uri_str = format!("{}/v2/posts/following_timeline", local_var_configuration.base_path);
    let mut local_var_req_builder = local_var_client.request(reqwest::Method::GET, local_var_uri_str.as_str());

    if let Some(ref local_var_str) = from {
        local_var_req_builder = local_var_req_builder.query(&[("from", &local_var_str.to_string())]);
    }
    if let Some(ref local_var_str) = from_post_id {
        local_var_req_builder = local_var_req_builder.query(&[("from_post_id", &local_var_str.to_string())]);
    }
    if let Some(ref local_var_str) = only_root {
        local_var_req_builder = local_var_req_builder.query(&[("only_root", &local_var_str.to_string())]);
    }
    if let Some(ref local_var_str) = order_by {
        local_var_req_builder = local_var_req_builder.query(&[("order_by", &local_var_str.to_string())]);
    }
    if let Some(ref local_var_str) = number {
        local_var_req_builder = local_var_req_builder.query(&[("number", &local_var_str.to_string())]);
    }
    if let Some(ref local_var_str) = mxn {
        local_var_req_builder = local_var_req_builder.query(&[("mxn", &local_var_str.to_string())]);
    }
    if let Some(ref local_var_str) = reduce_selfie {
        local_var_req_builder = local_var_req_builder.query(&[("reduce_selfie", &local_var_str.to_string())]);
    }
    if let Some(ref local_var_str) = custom_generation_range {
        local_var_req_builder = local_var_req_builder.query(&[("custom_generation_range", &local_var_str.to_string())]);
    }
    if let Some(ref local_var_user_agent) = local_var_configuration.user_agent {
        local_var_req_builder = local_var_req_builder.header(reqwest::header::USER_AGENT, local_var_user_agent.clone());
    }

    let local_var_req = local_var_req_builder.build()?;
    let local_var_resp = local_var_client.execute(local_var_req).await?;

    let local_var_status = local_var_resp.status();
    let local_var_content = local_var_resp.text().await?;

    if !local_var_status.is_client_error() && !local_var_status.is_server_error() {
        serde_json::from_str(&local_var_content).map_err(Error::from)
    } else {
        let local_var_entity: Option<GetFollowingTimelineError> = serde_json::from_str(&local_var_content).ok();
        let local_var_error = ResponseContent { status: local_var_status, content: local_var_content, entity: local_var_entity };
        Err(Error::ResponseError(local_var_error))
    }
}

pub async fn get_group_calls(configuration: &configuration::Configuration, number: Option<i32>, group_category_id: Option<i64>, from_timestamp: Option<i64>, scope: Option<&str>) -> Result<models::PostsResponse, Error<GetGroupCallsError>> {
    let local_var_configuration = configuration;

    let local_var_client = &local_var_configuration.client;

    let local_var_uri_str = format!("{}/v1/posts/group_calls", local_var_configuration.base_path);
    let mut local_var_req_builder = local_var_client.request(reqwest::Method::GET, local_var_uri_str.as_str());

    if let Some(ref local_var_str) = number {
        local_var_req_builder = local_var_req_builder.query(&[("number", &local_var_str.to_string())]);
    }
    if let Some(ref local_var_str) = group_category_id {
        local_var_req_builder = local_var_req_builder.query(&[("group_category_id", &local_var_str.to_string())]);
    }
    if let Some(ref local_var_str) = from_timestamp {
        local_var_req_builder = local_var_req_builder.query(&[("from_timestamp", &local_var_str.to_string())]);
    }
    if let Some(ref local_var_str) = scope {
        local_var_req_builder = local_var_req_builder.query(&[("scope", &local_var_str.to_string())]);
    }
    if let Some(ref local_var_user_agent) = local_var_configuration.user_agent {
        local_var_req_builder = local_var_req_builder.header(reqwest::header::USER_AGENT, local_var_user_agent.clone());
    }

    let local_var_req = local_var_req_builder.build()?;
    let local_var_resp = local_var_client.execute(local_var_req).await?;

    let local_var_status = local_var_resp.status();
    let local_var_content = local_var_resp.text().await?;

    if !local_var_status.is_client_error() && !local_var_status.is_server_error() {
        serde_json::from_str(&local_var_content).map_err(Error::from)
    } else {
        let local_var_entity: Option<GetGroupCallsError> = serde_json::from_str(&local_var_content).ok();
        let local_var_error = ResponseContent { status: local_var_status, content: local_var_content, entity: local_var_entity };
        Err(Error::ResponseError(local_var_error))
    }
}

pub async fn get_group_timeline(configuration: &configuration::Configuration, group_id: Option<i64>, from_post_id: Option<i64>, reverse: Option<bool>, post_type: Option<models::PostType>, number: Option<i32>, only_root: Option<bool>) -> Result<models::PostsResponse, Error<GetGroupTimelineError>> {
    let local_var_configuration = configuration;

    let local_var_client = &local_var_configuration.client;

    let local_var_uri_str = format!("{}/v2/posts/group_timeline", local_var_configuration.base_path);
    let mut local_var_req_builder = local_var_client.request(reqwest::Method::GET, local_var_uri_str.as_str());

    if let Some(ref local_var_str) = group_id {
        local_var_req_builder = local_var_req_builder.query(&[("group_id", &local_var_str.to_string())]);
    }
    if let Some(ref local_var_str) = from_post_id {
        local_var_req_builder = local_var_req_builder.query(&[("from_post_id", &local_var_str.to_string())]);
    }
    if let Some(ref local_var_str) = reverse {
        local_var_req_builder = local_var_req_builder.query(&[("reverse", &local_var_str.to_string())]);
    }
    if let Some(ref local_var_str) = post_type {
        local_var_req_builder = local_var_req_builder.query(&[("post_type", &local_var_str.to_string())]);
    }
    if let Some(ref local_var_str) = number {
        local_var_req_builder = local_var_req_builder.query(&[("number", &local_var_str.to_string())]);
    }
    if let Some(ref local_var_str) = only_root {
        local_var_req_builder = local_var_req_builder.query(&[("only_root", &local_var_str.to_string())]);
    }
    if let Some(ref local_var_user_agent) = local_var_configuration.user_agent {
        local_var_req_builder = local_var_req_builder.header(reqwest::header::USER_AGENT, local_var_user_agent.clone());
    }

    let local_var_req = local_var_req_builder.build()?;
    let local_var_resp = local_var_client.execute(local_var_req).await?;

    let local_var_status = local_var_resp.status();
    let local_var_content = local_var_resp.text().await?;

    if !local_var_status.is_client_error() && !local_var_status.is_server_error() {
        serde_json::from_str(&local_var_content).map_err(Error::from)
    } else {
        let local_var_entity: Option<GetGroupTimelineError> = serde_json::from_str(&local_var_content).ok();
        let local_var_error = ResponseContent { status: local_var_status, content: local_var_content, entity: local_var_entity };
        Err(Error::ResponseError(local_var_error))
    }
}

pub async fn get_my_posts(configuration: &configuration::Configuration, from_post_id: Option<i64>, number: Option<i32>, include_group_post: Option<bool>) -> Result<models::PostsResponse, Error<GetMyPostsError>> {
    let local_var_configuration = configuration;

    let local_var_client = &local_var_configuration.client;

    let local_var_uri_str = format!("{}/v2/posts/mine", local_var_configuration.base_path);
    let mut local_var_req_builder = local_var_client.request(reqwest::Method::GET, local_var_uri_str.as_str());

    if let Some(ref local_var_str) = from_post_id {
        local_var_req_builder = local_var_req_builder.query(&[("from_post_id", &local_var_str.to_string())]);
    }
    if let Some(ref local_var_str) = number {
        local_var_req_builder = local_var_req_builder.query(&[("number", &local_var_str.to_string())]);
    }
    if let Some(ref local_var_str) = include_group_post {
        local_var_req_builder = local_var_req_builder.query(&[("include_group_post", &local_var_str.to_string())]);
    }
    if let Some(ref local_var_user_agent) = local_var_configuration.user_agent {
        local_var_req_builder = local_var_req_builder.header(reqwest::header::USER_AGENT, local_var_user_agent.clone());
    }

    let local_var_req = local_var_req_builder.build()?;
    let local_var_resp = local_var_client.execute(local_var_req).await?;

    let local_var_status = local_var_resp.status();
    let local_var_content = local_var_resp.text().await?;

    if !local_var_status.is_client_error() && !local_var_status.is_server_error() {
        serde_json::from_str(&local_var_content).map_err(Error::from)
    } else {
        let local_var_entity: Option<GetMyPostsError> = serde_json::from_str(&local_var_content).ok();
        let local_var_error = ResponseContent { status: local_var_status, content: local_var_content, entity: local_var_entity };
        Err(Error::ResponseError(local_var_error))
    }
}

pub async fn get_post(configuration: &configuration::Configuration, id: i64, cache_control: Option<&str>) -> Result<models::PostResponse, Error<GetPostError>> {
    let local_var_configuration = configuration;

    let local_var_client = &local_var_configuration.client;

    let local_var_uri_str = format!("{}/v2/posts/{id}", local_var_configuration.base_path, id=id);
    let mut local_var_req_builder = local_var_client.request(reqwest::Method::GET, local_var_uri_str.as_str());

    if let Some(ref local_var_user_agent) = local_var_configuration.user_agent {
        local_var_req_builder = local_var_req_builder.header(reqwest::header::USER_AGENT, local_var_user_agent.clone());
    }
    if let Some(local_var_param_value) = cache_control {
        local_var_req_builder = local_var_req_builder.header("Cache-Control", local_var_param_value.to_string());
    }

    let local_var_req = local_var_req_builder.build()?;
    let local_var_resp = local_var_client.execute(local_var_req).await?;

    let local_var_status = local_var_resp.status();
    let local_var_content = local_var_resp.text().await?;

    if !local_var_status.is_client_error() && !local_var_status.is_server_error() {
        serde_json::from_str(&local_var_content).map_err(Error::from)
    } else {
        let local_var_entity: Option<GetPostError> = serde_json::from_str(&local_var_content).ok();
        let local_var_error = ResponseContent { status: local_var_status, content: local_var_content, entity: local_var_entity };
        Err(Error::ResponseError(local_var_error))
    }
}

pub async fn get_post_gift_transactions(configuration: &configuration::Configuration, post_id: i64, number: Option<i32>, from: Option<&str>) -> Result<models::GiftTransactionsResponse, Error<GetPostGiftTransactionsError>> {
    let local_var_configuration = configuration;

    let local_var_client = &local_var_configuration.client;

    let local_var_uri_str = format!("{}/v1/posts/{post_id}/gift_transactions", local_var_configuration.base_path, post_id=post_id);
    let mut local_var_req_builder = local_var_client.request(reqwest::Method::GET, local_var_uri_str.as_str());

    if let Some(ref local_var_str) = number {
        local_var_req_builder = local_var_req_builder.query(&[("number", &local_var_str.to_string())]);
    }
    if let Some(ref local_var_str) = from {
        local_var_req_builder = local_var_req_builder.query(&[("from", &local_var_str.to_string())]);
    }
    if let Some(ref local_var_user_agent) = local_var_configuration.user_agent {
        local_var_req_builder = local_var_req_builder.header(reqwest::header::USER_AGENT, local_var_user_agent.clone());
    }

    let local_var_req = local_var_req_builder.build()?;
    let local_var_resp = local_var_client.execute(local_var_req).await?;

    let local_var_status = local_var_resp.status();
    let local_var_content = local_var_resp.text().await?;

    if !local_var_status.is_client_error() && !local_var_status.is_server_error() {
        serde_json::from_str(&local_var_content).map_err(Error::from)
    } else {
        let local_var_entity: Option<GetPostGiftTransactionsError> = serde_json::from_str(&local_var_content).ok();
        let local_var_error = ResponseContent { status: local_var_status, content: local_var_content, entity: local_var_entity };
        Err(Error::ResponseError(local_var_error))
    }
}

pub async fn get_post_likers(configuration: &configuration::Configuration, id: i64, from_id: Option<i64>) -> Result<models::PostLikersResponse, Error<GetPostLikersError>> {
    let local_var_configuration = configuration;

    let local_var_client = &local_var_configuration.client;

    let local_var_uri_str = format!("{}/v1/posts/{id}/likers", local_var_configuration.base_path, id=id);
    let mut local_var_req_builder = local_var_client.request(reqwest::Method::GET, local_var_uri_str.as_str());

    if let Some(ref local_var_str) = from_id {
        local_var_req_builder = local_var_req_builder.query(&[("from_id", &local_var_str.to_string())]);
    }
    if let Some(ref local_var_user_agent) = local_var_configuration.user_agent {
        local_var_req_builder = local_var_req_builder.header(reqwest::header::USER_AGENT, local_var_user_agent.clone());
    }

    let local_var_req = local_var_req_builder.build()?;
    let local_var_resp = local_var_client.execute(local_var_req).await?;

    let local_var_status = local_var_resp.status();
    let local_var_content = local_var_resp.text().await?;

    if !local_var_status.is_client_error() && !local_var_status.is_server_error() {
        serde_json::from_str(&local_var_content).map_err(Error::from)
    } else {
        let local_var_entity: Option<GetPostLikersError> = serde_json::from_str(&local_var_content).ok();
        let local_var_error = ResponseContent { status: local_var_status, content: local_var_content, entity: local_var_entity };
        Err(Error::ResponseError(local_var_error))
    }
}

pub async fn get_post_reposts(configuration: &configuration::Configuration, id: i64, from_post_id: Option<i64>) -> Result<models::PostsResponse, Error<GetPostRepostsError>> {
    let local_var_configuration = configuration;

    let local_var_client = &local_var_configuration.client;

    let local_var_uri_str = format!("{}/v2/posts/{id}/reposts", local_var_configuration.base_path, id=id);
    let mut local_var_req_builder = local_var_client.request(reqwest::Method::GET, local_var_uri_str.as_str());

    if let Some(ref local_var_str) = from_post_id {
        local_var_req_builder = local_var_req_builder.query(&[("from_post_id", &local_var_str.to_string())]);
    }
    if let Some(ref local_var_user_agent) = local_var_configuration.user_agent {
        local_var_req_builder = local_var_req_builder.header(reqwest::header::USER_AGENT, local_var_user_agent.clone());
    }

    let local_var_req = local_var_req_builder.build()?;
    let local_var_resp = local_var_client.execute(local_var_req).await?;

    let local_var_status = local_var_resp.status();
    let local_var_content = local_var_resp.text().await?;

    if !local_var_status.is_client_error() && !local_var_status.is_server_error() {
        serde_json::from_str(&local_var_content).map_err(Error::from)
    } else {
        let local_var_entity: Option<GetPostRepostsError> = serde_json::from_str(&local_var_content).ok();
        let local_var_error = ResponseContent { status: local_var_status, content: local_var_content, entity: local_var_entity };
        Err(Error::ResponseError(local_var_error))
    }
}

pub async fn get_post_url_metadata(configuration: &configuration::Configuration, url: Option<&str>) -> Result<models::SharedUrl, Error<GetPostUrlMetadataError>> {
    let local_var_configuration = configuration;

    let local_var_client = &local_var_configuration.client;

    let local_var_uri_str = format!("{}/v2/posts/url_metadata", local_var_configuration.base_path);
    let mut local_var_req_builder = local_var_client.request(reqwest::Method::GET, local_var_uri_str.as_str());

    if let Some(ref local_var_str) = url {
        local_var_req_builder = local_var_req_builder.query(&[("url", &local_var_str.to_string())]);
    }
    if let Some(ref local_var_user_agent) = local_var_configuration.user_agent {
        local_var_req_builder = local_var_req_builder.header(reqwest::header::USER_AGENT, local_var_user_agent.clone());
    }

    let local_var_req = local_var_req_builder.build()?;
    let local_var_resp = local_var_client.execute(local_var_req).await?;

    let local_var_status = local_var_resp.status();
    let local_var_content = local_var_resp.text().await?;

    if !local_var_status.is_client_error() && !local_var_status.is_server_error() {
        serde_json::from_str(&local_var_content).map_err(Error::from)
    } else {
        let local_var_entity: Option<GetPostUrlMetadataError> = serde_json::from_str(&local_var_content).ok();
        let local_var_error = ResponseContent { status: local_var_status, content: local_var_content, entity: local_var_entity };
        Err(Error::ResponseError(local_var_error))
    }
}

pub async fn get_posts_by_ids(configuration: &configuration::Configuration, post_ids: Option<Vec<i64>>) -> Result<models::PostsResponse, Error<GetPostsByIdsError>> {
    let local_var_configuration = configuration;

    let local_var_client = &local_var_configuration.client;

    let local_var_uri_str = format!("{}/v2/posts/multiple", local_var_configuration.base_path);
    let mut local_var_req_builder = local_var_client.request(reqwest::Method::GET, local_var_uri_str.as_str());

    if let Some(ref local_var_str) = post_ids {
        local_var_req_builder = match "multi" {
            "multi" => local_var_req_builder.query(&local_var_str.into_iter().map(|p| ("post_ids[]".to_owned(), p.to_string())).collect::<Vec<(std::string::String, std::string::String)>>()),
            _ => local_var_req_builder.query(&[("post_ids[]", &local_var_str.into_iter().map(|p| p.to_string()).collect::<Vec<String>>().join(",").to_string())]),
        };
    }
    if let Some(ref local_var_user_agent) = local_var_configuration.user_agent {
        local_var_req_builder = local_var_req_builder.header(reqwest::header::USER_AGENT, local_var_user_agent.clone());
    }

    let local_var_req = local_var_req_builder.build()?;
    let local_var_resp = local_var_client.execute(local_var_req).await?;

    let local_var_status = local_var_resp.status();
    let local_var_content = local_var_resp.text().await?;

    if !local_var_status.is_client_error() && !local_var_status.is_server_error() {
        serde_json::from_str(&local_var_content).map_err(Error::from)
    } else {
        let local_var_entity: Option<GetPostsByIdsError> = serde_json::from_str(&local_var_content).ok();
        let local_var_error = ResponseContent { status: local_var_status, content: local_var_content, entity: local_var_entity };
        Err(Error::ResponseError(local_var_error))
    }
}

pub async fn get_posts_by_tag(configuration: &configuration::Configuration, tag: &str, from_post_id: Option<i64>, number: Option<i32>) -> Result<models::PostsResponse, Error<GetPostsByTagError>> {
    let local_var_configuration = configuration;

    let local_var_client = &local_var_configuration.client;

    let local_var_uri_str = format!("{}/v2/posts/tags/{tag}", local_var_configuration.base_path, tag=crate::gen::apis::urlencode(tag));
    let mut local_var_req_builder = local_var_client.request(reqwest::Method::GET, local_var_uri_str.as_str());

    if let Some(ref local_var_str) = from_post_id {
        local_var_req_builder = local_var_req_builder.query(&[("from_post_id", &local_var_str.to_string())]);
    }
    if let Some(ref local_var_str) = number {
        local_var_req_builder = local_var_req_builder.query(&[("number", &local_var_str.to_string())]);
    }
    if let Some(ref local_var_user_agent) = local_var_configuration.user_agent {
        local_var_req_builder = local_var_req_builder.header(reqwest::header::USER_AGENT, local_var_user_agent.clone());
    }

    let local_var_req = local_var_req_builder.build()?;
    let local_var_resp = local_var_client.execute(local_var_req).await?;

    let local_var_status = local_var_resp.status();
    let local_var_content = local_var_resp.text().await?;

    if !local_var_status.is_client_error() && !local_var_status.is_server_error() {
        serde_json::from_str(&local_var_content).map_err(Error::from)
    } else {
        let local_var_entity: Option<GetPostsByTagError> = serde_json::from_str(&local_var_content).ok();
        let local_var_error = ResponseContent { status: local_var_status, content: local_var_content, entity: local_var_entity };
        Err(Error::ResponseError(local_var_error))
    }
}

pub async fn get_recent_engagement_posts(configuration: &configuration::Configuration, number: Option<i32>) -> Result<models::PostsResponse, Error<GetRecentEngagementPostsError>> {
    let local_var_configuration = configuration;

    let local_var_client = &local_var_configuration.client;

    let local_var_uri_str = format!("{}/v2/posts/recent_engagement", local_var_configuration.base_path);
    let mut local_var_req_builder = local_var_client.request(reqwest::Method::GET, local_var_uri_str.as_str());

    if let Some(ref local_var_str) = number {
        local_var_req_builder = local_var_req_builder.query(&[("number", &local_var_str.to_string())]);
    }
    if let Some(ref local_var_user_agent) = local_var_configuration.user_agent {
        local_var_req_builder = local_var_req_builder.header(reqwest::header::USER_AGENT, local_var_user_agent.clone());
    }

    let local_var_req = local_var_req_builder.build()?;
    let local_var_resp = local_var_client.execute(local_var_req).await?;

    let local_var_status = local_var_resp.status();
    let local_var_content = local_var_resp.text().await?;

    if !local_var_status.is_client_error() && !local_var_status.is_server_error() {
        serde_json::from_str(&local_var_content).map_err(Error::from)
    } else {
        let local_var_entity: Option<GetRecentEngagementPostsError> = serde_json::from_str(&local_var_content).ok();
        let local_var_error = ResponseContent { status: local_var_status, content: local_var_content, entity: local_var_entity };
        Err(Error::ResponseError(local_var_error))
    }
}

pub async fn get_recommended_post_tags(configuration: &configuration::Configuration, save_recent_search: Option<bool>, tag: Option<&str>) -> Result<models::PostTagsResponse, Error<GetRecommendedPostTagsError>> {
    let local_var_configuration = configuration;

    let local_var_client = &local_var_configuration.client;

    let local_var_uri_str = format!("{}/v1/posts/recommended_tag", local_var_configuration.base_path);
    let mut local_var_req_builder = local_var_client.request(reqwest::Method::POST, local_var_uri_str.as_str());

    if let Some(ref local_var_user_agent) = local_var_configuration.user_agent {
        local_var_req_builder = local_var_req_builder.header(reqwest::header::USER_AGENT, local_var_user_agent.clone());
    }
    let mut local_var_form = reqwest::multipart::Form::new();
    if let Some(local_var_param_value) = save_recent_search {
        local_var_form = local_var_form.text("save_recent_search", local_var_param_value.to_string());
    }
    if let Some(local_var_param_value) = tag {
        local_var_form = local_var_form.text("tag", local_var_param_value.to_string());
    }
    local_var_req_builder = local_var_req_builder.multipart(local_var_form);

    let local_var_req = local_var_req_builder.build()?;
    let local_var_resp = local_var_client.execute(local_var_req).await?;

    let local_var_status = local_var_resp.status();
    let local_var_content = local_var_resp.text().await?;

    if !local_var_status.is_client_error() && !local_var_status.is_server_error() {
        serde_json::from_str(&local_var_content).map_err(Error::from)
    } else {
        let local_var_entity: Option<GetRecommendedPostTagsError> = serde_json::from_str(&local_var_content).ok();
        let local_var_error = ResponseContent { status: local_var_status, content: local_var_content, entity: local_var_entity };
        Err(Error::ResponseError(local_var_error))
    }
}

pub async fn get_recommended_timeline(configuration: &configuration::Configuration, experiment_num: Option<i32>, variant_num: Option<i32>, number: Option<i32>) -> Result<models::PostsResponse, Error<GetRecommendedTimelineError>> {
    let local_var_configuration = configuration;

    let local_var_client = &local_var_configuration.client;

    let local_var_uri_str = format!("{}/v2/posts/recommended_timeline", local_var_configuration.base_path);
    let mut local_var_req_builder = local_var_client.request(reqwest::Method::GET, local_var_uri_str.as_str());

    if let Some(ref local_var_str) = experiment_num {
        local_var_req_builder = local_var_req_builder.query(&[("experiment_num", &local_var_str.to_string())]);
    }
    if let Some(ref local_var_str) = variant_num {
        local_var_req_builder = local_var_req_builder.query(&[("variant_num", &local_var_str.to_string())]);
    }
    if let Some(ref local_var_str) = number {
        local_var_req_builder = local_var_req_builder.query(&[("number", &local_var_str.to_string())]);
    }
    if let Some(ref local_var_user_agent) = local_var_configuration.user_agent {
        local_var_req_builder = local_var_req_builder.header(reqwest::header::USER_AGENT, local_var_user_agent.clone());
    }

    let local_var_req = local_var_req_builder.build()?;
    let local_var_resp = local_var_client.execute(local_var_req).await?;

    let local_var_status = local_var_resp.status();
    let local_var_content = local_var_resp.text().await?;

    if !local_var_status.is_client_error() && !local_var_status.is_server_error() {
        serde_json::from_str(&local_var_content).map_err(Error::from)
    } else {
        let local_var_entity: Option<GetRecommendedTimelineError> = serde_json::from_str(&local_var_content).ok();
        let local_var_error = ResponseContent { status: local_var_status, content: local_var_content, entity: local_var_entity };
        Err(Error::ResponseError(local_var_error))
    }
}

pub async fn get_timeline(configuration: &configuration::Configuration, noreply_mode: models::NoreplyMode, order_by: Option<&str>, experiment_older_age_rules: Option<bool>, from: Option<&str>, from_post_id: Option<i64>, number: Option<i32>, mxn: Option<i32>, en: Option<i32>, vn: Option<i32>, reduce_selfie: Option<bool>, custom_generation_range: Option<bool>) -> Result<models::PostsResponse, Error<GetTimelineError>> {
    let local_var_configuration = configuration;

    let local_var_client = &local_var_configuration.client;

    let local_var_uri_str = format!("{}/v2/posts/{noreply_mode}timeline", local_var_configuration.base_path, noreply_mode=noreply_mode.to_string());
    let mut local_var_req_builder = local_var_client.request(reqwest::Method::GET, local_var_uri_str.as_str());

    if let Some(ref local_var_str) = order_by {
        local_var_req_builder = local_var_req_builder.query(&[("order_by", &local_var_str.to_string())]);
    }
    if let Some(ref local_var_str) = experiment_older_age_rules {
        local_var_req_builder = local_var_req_builder.query(&[("experiment_older_age_rules", &local_var_str.to_string())]);
    }
    if let Some(ref local_var_str) = from {
        local_var_req_builder = local_var_req_builder.query(&[("from", &local_var_str.to_string())]);
    }
    if let Some(ref local_var_str) = from_post_id {
        local_var_req_builder = local_var_req_builder.query(&[("from_post_id", &local_var_str.to_string())]);
    }
    if let Some(ref local_var_str) = number {
        local_var_req_builder = local_var_req_builder.query(&[("number", &local_var_str.to_string())]);
    }
    if let Some(ref local_var_str) = mxn {
        local_var_req_builder = local_var_req_builder.query(&[("mxn", &local_var_str.to_string())]);
    }
    if let Some(ref local_var_str) = en {
        local_var_req_builder = local_var_req_builder.query(&[("en", &local_var_str.to_string())]);
    }
    if let Some(ref local_var_str) = vn {
        local_var_req_builder = local_var_req_builder.query(&[("vn", &local_var_str.to_string())]);
    }
    if let Some(ref local_var_str) = reduce_selfie {
        local_var_req_builder = local_var_req_builder.query(&[("reduce_selfie", &local_var_str.to_string())]);
    }
    if let Some(ref local_var_str) = custom_generation_range {
        local_var_req_builder = local_var_req_builder.query(&[("custom_generation_range", &local_var_str.to_string())]);
    }
    if let Some(ref local_var_user_agent) = local_var_configuration.user_agent {
        local_var_req_builder = local_var_req_builder.header(reqwest::header::USER_AGENT, local_var_user_agent.clone());
    }

    let local_var_req = local_var_req_builder.build()?;
    let local_var_resp = local_var_client.execute(local_var_req).await?;

    let local_var_status = local_var_resp.status();
    let local_var_content = local_var_resp.text().await?;

    if !local_var_status.is_client_error() && !local_var_status.is_server_error() {
        serde_json::from_str(&local_var_content).map_err(Error::from)
    } else {
        let local_var_entity: Option<GetTimelineError> = serde_json::from_str(&local_var_content).ok();
        let local_var_error = ResponseContent { status: local_var_status, content: local_var_content, entity: local_var_entity };
        Err(Error::ResponseError(local_var_error))
    }
}

pub async fn get_user_timeline(configuration: &configuration::Configuration, user_id: Option<i64>, from_post_id: Option<i64>, post_type: Option<models::PostType>, number: Option<i32>) -> Result<models::PostsResponse, Error<GetUserTimelineError>> {
    let local_var_configuration = configuration;

    let local_var_client = &local_var_configuration.client;

    let local_var_uri_str = format!("{}/v2/posts/user_timeline", local_var_configuration.base_path);
    let mut local_var_req_builder = local_var_client.request(reqwest::Method::GET, local_var_uri_str.as_str());

    if let Some(ref local_var_str) = user_id {
        local_var_req_builder = local_var_req_builder.query(&[("user_id", &local_var_str.to_string())]);
    }
    if let Some(ref local_var_str) = from_post_id {
        local_var_req_builder = local_var_req_builder.query(&[("from_post_id", &local_var_str.to_string())]);
    }
    if let Some(ref local_var_str) = post_type {
        local_var_req_builder = local_var_req_builder.query(&[("post_type", &local_var_str.to_string())]);
    }
    if let Some(ref local_var_str) = number {
        local_var_req_builder = local_var_req_builder.query(&[("number", &local_var_str.to_string())]);
    }
    if let Some(ref local_var_user_agent) = local_var_configuration.user_agent {
        local_var_req_builder = local_var_req_builder.header(reqwest::header::USER_AGENT, local_var_user_agent.clone());
    }

    let local_var_req = local_var_req_builder.build()?;
    let local_var_resp = local_var_client.execute(local_var_req).await?;

    let local_var_status = local_var_resp.status();
    let local_var_content = local_var_resp.text().await?;

    if !local_var_status.is_client_error() && !local_var_status.is_server_error() {
        serde_json::from_str(&local_var_content).map_err(Error::from)
    } else {
        let local_var_entity: Option<GetUserTimelineError> = serde_json::from_str(&local_var_content).ok();
        let local_var_error = ResponseContent { status: local_var_status, content: local_var_content, entity: local_var_entity };
        Err(Error::ResponseError(local_var_error))
    }
}

pub async fn like_posts(configuration: &configuration::Configuration, post_ids: Option<Vec<i64>>) -> Result<models::LikePostsResponse, Error<LikePostsError>> {
    let local_var_configuration = configuration;

    let local_var_client = &local_var_configuration.client;

    let local_var_uri_str = format!("{}/v2/posts/like", local_var_configuration.base_path);
    let mut local_var_req_builder = local_var_client.request(reqwest::Method::POST, local_var_uri_str.as_str());

    if let Some(ref local_var_user_agent) = local_var_configuration.user_agent {
        local_var_req_builder = local_var_req_builder.header(reqwest::header::USER_AGENT, local_var_user_agent.clone());
    }
    let mut local_var_form = reqwest::multipart::Form::new();
    if let Some(local_var_param_value) = post_ids {
        local_var_form = local_var_form.text("post_ids[]", local_var_param_value.into_iter().map(|p| p.to_string()).collect::<Vec<String>>().join(",").to_string());
    }
    local_var_req_builder = local_var_req_builder.multipart(local_var_form);

    let local_var_req = local_var_req_builder.build()?;
    let local_var_resp = local_var_client.execute(local_var_req).await?;

    let local_var_status = local_var_resp.status();
    let local_var_content = local_var_resp.text().await?;

    if !local_var_status.is_client_error() && !local_var_status.is_server_error() {
        serde_json::from_str(&local_var_content).map_err(Error::from)
    } else {
        let local_var_entity: Option<LikePostsError> = serde_json::from_str(&local_var_content).ok();
        let local_var_error = ResponseContent { status: local_var_status, content: local_var_content, entity: local_var_entity };
        Err(Error::ResponseError(local_var_error))
    }
}

pub async fn move_post_to_specific_thread(configuration: &configuration::Configuration, id: i64, thread_id: i64) -> Result<models::ThreadInfo, Error<MovePostToSpecificThreadError>> {
    let local_var_configuration = configuration;

    let local_var_client = &local_var_configuration.client;

    let local_var_uri_str = format!("{}/v3/posts/{id}/move_to_thread/{thread_id}", local_var_configuration.base_path, id=id, thread_id=thread_id);
    let mut local_var_req_builder = local_var_client.request(reqwest::Method::PUT, local_var_uri_str.as_str());

    if let Some(ref local_var_user_agent) = local_var_configuration.user_agent {
        local_var_req_builder = local_var_req_builder.header(reqwest::header::USER_AGENT, local_var_user_agent.clone());
    }

    let local_var_req = local_var_req_builder.build()?;
    let local_var_resp = local_var_client.execute(local_var_req).await?;

    let local_var_status = local_var_resp.status();
    let local_var_content = local_var_resp.text().await?;

    if !local_var_status.is_client_error() && !local_var_status.is_server_error() {
        serde_json::from_str(&local_var_content).map_err(Error::from)
    } else {
        let local_var_entity: Option<MovePostToSpecificThreadError> = serde_json::from_str(&local_var_content).ok();
        let local_var_error = ResponseContent { status: local_var_status, content: local_var_content, entity: local_var_entity };
        Err(Error::ResponseError(local_var_error))
    }
}

pub async fn move_post_to_thread(configuration: &configuration::Configuration, id: i64, thread_icon_filename: Option<&str>, title: Option<&str>) -> Result<models::ThreadInfo, Error<MovePostToThreadError>> {
    let local_var_configuration = configuration;

    let local_var_client = &local_var_configuration.client;

    let local_var_uri_str = format!("{}/v3/posts/{id}/move_to_thread", local_var_configuration.base_path, id=id);
    let mut local_var_req_builder = local_var_client.request(reqwest::Method::POST, local_var_uri_str.as_str());

    if let Some(ref local_var_user_agent) = local_var_configuration.user_agent {
        local_var_req_builder = local_var_req_builder.header(reqwest::header::USER_AGENT, local_var_user_agent.clone());
    }
    let mut local_var_form = reqwest::multipart::Form::new();
    if let Some(local_var_param_value) = thread_icon_filename {
        local_var_form = local_var_form.text("thread_icon_filename", local_var_param_value.to_string());
    }
    if let Some(local_var_param_value) = title {
        local_var_form = local_var_form.text("title", local_var_param_value.to_string());
    }
    local_var_req_builder = local_var_req_builder.multipart(local_var_form);

    let local_var_req = local_var_req_builder.build()?;
    let local_var_resp = local_var_client.execute(local_var_req).await?;

    let local_var_status = local_var_resp.status();
    let local_var_content = local_var_resp.text().await?;

    if !local_var_status.is_client_error() && !local_var_status.is_server_error() {
        serde_json::from_str(&local_var_content).map_err(Error::from)
    } else {
        let local_var_entity: Option<MovePostToThreadError> = serde_json::from_str(&local_var_content).ok();
        let local_var_error = ResponseContent { status: local_var_status, content: local_var_content, entity: local_var_entity };
        Err(Error::ResponseError(local_var_error))
    }
}

pub async fn pin_group_post(configuration: &configuration::Configuration, group_id: Option<i64>, post_id: Option<i64>) -> Result<(), Error<PinGroupPostError>> {
    let local_var_configuration = configuration;

    let local_var_client = &local_var_configuration.client;

    let local_var_uri_str = format!("{}/v2/posts/group_pinned_post", local_var_configuration.base_path);
    let mut local_var_req_builder = local_var_client.request(reqwest::Method::PUT, local_var_uri_str.as_str());

    if let Some(ref local_var_user_agent) = local_var_configuration.user_agent {
        local_var_req_builder = local_var_req_builder.header(reqwest::header::USER_AGENT, local_var_user_agent.clone());
    }
    let mut local_var_form = reqwest::multipart::Form::new();
    if let Some(local_var_param_value) = group_id {
        local_var_form = local_var_form.text("group_id", local_var_param_value.to_string());
    }
    if let Some(local_var_param_value) = post_id {
        local_var_form = local_var_form.text("post_id", local_var_param_value.to_string());
    }
    local_var_req_builder = local_var_req_builder.multipart(local_var_form);

    let local_var_req = local_var_req_builder.build()?;
    let local_var_resp = local_var_client.execute(local_var_req).await?;

    let local_var_status = local_var_resp.status();
    let local_var_content = local_var_resp.text().await?;

    if !local_var_status.is_client_error() && !local_var_status.is_server_error() {
        Ok(())
    } else {
        let local_var_entity: Option<PinGroupPostError> = serde_json::from_str(&local_var_content).ok();
        let local_var_error = ResponseContent { status: local_var_status, content: local_var_content, entity: local_var_entity };
        Err(Error::ResponseError(local_var_error))
    }
}

pub async fn report_post(configuration: &configuration::Configuration, post_id: i64, category_id: Option<i64>, opponent_id: Option<i64>, reason: Option<&str>, screenshot_2_filename: Option<&str>, screenshot_3_filename: Option<&str>, screenshot_4_filename: Option<&str>, screenshot_filename: Option<&str>) -> Result<(), Error<ReportPostError>> {
    let local_var_configuration = configuration;

    let local_var_client = &local_var_configuration.client;

    let local_var_uri_str = format!("{}/v3/posts/{post_id}/report", local_var_configuration.base_path, post_id=post_id);
    let mut local_var_req_builder = local_var_client.request(reqwest::Method::POST, local_var_uri_str.as_str());

    if let Some(ref local_var_user_agent) = local_var_configuration.user_agent {
        local_var_req_builder = local_var_req_builder.header(reqwest::header::USER_AGENT, local_var_user_agent.clone());
    }
    let mut local_var_form = reqwest::multipart::Form::new();
    if let Some(local_var_param_value) = category_id {
        local_var_form = local_var_form.text("category_id", local_var_param_value.to_string());
    }
    if let Some(local_var_param_value) = opponent_id {
        local_var_form = local_var_form.text("opponent_id", local_var_param_value.to_string());
    }
    if let Some(local_var_param_value) = reason {
        local_var_form = local_var_form.text("reason", local_var_param_value.to_string());
    }
    if let Some(local_var_param_value) = screenshot_2_filename {
        local_var_form = local_var_form.text("screenshot_2_filename", local_var_param_value.to_string());
    }
    if let Some(local_var_param_value) = screenshot_3_filename {
        local_var_form = local_var_form.text("screenshot_3_filename", local_var_param_value.to_string());
    }
    if let Some(local_var_param_value) = screenshot_4_filename {
        local_var_form = local_var_form.text("screenshot_4_filename", local_var_param_value.to_string());
    }
    if let Some(local_var_param_value) = screenshot_filename {
        local_var_form = local_var_form.text("screenshot_filename", local_var_param_value.to_string());
    }
    local_var_req_builder = local_var_req_builder.multipart(local_var_form);

    let local_var_req = local_var_req_builder.build()?;
    let local_var_resp = local_var_client.execute(local_var_req).await?;

    let local_var_status = local_var_resp.status();
    let local_var_content = local_var_resp.text().await?;

    if !local_var_status.is_client_error() && !local_var_status.is_server_error() {
        Ok(())
    } else {
        let local_var_entity: Option<ReportPostError> = serde_json::from_str(&local_var_content).ok();
        let local_var_error = ResponseContent { status: local_var_status, content: local_var_content, entity: local_var_entity };
        Err(Error::ResponseError(local_var_error))
    }
}

pub async fn repost(configuration: &configuration::Configuration, x_jwt: Option<&str>, attachment_2_filename: Option<&str>, attachment_3_filename: Option<&str>, attachment_4_filename: Option<&str>, attachment_5_filename: Option<&str>, attachment_6_filename: Option<&str>, attachment_7_filename: Option<&str>, attachment_8_filename: Option<&str>, attachment_9_filename: Option<&str>, attachment_filename: Option<&str>, choices: Option<Vec<String>>, color: Option<i32>, font_size: Option<i32>, group_id: Option<i64>, in_reply_to: Option<i64>, language: Option<&str>, mention_ids: Option<Vec<i64>>, message_tags: Option<serde_json::Value>, post_id: Option<i64>, post_type: Option<models::PostType>, shared_url: Option<serde_json::Value>, text: Option<&str>, video_file_name: Option<&str>) -> Result<models::CreatePostResponse, Error<RepostError>> {
    let local_var_configuration = configuration;

    let local_var_client = &local_var_configuration.client;

    let local_var_uri_str = format!("{}/v3/posts/repost", local_var_configuration.base_path);
    let mut local_var_req_builder = local_var_client.request(reqwest::Method::POST, local_var_uri_str.as_str());

    if let Some(ref local_var_user_agent) = local_var_configuration.user_agent {
        local_var_req_builder = local_var_req_builder.header(reqwest::header::USER_AGENT, local_var_user_agent.clone());
    }
    if let Some(local_var_param_value) = x_jwt {
        local_var_req_builder = local_var_req_builder.header("X-Jwt", local_var_param_value.to_string());
    }
    let mut local_var_form = reqwest::multipart::Form::new();
    if let Some(local_var_param_value) = attachment_2_filename {
        local_var_form = local_var_form.text("attachment_2_filename", local_var_param_value.to_string());
    }
    if let Some(local_var_param_value) = attachment_3_filename {
        local_var_form = local_var_form.text("attachment_3_filename", local_var_param_value.to_string());
    }
    if let Some(local_var_param_value) = attachment_4_filename {
        local_var_form = local_var_form.text("attachment_4_filename", local_var_param_value.to_string());
    }
    if let Some(local_var_param_value) = attachment_5_filename {
        local_var_form = local_var_form.text("attachment_5_filename", local_var_param_value.to_string());
    }
    if let Some(local_var_param_value) = attachment_6_filename {
        local_var_form = local_var_form.text("attachment_6_filename", local_var_param_value.to_string());
    }
    if let Some(local_var_param_value) = attachment_7_filename {
        local_var_form = local_var_form.text("attachment_7_filename", local_var_param_value.to_string());
    }
    if let Some(local_var_param_value) = attachment_8_filename {
        local_var_form = local_var_form.text("attachment_8_filename", local_var_param_value.to_string());
    }
    if let Some(local_var_param_value) = attachment_9_filename {
        local_var_form = local_var_form.text("attachment_9_filename", local_var_param_value.to_string());
    }
    if let Some(local_var_param_value) = attachment_filename {
        local_var_form = local_var_form.text("attachment_filename", local_var_param_value.to_string());
    }
    if let Some(local_var_param_value) = choices {
        local_var_form = local_var_form.text("choices[]", local_var_param_value.into_iter().map(|p| p.to_string()).collect::<Vec<String>>().join(",").to_string());
    }
    if let Some(local_var_param_value) = color {
        local_var_form = local_var_form.text("color", local_var_param_value.to_string());
    }
    if let Some(local_var_param_value) = font_size {
        local_var_form = local_var_form.text("font_size", local_var_param_value.to_string());
    }
    if let Some(local_var_param_value) = group_id {
        local_var_form = local_var_form.text("group_id", local_var_param_value.to_string());
    }
    if let Some(local_var_param_value) = in_reply_to {
        local_var_form = local_var_form.text("in_reply_to", local_var_param_value.to_string());
    }
    if let Some(local_var_param_value) = language {
        local_var_form = local_var_form.text("language", local_var_param_value.to_string());
    }
    if let Some(local_var_param_value) = mention_ids {
        local_var_form = local_var_form.text("mention_ids[]", local_var_param_value.into_iter().map(|p| p.to_string()).collect::<Vec<String>>().join(",").to_string());
    }
    if let Some(local_var_param_value) = message_tags {
        local_var_form = local_var_form.text("message_tags", local_var_param_value.to_string());
    }
    if let Some(local_var_param_value) = post_id {
        local_var_form = local_var_form.text("post_id", local_var_param_value.to_string());
    }
    if let Some(local_var_param_value) = post_type {
        local_var_form = local_var_form.text("post_type", local_var_param_value.to_string());
    }
    if let Some(local_var_param_value) = shared_url {
        local_var_form = local_var_form.text("shared_url", local_var_param_value.to_string());
    }
    if let Some(local_var_param_value) = text {
        local_var_form = local_var_form.text("text", local_var_param_value.to_string());
    }
    if let Some(local_var_param_value) = video_file_name {
        local_var_form = local_var_form.text("video_file_name", local_var_param_value.to_string());
    }
    local_var_req_builder = local_var_req_builder.multipart(local_var_form);

    let local_var_req = local_var_req_builder.build()?;
    let local_var_resp = local_var_client.execute(local_var_req).await?;

    let local_var_status = local_var_resp.status();
    let local_var_content = local_var_resp.text().await?;

    if !local_var_status.is_client_error() && !local_var_status.is_server_error() {
        serde_json::from_str(&local_var_content).map_err(Error::from)
    } else {
        let local_var_entity: Option<RepostError> = serde_json::from_str(&local_var_content).ok();
        let local_var_error = ResponseContent { status: local_var_status, content: local_var_content, entity: local_var_entity };
        Err(Error::ResponseError(local_var_error))
    }
}

pub async fn search_posts(configuration: &configuration::Configuration, keyword: Option<&str>, post_owner_scope: Option<&str>, only_media: Option<bool>, from_post_id: Option<i64>, number: Option<i32>) -> Result<models::PostsResponse, Error<SearchPostsError>> {
    let local_var_configuration = configuration;

    let local_var_client = &local_var_configuration.client;

    let local_var_uri_str = format!("{}/v2/posts/search", local_var_configuration.base_path);
    let mut local_var_req_builder = local_var_client.request(reqwest::Method::GET, local_var_uri_str.as_str());

    if let Some(ref local_var_str) = keyword {
        local_var_req_builder = local_var_req_builder.query(&[("keyword", &local_var_str.to_string())]);
    }
    if let Some(ref local_var_str) = post_owner_scope {
        local_var_req_builder = local_var_req_builder.query(&[("post_owner_scope", &local_var_str.to_string())]);
    }
    if let Some(ref local_var_str) = only_media {
        local_var_req_builder = local_var_req_builder.query(&[("only_media", &local_var_str.to_string())]);
    }
    if let Some(ref local_var_str) = from_post_id {
        local_var_req_builder = local_var_req_builder.query(&[("from_post_id", &local_var_str.to_string())]);
    }
    if let Some(ref local_var_str) = number {
        local_var_req_builder = local_var_req_builder.query(&[("number", &local_var_str.to_string())]);
    }
    if let Some(ref local_var_user_agent) = local_var_configuration.user_agent {
        local_var_req_builder = local_var_req_builder.header(reqwest::header::USER_AGENT, local_var_user_agent.clone());
    }

    let local_var_req = local_var_req_builder.build()?;
    let local_var_resp = local_var_client.execute(local_var_req).await?;

    let local_var_status = local_var_resp.status();
    let local_var_content = local_var_resp.text().await?;

    if !local_var_status.is_client_error() && !local_var_status.is_server_error() {
        serde_json::from_str(&local_var_content).map_err(Error::from)
    } else {
        let local_var_entity: Option<SearchPostsError> = serde_json::from_str(&local_var_content).ok();
        let local_var_error = ResponseContent { status: local_var_status, content: local_var_content, entity: local_var_entity };
        Err(Error::ResponseError(local_var_error))
    }
}

pub async fn unlike_post(configuration: &configuration::Configuration, id: i64) -> Result<(), Error<UnlikePostError>> {
    let local_var_configuration = configuration;

    let local_var_client = &local_var_configuration.client;

    let local_var_uri_str = format!("{}/v1/posts/{id}/unlike", local_var_configuration.base_path, id=id);
    let mut local_var_req_builder = local_var_client.request(reqwest::Method::POST, local_var_uri_str.as_str());

    if let Some(ref local_var_user_agent) = local_var_configuration.user_agent {
        local_var_req_builder = local_var_req_builder.header(reqwest::header::USER_AGENT, local_var_user_agent.clone());
    }

    let local_var_req = local_var_req_builder.build()?;
    let local_var_resp = local_var_client.execute(local_var_req).await?;

    let local_var_status = local_var_resp.status();
    let local_var_content = local_var_resp.text().await?;

    if !local_var_status.is_client_error() && !local_var_status.is_server_error() {
        Ok(())
    } else {
        let local_var_entity: Option<UnlikePostError> = serde_json::from_str(&local_var_content).ok();
        let local_var_error = ResponseContent { status: local_var_status, content: local_var_content, entity: local_var_entity };
        Err(Error::ResponseError(local_var_error))
    }
}

pub async fn unpin_group_post(configuration: &configuration::Configuration, group_id: Option<i64>) -> Result<(), Error<UnpinGroupPostError>> {
    let local_var_configuration = configuration;

    let local_var_client = &local_var_configuration.client;

    let local_var_uri_str = format!("{}/v2/posts/group_pinned_post", local_var_configuration.base_path);
    let mut local_var_req_builder = local_var_client.request(reqwest::Method::DELETE, local_var_uri_str.as_str());

    if let Some(ref local_var_str) = group_id {
        local_var_req_builder = local_var_req_builder.query(&[("group_id", &local_var_str.to_string())]);
    }
    if let Some(ref local_var_user_agent) = local_var_configuration.user_agent {
        local_var_req_builder = local_var_req_builder.header(reqwest::header::USER_AGENT, local_var_user_agent.clone());
    }

    let local_var_req = local_var_req_builder.build()?;
    let local_var_resp = local_var_client.execute(local_var_req).await?;

    let local_var_status = local_var_resp.status();
    let local_var_content = local_var_resp.text().await?;

    if !local_var_status.is_client_error() && !local_var_status.is_server_error() {
        Ok(())
    } else {
        let local_var_entity: Option<UnpinGroupPostError> = serde_json::from_str(&local_var_content).ok();
        let local_var_error = ResponseContent { status: local_var_status, content: local_var_content, entity: local_var_entity };
        Err(Error::ResponseError(local_var_error))
    }
}

pub async fn update_post(configuration: &configuration::Configuration, id: i64, api_key: Option<&str>, color: Option<i32>, font_size: Option<i32>, language: Option<&str>, message_tags: Option<serde_json::Value>, signed_info: Option<&str>, text: Option<&str>, timestamp: Option<i64>, uuid: Option<&str>) -> Result<models::Post, Error<UpdatePostError>> {
    let local_var_configuration = configuration;

    let local_var_client = &local_var_configuration.client;

    let local_var_uri_str = format!("{}/v3/posts/{id}", local_var_configuration.base_path, id=id);
    let mut local_var_req_builder = local_var_client.request(reqwest::Method::PUT, local_var_uri_str.as_str());

    if let Some(ref local_var_user_agent) = local_var_configuration.user_agent {
        local_var_req_builder = local_var_req_builder.header(reqwest::header::USER_AGENT, local_var_user_agent.clone());
    }
    let mut local_var_form = reqwest::multipart::Form::new();
    if let Some(local_var_param_value) = api_key {
        local_var_form = local_var_form.text("api_key", local_var_param_value.to_string());
    }
    if let Some(local_var_param_value) = color {
        local_var_form = local_var_form.text("color", local_var_param_value.to_string());
    }
    if let Some(local_var_param_value) = font_size {
        local_var_form = local_var_form.text("font_size", local_var_param_value.to_string());
    }
    if let Some(local_var_param_value) = language {
        local_var_form = local_var_form.text("language", local_var_param_value.to_string());
    }
    if let Some(local_var_param_value) = message_tags {
        local_var_form = local_var_form.text("message_tags", local_var_param_value.to_string());
    }
    if let Some(local_var_param_value) = signed_info {
        local_var_form = local_var_form.text("signed_info", local_var_param_value.to_string());
    }
    if let Some(local_var_param_value) = text {
        local_var_form = local_var_form.text("text", local_var_param_value.to_string());
    }
    if let Some(local_var_param_value) = timestamp {
        local_var_form = local_var_form.text("timestamp", local_var_param_value.to_string());
    }
    if let Some(local_var_param_value) = uuid {
        local_var_form = local_var_form.text("uuid", local_var_param_value.to_string());
    }
    local_var_req_builder = local_var_req_builder.multipart(local_var_form);

    let local_var_req = local_var_req_builder.build()?;
    let local_var_resp = local_var_client.execute(local_var_req).await?;

    let local_var_status = local_var_resp.status();
    let local_var_content = local_var_resp.text().await?;

    if !local_var_status.is_client_error() && !local_var_status.is_server_error() {
        serde_json::from_str(&local_var_content).map_err(Error::from)
    } else {
        let local_var_entity: Option<UpdatePostError> = serde_json::from_str(&local_var_content).ok();
        let local_var_error = ResponseContent { status: local_var_status, content: local_var_content, entity: local_var_entity };
        Err(Error::ResponseError(local_var_error))
    }
}

pub async fn validate_post(configuration: &configuration::Configuration, group_id: Option<i64>, text: Option<&str>, thread_id: Option<i64>) -> Result<models::ValidationPostResponse, Error<ValidatePostError>> {
    let local_var_configuration = configuration;

    let local_var_client = &local_var_configuration.client;

    let local_var_uri_str = format!("{}/v1/posts/validate", local_var_configuration.base_path);
    let mut local_var_req_builder = local_var_client.request(reqwest::Method::POST, local_var_uri_str.as_str());

    if let Some(ref local_var_user_agent) = local_var_configuration.user_agent {
        local_var_req_builder = local_var_req_builder.header(reqwest::header::USER_AGENT, local_var_user_agent.clone());
    }
    let mut local_var_form = reqwest::multipart::Form::new();
    if let Some(local_var_param_value) = group_id {
        local_var_form = local_var_form.text("group_id", local_var_param_value.to_string());
    }
    if let Some(local_var_param_value) = text {
        local_var_form = local_var_form.text("text", local_var_param_value.to_string());
    }
    if let Some(local_var_param_value) = thread_id {
        local_var_form = local_var_form.text("thread_id", local_var_param_value.to_string());
    }
    local_var_req_builder = local_var_req_builder.multipart(local_var_form);

    let local_var_req = local_var_req_builder.build()?;
    let local_var_resp = local_var_client.execute(local_var_req).await?;

    let local_var_status = local_var_resp.status();
    let local_var_content = local_var_resp.text().await?;

    if !local_var_status.is_client_error() && !local_var_status.is_server_error() {
        serde_json::from_str(&local_var_content).map_err(Error::from)
    } else {
        let local_var_entity: Option<ValidatePostError> = serde_json::from_str(&local_var_content).ok();
        let local_var_error = ResponseContent { status: local_var_status, content: local_var_content, entity: local_var_entity };
        Err(Error::ResponseError(local_var_error))
    }
}

pub async fn view_post_video(configuration: &configuration::Configuration, video_id: i64) -> Result<(), Error<ViewPostVideoError>> {
    let local_var_configuration = configuration;

    let local_var_client = &local_var_configuration.client;

    let local_var_uri_str = format!("{}/v1/posts/videos/{videoId}/view", local_var_configuration.base_path, videoId=video_id);
    let mut local_var_req_builder = local_var_client.request(reqwest::Method::POST, local_var_uri_str.as_str());

    if let Some(ref local_var_user_agent) = local_var_configuration.user_agent {
        local_var_req_builder = local_var_req_builder.header(reqwest::header::USER_AGENT, local_var_user_agent.clone());
    }

    let local_var_req = local_var_req_builder.build()?;
    let local_var_resp = local_var_client.execute(local_var_req).await?;

    let local_var_status = local_var_resp.status();
    let local_var_content = local_var_resp.text().await?;

    if !local_var_status.is_client_error() && !local_var_status.is_server_error() {
        Ok(())
    } else {
        let local_var_entity: Option<ViewPostVideoError> = serde_json::from_str(&local_var_content).ok();
        let local_var_error = ResponseContent { status: local_var_status, content: local_var_content, entity: local_var_entity };
        Err(Error::ResponseError(local_var_error))
    }
}

