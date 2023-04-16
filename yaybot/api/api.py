import datetime
import logging
import os
import random
import requests
import time

from ..config import Endpoints as ep
from ..exceptions import (
    YayError,
    AuthenticationError,
    ForbiddenError,
    RateLimitError,
    ExceedCallQuotaError,
    InvalidSignedInfo,
    UnknownError
)
from ..support import console_print
from .api_auth import YayAuth
from .api_chat import (
    send_message,
    accept_chat_request,
    delete_chat_room
)
from .api_get import (
    get_user,
    get_users_from_dict,
    get_letters_from_dict,
    get_letters,
    get_joined_groups,
    get_user_followers,
    get_user_followings,
    get_follow_requests,
    get_user_active_call,
    get_blocked_users,
    get_post,
    get_posts_from_dict,
    get_timeline,
    get_following_timeline,
    get_conversation,
    get_reposts,
    get_post_likers,
    get_group,
    get_groups_from_dict,
    get_group_users_from_dict,
    get_group_timeline,
    get_group_call,
    get_group_members,
    get_pending_users_in_group,
    get_banned_user_from_group,
    get_chat_room,
    get_chat_rooms_from_dict,
    get_chat_messages_from_dict,
    get_chat_room_id_from_user,
    get_chat_messages,
    get_chat_rooms,
    get_chat_requests,
    get_activities_from_dict,
    get_notification,
)
from .api_group import (
    create_group,
    delete_group,
    change_group_settings,
    transfer_group_ownership,
    offer_group_sub_owner,
    undo_group_ownership_transfer,
    undo_group_sub_owner_offer,
    fire_group_sub_owner,
    accept_group_join_request,
    decline_group_join_request,
    invite_user_to_group,
    pin_group_post,
    unpin_group_post,
    ban_user_from_group,
    unban_user_from_group,
    join_group,
    leave_group,
)
from .api_post import (
    create_post,
    create_post_in_group,
    create_repost,
    create_reply,
    delete_post,
    pin_post,
    unpin_post,
    like_post,
    unlike_post,
)
from .api_user import (
    follow_user,
    unfollow_user,
    accept_follow_request,
    reject_follow_request,
    send_letter,
    block_user,
    unblock_user,
)

version = '0.2.5'  # also change .. __init__
current_path = os.path.abspath(os.getcwd())


