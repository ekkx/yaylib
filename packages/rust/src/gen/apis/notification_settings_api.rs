// Code generated from the Yay! API schema; DO NOT EDIT.


use reqwest;
use serde::{Deserialize, Serialize};
use crate::gen::{apis::ResponseContent, models};
use super::{Error, configuration};


/// struct for typed errors of method [`get_group_notification_settings`]
#[derive(Debug, Clone, Serialize, Deserialize)]
#[serde(untagged)]
pub enum GetGroupNotificationSettingsError {
    UnknownValue(serde_json::Value),
}

/// struct for typed errors of method [`update_chat_room_notification_settings`]
#[derive(Debug, Clone, Serialize, Deserialize)]
#[serde(untagged)]
pub enum UpdateChatRoomNotificationSettingsError {
    UnknownValue(serde_json::Value),
}

/// struct for typed errors of method [`update_group_notification_settings`]
#[derive(Debug, Clone, Serialize, Deserialize)]
#[serde(untagged)]
pub enum UpdateGroupNotificationSettingsError {
    UnknownValue(serde_json::Value),
}


pub async fn get_group_notification_settings(configuration: &configuration::Configuration, id: i64) -> Result<models::GroupNotificationSettingsResponse, Error<GetGroupNotificationSettingsError>> {
    let local_var_configuration = configuration;

    let local_var_client = &local_var_configuration.client;

    let local_var_uri_str = format!("{}/v2/notification_settings/groups/{id}", local_var_configuration.base_path, id=id);
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
        let local_var_entity: Option<GetGroupNotificationSettingsError> = serde_json::from_str(&local_var_content).ok();
        let local_var_error = ResponseContent { status: local_var_status, content: local_var_content, entity: local_var_entity };
        Err(Error::ResponseError(local_var_error))
    }
}

pub async fn update_chat_room_notification_settings(configuration: &configuration::Configuration, id: i64, notification_chat: Option<i32>) -> Result<models::NotificationSettingResponse, Error<UpdateChatRoomNotificationSettingsError>> {
    let local_var_configuration = configuration;

    let local_var_client = &local_var_configuration.client;

    let local_var_uri_str = format!("{}/v2/notification_settings/chat_rooms/{id}", local_var_configuration.base_path, id=id);
    let mut local_var_req_builder = local_var_client.request(reqwest::Method::POST, local_var_uri_str.as_str());

    if let Some(ref local_var_user_agent) = local_var_configuration.user_agent {
        local_var_req_builder = local_var_req_builder.header(reqwest::header::USER_AGENT, local_var_user_agent.clone());
    }
    let mut local_var_form = reqwest::multipart::Form::new();
    if let Some(local_var_param_value) = notification_chat {
        local_var_form = local_var_form.text("notification_chat", local_var_param_value.to_string());
    }
    local_var_req_builder = local_var_req_builder.multipart(local_var_form);

    let local_var_req = local_var_req_builder.build()?;
    let local_var_resp = local_var_client.execute(local_var_req).await?;

    let local_var_status = local_var_resp.status();
    let local_var_content = local_var_resp.text().await?;

    if !local_var_status.is_client_error() && !local_var_status.is_server_error() {
        serde_json::from_str(&local_var_content).map_err(Error::from)
    } else {
        let local_var_entity: Option<UpdateChatRoomNotificationSettingsError> = serde_json::from_str(&local_var_content).ok();
        let local_var_error = ResponseContent { status: local_var_status, content: local_var_content, entity: local_var_entity };
        Err(Error::ResponseError(local_var_error))
    }
}

pub async fn update_group_notification_settings(configuration: &configuration::Configuration, id: i64, notification_group_join: Option<i32>, notification_group_message_tag_all: Option<i32>, notification_group_post: Option<i32>, notification_group_request: Option<i32>) -> Result<models::AdditionalSettingsResponse, Error<UpdateGroupNotificationSettingsError>> {
    let local_var_configuration = configuration;

    let local_var_client = &local_var_configuration.client;

    let local_var_uri_str = format!("{}/v2/notification_settings/groups/{id}", local_var_configuration.base_path, id=id);
    let mut local_var_req_builder = local_var_client.request(reqwest::Method::POST, local_var_uri_str.as_str());

    if let Some(ref local_var_user_agent) = local_var_configuration.user_agent {
        local_var_req_builder = local_var_req_builder.header(reqwest::header::USER_AGENT, local_var_user_agent.clone());
    }
    let mut local_var_form = reqwest::multipart::Form::new();
    if let Some(local_var_param_value) = notification_group_join {
        local_var_form = local_var_form.text("notification_group_join", local_var_param_value.to_string());
    }
    if let Some(local_var_param_value) = notification_group_message_tag_all {
        local_var_form = local_var_form.text("notification_group_message_tag_all", local_var_param_value.to_string());
    }
    if let Some(local_var_param_value) = notification_group_post {
        local_var_form = local_var_form.text("notification_group_post", local_var_param_value.to_string());
    }
    if let Some(local_var_param_value) = notification_group_request {
        local_var_form = local_var_form.text("notification_group_request", local_var_param_value.to_string());
    }
    local_var_req_builder = local_var_req_builder.multipart(local_var_form);

    let local_var_req = local_var_req_builder.build()?;
    let local_var_resp = local_var_client.execute(local_var_req).await?;

    let local_var_status = local_var_resp.status();
    let local_var_content = local_var_resp.text().await?;

    if !local_var_status.is_client_error() && !local_var_status.is_server_error() {
        serde_json::from_str(&local_var_content).map_err(Error::from)
    } else {
        let local_var_entity: Option<UpdateGroupNotificationSettingsError> = serde_json::from_str(&local_var_content).ok();
        let local_var_error = ResponseContent { status: local_var_status, content: local_var_content, entity: local_var_entity };
        Err(Error::ResponseError(local_var_error))
    }
}

