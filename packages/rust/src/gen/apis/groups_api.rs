// Code generated from the Yay! API schema; DO NOT EDIT.


use reqwest;
use serde::{Deserialize, Serialize};
use crate::gen::{apis::ResponseContent, models};
use super::{Error, configuration};


/// struct for typed errors of method [`accept_group_join_request`]
#[derive(Debug, Clone, Serialize, Deserialize)]
#[serde(untagged)]
pub enum AcceptGroupJoinRequestError {
    UnknownValue(serde_json::Value),
}

/// struct for typed errors of method [`accept_group_transfer`]
#[derive(Debug, Clone, Serialize, Deserialize)]
#[serde(untagged)]
pub enum AcceptGroupTransferError {
    UnknownValue(serde_json::Value),
}

/// struct for typed errors of method [`ban_group_user`]
#[derive(Debug, Clone, Serialize, Deserialize)]
#[serde(untagged)]
pub enum BanGroupUserError {
    UnknownValue(serde_json::Value),
}

/// struct for typed errors of method [`cancel_group_transfer`]
#[derive(Debug, Clone, Serialize, Deserialize)]
#[serde(untagged)]
pub enum CancelGroupTransferError {
    UnknownValue(serde_json::Value),
}

/// struct for typed errors of method [`create_group`]
#[derive(Debug, Clone, Serialize, Deserialize)]
#[serde(untagged)]
pub enum CreateGroupError {
    UnknownValue(serde_json::Value),
}

/// struct for typed errors of method [`decline_group_join_request`]
#[derive(Debug, Clone, Serialize, Deserialize)]
#[serde(untagged)]
pub enum DeclineGroupJoinRequestError {
    UnknownValue(serde_json::Value),
}

/// struct for typed errors of method [`deputize_group_users`]
#[derive(Debug, Clone, Serialize, Deserialize)]
#[serde(untagged)]
pub enum DeputizeGroupUsersError {
    UnknownValue(serde_json::Value),
}

/// struct for typed errors of method [`deputize_group_users_mass`]
#[derive(Debug, Clone, Serialize, Deserialize)]
#[serde(untagged)]
pub enum DeputizeGroupUsersMassError {
    UnknownValue(serde_json::Value),
}

/// struct for typed errors of method [`fire_group_user`]
#[derive(Debug, Clone, Serialize, Deserialize)]
#[serde(untagged)]
pub enum FireGroupUserError {
    UnknownValue(serde_json::Value),
}

/// struct for typed errors of method [`get_community_campaign_group_leaderboard`]
#[derive(Debug, Clone, Serialize, Deserialize)]
#[serde(untagged)]
pub enum GetCommunityCampaignGroupLeaderboardError {
    UnknownValue(serde_json::Value),
}

/// struct for typed errors of method [`get_community_campaign_status`]
#[derive(Debug, Clone, Serialize, Deserialize)]
#[serde(untagged)]
pub enum GetCommunityCampaignStatusError {
    UnknownValue(serde_json::Value),
}

/// struct for typed errors of method [`get_community_campaign_user_leaderboard`]
#[derive(Debug, Clone, Serialize, Deserialize)]
#[serde(untagged)]
pub enum GetCommunityCampaignUserLeaderboardError {
    UnknownValue(serde_json::Value),
}

/// struct for typed errors of method [`get_group`]
#[derive(Debug, Clone, Serialize, Deserialize)]
#[serde(untagged)]
pub enum GetGroupError {
    UnknownValue(serde_json::Value),
}

/// struct for typed errors of method [`get_group_ban_list`]
#[derive(Debug, Clone, Serialize, Deserialize)]
#[serde(untagged)]
pub enum GetGroupBanListError {
    UnknownValue(serde_json::Value),
}

/// struct for typed errors of method [`get_group_create_quota`]
#[derive(Debug, Clone, Serialize, Deserialize)]
#[serde(untagged)]
pub enum GetGroupCreateQuotaError {
    UnknownValue(serde_json::Value),
}

/// struct for typed errors of method [`get_group_gift_history`]
#[derive(Debug, Clone, Serialize, Deserialize)]
#[serde(untagged)]
pub enum GetGroupGiftHistoryError {
    UnknownValue(serde_json::Value),
}

/// struct for typed errors of method [`get_group_gift_transactions`]
#[derive(Debug, Clone, Serialize, Deserialize)]
#[serde(untagged)]
pub enum GetGroupGiftTransactionsError {
    UnknownValue(serde_json::Value),
}

/// struct for typed errors of method [`get_group_highlights`]
#[derive(Debug, Clone, Serialize, Deserialize)]
#[serde(untagged)]
pub enum GetGroupHighlightsError {
    UnknownValue(serde_json::Value),
}

/// struct for typed errors of method [`get_group_member`]
#[derive(Debug, Clone, Serialize, Deserialize)]
#[serde(untagged)]
pub enum GetGroupMemberError {
    UnknownValue(serde_json::Value),
}

/// struct for typed errors of method [`get_group_members`]
#[derive(Debug, Clone, Serialize, Deserialize)]
#[serde(untagged)]
pub enum GetGroupMembersError {
    UnknownValue(serde_json::Value),
}

/// struct for typed errors of method [`get_group_received_gift_senders`]
#[derive(Debug, Clone, Serialize, Deserialize)]
#[serde(untagged)]
pub enum GetGroupReceivedGiftSendersError {
    UnknownValue(serde_json::Value),
}

/// struct for typed errors of method [`get_group_unread_status`]
#[derive(Debug, Clone, Serialize, Deserialize)]
#[serde(untagged)]
pub enum GetGroupUnreadStatusError {
    UnknownValue(serde_json::Value),
}

/// struct for typed errors of method [`get_invitable_group_users`]
#[derive(Debug, Clone, Serialize, Deserialize)]
#[serde(untagged)]
pub enum GetInvitableGroupUsersError {
    UnknownValue(serde_json::Value),
}

/// struct for typed errors of method [`get_joined_group_statuses`]
#[derive(Debug, Clone, Serialize, Deserialize)]
#[serde(untagged)]
pub enum GetJoinedGroupStatusesError {
    UnknownValue(serde_json::Value),
}

/// struct for typed errors of method [`get_relatable_groups`]
#[derive(Debug, Clone, Serialize, Deserialize)]
#[serde(untagged)]
pub enum GetRelatableGroupsError {
    UnknownValue(serde_json::Value),
}

/// struct for typed errors of method [`get_related_groups`]
#[derive(Debug, Clone, Serialize, Deserialize)]
#[serde(untagged)]
pub enum GetRelatedGroupsError {
    UnknownValue(serde_json::Value),
}

/// struct for typed errors of method [`get_user_group_list`]
#[derive(Debug, Clone, Serialize, Deserialize)]
#[serde(untagged)]
pub enum GetUserGroupListError {
    UnknownValue(serde_json::Value),
}

/// struct for typed errors of method [`invite_to_group`]
#[derive(Debug, Clone, Serialize, Deserialize)]
#[serde(untagged)]
pub enum InviteToGroupError {
    UnknownValue(serde_json::Value),
}

/// struct for typed errors of method [`join_group`]
#[derive(Debug, Clone, Serialize, Deserialize)]
#[serde(untagged)]
pub enum JoinGroupError {
    UnknownValue(serde_json::Value),
}

/// struct for typed errors of method [`leave_group`]
#[derive(Debug, Clone, Serialize, Deserialize)]
#[serde(untagged)]
pub enum LeaveGroupError {
    UnknownValue(serde_json::Value),
}

/// struct for typed errors of method [`list_group_categories`]
#[derive(Debug, Clone, Serialize, Deserialize)]
#[serde(untagged)]
pub enum ListGroupCategoriesError {
    UnknownValue(serde_json::Value),
}

/// struct for typed errors of method [`list_groups`]
#[derive(Debug, Clone, Serialize, Deserialize)]
#[serde(untagged)]
pub enum ListGroupsError {
    UnknownValue(serde_json::Value),
}

/// struct for typed errors of method [`list_my_groups`]
#[derive(Debug, Clone, Serialize, Deserialize)]
#[serde(untagged)]
pub enum ListMyGroupsError {
    UnknownValue(serde_json::Value),
}

/// struct for typed errors of method [`pin_group_highlight_post`]
#[derive(Debug, Clone, Serialize, Deserialize)]
#[serde(untagged)]
pub enum PinGroupHighlightPostError {
    UnknownValue(serde_json::Value),
}

