class Activity:

    def __init__(
        self,
        id: int,
        created_at: int,
        type: str,
        user: dict,
        from_post: dict,
        to_post: dict,
        group: dict,
        followers: list,
        from_post_ids: list,
        vip_reward: int,
        metadata: dict,
    ):
        self.id = id
        self.created_at = created_at
        self.type = type
        self.user = user
        self.from_post = from_post
        self.to_post = to_post
        self.group = group
        self.followers = followers
        self.from_post_ids = from_post_ids
        self.vip_reward = vip_reward
        self.metadata = metadata

    def __repr__(self):
        return f'Activity(id={self.id})'


class Footprint:

    def __init__(
        self,
        user: dict,
        visited_at: int,
        id: int,
    ):

        self.user = user
        self.visited_at = visited_at
        self.id = id

    def __repr__(self):
        return f'Footprint(id={self.id})'


class Group:

    def __init__(
        self,
        id: int,
        topic: str,
        description: str,
        user_id: int,
        groups_users_count: int,
        posts_count: int,
        threads_count: int,
        highlighted_count: int,
        views_count: int,
        related_count: int,
        secret: bool,
        gender: int,
        hide_reported_posts: bool,
        hide_conference_call: bool,
        is_private: bool,
        only_verified_age: bool,
        only_mobile_verified: bool,
        call_timeline_display: bool,
        updated_at: int,
        cover_image: str,
        cover_image_thumbnail: str,
        generation_groups_limit: int,
        owner: dict,
        is_joined: bool,
        is_pending: bool,
        group_category_id: int,
        unread_counts: int,
        moderator_ids: dict,
        seizable: bool,
        seizable_before: int,
        pending_count: int,
        pending_transfer_id: int,
        pending_deputize_ids: dict,
        safe_mode: bool,
        is_related: bool,
        allow_ownership_transfer: bool,
        sub_category_id: int,
        title: str,
        allow_thread_creation_by: str,
        hide_from_game_eight: bool,
        walkthrough_requested: bool,
        unread_threads_count: int,
        allow_members_to_post_image_and_video: bool,
        allow_members_to_post_url: bool,
        guidelines: str,
        invited_to_join: bool,
    ):
        self.id = id
        self.topic = topic
        self.description = description
        self.user_id = user_id
        self.groups_users_count = groups_users_count
        self.posts_count = posts_count
        self.threads_count = threads_count
        self.highlighted_count = highlighted_count
        self.views_count = views_count
        self.related_count = related_count
        self.secret = secret
        self.gender = gender
        self.hide_reported_posts = hide_reported_posts
        self.hide_conference_call = hide_conference_call
        self.is_private = is_private
        self.only_verified_age = only_verified_age
        self.only_mobile_verified = only_mobile_verified
        self.call_timeline_display = call_timeline_display
        self.updated_at = updated_at
        self.cover_image = cover_image
        self.cover_image_thumbnail = cover_image_thumbnail
        self.generation_groups_limit = generation_groups_limit
        self.owner = owner
        self.is_joined = is_joined
        self.is_pending = is_pending
        self.group_category_id = group_category_id
        self.unread_counts = unread_counts
        self.moderator_ids = moderator_ids
        self.seizable = seizable
        self.seizable_before = seizable_before
        self.pending_count = pending_count
        self.pending_transfer_id = pending_transfer_id
        self.pending_deputize_ids = pending_deputize_ids
        self.safe_mode = safe_mode
        self.is_related = is_related
        self.allow_ownership_transfer = allow_ownership_transfer
        self.sub_category_id = sub_category_id
        self.title = title
        self.allow_thread_creation_by = allow_thread_creation_by
        self.hide_from_game_eight = hide_from_game_eight
        self.walkthrough_requested = walkthrough_requested
        self.unread_threads_count = unread_threads_count
        self.allow_members_to_post_image_and_video = allow_members_to_post_image_and_video
        self.allow_members_to_post_url = allow_members_to_post_url
        self.guidelines = guidelines
        self.invited_to_join = invited_to_join

    def __repr__(self):
        return f'Group(id={self.id})'


