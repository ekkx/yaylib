from .utils import parse_datetime


class Activity:

    __slots__ = (
        "data", "id", "created_at", "type", "user", "from_post", "to_post",
        "group", "followers", "from_post_ids", "vip_reward", "metadata",
    )

    def __init__(self, data):
        self.data = data
        self.id = data.get("id")
        self.created_at = parse_datetime(data.get("created_at"))
        self.type = data.get("type")

        self.user = data.get("user")
        if self.user is not None:
            self.user = User(self.user)

        self.from_post = data.get("from_post")
        if self.from_post is not None:
            self.from_post = Post(self.from_post)

        self.to_post = data.get("to_post")
        if self.to_post is not None:
            self.to_post = Post(self.to_post)

        self.group = data.get("group")
        if self.group is not None:
            self.group = Group(self.group)

        self.followers = data.get("followers")
        if self.user is not None:
            self.user = [
                User(user) for user in self.user
            ]

        self.from_post_ids = data.get("from_post_ids")

        self.vip_reward = data.get("vip_reward")

        self.metadata = data.get("metadata")
        # if self.metadata is not None:
        #     self.metadata = Metadata(self.metadata)

    def __repr__(self):
        return f'Activity(id={self.id})'


class CallGiftHistory:

    __slots__ = ("data", "gifts_count", "sent_at", "sender")

    def __init__(self, data):
        self.data = data

        self.gifts_count = data.get("gifts_count")
        if self.gifts_count is not None:
            self.gifts_count = [
                GiftCount(gifts_count) for gifts_count in self.gifts_count
            ]

        self.sent_at = parse_datetime(data.get("sent_at"))

        self.sender = data.get("sender")
        if self.sender is not None:
            self.sender = User(self.sender)

    def __repr__(self):
        return f'CallGiftHistory(data={self.data}, gifts_count={self.gifts_count})'


class CoinAmount:

    __slots__ = ("data", "paid", "free", "total")

    def __init__(self, data):
        self.data = data
        self.paid = data.get("paid")
        self.free = data.get("free")
        self.total = data.get("total")

    def __repr__(self):
        return f'CoinAmount(data={self.data}, paid={self.paid})'


class CoinExpiration:

    __slots__ = ("data", "expired_at_seconds", "amount")

    def __init__(self, data):
        self.data = data
        self.expired_at_seconds = data.get("expired_at_seconds")
        self.amount = data.get("amount")

    def __repr__(self):
        return f'CoinExpiration(data={self.data}, expired_at_seconds={self.expired_at_seconds})'


class CoinProduct:

    __slots__ = ("data", "id", "purchasable", "amount")

    def __init__(self, data):
        self.data = data
        self.id = data.get("id")
        self.purchasable = data.get("purchasable")
        self.amount = data.get("amount")

    def __repr__(self):
        return f'CoinProduct(data={self.data}, id={self.id})'


class CoinProductQuota:

    __slots__ = ("data", "bought", "limit")

    def __init__(self, data):
        self.data = data
        self.bought = data.get("bought")
        self.limit = data.get("limit")

    def __repr__(self):
        return f'CoinProductQuota(data={self.data}, bought={self.bought})'


class CreateGroupQuota:

    __slots__ = ("data", "used_quota", "remaining_quota")

    def __init__(self, data):
        self.data = data
        self.used_quota = data.get("used_quota")
        self.remaining_quota = data.get("remaining_quota")

    def __repr__(self):
        return f'CreateGroupQuota(data={self.data}, used_quota={self.used_quota})'


class Footprint:

    __slots__ = ("data", "user", "visited_at", "id")

    def __init__(self, data):
        self.data = data

        self.user = data.get("user")
        if self.user is not None:
            self.user = User(self.user)

        self.visited_at = parse_datetime(data.get("visited_at"))
        self.id = data.get("id")

    def __repr__(self):
        return f'Footprint(data={self.data}, id={self.id})'