/// struct for typed errors of method [`remove_group_cover`]
#[derive(Debug, Clone, Serialize, Deserialize)]
#[serde(untagged)]
pub enum RemoveGroupCoverError {
    UnknownValue(serde_json::Value),
}

/// struct for typed errors of method [`remove_group_deputies`]
#[derive(Debug, Clone, Serialize, Deserialize)]
#[serde(untagged)]
pub enum RemoveGroupDeputiesError {
    UnknownValue(serde_json::Value),
}

/// struct for typed errors of method [`remove_group_icon`]
#[derive(Debug, Clone, Serialize, Deserialize)]
#[serde(untagged)]
pub enum RemoveGroupIconError {
    UnknownValue(serde_json::Value),
}

/// struct for typed errors of method [`remove_related_groups`]
#[derive(Debug, Clone, Serialize, Deserialize)]
#[serde(untagged)]
pub enum RemoveRelatedGroupsError {
    UnknownValue(serde_json::Value),
}

/// struct for typed errors of method [`report_group`]
#[derive(Debug, Clone, Serialize, Deserialize)]
#[serde(untagged)]
pub enum ReportGroupError {
    UnknownValue(serde_json::Value),
}

/// struct for typed errors of method [`request_group_walkthrough`]
#[derive(Debug, Clone, Serialize, Deserialize)]
#[serde(untagged)]
pub enum RequestGroupWalkthroughError {
    UnknownValue(serde_json::Value),
}

/// struct for typed errors of method [`search_group_posts`]
#[derive(Debug, Clone, Serialize, Deserialize)]
#[serde(untagged)]
pub enum SearchGroupPostsError {
    UnknownValue(serde_json::Value),
}

/// struct for typed errors of method [`set_group_title`]
#[derive(Debug, Clone, Serialize, Deserialize)]
#[serde(untagged)]
pub enum SetGroupTitleError {
    UnknownValue(serde_json::Value),
}

/// struct for typed errors of method [`take_over_group`]
#[derive(Debug, Clone, Serialize, Deserialize)]
#[serde(untagged)]
pub enum TakeOverGroupError {
    UnknownValue(serde_json::Value),
}

/// struct for typed errors of method [`transfer_group`]
#[derive(Debug, Clone, Serialize, Deserialize)]
#[serde(untagged)]
pub enum TransferGroupError {
    UnknownValue(serde_json::Value),
}

/// struct for typed errors of method [`unban_group_user`]
#[derive(Debug, Clone, Serialize, Deserialize)]
#[serde(untagged)]
pub enum UnbanGroupUserError {
    UnknownValue(serde_json::Value),
}

/// struct for typed errors of method [`unpin_group_highlight_post`]
#[derive(Debug, Clone, Serialize, Deserialize)]
#[serde(untagged)]
pub enum UnpinGroupHighlightPostError {
    UnknownValue(serde_json::Value),
}

/// struct for typed errors of method [`update_group`]
#[derive(Debug, Clone, Serialize, Deserialize)]
#[serde(untagged)]
pub enum UpdateGroupError {
    UnknownValue(serde_json::Value),
}

/// struct for typed errors of method [`update_related_groups`]
#[derive(Debug, Clone, Serialize, Deserialize)]
#[serde(untagged)]
pub enum UpdateRelatedGroupsError {
    UnknownValue(serde_json::Value),
}

/// struct for typed errors of method [`visit_group`]
#[derive(Debug, Clone, Serialize, Deserialize)]
#[serde(untagged)]
pub enum VisitGroupError {
    UnknownValue(serde_json::Value),
}

/// struct for typed errors of method [`withdraw_group_deputy`]
#[derive(Debug, Clone, Serialize, Deserialize)]
#[serde(untagged)]
pub enum WithdrawGroupDeputyError {
    UnknownValue(serde_json::Value),
}

/// struct for typed errors of method [`withdraw_group_transfer`]
#[derive(Debug, Clone, Serialize, Deserialize)]
#[serde(untagged)]
pub enum WithdrawGroupTransferError {
    UnknownValue(serde_json::Value),
}


pub async fn accept_group_join_request(configuration: &configuration::Configuration, id: i64, user_id: i64) -> Result<(), Error<AcceptGroupJoinRequestError>> {
    let local_var_configuration = configuration;

    let local_var_client = &local_var_configuration.client;

    let local_var_uri_str = format!("{}/v1/groups/{id}/accept/{userId}", local_var_configuration.base_path, id=id, userId=user_id);
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
        let local_var_entity: Option<AcceptGroupJoinRequestError> = serde_json::from_str(&local_var_content).ok();
        let local_var_error = ResponseContent { status: local_var_status, content: local_var_content, entity: local_var_entity };
        Err(Error::ResponseError(local_var_error))
    }
}

pub async fn accept_group_transfer(configuration: &configuration::Configuration, id: i64) -> Result<(), Error<AcceptGroupTransferError>> {
    let local_var_configuration = configuration;

    let local_var_client = &local_var_configuration.client;

    let local_var_uri_str = format!("{}/v1/groups/{id}/transfer", local_var_configuration.base_path, id=id);
    let mut local_var_req_builder = local_var_client.request(reqwest::Method::PUT, local_var_uri_str.as_str());

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
        let local_var_entity: Option<AcceptGroupTransferError> = serde_json::from_str(&local_var_content).ok();
        let local_var_error = ResponseContent { status: local_var_status, content: local_var_content, entity: local_var_entity };
        Err(Error::ResponseError(local_var_error))
    }
}

pub async fn ban_group_user(configuration: &configuration::Configuration, id: i64, user_id: i64) -> Result<(), Error<BanGroupUserError>> {
    let local_var_configuration = configuration;

    let local_var_client = &local_var_configuration.client;

    let local_var_uri_str = format!("{}/v1/groups/{id}/ban/{userId}", local_var_configuration.base_path, id=id, userId=user_id);
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
        let local_var_entity: Option<BanGroupUserError> = serde_json::from_str(&local_var_content).ok();
        let local_var_error = ResponseContent { status: local_var_status, content: local_var_content, entity: local_var_entity };
        Err(Error::ResponseError(local_var_error))
    }
}

pub async fn cancel_group_transfer(configuration: &configuration::Configuration, id: i64) -> Result<(), Error<CancelGroupTransferError>> {
    let local_var_configuration = configuration;

    let local_var_client = &local_var_configuration.client;

    let local_var_uri_str = format!("{}/v1/groups/{id}/transfer", local_var_configuration.base_path, id=id);
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
        let local_var_entity: Option<CancelGroupTransferError> = serde_json::from_str(&local_var_content).ok();
        let local_var_error = ResponseContent { status: local_var_status, content: local_var_content, entity: local_var_entity };
        Err(Error::ResponseError(local_var_error))
    }
}