class Post:

    def __init__(
        self,
        id: int,
        text: str,
        post_type: str,
        group_id: int,
        font_size: int,
        color: int,
        likes_count: int,
        created_at: int,
        updated_at: int,
        edited_at: int,
        liked: bool,
        likers: list,
        tag: str,
        reposts_count: int,
        reposted: int,
        repostable: bool,
        reported_count: int,
        conversation_id: int,
        in_reply_to: int,
        in_reply_to_post: dict,
        in_reply_to_post_count: int,
        user: dict,
        mentions: list,
        group: dict,
        conference_call: dict,
        attachment: str,
        attachment_thumbnail: str,
        attachment_2: str,
        attachment_2_thumbnail: str,
        attachment_3: str,
        attachment_3_thumbnail: str,
        attachment_4: str,
        attachment_4_thumbnail: str,
        attachment_5: str,
        attachment_5_thumbnail: str,
        attachment_6: str,
        attachment_6_thumbnail: str,
        attachment_7: str,
        attachment_7_thumbnail: str,
        attachment_8: str,
        attachment_8_thumbnail: str,
        attachment_9: str,
        attachment_9_thumbnail: str,
        shareable: dict,
        shared_url: dict,
        survey: dict,
        videos: dict,
        gifts_count: list,
        shared_thread: dict,
        thread_id: int,
        thread: dict,
        highlighted: bool,
        message_tags: list,
        is_fail_to_send: bool,
    ):
        self.id = id
        self.text = text
        self.post_type = post_type
        self.group_id = group_id
        self.font_size = font_size
        self.color = color
        self.likes_count = likes_count
        self.created_at = created_at
        self.updated_at = updated_at
        self.edited_at = edited_at
        self.liked = liked
        self.likers = likers
        self.tag = tag
        self.reposts_count = reposts_count
        self.reposted = reposted
        self.repostable = repostable
        self.reported_count = reported_count
        self.conversation_id = conversation_id
        self.in_reply_to = in_reply_to
        self.in_reply_to_post = in_reply_to_post
        self.in_reply_to_post_count = in_reply_to_post_count
        self.user = user
        self.mentions = mentions
        self.group = group
        self.conference_call = conference_call
        self.attachment = attachment
        self.attachment_thumbnail = attachment_thumbnail
        self.attachment_2 = attachment_2
        self.attachment_2_thumbnail = attachment_2_thumbnail
        self.attachment_3 = attachment_3
        self.attachment_3_thumbnail = attachment_3_thumbnail
        self.attachment_4 = attachment_4
        self.attachment_4_thumbnail = attachment_4_thumbnail
        self.attachment_5 = attachment_5
        self.attachment_5_thumbnail = attachment_5_thumbnail
        self.attachment_6 = attachment_6
        self.attachment_6_thumbnail = attachment_6_thumbnail
        self.attachment_7 = attachment_7
        self.attachment_7_thumbnail = attachment_7_thumbnail
        self.attachment_8 = attachment_8
        self.attachment_8_thumbnail = attachment_8_thumbnail
        self.attachment_9 = attachment_9
        self.attachment_9_thumbnail = attachment_9_thumbnail
        self.shareable = shareable
        self.shared_url = shared_url
        self.survey = survey
        self.videos = videos
        self.gifts_count = gifts_count
        self.shared_thread = shared_thread
        self.thread_id = thread_id
        self.thread = thread
        self.highlighted = highlighted
        self.message_tags = message_tags
        self.is_fail_to_send = is_fail_to_send

    def __repr__(self):
        return f'Post(id={self.id})'


class PostTag:

    def __init__(self, id: int, tag: str, post_hashtags_count: int):
        self.id = id
        self.tag = tag
        self.post_hashtags_count = post_hashtags_count

    def __repr__(self):
        return f'PostTag(id={self.id})'