class GiftCount:

    __slots__ = ("data", "id", "quantity")

    def __init__(self, data):
        self.data = data
        self.id = data.get("id")
        self.quantity = data.get("quantity")

    def __repr__(self):
        return f'GiftCount(data={self.data}, id={self.id})'


class Group:

    __slots__ = (
        "data", "id", "topic", "description", "user_id", "groups_users_count",
        "posts_count", "threads_count", "highlighted_count", "views_count",
        "related_count", "secret", "gender", "hide_reported_posts",
        "hide_conference_call", "is_private", "only_verified_age",
        "only_mobile_verified", "call_timeline_display", "updated_at",
        "cover_image", "cover_image_thumbnail", "generation_groups_limit",
        "owner", "is_joined", "is_pending", "group_category_id",
        "unread_counts", "moderator_ids", "seizable", "seizable_before",
        "pending_count", "pending_transfer_id", "pending_deputize_ids",
        "safe_mode", "is_related", "allow_ownership_transfer",
        "sub_category_id", "title", "allow_thread_creation_by",
        "hide_from_game_eight", "walkthrough_requested",
        "unread_threads_count", "allow_members_to_post_image_and_video",
        "allow_members_to_post_url", "guidelines", "invited_to_join"
    )

    def __init__(self, data):
        self.data = data
        self.id = data.get("id")
        self.topic = data.get("topic")
        self.description = data.get("description")
        self.user_id = data.get("user_id")
        self.groups_users_count = data.get("groups_users_count")
        self.posts_count = data.get("posts_count")
        self.threads_count = data.get("threads_count")
        self.highlighted_count = data.get("highlighted_count")
        self.views_count = data.get("views_count")
        self.related_count = data.get("related_count")
        self.secret = data.get("secret")
        self.gender = data.get("gender")
        self.hide_reported_posts = data.get("hide_reported_posts")
        self.hide_conference_call = data.get("hide_conference_call")
        self.is_private = data.get("is_private")
        self.only_verified_age = data.get("only_verified_age")
        self.only_mobile_verified = data.get("only_mobile_verified")
        self.call_timeline_display = data.get("call_timeline_display")
        self.updated_at = parse_datetime(data.get("updated_at"))
        self.cover_image = data.get("cover_image")
        self.cover_image_thumbnail = data.get("cover_image_thumbnail")
        self.generation_groups_limit = data.get("generation_groups_limit")

        self.owner = data.get("owner")
        if self.owner is not None:
            self.owner = User(self.owner)

        self.is_joined = data.get("is_joined")
        self.is_pending = data.get("is_pending")
        self.group_category_id = data.get("group_category_id")
        self.unread_counts = data.get("unread_counts")
        self.moderator_ids = data.get("moderator_ids")
        self.seizable = data.get("seizable")
        self.seizable_before = data.get("seizable_before")
        self.pending_count = data.get("pending_count")
        self.pending_transfer_id = data.get("pending_transfer_id")
        self.pending_deputize_ids = data.get("pending_deputize_ids")
        self.safe_mode = data.get("safe_mode")
        self.is_related = data.get("is_related")
        self.allow_ownership_transfer = data.get("allow_ownership_transfer")
        self.sub_category_id = data.get("sub_category_id")
        self.title = data.get("title")
        self.allow_thread_creation_by = data.get("allow_thread_creation_by")
        self.hide_from_game_eight = data.get("hide_from_game_eight")
        self.walkthrough_requested = data.get("walkthrough_requested")
        self.unread_threads_count = data.get("unread_threads_count")
        self.allow_members_to_post_image_and_video = data.get(
            "allow_members_to_post_image_and_video")
        self.allow_members_to_post_url = data.get("allow_members_to_post_url")
        self.guidelines = data.get("guidelines")
        self.invited_to_join = data.get("invited_to_join")

    def __repr__(self):
        return f'Group(data={self.data}, id={self.id})'


