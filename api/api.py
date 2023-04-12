import requests
import time


from tqdm import tqdm


from auth import YayAuth
from endpoints import Endpoints as ep
from handle_exceptions import *
from models import *


class Yay(object):

    def __init__(self, token=None, proxy=None, timeout=10):
        self.auth = YayAuth(proxy=proxy, timeout=timeout)

        if token:
            self.access_token = token
            self.auth.access_token = token
            self.auth.headers.setdefault('Authorization', f'Bearer {token}')

    def login(self, email, password):
        self.auth.login(email, password)
        self.access_token = self.auth.access_token
        self.api_key = self.auth.api_key
        self.logged_in_as = self.auth.logged_in_as

    def logout(self):
        self.auth.logout()
        self.access_token = None
        self.api_key = None
        self.logged_in_as = None

    def _get(self, url: str):
        resp = requests.get(url, headers=self.auth.headers,
                            proxies={'http': self.auth.proxy,
                                     'https': self.auth.proxy},
                            timeout=self.auth.timeout)
        self._handle_response(resp)
        return resp.json()

    def _post(self, url: str, data: dict = None):
        resp = requests.post(url, params=data,
                             headers=self.auth.headers,
                             proxies={'http': self.auth.proxy,
                                      'https': self.auth.proxy},
                             timeout=self.auth.timeout)
        self._handle_response(resp)
        return resp.json()

    def _put(self, url: str, data: dict = None):
        resp = requests.put(url, params=data,
                            headers=self.auth.headers,
                            proxies={'http': self.auth.proxy,
                                     'https': self.auth.proxy},
                            timeout=self.auth.timeout)
        self._handle_response(resp)
        return resp.json()

    def _delete(self, url: str, data: dict = None):
        resp = requests.delete(url, params=data,
                               headers=self.auth.headers,
                               proxies={'http': self.auth.proxy,
                                        'https': self.auth.proxy},
                               timeout=self.auth.timeout)
        self._handle_response(resp)
        return resp.json()

    def _handle_response(self, resp):
        if resp.status_code == 401:
            raise AuthenticationError('Failed to authenticate')
        if resp.status_code == 429:
            raise RateLimitError('Rate limit exceeded')

        resp_json = resp.json()

        if resp_json.get('error_code') == '-343':
            raise ExceedCallQuotaError('Exceed call quota')

    # ----- GET USER -----

    def get_user(self, user_id: str):
        resp = self._get(f'{ep.USER_v2}/{user_id}')
        user_data = resp.get('user')
        return self.user_object(user_data)

    def get_users_from_dict(self, resp: dict):
        assert 'users' in resp, "'users' key not found"
        users_data = resp.get('users')
        users = []
        for user_data in users_data:
            user = self.user_object(user_data)
            users.append(user)
        return users

    def get_letters_from_dict(self, resp: dict):
        assert 'reviews' in resp, "'reviews' key not found"
        reviews_data = resp.get('reviews')
        reviews = []
        for review_data in reviews_data:
            review = self.review_object(review_data)
            reviews.append(review)
        return reviews

    def get_letters(self, user_id, amount: int = None):
        amount = float('inf') if amount is None else amount
        number = min(amount, 100)

        resp = self._get(
            f'{ep.USER_v1}/reviews/{user_id}?not_active=false&number={number}')
        reviews = self.get_letters_from_dict(resp)

        next_item = resp.get('reviews')[-1]
        next_id = next_item.get('id')
        reviews_count = self.get_user(user_id).reviews_count if amount == float(
            'inf') else amount
        amount -= 100

        with tqdm(total=reviews_count, desc='Extracting Letters') as pbar:
            while next_id and amount > 0:
                number = min(amount, 100)

                resp = self._get(
                    f'{ep.USER_v1}/reviews/{user_id}?from_id={next_id}&not_active=false&number={number}')
                reviews.extend(self.get_letters_from_dict(resp))

                if len(resp.get('reviews')) == 0:
                    break
                next_item = resp.get('reviews')[-1]
                next_id = next_item.get('id')
                amount -= 100

                pbar.update(number)

        return reviews

    def get_joined_groups(self, user_id, amount=100):
        resp = self._get(
            f'{ep.GROUP_v1}/user_group_list?number={amount}&page=0&user_id={user_id}')
        return self.get_groups_from_dict(resp)

    def get_user_followers(self, user_id, amount: int = None):
        amount = float('inf') if amount is None else amount
        number = min(amount, 50)

        resp = self._get(
            f'{ep.USER_v2}/{user_id}/web_followers?number={number}')
        users = self.get_users_from_dict(resp)

        next_id = resp.get('last_follow_id')
        followers_count = self.get_user(user_id).followers_count if amount == float(
            'inf') else amount
        amount -= 50

        with tqdm(total=followers_count, desc='Extracting Followers') as pbar:
            while next_id and amount > 0:
                number = min(amount, 50)

                resp = self._get(
                    f'{ep.USER_v2}/{user_id}/web_followers?from_follow_id={next_id}&number={number}')
                users.extend(self.get_users_from_dict(resp))

                next_id = resp.get('last_follow_id')
                amount -= 50

                pbar.update(number)

        return users

    def get_user_followings(self, user_id, amount: int = None):
        amount = float('inf') if amount is None else amount
        number = min(amount, 50)

        resp = self._get(
            f'{ep.USER_v2}/{user_id}/web_followings?number={number}')
        users = self.get_users_from_dict(resp)

        next_id = resp.get('last_follow_id')
        followings_count = self.get_user(user_id).followings_count if amount == float(
            'inf') else amount
        amount -= 50

        with tqdm(total=followings_count, desc='Extracting Followings') as pbar:
            while next_id and amount > 0:
                number = min(amount, 50)

                resp = self._get(
                    f'{ep.USER_v2}/{user_id}/web_followings?from_follow_id={next_id}&number=50')
                users.extend(self.get_users_from_dict(resp))

                next_id = resp.get('last_follow_id')
                amount -= 50

                pbar.update(number)

        return users

    def get_follow_requests(self, amount=100):
        # has last_timestamp: 1658812381
        resp = self._get(
            f'{ep.USER_v2}/follow_requests?number=50')
        return self.get_users_from_dict(resp)

    def get_user_active_call(self, user_id):
        resp = self._get(
            f'{ep.POST_v1}/active_call?user_id={user_id}')
        post_id = resp['post'].get('id')
        return self.get_post(post_id)

    def get_blocked_users(self, amount: int = 100):
        resp = self._get(
            f'{ep.USER_v2}/blocked')
        return self.get_users_from_dict(resp)

    # ----- GET POST -----

    def get_post(self, post_id: str):
        resp = self._get(f'{ep.POST_v2}/{post_id}')
        post_data = resp.get('post')
        return self.post_object(post_data)

    def get_posts_from_dict(self, resp: dict):
        assert 'posts' in resp, "'posts' key not found"
        posts_data = resp.get('posts')
        posts = []
        for post_data in posts_data:
            post = self.post_object(post_data)
            posts.append(post)
        return posts

    def get_timeline(self, user_id=None, keyword=None, hashtag=None, amount=100):
        if user_id:
            resp = self._get(
                f'{ep.GET_USER_TIMELINE}?number={amount}&user_id={user_id}')
        elif keyword:
            resp = self._get(
                f'{ep.GET_TIMELINE_BY_KEYWORD}?keyword={keyword}&number={amount}')
        elif hashtag:
            resp = self._get(
                f'{ep.GET_TIMELINE_BY_HASHTAG}/{hashtag}?number={amount}')
        else:
            resp = self._get(
                f'{ep.GET_TIMELINE}?number={amount}')
        return self.get_posts_from_dict(resp)

    def get_following_timeline(self, amount=50):
        if amount > 50:
            pass
        # has next_page_value
        resp = self._get(
            f'{ep.GET_FOLLOWING_TIMELINE}?number=50')
        return self.get_posts_from_dict(resp)

    def get_conversation(self, conversation_id: str = None, post_id: str = None, amount=100):
        if post_id:
            conversation_id = self.get_post(post_id).conversation_id
        resp = self._get(
            f'{ep.GET_CONVERSATION}/{conversation_id}?number={amount}&reverse=true')
        return self.get_posts_from_dict(resp)

    def get_reposts(self, post_id, amount=100):
        resp = self._get(
            f'{ep.POST_v2}/{post_id}/reposts?number=100')
        return self.get_posts_from_dict(resp)

    def get_post_likers(self, post_id, amount=50):
        # has last_id
        resp = self._get(
            f'{ep.POST_v1}/{post_id}/likers')
        return self.get_users_from_dict(resp)

    # ----- GET GROUP -----

    def get_group(self, group_id: str):
        resp = self._get(f'{ep.GROUP_v1}/{group_id}')
        group_data = resp.get('group')
        return self.group_object(group_data)

    def get_groups_from_dict(self, resp: dict):
        assert 'groups' in resp, "'groups' key not found"
        groups_data = resp.get('groups')
        groups = []
        for group_data in groups_data:
            group = self.group_object(group_data)
            groups.append(group)
        return groups

    def get_group_users_from_dict(self, resp: dict):
        assert 'group_users' in resp, "'group_users' key not found"
        users_data = resp.get('group_users')
        users = []
        for user_data in users_data:
            user = self.group_user_object(user_data)
            users.append(user)
        return users

    def get_group_timeline(self, group_id, amount=100):
        # {ep.POST_v2}/group_timeline?from_post_id=111&group_id={group_id}&number=100
        resp = self._get(
            f'{ep.POST_v2}/group_timeline?group_id={group_id}&number=100')
        return self.get_posts_from_dict(resp)

    def get_group_call(self, group_id):
        resp = self._get(
            f'{ep.POST_v2}/call_timeline?group_id={group_id}&number=20')
        return self.get_posts_from_dict(resp)

    def get_group_members(self, group_id, amount=100):
        resp = self._get(
            f'{ep.GROUP_v2}/{group_id}/members?number=100')
        return self.get_group_users_from_dict(resp)

    def get_pending_users_in_group(self, group_id, amount=100):
        resp = self._get(
            f'{ep.GROUP_v2}/{group_id}/{group_id}/members?mode=pending&number=100')
        return self.get_users_from_dict(resp)

    def get_banned_user_from_group(self, group_id, amount=100):
        resp = self._get(
            f'{ep.GROUP_v1}/{group_id}/{group_id}/ban_list?number=100')
        return self.get_users_from_dict(resp)

    # ----- GET CHAT -----

    def get_chat_room(self, chatroom_id: str):
        resp = self._get(f'{ep.CHATROOM_v2}/{chatroom_id}')
        chat_room_data = resp.get('chat')
        return self.chat_room_object(chat_room_data)

    def get_chat_rooms_from_dict(self, resp: dict):
        assert 'chat_rooms' in resp, "'chat_rooms' key not found"
        chat_rooms_data = resp.get('chat_rooms')
        chats = []
        for chat_room_data in chat_rooms_data:
            chat = self.chat_room_object(chat_room_data)
            chats.append(chat)
        return chats

    def get_chat_messages_from_dict(self, resp: dict):
        assert 'messages' in resp, "'messages' key not found"
        messages_data = resp.get('messages')
        messages = []
        for message_data in messages_data:
            message = self.message_object(message_data)
            messages.append(message)
        return messages

    def get_chat_room_id_from_user(self, user_id):
        data = {'with_user_id': user_id}
        resp = self._post(
            f'{ep.CHATROOM_v1}/new', data)
        return resp['room_id']

    def get_chat_messages(self, chatroom_id=None, user_id=None, amount=None):
        resp = self._get(
            f'{ep.CHATROOM_v2}/{chatroom_id}/messages?number=100')
        return self.get_chat_messages_from_dict(resp)

    def get_chat_rooms(self, amount=None):
        resp = self._get(
            f'{ep.CHATROOM_v1}/main_list?number=100')
        return self.get_chat_rooms_from_dict(resp)

    def get_chat_requests(self, amount=None):
        resp = self._get(
            f'{ep.CHATROOM_v1}/request_list?number=100')
        return self.get_chat_rooms_from_dict(resp)

    # ----- GET NOTIFICATION -----

    def get_activities_from_dict(self, resp: dict):
        assert 'activities' in resp, "'activities' key not found"
        activities_data = resp.get('activities')
        activities = []
        for activity_data in activities_data:
            activity = self.activity_object(activity_data)
            activities.append(activity)
        return activities

    def get_notification(self, important=True):
        if important:
            resp = self._get(
                f'{ep.CAS_BASE_URL}/api/user_activities?important=true&number=100')
        else:
            resp = self._get(
                f'{ep.CAS_BASE_URL}/api/user_activities?important=false&number=100')
        return self.get_activities_from_dict(resp)

    # -----------------------

    # ----- ACTION USER -----

    # -----------------------

    def follow_user(self, user_id):
        resp = self._post(f'{ep.USER_v2}/{user_id}/follow')
        return resp

    def unfollow_user(self, user_id):
        resp = self._post(f'{ep.USER_v2}/{user_id}/unfollow')
        return resp

    def accept_follow_request(self, user_id):
        resp = self._post(
            f'{ep.USER_v2}/{user_id}/follow_request?action=accept')
        return resp

    def reject_follow_request(self, user_id):
        resp = self._post(
            f'{ep.USER_v2}/{user_id}/follow_request?action=reject')
        return resp

    def send_letter(self, user_id, message):
        data = {'comment': message}
        resp = self._post(
            f'{ep.USER_v1}/reviews/{user_id}', data)
        return resp

    def block_user(self, user_id):
        resp = self._post(
            f'{ep.USER_v1}/{user_id}/block')
        return resp

    def unblock_user(self, user_id):
        resp = self._post(
            f'{ep.USER_v1}/{user_id}/unblock')
        return resp

    # ----- ACTION POST -----

    # post_type -> text, questionaire, image, video, call, video call
    def create_post(self, text, color=0, font_size=0):
        """
        Color:
        -----
            normal colors: 0 - 7
            special colors: 1001 - 1007

        Font_size:
        -----
            size: 0 - 4 (The larger the number, the larger the font size.)
        """
        data = {
            'text': text,
            'color': color,
            'font_size': font_size
        }
        resp = self._post(
            f'{ep.BASE_URL}/v1/web/posts/new', data)
        return resp

    def create_post_in_group(self, group_id, text, color=0, font_size=0):
        data = {
            'color': color,
            'font_size': font_size,
            'group_id': group_id,
            'text': text,
            'post_type': 'text',
            'uuid': ''
        }
        resp = self._post('https://yay.space/api/posts', data)
        return self.get_post(resp['id'])

    def create_repost(self, text, post_id, color=0, font_size=0):
        data = {
            'text': text,
            'color': color,
            'font_size': font_size,
            'post_id': post_id,
            'message_tags': '[]',
            'post_type': 'text'
        }
        resp = self._post(
            f'{ep.POST_v3}/repost', data)
        return resp

    def create_reply(self, text, post_id, color=0, font_size=0):
        data = {
            'text': text,
            'color': color,
            'font_size': font_size,
            'in_reply_to': post_id
        }
        resp = self._post(
            f'{ep.BASE_URL}/v1/web/posts/new', data)
        return resp

    def delete_post(self, post_id):
        data = {'posts_ids[]': post_id}
        resp = self._post(
            f'{ep.POST_v2}/mass_destroy', data)
        return resp

    def pin_post(self, post_id):
        data = {'id': post_id}
        resp = self._post(
            f'{ep.PIN_v1}/posts', data)
        return resp

    def unpin_post(self, post_id):
        resp = self._post(
            f'{ep.PIN_v1}/posts/{post_id}')
        return resp

    def like_post(self, post_id):
        data = {'post_ids': post_id}
        resp = self._post(
            f'{ep.POST_v2}/like', data)
        return resp

    def unlike_post(self, post_id):
        resp = self._post(
            f'{ep.POST_v1}/{post_id}/unlike')
        return resp

    # ----- ACTION GROUP -----

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
        data = {
            'topic': group_name,
            'description': description,
            'guidelines': guidelines,
            'group_category_id': group_category_id,
            'sub_category_id': sub_category_id,
            'is_private': is_private,
            'call_timeline_display': call_timeline_display,
            'hide_reported_posts': hide_reported_posts,
            'allow_ownership_transfer': allow_ownership_transfer,
            'allow_thread_creation_by': allow_thread_creation_by,
            'allow_members_to_post_image_and_video': allow_members_to_post_image_and_video,
            'allow_members_to_post_url': allow_members_to_post_url,
            'hide_conference_call': hide_conference_call,
            'secret': secret,
            'only_verified_age': only_verified_age,
            'only_mobile_verified': only_mobile_verified,
            'gender': gender,
            'generation_groups_limit': generation_groups_limit
        }
        # {ep.BASE_URL}/v1/groups/new   method=post
        resp = self._post(
            f'{ep.GROUP_v1}/new', data)
        return resp

    def delete_group(self, group_id):
        data = {'groupId': group_id}
        resp = self._delete(
            f'{ep.GROUP_v1}/{group_id}/leave', data)
        return resp

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
        data = {
            'topic': group_name,
            'description': description,
            'guidelines': guidelines,
            'group_category_id': group_category_id,
            'sub_category_id': sub_category_id,
            'is_private': is_private,
            'call_timeline_display': call_timeline_display,
            'hide_reported_posts': hide_reported_posts,
            'allow_ownership_transfer': allow_ownership_transfer,
            'allow_thread_creation_by': allow_thread_creation_by,
            'allow_members_to_post_image_and_video': allow_members_to_post_image_and_video,
            'allow_members_to_post_url': allow_members_to_post_url,
            'hide_conference_call': hide_conference_call,
            'hide_from_game_eight': hide_from_game_eight,
            'secret': secret,
            'only_verified_age': only_verified_age,
            'only_mobile_verified': only_mobile_verified,
            'gender': gender,
            'generation_groups_limit': generation_groups_limit,
            'uuid': ''
        }
        resp = self._put(
            f'https://yay.space/api/groups/{group_id}', data)
        return self.group_object(resp.get('group'))

    def transfer_group_ownership(self, group_id, user_id):
        # need a fix. json.decoder.JSONDecodeError: Expecting value: line 1 column 1 (char 0)
        # the resp: <Response [400]>, probably the uuid is invalid.
        data = {
            'uuid': '',
            'user_id': user_id
        }
        resp = self._post(
            f'https://yay.space/api/groups/{group_id}/transfer', data)
        return resp

    def offer_group_sub_owner(self, group_id, user_id):
        # need a fix. json.decoder.JSONDecodeError: Expecting value: line 1 column 1 (char 0)
        # the resp: <Response [400]>, probably the uuid is invalid.
        data = {'uuid': '', 'user_ids[]': user_id}
        resp = self._post(
            f'https://yay.space/api/groups/{group_id}/deputize', data)
        return resp

    def undo_group_ownership_transfer(self, group_id, user_id):
        data = {'user_id': user_id}
        resp = self._put(
            f'{ep.GROUP_v1}/{group_id}/transfer/withdraw', data)
        return resp

    def undo_group_sub_owner_offer(self, group_id, user_id):
        resp = self._put(
            f'{ep.GROUP_v1}/{group_id}/deputize/{user_id}/withdraw')
        return resp

    def fire_group_sub_owner(self, group_id, user_id):
        resp = self._post(
            f'{ep.GROUP_v1}/{group_id}/fire/{user_id}')
        return resp

    def accept_group_join_request(self, group_id, user_id):
        data = {
            'groupId': group_id,
            'mode': 'accept',
            'userId': user_id
        }
        resp = self._post(
            f'{ep.GROUP_v1}/{group_id}/accept/{user_id}', data)
        return resp

    def decline_group_join_request(self, group_id, user_id):
        data = {
            'groupId': group_id,
            'mode': 'decline',
            'userId': user_id
        }
        resp = self._post(
            f'{ep.GROUP_v1}/{group_id}/decline/{user_id}', data)
        return resp

    def invite_user_to_group(self, group_id, user_id):
        data = {'user_ids[]': user_id}
        resp = self._post(
            f'{ep.GROUP_v1}/{group_id}/invite', data)
        return resp

    def pin_group_post(self, group_id, post_id):
        data = {'group_id': group_id, 'post_id': post_id}
        resp = self._put(
            f'{ep.POST_v2}/group_pinned_post', data)
        return resp

    def unpin_group_post(self, group_id):
        data = {'group_id': group_id}
        resp = self._delete(
            f'{ep.POST_v2}/group_pinned_post', data)
        return resp

    def ban_user_from_group(self, group_id, user_id):
        resp = self._post(
            f'{ep.GROUP_v1}/{group_id}/ban/{user_id}')
        return resp

    def unban_user_from_group(self, group_id, user_id):
        resp = self._post(
            f'{ep.GROUP_v1}/{group_id}/unban/{user_id}')
        return resp

    def join_group(self, group_id):
        self._post(f'{ep.GROUP_v1}/{group_id}/join')
        return self.get_group(group_id)

    def leave_group(self, group_id):
        data = {'groupId': group_id}
        resp = self._delete(
            f'{ep.GROUP_v1}/{group_id}/leave', data)
        return resp

    # ----- ACTION CHAT -----

    def send_message(self, message, user_id=None, chat_room_id=None):
        if user_id:
            chat_room_id = self.get_chat_room_id_from_user(user_id)
        data = {'message_type': 'text', 'text': message}
        resp = self._post(
            f'{ep.CHATROOM_v1}/{chat_room_id}/messages/new', data)
        return resp

    def accept_chat_request(self, chat_room_id):
        data = {'chat_room_ids[]': chat_room_id}
        resp = self._post(
            f'{ep.CHATROOM_v1}/accept_chat_request', data)
        return resp

    def delete_chat_room(self, chat_room_id):
        data = {'chat_room_ids[]': chat_room_id}
        resp = self._post(
            f'{ep.CHATROOM_v1}/mass_destroy', data)
        return resp

    # ----- CREATE OBJECTS -----

    def user_object(self, user_data: dict) -> User:
        def get_val(key):
            return user_data.get(key, None)

        user = User(
            id=get_val('id'),
            username=get_val('nickname'),
            bio=get_val('biography'),
            badge=user_data.get('title', None),
            num_followers=get_val('followers_count'),
            num_followings=get_val('followers_count'),
            private_user=get_val('is_private'),
            num_posts=get_val('posts_count'),
            num_joined_groups=get_val('groups_users_count'),
            num_reviews=get_val('reviews_count'),
            verified_age=get_val('age_verified'),
            country_code=get_val('country_code'),
            is_vip=get_val('vip'),
            hide_vip=get_val('hide_vip'),
            online_status=get_val('online_status'),
            profile_image=get_val('profile_icon'),
            profile_thumbnail=get_val(
                'profile_icon_thumbnail'),
            cover_image=get_val('cover_image'),
            cover_thumbnail=get_val('cover_image_thumbnail'),
            last_logged_in_at=get_val('last_loggedin_at'),
            mutual_chat_enabled=get_val('mutual_chat'),
            chat_request_enabled=get_val('chat_request'),
            chat_phone_verification_required=get_val(
                'chat_required_phone_verification'),
            age_restricted_review=get_val(
                'age_restricted_on_review'),
            following_restricted_review=get_val(
                'following_restricted_on_review'),
            review_restricted_by=get_val('restricted_review_by'),
            recently_banned=get_val('recently_kenta'),
            dangerous_user=get_val('dangerous_user'),
            new_user=get_val('new_user'),
            selected_interests=get_val('interests_selected')
        )
        return user

    def group_user_object(self, user_data: dict) -> User:
        def get_val(key):
            return user_data['user'].get(key, None)

        group_user = GroupUser(
            moderator=user_data.get('is_moderator', None),
            banned=user_data.get('banned', None),
            pending_transfer=user_data.get('pending_transfer', None),
            pending_deputize=user_data.get('pending_deputize', None),
            badge=user_data.get('title', None),
            id=get_val('id'),
            username=get_val('nickname'),
            bio=get_val('biography'),
            num_followers=get_val('followers_count'),
            is_private=get_val('is_private'),
            num_posts=get_val('posts_count'),
            num_joined_groups=get_val('groups_users_count'),
            num_reviews=get_val('reviews_count'),
            verified_age=get_val('age_verified'),
            country_code=get_val('country_code'),
            is_vip=get_val('vip'),
            hide_vip=get_val('hide_vip'),
            online_status=get_val('online_status'),
            profile_image=get_val('profile_icon'),
            profile_thumbnail=get_val(
                'profile_icon_thumbnail'),
            cover_image=get_val('cover_image'),
            cover_thumbnail=get_val('cover_image_thumbnail'),
            last_logged_in_at=get_val('last_loggedin_at'),
            mutual_chat_enabled=get_val('mutual_chat'),
            chat_request_enabled=get_val('chat_request'),
            chat_phone_verification_required=get_val(
                'chat_required_phone_verification'),
            age_restricted_review=get_val(
                'age_restricted_on_review'),
            following_restricted_review=get_val(
                'following_restricted_on_review'),
            review_restricted_by=get_val('restricted_review_by'),
            recently_banned=get_val('recently_kenta'),
            dangerous_user=get_val('dangerous_user'),
            new_user=get_val('new_user'),
            selected_interests=get_val('interests_selected')
        )
        return group_user

    def post_object(self, post_data: dict) -> Post:
        def get_val(key):
            return post_data.get(key, None)

        post = Post(
            id=get_val('id'),
            author_id=post_data['user'].get('id'),
            author_username=post_data['user'].get('nickname', None),
            text=get_val('text'),
            group_id=get_val('group_id'),
            font_size=get_val('font_size'),
            liked=get_val('liked'),
            num_likes=get_val('likes_count'),
            type=get_val('post_type'),
            color=get_val('color'),
            num_reposted=get_val('reposts_count'),
            created_at=get_val('created_at'),
            updated_at=get_val('updated_at'),
            edited_at=get_val('edited_at'),
            num_reported=get_val('reported_count'),
            reply_to_id=get_val('in_reply_to_post'),
            num_reply_to=get_val(
                'in_reply_to_post_count'),
            repostable=get_val('repostable'),
            highlighted=get_val('highlighted'),
            hidden=get_val('hidden'),
            thread_id=get_val('thread_id'),
            message_tags=get_val('message_tags'),
            conversation_id=get_val('conversation_id'),
            attachment=get_val('attachment'),
            attachment_thumbnail=get_val('attachment_thumbnail'),
            shared_url=post_data.get('shared_url', {}).get('url', None)
        )
        return post

    def review_object(self, review_data: dict) -> Post:
        def get_val(key):
            return review_data.get(key, None)

        review = Review(
            text=get_val('comment'),
            created_at=get_val('created_at'),
            id=get_val('id'),
            mutual_review_enabled=get_val('mutual_review'),
            num_reported=get_val('reported_count'),
            author_id=review_data['reviewer'].get('id', None),
            author_username=review_data['reviewer'].get('nickname', None),
        )

        return review

    def group_object(self, group_data: dict) -> Group:
        def get_val(key):
            return group_data.get(key, None)

        group = Group(
            members_can_post_image_and_video=get_val(
                'allow_members_to_post_image_and_video'),
            members_can_post_url=get_val('allow_members_to_post_url'),
            ownership_transfer_allowed=get_val('allow_ownership_transfer'),
            allowed_thread_creators=get_val('allow_thread_creation_by'),
            call_timeline=get_val('call_timeline_display'),
            cover_image=get_val('cover_image'),
            cover_thumbnail=get_val('cover_image_thumbnail'),
            description=get_val('description'),
            gender=get_val('gender'),
            generation_groups_limit=get_val('generation_groups_limit'),
            category_id=get_val('group_category_id'),
            icon=get_val('group_icon'),
            num_groups_members=get_val('groups_users_count'),
            guidelines=get_val('guidelines'),
            conference_call_hidden=get_val('hide_conference_call'),
            game_eight_hidden=get_val('hide_from_game_eight'),
            reported_posts_hidden=get_val('hide_reported_posts'),
            num_highlighted=get_val('highlighted_count'),
            homepage=get_val('homepage'),
            id=get_val('group_id'),
            invited_to_join=get_val('invited_to_join'),
            joined=get_val('is_joined'),
            pending=get_val('is_pending'),
            private_group=get_val('is_private'),
            moderator_ids=get_val('moderator_ids'),
            mobile_verified_only=get_val('only_mobile_verified'),
            verified_age_only=get_val('only_verified_age'),
            owner_id=get_val('user_id'),
            owner_username=group_data['owner'].get('nickname', None),
            num_pendings=get_val('pending_count'),
            pending_deputize_ids=get_val('pending_deputize_ids'),
            pending_transfer_id=get_val('pending_transfer_id'),
            place=get_val('place'),
            num_posts=get_val('posts_count'),
            num_related_groups=get_val('related_count'),
            safe_mode_enabled=get_val('safe_mode'),
            hidden_grouop=get_val('secret'),
            seizable=get_val('seizable'),
            seizable_before=get_val('seizable_before'),
            sub_category_id=get_val('sub_category_id'),
            num_threads=get_val('threads_count'),
            category=get_val('title'),
            group_name=get_val('topic'),
            num_unread_threads=get_val('unread_threads_count'),
            updated_at=get_val('updated_at'),
            num_views=get_val('views_count'),
            walkthrough_requested=get_val('walkthrough_requested'),
        )

        return group

    def chat_room_object(self, chat_room_data: dict) -> Post:
        def get_val(key):
            return chat_room_data.get(key, None)

        chat_room = ChatRoom(
            background_image=get_val('background'),
            background_thumbnail=get_val('background_thumbnail'),
            icon=get_val('icon'),
            icon_thumbnail=get_val('icon_thumbnail'),
            id=get_val('id'),
            is_group=get_val('is_group'),
            is_request=get_val('is_request'),
            last_message=chat_room_data['last_message'].get('text', None),
            member_ids=[member.get('id', None)
                        for member in chat_room_data['members']],
            member_usernames=[member.get(
                'nickname', None) for member in chat_room_data['members']],
            chat_title=get_val('name'),
            num_unread=get_val('unread_count'),
            updated_at=get_val('updated_at')
        )

        return chat_room

    def message_object(self, message_data: dict) -> Post:
        def get_val(key):
            return message_data.get(key, None)

        message = Message(
            attachment=get_val('attachment'),
            attachment_android=get_val('attachment_android'),
            attachment_thumbnail=get_val('attachment_thumbnail'),
            created_at=get_val('created_at'),
            font_size=get_val('font_size'),
            message_id=get_val('id'),
            type=get_val('message_type'),
            reacted=get_val('reacted'),
            num_reactions=get_val('reactions_count'),
            room_id=get_val('room_id'),
            text=get_val('text'),
            author_id=get_val('user_id'),
            video_processed=get_val('video_processed'),
            video_thumbnail_big_url=get_val('video_thumbnail_big_url'),
            video_thumbnail_url=get_val('video_thumbnail_url'),
            video_url=get_val('video_url')
        )

        return message

    def activity_object(self, activity_data: dict) -> Post:
        def get_val(key):
            return activity_data.get(key, None)

        def get_val_2(key, sub_key):
            return activity_data.get(key, {}).get(sub_key, None)

        activity = Activity(
            created_at=get_val('created_at'),
            from_post_id=get_val_2('from_post', 'id'),
            from_group_id=get_val_2('group', 'id'),
            from_group_topic=get_val_2('group', 'topic'),
            metadata=get_val('metadata'),
            type=get_val('type'),
            from_user_id=get_val_2('user', 'id'),
            from_user_username=get_val_2('user', 'nickname'),
            from_user_profile_thumbnail=get_val_2(
                'user', 'from_user_profile_icon_thumbnail')
        )

        return activity