pub async fn create_group(configuration: &configuration::Configuration, allow_members_to_post_image_and_video: Option<bool>, allow_members_to_post_url: Option<bool>, allow_ownership_transfer: Option<bool>, allow_thread_creation_by: Option<&str>, api_key: Option<&str>, call_timeline_display: Option<bool>, cover_image_filename: Option<&str>, description: Option<&str>, gender: Option<i32>, generation_groups_limit: Option<i32>, group_category_id: Option<i64>, group_icon_filename: Option<&str>, guidelines: Option<&str>, hide_conference_call: Option<bool>, hide_from_game_eight: Option<bool>, hide_reported_posts: Option<bool>, is_private: Option<bool>, only_mobile_verified: Option<bool>, only_verified_age: Option<bool>, secret: Option<bool>, signed_info: Option<&str>, sub_category_id: Option<&str>, timestamp: Option<i64>, topic: Option<&str>, uuid: Option<&str>) -> Result<models::CreateGroupResponse, Error<CreateGroupError>> {
    let local_var_configuration = configuration;

    let local_var_client = &local_var_configuration.client;

    let local_var_uri_str = format!("{}/v3/groups/new", local_var_configuration.base_path);
    let mut local_var_req_builder = local_var_client.request(reqwest::Method::POST, local_var_uri_str.as_str());

    if let Some(ref local_var_user_agent) = local_var_configuration.user_agent {
        local_var_req_builder = local_var_req_builder.header(reqwest::header::USER_AGENT, local_var_user_agent.clone());
    }
    let mut local_var_form = reqwest::multipart::Form::new();
    if let Some(local_var_param_value) = allow_members_to_post_image_and_video {
        local_var_form = local_var_form.text("allow_members_to_post_image_and_video", local_var_param_value.to_string());
    }
    if let Some(local_var_param_value) = allow_members_to_post_url {
        local_var_form = local_var_form.text("allow_members_to_post_url", local_var_param_value.to_string());
    }
    if let Some(local_var_param_value) = allow_ownership_transfer {
        local_var_form = local_var_form.text("allow_ownership_transfer", local_var_param_value.to_string());
    }
    if let Some(local_var_param_value) = allow_thread_creation_by {
        local_var_form = local_var_form.text("allow_thread_creation_by", local_var_param_value.to_string());
    }
    if let Some(local_var_param_value) = api_key {
        local_var_form = local_var_form.text("api_key", local_var_param_value.to_string());
    }
    if let Some(local_var_param_value) = call_timeline_display {
        local_var_form = local_var_form.text("call_timeline_display", local_var_param_value.to_string());
    }
    if let Some(local_var_param_value) = cover_image_filename {
        local_var_form = local_var_form.text("cover_image_filename", local_var_param_value.to_string());
    }
    if let Some(local_var_param_value) = description {
        local_var_form = local_var_form.text("description", local_var_param_value.to_string());
    }
    if let Some(local_var_param_value) = gender {
        local_var_form = local_var_form.text("gender", local_var_param_value.to_string());
    }
    if let Some(local_var_param_value) = generation_groups_limit {
        local_var_form = local_var_form.text("generation_groups_limit", local_var_param_value.to_string());
    }
    if let Some(local_var_param_value) = group_category_id {
        local_var_form = local_var_form.text("group_category_id", local_var_param_value.to_string());
    }
    if let Some(local_var_param_value) = group_icon_filename {
        local_var_form = local_var_form.text("group_icon_filename", local_var_param_value.to_string());
    }
    if let Some(local_var_param_value) = guidelines {
        local_var_form = local_var_form.text("guidelines", local_var_param_value.to_string());
    }
    if let Some(local_var_param_value) = hide_conference_call {
        local_var_form = local_var_form.text("hide_conference_call", local_var_param_value.to_string());
    }
    if let Some(local_var_param_value) = hide_from_game_eight {
        local_var_form = local_var_form.text("hide_from_game_eight", local_var_param_value.to_string());
    }
    if let Some(local_var_param_value) = hide_reported_posts {
        local_var_form = local_var_form.text("hide_reported_posts", local_var_param_value.to_string());
    }
    if let Some(local_var_param_value) = is_private {
        local_var_form = local_var_form.text("is_private", local_var_param_value.to_string());
    }
    if let Some(local_var_param_value) = only_mobile_verified {
        local_var_form = local_var_form.text("only_mobile_verified", local_var_param_value.to_string());
    }
    if let Some(local_var_param_value) = only_verified_age {
        local_var_form = local_var_form.text("only_verified_age", local_var_param_value.to_string());
    }
    if let Some(local_var_param_value) = secret {
        local_var_form = local_var_form.text("secret", local_var_param_value.to_string());
    }
    if let Some(local_var_param_value) = signed_info {
        local_var_form = local_var_form.text("signed_info", local_var_param_value.to_string());
    }
    if let Some(local_var_param_value) = sub_category_id {
        local_var_form = local_var_form.text("sub_category_id", local_var_param_value.to_string());
    }
    if let Some(local_var_param_value) = timestamp {
        local_var_form = local_var_form.text("timestamp", local_var_param_value.to_string());
    }
    if let Some(local_var_param_value) = topic {
        local_var_form = local_var_form.text("topic", local_var_param_value.to_string());
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
        let local_var_entity: Option<CreateGroupError> = serde_json::from_str(&local_var_content).ok();
        let local_var_error = ResponseContent { status: local_var_status, content: local_var_content, entity: local_var_entity };
        Err(Error::ResponseError(local_var_error))
    }
}

pub async fn decline_group_join_request(configuration: &configuration::Configuration, id: i64, user_id: i64) -> Result<(), Error<DeclineGroupJoinRequestError>> {
    let local_var_configuration = configuration;

    let local_var_client = &local_var_configuration.client;

    let local_var_uri_str = format!("{}/v1/groups/{id}/decline/{userId}", local_var_configuration.base_path, id=id, userId=user_id);
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
        let local_var_entity: Option<DeclineGroupJoinRequestError> = serde_json::from_str(&local_var_content).ok();
        let local_var_error = ResponseContent { status: local_var_status, content: local_var_content, entity: local_var_entity };
        Err(Error::ResponseError(local_var_error))
    }
}

pub async fn deputize_group_users(configuration: &configuration::Configuration, id: i64) -> Result<(), Error<DeputizeGroupUsersError>> {
    let local_var_configuration = configuration;

    let local_var_client = &local_var_configuration.client;

    let local_var_uri_str = format!("{}/v1/groups/{id}/deputize", local_var_configuration.base_path, id=id);
    let mut local_var_req_builder = local_var_client.request(reqwest::Method::PUT, local_var_uri_str.as_str());

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
        let local_var_entity: Option<DeputizeGroupUsersError> = serde_json::from_str(&local_var_content).ok();
        let local_var_error = ResponseContent { status: local_var_status, content: local_var_content, entity: local_var_entity };
        Err(Error::ResponseError(local_var_error))
    }
}

pub async fn deputize_group_users_mass(configuration: &configuration::Configuration, group_id: i64, api_key: Option<&str>, signed_info: Option<&str>, timestamp: Option<i64>, user_ids: Option<Vec<i64>>, uuid: Option<&str>) -> Result<(), Error<DeputizeGroupUsersMassError>> {
    let local_var_configuration = configuration;

    let local_var_client = &local_var_configuration.client;

    let local_var_uri_str = format!("{}/v3/groups/{group_id}/deputize/mass", local_var_configuration.base_path, group_id=group_id);
    let mut local_var_req_builder = local_var_client.request(reqwest::Method::POST, local_var_uri_str.as_str());

    if let Some(ref local_var_user_agent) = local_var_configuration.user_agent {
        local_var_req_builder = local_var_req_builder.header(reqwest::header::USER_AGENT, local_var_user_agent.clone());
    }
    let mut local_var_form = reqwest::multipart::Form::new();
    if let Some(local_var_param_value) = api_key {
        local_var_form = local_var_form.text("api_key", local_var_param_value.to_string());
    }
    if let Some(local_var_param_value) = signed_info {
        local_var_form = local_var_form.text("signed_info", local_var_param_value.to_string());
    }
    if let Some(local_var_param_value) = timestamp {
        local_var_form = local_var_form.text("timestamp", local_var_param_value.to_string());
    }
    if let Some(local_var_param_value) = user_ids {
        local_var_form = local_var_form.text("user_ids[]", local_var_param_value.into_iter().map(|p| p.to_string()).collect::<Vec<String>>().join(",").to_string());
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
        Ok(())
    } else {
        let local_var_entity: Option<DeputizeGroupUsersMassError> = serde_json::from_str(&local_var_content).ok();
        let local_var_error = ResponseContent { status: local_var_status, content: local_var_content, entity: local_var_entity };
        Err(Error::ResponseError(local_var_error))
    }
}

pub async fn fire_group_user(configuration: &configuration::Configuration, group_id: i64, user_id: i64) -> Result<(), Error<FireGroupUserError>> {
    let local_var_configuration = configuration;

    let local_var_client = &local_var_configuration.client;

    let local_var_uri_str = format!("{}/v1/groups/{group_id}/fire/{user_id}", local_var_configuration.base_path, group_id=group_id, user_id=user_id);
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
        let local_var_entity: Option<FireGroupUserError> = serde_json::from_str(&local_var_content).ok();
        let local_var_error = ResponseContent { status: local_var_status, content: local_var_content, entity: local_var_entity };
        Err(Error::ResponseError(local_var_error))
    }
}

pub async fn get_community_campaign_group_leaderboard(configuration: &configuration::Configuration, ) -> Result<models::GroupOverallLeaderboardResponse, Error<GetCommunityCampaignGroupLeaderboardError>> {
    let local_var_configuration = configuration;

    let local_var_client = &local_var_configuration.client;

    let local_var_uri_str = format!("{}/v1/community_campaign/group_leaderboard", local_var_configuration.base_path);
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
        let local_var_entity: Option<GetCommunityCampaignGroupLeaderboardError> = serde_json::from_str(&local_var_content).ok();
        let local_var_error = ResponseContent { status: local_var_status, content: local_var_content, entity: local_var_entity };
        Err(Error::ResponseError(local_var_error))
    }
}