class GroupGiftHistory:

    __slots__ = ("data", "gifts_count", "received_date", "user")

    def __init__(self, data):
        self.data = data

        self.gifts_count = data.get("gifts_count")
        if self.gifts_count is not None:
            self.gifts_count = [
                GiftCount(gifts_count) for gifts_count in self.gifts_count
            ]

        self.received_date = data.get("received_date")

        self.user = data.get("user")
        if self.user is not None:
            self.user = User(self.user)

    def __repr__(self):
        return f'CallGiftHistory(data={self.data}, gifts_count={self.gifts_count})'


class GroupUser:

    __slots__ = (
        "data", "user", "is_moderator", "pending_transfer", "pending_deputize", "title"
    )

    def __init__(self, data):
        self.data = data

        self.user = data.get("user")
        if self.user is not None:
            self.user = User(self.user)

        self.is_moderator = data.get("is_moderator")
        self.pending_transfer = data.get("pending_transfer")
        self.pending_deputize = data.get("pending_deputize")
        self.title = data.get("title")

    def __repr__(self):
        return f'GroupUser(data={self.data}, user={self.user})'


class Interest:

    __slots__ = ("data", "id", "name", "icon", "selected")

    def __init__(self, data):
        self.data = data
        self.id = data.get("id")
        self.name = data.get("name")
        self.icon = data.get("icon")
        self.selected = data.get("selected")

    def __repr__(self):
        return f'Interest(data={self.data}, user_id={self.user_id})'


class MessageTag:

    __slots__ = ("data", "user_id", "offset", "length", "type")

    def __init__(self, data):
        self.data = data
        self.user_id = data.get("user_id")
        self.offset = data.get("offset")
        self.length = data.get("length")
        self.type = data.get("type")

    def __repr__(self):
        return f'MessageTag(data={self.data}, user_id={self.user_id})'


