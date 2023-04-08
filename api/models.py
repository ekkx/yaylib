class User:
    def __init__(self, id, display_name, biography, followers_count,
                 is_private, posts_count, joined_groups_count, reviews_count,
                 age_verified, country_code, is_vip, hide_vip, online_status,
                 profile_icon, profile_icon_thumbnail, cover_image,
                 cover_image_thumbnail, last_loggedin_at, mutual_chat,
                 chat_request, chat_required_phone_verification,
                 age_restricted_on_review, following_restricted_on_review,
                 restricted_review_by, is_recently_banned, is_dangerous_user,
                 is_new_user, interests_selected):
        self.id = id
        self.display_name = display_name
        self.biography = biography
        self.followers_count = followers_count
        self.is_private = is_private
        self.posts_count = posts_count
        self.joined_groups_count = joined_groups_count
        self.reviews_count = reviews_count
        self.age_verified = age_verified
        self.country_code = country_code
        self.is_vip = is_vip
        self.hide_vip = hide_vip
        self.online_status = online_status
        self.profile_icon = profile_icon
        self.profile_icon_thumbnail = profile_icon_thumbnail
        self.cover_image = cover_image
        self.cover_image_thumbnail = cover_image_thumbnail
        self.last_loggedin_at = last_loggedin_at
        self.mutual_chat = mutual_chat
        self.chat_request = chat_request
        self.chat_required_phone_verification = chat_required_phone_verification
        self.age_restricted_on_review = age_restricted_on_review
        self.following_restricted_on_review = following_restricted_on_review
        self.restricted_review_by = restricted_review_by
        self.is_recently_banned = is_recently_banned
        self.is_dangerous_user = is_dangerous_user
        self.is_new_user = is_new_user
        self.interests_selected = interests_selected

    def __repr__(self):
        return f"User(id={self.id}, display_name='{self.display_name}', biography='{self.biography}')"


class GroupUser:
    def __init__(
        self,
        is_moderator,
        banned,
        pending_transfer,
        pending_deputize,
        title,
        # 残りはUserオブジェクトと共通
    ):
        self.is_moderator = is_moderator
        self.banned = banned
        self.pending_transfer = pending_transfer
        self.pending_deputize = pending_deputize
        self.title = title
        # 残りはUserオブジェクトと共通

    def __repr__(self):
        return f"GroupUser(id={self.id}, display_name='{self.display_name}', biography='{self.biography}')"


class Post:
    def __init__(self, post_id, author_id, author_display_name, text, group_id,
                 font_size, is_liked, likes_count, post_type, color, reposts_count, created_at,
                 updated_at, edited_at, reported_count, in_reply_to_post, in_reply_to_post_count,
                 is_repostable, is_highlighted, is_hidden, thread_id, message_tags,
                 conversation_id, attachment, attachment_thumbnail, shared_url):
        self.post_id = post_id
        self.author_id = author_id
        self.author_display_name = author_display_name
        self.text = text
        self.group_id = group_id
        self.font_size = font_size
        self.is_liked = is_liked
        self.likes_count = likes_count
        self.post_type = post_type
        self.color = color
        self.reposts_count = reposts_count
        self.created_at = created_at
        self.updated_at = updated_at
        self.edited_at = edited_at
        self.reported_count = reported_count
        self.in_reply_to_post = in_reply_to_post
        self.in_reply_to_post_count = in_reply_to_post_count
        self.is_repostable = is_repostable
        self.is_highlighted = is_highlighted
        self.is_hidden = is_hidden
        self.thread_id = thread_id
        self.message_tags = message_tags
        self.conversation_id = conversation_id
        self.attachment = attachment
        self.attachment_thumbnail = attachment_thumbnail
        self.shared_url = shared_url

    def __repr__(self):
        return f"Post(post_id={self.post_id}, author_display_name='{self.author_display_name}', text='{self.text}')"


class Review:
    def __init__(self, text, created_at, id, mutual_review,
                 reported_count, author_id, author_display_name):
        self.text = text
        self.created_at = created_at
        self.id = id
        self.mutual_review = mutual_review
        self.reported_count = reported_count
        self.author_id = author_id
        self.author_display_name = author_display_name

    def __repr__(self):
        return f"Review(text={self.text}, created_at='{self.created_at}', id='{self.id}'), mutual_review='{self.mutual_review}'), reported_count='{self.reported_count}'), author_id='{self.author_id}'), author_display_name='{self.author_display_name}')"


