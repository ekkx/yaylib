// Code generated from the Yay! API schema; DO NOT EDIT.


use reqwest;
use serde::{Deserialize, Serialize};
use crate::gen::{apis::ResponseContent, models};
use super::{Error, configuration};


/// struct for typed errors of method [`add_thread_member`]
#[derive(Debug, Clone, Serialize, Deserialize)]
#[serde(untagged)]
pub enum AddThreadMemberError {
    UnknownValue(serde_json::Value),
}

/// struct for typed errors of method [`create_thread`]
#[derive(Debug, Clone, Serialize, Deserialize)]
#[serde(untagged)]
pub enum CreateThreadError {
    UnknownValue(serde_json::Value),
}

/// struct for typed errors of method [`create_thread_post`]
#[derive(Debug, Clone, Serialize, Deserialize)]
#[serde(untagged)]
pub enum CreateThreadPostError {
    UnknownValue(serde_json::Value),
}

/// struct for typed errors of method [`delete_thread`]
#[derive(Debug, Clone, Serialize, Deserialize)]
#[serde(untagged)]
pub enum DeleteThreadError {
    UnknownValue(serde_json::Value),
}

/// struct for typed errors of method [`get_joined_thread_statuses`]
#[derive(Debug, Clone, Serialize, Deserialize)]
#[serde(untagged)]
pub enum GetJoinedThreadStatusesError {
    UnknownValue(serde_json::Value),
}

/// struct for typed errors of method [`get_thread`]
#[derive(Debug, Clone, Serialize, Deserialize)]
#[serde(untagged)]
pub enum GetThreadError {
    UnknownValue(serde_json::Value),
}

/// struct for typed errors of method [`get_thread_posts`]
#[derive(Debug, Clone, Serialize, Deserialize)]
#[serde(untagged)]
pub enum GetThreadPostsError {
    UnknownValue(serde_json::Value),
}

/// struct for typed errors of method [`list_threads`]
#[derive(Debug, Clone, Serialize, Deserialize)]
#[serde(untagged)]
pub enum ListThreadsError {
    UnknownValue(serde_json::Value),
}

/// struct for typed errors of method [`remove_thread_member`]
#[derive(Debug, Clone, Serialize, Deserialize)]
#[serde(untagged)]
pub enum RemoveThreadMemberError {
    UnknownValue(serde_json::Value),
}

/// struct for typed errors of method [`update_thread`]
#[derive(Debug, Clone, Serialize, Deserialize)]
#[serde(untagged)]
pub enum UpdateThreadError {
    UnknownValue(serde_json::Value),
}


pub async fn add_thread_member(configuration: &configuration::Configuration, thread_id: i64, id: i64) -> Result<(), Error<AddThreadMemberError>> {
    let local_var_configuration = configuration;

    let local_var_client = &local_var_configuration.client;

    let local_var_uri_str = format!("{}/v1/threads/{thread_id}/members/{id}", local_var_configuration.base_path, thread_id=thread_id, id=id);
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
        let local_var_entity: Option<AddThreadMemberError> = serde_json::from_str(&local_var_content).ok();
        let local_var_error = ResponseContent { status: local_var_status, content: local_var_content, entity: local_var_entity };
        Err(Error::ResponseError(local_var_error))
    }
}

pub async fn create_thread(configuration: &configuration::Configuration, create_group_thread_request: models::CreateGroupThreadRequest) -> Result<models::ThreadInfo, Error<CreateThreadError>> {
    let local_var_configuration = configuration;

    let local_var_client = &local_var_configuration.client;

    let local_var_uri_str = format!("{}/v1/threads", local_var_configuration.base_path);
    let mut local_var_req_builder = local_var_client.request(reqwest::Method::POST, local_var_uri_str.as_str());

    if let Some(ref local_var_user_agent) = local_var_configuration.user_agent {
        local_var_req_builder = local_var_req_builder.header(reqwest::header::USER_AGENT, local_var_user_agent.clone());
    }
    local_var_req_builder = local_var_req_builder.json(&create_group_thread_request);

    let local_var_req = local_var_req_builder.build()?;
    let local_var_resp = local_var_client.execute(local_var_req).await?;

    let local_var_status = local_var_resp.status();
    let local_var_content = local_var_resp.text().await?;

    if !local_var_status.is_client_error() && !local_var_status.is_server_error() {
        serde_json::from_str(&local_var_content).map_err(Error::from)
    } else {
        let local_var_entity: Option<CreateThreadError> = serde_json::from_str(&local_var_content).ok();
        let local_var_error = ResponseContent { status: local_var_status, content: local_var_content, entity: local_var_entity };
        Err(Error::ResponseError(local_var_error))
    }
}

