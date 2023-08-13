"""
MIT License

Copyright (c) 2023-present Qvco, Konn

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
"""

from typing import Dict, List, Union

from .api import API
from .api.call import (
    bump_call,
    get_bgms,
    get_call,
    get_call_invitable_users,
    get_call_status,
    get_games,
    get_genres,
    get_group_calls,
    get_user_active_call,
    invite_to_call_bulk,
    invite_users_to_call,
    invite_users_to_chat_call,
    kick_and_ban_from_call,
    set_call,
    set_user_role,
    start_call,
    start_anonymous_call,
    stop_call,
    stop__anonymous_call,
)
from .api.cassandra import (
    get_user_activities,
    get_user_merged_activities,
    received_notification,
)
from .api.chat import (
    accept_chat_requests,
    check_unread_status,
    create_group_chat,
    create_private_chat,
    delete_background,
    delete_message,
    edit_chat_room,
    get_chatable_users,
    get_gifs_data,
    get_hidden_chat_rooms,
    get_main_chat_rooms,
    get_messages,
    get_request_chat_rooms,
    get_chat_room,
    get_sticker_packs,
    get_total_chat_requests,
    hide_chat,
    invite_to_chat,
    kick_users_from_chat,
    pin_chat,
    read_message,
    refresh_chat_rooms,
    remove_chat_rooms,
    report_chat_room,
    send_message,
    unhide_chat,
    unpin_chat,
)
from .api.group import (
    accept_moderator_offer,
    accept_ownership_offer,
    accept_group_join_request,
    add_related_groups,
    ban_group_user,
    check_unread_status,
    create_group,
    create_pin_group,
    decline_moderator_offer,
    decline_ownership_offer,
    decline_group_join_request,
    delete_pin_group,
    get_banned_group_members,
    get_group_categories,
    get_create_group_quota,
    get_group,
    get_groups,
    get_invitable_users,
    get_joined_statuses,
    get_group_member,
    get_group_members,
    get_my_groups,
    get_relatable_groups,
    get_related_groups,
    get_user_groups,
    invite_users_to_group,
    join_group,
    leave_group,
    post_gruop_social_shared,
    remove_group_cover,
    remove_moderator,
    remove_related_groups,
    report_group,
    send_moderator_offers,
    send_ownership_offer,
    set_group_title,
    take_over_group_ownership,
    unban_group_member,
    update_group,
    withdraw_moderator_offer,
    withdraw_ownership_offer,
)
from .api.login import (
    change_email,
    change_password,
    get_token,
    is_valid_token,
    login_with_email,
    logout,
    resend_confirm_email,
    restore_user,
    revoke_tokens,
    save_account_with_email,
)
from .api.misc import (
    accept_policy_agreement,
    generate_sns_thumbnail,
    send_verification_code,
    get_email_grant_token,
    get_email_verification_presigned_url,
    get_file_upload_presigned_urls,
    get_id_checker_presigned_url,
    get_old_file_upload_presigned_url,
    get_web_socket_token,
    verify_device,
    upload_image,
    upload_video,
    get_app_config,
    get_banned_words,
    get_popular_words,
)
from .api.post import (
    add_bookmark,
    add_group_highlight_post,
    create_call_post,
    create_group_pin_post,
    create_pin_post,
    mention,
    create_post,
    create_repost,
    create_share_post,
    create_thread_post,
    delete_all_post,
    delete_group_pin_post,
    delete_pin_post,
    get_bookmark,
    get_timeline_calls,
    get_conversation,
    get_conversation_root_posts,
    get_following_call_timeline,
    get_following_timeline,
    get_group_highlight_posts,
    get_group_timeline_by_keyword,
    get_group_timeline,
    get_timeline_by_hashtag,
    get_my_posts,
    get_post,
    get_post_likers,
    get_post_reposts,
    get_posts,
    get_recommended_post_tags,
    get_recommended_posts,
    get_timeline_by_keyword,
    get_timeline,
    get_url_metadata,
    get_user_timeline,
    like_posts,
    remove_bookmark,
    remove_group_highlight_post,
    remove_posts,
    report_post,
    unlike_post,
    update_post,
    view_video,
    vote_survey,
)
from .api.review import (
    create_review,
    create_reviews,
    delete_reviews,
    get_my_reviews,
    get_reviews,
    pin_review,
    unpin_review,
)
from .api.thread import (
    add_post_to_thread,
    convert_post_to_thread,
    create_thread,
    get_group_thread_list,
    get_thread_joined_statuses,
    get_thread_posts,
    join_thread,
    leave_thread,
    remove_thread,
    update_thread,
)
from .api.user import (
    delete_footprint,
    destroy_user,
    follow_user,
    follow_users,
    get_active_followings,
    get_follow_recommendations,
    get_follow_request,
    get_follow_request_count,
    get_following_users_born,
    get_footprints,
    get_fresh_user,
    get_hima_users,
    get_user_ranking,
    get_refresh_counter_requests,
    get_social_shared_users,
    get_timestamp,
    get_user,
    get_user_email,
    get_user_followers,
    get_user_followings,
    get_user_from_qr,
    get_user_without_leaving_footprint,
    get_users,
    refresh_counter,
    register,
    remove_user_avatar,
    remove_user_cover,
    report_user,
    reset_password,
    search_lobi_users,
    search_users,
    set_follow_permission_enabled,
    set_setting_follow_recommendation_enabled,
    take_action_follow_request,
    turn_on_hima,
    unfollow_user,
    update_language,
    update_user,
    block_user,
    get_blocked_user_ids,
    get_blocked_users,
    unblock_user,
    get_hidden_users_list,
    hide_user,
    unhide_users,
)
from .models import (
    ApplicationConfig,
    Attachment,
    BanWord,
    Bgm,
    ChatRoom,
    ConferenceCall,
    CreateGroupQuota,
    Footprint,
    GifImageCategory,
    Group,
    Message,
    Post,
    PresignedUrl,
    PopularWord,
    SharedUrl,
    StickerPack,
    Survey,
    ThreadInfo,
    User,
    UserWrapper,
)
from .responses import (
    ActiveFollowingsResponse,
    ActivitiesResponse,
    BlockedUserIdsResponse,
    BlockedUsersResponse,
    BookmarkPostResponse,
    CallStatusResponse,
    ChatRoomsResponse,
    CreateGroupResponse,
    CreateChatRoomResponse,
    FollowRecommendationsResponse,
    FollowUsersResponse,
    GamesResponse,
    GenresResponse,
    GroupCategoriesResponse,
    GroupsRelatedResponse,
    GroupsResponse,
    GroupThreadListResponse,
    GroupUserResponse,
    GroupUsersResponse,
    HiddenResponse,
    LoginUserResponse,
    LoginUpdateResponse,
    MessageResponse,
    PostsResponse,
    PostLikersResponse,
    PostTagsResponse,
    LikePostsResponse,
    RefreshCounterRequestsResponse,
    ReviewsResponse,
    SocialShareUsersResponse,
    TokenResponse,
    UnreadStatusResponse,
    UserResponse,
    UsersResponse,
    RankingUsersResponse,
    UsersByTimestampResponse,
    UserTimestampResponse,
)


POST_TYPE_TEXT = "text"
r""" `text`: テキストのみ投稿タイプ"""

POST_TYPE_MEDIA = "media"
r""" `media`: メディアを含める投稿タイプ"""

POST_TYPE_IMAGE = "image"
r""" `image`: 画像を含める投稿タイプ"""

POST_TYPE_VIDEO = "video"
r""" `video`: ビデオを含める投稿タイプ"""

POST_TYPE_SURVEY = "survey"
r""" `survey`: アンケートを含める投稿タイプ"""

POST_TYPE_CALL = "call"
r""" `call`: 通話用の投稿タイプ"""

POST_TYPE_SHAREABLE_URL = "shareable_url"
r""" `shareable_url`: サークルやスレッド共有用の投稿タイプ"""


CALL_TYPE_VOICE = "voice"
r""" `voice`: 音声通話用の通話タイプ"""

CALL_TYPE_VIDEO = "vdo"
r""" `vdo`: ビデオ通話用の通話タイプ"""


IMAGE_TYPE_POST = "post"
r""" `post`: 投稿に画像をアップロードする際の画像タイプ"""

IMAGE_TYPE_CHAT_MESSAGE = "chat_message"
r""" `chat_message`: 個人チャットに画像をアップロードする際の画像タイプ"""

IMAGE_TYPE_CHAT_BACKGROUND = "chat_background"
r""" `chat_background`: 個人チャットの背景用に画像をアップロードする際の画像タイプ"""

IMAGE_TYPE_REPORT = "report"
r""" `report`: 通報用の画像をアップロードする際の画像タイプ"""

IMAGE_TYPE_USER_AVATAR = "user_avatar"
r""" `user_avatar`: プロフィール画像をアップロードする際の画像タイプ"""

IMAGE_TYPE_USER_COVER = "user_cover"
r""" `user_cover`: プロフィールの背景画像をアップロードする際の画像タイプ"""