class Post:

    __slots__ = (
        "data", "id", "text", "post_type", "group_id", "font_size", "color",
        "likes_count", "created_at", "updated_at", "edited_at", "liked",
        "likers", "tag", "reposts_count", "reposted", "repostable",
        "reported_count", "conversation_id", "in_reply_to",
        "in_reply_to_post", "in_reply_to_post_count", "user", "mentions",
        "group", "conference_call", "attachment", "attachment_thumbnail",
        "attachment_2", "attachment_2_thumbnail", "attachment_3",
        "attachment_3_thumbnail", "attachment_4", "attachment_4_thumbnail",
        "attachment_5", "attachment_5_thumbnail", "attachment_6",
        "attachment_6_thumbnail", "attachment_7", "attachment_7_thumbnail",
        "attachment_8", "attachment_8_thumbnail", "attachment_9",
        "attachment_9_thumbnail", "shareable", "shared_url", "survey", "videos",
        "gifts_count", "shared_thread", "thread_id", "thread", "highlighted",
        "message_tags", "is_fail_to_send"
    )

    def __init__(self, data):
        self.data = data
        self.id = data.get("id")
        self.text = data.get("text")
        self.post_type = data.get("post_type")
        self.group_id = data.get("group_id")
        self.font_size = data.get("font_size")
        self.color = data.get("color")
        self.likes_count = data.get("likes_count")
        self.created_at = data.get("created_at")
        self.updated_at = data.get("updated_at")
        self.edited_at = data.get("edited_at")
        self.liked = data.get("liked")

        self.likers = data.get("likers")
        if self.likers is not None:
            self.likers = [
                User(liker) for liker in self.likers
            ]

        self.tag = data.get("tag")
        self.reposts_count = data.get("reposts_count")
        self.reposted = data.get("reposted")
        self.repostable = data.get("repostable")
        self.reported_count = data.get("reported_count")
        self.conversation_id = data.get("conversation_id")
        self.in_reply_to = data.get("in_reply_to")
        self.in_reply_to_post = data.get("in_reply_to_post")
        self.in_reply_to_post_count = data.get("in_reply_to_post_count")

        self.user = data.get("user")
        if self.user is not None:
            self.user = User(self.user)

        self.mentions = data.get("mentions")
        if self.mentions is not None:
            self.mentions = [
                User(mention) for mention in self.mentions
            ]

        self.group = data.get("group")
        if self.group is not None:
            self.group = Group(self.group)

        self.conference_call = data.get("conference_call")
        # if self.conference_call is not None:
        #     self.conference_call = ConferenceCall(self.conference_call)

        self.attachment = data.get("attachment")
        self.attachment_thumbnail = data.get("attachment_thumbnail")
        self.attachment_2 = data.get("attachment_2")
        self.attachment_2_thumbnail = data.get("attachment_2_thumbnail")
        self.attachment_3 = data.get("attachment_3")
        self.attachment_3_thumbnail = data.get("attachment_3_thumbnail")
        self.attachment_4 = data.get("attachment_4")
        self.attachment_4_thumbnail = data.get("attachment_4_thumbnail")
        self.attachment_5 = data.get("attachment_5")
        self.attachment_5_thumbnail = data.get("attachment_5_thumbnail")
        self.attachment_6 = data.get("attachment_6")
        self.attachment_6_thumbnail = data.get("attachment_6_thumbnail")
        self.attachment_7 = data.get("attachment_7")
        self.attachment_7_thumbnail = data.get("attachment_7_thumbnail")
        self.attachment_8 = data.get("attachment_8")
        self.attachment_8_thumbnail = data.get("attachment_8_thumbnail")
        self.attachment_9 = data.get("attachment_9")
        self.attachment_9_thumbnail = data.get("attachment_9_thumbnail")
        self.shareable = data.get("shareable")

        self.shared_url = data.get("shared_url")
        if self.shared_url is not None:
            self.shared_url = SharedUrl(self.shared_url)

        self.survey = data.get("survey")
        if self.survey is not None:
            self.survey = Survey(self.survey)

        self.videos = data.get("videos")
        if self.videos is not None:
            self.videos = [
                Video(video) for video in self.videos
            ]

        self.gifts_count = data.get("gifts_count")
        if self.gifts_count is not None:
            self.gifts_count = [
                GiftCount(gifts_count) for gifts_count in self.gifts_count
            ]

        self.shared_thread = data.get("shared_thread")
        if self.shared_thread is not None:
            self.shared_thread = [
                ThreadInfo(shared_thread) for shared_thread in self.shared_thread
            ]

        self.thread_id = data.get("thread_id")

        self.thread = data.get("thread")
        if self.thread is not None:
            self.thread = [
                ThreadInfo(thread) for thread in self.thread
            ]

        self.highlighted = data.get("highlighted")

        self.message_tags = data.get("message_tags")
        if self.message_tags is not None:
            self.message_tags = [
                MessageTag(message_tag) for message_tag in self.message_tags
            ]

        self.is_fail_to_send = data.get("is_fail_to_send")

    def __repr__(self):
        return f'Post(data=({self.data}, id={self.id})'


class PostTag:

    __slots__ = ("data", "id", "tag", "post_hashtags_count")

    def __init__(self, data):
        self.data = data
        self.id = data.get("id")
        self.tag = data.get("tag")
        self.post_hashtags_count = data.get("post_hashtags_count")

    def __repr__(self):
        return f'PostTag(data={self.data}, id={self.id})'


class RecentSearch:

    __slots__ = ("data", "id", "type", "user", "hashtag", "keyword")

    def __init__(self, data):
        self.data = data
        self.id = data.get("id")
        self.type = data.get("type")

        self.user = data.get("user")
        if self.user is not None:
            self.user = User(self.user)

        self.hashtag = data.get("hashtag")
        if self.hashtag is not None:
            self.hashtag = PostTag(self.hashtag)

        self.keyword = data.get("keyword")

    def __repr__(self):
        return f'RecentSearch(data={self.data}, gifts_count={self.gifts_count})'


