import requests
import re
import time


from tqdm import tqdm


from auth import YayAuth
from endpoints import Endpoints as ep
from handle_exceptions import *
from models import *


class Yay(object):

    def __init__(self, token=None, proxy=None, timeout=10):
        self.timeout = timeout
        self.auth = YayAuth(proxy=proxy, timeout=timeout)

        if token:
            self.auth.access_token = token
            self.auth.headers.setdefault('Authorization', f'Bearer {token}')

    def login(self, email, password):
        self.auth.login(email, password)
        self.logged_in_as = self.auth.logged_in_as
        self.access_token = self.auth.access_token
        self.api_key = self.auth.api_key
        self.logged_in = True

    def logout(self):
        self.auth.logout()
        self.logged_in_as = None
        self.access_token = None
        self.api_key = None

    def create_user_object(self, post_data: dict) -> User:
        def get_val(key):
            return post_data.get(key, None)

        user = User(
            id=get_val('id'),
            display_name=get_val('nickname'),
            biography=get_val('biography'),
            followers_count=get_val('followers_count'),
            is_private=get_val('is_private'),
            posts_count=get_val('posts_count'),
            joined_groups_count=get_val('groups_users_count'),
            reviews_count=get_val('reviews_count'),
            age_verified=get_val('age_verified'),
            country_code=get_val('country_code'),
            is_vip=get_val('vip'),
            hide_vip=get_val('hide_vip'),
            online_status=get_val('online_status'),
            profile_icon=get_val('profile_icon'),
            profile_icon_thumbnail=get_val(
                'profile_icon_thumbnail'),
            cover_image=get_val('cover_image'),
            cover_image_thumbnail=get_val('cover_image_thumbnail'),
            last_loggedin_at=get_val('last_loggedin_at'),
            mutual_chat=get_val('mutual_chat'),
            chat_request=get_val('chat_request'),
            chat_required_phone_verification=get_val(
                'chat_required_phone_verification'),
            age_restricted_on_review=get_val(
                'age_restricted_on_review'),
            following_restricted_on_review=get_val(
                'following_restricted_on_review'),
            restricted_review_by=get_val('restricted_review_by'),
            is_recently_banned=get_val('recently_kenta'),
            is_dangerous_user=get_val('dangerous_user'),
            is_new_user=get_val('new_user'),
            interests_selected=get_val('interests_selected')
        )
        return user

    def create_post_object(self, post_data: dict) -> Post:
        def get_val(key):
            return post_data.get(key, None)

        post = Post(
            post_id=get_val('id'),
            author_id=post_data['user'].get('id'),
            author_display_name=post_data['user'].get('nickname', None),
            text=get_val('text'),
            group_id=get_val('group_id'),
            font_size=get_val('font_size'),
            is_liked=get_val('liked'),
            likes_count=get_val('likes_count'),
            post_type=get_val('post_type'),
            color=get_val('color'),
            reposts_count=get_val('reposts_count'),
            created_at=get_val('created_at'),
            updated_at=get_val('updated_at'),
            edited_at=get_val('edited_at'),
            reported_count=get_val('reported_count'),
            in_reply_to_post=get_val('in_reply_to_post'),
            in_reply_to_post_count=get_val(
                'in_reply_to_post_count'),
            is_repostable=get_val('repostable'),
            is_highlighted=get_val('highlighted'),
            is_hidden=get_val('hidden'),
            thread_id=get_val('thread_id'),
            message_tags=get_val('message_tags'),
            conversation_id=get_val('conversation_id'),
            attachment=get_val('attachment'),
            attachment_thumbnail=get_val('attachment_thumbnail'),
            shared_url=post_data.get('shared_url', {}).get('url', None)
        )
        return post

    def create_review_object(self, review_data: dict) -> Post:
        def get_val(key):
            return review_data.get(key, None)

        review = Review(
            text=get_val('comment'),
            created_at=get_val('created_at'),
            id=get_val('id'),
            mutual_review=get_val('mutual_review'),
            reported_count=get_val('reported_count'),
            author_id=review_data['reviewer'].get('id', None),
            author_display_name=review_data['reviewer'].get('nickname', None),
        )

        return review

    def create_group_object(self, group_data: dict) -> Group:
        def get_val(key):
            return group_data.get(key, None)

        group = Group(
            allow_members_to_post_image_and_video=get_val(
                'allow_members_to_post_image_and_video'),
            allow_members_to_post_url=get_val('allow_members_to_post_url'),
            allow_ownership_transfer=get_val('allow_ownership_transfer'),
            allow_thread_creation_by=get_val('allow_thread_creation_by'),
            call_timeline_display=get_val('call_timeline_display'),
            cover_image=get_val('cover_image'),
            cover_image_thumbnail=get_val('cover_image_thumbnail'),
            description=get_val('description'),
            gender=get_val('gender'),
            generation_groups_limit=get_val('generation_groups_limit'),
            group_category_id=get_val('group_category_id'),
            group_icon=get_val('group_icon'),
            groups_users_count=get_val('groups_users_count'),
            guidelines=get_val('guidelines'),
            hide_conference_call=get_val('hide_conference_call'),
            hide_from_game_eight=get_val('hide_from_game_eight'),
            hide_reported_posts=get_val('hide_reported_posts'),
            highlighted_count=get_val('highlighted_count'),
            homepage=get_val('homepage'),
            group_id=get_val('group_id'),
            invited_to_join=get_val('invited_to_join'),
            is_joined=get_val('is_joined'),
            is_pending=get_val('is_pending'),
            is_private=get_val('is_private'),
            moderator_ids=get_val('moderator_ids'),
            only_mobile_verified=get_val('only_mobile_verified'),
            only_verified_age=get_val('only_verified_age'),
            owner_id=get_val('user_id'),
            owner_display_name=group_data['owner'].get('nickname', None),
            pending_count=get_val('pending_count'),
            pending_deputize_ids=get_val('pending_deputize_ids'),
            pending_transfer_id=get_val('pending_transfer_id'),
            place=get_val('place'),
            posts_count=get_val('posts_count'),
            related_count=get_val('related_count'),
            safe_mode=get_val('safe_mode'),
            secret=get_val('secret'),
            seizable=get_val('seizable'),
            seizable_before=get_val('seizable_before'),
            sub_category_id=get_val('sub_category_id'),
            threads_count=get_val('threads_count'),
            title=get_val('title'),
            topic=get_val('topic'),
            unread_threads_count=get_val('unread_threads_count'),
            updated_at=get_val('updated_at'),
            views_count=get_val('views_count'),
            walkthrough_requested=get_val('walkthrough_requested'),
        )

        return group

    def create_chat_room_object(self, chat_room_data: dict) -> Post:
        def get_val(key):
            return chat_room_data.get(key, None)

        chat_room = ChatRoom(
            background=get_val('background'),
            background_thumbnail=get_val('background_thumbnail'),
            icon=get_val('icon'),
            icon_thumbnail=get_val('icon_thumbnail'),
            id=get_val('id'),
            is_group=get_val('is_group'),
            is_request=get_val('is_request'),
            last_message=chat_room_data['last_message'].get('text', None),
            member_ids=[member.get('id', None)
                        for member in chat_room_data['members']],
            member_display_names=[member.get(
                'nickname', None) for member in chat_room_data['members']],
            name=get_val('name'),
            unread_count=get_val('unread_count'),
            updated_at=get_val('updated_at')
        )

        return chat_room

    def create_message_object(self, message_data: dict) -> Post:
        def get_val(key):
            return message_data.get(key, None)

        message = Message(
            attachment=get_val('attachment'),
            attachment_android=get_val('attachment_android'),
            attachment_thumbnail=get_val('attachment_thumbnail'),
            created_at=get_val('created_at'),
            font_size=get_val('font_size'),
            id=get_val('id'),
            message_type=get_val('message_type'),
            reacted=get_val('reacted'),
            reactions_count=get_val('reactions_count'),
            room_id=get_val('room_id'),
            text=get_val('text'),
            user_id=get_val('user_id'),
            video_processed=get_val('video_processed'),
            video_thumbnail_big_url=get_val('video_thumbnail_big_url'),
            video_thumbnail_url=get_val('video_thumbnail_url'),
            video_url=get_val('video_url')
        )

        return message

    def create_activity_object(self, activity_data: dict) -> Post:
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
            from_user_display_name=get_val_2('user', 'nickname'),
            from_user_profile_icon_thumbnail=get_val_2(
                'user', 'from_user_profile_icon_thumbnail')
        )

        return activity

    def _get(self, url: str):
        response = requests.get(url, headers=self.auth.headers,
                                proxies={'http': self.auth.proxy,
                                         'https': self.auth.proxy},
                                timeout=self.timeout)
        self._handle_response(response)
        return response.json()

    def _post(self, url: str, data: dict = None):
        response = requests.post(url, params=data,
                                 headers=self.auth.headers,
                                 proxies={'http': self.auth.proxy,
                                          'https': self.auth.proxy},
                                 timeout=self.timeout)
        self._handle_response(response)
        return response.json()

    def _put(self, url: str, data: dict = None):
        response = requests.put(url, params=data,
                                headers=self.auth.headers,
                                proxies={'http': self.auth.proxy,
                                         'https': self.auth.proxy},
                                timeout=self.timeout)
        self._handle_response(response)
        return response.json()

    def _delete(self, url: str, data: dict = None):
        response = requests.delete(url, params=data,
                                   headers=self.auth.headers,
                                   proxies={'http': self.auth.proxy,
                                            'https': self.auth.proxy},
                                   timeout=self.timeout)
        self._handle_response(response)
        return response.json()

    def _handle_response(self, response):
        if response.status_code == 401:
            raise AuthenticationError('Failed to authenticate')
        if response.status_code == 429:
            raise RateLimitError('Rate limit exceeded')

    # ----- GET USER -----

    def get_user(self, user_id: str):
        response = self._get(f'{ep.USER_v2}/{user_id}')
        user_data = response.get('user')
        user = self.create_user_object(user_data)
        return user

    def get_users_from_dict(self, response: dict):
        assert 'users' in response, "'users' key not found"
        users_data = response.get('users')
        users = []
        for user_data in users_data:
            user = self.create_user_object(user_data)
            users.append(user)
        return users

    def get_letters(self, user_id, amount=100):
        # あとで100以上のレターを表示できるようにする
        # {ep.USER_v1}/reviews/{user_id}?from_id=111&not_active=false&number=100
        response = self._get(
            f'{ep.USER_v1}/reviews/{user_id}?not_active=false&number={amount}')
        reviews_data = response['reviews']
        reviews = []
        for review_data in reviews_data:
            review = self.create_review_object(review_data)
            reviews.append(review)
        return reviews

    def get_joined_groups(self, user_id, amount=100):
        response = self._get(
            f'{ep.GROUP_v1}/user_group_list?number={amount}&page=0&user_id={user_id}')
        return self.get_groups_from_dict(response)

    def get_user_followers(self, user_id, amount: int = None):
        response = self._get(
            f'{ep.USER_v2}/{user_id}/web_followers?number=50')
        return self.get_users_from_dict(response)

    def get_followings(self, user_id, amount=50):
        # need fix
        number = amount if amount < 50 else 50
        located = 0

        response = self._get(
            f'{ep.USER_v2}/{user_id}/web_followings?number={number}')
        last_value = response.get('last_follow_id')
        users = self.get_users_from_dict(response)
        located += number

        while located < amount:
            response = self._get(
                f'{ep.USER_v2}/{user_id}/web_followings?from_follow_id={last_value}&number=50')
            last_value = response.get('last_follow_id')
            users.append(self.get_users_from_dict(response))
            located += number

        return users

    def get_follow_requests(self, amount=100):
        # has last_timestamp: 1658812381
        response = self._get(
            f'{ep.USER_v2}/follow_requests?number=50')
        return self.get_users_from_dict(response)

    def get_user_active_call(self, user_id):
        response = self._get(
            f'{ep.POST_v1}/active_call?user_id={user_id}')
        post_id = response['post'].get('id')
        return self.get_post(post_id)

    def get_blocked_users(self, amount=100):
        response = self._get(
            f'{ep.USER_v2}/blocked')
        return self.get_users_from_dict(response)

    # ----- GET POST -----

    def get_post(self, post_id: str):
        response = self._get(f'{ep.POST_v2}/{post_id}')
        post_data = response.get('post')
        post = self.create_post_object(post_data)
        return post

    def get_posts_from_dict(self, response: dict):
        assert 'posts' in response, "'posts' key not found"
        posts_data = response.get('posts')
        posts = []
        for post_data in posts_data:
            post = self.create_post_object(post_data)
            posts.append(post)
        return posts

    def get_timeline(self, user_id=None, keyword=None, hashtag=None, amount=100):
        if user_id:
            response = self._get(
                f'{ep.GET_USER_TIMELINE}?number={amount}&user_id={user_id}')
        elif keyword:
            response = self._get(
                f'{ep.GET_TIMELINE_BY_KEYWORD}?keyword={keyword}&number={amount}')
        elif hashtag:
            response = self._get(
                f'{ep.GET_TIMELINE_BY_HASHTAG}/{hashtag}?number={amount}')
        else:
            response = self._get(
                f'{ep.GET_TIMELINE}?number={amount}')
        return self.get_posts_from_dict(response)

    def get_following_timeline(self, amount=50):
        if amount > 50:
            pass
        # has next_page_value
        response = self._get(
            f'{ep.GET_FOLLOWING_TIMELINE}?number=50')
        return self.get_posts_from_dict(response)

    def get_conversation(self, conversation_id, post_id=None, amount=100):
        if post_id:
            conversation_id = self.get_post(post_id).conversation_id
        response = self._get(
            f'{ep.GET_CONVERSATION}/{conversation_id}?number={amount}&reverse=true')
        return self.get_posts_from_dict(response)

    def get_reposts(self, post_id, amount=100):
        response = self._get(
            f'{ep.POST_v2}/{post_id}/reposts?number=100')
        return self.get_posts_from_dict(response)

    def get_post_likers(self, post_id, amount=50):
        # has last_id
        response = self._get(
            f'{ep.POST_v1}/{post_id}/likers')
        return self.get_users_from_dict(response)

    # ----- GET GROUP -----

    def get_group(self, group_id: str):
        response = self._get(f'{ep.GROUP_v1}/{group_id}')
        group_data = response.get('group')
        group = self.create_group_object(group_data)
        return group

    def get_groups_from_dict(self, response: dict):
        assert 'groups' in response, "'groups' key not found"
        groups_data = response.get('groups')
        groups = []
        for group_data in groups_data:
            group = self.create_group_object(group_data)
            groups.append(group)
        return groups

    def get_group_timeline(self, group_id, amount=100):
        # {ep.POST_v2}/group_timeline?from_post_id=111&group_id={group_id}&number=100
        response = self._get(
            f'{ep.POST_v2}/group_timeline?group_id={group_id}&number=100')
        return self.get_posts_from_dict(response)

    def get_group_call(self, group_id):
        response = self._get(
            f'{ep.POST_v2}/call_timeline?group_id={group_id}&number=20')
        return self.get_posts_from_dict(response)

    def get_group_members(self, group_id, amount=100):
        # need a fix. the model is different
        response = self._get(
            f'{ep.GROUP_v2}/{group_id}/members?number=100')
        response = {'users': response['group_users']}
        # response = response['group_users']
        print(response)
        return self.get_users_from_dict(response)

    def get_pending_users_in_group(self, group_id, amount=100):
        response = self._get(
            f'{ep.GROUP_v2}/{group_id}/{group_id}/members?mode=pending&number=100')
        return self.get_users_from_dict(response)

    def get_banned_user_from_group(self, group_id, amount=100):
        response = self._get(
            f'{ep.GROUP_v1}/{group_id}/{group_id}/ban_list?number=100')
        return self.get_users_from_dict(response)

    # ----- GET CHAT -----

    def get_chat_room(self, chatroom_id: str):
        response = self._get(f'{ep.CHATROOM_v2}/{chatroom_id}')
        chat_room_data = response.get('chat')
        chat = self.create_chat_room_object(chat_room_data)
        return chat

    def get_chat_rooms_from_dict(self, response: dict):
        assert 'chat_rooms' in response, "'chat_rooms' key not found"
        chat_rooms_data = response.get('chat_rooms')
        chats = []
        for chat_room_data in chat_rooms_data:
            chat = self.create_chat_room_object(chat_room_data)
            chats.append(chat)
        return chats

    def get_chat_messages(self, chatroom_id=None, user_id=None, amount=None):
        response = self._get(
            f'{ep.CHATROOM_v2}/{chatroom_id}/messages?number=100')
        return self.get_chat_messages_from_dict(response)

    def get_chat_messages_from_dict(self, response: dict):
        assert 'messages' in response, "'messages' key not found"
        messages_data = response.get('messages')
        messages = []
        for message_data in messages_data:
            message = self.create_message_object(message_data)
            messages.append(message)
        return messages

    def get_chat_rooms(self, amount=None):
        response = self._get(
            f'{ep.CHATROOM_v1}/main_list?number=100')
        return self.get_chat_rooms_from_dict(response)

    def get_chat_requests(self, amount=None):
        response = self._get(
            f'{ep.CHATROOM_v1}/request_list?number=100')
        return self.get_chat_rooms_from_dict(response)

    # ----- GET NOTIFICATION -----

    def get_activity_from_dict(self, response: dict):
        assert 'activities' in response, "'activities' key not found"
        activities_data = response.get('activities')
        activities = []
        for activity_data in activities_data:
            activity = self.create_activity_object(activity_data)
            activities.append(activity)
        return activities

    def get_notification(self, important=True):
        if important:
            response = self._get(
                f'{ep.CAS_BASE_URL}/api/user_activities?important=true&number=100')
        else:
            response = self._get(
                f'{ep.CAS_BASE_URL}/api/user_activities?important=false&number=100')
        return self.get_activity_from_dict(response)

    # def get_footprint_from_dict(self, response: dict):
    #     # need a lot to think about
    #     return

    # def get_profile_visitors(self, amount=None):
    #     response = self._get(
    #         f'{ep.USER_v2}/footprints?number=100')
    #     return

    # def get_visited_users(self, amount=None):
    #     return

    # -----------------------

    # ----- ACTION USER -----

    # -----------------------

    def follow_user(self, user_id):
        response_data = self._post(f'{ep.USER_v2}/{user_id}/follow')
        return response_data

    def unfollow_user(self, user_id):
        response_data = self._post(f'{ep.USER_v2}/{user_id}/unfollow')
        return response_data

    def accept_follow_request(self, user_id):
        response_data = self._post(
            f'{ep.USER_v2}/{user_id}/follow_request?action=accept')
        return response_data

    def reject_follow_request(self, user_id):
        response_data = self._post(
            f'{ep.USER_v2}/{user_id}/follow_request?action=reject')
        return response_data

    def send_letter(self, user_id, message):
        data = {'comment': message}
        response_data = self._post(
            f'{ep.USER_v1}/reviews/{user_id}', data)
        return response_data

    def block_user(self, user_id):
        response_data = self._post(
            f'{ep.USER_v1}/{user_id}/block')
        return response_data

    def unblock_user(self, user_id):
        response_data = self._post(
            f'{ep.USER_v1}/{user_id}/unblock')
        return response_data

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
        response_data = self._post(
            f'{ep.BASE_URL}/v1/web/posts/new', data)
        return response_data

    def create_post_in_group(self, group_id, text, color=0, font_size=0):
        data = {
            'color': color,
            'font_size': font_size,
            'group_id': group_id,
            'text': text,
            'post_type': 'text',
            'uuid': ''
        }
        response_data = self._post('https://yay.space/api/posts', data)
        return self.get_post(response_data['id'])

    def create_repost(self, text, post_id, color=0, font_size=0):
        data = {
            'text': text,
            'color': color,
            'font_size': font_size,
            'post_id': post_id,
            'message_tags': '[]',
            'post_type': 'text'
        }
        response_data = self._post(
            f'{ep.POST_v3}/repost', data)
        return response_data

    def create_reply(self, text, post_id, color=0, font_size=0):
        data = {
            'text': text,
            'color': color,
            'font_size': font_size,
            'in_reply_to': post_id
        }
        response_data = self._post(
            f'{ep.BASE_URL}/v1/web/posts/new', data)
        return response_data

    def delete_post(self, post_id):
        data = {'posts_ids[]': post_id}
        response_data = self._post(
            f'{ep.POST_v2}/mass_destroy', data)
        return response_data

    def pin_post(self, post_id):
        data = {'id': post_id}
        response_data = self._post(
            f'{ep.PIN_v1}/posts', data)
        return response_data

    def unpin_post(self, post_id):
        response_data = self._post(
            f'{ep.PIN_v1}/posts/{post_id}')
        return response_data

    def like_post(self, post_id):
        data = {'post_ids': post_id}
        response_data = self._post(
            f'{ep.POST_v2}/like', data)
        return response_data

    def unlike_post(self, post_id):
        response_data = self._post(
            f'{ep.POST_v1}/{post_id}/unlike')
        return response_data

    # ----- ACTION GROUP -----

    def create_group(self, group_name: str, description: str = None, guidelines: str = None,
                     group_category_id=21, sub_category_id: int = None, is_private=False,
                     call_timeline_display=True, hide_reported_posts=False,
                     allow_ownership_transfer=True, allow_thread_creation_by='member',
                     allow_members_to_post_image_and_video=True, allow_members_to_post_url=True,
                     hide_conference_call=False, secret=False, only_verified_age=False,
                     only_mobile_verified=False, gender=-1, generation_groups_limit=0):
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
        response_data = self._post(
            f'{ep.GROUP_v1}/new', data)
        return response_data

    def delete_group(self, group_id):
        data = {'groupId': group_id}
        response_data = self._delete(
            f'{ep.GROUP_v1}/{group_id}/leave', data)
        return response_data

    def change_group_settings(self, group_id, group_name: str, description: str = None,
                              guidelines: str = None, group_category_id=21, sub_category_id: int = None,
                              is_private=False, call_timeline_display=True, hide_reported_posts=False,
                              allow_ownership_transfer=True, allow_thread_creation_by='member',
                              allow_members_to_post_image_and_video=True, allow_members_to_post_url=True,
                              hide_conference_call=False, hide_from_game_eight=False, secret=False, only_verified_age=False,
                              only_mobile_verified=False, gender=-1, generation_groups_limit=0):
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
        response = self._put(
            f'https://yay.space/api/groups/{group_id}', data)
        return self.create_group_object(response.get('group'))

    def transfer_group_ownership(self, group_id, user_id):
        # https://yay.space/api/groups/285796/transfer method=post
        # data     {"uuid":"5b6e9ca8-60aa-4a6e-8756-0034c3b7d2ab","user_id":1}
        data = {
            'uuid': '',
            'user_id': user_id
        }
        response_data = self._post(
            f'https://yay.space/api/groups/{group_id}/transfer', data)
        return response_data

    def undo_group_ownership_transfer(self, group_id, user_id):
        # {ep.BASE_URL}/v1/groups/285796/transfer/withdraw   method=put
        # data     {"user_id":1}
        return

    def offer_group_sub_owner(self, group_id, user_id):
        # need a fix json.decoder.JSONDecodeError: Expecting value: line 1 column 1 (char 0)
        data = {'uuid': '', 'user_ids': user_id}
        response = self._post(
            f'https://yay.space/api/groups/{group_id}/deputize', data)
        return response
        # https://yay.space/api/groups/285796/deputize   method=post
        # data     {"uuid":"aaa","user_ids":["1"]}

    def undo_group_sub_owner_offer(self, group_id, user_id):
        # {ep.BASE_URL}/v1/groups/{group_id}/deputize/{user_id}/withdraw method=put
        return

    def fire_group_sub_owner(self, group_id, user_id):
        # {ep.BASE_URL}/v1/groups/{group_id}/fire/{user_id} method=post
        return

    def accept_group_join_request(self, group_id, user_id):
        return

    def decline_group_join_request(self, group_id, user_id):
        return

    def invite_user_to_group(self, group_id, user_id):
        return

    def pin_group_post(self, group_id, post_id):
        data = {'group_id': group_id, 'post_id': post_id}
        response_data = self._put(
            f'{ep.POST_v2}/group_pinned_post', data)
        return response_data

    def unpin_group_post(self, group_id):
        data = {'group_id': group_id}
        response_data = self._delete(
            f'{ep.POST_v2}/group_pinned_post', data)
        return response_data

    def join_group(self, group_id):
        self._post(f'{ep.GROUP_v1}/{group_id}/join')
        return self.get_group(group_id)

    def leave_group(self, group_id):
        data = {'groupId': group_id}
        response_data = self._delete(
            f'{ep.GROUP_v1}/{group_id}/leave', data)
        return response_data

    # ----- ACTION CHAT -----

    def send_message(self, user_id):
        return

    def accept_chat_request(self, chat_room_id):
        # {ep.BASE_URL}/v1/chat_rooms/accept_chat_request method=post
        # data       chat_room_ids[]: 139328860
        return

    def delete_chat_room(self, chat_room_id):
        # {ep.BASE_URL}/v1/chat_rooms/mass_destroy method=post
        # data     {"chat_room_ids":[128089710]}
        return