class Yay(object):

    def __init__(
            self,
            token: str = None,
            proxy: str = None,
            timeout=10,
            base_path=current_path + '/config/',
            save_logfile=False,
            log_filename: str = None,
            loglevel_file=logging.DEBUG,
            loglevel_stream=logging.INFO,
    ):
        """

        YayBot
        ---
            Unofficial API for Yay! (yay.space) - developed by qualia-5w4

        """
        self.base_path = base_path

        # Setup logging
        self.logger = logging.getLogger('YayBot version: ' + version)

        if save_logfile is True:
            if not os.path.exists(base_path):
                # create base_path if not exists
                os.makedirs(base_path)

            if not os.path.exists(base_path + '/log/'):
                # create log folder if not exists
                os.makedirs(base_path + '/log/')

            if log_filename is None:
                log_filename = os.path.join(
                    base_path, 'log/yaybot_{}.log'.format(id(self))
                )

            fh = logging.FileHandler(filename=log_filename)
            fh.setLevel(loglevel_file)
            fh.setFormatter(
                logging.Formatter(
                    '%(asctime)s - %(name)s (%(module)s %(pathname)s:%(lineno)s) - %(levelname)s - %(message)s'
                )
            )

            handler_existed = False
            for handler in self.logger.handlers:
                if isinstance(handler, logging.FileHandler):
                    handler_existed = True
                    break
            if not handler_existed:
                self.logger.addHandler(fh)

        ch = logging.StreamHandler()
        ch.setLevel(loglevel_stream)
        ch.setFormatter(logging.Formatter(
            '%(asctime)s - %(levelname)s - %(message)s'))

        handler_existed = False
        for handler in self.logger.handlers:
            if isinstance(handler, logging.StreamHandler):
                handler_existed = True
                break
        if not handler_existed:
            self.logger.addHandler(ch)
        self.logger.setLevel(logging.DEBUG)

        self.auth = YayAuth(proxy=proxy, timeout=timeout)

        if token:
            self.access_token = token
            self.auth.access_token = token
            self.auth.headers.setdefault('Authorization', f'Bearer {token}')

        self.logger.info('YayBot version: ' + version + ' Started!')

    def login(self, email, password):
        self.auth.login(email, password)
        self.set_login_status()

    def logout(self):
        self.auth.logout()
        self.pop_login_status()

    def set_login_status(self):
        self.access_token = self.auth.access_token
        self.api_key = self.auth.api_key
        self.logged_in_as = self.auth.logged_in_as

    def pop_login_status(self):
        self.access_token = None
        self.api_key = None
        self.logged_in_as = None

    def _get(self, url: str):
        resp = requests.get(url, headers=self.auth.headers,
                            proxies=self.auth.proxies, timeout=self.auth.timeout)
        self._handle_response(resp)
        return resp.json()

    def _post(self, url: str, data: dict = None):
        resp = requests.post(url, params=data,
                             headers=self.auth.headers,
                             proxies=self.auth.proxies,
                             timeout=self.auth.timeout)
        self._handle_response(resp)
        return resp.json()

    def _put(self, url: str, data: dict = None):
        resp = requests.put(url, params=data,
                            headers=self.auth.headers,
                            proxies=self.auth.proxies,
                            timeout=self.auth.timeout)
        self._handle_response(resp)
        return resp.json()

    def _delete(self, url: str, data: dict = None):
        resp = requests.delete(url, params=data,
                               headers=self.auth.headers,
                               proxies=self.auth.proxies,
                               timeout=self.auth.timeout)
        self._handle_response(resp)
        return resp.json()

    def _handle_response(self, resp):
        if resp.status_code == 401:
            raise AuthenticationError('Failed to authenticate')
        if resp.status_code == 403:
            raise ForbiddenError('Forbidden')
        if resp.status_code == 429:
            raise RateLimitError('Rate limit exceeded')

        resp_json = resp.json()

        if 'error_code' in resp_json:
            if resp_json['error_code'] == -343:
                raise ExceedCallQuotaError('Exceed call quota')
            if resp_json['error_code'] == -380:
                raise InvalidSignedInfo('Invalid signed info')

    # ====== GETTERS ======

    # user
    def get_user(self, user_id: str):
        return get_user(self, user_id)

    def get_users_from_dict(self, resp: dict):
        return get_users_from_dict(self, resp)

    def get_letters_from_dict(self, resp: dict):
        return get_letters_from_dict(self, resp)

    def get_letters(self, user_id: str):
        return get_letters(self, user_id)

    def get_joined_groups(self, user_id: str, amount=100):
        return get_joined_groups(self, user_id, amount)

    def get_user_followers(self, user_id: str, amount: int = None):
        return get_user_followers(self, user_id, amount)

    def get_user_followings(self, user_id, amount=None):
        return get_user_followings(self, user_id, amount)

    def get_follow_requests(self, amount=100):
        return get_follow_requests(self, amount)

    def get_user_active_call(self, user_id: str):
        return get_user_active_call(self, user_id)

    def get_blocked_users(self, amount=100):
        return get_blocked_users(self, amount)

    # post
    def get_post(self, post_id: str):
        return get_post(self, post_id)

    def get_posts_from_dict(self, resp: dict):
        return get_posts_from_dict(self, resp)

    def get_timeline(self, user_id: str = None, keyword: str = None, hashtag: str = None, amount=100):
        return get_timeline(self, user_id, keyword, hashtag, amount)

    def get_following_timeline(self, amount=50):
        return get_following_timeline(self, amount)

    def get_conversation(self, conversation_id: str = None, post_id: str = None, amount=100):
        return get_conversation(self, conversation_id, post_id, amount)

    def get_reposts(self, post_id, amount=100):
        return get_reposts(self, post_id, amount)

    def get_post_likers(self, post_id, amount=50):
        return get_post_likers(self, post_id, amount)

    # group
    def get_group(self, group_id: str):
        return get_group(self, group_id)

    def get_groups_from_dict(self, resp: dict):
        return get_groups_from_dict(self, resp)

    def get_group_users_from_dict(self, resp: dict):
        return get_group_users_from_dict(self, resp)

    def get_group_timeline(self, group_id: str, amount=100):
        return get_group_timeline(self, group_id, amount)

    def get_group_call(self, group_id: str):
        return get_group_call(self, group_id)

    def get_group_members(self, group_id: str, amount=100):
        return get_group_members(self, group_id, amount)

    def get_pending_users_in_group(self, group_id: str, amount=100):
        return get_pending_users_in_group(self, group_id, amount)

    def get_banned_user_from_group(self, group_id: str, amount=100):
        return get_banned_user_from_group(self, group_id, amount)

    # chat
    def get_chat_room(self, chatroom_id: str):
        return get_chat_room(self, chatroom_id)

    def get_chat_rooms_from_dict(self, resp: dict):
        return get_chat_rooms_from_dict(self, resp)

    def get_chat_messages_from_dict(self, resp: dict):
        return get_chat_messages_from_dict(self, resp)

    def get_chat_room_id_from_user(self, user_id: str):
        return get_chat_room_id_from_user(self, user_id)

    def get_chat_messages(self, chatroom_id: str = None, user_id: str = None, amount: int = None):
        return get_chat_messages(self, chatroom_id, user_id, amount)

    def get_chat_rooms(self, amount: int = None):
        return get_chat_rooms(self, amount)

    def get_chat_requests(self, amount: int = None):
        return get_chat_requests(self, amount)

    # notification
    def get_activities_from_dict(self, resp: dict):
        return get_activities_from_dict(self, resp)

    def get_notification(self, important=True):
        return get_notification(self, important)

    # ====== USER ======

    def follow_user(self, user_id: str):
        return follow_user(self, user_id)

    def unfollow_user(self, user_id: str):
        return unfollow_user(self, user_id)

    def accept_follow_request(self, user_id: str):
        return accept_follow_request(self, user_id)

    def reject_follow_request(self, user_id: str):
        return reject_follow_request(self, user_id)

    def send_letter(self, user_id: str, message: str):
        return send_letter(self, user_id, message)

    def block_user(self, user_id: str):
        return block_user(self, user_id)

    def unblock_user(self, user_id: str):
        return unblock_user(self, user_id)

    # ====== POST ======

    def create_post(self, text, color=0, font_size=0, choices: list = None, type: str = None):
        """
        Color:
        -----
            normal colors: 0 - 7
            special colors: 1001 - 1007

        Font_size:
        -----
            size: 0 - 4 (The larger the number, the larger the font size.)
        """
        return create_post(self, text, color, font_size, choices, type)

    def create_post_in_group(self, group_id: str, text: str, color=0, font_size=0):
        return create_post_in_group(self, group_id, text, color, font_size)

    def create_repost(self, text: str, post_id: str, color=0, font_size=0):
        return create_repost(self, text, post_id, color, font_size)

    def create_reply(self, text: str, post_id: str, color=0, font_size=0):
        return create_reply(self, text, post_id, color, font_size)

    def delete_post(self, post_id: str):
        return delete_post(self, post_id)

    def pin_post(self, post_id: str):
        return pin_post(self, post_id)

    def unpin_post(self, post_id: str):
        return unpin_post(self, post_id)

    def like_post(self, post_id: str):
        return like_post(self, post_id)

    def unlike_post(self, post_id: str):
        return unlike_post(self, post_id)

    # ====== GROUP ======

    def create_group(
        self,
        group_name: str,
        description: str = None,
        guidelines: str = None,
        group_category_id=21,
        sub_category_id: int = None,
        is_private=False,
        call_timeline_display=True,
        hide_reported_posts=False,
        allow_ownership_transfer=True,
        allow_thread_creation_by='member',
        allow_members_to_post_image_and_video=True,
        allow_members_to_post_url=True,
        hide_conference_call=False,
        secret=False,
        only_verified_age=False,
        only_mobile_verified=False,
        gender=-1,
        generation_groups_limit=0
    ):
        return create_group(
            self,
            group_name,
            description,
            guidelines,
            group_category_id,
            sub_category_id,
            is_private,
            call_timeline_display,
            hide_reported_posts,
            allow_ownership_transfer,
            allow_thread_creation_by,
            allow_members_to_post_image_and_video,
            allow_members_to_post_url,
            hide_conference_call,
            secret,
            only_verified_age,
            only_mobile_verified,
            gender,
            generation_groups_limit
        )

    def delete_group(self, group_id: str):
        return delete_group(self, group_id)

    def change_group_settings(
        self,
        group_id,
        group_name: str,
        description: str = None,
        guidelines: str = None,
        group_category_id=21,
        sub_category_id: int = None,
        is_private=False,
        call_timeline_display=True,
        hide_reported_posts=False,
        allow_ownership_transfer=True,
        allow_thread_creation_by='member',
        allow_members_to_post_image_and_video=True,
        allow_members_to_post_url=True,
        hide_conference_call=False,
        hide_from_game_eight=False,
        secret=False,
        only_verified_age=False,
        only_mobile_verified=False,
        gender=-1,
        generation_groups_limit=0
    ):
        return change_group_settings(
            self,
            group_id,
            group_name,
            description,
            guidelines,
            group_category_id,
            sub_category_id,
            is_private,
            call_timeline_display,
            hide_reported_posts,
            allow_ownership_transfer,
            allow_thread_creation_by,
            allow_members_to_post_image_and_video,
            allow_members_to_post_url,
            hide_conference_call,
            hide_from_game_eight,
            secret,
            only_verified_age,
            only_mobile_verified,
            gender,
            generation_groups_limit
        )

    def transfer_group_ownership(self, group_id: str, user_id: str):
        return transfer_group_ownership(self, group_id, user_id)

    def offer_group_sub_owner(self, group_id: str, user_id: str):
        return offer_group_sub_owner(self, group_id, user_id)

    def undo_group_ownership_transfer(self, group_id: str, user_id: str):
        return undo_group_ownership_transfer(self, group_id, user_id)

    def undo_group_sub_owner_offer(self, group_id: str, user_id: str):
        return undo_group_sub_owner_offer(self, group_id, user_id)

    def fire_group_sub_owner(self, group_id: str, user_id: str):
        return fire_group_sub_owner(self, group_id, user_id)

    def accept_group_join_request(self, group_id: str, user_id: str):
        return accept_group_join_request(self, group_id, user_id)

    def decline_group_join_request(self, group_id: str, user_id: str):
        return decline_group_join_request(self, group_id, user_id)

    def invite_user_to_group(self, group_id: str, user_id: str):
        return invite_user_to_group(self, group_id, user_id)

    def pin_group_post(self, group_id: str, post_id: str):
        return pin_group_post(self, group_id, post_id)

    def unpin_group_post(self, group_id: str):
        return unpin_group_post(self, group_id)

    def ban_user_from_group(self, group_id: str, user_id: str):
        return ban_user_from_group(self, group_id, user_id)

    def unban_user_from_group(self, group_id: str, user_id: str):
        return unban_user_from_group(self, group_id, user_id)

    def join_group(self, group_id: str):
        return join_group(self, group_id)

    def leave_group(self, group_id: str):
        return leave_group(self, group_id)

    # ====== CHAT ======

    def send_message(self, message: str, user_id: str = None, chat_room_id: str = None):
        return send_message(self, message, user_id, chat_room_id)

    def accept_chat_request(self, chat_room_id: str):
        return accept_chat_request(self, chat_room_id)

    def delete_chat_room(self, chat_room_id: str):
        return delete_chat_room(self, chat_room_id)

    # ====== SUPPORT ======

    def console_print(self, text: str, color: str = None):
        return console_print(text, color)