class Review:

    __slots__ = (
        "data", "id", "comment", "reported_count", "user", "reviewer",
        "created_at", "mutual_review",
    )

    def __init__(self, data):
        self.data = data
        self.id = data.get("id")
        self.comment = data.get("comment")
        self.reported_count = data.get("reported_count")

        self.user = data.get("user")
        if self.user is not None:
            self.user = User(self.user)

        self.reviewer = data.get("reviewer")
        self.created_at = data.get("created_at")
        self.mutual_review = data.get("mutual_review")

    def __repr__(self):
        return f'Review(data={self.data}, id={self.id})'


class SearchCriteria:

    __slots__ = (
        "data", "nickname", "username", "biography", "prefecture", "gender"
    )

    def __init__(self, data):
        self.data = data
        self.nickname = data.get("nickname")
        self.username = data.get("username")
        self.biography = data.get("biography")
        self.prefecture = data.get("prefecture")
        self.gender = data.get("gender")

    def __repr__(self):
        return f'SearchCriteria(data={self.data}, nickname={self.nickname})'


class Shareable:

    __slots__ = ("data", "post", "group", "thread")

    def __init__(self, data):
        self.data = data

        self.post = data.get("post")
        if self.post is not None:
            self.post = Post(self.post)

        self.group = data.get("group")
        if self.group is not None:
            self.group = Group(self.group)

        self.thread = data.get("thread")
        if self.thread is not None:
            self.thread = ThreadInfo(self.thread)

    def __repr__(self):
        return f'Shareable(data={self.data}, post={self.post})'


class SharedUrl:

    __slots__ = ("data", "url", "title", "description", "image_url")

    def __init__(self, data):
        self.data = data
        self.url = data.get("url")
        self.title = data.get("title")
        self.description = data.get("description")
        self.image_url = data.get("image_url")

    def __repr__(self):
        return f'SharedUrl(data={self.data}, url={self.url})'


class Survey:

    __slots__ = ("data", "id", "votes_count", "choices", "voted")

    def __init__(self, data):
        self.data = data
        self.id = data.get("id")
        self.votes_count = data.get("votes_count")

        self.choices = data.get("choices")
        # if self.choices is not None:
        #     self.choices = [
        #         Choise(choise) for choise in self.choices
        #     ]

        self.voted = data.get("voted")

    def __repr__(self):
        return f'Survey(id={self.id})'


class ThreadInfo:

    __slots__ = (
        "data", "id", "title", "owner", "last_post", "unread_count", "posts_count",
        "created_at", "updated_at", "thread_icon", "is_joined", "new_updates"
    )

    def __init__(self, data):
        self.data = data
        self.id = data.get("id")
        self.title = data.get("title")

        self.owner = data.get("owner")
        if self.owner is not None:
            self.owner = User(self.owner)

        self.last_post = data.get("last_post")
        if self.last_post is not None:
            self.last_post = Post(self.last_post)

        self.unread_count = data.get("unread_count")
        self.posts_count = data.get("posts_count")
        self.created_at = data.get("created_at")
        self.updated_at = data.get("updated_at")
        self.thread_icon = data.get("thread_icon")
        self.is_joined = data.get("is_joined")
        self.new_updates = data.get("new_updates")

    def __repr__(self):
        return f'ThreadInfo(data={self.data}, id={self.id})'


