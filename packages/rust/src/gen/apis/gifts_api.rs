// Code generated from the Yay! API schema; DO NOT EDIT.


use reqwest;
use serde::{Deserialize, Serialize};
use crate::gen::{apis::ResponseContent, models};
use super::{Error, configuration};


/// struct for typed errors of method [`list_gifts`]
#[derive(Debug, Clone, Serialize, Deserialize)]
#[serde(untagged)]
pub enum ListGiftsError {
    UnknownValue(serde_json::Value),
}


pub async fn list_gifts(configuration: &configuration::Configuration, currency: Option<&str>, number: Option<i32>, from: Option<&str>) -> Result<models::GiftsResponse, Error<ListGiftsError>> {
    let local_var_configuration = configuration;

    let local_var_client = &local_var_configuration.client;

    let local_var_uri_str = format!("{}/v2/gifts", local_var_configuration.base_path);
    let mut local_var_req_builder = local_var_client.request(reqwest::Method::GET, local_var_uri_str.as_str());

    if let Some(ref local_var_str) = currency {
        local_var_req_builder = local_var_req_builder.query(&[("currency", &local_var_str.to_string())]);
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
        let local_var_entity: Option<ListGiftsError> = serde_json::from_str(&local_var_content).ok();
        let local_var_error = ResponseContent { status: local_var_status, content: local_var_content, entity: local_var_entity };
        Err(Error::ResponseError(local_var_error))
    }
}