pub async fn get_community_campaign_status(configuration: &configuration::Configuration, ) -> Result<models::GroupCommunityCampaignResponse, Error<GetCommunityCampaignStatusError>> {
    let local_var_configuration = configuration;

    let local_var_client = &local_var_configuration.client;

    let local_var_uri_str = format!("{}/v1/community_campaign", local_var_configuration.base_path);
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
        let local_var_entity: Option<GetCommunityCampaignStatusError> = serde_json::from_str(&local_var_content).ok();
        let local_var_error = ResponseContent { status: local_var_status, content: local_var_content, entity: local_var_entity };
        Err(Error::ResponseError(local_var_error))
    }
}

pub async fn get_community_campaign_user_leaderboard(configuration: &configuration::Configuration, id: i64) -> Result<models::GroupInCircleUserLeaderboardResponse, Error<GetCommunityCampaignUserLeaderboardError>> {
    let local_var_configuration = configuration;

    let local_var_client = &local_var_configuration.client;

    let local_var_uri_str = format!("{}/v1/community_campaign/groups/{id}/user_leaderboard", local_var_configuration.base_path, id=id);
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
        let local_var_entity: Option<GetCommunityCampaignUserLeaderboardError> = serde_json::from_str(&local_var_content).ok();
        let local_var_error = ResponseContent { status: local_var_status, content: local_var_content, entity: local_var_entity };
        Err(Error::ResponseError(local_var_error))
    }
}

pub async fn get_group(configuration: &configuration::Configuration, id: i64) -> Result<models::GroupResponse, Error<GetGroupError>> {
    let local_var_configuration = configuration;

    let local_var_client = &local_var_configuration.client;

    let local_var_uri_str = format!("{}/v1/groups/{id}", local_var_configuration.base_path, id=id);
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
        let local_var_entity: Option<GetGroupError> = serde_json::from_str(&local_var_content).ok();
        let local_var_error = ResponseContent { status: local_var_status, content: local_var_content, entity: local_var_entity };
        Err(Error::ResponseError(local_var_error))
    }
}

pub async fn get_group_ban_list(configuration: &configuration::Configuration, id: i64, page: Option<i32>) -> Result<models::UsersResponse, Error<GetGroupBanListError>> {
    let local_var_configuration = configuration;

    let local_var_client = &local_var_configuration.client;

    let local_var_uri_str = format!("{}/v1/groups/{id}/ban_list", local_var_configuration.base_path, id=id);
    let mut local_var_req_builder = local_var_client.request(reqwest::Method::GET, local_var_uri_str.as_str());

    if let Some(ref local_var_str) = page {
        local_var_req_builder = local_var_req_builder.query(&[("page", &local_var_str.to_string())]);
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
        let local_var_entity: Option<GetGroupBanListError> = serde_json::from_str(&local_var_content).ok();
        let local_var_error = ResponseContent { status: local_var_status, content: local_var_content, entity: local_var_entity };
        Err(Error::ResponseError(local_var_error))
    }
}

pub async fn get_group_create_quota(configuration: &configuration::Configuration, ) -> Result<models::CreateQuotaResponse, Error<GetGroupCreateQuotaError>> {
    let local_var_configuration = configuration;

    let local_var_client = &local_var_configuration.client;

    let local_var_uri_str = format!("{}/v1/groups/created_quota", local_var_configuration.base_path);
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
        let local_var_entity: Option<GetGroupCreateQuotaError> = serde_json::from_str(&local_var_content).ok();
        let local_var_error = ResponseContent { status: local_var_status, content: local_var_content, entity: local_var_entity };
        Err(Error::ResponseError(local_var_error))
    }
}

pub async fn get_group_gift_history(configuration: &configuration::Configuration, group_id: i64, number: Option<i32>, from: Option<&str>) -> Result<models::GroupGiftHistoryResponse, Error<GetGroupGiftHistoryError>> {
    let local_var_configuration = configuration;

    let local_var_client = &local_var_configuration.client;

    let local_var_uri_str = format!("{}/v1/groups/{group_id}/gift_history", local_var_configuration.base_path, group_id=group_id);
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
        let local_var_entity: Option<GetGroupGiftHistoryError> = serde_json::from_str(&local_var_content).ok();
        let local_var_error = ResponseContent { status: local_var_status, content: local_var_content, entity: local_var_entity };
        Err(Error::ResponseError(local_var_error))
    }
}

pub async fn get_group_gift_transactions(configuration: &configuration::Configuration, group_id: i64, number: Option<i32>, from: Option<&str>) -> Result<models::GiftTransactionsResponse, Error<GetGroupGiftTransactionsError>> {
    let local_var_configuration = configuration;

    let local_var_client = &local_var_configuration.client;

    let local_var_uri_str = format!("{}/v1/groups/{group_id}/gift_transactions", local_var_configuration.base_path, group_id=group_id);
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
        let local_var_entity: Option<GetGroupGiftTransactionsError> = serde_json::from_str(&local_var_content).ok();
        let local_var_error = ResponseContent { status: local_var_status, content: local_var_content, entity: local_var_entity };
        Err(Error::ResponseError(local_var_error))
    }
}

pub async fn get_group_highlights(configuration: &configuration::Configuration, group_id: i64, from: Option<i64>, number: Option<i32>) -> Result<models::PostsResponse, Error<GetGroupHighlightsError>> {
    let local_var_configuration = configuration;

    let local_var_client = &local_var_configuration.client;

    let local_var_uri_str = format!("{}/v1/groups/{group_id}/highlights", local_var_configuration.base_path, group_id=group_id);
    let mut local_var_req_builder = local_var_client.request(reqwest::Method::GET, local_var_uri_str.as_str());

    if let Some(ref local_var_str) = from {
        local_var_req_builder = local_var_req_builder.query(&[("from", &local_var_str.to_string())]);
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
        let local_var_entity: Option<GetGroupHighlightsError> = serde_json::from_str(&local_var_content).ok();
        let local_var_error = ResponseContent { status: local_var_status, content: local_var_content, entity: local_var_entity };
        Err(Error::ResponseError(local_var_error))
    }
}

pub async fn get_group_member(configuration: &configuration::Configuration, id: i64, user_id: i64) -> Result<models::GroupUserResponse, Error<GetGroupMemberError>> {
    let local_var_configuration = configuration;

    let local_var_client = &local_var_configuration.client;

    let local_var_uri_str = format!("{}/v1/groups/{id}/members/{userId}", local_var_configuration.base_path, id=id, userId=user_id);
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
        let local_var_entity: Option<GetGroupMemberError> = serde_json::from_str(&local_var_content).ok();
        let local_var_error = ResponseContent { status: local_var_status, content: local_var_content, entity: local_var_entity };
        Err(Error::ResponseError(local_var_error))
    }
}

pub async fn get_group_members(configuration: &configuration::Configuration, id: i64, number: Option<i32>, from_id: Option<i64>) -> Result<models::GroupUsersResponse, Error<GetGroupMembersError>> {
    let local_var_configuration = configuration;

    let local_var_client = &local_var_configuration.client;

    let local_var_uri_str = format!("{}/v2/groups/{id}/members", local_var_configuration.base_path, id=id);
    let mut local_var_req_builder = local_var_client.request(reqwest::Method::GET, local_var_uri_str.as_str());

    if let Some(ref local_var_str) = number {
        local_var_req_builder = local_var_req_builder.query(&[("number", &local_var_str.to_string())]);
    }
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
        let local_var_entity: Option<GetGroupMembersError> = serde_json::from_str(&local_var_content).ok();
        let local_var_error = ResponseContent { status: local_var_status, content: local_var_content, entity: local_var_entity };
        Err(Error::ResponseError(local_var_error))
    }
}

pub async fn get_group_received_gift_senders(configuration: &configuration::Configuration, group_id: i64, gift_id: i64, user_id: Option<i64>, timestamp: Option<i64>, number: Option<i32>) -> Result<models::GiftSendersResponse, Error<GetGroupReceivedGiftSendersError>> {
    let local_var_configuration = configuration;

    let local_var_client = &local_var_configuration.client;

    let local_var_uri_str = format!("{}/v1/groups/{group_id}/received_gifts/{gift_id}/senders", local_var_configuration.base_path, group_id=group_id, gift_id=gift_id);
    let mut local_var_req_builder = local_var_client.request(reqwest::Method::GET, local_var_uri_str.as_str());

    if let Some(ref local_var_str) = user_id {
        local_var_req_builder = local_var_req_builder.query(&[("user_id", &local_var_str.to_string())]);
    }
    if let Some(ref local_var_str) = timestamp {
        local_var_req_builder = local_var_req_builder.query(&[("timestamp", &local_var_str.to_string())]);
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
        let local_var_entity: Option<GetGroupReceivedGiftSendersError> = serde_json::from_str(&local_var_content).ok();
        let local_var_error = ResponseContent { status: local_var_status, content: local_var_content, entity: local_var_entity };
        Err(Error::ResponseError(local_var_error))
    }
}