pub async fn create_thread_post(configuration: &configuration::Configuration, id: i64, x_jwt: Option<&str>, attachment_2_filename: Option<&str>, attachment_3_filename: Option<&str>, attachment_4_filename: Option<&str>, attachment_5_filename: Option<&str>, attachment_6_filename: Option<&str>, attachment_7_filename: Option<&str>, attachment_8_filename: Option<&str>, attachment_9_filename: Option<&str>, attachment_filename: Option<&str>, choices: Option<Vec<String>>, color: Option<i32>, font_size: Option<i32>, group_id: Option<i64>, in_reply_to: Option<i64>, language: Option<&str>, mention_ids: Option<Vec<i64>>, message_tags: Option<serde_json::Value>, post_type: Option<models::PostType>, shared_url: Option<serde_json::Value>, text: Option<&str>, video_file_name: Option<&str>) -> Result<models::Post, Error<CreateThreadPostError>> {
    let local_var_configuration = configuration;

    let local_var_client = &local_var_configuration.client;

    let local_var_uri_str = format!("{}/v1/threads/{id}/posts", local_var_configuration.base_path, id=id);
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
        let local_var_entity: Option<CreateThreadPostError> = serde_json::from_str(&local_var_content).ok();
        let local_var_error = ResponseContent { status: local_var_status, content: local_var_content, entity: local_var_entity };
        Err(Error::ResponseError(local_var_error))
    }
}

pub async fn delete_thread(configuration: &configuration::Configuration, id: i64) -> Result<(), Error<DeleteThreadError>> {
    let local_var_configuration = configuration;

    let local_var_client = &local_var_configuration.client;

    let local_var_uri_str = format!("{}/v1/threads/{id}", local_var_configuration.base_path, id=id);
    let mut local_var_req_builder = local_var_client.request(reqwest::Method::DELETE, local_var_uri_str.as_str());

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
        let local_var_entity: Option<DeleteThreadError> = serde_json::from_str(&local_var_content).ok();
        let local_var_error = ResponseContent { status: local_var_status, content: local_var_content, entity: local_var_entity };
        Err(Error::ResponseError(local_var_error))
    }
}

pub async fn get_joined_thread_statuses(configuration: &configuration::Configuration, ids: Option<Vec<i64>>) -> Result<std::collections::HashMap<String, String>, Error<GetJoinedThreadStatusesError>> {
    let local_var_configuration = configuration;

    let local_var_client = &local_var_configuration.client;

    let local_var_uri_str = format!("{}/v1/threads/joined_statuses", local_var_configuration.base_path);
    let mut local_var_req_builder = local_var_client.request(reqwest::Method::GET, local_var_uri_str.as_str());

    if let Some(ref local_var_str) = ids {
        local_var_req_builder = match "multi" {
            "multi" => local_var_req_builder.query(&local_var_str.into_iter().map(|p| ("ids[]".to_owned(), p.to_string())).collect::<Vec<(std::string::String, std::string::String)>>()),
            _ => local_var_req_builder.query(&[("ids[]", &local_var_str.into_iter().map(|p| p.to_string()).collect::<Vec<String>>().join(",").to_string())]),
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
        let local_var_entity: Option<GetJoinedThreadStatusesError> = serde_json::from_str(&local_var_content).ok();
        let local_var_error = ResponseContent { status: local_var_status, content: local_var_content, entity: local_var_entity };
        Err(Error::ResponseError(local_var_error))
    }
}

pub async fn get_thread(configuration: &configuration::Configuration, id: i64) -> Result<models::ThreadInfo, Error<GetThreadError>> {
    let local_var_configuration = configuration;

    let local_var_client = &local_var_configuration.client;

    let local_var_uri_str = format!("{}/v1/threads/{id}", local_var_configuration.base_path, id=id);
    let mut local_var_req_builder = local_var_client.request(reqwest::Method::GET, local_var_uri_str.as_str());

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
        let local_var_entity: Option<GetThreadError> = serde_json::from_str(&local_var_content).ok();
        let local_var_error = ResponseContent { status: local_var_status, content: local_var_content, entity: local_var_entity };
        Err(Error::ResponseError(local_var_error))
    }
}