IMAGE_TYPE_GROUP_COVER = "group_cover"
r""" `group_cover`: グループの背景画像をアップロードする際の画像タイプ"""

IMAGE_TYPE_GROUP_THREAD_ICON = "group_thread_icon"
r""" `group_thread_icon`: グループ内のスレッド用アイコンをアップロードする際の画像タイプ"""

IMAGE_TYPE_GROUP_ICON = "group_icon"
r""" `group_icon`: グループのアイコンをアップロードする際の画像タイプ"""


SHAREABLE_TYPE_GROUP = "group"
r""" `group`: サークル用の共有タイプ"""

SHAREABLE_TYPE_THREAD = "thread"
r""" `thread`: スレッド用の共有タイプ"""


class Client(API):
    """

    Yay!（イェイ）の API クライアント

    #### Useage

        >>> api = yaylib.Client()
        >>> timeline = api.get_timeline(number=30)

    #### Parameters

        - access_token: str - (optional)
        - proxy: str - (optional)
        - max_retries: int - (optional)
        - backoff_factor: float - (optional)
        - timeout: int - (optional)
        - err_lang: str - (optional)
        - base_path: str - (optional)
        - save_cookie_file: bool - (optional)
        - encrypt_cookie: bool - (optional)
        - cookie_filename: str - (optional)
        - loglevel: int - (optional)

    <https://github.com/qvco/yaylib>

    ---

    ### Yay! (nanameue, Inc.) API Client

    Copyright (c) 2023-present Qvco, Konn

    """

    # -CALL

    def get_user_active_call(self, user_id: int, access_token: str = None) -> Post:
        """

        ユーザーが参加中の通話を取得します

        """
        return get_user_active_call(self, user_id, access_token)

    def get_bgms(self, access_token: str = None) -> List[Bgm]:
        """

        通話のBGMを取得します

        """
        return get_bgms(self, access_token)

    def get_call(self, call_id: int, access_token: str = None) -> ConferenceCall:
        """

        通話の詳細を取得します

        """
        return get_call(self, call_id, access_token)

    def get_call_invitable_users(
        self, call_id: int, from_timestamp: int = None, access_token: str = None
    ) -> UsersByTimestampResponse:
        # @Nullable @Query("user[nickname]")
        """

        通話に招待可能なユーザーを取得します

        """
        return get_call_invitable_users(self, call_id, from_timestamp, access_token)

    def get_call_status(
        self, opponent_id: int, access_token: str = None
    ) -> CallStatusResponse:
        """

        通話の状態を取得します

        """
        return get_call_status(self, opponent_id, access_token)

    def get_games(self, access_token: str = None, **params) -> GamesResponse:
        """

        ゲームを取得します

        Parameters
        ----------
            - number: int - (optional)
            - ids: List[int] - (optional)
            - from_id: int - (optional)

        """
        return get_games(self, access_token, **params)

    def get_genres(self, access_token: str = None, **params) -> GenresResponse:
        """

        通話のジャンルを取得します

        Parameters
        ----------
            - number: int - (optional)
            - from: int - (optional)

        """
        return get_genres(self, access_token, **params)

    def get_group_calls(self, access_token: str = None, **params) -> PostsResponse:
        """

        サークルの通話を取得します

        Parameters
        ----------
            - number: int - (optional)
            - group_category_id: int - (optional)
            - from_timestamp: int - (optional)
            - scope: str - (optional)

        """
        return get_group_calls(self, access_token, **params)

    def invite_online_followings_to_call(
        self, call_id: int, group_id: int = None, access_token: str = None
    ) -> dict:
        """

        オンラインの友達をまとめて通話に招待します

        Parameters
        ----------
            - call_id: int - (required)
            - group_id: int - (optional)

        """
        return invite_to_call_bulk(self, call_id, group_id, access_token)

    def invite_users_to_call(
        self, call_id: int, user_ids: List[int], access_token: str = None
    ) -> dict:
        """

        ユーザーを通話に招待します

        Parameters
        ----------
            - call_id: int - (required)
            - user_ids: List[int] - (required)

        """
        return invite_users_to_call(self, call_id, user_ids, access_token)

    def invite_users_to_chat_call(
        self, chat_room_id: int, room_id: int, room_url: str, access_token: str = None
    ) -> dict:
        """

        チャット通話にユーザーを招待します

        Parameters
        ----------

            - chat_room_id: int - (required)
            - room_id: int - (required)
            - room_url: int - (required)

        """
        return invite_users_to_chat_call(
            self, chat_room_id, room_id, room_url, access_token
        )

    def kick_user_from_call(
        self, call_id: int, user_id: int, access_token: str = None
    ) -> dict:
        """

        ユーザーを通話からキックします

        Parameters
        ----------

            - call_id: int - (required)
            - user_id: int - (required)

        """
        return kick_and_ban_from_call(self, call_id, user_id, access_token)

    def set_call(
        self,
        call_id: int,
        joinable_by: str,
        game_title: str = None,
        category_id: str = None,
        access_token: str = None,
    ) -> dict:
        """

        通話を開始します

        """
        return set_call(
            self, call_id, joinable_by, game_title, category_id, access_token
        )

    def set_user_role(
        self, call_id: int, user_id: int, role: str, access_token: str = None
    ) -> dict:
        """

        通話に参加中ののユーザーに役職を与えます

        """
        return set_user_role(self, call_id, user_id, role, access_token)

    def join_call(
        self, conference_id: int, call_sid: str = None, access_token: str = None
    ) -> ConferenceCall:
        """

        通話に参加します

        """
        return start_call(self, conference_id, call_sid, access_token)

    def join_call_as_anonymous(
        self, conference_id: int, agora_uid: str
    ) -> ConferenceCall:
        """

        無名くんとして通話に参加します

        """
        return start_anonymous_call(self, conference_id, agora_uid)

    def leave_call(
        self, conference_id: int, call_sid: str = None, access_token: str = None
    ) -> dict:
        """

        通話から退出します

        """
        return stop_call(self, conference_id, call_sid, access_token)

    def leave_call_as_anonymous(self, conference_id: int, agora_uid: str = None):
        """

        通話から退出します

        """
        return stop__anonymous_call(self, conference_id, agora_uid)

    # -CASSANDRA

    def get_activities(self, access_token: str = None, **params) -> ActivitiesResponse:
        """

        通知を取得します

        Parameters
        ----------
            - important: bool - (required)
            - from_timestamp: int - (optional)
            - number: int - (optional)

        """
        return get_user_activities(self, access_token, **params)

    def get_merged_activities(
        self, access_token: str = None, **params
    ) -> ActivitiesResponse:
        """

        全種類の通知を取得します

        Parameters
        ----------

            - from_timestamp: int - (optional)
            - number: int - (optional)

        """
        return get_user_merged_activities(self, access_token, **params)

    def received_notification(
        self, pid: str, type: str, opened_at: int = None, access_token: str = None
    ) -> dict:
        return received_notification(self, pid, type, opened_at, access_token)

    # -CHAT

    def accept_chat_requests(
        self, chat_room_ids: List[int], access_token: str = None
    ) -> dict:
        """

        チャットリクエストを承認します

        Parameters
        ----------

            - chat_room_ids: List[int] - (required)

        """
        return accept_chat_requests(self, chat_room_ids, access_token)

    def check_unread_status(
        self, from_time: int, access_token: str = None
    ) -> UnreadStatusResponse:
        """

        チャットの未読ステータスを確認します

        """
        return check_unread_status(self, from_time, access_token)

    def create_group_chat(
        self,
        name: str,
        with_user_ids: List[int],
        icon_filename: str = None,
        background_filename: str = None,
        access_token: str = None,
    ) -> CreateChatRoomResponse:
        """

        グループチャットを作成します

        Parameters
        ----------

            - name: str - (required)
            - with_user_ids: List[int] - (required)
            - icon_filename: str - (optional)
            - background_filename: str - (optional)

        """
        return create_group_chat(
            self, name, with_user_ids, icon_filename, background_filename, access_token
        )

    def create_private_chat(
        self,
        with_user_id: int,
        matching_id: int = None,
        hima_chat: bool = False,
        access_token: str = None,
    ) -> CreateChatRoomResponse:
        """

        個人チャットを作成します

        Parameters
        ----------

            - with_user_id: int - (required)
            - matching_id: str - (optional)
            - hima_chat: bool - (optional)

        """
        return create_private_chat(
            self, with_user_id, matching_id, hima_chat, access_token
        )

    def delete_background(self, room_id: int, access_token: str = None) -> dict:
        """

        チャットの背景画像を削除します

        Parameters
        ----------

            - room_id: int - (required)

        """
        return delete_background(self, room_id, access_token)

    def delete_message(
        self, room_id: int, message_id: int, access_token: str = None
    ) -> dict:
        """

        メッセージを削除します

        Parameters
        ----------

            - room_id: int - (required)
            - message_id: int - (required)

        """
        return delete_message(self, room_id, message_id, access_token)

    def edit_chat_room(
        self,
        chat_room_id: int,
        name: str,
        icon_filename: str = None,
        background_filename: str = None,
        access_token: str = None,
    ) -> dict:
        """

        チャットルームを編集します

        """
        return edit_chat_room(
            self, chat_room_id, name, icon_filename, background_filename, access_token
        )

    def get_chatable_users(
        self,
        from_follow_id: int = None,
        from_timestamp: int = None,
        order_by: str = None,
        access_token: str = None,
    ) -> FollowUsersResponse:
        """

        チャット可能なユーザーを取得します

        """
        return get_chatable_users(
            from_follow_id, from_timestamp, order_by, access_token
        )

    def get_gifs_data(self, access_token: str = None) -> List[GifImageCategory]:
        """

        チャットルームのGIFデータを取得します

        """
        return get_gifs_data(self, access_token)

    def get_hidden_chat_rooms(
        self, access_token: str = None, **params
    ) -> ChatRoomsResponse:
        """

        非表示のチャットルームを取得します

        Parameters
        ----------

            - from_timestamp: int - (optional)
            - number: int - (optional)

        """
        return get_hidden_chat_rooms(self, access_token, **params)

    def get_main_chat_rooms(
        self, from_timestamp: int = None, access_token: str = None
    ) -> ChatRoomsResponse:
        """

        メインのチャットルームを取得します

        """
        return get_main_chat_rooms(self, from_timestamp, access_token)

    def get_messages(
        self, chat_room_id: int, access_token: str = None, **params
    ) -> List[Message]:
        """

        メッセージを取得します

        Parameters
        ---------------
            - from_message_id: int - (optional)
            - to_message_id: int - (optional)

        """
        return get_messages(self, chat_room_id, access_token, **params)

    def get_chat_requests(
        self, access_token: str = None, **params
    ) -> ChatRoomsResponse:
        """

        チャットリクエストを取得します

        Parameters:
        ----------

            - number: int (optional)
            - from_timestamp: int (optional)
            - access_token: str (optional)

        """
        return get_request_chat_rooms(self, access_token, **params)

    def get_chat_room(self, chat_room_id: int, access_token: str = None) -> ChatRoom:
        """

        チャットルームの詳細を取得します

        """
        return get_chat_room(self, chat_room_id, access_token)

    def get_sticker_packs(self, access_token: str = None) -> List[StickerPack]:
        """

        スタンプを取得します

        """
        return get_sticker_packs(self, access_token)

    def get_chat_requests_count(self, access_token: str = None) -> int:
        """

        チャットリクエストの数を取得します

        """
        return get_total_chat_requests(self, access_token)

    def hide_chat(self, chat_room_id: int, access_token: str = None) -> dict:
        """

        チャットルームを非表示にします

        """
        return hide_chat(self, chat_room_id, access_token)

    def invite_to_chat(
        self, chat_room_id: int, user_ids: List[int], access_token: str = None
    ) -> dict:
        """

        チャットルームにユーザーを招待します

        """
        return invite_to_chat(self, chat_room_id, user_ids, access_token)

    def kick_users_from_chat(
        self, chat_room_id: int, user_ids: List[int], access_token: str = None
    ) -> dict:
        """

        チャットルームからユーザーを追放します

        """
        return kick_users_from_chat(self, chat_room_id, user_ids, access_token)

    def pin_chat(self, room_id: int, access_token: str = None) -> dict:
        """

        チャットルームをピン止めします

        """
        return pin_chat(self, room_id, access_token)

    def read_message(
        self, chat_room_id: int, message_id: int, access_token: str = None
    ) -> dict:
        """

        メッセージを既読にします

        """
        return read_message(self, chat_room_id, message_id, access_token)

    def refresh_chat_rooms(
        self, from_time: int = None, access_token: str = None
    ) -> ChatRoomsResponse:
        """

        チャットルームを更新します

        """
        return refresh_chat_rooms(self, from_time, access_token)

    def remove_chat_rooms(
        self, chat_room_ids: List[int], access_token: str = None
    ) -> dict:
        """

        チャットルームを削除します

        """
        return remove_chat_rooms(self, chat_room_ids, access_token)

    def report_chat_room(
        self,
        chat_room_id: int,
        opponent_id: int,
        category_id: int,
        reason: str = None,
        screenshot_filename: str = None,
        screenshot_2_filename: str = None,
        screenshot_3_filename: str = None,
        screenshot_4_filename: str = None,
        access_token: str = None,
    ) -> dict:
        """

        チャットルームを通報します

        """
        return report_chat_room(
            self,
            chat_room_id,
            opponent_id,
            category_id,
            reason,
            screenshot_filename,
            screenshot_2_filename,
            screenshot_3_filename,
            screenshot_4_filename,
            access_token,
        )

    def send_message(
        self,
        chat_room_id: int,
        message_type: str,
        call_type: str = None,
        text: str = None,
        font_size: int = None,
        gif_image_id: int = None,
        attachment_file_name: str = None,
        sticker_pack_id: int = None,
        sticker_id: int = None,
        video_file_name: str = None,
        access_token: str = None,
    ) -> MessageResponse:
        """

        チャットルームにメッセージを送信します

        """
        return send_message(
            self,
            chat_room_id,
            message_type,
            call_type,
            text,
            font_size,
            gif_image_id,
            attachment_file_name,
            sticker_pack_id,
            sticker_id,
            video_file_name,
            access_token,
        )

    def unhide_chat(self, chat_room_ids: int, access_token: str = None) -> dict:
        """

        チャットの非表示を解除します

        """
        return unhide_chat(self, chat_room_ids, access_token)

    def unpin_chat(self, chat_room_id: int, access_token: str = None) -> dict:
        """

        チャットルームのピン止めを解除します

        """
        return unpin_chat(self, chat_room_id, access_token)

    # -GROUP

    def accept_moderator_offer(self, group_id: int, access_token: str = None) -> dict:
        """

        サークル副管理人の権限オファーを引き受けます

        """
        return accept_moderator_offer(self, group_id, access_token)

    def accept_ownership_offer(self, group_id: int, access_token: str = None) -> dict:
        """

        サークル管理人の権限オファーを引き受けます

        """
        return accept_ownership_offer(self, group_id, access_token)

    def accept_group_join_request(
        self, group_id: int, user_id: int, access_token: str = None
    ) -> dict:
        """

        サークル参加リクエストを承認します

        """
        return accept_group_join_request(self, group_id, user_id, access_token)

    def add_related_groups(
        self, group_id: int, related_group_id: List[int], access_token: str = None
    ) -> dict:
        """

        関連サークルを追加します

        """
        return add_related_groups(self, group_id, related_group_id, access_token)

    def ban_group_user(
        self, group_id: int, user_id: int, access_token: str = None
    ) -> dict:
        """

        サークルからユーザーを追放します

        """
        return ban_group_user(self, group_id, user_id, access_token)

    def check_unread_status(
        self, from_time: int = None, access_token: str = None
    ) -> UnreadStatusResponse:
        """

        サークルの未読ステータスを取得します

        """
        return check_unread_status(self, from_time, access_token)

    def create_group(
        self,
        topic: str,
        description: str = None,
        secret: bool = None,
        hide_reported_posts: bool = None,
        hide_conference_call: bool = None,
        is_private: bool = None,
        only_verified_age: bool = None,
        only_mobile_verified: bool = None,
        call_timeline_display: bool = None,
        allow_ownership_transfer: bool = None,
        allow_thread_creation_by: str = None,
        gender: int = None,
        generation_groups_limit: int = None,
        group_category_id: int = None,
        cover_image_filename: str = None,
        sub_category_id: str = None,
        hide_from_game_eight: bool = None,
        allow_members_to_post_media: bool = None,
        allow_members_to_post_url: bool = None,
        guidelines: str = None,
        access_token: str = None,
    ) -> CreateGroupResponse:
        """

        サークルを作成します

        """
        return create_group(
            self,
            topic,
            description,
            secret,
            hide_reported_posts,
            hide_conference_call,
            is_private,
            only_verified_age,
            only_mobile_verified,
            call_timeline_display,
            allow_ownership_transfer,
            allow_thread_creation_by,
            gender,
            generation_groups_limit,
            group_category_id,
            cover_image_filename,
            sub_category_id,
            hide_from_game_eight,
            allow_members_to_post_media,
            allow_members_to_post_url,
            guidelines,
            access_token,
        )

    def pin_group(self, group_id: int, access_token: str = None) -> dict:
        """

        サークルをピンします

        """
        return create_pin_group(self, group_id, access_token)

    def decline_moderator_offer(self, group_id: int, access_token: str = None) -> dict:
        """

        サークル副管理人の権限オファーを断ります

        """
        return decline_moderator_offer(self, group_id, access_token)

    def decline_ownership_offer(self, group_id: int, access_token: str = None) -> dict:
        """

        サークル管理人の権限オファーを断ります

        """
        return decline_ownership_offer(self, group_id, access_token)

    def decline_group_join_request(
        self, group_id: int, user_id: int, access_token: str = None
    ) -> dict:
        """

        サークル参加リクエストを断ります

        """
        return decline_group_join_request(self, group_id, user_id, access_token)

    def unpin_group(self, group_id: int, access_token: str = None) -> dict:
        """

        サークルのピン止めを解除します

        """
        return delete_pin_group(self, group_id, access_token)

    def get_banned_group_members(
        self, group_id: int, page: int = None, access_token: str = None
    ) -> UsersResponse:
        """

        追放されたサークルメンバーを取得します

        """
        return get_banned_group_members(self, group_id, page, access_token)

    def get_group_categories(
        self, access_token: str = None, **params
    ) -> GroupCategoriesResponse:
        """

        サークルのカテゴリーを取得します

        Parameters
        ----------

            - page: int - (optional)
            - number: int - (optional)

        """
        return get_group_categories(self, access_token, **params)

    def get_create_group_quota(self, access_token: str = None) -> CreateGroupQuota:
        """

        サークル作成可能な割当量を取得します

        """
        return get_create_group_quota(self, access_token)

    def get_group(self, group_id: int, access_token: str = None) -> Group:
        """

        サークルの詳細を取得します

        """
        return get_group(self, group_id, access_token)

    def get_groups(self, access_token: str = None, **params) -> GroupsResponse:
        """

        複数のサークルの詳細を取得します

        Parameters
        ----------

            - group_category_id: int = None
            - keyword: str = None
            - from_timestamp: int = None
            - sub_category_id: int = None

        """
        return get_groups(self, access_token, **params)

    def get_invitable_users(
        self, group_id: int, access_token: str = None, **params
    ) -> UsersByTimestampResponse:
        """

        サークルに招待可能なユーザーを取得します

        Parameters
        ----------

            - from_timestamp: int - (optional)
            - user[nickname]: str - (optional)

        """
        return get_invitable_users(self, group_id, access_token, **params)

    def get_joined_statuses(self, ids: List[int], access_token: str = None) -> dict:
        """

        サークルの参加ステータスを取得します

        """
        return get_joined_statuses(self, ids, access_token)

    def get_group_member(
        self, group_id: int, user_id: int, access_token: str = None
    ) -> GroupUserResponse:
        """

        特定のサークルメンバーの情報を取得します

        """
        return get_group_member(self, group_id, user_id, access_token)

    def get_group_members(
        self, group_id: int, access_token: str = None, **params
    ) -> GroupUsersResponse:
        """

        サークルメンバーを取得します

        Parameters
        ----------

            - id: int - (required)
            - mode: str - (optional)
            - keyword: str - (optional)
            - from_id: int - (optional)
            - from_timestamp: int - (optional)
            - order_by: str - (optional)
            - followed_by_me: bool - (optional)

        """
        return get_group_members(self, group_id, access_token, **params)

    def get_my_groups(
        self, from_timestamp: None, access_token: str = None
    ) -> GroupsResponse:
        """

        自分のサークルを取得します

        """
        return get_my_groups(self, from_timestamp, access_token)

    def get_relatable_groups(
        self, group_id: int, access_token: str = None, **params
    ) -> GroupsRelatedResponse:
        """

        関連がある可能性があるサークルを取得します

        Parameters
        ----------

            - group_id: int - (required)
            - keyword: str - (optional)
            - from: str - (optional)

        """
        return get_relatable_groups(self, group_id, access_token, **params)

    def get_related_groups(
        self, group_id: int, access_token: str = None, **params
    ) -> GroupsRelatedResponse:
        """

        関連があるサークルを取得します

        Parameters
        ----------

            - group_id: int - (required)
            - keyword: str - (optional)
            - from: str - (optional)

        """
        return get_related_groups(self, group_id, access_token, **params)

    def get_user_groups(self, access_token: str = None, **params) -> GroupsResponse:
        """

        特定のユーザーが参加しているサークルを取得します

        Parameters
        ----------

            - user_id: int - (required)
            - page: int - (optional)

        """
        return get_user_groups(self, access_token, **params)

    def invite_users_to_group(
        self, group_id: int, user_ids: List[int], access_token: str = None
    ) -> dict:
        """

        サークルにユーザーを招待します

        """
        return invite_users_to_group(self, group_id, user_ids, access_token)

    def join_group(self, group_id: int, access_token: str = None) -> dict:
        """

        サークルに参加します

        """
        return join_group(self, group_id, access_token)

    def leave_group(self, group_id: int, access_token: str = None) -> dict:
        """

        サークルから脱退します

        """
        return leave_group(self, group_id, access_token)

    def post_gruop_social_shared(
        self, group_id: int, sns_name: str, access_token: str = None
    ) -> dict:
        """

        サークルのソーシャルシェアを投稿します

        """
        return post_gruop_social_shared(self, group_id, sns_name, access_token)

    def remove_group_cover(self, group_id: int, access_token: str = None) -> dict:
        """

        サークルのカバー画像を削除します

        """
        return remove_group_cover(self, group_id, access_token)

    def remove_moderator(
        self, group_id: int, user_id: int, access_token: str = None
    ) -> dict:
        """

        サークルの副管理人を削除します

        """
        return remove_moderator(self, group_id, user_id, access_token)

    def remove_related_groups(
        self, group_id: int, related_group_ids: List[int], access_token: str = None
    ) -> dict:
        """

        関連のあるサークルを削除します

        """
        return remove_related_groups(self, group_id, related_group_ids, access_token)

    def report_group(
        self,
        group_id: int,
        category_id: int,
        reason: str = None,
        opponent_id: int = None,
        screenshot_filename: str = None,
        screenshot_2_filename: str = None,
        screenshot_3_filename: str = None,
        screenshot_4_filename: str = None,
        access_token: str = None,
    ) -> dict:
        """

        サークルを通報します

        """
        return report_group(
            self,
            group_id,
            category_id,
            reason,
            opponent_id,
            screenshot_filename,
            screenshot_2_filename,
            screenshot_3_filename,
            screenshot_4_filename,
            access_token,
        )

    def send_moderator_offers(
        self, group_id: int, user_ids: List[int], access_token: str = None
    ) -> dict:
        """

        複数人にサークル副管理人のオファーを送信します

        """
        return send_moderator_offers(self, group_id, user_ids, access_token)

    def send_ownership_offer(
        self, group_id: int, user_id: int, access_token: str = None
    ) -> dict:
        """

        サークル管理人権限のオファーを送信します

        """
        return send_ownership_offer(self, group_id, user_id, access_token)

    def set_group_title(
        self, group_id: int, title: str, access_token: str = None
    ) -> dict:
        """

        サークルのタイトルを設定します

        """
        return set_group_title(self, group_id, title, access_token)

    def take_over_group_ownership(
        self, group_id: int, access_token: str = None
    ) -> dict:
        """

        サークル管理人の権限を引き継ぎます

        """
        return take_over_group_ownership(self, group_id, access_token)

    def unban_group_member(
        self, group_id: int, user_id: int, access_token: str = None
    ) -> dict:
        """

        特定のサークルメンバーの追放を解除します

        """
        return unban_group_member(self, group_id, user_id, access_token)

    def update_group(
        self,
        group_id: int,
        topic: str,
        description: str = None,
        secret: bool = None,
        hide_reported_posts: bool = None,
        hide_conference_call: bool = None,
        is_private: bool = None,
        only_verified_age: bool = None,
        only_mobile_verified: bool = None,
        call_timeline_display: bool = None,
        allow_ownership_transfer: bool = None,
        allow_thread_creation_by: str = None,
        gender: int = None,
        generation_groups_limit: int = None,
        group_category_id: int = None,
        cover_image_filename: str = None,
        sub_category_id: str = None,
        hide_from_game_eight: bool = None,
        allow_members_to_post_media: bool = None,
        allow_members_to_post_url: bool = None,
        guidelines: str = None,
        access_token: str = None,
    ) -> Group:
        """

        サークルを編集します

        """
        return update_group(
            group_id,
            topic,
            description,
            secret,
            hide_reported_posts,
            hide_conference_call,
            is_private,
            only_verified_age,
            only_mobile_verified,
            call_timeline_display,
            allow_ownership_transfer,
            allow_thread_creation_by,
            gender,
            generation_groups_limit,
            group_category_id,
            cover_image_filename,
            sub_category_id,
            hide_from_game_eight,
            allow_members_to_post_media,
            allow_members_to_post_url,
            guidelines,
            access_token,
        )

    def withdraw_moderator_offer(
        self, group_id: int, user_id: int, access_token: str = None
    ) -> dict:
        """

        サークル副管理人のオファーを取り消します

        """
        return withdraw_moderator_offer(self, group_id, user_id, access_token)

    def withdraw_ownership_offer(
        self, group_id: int, user_id: int, access_token: str = None
    ) -> dict:
        """

        サークル管理人のオファーを取り消します

        """
        return withdraw_ownership_offer(self, group_id, user_id, access_token)

    # -LOGIN

    def is_valid_token(self, access_token: str):
        """

        アクセストークンが有効か検証します

        """
        return is_valid_token(self, access_token)

    def change_password(
        self, current_password: str, new_password: str, access_token: str = None
    ) -> LoginUpdateResponse:
        """

        パスワードを変更します

        """
        return change_password(self, current_password, new_password, access_token)

    def get_token(
        self,
        grant_type: str,
        refresh_token: str = None,
        email: str = None,
        password: str = None,
        access_token: str = None,
    ) -> TokenResponse:
        """

        トークンを再発行します

        """
        return get_token(self, grant_type, refresh_token, email, password, access_token)

    def login(
        self, email: str, password: str, secret_key: str = None
    ) -> LoginUserResponse:
        """

        メールアドレスでログインします

        ※ ローカルストレージのトークンの暗号化を利用するには、`Client` クラスの `encrypt_cookie` 引数を`True` に設定してください。


        """
        return login_with_email(self, email, password, secret_key)

    def logout(self, access_token: str = None) -> dict:
        """

        ログアウトします

        """
        return logout(self, access_token)

    def resend_confirm_email(self, access_token: str = None) -> dict:
        """

        確認メールを再送信します

        """
        return resend_confirm_email(self, access_token)

    def restore_user(self, user_id: int, access_token: str = None) -> LoginUserResponse:
        """

        ユーザーを復元します

        """
        return restore_user(self, user_id, access_token)

    def revoke_tokens(self, access_token: str = None) -> dict:
        """

        トークンを無効化します

        """
        return revoke_tokens(self, access_token)

    def save_account_with_email(
        self,
        email: str,
        password: str = None,
        current_password: str = None,
        email_grant_token: str = None,
        access_token: str = None,
    ) -> LoginUpdateResponse:
        """

        メールアドレスでアカウントを保存します

        """
        return save_account_with_email(
            self,
            email,
            password,
            current_password,
            email_grant_token,
            access_token,
        )

    # -MISC

    def send_verification_code(self, email: str):
        """

        認証コードを送信します

        """
        return send_verification_code(self, email)

    def get_email_grant_token(self, code: int, email: str) -> str:
        """

        email_grant_tokenを取得します

        """
        return get_email_grant_token(self, code, email)

    def get_email_verification_presigned_url(
        self, email: str, locale: str, intent: str = None, access_token: str = None
    ) -> str:
        """

        メールアドレス確認用の署名付きURLを取得します

        """
        return get_email_verification_presigned_url(
            self, email, locale, intent, access_token
        )

    def get_file_upload_presigned_urls(
        self, file_names: List[str], access_token: str = None
    ) -> List[PresignedUrl]:
        """

        ファイルアップロード用の署名付きURLを取得します

        """
        return get_file_upload_presigned_urls(self, file_names, access_token)

    # def get_id_checker_presigned_url(
    #         self,
    #         model: str,
    #         action: str,
    #         **params
    # ) -> str:
    #     return get_id_checker_presigned_url(self, model, action, **params)

    def get_old_file_upload_presigned_url(
        self, video_file_name: str, access_token: str = None
    ) -> str:
        """

        動画ファイルアップロード用の署名付きURLを取得します

        """
        return get_old_file_upload_presigned_url(self, video_file_name, access_token)

    def get_web_socket_token(
        self, headers: dict = None, access_token: str = None
    ) -> str:
        """

        Web Socket Tokenを取得します

        """
        return get_web_socket_token(self, headers, access_token)

    def upload_image(self, image_paths: List[str], image_type: str) -> List[Attachment]:
        """

        画像をアップロードして、サーバー上のファイル名のリストを返します。

        Parameteres
        -----------

            - image_path: List[str] - (required): 画像パスのリスト
            - image_type: str - (required): 画像の種類

        Examples
        --------

        投稿に画像を付与する場合

        >>> # サーバー上にアップロード
        >>> attachments = api.upload_image(
        >>>     image_type=yaylib.IMAGE_TYPE_POST,
        >>>     image_paths=["./test.jpg"],
        >>> )
        >>> # サーバー上のファイル名を指定
        >>> api.create_post(
        >>>     "Hello with yaylib!",
        >>>     attachment_filename=attachments[0].filename
        >>> )

        """
        return upload_image(self, image_paths, image_type)

    def get_app_config(self) -> ApplicationConfig:
        """

        アプリケーションの設定情報を取得します

        """
        return get_app_config(self)

    def get_banned_words(self, country_code: str = "jp") -> List[BanWord]:
        """

        禁止ワードの一覧を取得します

        """
        return get_banned_words(self, country_code)

    def get_popular_words(self, country_code: str = "jp") -> List[PopularWord]:
        """

        人気のワードの一覧を取得します

        """
        return get_popular_words(self, country_code)

    # -POST

    def add_bookmark(
        self, user_id: int, post_id: int, access_token: str = None
    ) -> BookmarkPostResponse:
        """

        ブックマークに追加します

        """
        return add_bookmark(self, user_id, post_id, access_token)

    def add_group_highlight_post(
        self, group_id: int, post_id: int, access_token: str = None
    ) -> dict:
        """

        投稿をグループのまとめに追加します

        """
        return add_group_highlight_post(self, group_id, post_id, access_token)

    def create_call_post(
        self,
        text: str = None,
        font_size: int = None,
        color: int = None,
        group_id: int = None,
        call_type: str = None,
        category_id: int = None,
        game_title: str = None,
        joinable_by: str = None,
        message_tags: str = "[]",
        attachment_filename: str = None,
        attachment_2_filename: str = None,
        attachment_3_filename: str = None,
        attachment_4_filename: str = None,
        attachment_5_filename: str = None,
        attachment_6_filename: str = None,
        attachment_7_filename: str = None,
        attachment_8_filename: str = None,
        attachment_9_filename: str = None,
        access_token: str = None,
    ) -> ConferenceCall:
        """

        通話の投稿を作成します

        """
        return create_call_post(
            self,
            text,
            font_size,
            color,
            group_id,
            call_type,
            category_id,
            game_title,
            joinable_by,
            message_tags,
            attachment_filename,
            attachment_2_filename,
            attachment_3_filename,
            attachment_4_filename,
            attachment_5_filename,
            attachment_6_filename,
            attachment_7_filename,
            attachment_8_filename,
            attachment_9_filename,
            access_token,
        )

    def pin_group_post(
        self, post_id: int, group_id: int, access_token: str = None
    ) -> dict:
        """

        グループの投稿をピンします

        """
        return create_group_pin_post(self, post_id, group_id, access_token)

    def pin_post(self, post_id: int, access_token: str = None) -> dict:
        """

        投稿をピンします

        """
        return create_pin_post(self, post_id, access_token)

    def mention(self, user_id: int) -> str:
        """

        メンション用文字列を返します

        """
        return mention(self, user_id)

    def create_post(
        self,
        text: str = None,
        font_size: int = 0,
        color: int = 0,
        in_reply_to: int = None,
        group_id: int = None,
        mention_ids: List[int] = None,
        choices: List[str] = None,
        shared_url: str = None,
        message_tags: str = "[]",
        attachment_filename: str = None,
        attachment_2_filename: str = None,
        attachment_3_filename: str = None,
        attachment_4_filename: str = None,
        attachment_5_filename: str = None,
        attachment_6_filename: str = None,
        attachment_7_filename: str = None,
        attachment_8_filename: str = None,
        attachment_9_filename: str = None,
        video_file_name: str = None,
        access_token: str = None,
    ) -> Post:
        """

        投稿を作成します

        """
        return create_post(
            self,
            text,
            font_size,
            color,
            in_reply_to,
            group_id,
            mention_ids,
            choices,
            shared_url,
            message_tags,
            attachment_filename,
            attachment_2_filename,
            attachment_3_filename,
            attachment_4_filename,
            attachment_5_filename,
            attachment_6_filename,
            attachment_7_filename,
            attachment_8_filename,
            attachment_9_filename,
            video_file_name,
            access_token,
        )

    def create_repost(
        self,
        post_id: int,
        text: str = None,
        font_size: int = None,
        color: int = None,
        in_reply_to: int = None,
        group_id: int = None,
        mention_ids: List[int] = None,
        choices: List[str] = None,
        shared_url: Dict[str, Union[str, int]] = None,
        message_tags: str = "[]",
        attachment_filename: str = None,
        attachment_2_filename: str = None,
        attachment_3_filename: str = None,
        attachment_4_filename: str = None,
        attachment_5_filename: str = None,
        attachment_6_filename: str = None,
        attachment_7_filename: str = None,
        attachment_8_filename: str = None,
        attachment_9_filename: str = None,
        video_file_name: str = None,
        access_token: str = None,
    ) -> Post:
        """

        投稿を(´∀｀∩)↑age↑します

        """
        return create_repost(
            self,
            post_id,
            text,
            font_size,
            color,
            in_reply_to,
            group_id,
            mention_ids,
            choices,
            shared_url,
            message_tags,
            attachment_filename,
            attachment_2_filename,
            attachment_3_filename,
            attachment_4_filename,
            attachment_5_filename,
            attachment_6_filename,
            attachment_7_filename,
            attachment_8_filename,
            attachment_9_filename,
            video_file_name,
            access_token,
        )

    def create_share_post(
        self,
        shareable_type: str,
        shareable_id: int,
        text: str = None,
        font_size: int = None,
        color: int = None,
        group_id: int = None,
        access_token: str = None,
    ) -> Post:
        """

        シェア投稿を作成します

        """
        return create_share_post(
            self,
            shareable_type,
            shareable_id,
            text,
            font_size,
            color,
            group_id,
            access_token,
        )

    def create_thread_post(
        self,
        post_id: int,
        text: str = None,
        font_size: int = None,
        color: int = None,
        in_reply_to: int = None,
        group_id: int = None,
        mention_ids: List[int] = None,
        choices: List[str] = None,
        shared_url: Dict[str, Union[str, int]] = None,
        message_tags: str = "[]",
        attachment_filename: str = None,
        attachment_2_filename: str = None,
        attachment_3_filename: str = None,
        attachment_4_filename: str = None,
        attachment_5_filename: str = None,
        attachment_6_filename: str = None,
        attachment_7_filename: str = None,
        attachment_8_filename: str = None,
        attachment_9_filename: str = None,
        video_file_name: str = None,
        access_token: str = None,
    ) -> Post:
        """

        スレッドの投稿を作成します

        """
        return create_thread_post(
            self,
            post_id,
            text,
            font_size,
            color,
            in_reply_to,
            group_id,
            mention_ids,
            choices,
            shared_url,
            message_tags,
            attachment_filename,
            attachment_2_filename,
            attachment_3_filename,
            attachment_4_filename,
            attachment_5_filename,
            attachment_6_filename,
            attachment_7_filename,
            attachment_8_filename,
            attachment_9_filename,
            video_file_name,
            access_token,
        )

    def delete_all_post(self, access_token: str = None) -> dict:
        """

        すべての自分の投稿を削除します

        """
        return delete_all_post(self, access_token)

    def unpin_group_post(self, group_id: int, access_token: str = None) -> dict:
        """

        グループのピン投稿を解除します

        """
        return delete_group_pin_post(self, group_id, access_token)

    def unpin_post(self, post_id: int, access_token: str = None) -> dict:
        """

        ピン投稿を削除します

        """
        return delete_pin_post(self, post_id, access_token)

    def get_bookmark(
        self, user_id: int, from_str: str = None, access_token: str = None
    ) -> PostsResponse:
        """

        ブックマークを取得します

        """
        return get_bookmark(self, user_id, from_str, access_token)

    def get_timeline_calls(self, access_token: str = None, **params) -> PostsResponse:
        """

        誰でも通話を取得します

        Parameters
        ----------

            - group_id: int = None
            - from_timestamp: int = None
            - number: int = None
            - category_id: int = None
            - call_type: str = "voice"
            - include_circle_call: bool = None
            - cross_generation: bool = None
            - exclude_recent_gomimushi: bool = None
            - shared_interest_categories: bool = None

        """
        return get_timeline_calls(self, access_token, **params)

    def get_conversation(
        self, conversation_id: int, access_token: str = None, **params
    ) -> PostsResponse:
        """

        リプライを含める投稿の会話を取得します

        Parameters
        ----------

            - conversation_id: int
            - group_id: int = None
            - thread_id: int = None
            - from_post_id: int = None
            - number: int = 50
            - reverse: bool = True

        """
        return get_conversation(self, conversation_id, access_token, **params)

    def get_conversation_root_posts(
        self, post_ids: List[int], access_token: str = None
    ) -> PostsResponse:
        """

        会話の原点の投稿を取得します

        """
        return get_conversation_root_posts(self, post_ids, access_token)

    def get_following_call_timeline(
        self, access_token: str = None, **params
    ) -> PostsResponse:
        """

        フォロー中の通話を取得します

        Parameters
        ----------

            - from_timestamp: int = None
            - number: int = None
            - category_id: int = None
            - call_type: str = None
            - include_circle_call: bool = None
            - exclude_recent_gomimushi: bool = None

        """
        return get_following_call_timeline(self, access_token, **params)

    def get_following_timeline(
        self, access_token: str = None, **params
    ) -> PostsResponse:
        """

        フォロー中のタイムラインを取得します

        Parameters
        ----------

            - from_str: str = None
            - only_root: bool = None
            - order_by: str = None
            - number: int = None
            - mxn: int = None
            - reduce_selfie: bool = None
            - custom_generation_range: bool = None

        """
        return get_following_timeline(self, access_token, **params)

    def get_group_highlight_posts(
        self, group_id: int, access_token: str = None, **params
    ) -> PostsResponse:
        """

        グループのハイライト投稿を取得します

        Parameters
        ----------

            - group_id: int
            - from_post: int = None
            - number: int = None

        """
        return get_group_highlight_posts(self, group_id, access_token, **params)

    def get_group_timeline_by_keyword(
        self, group_id: int, keyword: str, access_token: str = None, **params
    ) -> PostsResponse:
        """

        グループの投稿をキーワードで検索します

        Parameters
        ----------

            - group_id: int
            - keyword: str
            - from_post_id: int = None
            - number: int = None
            - only_thread_posts: bool = False

        """
        return get_group_timeline_by_keyword(
            self, group_id, keyword, access_token, **params
        )

    def get_group_timeline(
        self, group_id: int, access_token: str = None, **params
    ) -> PostsResponse:
        """

        グループのタイムラインを取得します

        Parameters
        ----------

            - group_id: int
            - from_post_id: int
            - reverse: bool
            - post_type: str
            - number: int
            - only_root: bool

        """
        return get_group_timeline(self, group_id, access_token, **params)

    def get_timeline_by_hashtag(
        self, hashtag: str, access_token: str = None, **params
    ) -> PostsResponse:
        """

        ハッシュタグでタイムラインを検索します

        Parameters
        ----------

            - hashtag: str - (required)
            - from_post_id: int - (optional)
            - number: int - (optional)

        """
        return get_timeline_by_hashtag(self, hashtag, access_token, **params)

    def get_my_posts(self, access_token: str = None, **params) -> PostsResponse:
        """

        自分の投稿を取得します

        Parameters
        ----------

            - from_post_id: int - (optional)
            - number: int - (optional)
            - include_group_post: bool - (optional)

        """
        return get_my_posts(self, access_token, **params)

    def get_post(self, post_id: int, access_token: str = None) -> Post:
        """

        投稿の詳細を取得します

        """
        return get_post(self, post_id, access_token)

    def get_post_likers(
        self, post_id: int, access_token: str = None, **params
    ) -> PostLikersResponse:
        """

        投稿にいいねしたユーザーを取得します

        Parameters
        ----------

            - from_id: int - (optional)
            - number: int - (optional)

        """
        return get_post_likers(self, post_id, access_token, **params)

    def get_reposts(
        self, post_id: int, access_token: str = None, **params: int
    ) -> PostsResponse:
        """

        投稿の(´∀｀∩)↑age↑を取得します

        Parameters
        ----------

            - post_id: int - (required)
            - from_post_id: int - (optional)
            - number: int - (optional)

        """
        return get_post_reposts(self, post_id, access_token, **params)

    def get_posts(self, post_ids: List[int], access_token: str = None) -> PostsResponse:
        """

        複数の投稿を取得します

        """
        return get_posts(self, post_ids, access_token)

    def get_recommended_post_tags(
        self,
        tag: str = None,
        save_recent_search: bool = False,
        access_token: str = None,
    ) -> PostTagsResponse:
        """

        おすすめのタグ候補を取得します

        """
        return get_recommended_post_tags(self, tag, save_recent_search, access_token)

    def get_recommended_posts(
        self, access_token: str = None, **params
    ) -> PostsResponse:
        """

        おすすめの投稿を取得します

        Parameters
        ----------

            - experiment_num: int
            - variant_num: int
            - number: int

        """
        return get_recommended_posts(self, access_token, **params)

    def get_timeline_by_keyword(
        self, keyword: str = None, access_token: str = None, **params
    ) -> PostsResponse:
        """

        キーワードでタイムラインを検索します

        Parameters
        ----------

            - keyword: str
            - from_post_id: int
            - number: int

        """
        return get_timeline_by_keyword(self, keyword, access_token, **params)

    def get_timeline(self, access_token: str = None, **params) -> PostsResponse:
        # - from: str - (optional)
        """

        タイムラインを取得します

        Parameters
        ----------

            - noreply_mode: bool - (optional)
            - from_post_id: int - (optional)
            - number: int - (optional)
            - order_by: str - (optional)
            - experiment_older_age_rules: bool - (optional)
            - shared_interest_categories: bool - (optional)
            - mxn: int - (optional)
            - en: int - (optional)
            - vn: int - (optional)
            - reduce_selfie: bool - (optional)
            - custom_generation_range: bool - (optional)

        """
        return get_timeline(self, access_token, **params)

    def get_url_metadata(self, url: str, access_token: str = None) -> SharedUrl:
        """

        URLのメタデータを取得します

        """
        return get_url_metadata(self, url, access_token)

    def get_user_timeline(
        self, user_id: int, access_token: str = None, **params
    ) -> PostsResponse:
        """

        ユーザーのタイムラインを取得します

        Parameters
        ----------

            - from_post_id: int - (optional)
            - number: int - (optional)
            - post_type: str - (optional)

        """
        return get_user_timeline(self, user_id, access_token, **params)

    def like(self, post_ids: List[int], access_token: str = None) -> LikePostsResponse:
        """

        投稿にいいねします

        ※ 一度にいいねできる投稿数は最大25個

        """
        return like_posts(self, post_ids, access_token)

    def remove_bookmark(
        self, user_id: int, post_id: int, access_token: str = None
    ) -> dict:
        """

        ブックマークを削除します

        """
        return remove_bookmark(self, user_id, post_id, access_token)

    def remove_group_highlight_post(
        self, group_id: int, post_id: int, access_token: str = None
    ) -> dict:
        """

        サークルのハイライト投稿を解除します

        """
        return remove_group_highlight_post(self, group_id, post_id, access_token)

    def remove_posts(self, post_ids: List[int], access_token: str = None) -> dict:
        """

        複数の投稿を削除します

        """
        return remove_posts(self, post_ids, access_token)

    def report_post(
        self,
        post_id: int,
        opponent_id: int,
        category_id: int,
        reason: str = None,
        screenshot_filename: str = None,
        screenshot_2_filename: str = None,
        screenshot_3_filename: str = None,
        screenshot_4_filename: str = None,
        access_token: str = None,
    ) -> dict:
        """

        投稿を通報します

        """
        return report_post(
            self,
            post_id,
            opponent_id,
            category_id,
            reason,
            screenshot_filename,
            screenshot_2_filename,
            screenshot_3_filename,
            screenshot_4_filename,
            access_token,
        )

    def unlike(self, post_id: int, access_token: str = None) -> dict:
        """

        投稿のいいねを解除します

        """
        return unlike_post(self, post_id, access_token)

    def update_post(
        self,
        post_id: int,
        text: str = None,
        font_size: int = None,
        color: int = None,
        message_tags: str = "[]",
        access_token: str = None,
    ) -> Post:
        """

        投稿を編集します

        """
        return update_post(
            self, post_id, text, font_size, color, message_tags, access_token
        )

    def view_video(self, video_id: int, access_token: str = None) -> dict:
        """

        動画を視聴します

        """
        return view_video(self, video_id, access_token)

    def vote_survey(
        self, survey_id: int, choice_id: int, access_token: str = None
    ) -> Survey:
        """

        アンケートに投票します

        """
        return vote_survey(self, survey_id, choice_id, access_token)

    # -REVIEW

    def create_review(
        self, user_id: int, comment: str, access_token: str = None
    ) -> dict:
        """

        レターを送信します

        """
        return create_review(self, user_id, comment, access_token)

    def create_reviews(
        self, user_ids: List[int], comment: str, access_token: str = None
    ) -> dict:
        """

        複数人にレターを送信します

        """
        return create_reviews(self, user_ids, comment, access_token)

    def delete_reviews(self, review_ids: List[int], access_token: str = None) -> dict:
        """

        レターを削除します

        """
        return delete_reviews(self, review_ids, access_token)

    def get_my_reviews(self, access_token: str = None, **params) -> ReviewsResponse:
        """

        送信したレターを取得します

        Parameters
        ----------

            - from_id: int (optional)
            - number: int = (optional)

        """
        return get_my_reviews(self, access_token, **params)

    def get_reviews(
        self, user_id: int, access_token: str = None, **params
    ) -> ReviewsResponse:
        """

        ユーザーが受け取ったレターを取得します

        Parameters
        ----------

            - user_id: int (required)
            - from_id: int = (optional)
            - number: int = (optional)

        """
        return get_reviews(self, user_id, access_token, **params)

    def pin_review(self, review_id: int, access_token: str = None) -> dict:
        """

        レターをピンします

        """
        return pin_review(self, review_id, access_token)

    def unpin_review(self, review_id: int, access_token: str = None) -> dict:
        """

        レターのピン止めを解除します

        """
        return unpin_review(self, review_id, access_token)

    # -THREAD

    def add_post_to_thread(
        self, post_id: int, thread_id: int, access_token: str = None
    ) -> ThreadInfo:
        """

        投稿をスレッドに追加します

        """
        return add_post_to_thread(self, post_id, thread_id, access_token)

    def convert_post_to_thread(
        self,
        post_id: int,
        title: str = None,
        thread_icon_filename: str = None,
        access_token: str = None,
    ) -> ThreadInfo:
        """

        投稿をスレッドに変換します

        """
        return convert_post_to_thread(
            self, post_id, title, thread_icon_filename, access_token
        )

    def create_thread(
        self,
        group_id: int,
        title: str,
        thread_icon_filename: str,
        access_token: str = None,
    ) -> ThreadInfo:
        """

        スレッドを作成します

        """
        return create_thread(self, group_id, title, thread_icon_filename, access_token)

    def get_group_thread_list(
        self, group_id: int, from_str: str = None, access_token: str = None, **params
    ) -> GroupThreadListResponse:
        """

        グループのスレッド一覧を取得します

        Parameters
        ----------

            - group_id: int
            - from_str: str = None
            - join_status: str = None

        """
        return get_group_thread_list(self, group_id, from_str, access_token, **params)

    def get_thread_joined_statuses(
        self, ids: List[int], access_token: str = None
    ) -> dict:
        """

        スレッド参加ステータスを取得します

        """
        return get_thread_joined_statuses(self, ids, access_token)

    def get_thread_posts(
        self, thread_id: int, from_str: str = None, access_token: str = None, **params
    ) -> PostsResponse:
        """

        スレッドの投稿を取得します

        Parameters
        ----------

            - post_type: str
            - number: int = None
            - from_str: str = None

        """
        return get_thread_posts(self, thread_id, from_str, access_token, **params)

    def join_thread(
        self, thread_id: int, user_id: int, access_token: str = None
    ) -> dict:
        """

        スレッドに参加します

        """
        return join_thread(self, thread_id, user_id, access_token)

    def leave_thread(
        self, thread_id: int, user_id: int, access_token: str = None
    ) -> dict:
        """

        スレッドから脱退します

        """
        return leave_thread(self, thread_id, user_id, access_token)

    def remove_thread(self, thread_id: int, access_token: str = None) -> dict:
        """

        スレッドを削除します

        """
        return remove_thread(self, thread_id, access_token)

    def update_thread(
        self,
        thread_id: int,
        title: str,
        thread_icon_filename: str,
        access_token: str = None,
    ) -> dict:
        """

        スレッドを編集します

        """
        return update_thread(self, thread_id, title, thread_icon_filename, access_token)

    # -USER

    def delete_footprint(
        self, user_id: int, footprint_id: int, access_token: str = None
    ) -> dict:
        """

        足跡を削除します

        """
        return delete_footprint(self, user_id, footprint_id, access_token)

    def destroy_user(self, access_token: str = None) -> dict:
        """

        アカウントを削除します

        """
        return destroy_user(self, access_token)

    def follow_user(self, user_id: int, access_token: str = None) -> dict:
        """

        ユーザーをフォローします

        """
        return follow_user(self, user_id, access_token)

    def follow_users(self, user_ids: List[int], access_token: str = None) -> dict:
        """

        複数のユーザーをフォローします

        """
        return follow_users(self, user_ids, access_token)

    def get_active_followings(
        self, access_token: str = None, **params
    ) -> ActiveFollowingsResponse:
        """

        アクティブなフォロー中のユーザーを取得します

        Parameters
        ----------

            - only_online: bool
            - from_loggedin_at: int = None

        """
        return get_active_followings(self, access_token, **params)

    def get_follow_recommendations(
        self, access_token: str = None, **params
    ) -> FollowRecommendationsResponse:
        """

        フォローするのにおすすめのユーザーを取得します

        Parameters
        ----------

            - from_timestamp: int = None,
            - number: int = None,
            - sources: List[str] = None

        """
        return get_follow_recommendations(self, access_token, **params)

    def get_follow_request(
        self, from_timestamp: int = None, access_token: str = None
    ) -> UsersByTimestampResponse:
        """

        フォローリクエストを取得します

        """
        return get_follow_request(self, from_timestamp, access_token)

    def get_follow_request_count(self, access_token: str = None) -> int:
        """

        フォローリクエストの数を取得します

        """
        return get_follow_request_count(self, access_token)

    def get_following_users_born(
        self, birthdate: int = None, access_token: str = None
    ) -> UsersResponse:
        """

        フォロー中のユーザーの誕生日を取得します

        """
        return get_following_users_born(self, birthdate, access_token)

    def get_footprints(self, access_token: str = None, **params) -> List[Footprint]:
        """

        足跡を取得します

        Parameters
        ----------

            - from_id: int = None
            - number: int = None
            - mode: str = None

        """
        return get_footprints(self, access_token, **params)

    def get_fresh_user(self, user_id: int, access_token: str = None) -> UserResponse:
        """

        認証情報などを含んだユーザー情報を取得します

        """
        return get_fresh_user(self, user_id, access_token)

    def get_hima_users(self, access_token: str = None, **params) -> List[UserWrapper]:
        """

        暇なユーザーを取得します

        Parameters
        ----------

            - from_hima_id: int = None
            - number: int = None

        """
        return get_hima_users(self, access_token, **params)

    def get_user_ranking(self, mode: str) -> RankingUsersResponse:
        """

        ユーザーのフォロワーランキングを取得します

        Examples
        --------

        >>> ルーキーを取得する場合:

        >>> api.get_user_ranking(mode="one_month")

        ---

        >>> ミドルを取得する場合:

        >>> api.get_user_ranking(mode="six_months")

        ---

        >>> 殿堂入りを取得する場合:

        >>> api.get_user_ranking(mode="all_time")

        """
        return get_user_ranking(self, mode)

    def get_refresh_counter_requests(
        self, access_token: str = None
    ) -> RefreshCounterRequestsResponse:
        """

        カウンター更新のリクエストを取得します

        """
        return get_refresh_counter_requests(self, access_token)

    def get_social_shared_users(
        self, access_token: str = None, **params
    ) -> SocialShareUsersResponse:
        """

        SNS共有をしたユーザーを取得します

        Parameters
        ----------

            - sns_name: str - (Required)
            - number: int - (Optional)
            - from_id: int - (Optional)

        """
        return get_social_shared_users(self, access_token, **params)

    def get_timestamp(self, access_token: str = None) -> UserTimestampResponse:
        """

        タイムスタンプを取得します

        """
        return get_timestamp(self, access_token)

    def get_user(self, user_id: int, access_token: str = None) -> User:
        """

        ユーザーの情報を取得します

        """
        return get_user(self, user_id, access_token)

    def get_user_email(self, user_id: int, access_token: str = None) -> str:
        """

        ユーザーのメールアドレスを取得します

        """
        return get_user_email(self, user_id, access_token)

    def get_user_followers(
        self, user_id: int, access_token: str = None, **params
    ) -> FollowUsersResponse:
        """

        ユーザーのフォロワーを取得します

        Parameters
        ----------

            - user_id: int
            - from_follow_id: int = None
            - followed_by_me: int = None

        """
        return get_user_followers(self, user_id, access_token, **params)

    def get_user_followings(
        self, user_id: int, access_token: str = None, **params
    ) -> FollowUsersResponse:
        """

        フォロー中のユーザーを取得します

        Parameters
        ----------

            - user_id: int
            - from_follow_id: int = None
            - from_timestamp: int = None
            - order_by: str = None

        """
        return get_user_followings(self, user_id, access_token, **params)

    def get_user_from_qr(self, qr: str, access_token: str = None) -> UserResponse:
        """

        QRコードからユーザーを取得します

        """
        return get_user_from_qr(self, qr, access_token)

    def get_user_without_leaving_footprint(
        self, user_id: int, access_token: str = None
    ) -> UserResponse:
        """

        足跡をつけずにユーザーの情報を取得します

        """
        return get_user_without_leaving_footprint(self, user_id, access_token)

    def get_users(self, user_ids: List[int], access_token: str = None) -> UsersResponse:
        """

        複数のユーザーの情報を取得します

        """
        return get_users(self, user_ids, access_token)

    def refresh_counter(self, counter: str, access_token: str = None) -> dict:
        """

        カウンターを更新します

        """
        return refresh_counter(self, counter, access_token)

    def register(
        self,
        email: str,
        email_grant_token: str,
        password: str,
        nickname: str,
        birth_date: str,
        gender: int = -1,
        country_code: str = "JP",
        biography: str = None,
        prefecture: str = None,
        profile_icon_filename: str = None,
        cover_image_filename: str = None,
        # @Nullable @Part("sns_info") SignUpSnsInfoRequest signUpSnsInfoRequest,
        en: int = None,
        vn: int = None,
    ):
        """

        Register user

        """
        return register(
            self,
            email,
            email_grant_token,
            password,
            nickname,
            birth_date,
            gender,
            country_code,
            biography,
            prefecture,
            profile_icon_filename,
            cover_image_filename,
            en,
            vn,
        )

    def remove_user_avatar(self, access_token: str = None) -> dict:
        """

        ユーザーのアイコンを削除します

        """
        return remove_user_avatar(self, access_token)

    def remove_user_cover(self, access_token: str = None) -> dict:
        """

        ユーザーのカバー画像を削除します

        """
        return remove_user_cover(self, access_token)

    def report_user(
        self,
        user_id: int,
        category_id: int,
        reason: str = None,
        screenshot_filename: str = None,
        screenshot_2_filename: str = None,
        screenshot_3_filename: str = None,
        screenshot_4_filename: str = None,
        access_token: str = None,
    ) -> dict:
        """

        ユーザーを通報します

        """
        return report_user(
            self,
            user_id,
            category_id,
            reason,
            screenshot_filename,
            screenshot_2_filename,
            screenshot_3_filename,
            screenshot_4_filename,
            access_token,
        )

    def reset_password(
        self,
        email: str,
        email_grant_token: str,
        password: str,
        access_token: str = None,
    ) -> dict:
        """

        パスワードをリセットします

        """
        return reset_password(self, email, email_grant_token, password, access_token)

    def search_lobi_users(self, access_token: str = None, **params) -> UsersResponse:
        """

        Lobiのユーザーを検索します

        Parameters
        ----------

            - nickname: str = None
            - number: int = None
            - from_str: str = None

        """
        return search_lobi_users(self, access_token, **params)

    def search_users(self, access_token: str = None, **params) -> UsersResponse:
        """

        ユーザーを検索します

        Parameters
        ----------

            - gender: int = None
            - nickname: str = None
            - title: str = None
            - biography: str = None
            - from_timestamp: int = None
            - similar_age: bool = None
            - not_recent_gomimushi: bool = None
            - recently_created: bool = None
            - same_prefecture: bool = None
            - save_recent_search: bool = None

        """
        return search_users(self, access_token, **params)

    def set_follow_permission_enabled(
        self, nickname: str, is_private: bool = None, access_token: str = None
    ) -> dict:
        """

        フォローを許可制に設定します

        """
        return set_follow_permission_enabled(self, nickname, is_private, access_token)

    def take_action_follow_request(
        self, target_id: int, action: str, access_token: str = None
    ) -> dict:
        return take_action_follow_request(self, target_id, action, access_token)

    def turn_on_hima(self, access_token: str = None) -> dict:
        """

        ひまなう

        """
        return turn_on_hima(self, access_token)

    def unfollow_user(self, user_id: int, access_token: str = None) -> dict:
        """

        ユーザーをアンフォローします

        """
        return unfollow_user(self, user_id, access_token)

    def update_user(
        self,
        nickname: str,
        biography: str = None,
        prefecture: str = None,
        gender: int = None,
        country_code: str = None,
        profile_icon_filename: str = None,
        cover_image_filename: str = None,
        username: str = None,
        access_token: str = None,
    ) -> dict:
        """

        プロフィールを更新します

        """
        return update_user(
            self,
            nickname,
            biography,
            prefecture,
            gender,
            country_code,
            profile_icon_filename,
            cover_image_filename,
            username,
            access_token,
        )

    def block_user(self, user_id: int, access_token: str = None) -> dict:
        """

        ユーザーをブロックします

        """
        return block_user(self, user_id, access_token)

    def get_blocked_user_ids(self, access_token: str = None) -> BlockedUserIdsResponse:
        """

        あなたをブロックしたユーザーを取得します

        """
        return get_blocked_user_ids(self, access_token)

    def get_blocked_users(
        self, from_id: int = None, access_token: str = None
    ) -> BlockedUsersResponse:
        """

        ブロックしたユーザーを取得します

        """
        return get_blocked_users(self, from_id, access_token)

    def unblock_user(self, user_id: int, access_token: str = None) -> dict:
        """

        ユーザーをアンブロックします

        """
        return unblock_user(self, user_id, access_token)

    def get_hidden_users_list(
        self, access_token: str = None, **params: Union[str, int]
    ) -> HiddenResponse:
        """

        非表示のユーザー一覧を取得します

        Parameters
        ----------

            - from: str = None
            - number: int = None

        """
        return get_hidden_users_list(self, access_token, **params)

    def hide_user(self, user_id: int, access_token: str = None) -> dict:
        """

        ユーザーを非表示にします

        """
        return hide_user(self, user_id, access_token)

    def unhide_users(self, user_ids: List[int], access_token: str = None) -> dict:
        """

        ユーザーの非表示を解除します

        """
        return unhide_users(self, user_ids, access_token)