pub async fn get_group_unread_status(configuration: &configuration::Configuration, from_time: Option<i64>) -> Result<models::UnreadStatusResponse, Error<GetGroupUnreadStatusError>> {
    let local_var_configuration = configuration;

    let local_var_client = &local_var_configuration.client;

    let local_var_uri_str = format!("{}/v1/groups/unread_status", local_var_configuration.base_path);
    let mut local_var_req_builder = local_var_client.request(reqwest::Method::GET, local_var_uri_str.as_str());

    if let Some(ref local_var_str) = from_time {
        local_var_req_builder = local_var_req_builder.query(&[("from_time", &local_var_str.to_string())]);
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
        let local_var_entity: Option<GetGroupUnreadStatusError> = serde_json::from_str(&local_var_content).ok();
        let local_var_error = ResponseContent { status: local_var_status, content: local_var_content, entity: local_var_entity };
        Err(Error::ResponseError(local_var_error))
    }
}

pub async fn get_invitable_group_users(configuration: &configuration::Configuration, group_id: i64, from_timestamp: Option<i64>, user_nickname: Option<&str>) -> Result<models::UsersByTimestampResponse, Error<GetInvitableGroupUsersError>> {
    let local_var_configuration = configuration;

    let local_var_client = &local_var_configuration.client;

    let local_var_uri_str = format!("{}/v1/groups/{group_id}/users/invitable", local_var_configuration.base_path, group_id=group_id);
    let mut local_var_req_builder = local_var_client.request(reqwest::Method::GET, local_var_uri_str.as_str());

    if let Some(ref local_var_str) = from_timestamp {
        local_var_req_builder = local_var_req_builder.query(&[("from_timestamp", &local_var_str.to_string())]);
    }
    if let Some(ref local_var_str) = user_nickname {
        local_var_req_builder = local_var_req_builder.query(&[("user[nickname]", &local_var_str.to_string())]);
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
        let local_var_entity: Option<GetInvitableGroupUsersError> = serde_json::from_str(&local_var_content).ok();
        let local_var_error = ResponseContent { status: local_var_status, content: local_var_content, entity: local_var_entity };
        Err(Error::ResponseError(local_var_error))
    }
}

pub async fn get_joined_group_statuses(configuration: &configuration::Configuration, ids: Option<Vec<i64>>) -> Result<std::collections::HashMap<String, String>, Error<GetJoinedGroupStatusesError>> {
    let local_var_configuration = configuration;

    let local_var_client = &local_var_configuration.client;

    let local_var_uri_str = format!("{}/v1/groups/joined_statuses", local_var_configuration.base_path);
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
        let local_var_entity: Option<GetJoinedGroupStatusesError> = serde_json::from_str(&local_var_content).ok();
        let local_var_error = ResponseContent { status: local_var_status, content: local_var_content, entity: local_var_entity };
        Err(Error::ResponseError(local_var_error))
    }
}

pub async fn get_relatable_groups(configuration: &configuration::Configuration, id: i64, keyword: Option<&str>, from: Option<&str>) -> Result<models::GroupsRelatedResponse, Error<GetRelatableGroupsError>> {
    let local_var_configuration = configuration;

    let local_var_client = &local_var_configuration.client;

    let local_var_uri_str = format!("{}/v1/groups/{id}/relatable", local_var_configuration.base_path, id=id);
    let mut local_var_req_builder = local_var_client.request(reqwest::Method::GET, local_var_uri_str.as_str());

    if let Some(ref local_var_str) = keyword {
        local_var_req_builder = local_var_req_builder.query(&[("keyword", &local_var_str.to_string())]);
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
        let local_var_entity: Option<GetRelatableGroupsError> = serde_json::from_str(&local_var_content).ok();
        let local_var_error = ResponseContent { status: local_var_status, content: local_var_content, entity: local_var_entity };
        Err(Error::ResponseError(local_var_error))
    }
}

pub async fn get_related_groups(configuration: &configuration::Configuration, id: i64, keyword: Option<&str>, from: Option<&str>) -> Result<models::GroupsRelatedResponse, Error<GetRelatedGroupsError>> {
    let local_var_configuration = configuration;

    let local_var_client = &local_var_configuration.client;

    let local_var_uri_str = format!("{}/v1/groups/{id}/related", local_var_configuration.base_path, id=id);
    let mut local_var_req_builder = local_var_client.request(reqwest::Method::GET, local_var_uri_str.as_str());

    if let Some(ref local_var_str) = keyword {
        local_var_req_builder = local_var_req_builder.query(&[("keyword", &local_var_str.to_string())]);
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
        let local_var_entity: Option<GetRelatedGroupsError> = serde_json::from_str(&local_var_content).ok();
        let local_var_error = ResponseContent { status: local_var_status, content: local_var_content, entity: local_var_entity };
        Err(Error::ResponseError(local_var_error))
    }
}

pub async fn get_user_group_list(configuration: &configuration::Configuration, page: Option<i32>, user_id: Option<i64>) -> Result<models::GroupsResponse, Error<GetUserGroupListError>> {
    let local_var_configuration = configuration;

    let local_var_client = &local_var_configuration.client;

    let local_var_uri_str = format!("{}/v1/groups/user_group_list", local_var_configuration.base_path);
    let mut local_var_req_builder = local_var_client.request(reqwest::Method::GET, local_var_uri_str.as_str());

    if let Some(ref local_var_str) = page {
        local_var_req_builder = local_var_req_builder.query(&[("page", &local_var_str.to_string())]);
    }
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
        let local_var_entity: Option<GetUserGroupListError> = serde_json::from_str(&local_var_content).ok();
        let local_var_error = ResponseContent { status: local_var_status, content: local_var_content, entity: local_var_entity };
        Err(Error::ResponseError(local_var_error))
    }
}

pub async fn invite_to_group(configuration: &configuration::Configuration, id: i64, user_ids: Option<Vec<i64>>) -> Result<(), Error<InviteToGroupError>> {
    let local_var_configuration = configuration;

    let local_var_client = &local_var_configuration.client;

    let local_var_uri_str = format!("{}/v1/groups/{id}/invite", local_var_configuration.base_path, id=id);
    let mut local_var_req_builder = local_var_client.request(reqwest::Method::POST, local_var_uri_str.as_str());

    if let Some(ref local_var_user_agent) = local_var_configuration.user_agent {
        local_var_req_builder = local_var_req_builder.header(reqwest::header::USER_AGENT, local_var_user_agent.clone());
    }
    let mut local_var_form = reqwest::multipart::Form::new();
    if let Some(local_var_param_value) = user_ids {
        local_var_form = local_var_form.text("user_ids[]", local_var_param_value.into_iter().map(|p| p.to_string()).collect::<Vec<String>>().join(",").to_string());
    }
    local_var_req_builder = local_var_req_builder.multipart(local_var_form);

    let local_var_req = local_var_req_builder.build()?;
    let local_var_resp = local_var_client.execute(local_var_req).await?;

    let local_var_status = local_var_resp.status();
    let local_var_content = local_var_resp.text().await?;

    if !local_var_status.is_client_error() && !local_var_status.is_server_error() {
        Ok(())
    } else {
        let local_var_entity: Option<InviteToGroupError> = serde_json::from_str(&local_var_content).ok();
        let local_var_error = ResponseContent { status: local_var_status, content: local_var_content, entity: local_var_entity };
        Err(Error::ResponseError(local_var_error))
    }
}

pub async fn join_group(configuration: &configuration::Configuration, id: i64) -> Result<(), Error<JoinGroupError>> {
    let local_var_configuration = configuration;

    let local_var_client = &local_var_configuration.client;

    let local_var_uri_str = format!("{}/v1/groups/{id}/join", local_var_configuration.base_path, id=id);
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
        let local_var_entity: Option<JoinGroupError> = serde_json::from_str(&local_var_content).ok();
        let local_var_error = ResponseContent { status: local_var_status, content: local_var_content, entity: local_var_entity };
        Err(Error::ResponseError(local_var_error))
    }
}

