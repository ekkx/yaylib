from enum import Enum


class Configs:

    YAYLIB_VERSION = "0.1.2"
    YAY_API_VERSION = "3.16"
    YAY_VERSION_NAME = "3.16.1"
    YAY_API_VERSION_KEY = "e83a1d2588918c2061280427c88e6f56"
    YAY_API_KEY = "ccd59ee269c01511ba763467045c115779fcae3050238a252f1bd1a4b65cfec6"
    YAY_SHARED_KEY = "yayZ1"
    YAY_STORE_KEY = "yayZ1payment"
    ID_CARD_CHECK_SECRET_KEY = "4aa6d1c301a97154bc1098c2"
    YAY_REVIEW_HOST_1 = "review.yay.space"
    YAY_REVIEW_HOST_2 = "cas-stg.yay.space"
    YAY_STAGING_HOST_1 = "stg.yay.space"
    YAY_STAGING_HOST_2 = "cas.yay.space"
    YAY_PRODUCTION_HOST = "api.yay.space"
    YAY_API_URL = "https://" + YAY_PRODUCTION_HOST
    ID_CARD_CHECK_HOST_PRODUCTION = "idcardcheck.com"
    ID_CARD_CHECK_HOST_STAGING = "stg.idcardcheck.com"
    USER_AGENT = "android 11 (3.5x 1440x2960 Galaxy S9)"
    REQUEST_HEADERS = {
        "Host": YAY_PRODUCTION_HOST,
        "X-App-Version": YAY_API_VERSION,
        "User-Agent": USER_AGENT,
        "X-Device-Info": f"yay {YAY_VERSION_NAME} {USER_AGENT}",
        "X-Device-Uuid": "",
        "X-Connection-Type": "wifi",
        "Accept-Language": "ja",
        "Content-Type": "application/json;charset=UTF-8"
    }


