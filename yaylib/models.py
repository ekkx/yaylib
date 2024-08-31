"""
MIT License

Copyright (c) 2023 ekkx

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

import json
from typing import List, Optional


class Model:
    pass


class Activity(Model):
    __slots__ = (
        "data",
        "id",
        "created_at",
        "type",
        "user",
        "from_post",
        "to_post",
        "group",
        "followers",
        "followers_count",
        "from_post_ids",
        "vip_reward",
        "metadata",
        "birthday_users",
        "birthday_users_count",
    )

    def __init__(self, data: dict):
        self.data = data
        self.id = data.get("id")
        self.created_at = data.get("created_at")
        self.type = data.get("type")

        self.user: User = data.get("user")
        if self.user is not None:
            self.user = User(self.user)

        self.from_post: Post = data.get("from_post")
        if self.from_post is not None:
            self.from_post = Post(self.from_post)

        self.to_post: Post = data.get("to_post")
        if self.to_post is not None:
            self.to_post = Post(self.to_post)

        self.group: Group = data.get("group")
        if self.group is not None:
            self.group = Group(self.group)

        self.followers: List[User] = data.get("followers")
        if self.followers is not None:
            self.followers = [User(follower) for follower in self.followers]

        self.followers_count = data.get("followers_count")
        self.from_post_ids = data.get("from_post_ids")
        self.vip_reward = data.get("vip_reward")

        self.metadata: Metadata = data.get("metadata")
        if self.metadata is not None:
            self.metadata = Metadata(self.metadata)

        self.birthday_users: List[User] = data.get("birthday_users")
        if self.birthday_users is not None:
            self.birthday_users = [
                User(birthday_user) for birthday_user in self.birthday_users
            ]

        self.birthday_users_count = data.get("birthday_users_count")

    def __repr__(self):
        return f"Activity(data={self.data})"


class Metadata(Model):
    __slots__ = (
        "data",
        "body",
        "bulk_invitation",
        "content_preview",
        "title",
        "url",
    )

    def __init__(self, data: dict):
        self.data = data.get("data")
        self.body = data.get("body")
        self.bulk_invitation = data.get("bulk_invitation")
        self.content_preview = data.get("content_preview")
        self.title = data.get("title")
        self.url = data.get("url")

    def __repr__(self):
        return f"Metadata(data={self.data})"


class ApplicationConfig(Model):
    __slots__ = (
        "data",
        "id",
        "name",
        "server_name",
        "description",
        "itunes_app_id",
        "settings",
    )

    def __init__(self, data: dict):
        self.data = data
        self.id = data.get("id")
        self.name = data.get("name")
        self.server_name = data.get("server_name")
        self.description = data.get("description")
        self.itunes_app_id = data.get("itunes_app_id")

        self.settings: ApplicationConfigSettings = data.get("settings")
        if self.settings is not None:
            self.settings = ApplicationConfigSettings(self.settings)

    def __repr__(self):
        return f"ApplicationConfig(data={self.data})"


class ApplicationConfigSettings(Model):
    __slots__ = (
        "data",
        "minimum_app_version_required",
        "minimum_android_app_version_required",
        "line_official_account_id",
        "use_random_message_refresh_rate",
        "id_card_check_required_group_category_ids",
        "is_id_card_check_required_blocker_view_enabled",
        "support_email_address",
        "minimum_ios_app_version_required",
        "is_id_card_and_phone_verification_check_for_review_enabled",
        "minimum_android_version_supported",
        "is_maintenanced",
        "promotion_sticker_pack_ids",
        "minimum_ios_version_supported",
        "is_direct_vip_purchase_enabled",
        "is_speed_test_enabled",
        "is_initial_post_enabled",
        "is_phone_verification_required_blocker_view_enabled",
        "latest_android_app_version",
        "is_web_phone_verification_enabled",
        "latest_ios_app_version",
        "max_image_frame_count",
        "localized_maintenance_url",
        "ad_tester_user_ids",
        "should_append_user_id_to_news_url",
        "is_star_enabled",
        "news_version",
        "localized_news_title",
        "is_chat_websocket_enabled",
        "is_oauth_enabled",
        "is_call_moderator_enabled",
        "is_gift_features_enabled",
        "is_call_chat_deletion_enabled",
        "twitter_official_account_id",
        "is_web_subscription_enabled",
        "localized_news_url",
    )

    def __init__(self, data: dict):
        self.data = data
        self.minimum_app_version_required = data.get("minimum_app_version_required")
        self.minimum_android_app_version_required = data.get(
            "minimum_android_app_version_required"
        )
        self.line_official_account_id = data.get("line_official_account_id")
        self.use_random_message_refresh_rate = data.get(
            "use_random_message_refresh_rate"
        )
        self.id_card_check_required_group_category_ids = data.get(
            "id_card_check_required_group_category_ids"
        )
        self.is_id_card_check_required_blocker_view_enabled = data.get(
            "is_id_card_check_required_blocker_view_enabled"
        )
        self.support_email_address = data.get("support_email_address")
        self.minimum_ios_app_version_required = data.get(
            "minimum_ios_app_version_required"
        )
        self.is_id_card_and_phone_verification_check_for_review_enabled = data.get(
            "is_id_card_and_phone_verification_check_for_review_enabled"
        )
        self.minimum_android_version_supported = data.get(
            "minimum_android_version_supported"
        )
        self.is_maintenanced = data.get("is_maintenanced")
        self.promotion_sticker_pack_ids = data.get("promotion_sticker_pack_ids")
        self.minimum_ios_version_supported = data.get("minimum_ios_version_supported")
        self.is_direct_vip_purchase_enabled = data.get("is_direct_vip_purchase_enabled")
        self.is_speed_test_enabled = data.get("is_speed_test_enabled")
        self.is_initial_post_enabled = data.get("is_initial_post_enabled")
        self.is_phone_verification_required_blocker_view_enabled = data.get(
            "is_phone_verification_required_blocker_view_enabled"
        )
        self.latest_android_app_version = data.get("latest_android_app_version")
        self.is_web_phone_verification_enabled = data.get(
            "is_web_phone_verification_enabled"
        )
        self.latest_ios_app_version = data.get("latest_ios_app_version")
        self.max_image_frame_count = data.get("max_image_frame_count")
        self.localized_maintenance_url = data.get("localized_maintenance_url")
        self.ad_tester_user_ids = data.get("ad_tester_user_ids")
        self.should_append_user_id_to_news_url = data.get(
            "should_append_user_id_to_news_url"
        )
        self.is_star_enabled = data.get("is_star_enabled")
        self.news_version = data.get("news_version")
        self.localized_news_title = data.get("localized_news_title")
        self.is_chat_websocket_enabled = data.get("is_chat_websocket_enabled")
        self.is_oauth_enabled = data.get("is_oauth_enabled")
        self.is_call_moderator_enabled = data.get("is_call_moderator_enabled")
        self.is_gift_features_enabled = data.get("is_gift_features_enabled")
        self.is_call_chat_deletion_enabled = data.get("is_call_chat_deletion_enabled")
        self.twitter_official_account_id = data.get("twitter_official_account_id")
        self.is_web_subscription_enabled = data.get("is_web_subscription_enabled")
        self.localized_news_url = data.get("localized_news_url")

    def __repr__(self):
        return f"ApplicationConfigSettings(data={self.data})"


class Attachment(Model):
    def __init__(
        self,
        file,
        filename,
        original_file_name,
        original_file_extension,
        natural_width,
        natural_height,
        is_thumb: bool,
    ):
        self.file = file
        self.filename = filename
        self.original_file_name = original_file_name
        self.original_file_extension = original_file_extension
        self.natural_width = natural_width
        self.natural_height = natural_height
        self.is_thumb = is_thumb

    def __repr__(self):
        return f"Attachment(filename={self.filename})"


class BanWord(Model):
    __slots__ = ("data", "id", "type", "word")

    def __init__(self, data: dict):
        self.data = data
        self.id = data.get("id")
        self.type = data.get("type")
        self.word = data.get("word")

    def __repr__(self):
        return f"BanWord(data={self.data})"


class Bgm(Model):
    __slots__ = ("data", "id", "title", "music_url", "order")

    def __init__(self, data: dict):
        self.data = data
        self.id = data.get("id")
        self.title = data.get("title")
        self.music_url = data.get("music_url")
        self.order = data.get("order")

    def __repr__(self):
        return f"Bgm(data={self.data})"


class CallGiftHistory(Model):
    __slots__ = ("data", "gifts_count", "sent_at", "sender")

    def __init__(self, data: dict):
        self.data = data

        self.gifts_count: List[GiftCount] = data.get("gifts_count")
        if self.gifts_count is not None:
            self.gifts_count = [
                GiftCount(gifts_count) for gifts_count in self.gifts_count
            ]

        self.sent_at = data.get("sent_at")

        self.sender: User = data.get("sender")
        if self.sender is not None:
            self.sender = User(self.sender)

    def __repr__(self):
        return f"CallGiftHistory(data={self.data})"


class ChatRoom(Model):
    __slots__ = (
        "data",
        "id",
        "unread_count",
        "updated_at",
        "members",
        "background",
        "last_message",
        "name",
        "is_group",
        "owner",
        "is_request",
        "is_pinned",
        "is_notification_on",
    )

    def __init__(self, data: dict):
        self.data = data
        self.id = data.get("id")
        self.unread_count = data.get("unread_count")
        self.updated_at = data.get("updated_at")

        self.members: List[User] = data.get("members")
        if self.members is not None:
            self.members = [User(member) for member in self.members]

        self.background = data.get("background")

        self.last_message: Message = data.get("last_message")
        if self.last_message is not None:
            self.last_message = Message(self.last_message)

        self.name = data.get("name")
        self.is_group = data.get("is_group")

        self.owner: User = data.get("owner")
        if self.owner is not None:
            self.owner = User(self.owner)

        self.is_request = data.get("is_request")
        self.is_pinned = data.get("is_pinned")
        self.is_notification_on = data.get("is_notification_on")

    def __repr__(self):
        return f"ChatRoom(data={self.data})"


class ChatRoomDraft(Model):
    __slots__ = ("data", "id", "text")

    def __init__(self, data: dict):
        self.data = data
        self.id = data.get("id")
        self.text = data.get("text")

    def __repr__(self):
        return f"ChatRoomDraft(data={self.data})"


class MessageEvent(Model):
    __slots__ = (
        "data",
        "message",
        "event",
    )

    def __init__(self, data: dict):
        self.data = data

        self.message: Message = data.get("data")
        if self.message is not None:
            self.message = Message(self.message)

        self.event = data.get("event")

    def __repr__(self):
        return f"MessageEvent(data={self.data})"


class ChatRoomEvent(Model):
    __slots__ = (
        "data",
        "icon_thumbnail",
        "id",
        "last_message",
        "name",
        "unread_count",
    )

    def __init__(self, data: dict):
        self.data = data
        self.icon_thumbnail = data.get("icon_thumbnail")
        self.id = data.get("id")

        self.last_message: Message = data.get("last_message")
        if self.last_message is not None:
            self.last_message = Message(self.last_message)

        self.name = data.get("name")
        self.unread_count = data.get("unread_count")

    def __repr__(self):
        return f"ChatRoomEvent(data={self.data})"


class GroupUpdatesEvent(Model):
    __slots__ = (
        "response",
        "data",
        "event",
    )

    def __init__(self, data: dict):
        self.response = data
        self.data = data.get("data")
        self.event = data.get("event")

    def __repr__(self):
        return f"GroupUpdateEvent(data={self.response})"


class Choice(Model):
    __slots__ = ("data", "id", "label", "votes_count")

    def __init__(self, data: dict):
        self.data = data
        self.id = data.get("id")
        self.label = data.get("label")
        self.votes_count = data.get("votes_count")

    def __repr__(self):
        return f"Choice(data={self.data})"


class CoinAmount(Model):
    __slots__ = ("data", "paid", "free", "total")

    def __init__(self, data: dict):
        self.data = data
        self.paid = data.get("paid")
        self.free = data.get("free")
        self.total = data.get("total")

    def __repr__(self):
        return f"CoinAmount(data={self.data})"


class CoinExpiration(Model):
    __slots__ = ("data", "expired_at", "amount")

    def __init__(self, data: dict):
        self.data = data
        self.expired_at = data.get("expired_at")
        self.amount = data.get("amount")

    def __repr__(self):
        return f"CoinExpiration(data={self.data})"


class CoinProduct(Model):
    __slots__ = ("data", "id", "purchasable", "amount")

    def __init__(self, data: dict):
        self.data = data
        self.id = data.get("id")
        self.purchasable = data.get("purchasable")
        self.amount = data.get("amount")

    def __repr__(self):
        return f"CoinProduct(data={self.data})"


class CoinProductQuota(Model):
    __slots__ = ("data", "bought", "limit")

    def __init__(self, data: dict):
        self.data = data
        self.bought = data.get("bought")
        self.limit = data.get("limit")

    def __repr__(self):
        return f"CoinProductQuota(data={self.data})"


class ConferenceCall(Model):
    __slots__ = (
        "data",
        "id",
        "post_id",
        "group_id",
        "is_active",
        "anonymous_call_users_count",
        "agora_channel",
        "agora_token",
        "call_type",
        "joinable_by",
        "game",
        "genre",
        "duration_seconds",
        "max_participants",
        "conference_call_users",
        "bump_params",
        "conference_call_user_roles",
    )

    def __init__(self, data: dict):
        self.data = data
        self.id = data.get("id")
        self.post_id = data.get("post_id")
        self.group_id = data.get("group_id")
        self.is_active = data.get("is_active")
        self.anonymous_call_users_count = data.get("anonymous_call_users_count")
        self.agora_channel = data.get("agora_channel")
        self.agora_token = data.get("agora_token")
        self.call_type = data.get("call_type")
        self.joinable_by = data.get("joinable_by")

        self.game: Game = data.get("game")
        if self.game is not None:
            self.game = Game(self.game)

        self.genre: Genre = data.get("genre")
        if self.genre is not None:
            self.genre = Genre(self.genre)

        self.duration_seconds = data.get("duration_seconds")
        self.max_participants = data.get("max_participants")

        self.conference_call_users: List[User] = data.get("conference_call_users")
        if self.conference_call_users is not None:
            self.conference_call_users = [
                User(conference_call_user)
                for conference_call_user in self.conference_call_users
            ]

        self.bump_params = data.get("bump_params")
        # if self.bump_params is not None:
        #     self.bump_params = BumpParams(self.bump_params)

        self.conference_call_user_roles: List[ConferenceCallUserRole] = data.get(
            "conference_call_user_roles"
        )
        if self.conference_call_user_roles is not None:
            self.conference_call_user_roles = [
                ConferenceCallUserRole(conference_call_user_role)
                for conference_call_user_role in self.conference_call_user_roles
            ]

    def __repr__(self):
        return f"ConferenceCall(data={self.data})"


class ConferenceCallUserRole(Model):
    __slots__ = ("data", "id", "user_id", "role")

    def __init__(self, data: dict):
        self.data = data
        self.id = data.get("id")
        self.user_id = data.get("user_id")
        self.role = data.get("role")

    def __repr__(self):
        return f"ConferenceCallUserRole(data={self.data})"


class ContactStatus(Model):
    __slots__ = ("data", "status", "user_id")

    def __init__(self, data: dict):
        self.data = data
        self.status = data.get("status")
        self.user_id = data.get("user_id")

    def __repr__(self):
        return f"ContactStatus(data={self.data})"


class CreateGroupQuota(Model):
    __slots__ = ("data", "used_quota", "remaining_quota")

    def __init__(self, data: dict):
        self.data = data
        self.used_quota = data.get("used_quota")
        self.remaining_quota = data.get("remaining_quota")

    def __repr__(self):
        return f"CreateGroupQuota(data={self.data})"


class TimelineSettings(Model):
    __slots__ = (
        "data",
        "hide_hot_post",
        "hide_reply_public_timeline",
        "hide_reply_following_timeline",
        "faves_filter",
    )

    def __init__(self, data: dict):
        self.data = data
        self.hide_hot_post = data.get("hide_hot_post")
        self.hide_reply_public_timeline = data.get("hide_reply_public_timeline")
        self.hide_reply_following_timeline = data.get("hide_reply_following_timeline")
        self.faves_filter = data.get("faves_filter")

    def __repr__(self):
        return f"TimelineSettings(data={self.data})"


class Error(Model):
    __slots__ = ("data", "throwable", "type", "action")

    def __init__(self, data: dict):
        self.data = data
        self.throwable = data.get("throwable")
        self.type = data.get("type")
        self.action = data.get("action")

    def __repr__(self):
        return f"Error(data={self.data})"


class Footprint(Model):
    __slots__ = ("data", "user", "visited_at", "id")

    def __init__(self, data: dict):
        self.data = data

        self.user: User = data.get("user")
        if self.user is not None:
            self.user = User(self.user)

        self.visited_at = data.get("visited_at")
        self.id = data.get("id")

    def __repr__(self):
        return f"Footprint(data={self.data})"


class Game(Model):
    __slots__ = ("data", "id", "type", "title", "icon_url", "platform_details")

    def __init__(self, data: dict):
        self.data = data
        self.id = data.get("id")
        self.type = data.get("type")
        self.title = data.get("title")
        self.icon_url = data.get("icon_url")

        self.platform_details: PlatformDetails = data.get("platform_details")
        if self.platform_details is not None:
            self.platform_details = PlatformDetails(self.platform_details)

    def __repr__(self):
        return f"Game(data={self.data})"


class Genre(Model):
    __slots__ = ("data", "id", "type", "title", "icon_url")

    def __init__(self, data: dict):
        self.data = data
        self.id = data.get("id")
        self.type = data.get("type")
        self.title = data.get("title")
        self.icon_url = data.get("icon_url")

    def __repr__(self):
        return f"Genre(data={self.data})"


class GifImage(Model):
    __slots__ = ("data", "id", "url", "width", "height")

    def __init__(self, data: dict):
        self.data = data
        self.id = data.get("id")
        self.url = data.get("url")
        self.width = data.get("width")
        self.height = data.get("height")

    def __repr__(self):
        return f"GifImage(data={self.data})"


class GifImageCategory(Model):
    __slots__ = ("data", "id", "name", "language", "gifs")

    def __init__(self, data: dict):
        self.data = data
        self.id = data.get("id")
        self.name = data.get("name")
        self.language = data.get("language")

        self.gifs: List[GifImage] = data.get("gifs")
        if self.gifs is not None:
            self.gifs = [GifImage(gif) for gif in self.gifs]

    def __repr__(self):
        return f"GifImageCategory(data={self.data})"


class Gift(Model):
    __slots__ = ("data", "id", "title", "icon", "iconThumbnail", "price")

    def __init__(self, data: dict):
        self.data = data
        self.id = data.get("id")
        self.type = data.get("title")
        self.title = data.get("icon")
        self.icon_url = data.get("iconThumbnail")
        self.icon_url = data.get("price")

    def __repr__(self):
        return f"Gift(data={self.data})"


class GiftCount(Model):
    __slots__ = ("data", "id", "quantity")

    def __init__(self, data: dict):
        self.data = data
        self.id = data.get("id")
        self.quantity = data.get("quantity")

    def __repr__(self):
        return f"GiftCount(data={self.data})"


class GiftHistory(Model):
    __slots__ = ("data", "transaction_at_seconds", "user", "gifts")

    def __init__(self, data: dict):
        self.data = data
        self.transaction_at_seconds = data.get("transaction_at_seconds")

        self.user: User = data.get("user")
        if self.user is not None:
            self.user = User(self.user)

        self.gifts: List[ReceivedGift] = data.get("gifts")
        if self.gifts is not None:
            self.gifts = [ReceivedGift(gift) for gift in self.gifts]

    def __repr__(self):
        return f"GiftHistory(data={self.data})"


class GiftingAbility(Model):
    __slots__ = ("data", "user_id", "enabled", "can_send", "can_receive")

    def __init__(self, data: dict):
        self.data = data
        self.user_id = data.get("user_id")
        self.enabled = data.get("enabled")
        self.can_send = data.get("can_send")
        self.can_receive = data.get("can_receive")

    def __repr__(self):
        return f"GiftingAbility(data={self.data})"


class Group(Model):
    __slots__ = (
        "data",
        "id",
        "topic",
        "description",
        "groups_users_count",
        "posts_count",
        "user_id",
        "threads_count",
        "highlighted_count",
        "views_count",
        "related_count",
        "secret",
        "gender",
        "place",
        "hide_reported_posts",
        "hide_conference_call",
        "is_private",
        "only_verified_age",
        "only_mobile_verified",
        "call_timeline_display",
        "updated_at",
        "group_icon",
        "group_icon_thumbnail",
        "cover_image",
        "cover_image_thumbnail",
        "generation_groups_limit",
        "owner",
        "is_joined",
        "is_pending",
        "group_category_id",
        "homepage",
        "unread_counts",
        "moderator_ids",
        "seizable",
        "seizable_before",
        "pending_count",
        "pending_transfer_id",
        "pending_deputize_ids",
        "safe_mode",
        "is_related",
        "allow_ownership_transfer",
        "sub_category_id",
        "title",
        "allow_thread_creation_by",
        "hide_from_game_eight",
        "walkthrough_requested",
        "unread_threads_count",
        "allow_members_to_post_image_and_video",
        "allow_members_to_post_url",
        "guidelines",
        "invited_to_join",
    )

    def __init__(self, data: dict):
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
        self.place = data.get("place")
        self.hide_reported_posts = data.get("hide_reported_posts")
        self.hide_conference_call = data.get("hide_conference_call")
        self.is_private = data.get("is_private")
        self.only_verified_age = data.get("only_verified_age")
        self.only_mobile_verified = data.get("only_mobile_verified")
        self.call_timeline_display = data.get("call_timeline_display")
        self.updated_at = data.get("updated_at")
        self.group_icon = data.get("group_icon")
        self.group_icon_thumbnail = data.get("group_icon_thumbnail")
        self.cover_image = data.get("cover_image")
        self.cover_image_thumbnail = data.get("cover_image_thumbnail")
        self.generation_groups_limit = data.get("generation_groups_limit")

        self.owner: User = data.get("owner")
        if self.owner is not None:
            self.owner = User(self.owner)

        self.is_joined = data.get("is_joined")
        self.is_pending = data.get("is_pending")
        self.group_category_id = data.get("group_category_id")
        self.homepage = data.get("homepage")
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
            "allow_members_to_post_image_and_video"
        )
        self.allow_members_to_post_url = data.get("allow_members_to_post_url")
        self.guidelines = data.get("guidelines")
        self.invited_to_join = data.get("invited_to_join")

    def __repr__(self):
        return f"Group(data={self.data})"


class GroupCategory(Model):
    __slots__ = ("data", "id", "name", "icon", "rank")

    def __init__(self, data: dict):
        self.data = data
        self.id = data.get("id")
        self.name = data.get("name")
        self.icon = data.get("icon")
        self.rank = data.get("rank")

    def __repr__(self):
        return f"GroupCategory(data={self.data})"


class GroupGiftHistory(Model):
    __slots__ = ("data", "gifts_count", "received_date", "user")

    def __init__(self, data: dict):
        self.data = data

        self.gifts_count: List[GiftCount] = data.get("gifts_count")
        if self.gifts_count is not None:
            self.gifts_count = [
                GiftCount(gifts_count) for gifts_count in self.gifts_count
            ]

        self.received_date = data.get("received_date")

        self.user: User = data.get("user")
        if self.user is not None:
            self.user = User(self.user)

    def __repr__(self):
        return f"GroupGiftHistory(data={self.data})"


class GroupUser(Model):
    __slots__ = (
        "data",
        "user",
        "is_moderator",
        "pending_transfer",
        "pending_deputize",
        "title",
    )

    def __init__(self, data: dict):
        self.data = data

        self.user: User = data.get("user")
        if self.user is not None:
            self.user = User(self.user)

        self.is_moderator = data.get("is_moderator")
        self.pending_transfer = data.get("pending_transfer")
        self.pending_deputize = data.get("pending_deputize")
        self.title = data.get("title")

    def __repr__(self):
        return f"GroupUser(data={self.data})"


class HiddenRecommendedPost(Model):
    __slots__ = ("data", "post")

    def __init__(self, data: dict):
        self.data = data
        self.post: Post = data.get("post")
        if self.post is not None:
            self.post = Post(self.post)

    def __repr__(self):
        return f"HiddenRecommendedPost(data={self.data})"


class Interest(Model):
    __slots__ = ("data", "id", "name", "icon", "selected")

    def __init__(self, data: dict):
        self.data = data
        self.id = data.get("id")
        self.name = data.get("name")
        self.icon = data.get("icon")
        self.selected = data.get("selected")

    def __repr__(self):
        return f"Interest(data={self.data})"


class Message(Model):
    __slots__ = (
        "data",
        "attachment",
        "attachment_android",
        "attachment_read_count",
        "attachment_thumbnail",
        "conference_call",
        "parent",
        "created_at",
        "font_size",
        "gif",
        "id",
        "message_type",
        "reactions_count",
        "reacted",
        "room_id",
        "sticker",
        "is_sent",
        "refresh_retry_count",
        "is_error",
        "text",
        "user_id",
        "video_processed",
        "video_thumbnail_big_url",
        "video_thumbnail_url",
        "video_url",
    )

    def __init__(self, data: dict):
        self.data = data
        self.attachment = data.get("attachment")
        self.attachment_android = data.get("attachment_android")
        self.attachment_read_count = data.get("attachment_read_count")
        self.attachment_thumbnail = data.get("attachment_thumbnail")

        self.conference_call: ConferenceCall = data.get("conference_call")
        if self.conference_call is not None:
            self.conference_call = ConferenceCall(self.conference_call)

        self.parent: ParentMessage = data.get("parent")
        if self.parent is not None:
            self.parent = ParentMessage(self.parent)

        self.created_at = data.get("created_at")
        self.font_size = data.get("font_size")

        self.gif: GifImage = data.get("gif")
        if self.gif is not None:
            self.gif = GifImage(self.gif)

        self.id = data.get("id")
        self.message_type = data.get("message_type")
        self.reactions_count = data.get("reactions_count")
        self.reacted = data.get("reacted")
        self.room_id = data.get("room_id")

        self.sticker: Sticker = data.get("sticker")
        if self.sticker is not None:
            self.sticker = Sticker(self.sticker)

        self.is_sent = data.get("is_sent")
        self.refresh_retry_count = data.get("refresh_retry_count")
        self.is_error = data.get("is_error")
        self.text = data.get("text")
        self.user_id = data.get("user_id")
        self.video_processed = data.get("video_processed")
        self.video_thumbnail_big_url = data.get("video_thumbnail_big_url")
        self.video_thumbnail_url = data.get("video_thumbnail_url")
        self.video_url = data.get("video_url")

    def __repr__(self):
        return f"Message(data={self.data})"


class ParentMessage(Model):
    __slots__ = (
        "data",
        "attachment",
        "attachment_android",
        "attachment_thumbnail",
        "created_at",
        "font_size",
        "gif",
        "id",
        "message_type",
        "reactions_count",
        "room_id",
        "sticker",
        "text",
        "user_id",
        "video_processed",
        "video_thumbnail_url",
        "video_thumbnail_big_url",
        "video_url",
        "parent_id",
        "reacted",
    )

    def __init__(self, data: dict):
        self.data = data
        self.attachment = data.get("attachment")
        self.attachment_android = data.get("attachment_android")
        self.attachment_thumbnail = data.get("attachment_thumbnail")
        self.created_at = data.get("created_at")
        self.font_size = data.get("font_size")

        self.gif: GifImage = data.get("gif")
        if self.gif is not None:
            self.gif = GifImage(self.gif)

        self.id = data.get("id")
        self.message_type = data.get("message_type")
        self.reactions_count = data.get("reactions_count")
        self.room_id = data.get("room_id")

        self.sticker: Sticker = data.get("sticker")
        if self.sticker is not None:
            self.sticker = Sticker(self.sticker)

        self.text = data.get("text")
        self.user_id = data.get("user_id")
        self.video_processed = data.get("video_processed")
        self.video_thumbnail_url = data.get("video_thumbnail_url")
        self.video_thumbnail_big_url = data.get("video_thumbnail_big_url")
        self.video_url = data.get("video_url")
        self.parent_id = data.get("parent_id")
        self.reacted = data.get("reacted")

    def __repr__(self):
        return f"ParentMessage(data={self.data})"


class MessageTag(Model):
    __slots__ = ("data", "user_id", "offset", "length", "type")

    def __init__(self, data: dict):
        self.data = data
        self.user_id = data.get("user_id")
        self.offset = data.get("offset")
        self.length = data.get("length")
        self.type = data.get("type")

    def __repr__(self):
        return f"MessageTag(data={self.data})"


class MuteKeyword(Model):
    __slots__ = ("data", "id", "word", "context")

    def __init__(self, data: dict):
        self.data = data
        self.id = data.get("id")
        self.word = data.get("word")
        self.context = data.get("context")

    def __repr__(self):
        return f"MuteKeyword(data={self.data})"


class PlatformDetails(Model):
    __slots__ = ("data", "package_id", "affiliate_url")

    def __init__(self, data: dict):
        self.data = data
        self.package_id = data.get("package_id")
        self.affiliate_url = data.get("affiliate_url")

    def __repr__(self):
        return f"PlatformDetails(data={self.data})"


class PopularWord(Model):
    __slots__ = ("data", "id", "word", "type")

    def __init__(self, data: dict):
        self.data = data
        self.id = data.get("id")
        self.word = data.get("word")
        self.type = data.get("type")

    def __repr__(self):
        return f"PopularWord(data={self.data})"


class Post(Model):
    __slots__ = (
        "data",
        "id",
        "text",
        "post_type",
        "group_id",
        "font_size",
        "color",
        "likes_count",
        "created_at",
        "updated_at",
        "edited_at",
        "liked",
        "tag",
        "reposts_count",
        "reposted",
        "repostable",
        "reported_count",
        "conversation_id",
        "in_reply_to",
        "in_reply_to_post",
        "in_reply_to_post_count",
        "user",
        "mentions",
        "group",
        "conference_call",
        "attachment",
        "attachment_thumbnail",
        "attachment_2",
        "attachment_2_thumbnail",
        "attachment_3",
        "attachment_3_thumbnail",
        "attachment_4",
        "attachment_4_thumbnail",
        "attachment_5",
        "attachment_5_thumbnail",
        "attachment_6",
        "attachment_6_thumbnail",
        "attachment_7",
        "attachment_7_thumbnail",
        "attachment_8",
        "attachment_8_thumbnail",
        "attachment_9",
        "attachment_9_thumbnail",
        "shareable",
        "shared_url",
        "survey",
        "videos",
        "gifts_count",
        "shared_thread",
        "thread_id",
        "thread",
        "highlighted",
        "message_tags",
        "is_fail_to_send",
    )

    def __init__(self, data: dict):
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
        self.tag = data.get("tag")
        self.reposts_count = data.get("reposts_count")
        self.reposted = data.get("reposted")
        self.repostable = data.get("repostable")
        self.reported_count = data.get("reported_count")
        self.conversation_id = data.get("conversation_id")
        self.in_reply_to = data.get("in_reply_to")
        self.in_reply_to_post = data.get("in_reply_to_post")
        self.in_reply_to_post_count = data.get("in_reply_to_post_count")

        self.user: User = data.get("user")
        if self.user is not None:
            self.user = User(self.user)

        self.mentions: List[User] = data.get("mentions")
        if self.mentions is not None:
            self.mentions = [User(mention) for mention in self.mentions]

        self.group: Group = data.get("group")
        if self.group is not None:
            self.group = Group(self.group)

        self.conference_call: ConferenceCall = data.get("conference_call")
        if self.conference_call is not None:
            self.conference_call = ConferenceCall(self.conference_call)

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

        self.shared_url: SharedUrl = data.get("shared_url")
        if self.shared_url is not None:
            self.shared_url = SharedUrl(self.shared_url)

        self.survey: Survey = data.get("survey")
        if self.survey is not None:
            self.survey = Survey(self.survey)

        self.videos: List[Video] = data.get("videos")
        if self.videos is not None:
            self.videos = [Video(video) for video in self.videos]

        self.gifts_count: List[GiftCount] = data.get("gifts_count")
        if self.gifts_count is not None:
            self.gifts_count = [
                GiftCount(gifts_count) for gifts_count in self.gifts_count
            ]

        self.shared_thread: List[ThreadInfo] = data.get("shared_thread")
        if self.shared_thread is not None:
            self.shared_thread = [
                ThreadInfo(shared_thread) for shared_thread in self.shared_thread
            ]

        self.thread_id = data.get("thread_id")

        self.thread: List[ThreadInfo] = data.get("thread")
        if self.thread is not None:
            self.thread = [ThreadInfo(thread) for thread in self.thread]

        self.highlighted = data.get("highlighted")

        self.message_tags: List[MessageTag] = data.get("message_tags")
        if self.message_tags is not None:
            self.message_tags = [
                MessageTag(message_tag) for message_tag in self.message_tags
            ]

        self.is_fail_to_send = data.get("is_fail_to_send")

    def __repr__(self):
        return f"Post(data={self.data})"


class PostGift(Model):
    __slots__ = ("data", "count", "gift")

    def __init__(self, data: dict):
        self.data = data
        self.count = data.get("count")

        self.gift: Gift = data.get("gift")
        if self.gift is not None:
            self.gift = Gift(self.gift)

    def __repr__(self):
        return f"PostGift(data={self.data})"


class PostTag(Model):
    __slots__ = ("data", "id", "tag", "post_hashtags_count")

    def __init__(self, data: dict):
        self.data = data
        self.id = data.get("id")
        self.tag = data.get("tag")
        self.post_hashtags_count = data.get("post_hashtags_count")

    def __repr__(self):
        return f"PostTag(data={self.data})"


class PresignedUrl(Model):
    __slots__ = ("data", "filename", "url")

    def __init__(self, data: dict):
        self.data = data
        self.filename = data.get("filename")
        self.url = data.get("url")

    def __repr__(self):
        return f"PresignedUrl(data={self.data})"


class Promotion(Model):
    __slots__ = ("data", "id", "title", "image_url", "promotion_url", "order")

    def __init__(self, data: dict):
        self.data = data
        self.id = data.get("id")
        self.title = data.get("title")
        self.image_url = data.get("image_url")
        self.promotion_url = data.get("promotion_url")
        self.order = data.get("order")

    def __repr__(self):
        return f"Promotion(data={self.data})"


class ReceivedGift(Model):
    __slots__ = ("data", "gift", "received_count", "senders", "total_senders_count")

    def __init__(self, data: dict):
        self.data = data

        self.gift: Gift = data.get("gift")
        if self.gift is not None:
            self.gift = Gift(self.gift)

        self.received_count = data.get("received_count")

        self.senders: List[User] = data.get("senders")
        if self.senders is not None:
            self.senders = [User(sender) for sender in self.senders]

        self.total_senders_count = data.get("total_senders_count")

    def __repr__(self):
        return f"ReceivedGift(data={self.data})"


class RecentSearch(Model):
    __slots__ = ("data", "id", "type", "user", "hashtag", "keyword")

    def __init__(self, data: dict):
        self.data = data
        self.id = data.get("id")
        self.type = data.get("type")

        self.user: User = data.get("user")
        if self.user is not None:
            self.user = User(self.user)

        self.hashtag: PostTag = data.get("hashtag")
        if self.hashtag is not None:
            self.hashtag = PostTag(self.hashtag)

        self.keyword = data.get("keyword")

    def __repr__(self):
        return f"RecentSearch(data={self.data})"


class RefreshCounterRequest(Model):
    __slots__ = ("data", "counter", "status", "last_requested_at")

    def __init__(self, data: dict):
        self.data = data
        self.counter = data.get("counter")
        self.status = data.get("status")
        self.last_requested_at = data.get("last_requested_at")

    def __repr__(self):
        return f"RefreshCounterRequest(data={self.data})"


class Review(Model):
    __slots__ = (
        "data",
        "id",
        "comment",
        "reported_count",
        "user",
        "reviewer",
        "created_at",
        "mutual_review",
    )

    def __init__(self, data: dict):
        self.data = data
        self.id = data.get("id")
        self.comment = data.get("comment")
        self.reported_count = data.get("reported_count")

        self.user: User = data.get("user")
        if self.user is not None:
            self.user = User(self.user)

        self.reviewer = data.get("reviewer")
        self.created_at = data.get("created_at")
        self.mutual_review = data.get("mutual_review")

    def __repr__(self):
        return f"Review(data={self.data})"


class SearchCriteria(Model):
    __slots__ = ("data", "nickname", "username", "biography", "prefecture", "gender")

    def __init__(self, data: dict):
        self.data = data
        self.nickname = data.get("nickname")
        self.username = data.get("username")
        self.biography = data.get("biography")
        self.prefecture = data.get("prefecture")
        self.gender = data.get("gender")

    def __repr__(self):
        return f"SearchCriteria(data={self.data})"


class Setting(Model):
    __slots__ = (
        "data",
        "notification_group_request",
        "notification_group_join",
        "notification_group_post",
        "notification_group_message_tag_all",
    )

    def __init__(self, data: dict):
        self.data = data
        self.notification_group_request = data.get("notification_group_request")
        self.notification_group_join = data.get("notification_group_join")
        self.notification_group_post = data.get("notification_group_post")
        self.notification_group_message_tag_all = data.get(
            "notification_group_message_tag_all"
        )

    def __repr__(self):
        return f"Setting(data={self.data})"


class Settings(Model):
    __slots__ = (
        "data",
        "notification_like",
        "notification_reply",
        "notification_repost",
        "notification_follow",
        "notification_chat",
        "notification_chat_delete",
        "notification_follow_request",
        "notification_message_tag",
        "notification_follow_accept",
        "notification_group_request",
        "notification_group_accept",
        "notification_group_join",
        "notification_group_post",
        "notification_group_invite",
        "notification_group_message_tag_all",
        "notification_profile_screenshot",
        "notification_following_birthdate_on",
        "notification_follower_create_group",
        "notification_follower_conference_call",
        "notification_group_conference_call",
        "notification_review",
        "notification_call_invite",
        "notification_bulk_call_invite",
        "notification_group_moderator",
        "notification_daily_summary",
        "notification_footprint",
        "notification_latest_news",
        "notification_popular_post",
        "notification_following_post_after_break",
        "notification_hima_now",
        "notification_birthday_to_followers",
        "notification_twitter_friend",
        "notification_contact_friend",
        "notification_security_warning",
        "notification_followings_in_call",
        "hide_active_call",
        "privacy_mode",
        "private_post",
        "private_user_timeline",
        "vip_invisible_footprint_mode",
        "allow_reposts",
        "invisible_on_user_search",
        "age_restricted_on_review",
        "hide_online_status",
        "following_restricted_on_review",
        "visible_on_sns_friend_recommendation",
        "following_only_call_invite",
        "following_only_group_invite",
        "caution_user_chat",
        "hide_vip",
        "hide_on_invitable",
        "hide_hot_post",
        "no_reply_public_timeline",
        "no_reply_following_timeline",
        "no_reply_group_timeline",
    )

    def __init__(self, data: dict):
        self.data = data
        self.notification_like = data.get("notification_like")
        self.notification_reply = data.get("notification_reply")
        self.notification_repost = data.get("notification_repost")
        self.notification_follow = data.get("notification_follow")
        self.notification_chat = data.get("notification_chat")
        self.notification_chat_delete = data.get("notification_chat_delete")
        self.notification_follow_request = data.get("notification_follow_request")
        self.notification_message_tag = data.get("notification_message_tag")
        self.notification_follow_accept = data.get("notification_follow_accept")
        self.notification_group_request = data.get("notification_group_request")
        self.notification_group_accept = data.get("notification_group_accept")
        self.notification_group_join = data.get("notification_group_join")
        self.notification_group_post = data.get("notification_group_post")
        self.notification_group_invite = data.get("notification_group_invite")
        self.notification_group_message_tag_all = data.get(
            "notification_group_message_tag_all"
        )
        self.notification_profile_screenshot = data.get(
            "notification_profile_screenshot"
        )
        self.notification_following_birthdate_on = data.get(
            "notification_following_birthdate_on"
        )
        self.notification_follower_create_group = data.get(
            "notification_follower_create_group"
        )
        self.notification_follower_conference_call = data.get(
            "notification_follower_conference_call"
        )
        self.notification_group_conference_call = data.get(
            "notification_group_conference_call"
        )
        self.notification_review = data.get("notification_review")
        self.notification_call_invite = data.get("notification_call_invite")
        self.notification_bulk_call_invite = data.get("notification_bulk_call_invite")
        self.notification_group_moderator = data.get("notification_group_moderator")
        self.notification_daily_summary = data.get("notification_daily_summary")
        self.notification_footprint = data.get("notification_footprint")
        self.notification_latest_news = data.get("notification_latest_news")
        self.notification_popular_post = data.get("notification_popular_post")
        self.notification_following_post_after_break = data.get(
            "notification_following_post_after_break"
        )
        self.notification_hima_now = data.get("notification_hima_now")
        self.notification_birthday_to_followers = data.get(
            "notification_birthday_to_followers"
        )
        self.notification_twitter_friend = data.get("notification_twitter_friend")
        self.notification_contact_friend = data.get("notification_contact_friend")
        self.notification_security_warning = data.get("notification_security_warning")
        self.notification_followings_in_call = data.get(
            "notification_followings_in_call"
        )
        self.hide_active_call = data.get("hide_active_call")
        self.privacy_mode = data.get("privacy_mode")
        self.private_post = data.get("private_post")
        self.private_user_timeline = data.get("private_user_timeline")
        self.vip_invisible_footprint_mode = data.get("vip_invisible_footprint_mode")
        self.allow_reposts = data.get("allow_reposts")
        self.invisible_on_user_search = data.get("invisible_on_user_search")
        self.age_restricted_on_review = data.get("age_restricted_on_review")
        self.hide_online_status = data.get("hide_online_status")
        self.following_restricted_on_review = data.get("following_restricted_on_review")
        self.visible_on_sns_friend_recommendation = data.get(
            "visible_on_sns_friend_recommendation"
        )
        self.following_only_call_invite = data.get("following_only_call_invite")
        self.following_only_group_invite = data.get("following_only_group_invite")
        self.caution_user_chat = data.get("caution_user_chat")
        self.hide_vip = data.get("hide_vip")
        self.hide_on_invitable = data.get("hide_on_invitable")
        self.hide_hot_post = data.get("hide_hot_post")
        self.no_reply_public_timeline = data.get("no_reply_public_timeline")
        self.no_reply_following_timeline = data.get("no_reply_following_timeline")
        self.no_reply_group_timeline = data.get("no_reply_group_timeline")

    def __repr__(self):
        return f"Settings(data={self.data})"


class Shareable(Model):
    __slots__ = ("data", "post", "group", "thread")

    def __init__(self, data: dict):
        self.data = data

        self.post: Post = data.get("post")
        if self.post is not None:
            self.post = Post(self.post)

        self.group: Group = data.get("group")
        if self.group is not None:
            self.group = Group(self.group)

        self.thread: ThreadInfo = data.get("thread")
        if self.thread is not None:
            self.thread = ThreadInfo(self.thread)

    def __repr__(self):
        return f"Shareable(data={self.data})"


class SharedUrl(Model):
    __slots__ = ("data", "url", "title", "description", "image_url")

    def __init__(self, data: dict):
        self.data = data
        self.url = data.get("url")
        self.title = data.get("title")
        self.description = data.get("description")
        self.image_url = data.get("image_url")

    def __repr__(self):
        return f"SharedUrl(data={self.data})"


class SNSInfo(Model):
    __slots__ = (
        "data",
        "type",
        "uid",
        "nickname",
        "biography",
        "profile_image",
        "gender",
    )

    def __init__(self, data: dict):
        self.data = data
        self.type = data.get("type")
        self.uid = data.get("uid")
        self.nickname = data.get("nickname")
        self.biography = data.get("biography")
        self.profile_image = data.get("profile_image")
        self.gender = data.get("gender")

    def __repr__(self):
        return f"SNSInfo(data={self.data})"


class Sticker(Model):
    __slots__ = ("data", "id", "sticker_pack_id", "width", "height", "url", "extension")

    def __init__(self, data: dict):
        self.data = data
        self.id = data.get("id")
        self.sticker_pack_id = data.get("sticker_pack_id")
        self.width = data.get("width")
        self.height = data.get("height")
        self.url = data.get("url")
        self.extension = data.get("extension")

    def __repr__(self):
        return f"Sticker(data={self.data})"


class StickerPack(Model):
    __slots__ = ("data", "id", "name", "description", "cover", "stickers", "order")

    def __init__(self, data: dict):
        self.data = data
        self.id = data.get("id")
        self.name = data.get("name")
        self.description = data.get("description")
        self.cover = data.get("cover")

        self.stickers: List[Sticker] = data.get("stickers")
        if self.stickers is not None:
            self.stickers = [Sticker(sticker) for sticker in self.stickers]

        self.order = data.get("order")

    def __repr__(self):
        return f"StickerPack(data={self.data})"


class Survey(Model):
    __slots__ = ("data", "id", "votes_count", "choices", "voted")

    def __init__(self, data: dict):
        self.data = data
        self.id = data.get("id")
        self.votes_count = data.get("votes_count")

        self.choices: List[Choice] = data.get("choices")
        if self.choices is not None:
            self.choices = [Choice(choice) for choice in self.choices]

        self.voted = data.get("voted")

    def __repr__(self):
        return f"Survey(data={self.data})"


class ThreadInfo(Model):
    __slots__ = (
        "data",
        "id",
        "title",
        "owner",
        "last_post",
        "unread_count",
        "posts_count",
        "created_at",
        "updated_at",
        "thread_icon",
        "is_joined",
        "new_updates",
    )

    def __init__(self, data: dict):
        self.data = data
        self.id = data.get("id")
        self.title = data.get("title")

        self.owner: User = data.get("owner")
        if self.owner is not None:
            self.owner = User(self.owner)

        self.last_post: Post = data.get("last_post")
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
        return f"ThreadInfo(data={self.data})"


class User(Model):
    # last_logged_in_at, created_at, updated_time_millis 確かめる

    __slots__ = (
        "data",
        "id",
        "nickname",
        "prefecture",
        "biography",
        "gender",
        "generation",
        "last_logged_in_at",
        "created_at",
        "badge",
        "followers_count",
        "followings_count",
        "posts_count",
        "groups_users_count",
        "reviews_count",
        "login_streak_count",
        "profile_icon",
        "profile_icon_thumbnail",
        "cover_image",
        "cover_image_thumbnail",
        "is_private",
        "is_vip",
        "is_vip_hidden",
        "is_chat_request_on",
        "mobile_number",
        "is_age_verified",
        "is_new_user",
        "online_status",
        "country_code",
        "is_recently_banned",
        "is_dangerous_user",
        "is_trusted_different_generation",
        "is_selected_interests",
        "group_user",
        "restricted_review_by",
        "is_following",
        "is_followed_by",
        "is_follow_pending",
        "is_hidden",
        "connected_by",
        "contact_phones",
        "updated_time_millis",
    )

    def __init__(self, data: dict):
        self.data = data
        self.id = data.get("id")
        self.nickname = data.get("nickname")
        self.prefecture = data.get("prefecture")
        self.biography = data.get("biography")
        self.gender = data.get("gender")
        self.generation = data.get("generation")
        self.last_logged_in_at = data.get("last_loggedin_at")
        self.created_at = data.get("created_at")
        self.badge = data.get("title")
        self.followers_count = data.get("followers_count")
        self.followings_count = data.get("followings_count")
        self.posts_count = data.get("posts_count")
        self.groups_users_count = data.get("groups_users_count")
        self.reviews_count = data.get("reviews_count")
        self.login_streak_count = data.get("login_streak_count")
        self.profile_icon = data.get("profile_icon")
        self.profile_icon_thumbnail = data.get("profile_icon_thumbnail")
        self.cover_image = data.get("cover_image")
        self.cover_image_thumbnail = data.get("cover_image_thumbnail")
        self.is_private = data.get("is_private")
        self.is_vip = data.get("vip")
        self.is_vip_hidden = data.get("hide_vip")
        self.is_chat_request_on = data.get("chat_request")
        self.mobile_number = data.get("mobile_number")
        self.is_age_verified = data.get("age_verified")
        self.is_new_user = data.get("new_user")
        self.online_status = data.get("online_status")
        self.country_code = data.get("country_code")
        self.is_recently_banned = data.get("recently_kenta")
        self.is_dangerous_user = data.get("dangerous_user")
        self.is_trusted_different_generation = data.get(
            "from_different_generation_and_trusted"
        )
        self.is_selected_interests = data.get("interests_selected")

        self.group_user: GroupUser = data.get("group_user")
        if self.group_user is not None:
            self.group_user = GroupUser(self.group_user)

        self.restricted_review_by = data.get("restricted_review_by")
        self.is_following = data.get("following")
        self.is_followed_by = data.get("followed_by")
        self.is_follow_pending = data.get("follow_pending")
        self.is_hidden = data.get("hidden")
        self.connected_by = data.get("connected_by")
        self.contact_phones = data.get("contact_phones")
        self.updated_time_millis = data.get("updated_time_millis")

    def __repr__(self):
        return f"User(data={self.data})"


class UserAuth(Model):
    __slots__ = ("data", "user_id", "access_token", "refresh_token", "expires_in")

    def __init__(self, data: dict):
        self.data = data
        self.user_id = data.get("user_id")
        self.access_token = data.get("access_token")
        self.refresh_token = data.get("refresh_token")
        self.expires_in = data.get("expires_in")

        self.user: User = data.get("user")
        if self.user is not None:
            self.user = User(self.user)

    def __repr__(self):
        return f"UserWrapper(data={self.data})"


class UserWrapper(Model):
    __slots__ = ("data", "id", "user")

    def __init__(self, data: dict):
        self.data = data
        self.id = data.get("id")

        self.user: User = data.get("user")
        if self.user is not None:
            self.user = User(self.user)

    def __repr__(self):
        return f"UserWrapper(data={self.data})"


class Video(Model):
    __slots__ = (
        "data",
        "id",
        "completed",
        "width",
        "height",
        "bitrate",
        "views_count",
        "video_url",
        "thumbnail_url",
        "thumbnail_big_url",
    )

    def __init__(self, data: dict):
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
        return f"Video(data={self.data})"


class Walkthrough(Model):
    __slots__ = ("data", "title", "url")

    def __init__(self, data: dict):
        self.data = data
        self.title = data.get("title")
        self.url = data.get("url")

    def __repr__(self):
        return f"Walkthrough(data={self.data})"


class WalletTransaction(Model):
    __slots__ = ("data", "id", "created_at", "description", "amount", "coins")

    def __init__(self, data: dict):
        self.data = data
        self.id = data.get("id")
        self.created_at = data.get("created_at")
        self.description = data.get("description")
        self.amount = data.get("amount")

        self.coins: CoinAmount = data.get("coins")
        if self.coins is not None:
            self.coins = CoinAmount(self.coins)

    def __repr__(self):
        return f"WalletTransaction(data={self.data})"


class WSIdentifier(Model):
    __slots__ = ("data", "channel")

    def __init__(self, data: dict) -> None:
        self.data = data
        self.channel: Optional[str] = data.get("channel")

    def __repr__(self):
        return f"WSIdentifier(data={self.data})"


class WSMessage(Model):
    __slots__ = ("data", "event", "message", "data")

    def __init__(self, data: dict) -> None:
        self.data = data
        self.event: Optional[str] = data.get("event")
        self.message: Optional[dict] = data.get("message")
        self.data: Optional[dict] = data.get("data")

    def __repr__(self):
        return f"WSMessage(data={self.data})"


class WSChannelMessage(Model):
    __slots__ = ("data", "type", "message", "identifier", "sid", "reason")

    def __init__(self, data: dict) -> None:
        self.data = data
        self.type: Optional[str] = data.get("type")

        self.message: Optional[WSMessage] = data.get("message")
        if self.message is not None and isinstance(self.message, dict):
            self.message = WSMessage(self.message)

        self.identifier: Optional[WSIdentifier] = data.get("identifier")
        if self.identifier is not None:
            self.identifier = WSIdentifier(json.loads(self.identifier))

        self.sid: Optional[str] = data.get("sid")
        self.reason: Optional[str] = data.get("reason")

    def __repr__(self):
        return f"WSChannelMessage(data={self.data})"


class SignaturePayload(Model):
    __slots__ = (
        "data",
        "action",
        "call_id",
        "receiver_uuid",
        "sender_uuid",
        "signature",
        "timestamp",
    )

    def __init__(self, data: dict) -> None:
        self.data = data
        self.action = data.get("action")
        self.call_id = data.get("call_id")
        self.receiver_uuid = data.get("receiver_uuid")
        self.sender_uuid = data.get("sender_uuid")
        self.signature = data.get("signature")
        self.timestamp = data.get("timestamp")

    def __repr__(self):
        return f"SignaturePayload(data={self.data})"