pub async fn leave_group(configuration: &configuration::Configuration, id: i64) -> Result<(), Error<LeaveGroupError>> {
    let local_var_configuration = configuration;

    let local_var_client = &local_var_configuration.client;

    let local_var_uri_str = format!("{}/v1/groups/{id}/leave", local_var_configuration.base_path, id=id);
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
        let local_var_entity: Option<LeaveGroupError> = serde_json::from_str(&local_var_content).ok();
        let local_var_error = ResponseContent { status: local_var_status, content: local_var_content, entity: local_var_entity };
        Err(Error::ResponseError(local_var_error))
    }
}

pub async fn list_group_categories(configuration: &configuration::Configuration, page: Option<i32>, number: Option<i32>) -> Result<models::GroupCategoriesResponse, Error<ListGroupCategoriesError>> {
    let local_var_configuration = configuration;

    let local_var_client = &local_var_configuration.client;

    let local_var_uri_str = format!("{}/v1/groups/categories", local_var_configuration.base_path);
    let mut local_var_req_builder = local_var_client.request(reqwest::Method::GET, local_var_uri_str.as_str());

    if let Some(ref local_var_str) = page {
        local_var_req_builder = local_var_req_builder.query(&[("page", &local_var_str.to_string())]);
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
        let local_var_entity: Option<ListGroupCategoriesError> = serde_json::from_str(&local_var_content).ok();
        let local_var_error = ResponseContent { status: local_var_status, content: local_var_content, entity: local_var_entity };
        Err(Error::ResponseError(local_var_error))
    }
}

pub async fn list_groups(configuration: &configuration::Configuration, group_category_id: Option<i64>, keyword: Option<&str>, from_timestamp: Option<i64>, sub_category_id: Option<i64>) -> Result<models::GroupsResponse, Error<ListGroupsError>> {
    let local_var_configuration = configuration;

    let local_var_client = &local_var_configuration.client;

    let local_var_uri_str = format!("{}/v2/groups", local_var_configuration.base_path);
    let mut local_var_req_builder = local_var_client.request(reqwest::Method::GET, local_var_uri_str.as_str());

    if let Some(ref local_var_str) = group_category_id {
        local_var_req_builder = local_var_req_builder.query(&[("group_category_id", &local_var_str.to_string())]);
    }
    if let Some(ref local_var_str) = keyword {
        local_var_req_builder = local_var_req_builder.query(&[("keyword", &local_var_str.to_string())]);
    }
    if let Some(ref local_var_str) = from_timestamp {
        local_var_req_builder = local_var_req_builder.query(&[("from_timestamp", &local_var_str.to_string())]);
    }
    if let Some(ref local_var_str) = sub_category_id {
        local_var_req_builder = local_var_req_builder.query(&[("sub_category_id", &local_var_str.to_string())]);
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
        let local_var_entity: Option<ListGroupsError> = serde_json::from_str(&local_var_content).ok();
        let local_var_error = ResponseContent { status: local_var_status, content: local_var_content, entity: local_var_entity };
        Err(Error::ResponseError(local_var_error))
    }
}

pub async fn list_my_groups(configuration: &configuration::Configuration, from_timestamp: Option<i64>) -> Result<models::GroupsResponse, Error<ListMyGroupsError>> {
    let local_var_configuration = configuration;

    let local_var_client = &local_var_configuration.client;

    let local_var_uri_str = format!("{}/v2/groups/mine", local_var_configuration.base_path);
    let mut local_var_req_builder = local_var_client.request(reqwest::Method::GET, local_var_uri_str.as_str());

    if let Some(ref local_var_str) = from_timestamp {
        local_var_req_builder = local_var_req_builder.query(&[("from_timestamp", &local_var_str.to_string())]);
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
        let local_var_entity: Option<ListMyGroupsError> = serde_json::from_str(&local_var_content).ok();
        let local_var_error = ResponseContent { status: local_var_status, content: local_var_content, entity: local_var_entity };
        Err(Error::ResponseError(local_var_error))
    }
}

pub async fn pin_group_highlight_post(configuration: &configuration::Configuration, group_id: i64, post_id: i64) -> Result<(), Error<PinGroupHighlightPostError>> {
    let local_var_configuration = configuration;

    let local_var_client = &local_var_configuration.client;

    let local_var_uri_str = format!("{}/v1/groups/{group_id}/highlights/{post_id}", local_var_configuration.base_path, group_id=group_id, post_id=post_id);
    let mut local_var_req_builder = local_var_client.request(reqwest::Method::PUT, local_var_uri_str.as_str());

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
        let local_var_entity: Option<PinGroupHighlightPostError> = serde_json::from_str(&local_var_content).ok();
        let local_var_error = ResponseContent { status: local_var_status, content: local_var_content, entity: local_var_entity };
        Err(Error::ResponseError(local_var_error))
    }
}

pub async fn remove_group_cover(configuration: &configuration::Configuration, id: i64) -> Result<(), Error<RemoveGroupCoverError>> {
    let local_var_configuration = configuration;

    let local_var_client = &local_var_configuration.client;

    let local_var_uri_str = format!("{}/v3/groups/{id}/cover", local_var_configuration.base_path, id=id);
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
        let local_var_entity: Option<RemoveGroupCoverError> = serde_json::from_str(&local_var_content).ok();
        let local_var_error = ResponseContent { status: local_var_status, content: local_var_content, entity: local_var_entity };
        Err(Error::ResponseError(local_var_error))
    }
}

pub async fn remove_group_deputies(configuration: &configuration::Configuration, id: i64) -> Result<(), Error<RemoveGroupDeputiesError>> {
    let local_var_configuration = configuration;

    let local_var_client = &local_var_configuration.client;

    let local_var_uri_str = format!("{}/v1/groups/{id}/deputize", local_var_configuration.base_path, id=id);
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
        let local_var_entity: Option<RemoveGroupDeputiesError> = serde_json::from_str(&local_var_content).ok();
        let local_var_error = ResponseContent { status: local_var_status, content: local_var_content, entity: local_var_entity };
        Err(Error::ResponseError(local_var_error))
    }
}

pub async fn remove_group_icon(configuration: &configuration::Configuration, id: i64) -> Result<(), Error<RemoveGroupIconError>> {
    let local_var_configuration = configuration;

    let local_var_client = &local_var_configuration.client;

    let local_var_uri_str = format!("{}/v3/groups/{id}/icon", local_var_configuration.base_path, id=id);
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
        let local_var_entity: Option<RemoveGroupIconError> = serde_json::from_str(&local_var_content).ok();
        let local_var_error = ResponseContent { status: local_var_status, content: local_var_content, entity: local_var_entity };
        Err(Error::ResponseError(local_var_error))
    }
}

pub async fn remove_related_groups(configuration: &configuration::Configuration, id: i64, related_group_id: Option<Vec<i64>>) -> Result<(), Error<RemoveRelatedGroupsError>> {
    let local_var_configuration = configuration;

    let local_var_client = &local_var_configuration.client;

    let local_var_uri_str = format!("{}/v1/groups/{id}/related", local_var_configuration.base_path, id=id);
    let mut local_var_req_builder = local_var_client.request(reqwest::Method::DELETE, local_var_uri_str.as_str());

    if let Some(ref local_var_str) = related_group_id {
        local_var_req_builder = match "multi" {
            "multi" => local_var_req_builder.query(&local_var_str.into_iter().map(|p| ("related_group_id[]".to_owned(), p.to_string())).collect::<Vec<(std::string::String, std::string::String)>>()),
            _ => local_var_req_builder.query(&[("related_group_id[]", &local_var_str.into_iter().map(|p| p.to_string()).collect::<Vec<String>>().join(",").to_string())]),
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
        Ok(())
    } else {
        let local_var_entity: Option<RemoveRelatedGroupsError> = serde_json::from_str(&local_var_content).ok();
        let local_var_error = ResponseContent { status: local_var_status, content: local_var_content, entity: local_var_entity };
        Err(Error::ResponseError(local_var_error))
    }
}

pub async fn report_group(configuration: &configuration::Configuration, group_id: i64, category_id: Option<i64>, opponent_id: Option<i64>, reason: Option<&str>, screenshot_2_filename: Option<&str>, screenshot_3_filename: Option<&str>, screenshot_4_filename: Option<&str>, screenshot_filename: Option<&str>) -> Result<(), Error<ReportGroupError>> {
    let local_var_configuration = configuration;

    let local_var_client = &local_var_configuration.client;

    let local_var_uri_str = format!("{}/v3/groups/{group_id}/report", local_var_configuration.base_path, group_id=group_id);
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
        let local_var_entity: Option<ReportGroupError> = serde_json::from_str(&local_var_content).ok();
        let local_var_error = ResponseContent { status: local_var_status, content: local_var_content, entity: local_var_entity };
        Err(Error::ResponseError(local_var_error))
    }
}