pub async fn get_thread_posts(configuration: &configuration::Configuration, id: i64, post_type: Option<models::PostType>, number: Option<i32>, from: Option<i64>) -> Result<models::PostsResponse, Error<GetThreadPostsError>> {
    let local_var_configuration = configuration;

    let local_var_client = &local_var_configuration.client;

    let local_var_uri_str = format!("{}/v1/threads/{id}/posts", local_var_configuration.base_path, id=id);
    let mut local_var_req_builder = local_var_client.request(reqwest::Method::GET, local_var_uri_str.as_str());

    if let Some(ref local_var_str) = post_type {
        local_var_req_builder = local_var_req_builder.query(&[("post_type", &local_var_str.to_string())]);
    }
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
        let local_var_entity: Option<GetThreadPostsError> = serde_json::from_str(&local_var_content).ok();
        let local_var_error = ResponseContent { status: local_var_status, content: local_var_content, entity: local_var_entity };
        Err(Error::ResponseError(local_var_error))
    }
}

pub async fn list_threads(configuration: &configuration::Configuration, group_id: Option<i64>, from: Option<&str>, join_status: Option<&str>) -> Result<models::GroupThreadListResponse, Error<ListThreadsError>> {
    let local_var_configuration = configuration;

    let local_var_client = &local_var_configuration.client;

    let local_var_uri_str = format!("{}/v1/threads", local_var_configuration.base_path);
    let mut local_var_req_builder = local_var_client.request(reqwest::Method::GET, local_var_uri_str.as_str());

    if let Some(ref local_var_str) = group_id {
        local_var_req_builder = local_var_req_builder.query(&[("group_id", &local_var_str.to_string())]);
    }
    if let Some(ref local_var_str) = from {
        local_var_req_builder = local_var_req_builder.query(&[("from", &local_var_str.to_string())]);
    }
    if let Some(ref local_var_str) = join_status {
        local_var_req_builder = local_var_req_builder.query(&[("join_status", &local_var_str.to_string())]);
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
        let local_var_entity: Option<ListThreadsError> = serde_json::from_str(&local_var_content).ok();
        let local_var_error = ResponseContent { status: local_var_status, content: local_var_content, entity: local_var_entity };
        Err(Error::ResponseError(local_var_error))
    }
}

pub async fn remove_thread_member(configuration: &configuration::Configuration, thread_id: i64, id: i64) -> Result<(), Error<RemoveThreadMemberError>> {
    let local_var_configuration = configuration;

    let local_var_client = &local_var_configuration.client;

    let local_var_uri_str = format!("{}/v1/threads/{thread_id}/members/{id}", local_var_configuration.base_path, thread_id=thread_id, id=id);
    let mut local_var_req_builder = local_var_client.request(reqwest::Method::DELETE, local_var_uri_str.as_str());

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
        let local_var_entity: Option<RemoveThreadMemberError> = serde_json::from_str(&local_var_content).ok();
        let local_var_error = ResponseContent { status: local_var_status, content: local_var_content, entity: local_var_entity };
        Err(Error::ResponseError(local_var_error))
    }
}

pub async fn update_thread(configuration: &configuration::Configuration, id: i64, thread_icon_filename: Option<&str>, title: Option<&str>) -> Result<(), Error<UpdateThreadError>> {
    let local_var_configuration = configuration;

    let local_var_client = &local_var_configuration.client;

    let local_var_uri_str = format!("{}/v1/threads/{id}", local_var_configuration.base_path, id=id);
    let mut local_var_req_builder = local_var_client.request(reqwest::Method::PUT, local_var_uri_str.as_str());

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
        Ok(())
    } else {
        let local_var_entity: Option<UpdateThreadError> = serde_json::from_str(&local_var_content).ok();
        let local_var_error = ResponseContent { status: local_var_status, content: local_var_content, entity: local_var_entity };
        Err(Error::ResponseError(local_var_error))
    }
}