class Endpoints:

    # api v1 endpoints
    USERS_V1 = Configs.YAY_API_URL + "/v1/users"
    PAYMENTS_V1 = Configs.YAY_API_URL + "/v1/payments"
    THREADS_V1 = Configs.YAY_API_URL + "/v1/threads"
    PINNED_V1 = Configs.YAY_API_URL + "/v1/pinned"
    POSTS_V1 = Configs.YAY_API_URL + "/v1/posts"
    CONVERSATIONS_V1 = Configs.YAY_API_URL + "/v1/conversations"
    HIDDEN_V1 = Configs.YAY_API_URL + "/v1/hidden"
    GROUPS_V1 = Configs.YAY_API_URL + "/v1/groups"
    CHAT_ROOMS_V1 = Configs.YAY_API_URL + "/v1/chat_rooms"
    CALLS_V1 = Configs.YAY_API_URL + "/v1/calls"
    SURVEYS_V1 = Configs.YAY_API_URL + "/v1/surveys"
    FRIENDS_V1 = Configs.YAY_API_URL + "/v1/friends"
    GAMES_V1 = Configs.YAY_API_URL + "/v1/games"
    GENRES_V1 = Configs.YAY_API_URL + "/v1/genres"
    ANONYMOUS_CALLS_V1 = Configs.YAY_API_URL + "/v1/anonymous_calls"
    NOTIFICATION_SETTINGS_V1 = Configs.YAY_API_URL + "/v1/notification_settings"
    STICKER_PACKS_V1 = Configs.YAY_API_URL + "/v1/sticker_packs"
    SNS_THUMBNAIL_V1 = Configs.YAY_API_URL + "/v1/sns_thumbnail"
    EMAIL_VERIFICATION_URL_V1 = Configs.YAY_API_URL + "/v1/email_verification_urls"
    BUCKETS_V1 = Configs.YAY_API_URL + "/v1/buckets"
    ID_CHECK_V1 = Configs.YAY_API_URL + "/v1/id_check"
    PROMOTIONS_V1 = Configs.YAY_API_URL + "/v1/promotions"
    SKYFALL_V1 = Configs.YAY_API_URL + "/v1/skyfall"
    GENUINE_DEVICES_V1 = Configs.YAY_API_URL + "/v1/genuine_devices"
    LOBI_FRIENDS_V1 = Configs.YAY_API_URL + "/v1/lobi_friends"

    # api v2 endpoints
    USERS_V2 = Configs.YAY_API_URL + "/v2/users"
    PAYMENTS_V2 = Configs.YAY_API_URL + "/v2/payments"
    THREADS_V2 = Configs.YAY_API_URL + "/v2/threads"
    PINNED_V2 = Configs.YAY_API_URL + "/v2/pinned"
    POSTS_V2 = Configs.YAY_API_URL + "/v2/posts"
    CONVERSATIONS_V2 = Configs.YAY_API_URL + "/v2/conversations"
    HIDDEN_V2 = Configs.YAY_API_URL + "/v2/hidden"
    GROUPS_V2 = Configs.YAY_API_URL + "/v2/groups"
    CHAT_ROOMS_V2 = Configs.YAY_API_URL + "/v2/chat_rooms"
    CALLS_V2 = Configs.YAY_API_URL + "/v2/calls"
    SURVEYS_V2 = Configs.YAY_API_URL + "/v2/surveys"
    FRIENDS_V2 = Configs.YAY_API_URL + "/v2/friends"
    GAMES_V2 = Configs.YAY_API_URL + "/v2/games"
    GENRES_V2 = Configs.YAY_API_URL + "/v2/genres"
    ANONYMOUS_CALLS_V2 = Configs.YAY_API_URL + "/v2/anonymous_calls"
    NOTIFICATION_SETTINGS_V2 = Configs.YAY_API_URL + "/v2/notification_settings"
    STICKER_PACKS_V2 = Configs.YAY_API_URL + "/v2/sticker_packs"
    SNS_THUMBNAIL_V2 = Configs.YAY_API_URL + "/v2/sns_thumbnail"
    EMAIL_VERIFICATION_URL_V2 = Configs.YAY_API_URL + "/v2/email_verification_urls"
    BUCKETS_V2 = Configs.YAY_API_URL + "/v2/buckets"
    ID_CHECK_V2 = Configs.YAY_API_URL + "/v2/id_check"
    PROMOTIONS_V2 = Configs.YAY_API_URL + "/v2/promotions"
    SKYFALL_V2 = Configs.YAY_API_URL + "/v2/skyfall"
    GENUINE_DEVICES_V2 = Configs.YAY_API_URL + "/v2/genuine_devices"
    LOBI_FRIENDS_V2 = Configs.YAY_API_URL + "/v2/lobi_friends"

    # api v3 endpoints
    USERS_V3 = Configs.YAY_API_URL + "/v3/users"
    PAYMENTS_V3 = Configs.YAY_API_URL + "/v3/payments"
    THREADS_V3 = Configs.YAY_API_URL + "/v3/threads"
    PINNED_V3 = Configs.YAY_API_URL + "/v3/pinned"
    POSTS_V3 = Configs.YAY_API_URL + "/v3/posts"
    CONVERSATIONS_V3 = Configs.YAY_API_URL + "/v3/conversations"
    HIDDEN_V3 = Configs.YAY_API_URL + "/v3/hidden"
    GROUPS_V3 = Configs.YAY_API_URL + "/v3/groups"
    CHAT_ROOMS_V3 = Configs.YAY_API_URL + "/v3/chat_rooms"
    CALLS_V3 = Configs.YAY_API_URL + "/v3/calls"
    SURVEYS_V3 = Configs.YAY_API_URL + "/v3/surveys"
    FRIENDS_V3 = Configs.YAY_API_URL + "/v3/friends"
    GAMES_V3 = Configs.YAY_API_URL + "/v3/games"
    GENRES_V3 = Configs.YAY_API_URL + "/v3/genres"
    ANONYMOUS_CALLS_V3 = Configs.YAY_API_URL + "/v3/anonymous_calls"
    NOTIFICATION_SETTINGS_V3 = Configs.YAY_API_URL + "/v3/notification_settings"
    STICKER_PACKS_V3 = Configs.YAY_API_URL + "/v3/sticker_packs"
    SNS_THUMBNAIL_V3 = Configs.YAY_API_URL + "/v3/sns_thumbnail"
    EMAIL_VERIFICATION_URL_V3 = Configs.YAY_API_URL + "/v3/email_verification_urls"
    BUCKETS_V3 = Configs.YAY_API_URL + "/v3/buckets"
    ID_CHECK_V3 = Configs.YAY_API_URL + "/v3/id_check"
    PROMOTIONS_V3 = Configs.YAY_API_URL + "/v3/promotions"
    SKYFALL_V3 = Configs.YAY_API_URL + "/v3/skyfall"
    GENUINE_DEVICES_V3 = Configs.YAY_API_URL + "/v3/genuine_devices"
    LOBI_FRIENDS_V3 = Configs.YAY_API_URL + "/v3/lobi_friends"


