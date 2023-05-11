class User:

    def __init__(
        self,
        id,
        username,
        bio,
        badge,
        num_followers,
        num_followings,
        private_user,
        num_posts,
        num_joined_groups,
        num_reviews,
        verified_age,
        country_code,
        is_vip,
        hide_vip,
        online_status,
        prefecture,
        gender,
        generation,
        created_at,
        profile_image,
        profile_thumbnail,
        cover_image,
        cover_thumbnail,
        last_logged_in_at,
        mutual_chat_enabled,
        chat_request_enabled,
        chat_phone_verification_required,
        age_restricted_review,
        following_restricted_review,
        review_restricted_by,
        recently_banned,
        dangerous_user,
        new_user=None,
        selected_interests=None
    ):
        self.id = id
        self.username = username
        self.bio = bio
        self.badge = badge
        self.num_followers = num_followers
        self.num_followings = num_followings
        self.private_user = private_user
        self.num_posts = num_posts
        self.num_joined_groups = num_joined_groups
        self.num_reviews = num_reviews
        self.verified_age = verified_age
        self.country_code = country_code
        self.is_vip = is_vip
        self.hide_vip = hide_vip
        self.online_status = online_status
        self.prefecture = prefecture
        self.gender = gender
        self.generation = generation
        self.created_at = created_at
        self.profile_image = profile_image
        self.profile_thumbnail = profile_thumbnail
        self.cover_image = cover_image
        self.cover_thumbnail = cover_thumbnail
        self.last_logged_in_at = last_logged_in_at
        self.mutual_chat_enabled = mutual_chat_enabled
        self.chat_request_enabled = chat_request_enabled
        self.chat_phone_verification_required = chat_phone_verification_required
        self.age_restricted_review = age_restricted_review
        self.following_restricted_review = following_restricted_review
        self.review_restricted_by = review_restricted_by
        self.recently_banned = recently_banned
        self.dangerous_user = dangerous_user
        self.new_user = new_user
        self.selected_interests = selected_interests

    def __repr__(self):
        return f'User(username={self.username})'


class GroupUser(User):

    def __init__(
        self,
        moderator,
        banned,
        pending_transfer,
        pending_deputize,
        badge,
        id,
        username,
        bio,
        num_followers,
        is_private,
        num_posts,
        num_joined_groups,
        num_reviews,
        verified_age,
        country_code,
        is_vip,
        hide_vip,
        online_status,
        profile_image,
        profile_thumbnail,
        cover_image,
        cover_thumbnail,
        last_logged_in_at,
        mutual_chat_enabled,
        chat_request_enabled,
        chat_phone_verification_required,
        age_restricted_review,
        following_restricted_review,
        review_restricted_by,
        recently_banned,
        dangerous_user,
        new_user,
        selected_interests
    ):
        super().__init__(
            id,
            username,
            bio,
            num_followers,
            is_private,
            num_posts,
            num_joined_groups,
            num_reviews,
            verified_age,
            country_code,
            is_vip,
            hide_vip,
            online_status,
            profile_image,
            profile_thumbnail,
            cover_image,
            cover_thumbnail,
            last_logged_in_at,
            mutual_chat_enabled,
            chat_request_enabled,
            chat_phone_verification_required,
            age_restricted_review,
            following_restricted_review,
            review_restricted_by,
            recently_banned,
            dangerous_user,
            new_user,
            selected_interests
        )
        self.moderator = moderator
        self.banned = banned
        self.pending_transfer = pending_transfer
        self.pending_deputize = pending_deputize
        self.badge = badge

    def __repr__(self):
        return f''


class Post:

    def __init__(
        self,
        id,
        author_id,
        author_username,
        text,
        group_id,
        font_size,
        liked,
        num_likes,
        type,
        color,
        num_reposts,
        created_at,
        updated_at,
        edited_at,
        num_reported,
        reply_to_id,
        num_reply_to,
        repostable,
        highlighted,
        hidden,
        thread_id,
        message_tags,
        tag_type,
        mentioned_user_id,
        conversation_id,
        attachment,
        attachment_thumbnail,
        shared_url
    ):
        self.id = id
        self.author_id = author_id
        self.author_username = author_username
        self.text = text
        self.group_id = group_id
        self.font_size = font_size
        self.liked = liked
        self.num_likes = num_likes
        self.type = type
        self.color = color
        self.num_reposts = num_reposts
        self.created_at = created_at
        self.updated_at = updated_at
        self.edited_at = edited_at
        self.num_reported = num_reported
        self.reply_to_id = reply_to_id
        self.num_reply_to = num_reply_to
        self.repostable = repostable
        self.highlighted = highlighted
        self.hidden = hidden
        self.thread_id = thread_id
        self.message_tags = message_tags
        self.tag_type = tag_type
        self.mentioned_user_id = mentioned_user_id
        self.conversation_id = conversation_id
        self.attachment = attachment
        self.attachment_thumbnail = attachment_thumbnail
        self.shared_url = shared_url

    def __repr__(self):
        return str(self)