class Group:
    def __init__(
        self,
        allow_members_to_post_image_and_video,
        allow_members_to_post_url,
        allow_ownership_transfer,
        allow_thread_creation_by,
        call_timeline_display,
        cover_image,
        cover_image_thumbnail,
        description,
        gender,
        generation_groups_limit,
        group_category_id,
        group_icon,
        groups_users_count,
        guidelines,
        hide_conference_call,
        hide_from_game_eight,
        hide_reported_posts,
        highlighted_count,
        homepage,
        group_id,
        invited_to_join,
        is_joined,
        is_pending,
        is_private,
        moderator_ids,
        only_mobile_verified,
        only_verified_age,
        owner_id,
        owner_display_name,
        pending_count,
        pending_deputize_ids,
        pending_transfer_id,
        place,
        posts_count,
        related_count,
        safe_mode,
        secret,
        seizable,
        seizable_before,
        sub_category_id,
        threads_count,
        title,
        topic,
        unread_threads_count,
        updated_at,
        views_count,
        walkthrough_requested
    ):
        self.allow_members_to_post_image_and_video = allow_members_to_post_image_and_video
        self.allow_members_to_post_url = allow_members_to_post_url
        self.allow_ownership_transfer = allow_ownership_transfer
        self.allow_thread_creation_by = allow_thread_creation_by
        self.call_timeline_display = call_timeline_display
        self.cover_image = cover_image
        self.cover_image_thumbnail = cover_image_thumbnail
        self.description = description
        self.gender = gender
        self.generation_groups_limit = generation_groups_limit
        self.group_category_id = group_category_id
        self.groups_users_count = groups_users_count
        self.guidelines = guidelines
        self.hide_conference_call = hide_conference_call
        self.hide_from_game_eight = hide_from_game_eight
        self.hide_reported_posts = hide_reported_posts
        self.highlighted_count = highlighted_count
        self.homepage = homepage
        self.group_icon = group_icon
        self.group_id = group_id
        self.invited_to_join = invited_to_join
        self.is_joined = is_joined
        self.is_pending = is_pending
        self.is_private = is_private
        self.moderator_ids = moderator_ids
        self.only_mobile_verified = only_mobile_verified
        self.only_verified_age = only_verified_age
        self.owner_id = owner_id
        self.owner_display_name = owner_display_name
        self.pending_count = pending_count
        self.pending_deputize_ids = pending_deputize_ids
        self.pending_transfer_id = pending_transfer_id
        self.place = place
        self.posts_count = posts_count
        self.related_count = related_count
        self.safe_mode = safe_mode
        self.secret = secret
        self.seizable = seizable
        self.seizable_before = seizable_before
        self.sub_category_id = sub_category_id
        self.threads_count = threads_count
        self.title = title
        self.topic = topic
        self.unread_threads_count = unread_threads_count
        self.updated_at = updated_at
        self.views_count = views_count
        self.walkthrough_requested = walkthrough_requested

    def __repr__(self):
        return f"Group(topic={self.topic}, description='{self.description}', owner_display_name='{self.owner_display_name}')"


class ChatRoom:
    def __init__(
            self,
            background,
            background_thumbnail,
            icon,
            icon_thumbnail,
            id,
            is_group,
            is_request,
            last_message,
            member_ids,
            member_display_names,
            name,
            unread_count,
            updated_at
    ):
        self.background = background
        self.background_thumbnail = background_thumbnail
        self.icon = icon
        self.icon_thumbnail = icon_thumbnail
        self.id = id
        self.is_group = is_group
        self.is_request = is_request
        self.last_message = last_message
        self.member_ids = member_ids
        self.member_display_names = member_display_names
        self.name = name
        self.unread_count = unread_count
        self.updated_at = updated_at

    def __repr__(self):
        return f"ChatRoom(id={self.id}, member_ids='{self.member_ids}', unread_count='{self.unread_count}', is_group='{self.is_group}')"


class Message:
    def __init__(
            self,
            attachment,
            attachment_android,
            attachment_thumbnail,
            created_at,
            font_size,
            id,
            message_type,
            reacted,
            reactions_count,
            room_id,
            text,
            user_id,
            video_processed,
            video_thumbnail_big_url,
            video_thumbnail_url,
            video_url
    ):
        self.attachment = attachment
        self.attachment_android = attachment_android
        self.attachment_thumbnail = attachment_thumbnail
        self.created_at = created_at
        self.font_size = font_size
        self.id = id
        self.message_type = message_type
        self.reacted = reacted
        self.reactions_count = reactions_count
        self.room_id = room_id
        self.text = text
        self.user_id = user_id
        self.video_processed = video_processed
        self.video_thumbnail_big_url = video_thumbnail_big_url
        self.video_thumbnail_url = video_thumbnail_url
        self.video_url = video_url

    def __repr__(self):
        return f"Message(user_id={self.user_id}, message_type='{self.message_type}', text='{self.text}', attachment='{self.attachment}')"


class Activity:
    def __init__(
            self,
            created_at,
            from_post_id,
            from_group_id,
            from_group_topic,
            metadata,
            type,
            from_user_id,
            from_user_display_name,
            from_user_profile_icon_thumbnail
    ):
        self.created_at = created_at
        self.from_post_id = from_post_id
        self.from_group_id = from_group_id
        self.from_group_topic = from_group_topic
        self.metadata = metadata
        self.type = type
        self.from_user_id = from_user_id
        self.from_user_display_name = from_user_display_name
        self.from_user_profile_icon_thumbnail = from_user_profile_icon_thumbnail

    def __repr__(self):
        return f"Activity(type={self.type}, from_post_id='{self.from_post_id}', from_group_id='{self.from_group_id}', from_user_id='{self.from_user_id}')"


class Like:
    def __init__(self, like_id, user, post):
        self.like_id = like_id
        self.user = user
        self.post = post

    def __repr__(self):
        return f"Like(like_id={self.like_id}, user={self.user}, post={self.post})"