class Review:

    def __init__(
        self,
        id: int,
        comment: str,
        reportedCount: int,
        user: dict,
        reviewer: dict,
        createdAt: int,
        mutualReview: bool,
    ):
        self.id = id
        self.comment = comment
        self.reportedCount = reportedCount
        self.user = user
        self.reviewer = reviewer
        self.createdAt = createdAt
        self.mutualReview = mutualReview

    def __repr__(self):
        return f'Review(id={self.id})'


class Survey:

    def __init__(self, id: int, votes_count: int, choices: dict, voted: bool):
        self.id = id
        self.votes_count = votes_count
        self.choices = choices
        self.voted = voted

    def __repr__(self):
        return f'Survey(id={self.id})'


class User:

    def __init__(
        self,
        user_id: int,
        biography: str,
        birth_date: str,
        cover_image: str,
        gender: str,
        cover_image_thumbnail: str,
        followers_count: int,
        followings_count: int,
        groups_users_count: int,
        is_age_verified: bool,
        is_chat_request_on: bool,
        is_dangerous: bool,
        is_email_confirmed: bool,
        is_facebook_connected: bool,
        is_follow_pending: bool,
        is_followed_by: bool,
        is_following: bool,
        is_from_different_generation_and_trusted: bool,
        is_group_phone_on: bool,
        is_hidden: bool,
        is_hide_vip: bool,
        is_interests_selected: bool,
        is_line_connected: bool,
        is_lobi_connected: bool,
        is_new: bool,
        is_phone_on: bool,
        is_private: bool,
        is_recently_penalized: bool,
        is_registered: bool,
        is_video_pn: bool,
        is_vip: bool,
        masked_email: str,
        mobile_number: str,
        nickname: str,
        posts_count: int,
        prefecture: str,
        profile_icon: str,
        profile_icon_thumbnail: str,
        reviews_count: int,
    ):
        self.id = user_id
        self.biography = biography
        self.birth_date = birth_date
        self.cover_image = cover_image
        self.gender = gender
        self.cover_image_thumbnail = cover_image_thumbnail
        self.followers_count = followers_count
        self.followings_count = followings_count
        self.groups_users_count = groups_users_count
        self.is_age_verified = is_age_verified
        self.is_chat_request_on = is_chat_request_on
        self.is_dangerous = is_dangerous
        self.is_email_confirmed = is_email_confirmed
        self.is_facebook_connected = is_facebook_connected
        self.is_follow_pending = is_follow_pending
        self.is_followed_by = is_followed_by
        self.is_following = is_following
        self.is_from_different_generation_and_trusted = is_from_different_generation_and_trusted
        self.is_group_phone_on = is_group_phone_on
        self.is_hidden = is_hidden
        self.is_hide_vip = is_hide_vip
        self.is_interests_selected = is_interests_selected
        self.is_line_connected = is_line_connected
        self.is_lobi_connected = is_lobi_connected
        self.is_new = is_new
        self.is_phone_on = is_phone_on
        self.is_private = is_private
        self.is_recently_penalized = is_recently_penalized
        self.is_registered = is_registered
        self.is_video_pn = is_video_pn
        self.is_vip = is_vip
        self.masked_email = masked_email
        self.mobile_number = mobile_number
        self.nickname = nickname
        self.posts_count = posts_count
        self.prefecture = prefecture
        self.profile_icon = profile_icon
        self.profile_icon_thumbnail = profile_icon_thumbnail
        self.reviews_count = reviews_count

    def __repr__(self):
        return f'User(username={self.nickname})'


class Video:

    def __init__(
        self,
        id: int,
        completed: bool,
        width: int,
        height: int,
        bitrate: int,
        views_count: int,
        video_url: str,
        thumbnail_url: str,
        thumbnail_big_url: str,
    ):
        self.id = id
        self.completed = completed
        self.width = width
        self.height = height
        self.bitrate = bitrate
        self.views_count = views_count
        self.video_url = video_url
        self.thumbnail_url = thumbnail_url
        self.thumbnail_big_url = thumbnail_big_url

    def __repr__(self):
        return f'Video(id={self.id})'