class Review:

    def __init__(self, text, created_at, id, mutual_review_enabled,
                 num_reported, author_id, author_username):
        self.text = text
        self.created_at = created_at
        self.id = id
        self.mutual_review_enabled = mutual_review_enabled
        self.num_reported = num_reported
        self.author_id = author_id
        self.author_username = author_username

    def __repr__(self):
        return f'Review(author_username={self.author_username}, text={self.text})'


class Group:

    def __init__(
        self,
        members_can_post_image_and_video,
        members_can_post_url,
        ownership_transfer_allowed,
        allowed_thread_creators,
        call_timeline,
        cover_image,
        cover_thumbnail,
        description,
        gender,
        generation_groups_limit,
        category_id,
        icon,
        num_groups_members,
        guidelines,
        conference_call_hidden,
        game_eight_hidden,
        reported_posts_hidden,
        num_highlighted,
        homepage,
        id,
        invited_to_join,
        joined,
        pending,
        private_group,
        moderator_ids,
        mobile_verified_only,
        verified_age_only,
        owner_id,
        owner_username,
        num_pendings,
        pending_deputize_ids,
        pending_transfer_id,
        place,
        num_posts,
        num_related_groups,
        safe_mode_enabled,
        hidden_grouop,
        seizable,
        seizable_before,
        sub_category_id,
        num_threads,
        category,
        group_name,
        num_unread_threads,
        updated_at,
        num_views,
        walkthrough_requested
    ):
        self.members_can_post_image_and_video = members_can_post_image_and_video
        self.members_can_post_url = members_can_post_url
        self.ownership_transfer_allowed = ownership_transfer_allowed
        self.allowed_thread_creators = allowed_thread_creators
        self.call_timeline = call_timeline
        self.cover_image = cover_image
        self.cover_thumbnail = cover_thumbnail
        self.description = description
        self.gender = gender
        self.generation_groups_limit = generation_groups_limit
        self.category_id = category_id
        self.num_groups_members = num_groups_members
        self.guidelines = guidelines
        self.conference_call_hidden = conference_call_hidden
        self.game_eight_hidden = game_eight_hidden
        self.reported_posts_hidden = reported_posts_hidden
        self.num_highlighted = num_highlighted
        self.homepage = homepage
        self.icon = icon
        self.id = id
        self.invited_to_join = invited_to_join
        self.joined = joined
        self.pending = pending
        self.private_group = private_group
        self.moderator_ids = moderator_ids
        self.mobile_verified_only = mobile_verified_only
        self.verified_age_only = verified_age_only
        self.owner_id = owner_id
        self.owner_username = owner_username
        self.num_pendings = num_pendings
        self.pending_deputize_ids = pending_deputize_ids
        self.pending_transfer_id = pending_transfer_id
        self.place = place
        self.num_posts = num_posts
        self.num_related_groups = num_related_groups
        self.safe_mode_enabled = safe_mode_enabled
        self.hidden_grouop = hidden_grouop
        self.seizable = seizable
        self.seizable_before = seizable_before
        self.sub_category_id = sub_category_id
        self.num_threads = num_threads
        self.category = category
        self.group_name = group_name
        self.num_unread_threads = num_unread_threads
        self.updated_at = updated_at
        self.num_views = num_views
        self.walkthrough_requested = walkthrough_requested

    def __repr__(self):
        return str(self)


class ChatRoom:

    def __init__(
            self,
            background_image,
            background_thumbnail,
            icon,
            icon_thumbnail,
            id,
            is_group,
            is_request,
            last_message,
            member_ids,
            member_usernames,
            chat_title,
            num_unread,
            updated_at
    ):
        self.background_image = background_image
        self.background_thumbnail = background_thumbnail
        self.icon = icon
        self.icon_thumbnail = icon_thumbnail
        self.id = id
        self.is_group = is_group
        self.is_request = is_request
        self.last_message = last_message
        self.member_ids = member_ids
        self.member_usernames = member_usernames
        self.chat_title = chat_title
        self.num_unread = num_unread
        self.updated_at = updated_at

    def __repr__(self):
        return str(self)


class Message:

    def __init__(
            self,
            attachment,
            attachment_android,
            attachment_thumbnail,
            created_at,
            font_size,
            message_id,
            type,
            reacted,
            num_reactions,
            room_id,
            text,
            author_id,
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
        self.message_id = message_id
        self.type = type
        self.reacted = reacted
        self.num_reactions = num_reactions
        self.room_id = room_id
        self.text = text
        self.author_id = author_id
        self.video_processed = video_processed
        self.video_thumbnail_big_url = video_thumbnail_big_url
        self.video_thumbnail_url = video_thumbnail_url
        self.video_url = video_url

    def __repr__(self):
        return str(self)


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
            from_username,
            from_user_profile_thumbnail
    ):
        self.created_at = created_at
        self.from_post_id = from_post_id
        self.from_group_id = from_group_id
        self.from_group_topic = from_group_topic
        self.metadata = metadata
        self.type = type
        self.from_user_id = from_user_id
        self.from_username = from_username
        self.from_user_profile_thumbnail = from_user_profile_thumbnail

    def __repr__(self):
        return str(self)
