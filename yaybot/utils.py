import huepy

from .exceptions import (
    YayError,
    AuthenticationError,
    ForbiddenError,
    RateLimitError,
    ExceedCallQuotaError,
    InvalidSignedInfoError,
    UnknownError
)

from .models import (
    User,
    GroupUser,
    Post,
    Review,
    Group,
    ChatRoom,
    Message,
    Activity
)


def handle_response(resp):
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
            raise InvalidSignedInfoError('Invalid signed info')


def console_print(text, color=None):
    text = '\n' + text
    if color is not None:
        text = getattr(huepy, color)(text)
    print(text)


class ObjectGenerator(object):

    def user_object(self, user_data: dict) -> User:
        def get_val(key):
            return user_data.get(key, None)

        user = User(
            id=get_val('id'),
            username=get_val('nickname'),
            bio=get_val('biography'),
            badge=user_data.get('title', None),
            num_followers=get_val('followers_count'),
            num_followings=get_val('followings_count'),
            private_user=get_val('is_private'),
            num_posts=get_val('posts_count'),
            num_joined_groups=get_val('groups_users_count'),
            num_reviews=get_val('reviews_count'),
            verified_age=get_val('age_verified'),
            country_code=get_val('country_code'),
            is_vip=get_val('vip'),
            hide_vip=get_val('hide_vip'),
            online_status=get_val('online_status'),
            prefecture=get_val('prefecture'),
            gender=get_val('gender'),
            generation=get_val('generation'),
            created_at=get_val('created_at'),
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

    def group_user_object(self, user_data: dict) -> GroupUser:
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
            author_id=post_data['user'].get('id', None),
            author_username=post_data['user'].get('nickname', None),
            text=get_val('text'),
            group_id=get_val('group_id'),
            font_size=get_val('font_size'),
            liked=get_val('liked'),
            num_likes=get_val('likes_count'),
            type=get_val('post_type'),
            color=get_val('color'),
            num_reposts=get_val('reposts_count'),
            created_at=get_val('created_at'),
            updated_at=get_val('updated_at'),
            edited_at=get_val('edited_at'),
            num_reported=get_val('reported_count'),
            reply_to_id=get_val('in_reply_to'),
            num_reply_to=get_val(
                'in_reply_to_post_count'),
            repostable=get_val('repostable'),
            highlighted=get_val('highlighted'),
            hidden=get_val('hidden'),
            thread_id=get_val('thread_id'),
            message_tags=get_val('message_tags'),
            tag_type=post_data['message_tags'][0].get(
                'type') if post_data.get('message_tags') else None,
            mentioned_user_id=post_data['message_tags'][0].get(
                'user_id') if post_data.get('message_tags') else None,
            conversation_id=get_val('conversation_id'),
            attachment=get_val('attachment'),
            attachment_thumbnail=get_val('attachment_thumbnail'),
            shared_url=post_data.get('shared_url', {}).get('url', None)
        )
        return post

    def review_object(self, review_data: dict) -> Review:
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

    def chat_room_object(self, chat_room_data: dict) -> ChatRoom:
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

    def message_object(self, message_data: dict) -> Message:
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

    def activity_object(self, activity_data: dict) -> Activity:
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
            from_username=get_val_2('user', 'nickname'),
            from_user_profile_thumbnail=get_val_2(
                'user', 'from_user_profile_icon_thumbnail')
        )

        return activity