pub async fn request_group_walkthrough(configuration: &configuration::Configuration, id: i64) -> Result<(), Error<RequestGroupWalkthroughError>> {
    let local_var_configuration = configuration;

    let local_var_client = &local_var_configuration.client;

    let local_var_uri_str = format!("{}/v1/groups/{id}/request_walkthrough", local_var_configuration.base_path, id=id);
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
        let local_var_entity: Option<RequestGroupWalkthroughError> = serde_json::from_str(&local_var_content).ok();
        let local_var_error = ResponseContent { status: local_var_status, content: local_var_content, entity: local_var_entity };
        Err(Error::ResponseError(local_var_error))
    }
}

pub async fn search_group_posts(configuration: &configuration::Configuration, id: i64, keyword: Option<&str>, from_post_id: Option<i64>, number: Option<i32>, only_thread_posts: Option<bool>) -> Result<models::PostsResponse, Error<SearchGroupPostsError>> {
    let local_var_configuration = configuration;

    let local_var_client = &local_var_configuration.client;

    let local_var_uri_str = format!("{}/v2/groups/{id}/posts/search", local_var_configuration.base_path, id=id);
    let mut local_var_req_builder = local_var_client.request(reqwest::Method::GET, local_var_uri_str.as_str());

    if let Some(ref local_var_str) = keyword {
        local_var_req_builder = local_var_req_builder.query(&[("keyword", &local_var_str.to_string())]);
    }
    if let Some(ref local_var_str) = from_post_id {
        local_var_req_builder = local_var_req_builder.query(&[("from_post_id", &local_var_str.to_string())]);
    }
    if let Some(ref local_var_str) = number {
        local_var_req_builder = local_var_req_builder.query(&[("number", &local_var_str.to_string())]);
    }
    if let Some(ref local_var_str) = only_thread_posts {
        local_var_req_builder = local_var_req_builder.query(&[("only_thread_posts", &local_var_str.to_string())]);
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
        let local_var_entity: Option<SearchGroupPostsError> = serde_json::from_str(&local_var_content).ok();
        let local_var_error = ResponseContent { status: local_var_status, content: local_var_content, entity: local_var_entity };
        Err(Error::ResponseError(local_var_error))
    }
}

pub async fn set_group_title(configuration: &configuration::Configuration, id: i64, title: Option<&str>) -> Result<(), Error<SetGroupTitleError>> {
    let local_var_configuration = configuration;

    let local_var_client = &local_var_configuration.client;

    let local_var_uri_str = format!("{}/v1/groups/{id}/set_title", local_var_configuration.base_path, id=id);
    let mut local_var_req_builder = local_var_client.request(reqwest::Method::POST, local_var_uri_str.as_str());

    if let Some(ref local_var_user_agent) = local_var_configuration.user_agent {
        local_var_req_builder = local_var_req_builder.header(reqwest::header::USER_AGENT, local_var_user_agent.clone());
    }
    let mut local_var_form = reqwest::multipart::Form::new();
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
        let local_var_entity: Option<SetGroupTitleError> = serde_json::from_str(&local_var_content).ok();
        let local_var_error = ResponseContent { status: local_var_status, content: local_var_content, entity: local_var_entity };
        Err(Error::ResponseError(local_var_error))
    }
}

pub async fn take_over_group(configuration: &configuration::Configuration, id: i64) -> Result<(), Error<TakeOverGroupError>> {
    let local_var_configuration = configuration;

    let local_var_client = &local_var_configuration.client;

    let local_var_uri_str = format!("{}/v1/groups/{id}/take_over", local_var_configuration.base_path, id=id);
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
        let local_var_entity: Option<TakeOverGroupError> = serde_json::from_str(&local_var_content).ok();
        let local_var_error = ResponseContent { status: local_var_status, content: local_var_content, entity: local_var_entity };
        Err(Error::ResponseError(local_var_error))
    }
}

pub async fn transfer_group(configuration: &configuration::Configuration, id: i64, api_key: Option<&str>, signed_info: Option<&str>, timestamp: Option<i64>, user_id: Option<i64>, uuid: Option<&str>) -> Result<(), Error<TransferGroupError>> {
    let local_var_configuration = configuration;

    let local_var_client = &local_var_configuration.client;

    let local_var_uri_str = format!("{}/v3/groups/{id}/transfer", local_var_configuration.base_path, id=id);
    let mut local_var_req_builder = local_var_client.request(reqwest::Method::POST, local_var_uri_str.as_str());

    if let Some(ref local_var_user_agent) = local_var_configuration.user_agent {
        local_var_req_builder = local_var_req_builder.header(reqwest::header::USER_AGENT, local_var_user_agent.clone());
    }
    let mut local_var_form = reqwest::multipart::Form::new();
    if let Some(local_var_param_value) = api_key {
        local_var_form = local_var_form.text("api_key", local_var_param_value.to_string());
    }
    if let Some(local_var_param_value) = signed_info {
        local_var_form = local_var_form.text("signed_info", local_var_param_value.to_string());
    }
    if let Some(local_var_param_value) = timestamp {
        local_var_form = local_var_form.text("timestamp", local_var_param_value.to_string());
    }
    if let Some(local_var_param_value) = user_id {
        local_var_form = local_var_form.text("user_id", local_var_param_value.to_string());
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
        Ok(())
    } else {
        let local_var_entity: Option<TransferGroupError> = serde_json::from_str(&local_var_content).ok();
        let local_var_error = ResponseContent { status: local_var_status, content: local_var_content, entity: local_var_entity };
        Err(Error::ResponseError(local_var_error))
    }
}

pub async fn unban_group_user(configuration: &configuration::Configuration, id: i64, user_id: i64) -> Result<(), Error<UnbanGroupUserError>> {
    let local_var_configuration = configuration;

    let local_var_client = &local_var_configuration.client;

    let local_var_uri_str = format!("{}/v1/groups/{id}/unban/{userId}", local_var_configuration.base_path, id=id, userId=user_id);
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
        let local_var_entity: Option<UnbanGroupUserError> = serde_json::from_str(&local_var_content).ok();
        let local_var_error = ResponseContent { status: local_var_status, content: local_var_content, entity: local_var_entity };
        Err(Error::ResponseError(local_var_error))
    }
}

pub async fn unpin_group_highlight_post(configuration: &configuration::Configuration, group_id: i64, post_id: i64) -> Result<(), Error<UnpinGroupHighlightPostError>> {
    let local_var_configuration = configuration;

    let local_var_client = &local_var_configuration.client;

    let local_var_uri_str = format!("{}/v1/groups/{group_id}/highlights/{post_id}", local_var_configuration.base_path, group_id=group_id, post_id=post_id);
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
        let local_var_entity: Option<UnpinGroupHighlightPostError> = serde_json::from_str(&local_var_content).ok();
        let local_var_error = ResponseContent { status: local_var_status, content: local_var_content, entity: local_var_entity };
        Err(Error::ResponseError(local_var_error))
    }
}

