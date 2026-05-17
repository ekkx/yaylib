# Code generated; DO NOT EDIT.
#
# Flat operation facade mixed into Client: every operation is
# reachable as client.<op>, delegating to the per-service
# attribute. Client's own hand-written methods override these
# via MRO (the embed-shadowing rule); such ops are also omitted
# here. The plain def returns the coroutine for the caller to
# await, matching the generated op's own call style.


class GeneratedFacade:
    def get_user_activities(self, *args, **kwargs):
        return self.activities_api.get_user_activities(*args, **kwargs)

    def get_user_activities_v1(self, *args, **kwargs):
        return self.activities_api.get_user_activities_v1(*args, **kwargs)

    def get_app_config(self, *args, **kwargs):
        return self.apps_api.get_app_config(*args, **kwargs)

    def oauth_token(self, *args, **kwargs):
        return self.auth_api.oauth_token(*args, **kwargs)

    def get_bucket_presigned_urls(self, *args, **kwargs):
        return self.buckets_api.get_bucket_presigned_urls(*args, **kwargs)

    def bulk_invite_to_call(self, *args, **kwargs):
        return self.calls_api.bulk_invite_to_call(*args, **kwargs)

    def bump_call(self, *args, **kwargs):
        return self.calls_api.bump_call(*args, **kwargs)

    def generate_call_action_signature(self, *args, **kwargs):
        return self.calls_api.generate_call_action_signature(*args, **kwargs)

    def get_agora_rtm_token(self, *args, **kwargs):
        return self.calls_api.get_agora_rtm_token(*args, **kwargs)

    def get_call_bgms(self, *args, **kwargs):
        return self.calls_api.get_call_bgms(*args, **kwargs)

    def get_call_gift_history(self, *args, **kwargs):
        return self.calls_api.get_call_gift_history(*args, **kwargs)

    def get_conference_call(self, *args, **kwargs):
        return self.calls_api.get_conference_call(*args, **kwargs)

    def get_invitable_call_users(self, *args, **kwargs):
        return self.calls_api.get_invitable_call_users(*args, **kwargs)

    def get_phone_status(self, *args, **kwargs):
        return self.calls_api.get_phone_status(*args, **kwargs)

    def invite_to_call(self, *args, **kwargs):
        return self.calls_api.invite_to_call(*args, **kwargs)

    def invite_to_conference_call(self, *args, **kwargs):
        return self.calls_api.invite_to_conference_call(*args, **kwargs)

    def kick_from_conference_call(self, *args, **kwargs):
        return self.calls_api.kick_from_conference_call(*args, **kwargs)

    def leave_agora_channel(self, *args, **kwargs):
        return self.calls_api.leave_agora_channel(*args, **kwargs)

    def leave_conference_call(self, *args, **kwargs):
        return self.calls_api.leave_conference_call(*args, **kwargs)

    def start_conference_call(self, *args, **kwargs):
        return self.calls_api.start_conference_call(*args, **kwargs)

    def update_call(self, *args, **kwargs):
        return self.calls_api.update_call(*args, **kwargs)

    def update_call_user(self, *args, **kwargs):
        return self.calls_api.update_call_user(*args, **kwargs)

    def validate_call_action_signature(self, *args, **kwargs):
        return self.calls_api.validate_call_action_signature(*args, **kwargs)

    def accept_chat_request(self, *args, **kwargs):
        return self.chat_rooms_api.accept_chat_request(*args, **kwargs)

    def create_chat_room(self, *args, **kwargs):
        return self.chat_rooms_api.create_chat_room(*args, **kwargs)

    def create_chat_room_v1(self, *args, **kwargs):
        return self.chat_rooms_api.create_chat_room_v1(*args, **kwargs)

    def delete_chat_message(self, *args, **kwargs):
        return self.chat_rooms_api.delete_chat_message(*args, **kwargs)

    def delete_chat_rooms(self, *args, **kwargs):
        return self.chat_rooms_api.delete_chat_rooms(*args, **kwargs)

    def get_chat_messages(self, *args, **kwargs):
        return self.chat_rooms_api.get_chat_messages(*args, **kwargs)

    def get_chat_request_count(self, *args, **kwargs):
        return self.chat_rooms_api.get_chat_request_count(*args, **kwargs)

    def get_chat_requests(self, *args, **kwargs):
        return self.chat_rooms_api.get_chat_requests(*args, **kwargs)

    def get_chat_room(self, *args, **kwargs):
        return self.chat_rooms_api.get_chat_room(*args, **kwargs)

    def get_chat_unread_status(self, *args, **kwargs):
        return self.chat_rooms_api.get_chat_unread_status(*args, **kwargs)

    def get_main_chat_rooms(self, *args, **kwargs):
        return self.chat_rooms_api.get_main_chat_rooms(*args, **kwargs)

    def get_updated_chat_rooms(self, *args, **kwargs):
        return self.chat_rooms_api.get_updated_chat_rooms(*args, **kwargs)

    def invite_to_chat_room(self, *args, **kwargs):
        return self.chat_rooms_api.invite_to_chat_room(*args, **kwargs)

    def kick_from_chat_room(self, *args, **kwargs):
        return self.chat_rooms_api.kick_from_chat_room(*args, **kwargs)

    def pin_chat_room(self, *args, **kwargs):
        return self.chat_rooms_api.pin_chat_room(*args, **kwargs)

    def read_chat_attachments(self, *args, **kwargs):
        return self.chat_rooms_api.read_chat_attachments(*args, **kwargs)

    def read_chat_message(self, *args, **kwargs):
        return self.chat_rooms_api.read_chat_message(*args, **kwargs)

    def read_chat_videos(self, *args, **kwargs):
        return self.chat_rooms_api.read_chat_videos(*args, **kwargs)

    def remove_chat_room_background(self, *args, **kwargs):
        return self.chat_rooms_api.remove_chat_room_background(*args, **kwargs)

    def report_chat_room(self, *args, **kwargs):
        return self.chat_rooms_api.report_chat_room(*args, **kwargs)

    def send_chat_message(self, *args, **kwargs):
        return self.chat_rooms_api.send_chat_message(*args, **kwargs)

    def unpin_chat_room(self, *args, **kwargs):
        return self.chat_rooms_api.unpin_chat_room(*args, **kwargs)

    def update_chat_room(self, *args, **kwargs):
        return self.chat_rooms_api.update_chat_room(*args, **kwargs)

    def get_conversation(self, *args, **kwargs):
        return self.conversations_api.get_conversation(*args, **kwargs)

    def get_root_posts(self, *args, **kwargs):
        return self.conversations_api.get_root_posts(*args, **kwargs)

    def request_email_verification(self, *args, **kwargs):
        return self.email_verification_urls_api.request_email_verification(*args, **kwargs)

    def list_game_apps(self, *args, **kwargs):
        return self.games_api.list_game_apps(*args, **kwargs)

    def list_game_walkthroughs(self, *args, **kwargs):
        return self.games_api.list_game_walkthroughs(*args, **kwargs)

    def list_genres(self, *args, **kwargs):
        return self.genres_api.list_genres(*args, **kwargs)

    def list_gifts(self, *args, **kwargs):
        return self.gifts_api.list_gifts(*args, **kwargs)

    def list_muted_group_users(self, *args, **kwargs):
        return self.group_mute_api.list_muted_group_users(*args, **kwargs)

    def mute_group_user(self, *args, **kwargs):
        return self.group_mute_api.mute_group_user(*args, **kwargs)

    def unmute_group_user(self, *args, **kwargs):
        return self.group_mute_api.unmute_group_user(*args, **kwargs)

    def accept_group_join_request(self, *args, **kwargs):
        return self.groups_api.accept_group_join_request(*args, **kwargs)

    def accept_group_transfer(self, *args, **kwargs):
        return self.groups_api.accept_group_transfer(*args, **kwargs)

    def ban_group_user(self, *args, **kwargs):
        return self.groups_api.ban_group_user(*args, **kwargs)

    def cancel_group_transfer(self, *args, **kwargs):
        return self.groups_api.cancel_group_transfer(*args, **kwargs)

    def create_group(self, *args, **kwargs):
        return self.groups_api.create_group(*args, **kwargs)

    def decline_group_join_request(self, *args, **kwargs):
        return self.groups_api.decline_group_join_request(*args, **kwargs)

    def deputize_group_users(self, *args, **kwargs):
        return self.groups_api.deputize_group_users(*args, **kwargs)

    def deputize_group_users_mass(self, *args, **kwargs):
        return self.groups_api.deputize_group_users_mass(*args, **kwargs)

    def fire_group_user(self, *args, **kwargs):
        return self.groups_api.fire_group_user(*args, **kwargs)

    def get_group(self, *args, **kwargs):
        return self.groups_api.get_group(*args, **kwargs)

    def get_group_ban_list(self, *args, **kwargs):
        return self.groups_api.get_group_ban_list(*args, **kwargs)

    def get_group_create_quota(self, *args, **kwargs):
        return self.groups_api.get_group_create_quota(*args, **kwargs)

    def get_group_gift_history(self, *args, **kwargs):
        return self.groups_api.get_group_gift_history(*args, **kwargs)

    def get_group_gift_transactions(self, *args, **kwargs):
        return self.groups_api.get_group_gift_transactions(*args, **kwargs)

    def get_group_highlights(self, *args, **kwargs):
        return self.groups_api.get_group_highlights(*args, **kwargs)

    def get_group_member(self, *args, **kwargs):
        return self.groups_api.get_group_member(*args, **kwargs)

    def get_group_members(self, *args, **kwargs):
        return self.groups_api.get_group_members(*args, **kwargs)

    def get_group_received_gift_senders(self, *args, **kwargs):
        return self.groups_api.get_group_received_gift_senders(*args, **kwargs)

    def get_group_unread_status(self, *args, **kwargs):
        return self.groups_api.get_group_unread_status(*args, **kwargs)

    def get_invitable_group_users(self, *args, **kwargs):
        return self.groups_api.get_invitable_group_users(*args, **kwargs)

    def get_joined_group_statuses(self, *args, **kwargs):
        return self.groups_api.get_joined_group_statuses(*args, **kwargs)

    def get_relatable_groups(self, *args, **kwargs):
        return self.groups_api.get_relatable_groups(*args, **kwargs)

    def get_related_groups(self, *args, **kwargs):
        return self.groups_api.get_related_groups(*args, **kwargs)

    def get_user_group_list(self, *args, **kwargs):
        return self.groups_api.get_user_group_list(*args, **kwargs)

    def invite_to_group(self, *args, **kwargs):
        return self.groups_api.invite_to_group(*args, **kwargs)

    def join_group(self, *args, **kwargs):
        return self.groups_api.join_group(*args, **kwargs)

    def leave_group(self, *args, **kwargs):
        return self.groups_api.leave_group(*args, **kwargs)

    def list_group_categories(self, *args, **kwargs):
        return self.groups_api.list_group_categories(*args, **kwargs)

    def list_groups(self, *args, **kwargs):
        return self.groups_api.list_groups(*args, **kwargs)

    def list_my_groups(self, *args, **kwargs):
        return self.groups_api.list_my_groups(*args, **kwargs)

    def pin_group_highlight_post(self, *args, **kwargs):
        return self.groups_api.pin_group_highlight_post(*args, **kwargs)

    def remove_group_cover(self, *args, **kwargs):
        return self.groups_api.remove_group_cover(*args, **kwargs)

    def remove_group_deputies(self, *args, **kwargs):
        return self.groups_api.remove_group_deputies(*args, **kwargs)

    def remove_group_icon(self, *args, **kwargs):
        return self.groups_api.remove_group_icon(*args, **kwargs)

    def remove_related_groups(self, *args, **kwargs):
        return self.groups_api.remove_related_groups(*args, **kwargs)

    def report_group(self, *args, **kwargs):
        return self.groups_api.report_group(*args, **kwargs)

    def request_group_walkthrough(self, *args, **kwargs):
        return self.groups_api.request_group_walkthrough(*args, **kwargs)

    def search_group_posts(self, *args, **kwargs):
        return self.groups_api.search_group_posts(*args, **kwargs)

    def set_group_title(self, *args, **kwargs):
        return self.groups_api.set_group_title(*args, **kwargs)

    def take_over_group(self, *args, **kwargs):
        return self.groups_api.take_over_group(*args, **kwargs)

    def transfer_group(self, *args, **kwargs):
        return self.groups_api.transfer_group(*args, **kwargs)

    def unban_group_user(self, *args, **kwargs):
        return self.groups_api.unban_group_user(*args, **kwargs)

    def unpin_group_highlight_post(self, *args, **kwargs):
        return self.groups_api.unpin_group_highlight_post(*args, **kwargs)

    def update_group(self, *args, **kwargs):
        return self.groups_api.update_group(*args, **kwargs)

    def update_related_groups(self, *args, **kwargs):
        return self.groups_api.update_related_groups(*args, **kwargs)

    def visit_group(self, *args, **kwargs):
        return self.groups_api.visit_group(*args, **kwargs)

    def withdraw_group_deputy(self, *args, **kwargs):
        return self.groups_api.withdraw_group_deputy(*args, **kwargs)

    def withdraw_group_transfer(self, *args, **kwargs):
        return self.groups_api.withdraw_group_transfer(*args, **kwargs)

    def create_mute_keyword(self, *args, **kwargs):
        return self.hidden_api.create_mute_keyword(*args, **kwargs)

    def delete_mute_keyword(self, *args, **kwargs):
        return self.hidden_api.delete_mute_keyword(*args, **kwargs)

    def hide_chats(self, *args, **kwargs):
        return self.hidden_api.hide_chats(*args, **kwargs)

    def hide_users(self, *args, **kwargs):
        return self.hidden_api.hide_users(*args, **kwargs)

    def list_hidden_chats(self, *args, **kwargs):
        return self.hidden_api.list_hidden_chats(*args, **kwargs)

    def list_hidden_users(self, *args, **kwargs):
        return self.hidden_api.list_hidden_users(*args, **kwargs)

    def list_mute_keywords(self, *args, **kwargs):
        return self.hidden_api.list_mute_keywords(*args, **kwargs)

    def unhide_chats(self, *args, **kwargs):
        return self.hidden_api.unhide_chats(*args, **kwargs)

    def unhide_users(self, *args, **kwargs):
        return self.hidden_api.unhide_users(*args, **kwargs)

    def get_banned_words(self, *args, **kwargs):
        return self.moderation_api.get_banned_words(*args, **kwargs)

    def get_popular_words(self, *args, **kwargs):
        return self.moderation_api.get_popular_words(*args, **kwargs)

    def get_group_notification_settings(self, *args, **kwargs):
        return self.notification_settings_api.get_group_notification_settings(*args, **kwargs)

    def update_chat_room_notification_settings(self, *args, **kwargs):
        return self.notification_settings_api.update_chat_room_notification_settings(*args, **kwargs)

    def update_group_notification_settings(self, *args, **kwargs):
        return self.notification_settings_api.update_group_notification_settings(*args, **kwargs)

    def pin_group(self, *args, **kwargs):
        return self.pinned_api.pin_group(*args, **kwargs)

    def pin_post(self, *args, **kwargs):
        return self.pinned_api.pin_post(*args, **kwargs)

    def pin_review(self, *args, **kwargs):
        return self.pinned_api.pin_review(*args, **kwargs)

    def unpin_group(self, *args, **kwargs):
        return self.pinned_api.unpin_group(*args, **kwargs)

    def unpin_review(self, *args, **kwargs):
        return self.pinned_api.unpin_review(*args, **kwargs)

    def create_conference_call_post(self, *args, **kwargs):
        return self.posts_api.create_conference_call_post(*args, **kwargs)

    def create_post(self, *args, **kwargs):
        return self.posts_api.create_post(*args, **kwargs)

    def create_share_post(self, *args, **kwargs):
        return self.posts_api.create_share_post(*args, **kwargs)

    def delete_all_posts(self, *args, **kwargs):
        return self.posts_api.delete_all_posts(*args, **kwargs)

    def delete_posts(self, *args, **kwargs):
        return self.posts_api.delete_posts(*args, **kwargs)

    def get_active_call_post(self, *args, **kwargs):
        return self.posts_api.get_active_call_post(*args, **kwargs)

    def get_call_followers_timeline(self, *args, **kwargs):
        return self.posts_api.get_call_followers_timeline(*args, **kwargs)

    def get_call_timeline(self, *args, **kwargs):
        return self.posts_api.get_call_timeline(*args, **kwargs)

    def get_following_timeline(self, *args, **kwargs):
        return self.posts_api.get_following_timeline(*args, **kwargs)

    def get_group_timeline(self, *args, **kwargs):
        return self.posts_api.get_group_timeline(*args, **kwargs)

    def get_my_posts(self, *args, **kwargs):
        return self.posts_api.get_my_posts(*args, **kwargs)

    def get_post(self, *args, **kwargs):
        return self.posts_api.get_post(*args, **kwargs)

    def get_post_gift_transactions(self, *args, **kwargs):
        return self.posts_api.get_post_gift_transactions(*args, **kwargs)

    def get_post_likers(self, *args, **kwargs):
        return self.posts_api.get_post_likers(*args, **kwargs)

    def get_post_reposts(self, *args, **kwargs):
        return self.posts_api.get_post_reposts(*args, **kwargs)

    def get_post_url_metadata(self, *args, **kwargs):
        return self.posts_api.get_post_url_metadata(*args, **kwargs)

    def get_posts_by_ids(self, *args, **kwargs):
        return self.posts_api.get_posts_by_ids(*args, **kwargs)

    def get_posts_by_tag(self, *args, **kwargs):
        return self.posts_api.get_posts_by_tag(*args, **kwargs)

    def get_recent_engagement_posts(self, *args, **kwargs):
        return self.posts_api.get_recent_engagement_posts(*args, **kwargs)

    def get_recommended_post_tags(self, *args, **kwargs):
        return self.posts_api.get_recommended_post_tags(*args, **kwargs)

    def get_recommended_timeline(self, *args, **kwargs):
        return self.posts_api.get_recommended_timeline(*args, **kwargs)

    def get_timeline(self, *args, **kwargs):
        return self.posts_api.get_timeline(*args, **kwargs)

    def get_user_timeline(self, *args, **kwargs):
        return self.posts_api.get_user_timeline(*args, **kwargs)

    def like_posts(self, *args, **kwargs):
        return self.posts_api.like_posts(*args, **kwargs)

    def move_post_to_specific_thread(self, *args, **kwargs):
        return self.posts_api.move_post_to_specific_thread(*args, **kwargs)

    def move_post_to_thread(self, *args, **kwargs):
        return self.posts_api.move_post_to_thread(*args, **kwargs)

    def pin_group_post(self, *args, **kwargs):
        return self.posts_api.pin_group_post(*args, **kwargs)

    def report_post(self, *args, **kwargs):
        return self.posts_api.report_post(*args, **kwargs)

    def repost(self, *args, **kwargs):
        return self.posts_api.repost(*args, **kwargs)

    def search_posts(self, *args, **kwargs):
        return self.posts_api.search_posts(*args, **kwargs)

    def unlike_post(self, *args, **kwargs):
        return self.posts_api.unlike_post(*args, **kwargs)

    def unpin_group_post(self, *args, **kwargs):
        return self.posts_api.unpin_group_post(*args, **kwargs)

    def update_post(self, *args, **kwargs):
        return self.posts_api.update_post(*args, **kwargs)

    def validate_post(self, *args, **kwargs):
        return self.posts_api.validate_post(*args, **kwargs)

    def view_post_video(self, *args, **kwargs):
        return self.posts_api.view_post_video(*args, **kwargs)

    def get_received_gift_senders(self, *args, **kwargs):
        return self.received_gifts_api.get_received_gift_senders(*args, **kwargs)

    def get_received_gift_transaction(self, *args, **kwargs):
        return self.received_gifts_api.get_received_gift_transaction(*args, **kwargs)

    def list_received_gifts(self, *args, **kwargs):
        return self.received_gifts_api.list_received_gifts(*args, **kwargs)

    def list_received_gifts_v1(self, *args, **kwargs):
        return self.received_gifts_api.list_received_gifts_v1(*args, **kwargs)

    def list_sticker_packs(self, *args, **kwargs):
        return self.sticker_packs_api.list_sticker_packs(*args, **kwargs)

    def vote_survey(self, *args, **kwargs):
        return self.surveys_api.vote_survey(*args, **kwargs)

    def add_thread_member(self, *args, **kwargs):
        return self.threads_api.add_thread_member(*args, **kwargs)

    def create_thread(self, *args, **kwargs):
        return self.threads_api.create_thread(*args, **kwargs)

    def create_thread_post(self, *args, **kwargs):
        return self.threads_api.create_thread_post(*args, **kwargs)

    def delete_thread(self, *args, **kwargs):
        return self.threads_api.delete_thread(*args, **kwargs)

    def get_joined_thread_statuses(self, *args, **kwargs):
        return self.threads_api.get_joined_thread_statuses(*args, **kwargs)

    def get_thread(self, *args, **kwargs):
        return self.threads_api.get_thread(*args, **kwargs)

    def get_thread_posts(self, *args, **kwargs):
        return self.threads_api.get_thread_posts(*args, **kwargs)

    def list_threads(self, *args, **kwargs):
        return self.threads_api.list_threads(*args, **kwargs)

    def remove_thread_member(self, *args, **kwargs):
        return self.threads_api.remove_thread_member(*args, **kwargs)

    def update_thread(self, *args, **kwargs):
        return self.threads_api.update_thread(*args, **kwargs)

    def agree_policy(self, *args, **kwargs):
        return self.users_api.agree_policy(*args, **kwargs)

    def block_user(self, *args, **kwargs):
        return self.users_api.block_user(*args, **kwargs)

    def change_email(self, *args, **kwargs):
        return self.users_api.change_email(*args, **kwargs)

    def change_password(self, *args, **kwargs):
        return self.users_api.change_password(*args, **kwargs)

    def create_bookmark(self, *args, **kwargs):
        return self.users_api.create_bookmark(*args, **kwargs)

    def create_review(self, *args, **kwargs):
        return self.users_api.create_review(*args, **kwargs)

    def delete_bookmark(self, *args, **kwargs):
        return self.users_api.delete_bookmark(*args, **kwargs)

    def delete_footprint(self, *args, **kwargs):
        return self.users_api.delete_footprint(*args, **kwargs)

    def delete_my_reviews(self, *args, **kwargs):
        return self.users_api.delete_my_reviews(*args, **kwargs)

    def disable_two_factor_auth(self, *args, **kwargs):
        return self.users_api.disable_two_factor_auth(*args, **kwargs)

    def edit_user(self, *args, **kwargs):
        return self.users_api.edit_user(*args, **kwargs)

    def edit_user_v2(self, *args, **kwargs):
        return self.users_api.edit_user_v2(*args, **kwargs)

    def enable_two_factor_auth(self, *args, **kwargs):
        return self.users_api.enable_two_factor_auth(*args, **kwargs)

    def follow_user(self, *args, **kwargs):
        return self.users_api.follow_user(*args, **kwargs)

    def follow_users(self, *args, **kwargs):
        return self.users_api.follow_users(*args, **kwargs)

    def get_active_followings(self, *args, **kwargs):
        return self.users_api.get_active_followings(*args, **kwargs)

    def get_additional_notification_setting(self, *args, **kwargs):
        return self.users_api.get_additional_notification_setting(*args, **kwargs)

    def get_blocked_user_ids(self, *args, **kwargs):
        return self.users_api.get_blocked_user_ids(*args, **kwargs)

    def get_blocked_users(self, *args, **kwargs):
        return self.users_api.get_blocked_users(*args, **kwargs)

    def get_bookmarked_posts(self, *args, **kwargs):
        return self.users_api.get_bookmarked_posts(*args, **kwargs)

    def get_chatable_followings(self, *args, **kwargs):
        return self.users_api.get_chatable_followings(*args, **kwargs)

    def get_default_settings(self, *args, **kwargs):
        return self.users_api.get_default_settings(*args, **kwargs)

    def get_follow_request_count(self, *args, **kwargs):
        return self.users_api.get_follow_request_count(*args, **kwargs)

    def get_follow_requests(self, *args, **kwargs):
        return self.users_api.get_follow_requests(*args, **kwargs)

    def get_followings_born_today(self, *args, **kwargs):
        return self.users_api.get_followings_born_today(*args, **kwargs)

    def get_footprints(self, *args, **kwargs):
        return self.users_api.get_footprints(*args, **kwargs)

    def get_fresh_user(self, *args, **kwargs):
        return self.users_api.get_fresh_user(*args, **kwargs)

    def get_hima_users(self, *args, **kwargs):
        return self.users_api.get_hima_users(*args, **kwargs)

    def get_my_reviews(self, *args, **kwargs):
        return self.users_api.get_my_reviews(*args, **kwargs)

    def get_policy_agreements(self, *args, **kwargs):
        return self.users_api.get_policy_agreements(*args, **kwargs)

    def get_recommended_follow_users(self, *args, **kwargs):
        return self.users_api.get_recommended_follow_users(*args, **kwargs)

    def get_reset_counters(self, *args, **kwargs):
        return self.users_api.get_reset_counters(*args, **kwargs)

    def get_two_factor_auth_request_info(self, *args, **kwargs):
        return self.users_api.get_two_factor_auth_request_info(*args, **kwargs)

    def get_two_factor_auth_status(self, *args, **kwargs):
        return self.users_api.get_two_factor_auth_status(*args, **kwargs)

    def get_user(self, *args, **kwargs):
        return self.users_api.get_user(*args, **kwargs)

    def get_user_by_qr(self, *args, **kwargs):
        return self.users_api.get_user_by_qr(*args, **kwargs)

    def get_user_custom_definitions(self, *args, **kwargs):
        return self.users_api.get_user_custom_definitions(*args, **kwargs)

    def get_user_followers(self, *args, **kwargs):
        return self.users_api.get_user_followers(*args, **kwargs)

    def get_user_followings(self, *args, **kwargs):
        return self.users_api.get_user_followings(*args, **kwargs)

    def get_user_gift_transactions(self, *args, **kwargs):
        return self.users_api.get_user_gift_transactions(*args, **kwargs)

    def get_user_info(self, *args, **kwargs):
        return self.users_api.get_user_info(*args, **kwargs)

    def get_user_interests(self, *args, **kwargs):
        return self.users_api.get_user_interests(*args, **kwargs)

    def get_user_presigned_url(self, *args, **kwargs):
        return self.users_api.get_user_presigned_url(*args, **kwargs)

    def get_user_reviews(self, *args, **kwargs):
        return self.users_api.get_user_reviews(*args, **kwargs)

    def get_user_timestamp(self, *args, **kwargs):
        return self.users_api.get_user_timestamp(*args, **kwargs)

    def get_users_by_ids(self, *args, **kwargs):
        return self.users_api.get_users_by_ids(*args, **kwargs)

    def get_web_socket_token(self, *args, **kwargs):
        return self.users_api.get_web_socket_token(*args, **kwargs)

    def logout(self, *args, **kwargs):
        return self.users_api.logout(*args, **kwargs)

    def ping_alive(self, *args, **kwargs):
        return self.users_api.ping_alive(*args, **kwargs)

    def remove_cover_image(self, *args, **kwargs):
        return self.users_api.remove_cover_image(*args, **kwargs)

    def remove_profile_photo(self, *args, **kwargs):
        return self.users_api.remove_profile_photo(*args, **kwargs)

    def reply_to_review(self, *args, **kwargs):
        return self.users_api.reply_to_review(*args, **kwargs)

    def report_user(self, *args, **kwargs):
        return self.users_api.report_user(*args, **kwargs)

    def request_follow(self, *args, **kwargs):
        return self.users_api.request_follow(*args, **kwargs)

    def resend_confirm_email(self, *args, **kwargs):
        return self.users_api.resend_confirm_email(*args, **kwargs)

    def reset_counters(self, *args, **kwargs):
        return self.users_api.reset_counters(*args, **kwargs)

    def reset_password(self, *args, **kwargs):
        return self.users_api.reset_password(*args, **kwargs)

    def search_users(self, *args, **kwargs):
        return self.users_api.search_users(*args, **kwargs)

    def set_hima(self, *args, **kwargs):
        return self.users_api.set_hima(*args, **kwargs)

    def unblock_user(self, *args, **kwargs):
        return self.users_api.unblock_user(*args, **kwargs)

    def unfollow_user(self, *args, **kwargs):
        return self.users_api.unfollow_user(*args, **kwargs)

    def update_additional_notification_setting(self, *args, **kwargs):
        return self.users_api.update_additional_notification_setting(*args, **kwargs)

    def update_language(self, *args, **kwargs):
        return self.users_api.update_language(*args, **kwargs)

    def update_login(self, *args, **kwargs):
        return self.users_api.update_login(*args, **kwargs)

    def update_user_interests(self, *args, **kwargs):
        return self.users_api.update_user_interests(*args, **kwargs)

