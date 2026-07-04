// Code generated; DO NOT EDIT.
//
// Flat operation facade: every operation is reachable as
// client.<op>(), delegating to the per-module free function.
// A hand-written method of the same name takes precedence
// (it is excluded from this file).

use crate::Client;

impl Client {
    pub async fn get_user_activities(&self, from_timestamp: Option<i64>) -> Result<crate::gen::models::ActivitiesResponse, crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::activities_api::get_user_activities(&config, from_timestamp).await.map_err(Into::into)
    }

    pub async fn get_user_activities_v1(&self, important: Option<bool>, from_timestamp: Option<i64>, number: Option<i32>) -> Result<crate::gen::models::ActivitiesResponse, crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::activities_api::get_user_activities_v1(&config, important, from_timestamp, number).await.map_err(Into::into)
    }

    pub async fn get_app_config(&self, app: &str) -> Result<crate::gen::models::ApplicationConfigResponse, crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::apps_api::get_app_config(&config, app).await.map_err(Into::into)
    }

    pub async fn oauth_token(&self, email: Option<&str>, grant_type: Option<&str>, password: Option<&str>, refresh_token: Option<&str>) -> Result<crate::gen::models::TokenResponse, crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::auth_api::oauth_token(&config, email, grant_type, password, refresh_token).await.map_err(Into::into)
    }

    pub async fn get_bucket_presigned_urls(&self, file_names: Option<Vec<String>>) -> Result<crate::gen::models::PresignedUrlsResponse, crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::buckets_api::get_bucket_presigned_urls(&config, file_names).await.map_err(Into::into)
    }

    pub async fn bulk_invite_to_call(&self, call_id: i64, group_id: Option<i64>) -> Result<(), crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::calls_api::bulk_invite_to_call(&config, call_id, group_id).await.map_err(Into::into)
    }

    pub async fn bump_call(&self, call_id: i64, participant_limit: Option<i32>) -> Result<(), crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::calls_api::bump_call(&config, call_id, participant_limit).await.map_err(Into::into)
    }

    pub async fn generate_call_action_signature(&self, action: Option<&str>, conference_id: Option<i64>, target_user_uuid: Option<&str>) -> Result<crate::gen::models::CallActionSignatureResponse, crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::calls_api::generate_call_action_signature(&config, action, conference_id, target_user_uuid).await.map_err(Into::into)
    }

    pub async fn get_agora_rtm_token(&self, call_id: i64) -> Result<crate::gen::models::RtmTokenResponse, crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::calls_api::get_agora_rtm_token(&config, call_id).await.map_err(Into::into)
    }

    pub async fn get_call_bgms(&self) -> Result<crate::gen::models::BgmsResponse, crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::calls_api::get_call_bgms(&config).await.map_err(Into::into)
    }

    pub async fn get_call_gift_history(&self, call_id: i64, from: Option<&str>, number: Option<i32>) -> Result<crate::gen::models::CallGiftHistoryResponse, crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::calls_api::get_call_gift_history(&config, call_id, from, number).await.map_err(Into::into)
    }

    pub async fn get_call_participant(&self, call_id: i64, call_user_id: &str) -> Result<crate::gen::models::UserResponse, crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::calls_api::get_call_participant(&config, call_id, call_user_id).await.map_err(Into::into)
    }

    pub async fn get_conference_call(&self, call_id: i64) -> Result<crate::gen::models::ConferenceCallResponse, crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::calls_api::get_conference_call(&config, call_id).await.map_err(Into::into)
    }

    pub async fn get_invitable_call_users(&self, call_id: i64, from_timestamp: Option<i64>, user_nickname: Option<&str>) -> Result<crate::gen::models::UsersByTimestampResponse, crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::calls_api::get_invitable_call_users(&config, call_id, from_timestamp, user_nickname).await.map_err(Into::into)
    }

    pub async fn get_phone_status(&self, opponent_id: i64) -> Result<crate::gen::models::CallStatusResponse, crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::calls_api::get_phone_status(&config, opponent_id).await.map_err(Into::into)
    }

    pub async fn invite_to_call(&self, chat_room_id: Option<i64>, room_id: Option<i64>, room_url: Option<&str>) -> Result<(), crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::calls_api::invite_to_call(&config, chat_room_id, room_id, room_url).await.map_err(Into::into)
    }

    pub async fn invite_to_conference_call(&self, call_id: i64, user_ids: Option<Vec<i64>>) -> Result<(), crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::calls_api::invite_to_conference_call(&config, call_id, user_ids).await.map_err(Into::into)
    }

    pub async fn kick_from_conference_call(&self, call_id: i64, ban: Option<bool>, uuid: Option<&str>) -> Result<crate::gen::models::CallActionSignatureResponse, crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::calls_api::kick_from_conference_call(&config, call_id, ban, uuid).await.map_err(Into::into)
    }

    pub async fn leave_agora_channel(&self, conference_id: Option<i64>, user_id: Option<i64>) -> Result<(), crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::calls_api::leave_agora_channel(&config, conference_id, user_id).await.map_err(Into::into)
    }

    pub async fn leave_conference_call(&self, call_sid: Option<&str>, conference_id: Option<i64>, duration: Option<i64>, total_users_in_call: Option<i32>) -> Result<(), crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::calls_api::leave_conference_call(&config, call_sid, conference_id, duration, total_users_in_call).await.map_err(Into::into)
    }

    pub async fn start_conference_call(&self, call_sid: Option<&str>, conference_id: Option<i64>) -> Result<crate::gen::models::ConferenceCallResponse, crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::calls_api::start_conference_call(&config, call_sid, conference_id).await.map_err(Into::into)
    }

    pub async fn update_call(&self, call_id: i64, category_id: Option<&str>, game_title: Option<&str>, joinable_by: Option<crate::gen::models::JoinableBy>) -> Result<(), crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::calls_api::update_call(&config, call_id, category_id, game_title, joinable_by).await.map_err(Into::into)
    }

    pub async fn update_call_user(&self, call_id: i64, user_id: i64, role: Option<&str>) -> Result<(), crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::calls_api::update_call_user(&config, call_id, user_id, role).await.map_err(Into::into)
    }

    pub async fn validate_call_action_signature(&self, call_id: Option<i64>, sender_uuid: Option<&str>, receiver_uuid: Option<&str>, action: Option<&str>, timestamp: Option<i64>, signature: Option<&str>) -> Result<(), crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::calls_api::validate_call_action_signature(&config, call_id, sender_uuid, receiver_uuid, action, timestamp, signature).await.map_err(Into::into)
    }

    pub async fn accept_chat_request(&self, chat_room_ids: Option<Vec<i64>>) -> Result<(), crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::chat_rooms_api::accept_chat_request(&config, chat_room_ids).await.map_err(Into::into)
    }

    pub async fn create_chat_room(&self, background_filename: Option<&str>, icon_filename: Option<&str>, name: Option<&str>, with_user_ids: Option<Vec<i64>>) -> Result<crate::gen::models::CreateChatRoomResponse, crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::chat_rooms_api::create_chat_room(&config, background_filename, icon_filename, name, with_user_ids).await.map_err(Into::into)
    }

    pub async fn create_chat_room_v1(&self, hima_chat: Option<bool>, matching_id: Option<i64>, with_user_id: Option<i64>) -> Result<crate::gen::models::CreateChatRoomResponse, crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::chat_rooms_api::create_chat_room_v1(&config, hima_chat, matching_id, with_user_id).await.map_err(Into::into)
    }

    pub async fn delete_chat_message(&self, room_id: i64, message_id: i64) -> Result<(), crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::chat_rooms_api::delete_chat_message(&config, room_id, message_id).await.map_err(Into::into)
    }

    pub async fn delete_chat_rooms(&self, chat_room_ids: Option<Vec<i64>>) -> Result<(), crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::chat_rooms_api::delete_chat_rooms(&config, chat_room_ids).await.map_err(Into::into)
    }

    pub async fn get_chat_messages(&self, id: i64, number: Option<i32>, from_message_id: Option<i64>, to_message_id: Option<i64>) -> Result<crate::gen::models::MessagesResponse, crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::chat_rooms_api::get_chat_messages(&config, id, number, from_message_id, to_message_id).await.map_err(Into::into)
    }

    pub async fn get_chat_request_count(&self) -> Result<crate::gen::models::TotalChatRequestResponse, crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::chat_rooms_api::get_chat_request_count(&config).await.map_err(Into::into)
    }

    pub async fn get_chat_requests(&self, from_timestamp: Option<i64>) -> Result<crate::gen::models::ChatRoomsResponse, crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::chat_rooms_api::get_chat_requests(&config, from_timestamp).await.map_err(Into::into)
    }

    pub async fn get_chat_room(&self, id: i64) -> Result<crate::gen::models::ChatRoomResponse, crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::chat_rooms_api::get_chat_room(&config, id).await.map_err(Into::into)
    }

    pub async fn get_chat_unread_status(&self, from_time: Option<i64>) -> Result<crate::gen::models::UnreadStatusResponse, crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::chat_rooms_api::get_chat_unread_status(&config, from_time).await.map_err(Into::into)
    }

    pub async fn get_main_chat_rooms(&self, from_timestamp: Option<i64>) -> Result<crate::gen::models::ChatRoomsResponse, crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::chat_rooms_api::get_main_chat_rooms(&config, from_timestamp).await.map_err(Into::into)
    }

    pub async fn get_updated_chat_rooms(&self, from_time: Option<i64>) -> Result<crate::gen::models::ChatRoomsResponse, crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::chat_rooms_api::get_updated_chat_rooms(&config, from_time).await.map_err(Into::into)
    }

    pub async fn invite_to_chat_room(&self, id: i64, with_user_ids: Option<Vec<i64>>) -> Result<(), crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::chat_rooms_api::invite_to_chat_room(&config, id, with_user_ids).await.map_err(Into::into)
    }

    pub async fn kick_from_chat_room(&self, id: i64, with_user_ids: Option<Vec<i64>>) -> Result<(), crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::chat_rooms_api::kick_from_chat_room(&config, id, with_user_ids).await.map_err(Into::into)
    }

    pub async fn pin_chat_room(&self, id: i64) -> Result<(), crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::chat_rooms_api::pin_chat_room(&config, id).await.map_err(Into::into)
    }

    pub async fn read_chat_attachments(&self, id: i64, read_attachment_request: crate::gen::models::ReadAttachmentRequest) -> Result<(), crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::chat_rooms_api::read_chat_attachments(&config, id, read_attachment_request).await.map_err(Into::into)
    }

    pub async fn read_chat_message(&self, id: i64, message_id: i64) -> Result<(), crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::chat_rooms_api::read_chat_message(&config, id, message_id).await.map_err(Into::into)
    }

    pub async fn read_chat_videos(&self, id: i64, video_msg_ids: Option<Vec<i64>>) -> Result<(), crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::chat_rooms_api::read_chat_videos(&config, id, video_msg_ids).await.map_err(Into::into)
    }

    pub async fn remove_chat_room_background(&self, id: i64) -> Result<(), crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::chat_rooms_api::remove_chat_room_background(&config, id).await.map_err(Into::into)
    }

    pub async fn report_chat_room(&self, chat_room_id: i64, category_id: Option<i64>, opponent_id: Option<i64>, reason: Option<&str>, screenshot_2_filename: Option<&str>, screenshot_3_filename: Option<&str>, screenshot_4_filename: Option<&str>, screenshot_filename: Option<&str>) -> Result<(), crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::chat_rooms_api::report_chat_room(&config, chat_room_id, category_id, opponent_id, reason, screenshot_2_filename, screenshot_3_filename, screenshot_4_filename, screenshot_filename).await.map_err(Into::into)
    }

    pub async fn send_chat_message(&self, id: i64, attachment_file_name: Option<&str>, call_type: Option<crate::gen::models::CallType>, font_size: Option<i32>, gif_image_id: Option<i64>, message_type: Option<crate::gen::models::MessageType>, parent_id: Option<i64>, sticker_id: Option<i64>, sticker_pack_id: Option<i64>, text: Option<&str>, video_file_name: Option<&str>) -> Result<crate::gen::models::MessageResponse, crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::chat_rooms_api::send_chat_message(&config, id, attachment_file_name, call_type, font_size, gif_image_id, message_type, parent_id, sticker_id, sticker_pack_id, text, video_file_name).await.map_err(Into::into)
    }

    pub async fn unpin_chat_room(&self, id: i64) -> Result<(), crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::chat_rooms_api::unpin_chat_room(&config, id).await.map_err(Into::into)
    }

    pub async fn update_chat_room(&self, id: i64, background_filename: Option<&str>, icon_filename: Option<&str>, name: Option<&str>) -> Result<(), crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::chat_rooms_api::update_chat_room(&config, id, background_filename, icon_filename, name).await.map_err(Into::into)
    }

    pub async fn get_conversation(&self, id: i64, group_id: Option<i64>, thread_id: Option<i64>, number: Option<i32>, from_post_id: Option<i64>, reverse: Option<bool>) -> Result<crate::gen::models::PostsResponse, crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::conversations_api::get_conversation(&config, id, group_id, thread_id, number, from_post_id, reverse).await.map_err(Into::into)
    }

    pub async fn get_root_posts(&self, ids: Option<Vec<i64>>) -> Result<crate::gen::models::PostsResponse, crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::conversations_api::get_root_posts(&config, ids).await.map_err(Into::into)
    }

    pub async fn request_email_verification(&self, device_uuid: Option<&str>, email: Option<&str>, intent: Option<&str>, locale: Option<&str>) -> Result<crate::gen::models::CommonUrlResponse, crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::email_verification_urls_api::request_email_verification(&config, device_uuid, email, intent, locale).await.map_err(Into::into)
    }

    pub async fn list_game_apps(&self, number: Option<i32>, from_id: Option<i64>, ids: Option<Vec<i64>>) -> Result<crate::gen::models::GamesResponse, crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::games_api::list_game_apps(&config, number, from_id, ids).await.map_err(Into::into)
    }

    pub async fn list_game_walkthroughs(&self, app_id: i64) -> Result<Vec<crate::gen::models::Walkthrough>, crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::games_api::list_game_walkthroughs(&config, app_id).await.map_err(Into::into)
    }

    pub async fn list_genres(&self, number: Option<i32>, from: Option<&str>) -> Result<crate::gen::models::GenresResponse, crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::genres_api::list_genres(&config, number, from).await.map_err(Into::into)
    }

    pub async fn list_gifts(&self, currency: Option<&str>, number: Option<i32>, from: Option<&str>) -> Result<crate::gen::models::GiftsResponse, crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::gifts_api::list_gifts(&config, currency, number, from).await.map_err(Into::into)
    }

    pub async fn list_muted_group_users(&self, id: i64, keyword: Option<&str>, cursor: Option<&str>, size: Option<i32>) -> Result<crate::gen::models::GroupMuteUsersResponse, crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::group_mute_api::list_muted_group_users(&config, id, keyword, cursor, size).await.map_err(Into::into)
    }

    pub async fn mute_group_user(&self, id: i64, user_id: i64) -> Result<(), crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::group_mute_api::mute_group_user(&config, id, user_id).await.map_err(Into::into)
    }

    pub async fn unmute_group_user(&self, id: i64, user_id: i64) -> Result<(), crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::group_mute_api::unmute_group_user(&config, id, user_id).await.map_err(Into::into)
    }

    pub async fn accept_group_join_request(&self, id: i64, user_id: i64) -> Result<(), crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::groups_api::accept_group_join_request(&config, id, user_id).await.map_err(Into::into)
    }

    pub async fn accept_group_transfer(&self, id: i64) -> Result<(), crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::groups_api::accept_group_transfer(&config, id).await.map_err(Into::into)
    }

    pub async fn ban_group_user(&self, id: i64, user_id: i64) -> Result<(), crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::groups_api::ban_group_user(&config, id, user_id).await.map_err(Into::into)
    }

    pub async fn cancel_group_transfer(&self, id: i64) -> Result<(), crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::groups_api::cancel_group_transfer(&config, id).await.map_err(Into::into)
    }

    pub async fn create_group(&self, allow_members_to_post_image_and_video: Option<bool>, allow_members_to_post_url: Option<bool>, allow_ownership_transfer: Option<bool>, allow_thread_creation_by: Option<&str>, api_key: Option<&str>, call_timeline_display: Option<bool>, cover_image_filename: Option<&str>, description: Option<&str>, gender: Option<i32>, generation_groups_limit: Option<i32>, group_category_id: Option<i64>, group_icon_filename: Option<&str>, guidelines: Option<&str>, hide_conference_call: Option<bool>, hide_from_game_eight: Option<bool>, hide_reported_posts: Option<bool>, is_private: Option<bool>, only_mobile_verified: Option<bool>, only_verified_age: Option<bool>, secret: Option<bool>, signed_info: Option<&str>, sub_category_id: Option<&str>, timestamp: Option<i64>, topic: Option<&str>, uuid: Option<&str>) -> Result<crate::gen::models::CreateGroupResponse, crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::groups_api::create_group(&config, allow_members_to_post_image_and_video, allow_members_to_post_url, allow_ownership_transfer, allow_thread_creation_by, api_key, call_timeline_display, cover_image_filename, description, gender, generation_groups_limit, group_category_id, group_icon_filename, guidelines, hide_conference_call, hide_from_game_eight, hide_reported_posts, is_private, only_mobile_verified, only_verified_age, secret, signed_info, sub_category_id, timestamp, topic, uuid).await.map_err(Into::into)
    }

    pub async fn decline_group_join_request(&self, id: i64, user_id: i64) -> Result<(), crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::groups_api::decline_group_join_request(&config, id, user_id).await.map_err(Into::into)
    }

    pub async fn deputize_group_users(&self, id: i64) -> Result<(), crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::groups_api::deputize_group_users(&config, id).await.map_err(Into::into)
    }

    pub async fn deputize_group_users_mass(&self, group_id: i64, api_key: Option<&str>, signed_info: Option<&str>, timestamp: Option<i64>, user_ids: Option<Vec<i64>>, uuid: Option<&str>) -> Result<(), crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::groups_api::deputize_group_users_mass(&config, group_id, api_key, signed_info, timestamp, user_ids, uuid).await.map_err(Into::into)
    }

    pub async fn fire_group_user(&self, group_id: i64, user_id: i64) -> Result<(), crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::groups_api::fire_group_user(&config, group_id, user_id).await.map_err(Into::into)
    }

    pub async fn get_community_campaign_group_leaderboard(&self) -> Result<crate::gen::models::GroupOverallLeaderboardResponse, crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::groups_api::get_community_campaign_group_leaderboard(&config).await.map_err(Into::into)
    }

    pub async fn get_community_campaign_status(&self) -> Result<crate::gen::models::GroupCommunityCampaignResponse, crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::groups_api::get_community_campaign_status(&config).await.map_err(Into::into)
    }

    pub async fn get_community_campaign_user_leaderboard(&self, id: i64) -> Result<crate::gen::models::GroupInCircleUserLeaderboardResponse, crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::groups_api::get_community_campaign_user_leaderboard(&config, id).await.map_err(Into::into)
    }

    pub async fn get_group(&self, id: i64) -> Result<crate::gen::models::GroupResponse, crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::groups_api::get_group(&config, id).await.map_err(Into::into)
    }

    pub async fn get_group_ban_list(&self, id: i64, page: Option<i32>) -> Result<crate::gen::models::UsersResponse, crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::groups_api::get_group_ban_list(&config, id, page).await.map_err(Into::into)
    }

    pub async fn get_group_create_quota(&self) -> Result<crate::gen::models::CreateQuotaResponse, crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::groups_api::get_group_create_quota(&config).await.map_err(Into::into)
    }

    pub async fn get_group_gift_history(&self, group_id: i64, number: Option<i32>, from: Option<&str>) -> Result<crate::gen::models::GroupGiftHistoryResponse, crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::groups_api::get_group_gift_history(&config, group_id, number, from).await.map_err(Into::into)
    }

    pub async fn get_group_gift_transactions(&self, group_id: i64, number: Option<i32>, from: Option<&str>) -> Result<crate::gen::models::GiftTransactionsResponse, crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::groups_api::get_group_gift_transactions(&config, group_id, number, from).await.map_err(Into::into)
    }

    pub async fn get_group_highlights(&self, group_id: i64, from: Option<i64>, number: Option<i32>) -> Result<crate::gen::models::PostsResponse, crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::groups_api::get_group_highlights(&config, group_id, from, number).await.map_err(Into::into)
    }

    pub async fn get_group_member(&self, id: i64, user_id: i64) -> Result<crate::gen::models::GroupUserResponse, crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::groups_api::get_group_member(&config, id, user_id).await.map_err(Into::into)
    }

    pub async fn get_group_members(&self, id: i64, number: Option<i32>, from_id: Option<i64>) -> Result<crate::gen::models::GroupUsersResponse, crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::groups_api::get_group_members(&config, id, number, from_id).await.map_err(Into::into)
    }

    pub async fn get_group_received_gift_senders(&self, group_id: i64, gift_id: i64, user_id: Option<i64>, timestamp: Option<i64>, number: Option<i32>) -> Result<crate::gen::models::GiftSendersResponse, crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::groups_api::get_group_received_gift_senders(&config, group_id, gift_id, user_id, timestamp, number).await.map_err(Into::into)
    }

    pub async fn get_group_unread_status(&self, from_time: Option<i64>) -> Result<crate::gen::models::UnreadStatusResponse, crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::groups_api::get_group_unread_status(&config, from_time).await.map_err(Into::into)
    }

    pub async fn get_invitable_group_users(&self, group_id: i64, from_timestamp: Option<i64>, user_nickname: Option<&str>) -> Result<crate::gen::models::UsersByTimestampResponse, crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::groups_api::get_invitable_group_users(&config, group_id, from_timestamp, user_nickname).await.map_err(Into::into)
    }

    pub async fn get_joined_group_statuses(&self, ids: Option<Vec<i64>>) -> Result<std::collections::HashMap<String, String>, crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::groups_api::get_joined_group_statuses(&config, ids).await.map_err(Into::into)
    }

    pub async fn get_relatable_groups(&self, id: i64, keyword: Option<&str>, from: Option<&str>) -> Result<crate::gen::models::GroupsRelatedResponse, crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::groups_api::get_relatable_groups(&config, id, keyword, from).await.map_err(Into::into)
    }

    pub async fn get_related_groups(&self, id: i64, keyword: Option<&str>, from: Option<&str>) -> Result<crate::gen::models::GroupsRelatedResponse, crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::groups_api::get_related_groups(&config, id, keyword, from).await.map_err(Into::into)
    }

    pub async fn get_user_group_list(&self, page: Option<i32>, user_id: Option<i64>) -> Result<crate::gen::models::GroupsResponse, crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::groups_api::get_user_group_list(&config, page, user_id).await.map_err(Into::into)
    }

    pub async fn invite_to_group(&self, id: i64, user_ids: Option<Vec<i64>>) -> Result<(), crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::groups_api::invite_to_group(&config, id, user_ids).await.map_err(Into::into)
    }

    pub async fn join_group(&self, id: i64) -> Result<(), crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::groups_api::join_group(&config, id).await.map_err(Into::into)
    }

    pub async fn leave_group(&self, id: i64) -> Result<(), crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::groups_api::leave_group(&config, id).await.map_err(Into::into)
    }

    pub async fn list_group_categories(&self, page: Option<i32>, number: Option<i32>) -> Result<crate::gen::models::GroupCategoriesResponse, crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::groups_api::list_group_categories(&config, page, number).await.map_err(Into::into)
    }

    pub async fn list_groups(&self, group_category_id: Option<i64>, keyword: Option<&str>, from_timestamp: Option<i64>, sub_category_id: Option<i64>) -> Result<crate::gen::models::GroupsResponse, crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::groups_api::list_groups(&config, group_category_id, keyword, from_timestamp, sub_category_id).await.map_err(Into::into)
    }

    pub async fn list_my_groups(&self, from_timestamp: Option<i64>) -> Result<crate::gen::models::GroupsResponse, crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::groups_api::list_my_groups(&config, from_timestamp).await.map_err(Into::into)
    }

    pub async fn pin_group_highlight_post(&self, group_id: i64, post_id: i64) -> Result<(), crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::groups_api::pin_group_highlight_post(&config, group_id, post_id).await.map_err(Into::into)
    }

    pub async fn remove_group_cover(&self, id: i64) -> Result<(), crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::groups_api::remove_group_cover(&config, id).await.map_err(Into::into)
    }

    pub async fn remove_group_deputies(&self, id: i64) -> Result<(), crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::groups_api::remove_group_deputies(&config, id).await.map_err(Into::into)
    }

    pub async fn remove_group_icon(&self, id: i64) -> Result<(), crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::groups_api::remove_group_icon(&config, id).await.map_err(Into::into)
    }

    pub async fn remove_related_groups(&self, id: i64, related_group_id: Option<Vec<i64>>) -> Result<(), crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::groups_api::remove_related_groups(&config, id, related_group_id).await.map_err(Into::into)
    }

    pub async fn report_group(&self, group_id: i64, category_id: Option<i64>, opponent_id: Option<i64>, reason: Option<&str>, screenshot_2_filename: Option<&str>, screenshot_3_filename: Option<&str>, screenshot_4_filename: Option<&str>, screenshot_filename: Option<&str>) -> Result<(), crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::groups_api::report_group(&config, group_id, category_id, opponent_id, reason, screenshot_2_filename, screenshot_3_filename, screenshot_4_filename, screenshot_filename).await.map_err(Into::into)
    }

    pub async fn request_group_walkthrough(&self, id: i64) -> Result<(), crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::groups_api::request_group_walkthrough(&config, id).await.map_err(Into::into)
    }

    pub async fn search_group_posts(&self, id: i64, keyword: Option<&str>, from_post_id: Option<i64>, number: Option<i32>, only_thread_posts: Option<bool>) -> Result<crate::gen::models::PostsResponse, crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::groups_api::search_group_posts(&config, id, keyword, from_post_id, number, only_thread_posts).await.map_err(Into::into)
    }

    pub async fn set_group_title(&self, id: i64, title: Option<&str>) -> Result<(), crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::groups_api::set_group_title(&config, id, title).await.map_err(Into::into)
    }

    pub async fn take_over_group(&self, id: i64) -> Result<(), crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::groups_api::take_over_group(&config, id).await.map_err(Into::into)
    }

    pub async fn transfer_group(&self, id: i64, api_key: Option<&str>, signed_info: Option<&str>, timestamp: Option<i64>, user_id: Option<i64>, uuid: Option<&str>) -> Result<(), crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::groups_api::transfer_group(&config, id, api_key, signed_info, timestamp, user_id, uuid).await.map_err(Into::into)
    }

    pub async fn unban_group_user(&self, id: i64, user_id: i64) -> Result<(), crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::groups_api::unban_group_user(&config, id, user_id).await.map_err(Into::into)
    }

    pub async fn unpin_group_highlight_post(&self, group_id: i64, post_id: i64) -> Result<(), crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::groups_api::unpin_group_highlight_post(&config, group_id, post_id).await.map_err(Into::into)
    }

    pub async fn update_group(&self, id: i64, allow_members_to_post_image_and_video: Option<bool>, allow_members_to_post_url: Option<bool>, allow_ownership_transfer: Option<bool>, allow_thread_creation_by: Option<&str>, api_key: Option<&str>, call_timeline_display: Option<bool>, cover_image_filename: Option<&str>, description: Option<&str>, gender: Option<i32>, generation_groups_limit: Option<i32>, group_category_id: Option<i64>, group_icon_filename: Option<&str>, guidelines: Option<&str>, hide_conference_call: Option<bool>, hide_from_game_eight: Option<bool>, hide_reported_posts: Option<bool>, is_private: Option<bool>, only_mobile_verified: Option<bool>, only_verified_age: Option<bool>, secret: Option<bool>, signed_info: Option<&str>, sub_category_id: Option<&str>, timestamp: Option<i64>, topic: Option<&str>, uuid: Option<&str>) -> Result<crate::gen::models::GroupResponse, crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::groups_api::update_group(&config, id, allow_members_to_post_image_and_video, allow_members_to_post_url, allow_ownership_transfer, allow_thread_creation_by, api_key, call_timeline_display, cover_image_filename, description, gender, generation_groups_limit, group_category_id, group_icon_filename, guidelines, hide_conference_call, hide_from_game_eight, hide_reported_posts, is_private, only_mobile_verified, only_verified_age, secret, signed_info, sub_category_id, timestamp, topic, uuid).await.map_err(Into::into)
    }

    pub async fn update_related_groups(&self, id: i64, related_group_id: Option<Vec<i64>>) -> Result<(), crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::groups_api::update_related_groups(&config, id, related_group_id).await.map_err(Into::into)
    }

    pub async fn visit_group(&self, id: i64) -> Result<(), crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::groups_api::visit_group(&config, id).await.map_err(Into::into)
    }

    pub async fn withdraw_group_deputy(&self, group_id: i64, user_id: i64) -> Result<(), crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::groups_api::withdraw_group_deputy(&config, group_id, user_id).await.map_err(Into::into)
    }

    pub async fn withdraw_group_transfer(&self, id: i64, user_id: Option<i64>) -> Result<(), crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::groups_api::withdraw_group_transfer(&config, id, user_id).await.map_err(Into::into)
    }

    pub async fn create_mute_keyword(&self, mute_keyword_request: crate::gen::models::MuteKeywordRequest) -> Result<crate::gen::models::CreateMuteKeywordResponse, crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::hidden_api::create_mute_keyword(&config, mute_keyword_request).await.map_err(Into::into)
    }

    pub async fn delete_mute_keyword(&self, ids: Option<Vec<i64>>) -> Result<(), crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::hidden_api::delete_mute_keyword(&config, ids).await.map_err(Into::into)
    }

    pub async fn hide_chats(&self, chat_room_id: Option<i64>) -> Result<(), crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::hidden_api::hide_chats(&config, chat_room_id).await.map_err(Into::into)
    }

    pub async fn hide_users(&self, user_id: Option<i64>) -> Result<(), crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::hidden_api::hide_users(&config, user_id).await.map_err(Into::into)
    }

    pub async fn list_hidden_chats(&self, number: Option<i32>, from_timestamp: Option<i64>) -> Result<crate::gen::models::ChatRoomsResponse, crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::hidden_api::list_hidden_chats(&config, number, from_timestamp).await.map_err(Into::into)
    }

    pub async fn list_hidden_users(&self, from: Option<&str>, number: Option<i32>) -> Result<crate::gen::models::HiddenResponse, crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::hidden_api::list_hidden_users(&config, from, number).await.map_err(Into::into)
    }

    pub async fn list_mute_keywords(&self) -> Result<crate::gen::models::MuteKeywordResponse, crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::hidden_api::list_mute_keywords(&config).await.map_err(Into::into)
    }

    pub async fn unhide_chats(&self, chat_room_ids: Option<i64>) -> Result<(), crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::hidden_api::unhide_chats(&config, chat_room_ids).await.map_err(Into::into)
    }

    pub async fn unhide_users(&self, user_ids: Option<Vec<i64>>) -> Result<(), crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::hidden_api::unhide_users(&config, user_ids).await.map_err(Into::into)
    }

    pub async fn update_mute_keyword(&self, id: i64, mute_keyword_request: crate::gen::models::MuteKeywordRequest) -> Result<crate::gen::models::CreateMuteKeywordResponse, crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::hidden_api::update_mute_keyword(&config, id, mute_keyword_request).await.map_err(Into::into)
    }

    pub async fn get_banned_words(&self, country_api_value: &str) -> Result<crate::gen::models::BanWordsResponse, crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::moderation_api::get_banned_words(&config, country_api_value).await.map_err(Into::into)
    }

    pub async fn get_popular_words(&self, country_api_value: &str, app: &str) -> Result<crate::gen::models::PopularWordsResponse, crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::moderation_api::get_popular_words(&config, country_api_value, app).await.map_err(Into::into)
    }

    pub async fn get_group_notification_settings(&self, id: i64) -> Result<crate::gen::models::GroupNotificationSettingsResponse, crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::notification_settings_api::get_group_notification_settings(&config, id).await.map_err(Into::into)
    }

    pub async fn update_chat_room_notification_settings(&self, id: i64, notification_chat: Option<i32>) -> Result<crate::gen::models::NotificationSettingResponse, crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::notification_settings_api::update_chat_room_notification_settings(&config, id, notification_chat).await.map_err(Into::into)
    }

    pub async fn update_group_notification_settings(&self, id: i64, notification_group_join: Option<i32>, notification_group_message_tag_all: Option<i32>, notification_group_post: Option<i32>, notification_group_request: Option<i32>) -> Result<crate::gen::models::AdditionalSettingsResponse, crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::notification_settings_api::update_group_notification_settings(&config, id, notification_group_join, notification_group_message_tag_all, notification_group_post, notification_group_request).await.map_err(Into::into)
    }

    pub async fn pin_group(&self, id: Option<i64>) -> Result<(), crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::pinned_api::pin_group(&config, id).await.map_err(Into::into)
    }

    pub async fn pin_post(&self, id: Option<i64>) -> Result<(), crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::pinned_api::pin_post(&config, id).await.map_err(Into::into)
    }

    pub async fn pin_review(&self, id: Option<i64>) -> Result<(), crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::pinned_api::pin_review(&config, id).await.map_err(Into::into)
    }

    pub async fn unpin_group(&self, id: i64) -> Result<(), crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::pinned_api::unpin_group(&config, id).await.map_err(Into::into)
    }

    pub async fn unpin_post(&self, id: i64) -> Result<(), crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::pinned_api::unpin_post(&config, id).await.map_err(Into::into)
    }

    pub async fn unpin_review(&self, id: i64) -> Result<(), crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::pinned_api::unpin_review(&config, id).await.map_err(Into::into)
    }

    pub async fn create_conference_call_post(&self, api_key: Option<&str>, attachment_2_filename: Option<&str>, attachment_3_filename: Option<&str>, attachment_4_filename: Option<&str>, attachment_5_filename: Option<&str>, attachment_6_filename: Option<&str>, attachment_7_filename: Option<&str>, attachment_8_filename: Option<&str>, attachment_9_filename: Option<&str>, attachment_filename: Option<&str>, call_type: Option<crate::gen::models::CallType>, category_id: Option<i64>, color: Option<i32>, font_size: Option<i32>, game_title: Option<&str>, group_id: Option<i64>, joinable_by: Option<crate::gen::models::JoinableBy>, language: Option<&str>, message_tags: Option<serde_json::Value>, signed_info: Option<&str>, text: Option<&str>, timestamp: Option<i64>, uuid: Option<&str>) -> Result<crate::gen::models::CreatePostResponse, crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::posts_api::create_conference_call_post(&config, api_key, attachment_2_filename, attachment_3_filename, attachment_4_filename, attachment_5_filename, attachment_6_filename, attachment_7_filename, attachment_8_filename, attachment_9_filename, attachment_filename, call_type, category_id, color, font_size, game_title, group_id, joinable_by, language, message_tags, signed_info, text, timestamp, uuid).await.map_err(Into::into)
    }

    pub async fn create_post(&self, x_jwt: Option<&str>, attachment_2_filename: Option<&str>, attachment_3_filename: Option<&str>, attachment_4_filename: Option<&str>, attachment_5_filename: Option<&str>, attachment_6_filename: Option<&str>, attachment_7_filename: Option<&str>, attachment_8_filename: Option<&str>, attachment_9_filename: Option<&str>, attachment_filename: Option<&str>, choices: Option<Vec<String>>, color: Option<i32>, font_size: Option<i32>, group_id: Option<i64>, in_reply_to: Option<i64>, language: Option<&str>, mention_ids: Option<Vec<i64>>, message_tags: Option<serde_json::Value>, post_type: Option<crate::gen::models::PostType>, shared_url: Option<serde_json::Value>, text: Option<&str>, video_file_name: Option<&str>) -> Result<crate::gen::models::Post, crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::posts_api::create_post(&config, x_jwt, attachment_2_filename, attachment_3_filename, attachment_4_filename, attachment_5_filename, attachment_6_filename, attachment_7_filename, attachment_8_filename, attachment_9_filename, attachment_filename, choices, color, font_size, group_id, in_reply_to, language, mention_ids, message_tags, post_type, shared_url, text, video_file_name).await.map_err(Into::into)
    }

    pub async fn create_share_post(&self, api_key: Option<&str>, color: Option<i32>, font_size: Option<i32>, group_id: Option<i64>, language: Option<&str>, message_tags: Option<serde_json::Value>, shareable_id: Option<i64>, shareable_type: Option<&str>, signed_info: Option<&str>, text: Option<&str>, timestamp: Option<i64>, uuid: Option<&str>) -> Result<crate::gen::models::Post, crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::posts_api::create_share_post(&config, api_key, color, font_size, group_id, language, message_tags, shareable_id, shareable_type, signed_info, text, timestamp, uuid).await.map_err(Into::into)
    }

    pub async fn delete_all_posts(&self) -> Result<(), crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::posts_api::delete_all_posts(&config).await.map_err(Into::into)
    }

    pub async fn delete_posts(&self, posts_ids: Option<Vec<i64>>) -> Result<(), crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::posts_api::delete_posts(&config, posts_ids).await.map_err(Into::into)
    }

    pub async fn get_active_call_post(&self, user_id: Option<i64>) -> Result<crate::gen::models::PostResponse, crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::posts_api::get_active_call_post(&config, user_id).await.map_err(Into::into)
    }

    pub async fn get_call_followers_timeline(&self, from_timestamp: Option<i64>, number: Option<i32>, category_id: Option<i64>, call_type: Option<crate::gen::models::CallType>, include_circle_call: Option<bool>, exclude_recent_gomimushi: Option<bool>) -> Result<crate::gen::models::PostsResponse, crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::posts_api::get_call_followers_timeline(&config, from_timestamp, number, category_id, call_type, include_circle_call, exclude_recent_gomimushi).await.map_err(Into::into)
    }

    pub async fn get_call_timeline(&self, group_id: Option<i64>, from_timestamp: Option<i64>, number: Option<i32>, category_id: Option<i64>, call_type: Option<crate::gen::models::CallType>, include_circle_call: Option<bool>, cross_generation: Option<bool>, exclude_recent_gomimushi: Option<bool>) -> Result<crate::gen::models::PostsResponse, crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::posts_api::get_call_timeline(&config, group_id, from_timestamp, number, category_id, call_type, include_circle_call, cross_generation, exclude_recent_gomimushi).await.map_err(Into::into)
    }

    pub async fn get_following_timeline(&self, from: Option<&str>, from_post_id: Option<i64>, only_root: Option<bool>, order_by: Option<&str>, number: Option<i32>, mxn: Option<i32>, reduce_selfie: Option<bool>, custom_generation_range: Option<bool>) -> Result<crate::gen::models::PostsResponse, crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::posts_api::get_following_timeline(&config, from, from_post_id, only_root, order_by, number, mxn, reduce_selfie, custom_generation_range).await.map_err(Into::into)
    }

    pub async fn get_group_calls(&self, number: Option<i32>, group_category_id: Option<i64>, from_timestamp: Option<i64>, scope: Option<&str>) -> Result<crate::gen::models::PostsResponse, crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::posts_api::get_group_calls(&config, number, group_category_id, from_timestamp, scope).await.map_err(Into::into)
    }

    pub async fn get_group_timeline(&self, group_id: Option<i64>, from_post_id: Option<i64>, reverse: Option<bool>, post_type: Option<crate::gen::models::PostType>, number: Option<i32>, only_root: Option<bool>) -> Result<crate::gen::models::PostsResponse, crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::posts_api::get_group_timeline(&config, group_id, from_post_id, reverse, post_type, number, only_root).await.map_err(Into::into)
    }

    pub async fn get_my_posts(&self, from_post_id: Option<i64>, number: Option<i32>, include_group_post: Option<bool>) -> Result<crate::gen::models::PostsResponse, crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::posts_api::get_my_posts(&config, from_post_id, number, include_group_post).await.map_err(Into::into)
    }

    pub async fn get_post(&self, id: i64, cache_control: Option<&str>) -> Result<crate::gen::models::PostResponse, crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::posts_api::get_post(&config, id, cache_control).await.map_err(Into::into)
    }

    pub async fn get_post_gift_transactions(&self, post_id: i64, number: Option<i32>, from: Option<&str>) -> Result<crate::gen::models::GiftTransactionsResponse, crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::posts_api::get_post_gift_transactions(&config, post_id, number, from).await.map_err(Into::into)
    }

    pub async fn get_post_likers(&self, id: i64, from_id: Option<i64>) -> Result<crate::gen::models::PostLikersResponse, crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::posts_api::get_post_likers(&config, id, from_id).await.map_err(Into::into)
    }

    pub async fn get_post_reposts(&self, id: i64, from_post_id: Option<i64>) -> Result<crate::gen::models::PostsResponse, crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::posts_api::get_post_reposts(&config, id, from_post_id).await.map_err(Into::into)
    }

    pub async fn get_post_url_metadata(&self, url: Option<&str>) -> Result<crate::gen::models::SharedUrl, crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::posts_api::get_post_url_metadata(&config, url).await.map_err(Into::into)
    }

    pub async fn get_posts_by_ids(&self, post_ids: Option<Vec<i64>>) -> Result<crate::gen::models::PostsResponse, crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::posts_api::get_posts_by_ids(&config, post_ids).await.map_err(Into::into)
    }

    pub async fn get_posts_by_tag(&self, tag: &str, from_post_id: Option<i64>, number: Option<i32>) -> Result<crate::gen::models::PostsResponse, crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::posts_api::get_posts_by_tag(&config, tag, from_post_id, number).await.map_err(Into::into)
    }

    pub async fn get_recent_engagement_posts(&self, number: Option<i32>) -> Result<crate::gen::models::PostsResponse, crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::posts_api::get_recent_engagement_posts(&config, number).await.map_err(Into::into)
    }

    pub async fn get_recommended_post_tags(&self, save_recent_search: Option<bool>, tag: Option<&str>) -> Result<crate::gen::models::PostTagsResponse, crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::posts_api::get_recommended_post_tags(&config, save_recent_search, tag).await.map_err(Into::into)
    }

    pub async fn get_recommended_timeline(&self, experiment_num: Option<i32>, variant_num: Option<i32>, number: Option<i32>) -> Result<crate::gen::models::PostsResponse, crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::posts_api::get_recommended_timeline(&config, experiment_num, variant_num, number).await.map_err(Into::into)
    }

    pub async fn get_timeline(&self, noreply_mode: crate::gen::models::NoreplyMode, order_by: Option<&str>, experiment_older_age_rules: Option<bool>, from: Option<&str>, from_post_id: Option<i64>, number: Option<i32>, mxn: Option<i32>, en: Option<i32>, vn: Option<i32>, reduce_selfie: Option<bool>, custom_generation_range: Option<bool>) -> Result<crate::gen::models::PostsResponse, crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::posts_api::get_timeline(&config, noreply_mode, order_by, experiment_older_age_rules, from, from_post_id, number, mxn, en, vn, reduce_selfie, custom_generation_range).await.map_err(Into::into)
    }

    pub async fn get_user_timeline(&self, user_id: Option<i64>, from_post_id: Option<i64>, post_type: Option<crate::gen::models::PostType>, number: Option<i32>) -> Result<crate::gen::models::PostsResponse, crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::posts_api::get_user_timeline(&config, user_id, from_post_id, post_type, number).await.map_err(Into::into)
    }

    pub async fn like_posts(&self, post_ids: Option<Vec<i64>>) -> Result<crate::gen::models::LikePostsResponse, crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::posts_api::like_posts(&config, post_ids).await.map_err(Into::into)
    }

    pub async fn move_post_to_specific_thread(&self, id: i64, thread_id: i64) -> Result<crate::gen::models::ThreadInfo, crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::posts_api::move_post_to_specific_thread(&config, id, thread_id).await.map_err(Into::into)
    }

    pub async fn move_post_to_thread(&self, id: i64, thread_icon_filename: Option<&str>, title: Option<&str>) -> Result<crate::gen::models::ThreadInfo, crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::posts_api::move_post_to_thread(&config, id, thread_icon_filename, title).await.map_err(Into::into)
    }

    pub async fn pin_group_post(&self, group_id: Option<i64>, post_id: Option<i64>) -> Result<(), crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::posts_api::pin_group_post(&config, group_id, post_id).await.map_err(Into::into)
    }

    pub async fn report_post(&self, post_id: i64, category_id: Option<i64>, opponent_id: Option<i64>, reason: Option<&str>, screenshot_2_filename: Option<&str>, screenshot_3_filename: Option<&str>, screenshot_4_filename: Option<&str>, screenshot_filename: Option<&str>) -> Result<(), crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::posts_api::report_post(&config, post_id, category_id, opponent_id, reason, screenshot_2_filename, screenshot_3_filename, screenshot_4_filename, screenshot_filename).await.map_err(Into::into)
    }

    pub async fn repost(&self, x_jwt: Option<&str>, attachment_2_filename: Option<&str>, attachment_3_filename: Option<&str>, attachment_4_filename: Option<&str>, attachment_5_filename: Option<&str>, attachment_6_filename: Option<&str>, attachment_7_filename: Option<&str>, attachment_8_filename: Option<&str>, attachment_9_filename: Option<&str>, attachment_filename: Option<&str>, choices: Option<Vec<String>>, color: Option<i32>, font_size: Option<i32>, group_id: Option<i64>, in_reply_to: Option<i64>, language: Option<&str>, mention_ids: Option<Vec<i64>>, message_tags: Option<serde_json::Value>, post_id: Option<i64>, post_type: Option<crate::gen::models::PostType>, shared_url: Option<serde_json::Value>, text: Option<&str>, video_file_name: Option<&str>) -> Result<crate::gen::models::CreatePostResponse, crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::posts_api::repost(&config, x_jwt, attachment_2_filename, attachment_3_filename, attachment_4_filename, attachment_5_filename, attachment_6_filename, attachment_7_filename, attachment_8_filename, attachment_9_filename, attachment_filename, choices, color, font_size, group_id, in_reply_to, language, mention_ids, message_tags, post_id, post_type, shared_url, text, video_file_name).await.map_err(Into::into)
    }

    pub async fn search_posts(&self, keyword: Option<&str>, post_owner_scope: Option<&str>, only_media: Option<bool>, from_post_id: Option<i64>, number: Option<i32>) -> Result<crate::gen::models::PostsResponse, crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::posts_api::search_posts(&config, keyword, post_owner_scope, only_media, from_post_id, number).await.map_err(Into::into)
    }

    pub async fn unlike_post(&self, id: i64) -> Result<(), crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::posts_api::unlike_post(&config, id).await.map_err(Into::into)
    }

    pub async fn unpin_group_post(&self, group_id: Option<i64>) -> Result<(), crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::posts_api::unpin_group_post(&config, group_id).await.map_err(Into::into)
    }

    pub async fn update_post(&self, id: i64, api_key: Option<&str>, color: Option<i32>, font_size: Option<i32>, language: Option<&str>, message_tags: Option<serde_json::Value>, signed_info: Option<&str>, text: Option<&str>, timestamp: Option<i64>, uuid: Option<&str>) -> Result<crate::gen::models::Post, crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::posts_api::update_post(&config, id, api_key, color, font_size, language, message_tags, signed_info, text, timestamp, uuid).await.map_err(Into::into)
    }

    pub async fn validate_post(&self, group_id: Option<i64>, text: Option<&str>, thread_id: Option<i64>) -> Result<crate::gen::models::ValidationPostResponse, crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::posts_api::validate_post(&config, group_id, text, thread_id).await.map_err(Into::into)
    }

    pub async fn view_post_video(&self, video_id: i64) -> Result<(), crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::posts_api::view_post_video(&config, video_id).await.map_err(Into::into)
    }

    pub async fn get_received_gift_senders(&self, gift_id: i64, number: Option<i32>) -> Result<crate::gen::models::GiftSendersResponse, crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::received_gifts_api::get_received_gift_senders(&config, gift_id, number).await.map_err(Into::into)
    }

    pub async fn get_received_gift_transaction(&self, gift_transaction_uuid: &str) -> Result<crate::gen::models::GiftReceivedTransactionResponse, crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::received_gifts_api::get_received_gift_transaction(&config, gift_transaction_uuid).await.map_err(Into::into)
    }

    pub async fn list_received_gifts(&self) -> Result<crate::gen::models::GiftReceivedResponse, crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::received_gifts_api::list_received_gifts(&config).await.map_err(Into::into)
    }

    pub async fn list_received_gifts_v1(&self, from: Option<&str>, number: Option<i32>) -> Result<crate::gen::models::GiftReceivedResponse, crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::received_gifts_api::list_received_gifts_v1(&config, from, number).await.map_err(Into::into)
    }

    pub async fn list_sticker_packs(&self) -> Result<crate::gen::models::StickerPacksResponse, crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::sticker_packs_api::list_sticker_packs(&config).await.map_err(Into::into)
    }

    pub async fn vote_survey(&self, id: i64, choice_id: Option<i64>) -> Result<crate::gen::models::VoteSurveyResponse, crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::surveys_api::vote_survey(&config, id, choice_id).await.map_err(Into::into)
    }

    pub async fn add_thread_member(&self, thread_id: i64, id: i64) -> Result<(), crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::threads_api::add_thread_member(&config, thread_id, id).await.map_err(Into::into)
    }

    pub async fn create_thread(&self, create_group_thread_request: crate::gen::models::CreateGroupThreadRequest) -> Result<crate::gen::models::ThreadInfo, crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::threads_api::create_thread(&config, create_group_thread_request).await.map_err(Into::into)
    }

    pub async fn create_thread_post(&self, id: i64, x_jwt: Option<&str>, attachment_2_filename: Option<&str>, attachment_3_filename: Option<&str>, attachment_4_filename: Option<&str>, attachment_5_filename: Option<&str>, attachment_6_filename: Option<&str>, attachment_7_filename: Option<&str>, attachment_8_filename: Option<&str>, attachment_9_filename: Option<&str>, attachment_filename: Option<&str>, choices: Option<Vec<String>>, color: Option<i32>, font_size: Option<i32>, group_id: Option<i64>, in_reply_to: Option<i64>, language: Option<&str>, mention_ids: Option<Vec<i64>>, message_tags: Option<serde_json::Value>, post_type: Option<crate::gen::models::PostType>, shared_url: Option<serde_json::Value>, text: Option<&str>, video_file_name: Option<&str>) -> Result<crate::gen::models::Post, crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::threads_api::create_thread_post(&config, id, x_jwt, attachment_2_filename, attachment_3_filename, attachment_4_filename, attachment_5_filename, attachment_6_filename, attachment_7_filename, attachment_8_filename, attachment_9_filename, attachment_filename, choices, color, font_size, group_id, in_reply_to, language, mention_ids, message_tags, post_type, shared_url, text, video_file_name).await.map_err(Into::into)
    }

    pub async fn delete_thread(&self, id: i64) -> Result<(), crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::threads_api::delete_thread(&config, id).await.map_err(Into::into)
    }

    pub async fn get_joined_thread_statuses(&self, ids: Option<Vec<i64>>) -> Result<std::collections::HashMap<String, String>, crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::threads_api::get_joined_thread_statuses(&config, ids).await.map_err(Into::into)
    }

    pub async fn get_thread(&self, id: i64) -> Result<crate::gen::models::ThreadInfo, crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::threads_api::get_thread(&config, id).await.map_err(Into::into)
    }

    pub async fn get_thread_posts(&self, id: i64, post_type: Option<crate::gen::models::PostType>, number: Option<i32>, from: Option<i64>) -> Result<crate::gen::models::PostsResponse, crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::threads_api::get_thread_posts(&config, id, post_type, number, from).await.map_err(Into::into)
    }

    pub async fn list_threads(&self, group_id: Option<i64>, from: Option<&str>, join_status: Option<&str>) -> Result<crate::gen::models::GroupThreadListResponse, crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::threads_api::list_threads(&config, group_id, from, join_status).await.map_err(Into::into)
    }

    pub async fn remove_thread_member(&self, thread_id: i64, id: i64) -> Result<(), crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::threads_api::remove_thread_member(&config, thread_id, id).await.map_err(Into::into)
    }

    pub async fn update_thread(&self, id: i64, thread_icon_filename: Option<&str>, title: Option<&str>) -> Result<(), crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::threads_api::update_thread(&config, id, thread_icon_filename, title).await.map_err(Into::into)
    }

    pub async fn agree_policy(&self, r#type: &str) -> Result<(), crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::users_api::agree_policy(&config, r#type).await.map_err(Into::into)
    }

    pub async fn block_user(&self, id: i64) -> Result<(), crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::users_api::block_user(&config, id).await.map_err(Into::into)
    }

    pub async fn change_email(&self, api_key: Option<&str>, email: Option<&str>, email_grant_token: Option<&str>, password: Option<&str>) -> Result<crate::gen::models::LoginUpdateResponse, crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::users_api::change_email(&config, api_key, email, email_grant_token, password).await.map_err(Into::into)
    }

    pub async fn change_password(&self, api_key: Option<&str>, current_password: Option<&str>, password: Option<&str>) -> Result<crate::gen::models::LoginUpdateResponse, crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::users_api::change_password(&config, api_key, current_password, password).await.map_err(Into::into)
    }

    pub async fn create_bookmark(&self, user_id: i64, id: i64) -> Result<crate::gen::models::BookmarkPostResponse, crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::users_api::create_bookmark(&config, user_id, id).await.map_err(Into::into)
    }

    pub async fn create_review(&self, api_key: Option<&str>, comment: Option<&str>, signed_info: Option<&str>, timestamp: Option<i64>, user_ids: Option<Vec<i64>>, uuid: Option<&str>) -> Result<(), crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::users_api::create_review(&config, api_key, comment, signed_info, timestamp, user_ids, uuid).await.map_err(Into::into)
    }

    pub async fn delete_bookmark(&self, user_id: i64, id: i64) -> Result<(), crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::users_api::delete_bookmark(&config, user_id, id).await.map_err(Into::into)
    }

    pub async fn delete_footprint(&self, user_id: i64, footprint_id: i64) -> Result<(), crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::users_api::delete_footprint(&config, user_id, footprint_id).await.map_err(Into::into)
    }

    pub async fn delete_my_reviews(&self, review_ids: Option<Vec<i64>>) -> Result<(), crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::users_api::delete_my_reviews(&config, review_ids).await.map_err(Into::into)
    }

    pub async fn disable_two_factor_auth(&self, code: Option<&str>) -> Result<(), crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::users_api::disable_two_factor_auth(&config, code).await.map_err(Into::into)
    }

    pub async fn edit_user(&self, api_key: Option<&str>, biography: Option<&str>, country_code: Option<&str>, cover_image_filename: Option<&str>, gender: Option<i32>, nickname: Option<&str>, prefecture: Option<&str>, profile_icon_filename: Option<&str>, signed_info: Option<&str>, timestamp: Option<i64>, username: Option<&str>, uuid: Option<&str>) -> Result<(), crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::users_api::edit_user(&config, api_key, biography, country_code, cover_image_filename, gender, nickname, prefecture, profile_icon_filename, signed_info, timestamp, username, uuid).await.map_err(Into::into)
    }

    pub async fn edit_user_v2(&self, api_key: Option<&str>, is_private: Option<bool>, nickname: Option<&str>, signed_info: Option<&str>, timestamp: Option<i64>, uuid: Option<&str>) -> Result<(), crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::users_api::edit_user_v2(&config, api_key, is_private, nickname, signed_info, timestamp, uuid).await.map_err(Into::into)
    }

    pub async fn enable_two_factor_auth(&self, code: Option<&str>, r#type: Option<&str>) -> Result<crate::gen::models::TwoStepAuthEnabledResponse, crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::users_api::enable_two_factor_auth(&config, code, r#type).await.map_err(Into::into)
    }

    pub async fn follow_user(&self, id: i64) -> Result<(), crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::users_api::follow_user(&config, id).await.map_err(Into::into)
    }

    pub async fn follow_users(&self, user_ids: Option<Vec<i64>>) -> Result<(), crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::users_api::follow_users(&config, user_ids).await.map_err(Into::into)
    }

    pub async fn get_active_followings(&self, only_online: Option<bool>, from_loggedin_at: Option<i64>) -> Result<crate::gen::models::ActiveFollowingsResponse, crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::users_api::get_active_followings(&config, only_online, from_loggedin_at).await.map_err(Into::into)
    }

    pub async fn get_additional_notification_setting(&self) -> Result<crate::gen::models::AdditionalSettingsResponse, crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::users_api::get_additional_notification_setting(&config).await.map_err(Into::into)
    }

    pub async fn get_blocked_user_ids(&self) -> Result<crate::gen::models::BlockedUserIdsResponse, crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::users_api::get_blocked_user_ids(&config).await.map_err(Into::into)
    }

    pub async fn get_blocked_users(&self, search_users_request: crate::gen::models::SearchUsersRequest, from_id: Option<i64>) -> Result<crate::gen::models::BlockedUsersResponse, crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::users_api::get_blocked_users(&config, search_users_request, from_id).await.map_err(Into::into)
    }

    pub async fn get_bookmarked_posts(&self, user_id: i64, from: Option<&str>) -> Result<crate::gen::models::PostsResponse, crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::users_api::get_bookmarked_posts(&config, user_id, from).await.map_err(Into::into)
    }

    pub async fn get_chatable_followings(&self, search_users_request: crate::gen::models::SearchUsersRequest, from_follow_id: Option<i64>, from_timestamp: Option<i64>, order_by: Option<&str>) -> Result<crate::gen::models::FollowUsersResponse, crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::users_api::get_chatable_followings(&config, search_users_request, from_follow_id, from_timestamp, order_by).await.map_err(Into::into)
    }

    pub async fn get_default_settings(&self) -> Result<crate::gen::models::DefaultSettingsResponse, crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::users_api::get_default_settings(&config).await.map_err(Into::into)
    }

    pub async fn get_follow_request_count(&self) -> Result<crate::gen::models::FollowRequestCountResponse, crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::users_api::get_follow_request_count(&config).await.map_err(Into::into)
    }

    pub async fn get_follow_requests(&self, from_timestamp: Option<i64>) -> Result<crate::gen::models::UsersByTimestampResponse, crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::users_api::get_follow_requests(&config, from_timestamp).await.map_err(Into::into)
    }

    pub async fn get_followings_born_today(&self, birthdate: Option<i64>) -> Result<crate::gen::models::UsersResponse, crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::users_api::get_followings_born_today(&config, birthdate).await.map_err(Into::into)
    }

    pub async fn get_footprints(&self, mode: Option<&str>, number: Option<i32>, from: Option<&str>) -> Result<crate::gen::models::FootprintsResponse, crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::users_api::get_footprints(&config, mode, number, from).await.map_err(Into::into)
    }

    pub async fn get_fresh_user(&self, id: i64) -> Result<crate::gen::models::UserResponse, crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::users_api::get_fresh_user(&config, id).await.map_err(Into::into)
    }

    pub async fn get_hima_users(&self, from_hima_id: Option<i64>, number: Option<i32>) -> Result<crate::gen::models::HimaUsersResponse, crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::users_api::get_hima_users(&config, from_hima_id, number).await.map_err(Into::into)
    }

    pub async fn get_my_reviews(&self, from_id: Option<i64>) -> Result<crate::gen::models::ReviewsResponse, crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::users_api::get_my_reviews(&config, from_id).await.map_err(Into::into)
    }

    pub async fn get_policy_agreements(&self) -> Result<crate::gen::models::PolicyAgreementsResponse, crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::users_api::get_policy_agreements(&config).await.map_err(Into::into)
    }

    pub async fn get_recommended_follow_users(&self, id: i64, number: Option<i32>, page: Option<i32>) -> Result<crate::gen::models::UsersResponse, crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::users_api::get_recommended_follow_users(&config, id, number, page).await.map_err(Into::into)
    }

    pub async fn get_reset_counters(&self) -> Result<crate::gen::models::RefreshCounterRequestsResponse, crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::users_api::get_reset_counters(&config).await.map_err(Into::into)
    }

    pub async fn get_two_factor_auth_request_info(&self) -> Result<crate::gen::models::TwoStepAuthRequestInfoResponse, crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::users_api::get_two_factor_auth_request_info(&config).await.map_err(Into::into)
    }

    pub async fn get_two_factor_auth_status(&self) -> Result<crate::gen::models::TwoFaStatusResponse, crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::users_api::get_two_factor_auth_status(&config).await.map_err(Into::into)
    }

    pub async fn get_user(&self, id: i64) -> Result<crate::gen::models::UserResponse, crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::users_api::get_user(&config, id).await.map_err(Into::into)
    }

    pub async fn get_user_by_qr(&self, qr: &str) -> Result<crate::gen::models::UserResponse, crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::users_api::get_user_by_qr(&config, qr).await.map_err(Into::into)
    }

    pub async fn get_user_custom_definitions(&self) -> Result<crate::gen::models::UserCustomDefinitionsResponse, crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::users_api::get_user_custom_definitions(&config).await.map_err(Into::into)
    }

    pub async fn get_user_followers(&self, id: i64, from: Option<i64>, followed_by_me: Option<bool>, user_nickname: Option<&str>) -> Result<crate::gen::models::FollowUsersResponse, crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::users_api::get_user_followers(&config, id, from, followed_by_me, user_nickname).await.map_err(Into::into)
    }

    pub async fn get_user_followings(&self, id: i64, from: Option<i64>, from_timestamp: Option<i64>, order_by: Option<&str>, user_nickname: Option<&str>) -> Result<crate::gen::models::FollowUsersResponse, crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::users_api::get_user_followings(&config, id, from, from_timestamp, order_by, user_nickname).await.map_err(Into::into)
    }

    pub async fn get_user_gift_transactions(&self, user_id: i64, number: Option<i32>, from: Option<&str>) -> Result<crate::gen::models::GiftTransactionsResponse, crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::users_api::get_user_gift_transactions(&config, user_id, number, from).await.map_err(Into::into)
    }

    pub async fn get_user_info(&self, id: i64) -> Result<crate::gen::models::UserResponse, crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::users_api::get_user_info(&config, id).await.map_err(Into::into)
    }

    pub async fn get_user_interests(&self) -> Result<crate::gen::models::UserInterestsResponse, crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::users_api::get_user_interests(&config).await.map_err(Into::into)
    }

    pub async fn get_user_presigned_url(&self, video_file_name: Option<&str>) -> Result<crate::gen::models::PresignedUrlResponse, crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::users_api::get_user_presigned_url(&config, video_file_name).await.map_err(Into::into)
    }

    pub async fn get_user_reviews(&self, id: i64, from_id: Option<i64>) -> Result<crate::gen::models::ReviewsResponse, crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::users_api::get_user_reviews(&config, id, from_id).await.map_err(Into::into)
    }

    pub async fn get_user_timestamp(&self) -> Result<crate::gen::models::UserTimestampResponse, crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::users_api::get_user_timestamp(&config).await.map_err(Into::into)
    }

    pub async fn get_users_by_ids(&self, x_jwt: Option<&str>, user_ids: Option<Vec<i64>>) -> Result<crate::gen::models::UsersResponse, crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::users_api::get_users_by_ids(&config, x_jwt, user_ids).await.map_err(Into::into)
    }

    pub async fn get_web_socket_token(&self) -> Result<crate::gen::models::WebSocketTokenResponse, crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::users_api::get_web_socket_token(&config).await.map_err(Into::into)
    }

    pub async fn logout(&self, uuid: Option<&str>) -> Result<(), crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::users_api::logout(&config, uuid).await.map_err(Into::into)
    }

    pub async fn ping_alive(&self) -> Result<(), crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::users_api::ping_alive(&config).await.map_err(Into::into)
    }

    pub async fn remove_cover_image(&self) -> Result<(), crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::users_api::remove_cover_image(&config).await.map_err(Into::into)
    }

    pub async fn remove_profile_photo(&self) -> Result<(), crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::users_api::remove_profile_photo(&config).await.map_err(Into::into)
    }

    pub async fn reply_to_review(&self, id: i64, api_key: Option<&str>, comment: Option<&str>, context: Option<&str>, signed_info: Option<&str>, timestamp: Option<i64>, uuid: Option<&str>) -> Result<(), crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::users_api::reply_to_review(&config, id, api_key, comment, context, signed_info, timestamp, uuid).await.map_err(Into::into)
    }

    pub async fn report_user(&self, user_id: i64, category_id: Option<i64>, reason: Option<&str>, screenshot_2_filename: Option<&str>, screenshot_3_filename: Option<&str>, screenshot_4_filename: Option<&str>, screenshot_filename: Option<&str>) -> Result<(), crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::users_api::report_user(&config, user_id, category_id, reason, screenshot_2_filename, screenshot_3_filename, screenshot_4_filename, screenshot_filename).await.map_err(Into::into)
    }

    pub async fn request_follow(&self, target_id: i64, action: Option<&str>) -> Result<(), crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::users_api::request_follow(&config, target_id, action).await.map_err(Into::into)
    }

    pub async fn resend_confirm_email(&self) -> Result<(), crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::users_api::resend_confirm_email(&config).await.map_err(Into::into)
    }

    pub async fn reset_counters(&self, counter: Option<&str>) -> Result<(), crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::users_api::reset_counters(&config, counter).await.map_err(Into::into)
    }

    pub async fn reset_password(&self, email: Option<&str>, email_grant_token: Option<&str>, password: Option<&str>) -> Result<(), crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::users_api::reset_password(&config, email, email_grant_token, password).await.map_err(Into::into)
    }

    pub async fn search_users(&self, gender: Option<i32>, nickname: Option<&str>, title: Option<&str>, biography: Option<&str>, from_timestamp: Option<i64>, similar_age: Option<bool>, not_recent_gomimushi: Option<bool>, recently_created: Option<bool>, same_prefecture: Option<bool>, save_recent_search: Option<bool>) -> Result<crate::gen::models::UsersResponse, crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::users_api::search_users(&config, gender, nickname, title, biography, from_timestamp, similar_age, not_recent_gomimushi, recently_created, same_prefecture, save_recent_search).await.map_err(Into::into)
    }

    pub async fn set_hima(&self) -> Result<(), crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::users_api::set_hima(&config).await.map_err(Into::into)
    }

    pub async fn unblock_user(&self, id: i64) -> Result<(), crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::users_api::unblock_user(&config, id).await.map_err(Into::into)
    }

    pub async fn unfollow_user(&self, id: i64) -> Result<(), crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::users_api::unfollow_user(&config, id).await.map_err(Into::into)
    }

    pub async fn update_additional_notification_setting(&self, mode: Option<&str>, on: Option<i32>) -> Result<(), crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::users_api::update_additional_notification_setting(&config, mode, on).await.map_err(Into::into)
    }

    pub async fn update_language(&self, api_key: Option<&str>, language: Option<&str>, signed_info: Option<&str>, timestamp: Option<i64>, uuid: Option<&str>) -> Result<(), crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::users_api::update_language(&config, api_key, language, signed_info, timestamp, uuid).await.map_err(Into::into)
    }

    pub async fn update_login(&self, api_key: Option<&str>, current_password: Option<&str>, email: Option<&str>, email_grant_token: Option<&str>, password: Option<&str>) -> Result<crate::gen::models::LoginUpdateResponse, crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::users_api::update_login(&config, api_key, current_password, email, email_grant_token, password).await.map_err(Into::into)
    }

    pub async fn update_user_interests(&self, common_ids_request: crate::gen::models::CommonIdsRequest) -> Result<(), crate::ApiError> {
        let config = self.gen_configuration();
        crate::gen::apis::users_api::update_user_interests(&config, common_ids_request).await.map_err(Into::into)
    }

}