pub async fn update_group(configuration: &configuration::Configuration, id: i64, allow_members_to_post_image_and_video: Option<bool>, allow_members_to_post_url: Option<bool>, allow_ownership_transfer: Option<bool>, allow_thread_creation_by: Option<&str>, api_key: Option<&str>, call_timeline_display: Option<bool>, cover_image_filename: Option<&str>, description: Option<&str>, gender: Option<i32>, generation_groups_limit: Option<i32>, group_category_id: Option<i64>, group_icon_filename: Option<&str>, guidelines: Option<&str>, hide_conference_call: Option<bool>, hide_from_game_eight: Option<bool>, hide_reported_posts: Option<bool>, is_private: Option<bool>, only_mobile_verified: Option<bool>, only_verified_age: Option<bool>, secret: Option<bool>, signed_info: Option<&str>, sub_category_id: Option<&str>, timestamp: Option<i64>, topic: Option<&str>, uuid: Option<&str>) -> Result<models::GroupResponse, Error<UpdateGroupError>> {
    let local_var_configuration = configuration;

    let local_var_client = &local_var_configuration.client;

    let local_var_uri_str = format!("{}/v3/groups/{id}/update", local_var_configuration.base_path, id=id);
    let mut local_var_req_builder = local_var_client.request(reqwest::Method::POST, local_var_uri_str.as_str());

    if let Some(ref local_var_user_agent) = local_var_configuration.user_agent {
        local_var_req_builder = local_var_req_builder.header(reqwest::header::USER_AGENT, local_var_user_agent.clone());
    }
    let mut local_var_form = reqwest::multipart::Form::new();
    if let Some(local_var_param_value) = allow_members_to_post_image_and_video {
        local_var_form = local_var_form.text("allow_members_to_post_image_and_video", local_var_param_value.to_string());
    }
    if let Some(local_var_param_value) = allow_members_to_post_url {
        local_var_form = local_var_form.text("allow_members_to_post_url", local_var_param_value.to_string());
    }
    if let Some(local_var_param_value) = allow_ownership_transfer {
        local_var_form = local_var_form.text("allow_ownership_transfer", local_var_param_value.to_string());
    }
    if let Some(local_var_param_value) = allow_thread_creation_by {
        local_var_form = local_var_form.text("allow_thread_creation_by", local_var_param_value.to_string());
    }
    if let Some(local_var_param_value) = api_key {
        local_var_form = local_var_form.text("api_key", local_var_param_value.to_string());
    }
    if let Some(local_var_param_value) = call_timeline_display {
        local_var_form = local_var_form.text("call_timeline_display", local_var_param_value.to_string());
    }
    if let Some(local_var_param_value) = cover_image_filename {
        local_var_form = local_var_form.text("cover_image_filename", local_var_param_value.to_string());
    }
    if let Some(local_var_param_value) = description {
        local_var_form = local_var_form.text("description", local_var_param_value.to_string());
    }
    if let Some(local_var_param_value) = gender {
        local_var_form = local_var_form.text("gender", local_var_param_value.to_string());
    }
    if let Some(local_var_param_value) = generation_groups_limit {
        local_var_form = local_var_form.text("generation_groups_limit", local_var_param_value.to_string());
    }
    if let Some(local_var_param_value) = group_category_id {
        local_var_form = local_var_form.text("group_category_id", local_var_param_value.to_string());
    }
    if let Some(local_var_param_value) = group_icon_filename {
        local_var_form = local_var_form.text("group_icon_filename", local_var_param_value.to_string());
    }
    if let Some(local_var_param_value) = guidelines {
        local_var_form = local_var_form.text("guidelines", local_var_param_value.to_string());
    }
    if let Some(local_var_param_value) = hide_conference_call {
        local_var_form = local_var_form.text("hide_conference_call", local_var_param_value.to_string());
    }
    if let Some(local_var_param_value) = hide_from_game_eight {
        local_var_form = local_var_form.text("hide_from_game_eight", local_var_param_value.to_string());
    }
    if let Some(local_var_param_value) = hide_reported_posts {
        local_var_form = local_var_form.text("hide_reported_posts", local_var_param_value.to_string());
    }
    if let Some(local_var_param_value) = is_private {
        local_var_form = local_var_form.text("is_private", local_var_param_value.to_string());
    }
    if let Some(local_var_param_value) = only_mobile_verified {
        local_var_form = local_var_form.text("only_mobile_verified", local_var_param_value.to_string());
    }
    if let Some(local_var_param_value) = only_verified_age {
        local_var_form = local_var_form.text("only_verified_age", local_var_param_value.to_string());
    }
    if let Some(local_var_param_value) = secret {
        local_var_form = local_var_form.text("secret", local_var_param_value.to_string());
    }
    if let Some(local_var_param_value) = signed_info {
        local_var_form = local_var_form.text("signed_info", local_var_param_value.to_string());
    }
    if let Some(local_var_param_value) = sub_category_id {
        local_var_form = local_var_form.text("sub_category_id", local_var_param_value.to_string());
    }
    if let Some(local_var_param_value) = timestamp {
        local_var_form = local_var_form.text("timestamp", local_var_param_value.to_string());
    }
    if let Some(local_var_param_value) = topic {
        local_var_form = local_var_form.text("topic", local_var_param_value.to_string());
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
        let local_var_entity: Option<UpdateGroupError> = serde_json::from_str(&local_var_content).ok();
        let local_var_error = ResponseContent { status: local_var_status, content: local_var_content, entity: local_var_entity };
        Err(Error::ResponseError(local_var_error))
    }
}

pub async fn update_related_groups(configuration: &configuration::Configuration, id: i64, related_group_id: Option<Vec<i64>>) -> Result<(), Error<UpdateRelatedGroupsError>> {
    let local_var_configuration = configuration;

    let local_var_client = &local_var_configuration.client;

    let local_var_uri_str = format!("{}/v1/groups/{id}/related", local_var_configuration.base_path, id=id);
    let mut local_var_req_builder = local_var_client.request(reqwest::Method::PUT, local_var_uri_str.as_str());

    if let Some(ref local_var_str) = related_group_id {
        local_var_req_builder = match "multi" {
            "multi" => local_var_req_builder.query(&local_var_str.into_iter().map(|p| ("related_group_id[]".to_owned(), p.to_string())).collect::<Vec<(std::string::String, std::string::String)>>()),
            _ => local_var_req_builder.query(&[("related_group_id[]", &local_var_str.into_iter().map(|p| p.to_string()).collect::<Vec<String>>().join(",").to_string())]),
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
        Ok(())
    } else {
        let local_var_entity: Option<UpdateRelatedGroupsError> = serde_json::from_str(&local_var_content).ok();
        let local_var_error = ResponseContent { status: local_var_status, content: local_var_content, entity: local_var_entity };
        Err(Error::ResponseError(local_var_error))
    }
}

pub async fn visit_group(configuration: &configuration::Configuration, id: i64) -> Result<(), Error<VisitGroupError>> {
    let local_var_configuration = configuration;

    let local_var_client = &local_var_configuration.client;

    let local_var_uri_str = format!("{}/v1/groups/{id}/visit", local_var_configuration.base_path, id=id);
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
        let local_var_entity: Option<VisitGroupError> = serde_json::from_str(&local_var_content).ok();
        let local_var_error = ResponseContent { status: local_var_status, content: local_var_content, entity: local_var_entity };
        Err(Error::ResponseError(local_var_error))
    }
}

pub async fn withdraw_group_deputy(configuration: &configuration::Configuration, group_id: i64, user_id: i64) -> Result<(), Error<WithdrawGroupDeputyError>> {
    let local_var_configuration = configuration;

    let local_var_client = &local_var_configuration.client;

    let local_var_uri_str = format!("{}/v1/groups/{group_id}/deputize/{user_id}/withdraw", local_var_configuration.base_path, group_id=group_id, user_id=user_id);
    let mut local_var_req_builder = local_var_client.request(reqwest::Method::PUT, local_var_uri_str.as_str());

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
        let local_var_entity: Option<WithdrawGroupDeputyError> = serde_json::from_str(&local_var_content).ok();
        let local_var_error = ResponseContent { status: local_var_status, content: local_var_content, entity: local_var_entity };
        Err(Error::ResponseError(local_var_error))
    }
}

pub async fn withdraw_group_transfer(configuration: &configuration::Configuration, id: i64, user_id: Option<i64>) -> Result<(), Error<WithdrawGroupTransferError>> {
    let local_var_configuration = configuration;

    let local_var_client = &local_var_configuration.client;

    let local_var_uri_str = format!("{}/v1/groups/{id}/transfer/withdraw", local_var_configuration.base_path, id=id);
    let mut local_var_req_builder = local_var_client.request(reqwest::Method::PUT, local_var_uri_str.as_str());

    if let Some(ref local_var_user_agent) = local_var_configuration.user_agent {
        local_var_req_builder = local_var_req_builder.header(reqwest::header::USER_AGENT, local_var_user_agent.clone());
    }
    let mut local_var_form = reqwest::multipart::Form::new();
    if let Some(local_var_param_value) = user_id {
        local_var_form = local_var_form.text("user_id", local_var_param_value.to_string());
    }
    local_var_req_builder = local_var_req_builder.multipart(local_var_form);

    let local_var_req = local_var_req_builder.build()?;
    let local_var_resp = local_var_client.execute(local_var_req).await?;

    let local_var_status = local_var_resp.status();
    let local_var_content = local_var_resp.text().await?;

    if !local_var_status.is_client_error() && !local_var_status.is_server_error() {
        Ok(())
    } else {
        let local_var_entity: Option<WithdrawGroupTransferError> = serde_json::from_str(&local_var_content).ok();
        let local_var_error = ResponseContent { status: local_var_status, content: local_var_content, entity: local_var_entity };
        Err(Error::ResponseError(local_var_error))
    }
}