class LogMessage(Enum):

    # <https://yay.space/locales/en/translation.json?v=3.18.0>
    # <https://yay.space/locales/ja/translation.json?v=3.18.0>

    # logger(self, fname="invite_users_to_call")

    # -CALL
    bump_call = {
        "en": "Call bumped.",
        "ja": ""
    }
    invite_to_call_bulk = {
        "en": "Invited your online followings to the call",
        "ja": "オンラインを友達をまとめて通話に招待しました"
    }
    invite_users_to_call = {
        "en": "Invitation sent.",
        "ja": ""
    }
    invite_users_to_chat_call = {
        "en": "Invitation of chat call sent.",
        "ja": ""
    }
    kick_and_ban_from_call = {
        "en": "User has been banned from the call.",
        "ja": ""
    }
    notify_anonymous_user_leave_agora_channel = {
        "en": "Notified anonymous user left the call.",
        "ja": ""
    }
    notify_user_leave_agora_channel = {
        "en": "Notified user '{user_id}' left the call.",
        "ja": ""
    }
    send_call_screenshot = {
        "en": "Call screenshot sent.",
        "ja": ""
    }
    set_call = {
        "en": "Call set.",
        "ja": ""
    }
    set_user_role = {
        "en": "User has been given a role.",
        "ja": ""
    }
    start_call = {
        "en": "Joined the call.",
        "ja": ""
    }
    stop_call = {
        "en": "Left the call.",
        "ja": ""
    }

    # -CHAT
    accept_chat_request = {
        "en": "Chat request accepted.",
        "ja": ""
    }
    create_group_chat = {
        "en": "Group chat '{name}' created.",
        "ja": ""
    }
    create_private_chat = {
        "en": "Private chatroom with '{with_user_id}' created.",
        "ja": ""
    }
    delete_background = {
        "en": "Background deleted.",
        "ja": ""
    }
    delete_message = {
        "en": "Message deleted.",
        "ja": ""
    }
    edit_chat_room = {
        "en": "Chatroom updated.",
        "ja": ""
    }
    invite_to_chat = {
        "en": "Chatroom invitation sent.",
        "ja": ""
    }
    kick_users_from_chat = {
        "en": "Users have been kicked from the chatroom.",
        "ja": ""
    }
    pin_chat = {
        "en": "Pinned chatroom.",
        "ja": ""
    }
    read_attachment = {
        "en": "Attachment has been read.",
        "ja": ""
    }
    read_message = {
        "en": "Message has been read.",
        "ja": ""
    }
    read_video_message = {
        "en": "Video message has been read.",
        "ja": ""
    }
    remove_chat_rooms = {
        "en": "Chat rooms removed.",
        "ja": ""
    }
    report_chat_room = {
        "en": "Chatroom '{chat_room_id}' has been reported.",
        "ja": ""
    }
    send_media_screenshot_notification = {
        "en": "Media screenshot notification sent.",
        "ja": ""
    }
    send_message = {
        "en": "Your message has been sent.",
        "ja": ""
    }
    unhide_chat = {
        "en": "Chatroom unhidden.",
        "ja": ""
    }
    unpin_chat = {
        "en": "Chatroom unpinned.",
        "ja": ""
    }

    # -GROUP
    accept_moderator_offer = {
        "en": "Moderator offer accepted.",
        "ja": ""
    }
    accept_ownership_offer = {
        "en": "Accepted ownership offer.",
        "ja": ""
    }
    accept_group_join_request = {
        "en": "Accepted group join request.",
        "ja": ""
    }
    add_related_groups = {
        "en": "Group added to the related groups",
        "ja": ""
    }
    ban_group_user = {
        "en": "User has been banned from the group.",
        "ja": ""
    }
    create_group = {
        "en": "Group created.",
        "ja": ""
    }
    create_pin_group = {
        "en": "Pinned post in the group.",
        "ja": ""
    }
    decline_moderator_offer = {
        "en": "Declined moderator offer.",
        "ja": ""
    }
    decline_ownership_offer = {
        "en": "Declined ownership offer.",
        "ja": ""
    }
    decline_group_join_request = {
        "en": "Declined group join request",
        "ja": ""
    }
    delete_pin_group = {
        "en": "Unpinned group post.",
        "ja": ""
    }
    invite_users_to_group = {
        "en": "Group invitation sent.",
        "ja": ""
    }
    join_group = {
        "en": "Joined the group.",
        "ja": ""
    }
    leave_group = {
        "en": "Left the group.",
        "ja": ""
    }
    post_gruop_social_shared = {
        "en": "Group social shared posted.",
        "ja": ""
    }
    remove_group_cover = {
        "en": "Group cover removed.",
        "ja": ""
    }
    remove_moderator = {
        "en": "Group moderator removed.",
        "ja": ""
    }
    remove_related_groups = {
        "en": "Related groups removed.",
        "ja": ""
    }
    report_group = {
        "en": "Group reported.",
        "ja": ""
    }
    send_moderator_offers = {
        "en": "Group moderator offer sent.",
        "ja": ""
    }
    send_ownership_offer = {
        "en": "Group ownership offer sent.",
        "ja": ""
    }
    set_group_title = {
        "en": "Group title set.",
        "ja": ""
    }
    take_over_group_ownership = {
        "en": "Took over the group ownership of {group_id}",
        "ja": ""
    }
    unban_group_member = {
        "en": "Group user has been unbanned.",
        "ja": ""
    }
    update_group = {
        "en": "Group details have been updated.",
        "ja": ""
    }
    visit_group = {
        "en": "Visited the group.",
        "ja": ""
    }
    withdraw_moderator_offer = {
        "en": "Group moderator offer withdrawn",
        "ja": ""
    }
    withdraw_ownership_offer = {
        "en": "Group ownershipt offer withdrawn",
        "ja": ""
    }

    # -LOGIN
    change_email = {
        "en": "Your email has been changed.",
        "ja": "実行完了"
    }
    change_password = {
        "en": "Password has been changed.",
        "ja": "実行完了"
    }
    login = {
        "en": "Successfully logged in as '{user_id}'",
        "ja": "実行完了"
    }
    logout = {
        "en": "User has Logged out.",
        "ja": "実行完了"
    }
    restore_user = {
        "en": "User restored.",
        "ja": "実行完了"
    }
    revoke_tokens = {
        "en": "Token revoked.",
        "ja": "実行完了"
    }
    save_account_with_email = {
        "en": "Account has been save with email.",
        "ja": "実行完了"
    }

    # -MISC
    accept_policy_agreement = {
        "en": "Accepted to policy agreement.",
        "ja": "実行完了"
    }
    generate_sns_thumbnail = {
        "en": "SNS thumbnail generated.",
        "ja": "実行完了"
    }
    verify_device = {
        "en": "Device verified.",
        "ja": "実行完了"
    }
    upload_image = {
        "en": "Image '{filename}{ext}' is uploaded",
        "ja": "実行完了"
    }

    # -POST
    add_bookmark = {
        "en": "Post added to bookmarks.",
        "ja": "実行完了"
    }
    add_group_highlight_post = {
        "en": "Post added to group highlight.",
        "ja": "実行完了"
    }
    create_call_post = {
        "en": "Call post created.",
        "ja": "実行完了"
    }
    create_group_pin_post = {
        "en": "Pinned post in a group.",
        "ja": "実行完了"
    }
    create_pin_post = {
        "en": "Pinned post.",
        "ja": "実行完了"
    }
    create_post = {
        "en": "Post has been created.",
        "ja": "実行完了"
    }
    create_repost = {
        "en": "Repost has been created.",
        "ja": "実行完了"
    }
    create_share_post = {
        "en": "Share post has been created.",
        "ja": "実行完了"
    }
    create_thread_post = {
        "en": "Thread post has been created.",
        "ja": "実行完了"
    }
    delete_all_post = {
        "en": "Post deletion requested.",
        "ja": "実行完了"
    }
    delete_all_post_not_found = {
        "en": "Post not found.",
        "ja": "投稿が見つかりません"
    }
    delete_group_pin_post = {
        "en": "Unpinned post in a group.",
        "ja": "実行完了"
    }
    delete_pin_post = {
        "en": "Unpinned post.",
        "ja": "実行完了"
    }
    like_posts = {
        "en": "Posts liked.",
        "ja": "実行完了"
    }
    remove_bookmark = {
        "en": "Bookmark has been removed.",
        "ja": "実行完了"
    }
    remove_group_highlight_post = {
        "en": "Group hightlight post removed.",
        "ja": "実行完了"
    }
    remove_posts = {
        "en": "Posts have been removed.",
        "ja": "実行完了"
    }
    report_post = {
        "en": "Post has been reported.",
        "ja": "実行完了"
    }
    unlike_post = {
        "en": "Post unliked.",
        "ja": "実行完了"
    }
    update_post = {
        "en": "Post has been updated.",
        "ja": "実行完了"
    }
    update_recommendation_feedback = {
        "en": "Recommendation feedback updated.",
        "ja": "実行完了"
    }
    validate_post = {
        "en": "Post validated.",
        "ja": "実行完了"
    }
    view_video = {
        "en": "Viewed the video",
        "ja": "実行完了"
    }
    vote_survey = {
        "en": "Survey voted.",
        "ja": "実行完了"
    }

    # -REVIEW
    send_call_screenshot = {
        "en": "aaaaaaaaaaaaaaaaaa",
        "ja": "実行完了"
    }
