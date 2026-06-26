// Code generated from the Yay! API schema; DO NOT EDIT.


use reqwest;
use serde::{Deserialize, Serialize};
use crate::gen::{apis::ResponseContent, models};
use super::{Error, configuration};


/// struct for typed errors of method [`request_email_verification`]
#[derive(Debug, Clone, Serialize, Deserialize)]
#[serde(untagged)]
pub enum RequestEmailVerificationError {
    UnknownValue(serde_json::Value),
}


pub async fn request_email_verification(configuration: &configuration::Configuration, device_uuid: Option<&str>, email: Option<&str>, intent: Option<&str>, locale: Option<&str>) -> Result<models::CommonUrlResponse, Error<RequestEmailVerificationError>> {
    let local_var_configuration = configuration;

    let local_var_client = &local_var_configuration.client;

    let local_var_uri_str = format!("{}/v1/email_verification_urls", local_var_configuration.base_path);
    let mut local_var_req_builder = local_var_client.request(reqwest::Method::POST, local_var_uri_str.as_str());

    if let Some(ref local_var_user_agent) = local_var_configuration.user_agent {
        local_var_req_builder = local_var_req_builder.header(reqwest::header::USER_AGENT, local_var_user_agent.clone());
    }
    let mut local_var_form = reqwest::multipart::Form::new();
    if let Some(local_var_param_value) = device_uuid {
        local_var_form = local_var_form.text("device_uuid", local_var_param_value.to_string());
    }
    if let Some(local_var_param_value) = email {
        local_var_form = local_var_form.text("email", local_var_param_value.to_string());
    }
    if let Some(local_var_param_value) = intent {
        local_var_form = local_var_form.text("intent", local_var_param_value.to_string());
    }
    if let Some(local_var_param_value) = locale {
        local_var_form = local_var_form.text("locale", local_var_param_value.to_string());
    }
    local_var_req_builder = local_var_req_builder.multipart(local_var_form);

    let local_var_req = local_var_req_builder.build()?;
    let local_var_resp = local_var_client.execute(local_var_req).await?;

    let local_var_status = local_var_resp.status();
    let local_var_content = local_var_resp.text().await?;

    if !local_var_status.is_client_error() && !local_var_status.is_server_error() {
        serde_json::from_str(&local_var_content).map_err(Error::from)
    } else {
        let local_var_entity: Option<RequestEmailVerificationError> = serde_json::from_str(&local_var_content).ok();
        let local_var_error = ResponseContent { status: local_var_status, content: local_var_content, entity: local_var_entity };
        Err(Error::ResponseError(local_var_error))
    }
}