class User:

    # TODO: レスポンスのキーとuserの属性名が違うので合わせましょう。

    __slots__ = (
        "data", "id", "biography", "birth_date", "cover_image", "gender",
        "cover_image_thumbnail", "followers_count", "followings_count",
        "groups_users_count", "is_age_verified", "is_chat_request_on",
        "is_dangerous", "is_email_confirmed", "is_facebook_connected",
        "is_follow_pending", "is_followed_by", "is_following",
        "is_from_different_generation_and_trusted", "is_group_phone_on",
        "is_hidden", "is_hide_vip", "is_interests_selected", "is_line_connected",
        "is_lobi_connected", "is_new", "is_phone_on", "is_private",
        "is_recently_penalized", "is_registered", "is_video_pn", "is_vip",
        "masked_email", "mobile_number", "nickname", "posts_count", "prefecture",
        "profile_icon", "profile_icon_thumbnail", "reviews_count"
    )

    def __init__(self, data):
        self.data = data
        self.id = data.get("id")
        self.biography = data.get("biography")
        self.birth_date = data.get("birth_date")
        self.cover_image = data.get("cover_image")
        self.gender = data.get("gender")
        self.cover_image_thumbnail = data.get("cover_image_thumbnail")
        self.followers_count = data.get("followers_count")
        self.followings_count = data.get("followings_count")
        self.groups_users_count = data.get("groups_users_count")
        self.is_age_verified = data.get("is_age_verified")
        self.is_chat_request_on = data.get("is_chat_request_on")
        self.is_dangerous = data.get("dangerous_user")
        self.is_email_confirmed = data.get("is_email_confirmed")
        self.is_facebook_connected = data.get("is_facebook_connected")
        self.is_follow_pending = data.get("is_follow_pending")
        self.is_followed_by = data.get("is_followed_by")
        self.is_following = data.get("is_following")
        self.is_from_different_generation_and_trusted = data.get(
            "is_from_different_generation_and_trusted")
        self.is_group_phone_on = data.get("is_group_phone_on")
        self.is_hidden = data.get("is_hidden")
        self.is_hide_vip = data.get("is_hide_vip")
        self.is_interests_selected = data.get("interests_selected")
        self.is_line_connected = data.get("is_line_connected")
        self.is_lobi_connected = data.get("is_lobi_connected")
        self.is_new = data.get("new_user")
        self.is_phone_on = data.get("is_phone_on")
        self.is_private = data.get("is_private")
        self.is_recently_penalized = data.get("recently_kenta")
        self.is_registered = data.get("is_registered")
        self.is_video_pn = data.get("is_video_pn")
        self.is_vip = data.get("is_vip")
        self.masked_email = data.get("masked_email")
        self.mobile_number = data.get("mobile_number")
        self.nickname = data.get("nickname")
        self.posts_count = data.get("posts_count")
        self.prefecture = data.get("prefecture")
        self.profile_icon = data.get("profile_icon")
        self.profile_icon_thumbnail = data.get("profile_icon_thumbnail")
        self.reviews_count = data.get("reviews_count")

    def __repr__(self):
        return f'User(data={self.data}, id={self.id})'


class UserWrapper:

    __slots__ = ("data", "id", "user")

    def __init__(self, data):
        self.data = data
        self.id = data.get("id")

        self.user = data.get("user")
        if self.user is not None:
            self.user = User(self.user)

    def __repr__(self):
        return f'UserWrapper(data={self.data}, id={self.id})'


class Video:

    __slots__ = (
        "data", "id", "completed", "width", "height", "bitrate",
        "views_count", "video_url", "thumbnail_url", "thumbnail_big_url"
    )

    def __init__(self, data):
        self.data = data
        self.id = data.get("id")
        self.completed = data.get("completed")
        self.width = data.get("width")
        self.height = data.get("height")
        self.bitrate = data.get("bitrate")
        self.views_count = data.get("views_count")
        self.video_url = data.get("video_url")
        self.thumbnail_url = data.get("thumbnail_url")
        self.thumbnail_big_url = data.get("thumbnail_big_url")

    def __repr__(self):
        return f'Video(data={self.data}, id={self.id})'


class Walkthrough:

    __slots__ = ("data", "title", "url")

    def __init__(self, data):
        self.data = data
        self.title = data.get("title")
        self.url = data.get("url")

    def __repr__(self):
        return f'Shareable(data={self.data}, title={self.title})'


class WalletTransaction:

    __slots__ = ("data", "id", "created_at", "description", "amount", "coins")

    def __init__(self, data):
        self.data = data
        self.id = data.get("id")
        self.created_at = data.get("created_at")
        self.description = data.get("description")
        self.amount = data.get("amount")

        self.coins = data.get("coins")
        if self.coins is not None:
            self.coins = CoinAmount(self.coins)

    def __repr__(self):
        return f'WalletTransaction(data={self.data}, id={self.id})'
