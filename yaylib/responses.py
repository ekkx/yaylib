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

from typing import List

from .models import (
    Activity,
    ApplicationConfig,
    BanWord,
    Bgm,
    ChatRoom,
    ConferenceCall,
    CreateGroupQuota,
    Footprint,
    Game,
    Genre,
    GifImageCategory,
    GiftingAbility,
    Group,
    GroupCategory,
    GroupGiftHistory,
    GroupUser,
    Message,
    Model,
    PopularWord,
    Post,
    PostTag,
    PresignedUrl,
    Promotion,
    RefreshCounterRequest,
    Review,
    Setting,
    Settings,
    SignaturePayload,
    SNSInfo,
    StickerPack,
    Survey,
    ThreadInfo,
    TimelineSettings,
    User,
    UserWrapper,
)


class Response(Model):
    __slots__ = ("data", "result")

    def __init__(self, data: dict):
        self.data: dict = data
        self.result: str = data.get("result")

    def __repr__(self):
        return f"Response(data={self.data})"


class ErrorResponse(Response):
    __slots__ = ("message", "error_code", "ban_until", "retry_in")

    def __init__(self, data: dict):
        super().__init__(data)

        self.message = data.get("message")
        self.error_code = data.get("error_code")
        self.ban_until = data.get("ban_until")
        self.retry_in = data.get("retry_in")

    def __repr__(self):
        return f"ErrorResponse(data={self.data})"


class ActiveFollowingsResponse(Response):
    __slots__ = ("last_loggedin_at", "users")

    def __init__(self, data: dict):
        super().__init__(data)

        self.last_loggedin_at = data.get("last_loggedin_at")

        self.users: List[User] = data.get("users")
        if self.users is not None:
            self.users = [User(user) for user in self.users]

    def __repr__(self):
        return f"ActiveFollowingsResponse(data={self.data})"


class ActivitiesResponse(Response):
    __slots__ = ("activities", "last_timestamp")

    def __init__(self, data: dict):
        super().__init__(data)

        self.activities: List[Activity] = data.get("activities")
        if self.activities is not None:
            self.activities = [Activity(activity) for activity in self.activities]

        self.last_timestamp = data.get("last_timestamp")

    def __repr__(self):
        return f"ActivitiesResponse(data={self.data})"


class AdditionalSettingsResponse(Response):
    __slots__ = "settings"

    def __init__(self, data: dict):
        super().__init__(data)

        self.settings: Setting = data.get("settings")
        if self.settings is not None:
            self.settings = Settings(self.settings)

    def __repr__(self):
        return f"AdditionalSettingsResponse(data={self.data})"


class AppReviewStatusResponse(Response):
    __slots__ = "is_app_reviewed"

    def __init__(self, data: dict):
        super().__init__(data)

        self.is_app_reviewed = data.get("is_app_reviewed")

    def __repr__(self):
        return f"AppReviewStatusResponse(data={self.data})"


class BgmsResponse(Response):
    __slots__ = "bgm"

    def __init__(self, data: dict):
        super().__init__(data)

        self.bgm: List[Bgm] = data.get("bgm")
        if self.bgm is not None:
            self.bgm = [Bgm(bgm) for bgm in self.bgm]

    def __repr__(self):
        return f"BgmsResponse(data={self.data})"


class BlockedUserIdsResponse(Response):
    __slots__ = "block_ids"

    def __init__(self, data: dict):
        super().__init__(data)

        self.block_ids = data.get("block_ids")

    def __repr__(self):
        return f"BlockedUserIdsResponse(data={self.data})"


class BlockedUsersResponse(Response):
    __slots__ = ("blocked_count", "last_id", "users")

    def __init__(self, data: dict):
        super().__init__(data)

        self.blocked_count = data.get("blocked_count")
        self.last_id = data.get("last_id")

        self.users: List[User] = data.get("users")
        if self.users is not None:
            self.users = [User(user) for user in self.users]

    def __repr__(self):
        return f"BlockedUsersResponse(data={self.data})"


class BookmarkPostResponse(Response):
    __slots__ = "is_bookmarked"

    def __init__(self, data: dict):
        super().__init__(data)

        self.is_bookmarked = data.get("is_bookmarked")

    def __repr__(self):
        return f"BookmarkPostResponse(data={self.data})"


class CallStatusResponse(Response):
    __slots__ = ("phone_status", "video_status", "room_url")

    def __init__(self, data: dict):
        super().__init__(data)

        self.phone_status = data.get("phone_status")
        self.video_status = data.get("video_status")
        self.room_url = data.get("room_url")

    def __repr__(self):
        return f"CallStatusResponse(data={self.data})"


class ChatRoomResponse(Response):
    __slots__ = "chat"

    def __init__(self, data: dict):
        super().__init__(data)

        self.chat: ChatRoom = data.get("chat")
        if self.chat is not None:
            self.chat = ChatRoom(self.chat)

    def __repr__(self):
        return f"ChatRoomResponse(data={self.data})"


class ChatRoomsResponse(Response):
    __slots__ = ("pinned_chat_rooms", "chat_rooms", "next_page_value")

    def __init__(self, data: dict):
        super().__init__(data)

        self.pinned_chat_rooms: List[ChatRoom] = data.get("pinned_chat_rooms")
        if self.pinned_chat_rooms is not None:
            self.pinned_chat_rooms = [
                ChatRoom(pinned_chat_room)
                for pinned_chat_room in self.pinned_chat_rooms
            ]

        self.chat_rooms: List[ChatRoom] = data.get("chat_rooms")
        if self.chat_rooms is not None:
            self.chat_rooms = [ChatRoom(chat_room) for chat_room in self.chat_rooms]

        self.next_page_value = data.get("next_page_value")

    def __repr__(self):
        return f"ChatRoomsResponse(data={self.data})"


class TotalChatRequestResponse(Response):
    __slots__ = "total"

    def __init__(self, data: dict):
        super().__init__(data)

        self.total = data.get("total")

    def __repr__(self):
        return f"TotalChatRequestResponse(data={self.data})"


class ConferenceCallResponse(Response):
    __slots__ = "conference_call"

    def __init__(self, data: dict):
        super().__init__(data)

        self.conference_call: ConferenceCall = data.get("conference_call")
        if self.conference_call is not None:
            self.conference_call = ConferenceCall(self.conference_call)

    def __repr__(self):
        return f"ConferenceCallResponse(data={self.data})"


class ContactStatusResponse(Response):
    __slots__ = "contacts"

    def __init__(self, data: dict):
        super().__init__(data)

        self.contacts = data.get("contacts")

    def __repr__(self):
        return f"ContactStatusResponse(data={self.data})"


class CreateGroupResponse(Response):
    __slots__ = "group_id"

    def __init__(self, data: dict):
        super().__init__(data)

        self.group_id = data.get("group_id")

    def __repr__(self):
        return f"CreateGroupResponse(data={self.data})"


class CreateQuotaResponse(Response):
    __slots__ = "create"

    def __init__(self, data: dict):
        super().__init__(data)

        self.create: CreateGroupQuota = data.get("create")
        if self.create is not None:
            self.create = CreateGroupQuota(self.create)

    def __repr__(self):
        return f"CreateQuotaResponse(data={self.data})"


class CreateChatRoomResponse(Response):
    __slots__ = "room_id"

    def __init__(self, data: dict):
        super().__init__(data)

        self.room_id = data.get("room_id")

    def __repr__(self):
        return f"CreateChatRoomResponse(data={self.data})"


class CreatePostResponse(Response):
    __slots__ = ("conference_call", "post")

    def __init__(self, data: dict):
        super().__init__(data)

        self.conference_call: ConferenceCall = data.get("conference_call")
        if self.conference_call is not None:
            self.conference_call = ConferenceCall(self.conference_call)

        self.post = data.get("post")
        if self.post is not None:
            self.post = Post(self.post)

    def __repr__(self):
        return f"CreatePostResponse(data={self.data})"


class CreateUserResponse(Response):
    __slots__ = ("id", "access_token", "refresh_token", "expires_in")

    def __init__(self, data: dict):
        super().__init__(data)

        self.id = data.get("id")
        self.access_token = data.get("access_token")
        self.refresh_token = data.get("refresh_token")
        self.expires_in = data.get("expires_in")

    def __repr__(self):
        return f"CreateUserResponse(data={self.data})"


class EmailGrantTokenResponse(Response):
    __slots__ = "email_grant_token"

    def __init__(self, data: dict):
        super().__init__(data)

        self.email_grant_token = data.get("email_grant_token")

    def __repr__(self):
        return f"EmailGrantTokenResponse(data={self.data})"


class EmailVerificationPresignedUrlResponse(Response):
    __slots__ = ("url", "presigned_url", "expires_at", "method_type")

    def __init__(self, data: dict):
        super().__init__(data)

        self.url = data.get("url")
        self.presigned_url = data.get("presigned_url")
        self.expires_at = data.get("expires_at")
        self.method_type = data.get("method_type")

    def __repr__(self):
        return f"EmailVerificationPresignedUrlResponse(data={self.data})"


class PresignedUrlResponse(Response):
    __slots__ = "presigned_url"

    def __init__(self, data: dict):
        super().__init__(data)

        self.presigned_url = data.get("presigned_url")

    def __repr__(self):
        return f"PresignedUrlResponse(data={self.data})"


class PresignedUrlsResponse(Response):
    __slots__ = "presigned_urls"

    def __init__(self, data: dict):
        super().__init__(data)

        self.presigned_urls: List[PresignedUrl] = data.get("presigned_urls")
        if self.presigned_urls is not None:
            self.presigned_urls = [
                PresignedUrl(presigned_url) for presigned_url in self.presigned_urls
            ]

    def __repr__(self):
        return f"PresignedUrlsResponse(data={self.data})"


class IdCheckerPresignedUrlResponse(Response):
    __slots__ = "presigned_url"

    def __init__(self, data: dict):
        super().__init__(data)

        self.presigned_url = data.get("presigned_url")

    def __repr__(self):
        return f"IdCheckerPresignedUrlResponse(data={self.data})"


class DefaultSettingsResponse(Response):
    __slots__ = "timeline_settings"

    def __init__(self, data: dict):
        super().__init__(data)

        self.timeline_settings: TimelineSettings = data.get("timeline_settings")
        if self.timeline_settings is not None:
            self.timeline_settings = TimelineSettings(self.timeline_settings)

    def __repr__(self):
        return f"DefaultSettingsResponse(data={self.data})"


class FollowRecommendationsResponse(Response):
    __slots__ = ("total", "users", "next")

    def __init__(self, data: dict):
        super().__init__(data)

        self.total = data.get("total")

        self.users: List[User] = data.get("users")
        if self.users is not None:
            self.users = [User(user) for user in self.users]

        self.next = data.get("next")

    def __repr__(self):
        return f"FollowRecommendationsResponse(data={self.data})"


class FollowRequestCountResponse(Response):
    __slots__ = "users_count"

    def __init__(self, data: dict):
        super().__init__(data)

        self.users_count = data.get("users_count")

    def __repr__(self):
        return f"FollowRequestCountResponse(data={self.data})"


class FollowUsersResponse(Response):
    __slots__ = ("last_follow_id", "users")

    def __init__(self, data: dict):
        super().__init__(data)

        self.last_follow_id = data.get("last_follow_id")

        self.users: List[User] = data.get("users")
        if self.users is not None:
            self.users = [User(user) for user in self.users]

    def __repr__(self):
        return f"FollowUsersResponse(data={self.data})"


class FootprintsResponse(Response):
    __slots__ = "footprints"

    def __init__(self, data: dict):
        super().__init__(data)

        self.footprints: List[Footprint] = data.get("footprints")
        if self.footprints is not None:
            self.footprints = [Footprint(footprint) for footprint in self.footprints]

    def __repr__(self):
        return f"FootprintsResponse(data={self.data})"


class GamesResponse(Response):
    __slots__ = ("games", "from_id")

    def __init__(self, data: dict):
        super().__init__(data)

        self.games: List[Game] = data.get("games")
        if self.games is not None:
            self.games = [Game(game) for game in self.games]

        self.from_id = data.get("from_id")

    def __repr__(self):
        return f"GamesResponse(data={self.data})"


class GenresResponse(Response):
    __slots__ = ("genres", "next_page_value")

    def __init__(self, data: dict):
        super().__init__(data)

        self.genres: List[Genre] = data.get("genres")
        if self.genres is not None:
            self.genres = [Genre(genre) for genre in self.genres]

        self.next_page_value = data.get("next_page_value")

    def __repr__(self):
        return f"GenresResponse(data={self.data})"


class GroupCategoriesResponse(Response):
    __slots__ = "group_categories"

    def __init__(self, data: dict):
        super().__init__(data)

        self.group_categories: List[GroupCategory] = data.get("group_categories")
        if self.group_categories is not None:
            self.group_categories = [
                GroupCategory(group_categorie)
                for group_categorie in self.group_categories
            ]

    def __repr__(self):
        return f"GroupCategoriesResponse(data={self.data})"


class GroupGiftHistoryResponse(Response):
    __slots__ = ("gift_history", "next_page_value")

    def __init__(self, data: dict):
        super().__init__(data)

        self.gift_history: List[GroupGiftHistory] = data.get("gift_history")
        if self.gift_history is not None:
            self.gift_history = [
                GroupGiftHistory(gift_history) for gift_history in self.gift_history
            ]

        self.next_page_value = data.get("next_page_value")

    def __repr__(self):
        return f"GroupGiftHistoryResponse(data={self.data})"


class GroupNotificationSettingsResponse(Response):
    __slots__ = "setting"

    def __init__(self, data: dict):
        super().__init__(data)

        self.setting: Setting = data.get("setting")
        if self.setting is not None:
            self.setting = Setting(self.setting)

    def __repr__(self):
        return f"GroupNotificationSettingsResponse(data={self.data})"


class GroupResponse(Response):
    __slots__ = "group"

    def __init__(self, data: dict):
        super().__init__(data)

        self.group: Group = data.get("group")
        if self.group is not None:
            self.group = Group(self.group)

    def __repr__(self):
        return f"GroupResponse(data={self.data})"


class GroupsRelatedResponse(Response):
    __slots__ = ("groups", "next_page_value")

    def __init__(self, data: dict):
        super().__init__(data)

        self.groups: List[Group] = data.get("groups")
        if self.groups is not None:
            self.groups = [Group(group) for group in self.groups]

        self.next_page_value = data.get("next_page_value")

    def __repr__(self):
        return f"GroupsRelatedResponse(data={self.data})"


class GroupsResponse(Response):
    __slots__ = ("pinned_groups", "groups")

    def __init__(self, data: dict):
        super().__init__(data)

        self.pinned_groups: List[Group] = data.get("pinned_groups")
        if self.pinned_groups is not None:
            self.pinned_groups = [
                Group(pinned_group) for pinned_group in self.pinned_groups
            ]

        self.groups: List[Group] = data.get("groups")
        if self.groups is not None:
            self.groups = [Group(group) for group in self.groups]

    def __repr__(self):
        return f"GroupsResponse(data={self.data})"


class GroupThreadListResponse(Response):
    __slots__ = ("threads", "next_page_value")

    def __init__(self, data: dict):
        super().__init__(data)

        self.threads: List[ThreadInfo] = data.get("threads")
        if self.threads is not None:
            self.threads = [ThreadInfo(thread) for thread in self.threads]

        self.next_page_value = data.get("next_page_value")

    def __repr__(self):
        return f"GroupThreadListResponse(data={self.data})"


class GroupUserResponse(Response):
    __slots__ = "group_user"

    def __init__(self, data: dict):
        super().__init__(data)

        self.group_user: GroupUser = data.get("group_user")
        if self.group_user is not None:
            self.group_user = GroupUser(self.group_user)

    def __repr__(self):
        return f"GroupUserResponse(data={self.data})"


class GroupUsersResponse(Response):
    __slots__ = "group_users"

    def __init__(self, data: dict):
        super().__init__(data)

        self.group_users: List[GroupUser] = data.get("group_users")
        if self.group_users is not None:
            self.group_users = [
                GroupUser(group_user) for group_user in self.group_users
            ]

    def __repr__(self):
        return f"GroupUsersResponse(data={self.data})"


class GifsDataResponse(Response):
    __slots__ = "gif_categories"

    def __init__(self, data: dict):
        super().__init__(data)

        self.gif_categories: List[GifImageCategory] = data.get("gif_categories")
        if self.gif_categories is not None:
            self.gif_categories = [
                GifImageCategory(gif_category) for gif_category in self.gif_categories
            ]

    def __repr__(self):
        return f"GifsDataResponse(data={self.data})"


class HiddenResponse(Response):
    __slots__ = ("hidden_users", "next_page_value", "total_count", "limit")

    def __init__(self, data: dict):
        super().__init__(data)

        self.hidden_users: List[User] = data.get("hidden_users")
        if self.hidden_users is not None:
            self.hidden_users = [User(hidden_user) for hidden_user in self.hidden_users]

        self.next_page_value = data.get("next_page_value")
        self.total_count = data.get("total_count")
        self.limit = data.get("limit")

    def __repr__(self):
        return f"HiddenResponse(data={self.data})"


class LoginUserResponse(Response):
    __slots__ = (
        "data",
        "user_id",
        "username",
        "is_new",
        "sns_info",
        "access_token",
        "refresh_token",
        "expires_in",
    )

    def __init__(self, data: dict):
        super().__init__(data)

        self.user_id = data.get("user_id")
        self.username = data.get("username")
        self.is_new = data.get("is_new")

        self.sns_info: SNSInfo = data.get("sns_info")
        if self.sns_info is not None:
            self.sns_info = SNSInfo(self.sns_info)

        self.access_token = data.get("access_token")
        self.refresh_token = data.get("refresh_token")
        self.expires_in = data.get("expires_in")

    def __repr__(self):
        return f"LoginUserResponse(data={self.data})"


class LoginUpdateResponse(Response):
    __slots__ = (
        "data",
        "user_id",
        "username",
        "access_token",
        "refresh_token",
        "expires_in",
    )

    def __init__(self, data: dict):
        super().__init__(data)

        self.user_id = data.get("user_id")
        self.username = data.get("username")
        self.access_token = data.get("access_token")
        self.refresh_token = data.get("refresh_token")
        self.expires_in = data.get("expires_in")

    def __repr__(self):
        return f"LoginUpdateResponse(data={self.data})"


class MessageResponse(Response):
    __slots__ = ("id", "conference_call")

    def __init__(self, data: dict):
        super().__init__(data)

        self.id = data.get("id")

        self.conference_call: ConferenceCall = data.get("conference_call")
        if self.conference_call is not None:
            self.conference_call = ConferenceCall(self.conference_call)

    def __repr__(self):
        return f"MessageResponse(data={self.data})"


class MessagesResponse(Response):
    __slots__ = "messages"

    def __init__(self, data: dict):
        super().__init__(data)

        self.messages: List[Message] = data.get("messages")
        if self.messages is not None:
            self.messages = [Message(message) for message in self.messages]

    def __repr__(self):
        return f"MessagesResponse(data={self.data})"


class PolicyAgreementsResponse(Response):
    __slots__ = ("latest_privacy_policy_agreed", "latest_terms_of_use_agreed")

    def __init__(self, data: dict):
        super().__init__(data)

        self.latest_privacy_policy_agreed = data.get("latest_privacy_policy_agreed")
        self.latest_terms_of_use_agreed = data.get("latest_terms_of_use_agreed")

    def __repr__(self):
        return f"PolicyAgreementsResponse(data={self.data})"


class PostResponse(Response):
    __slots__ = "post"

    def __init__(self, data: dict):
        super().__init__(data)

        self.post: Post = data.get("post")
        if self.post is not None:
            self.post = Post(self.post)

    def __repr__(self):
        return f"PostResponse(data={self.data})"


class PostsResponse(Response):
    __slots__ = ("next_page_value", "posts", "pinned_posts")

    def __init__(self, data: dict):
        super().__init__(data)

        self.next_page_value = data.get("next_page_value")

        self.posts: List[Post] = data.get("posts")
        if self.posts is not None:
            self.posts = [Post(post) for post in self.posts]

        self.pinned_posts: List[Post] = data.get("pinned_posts")
        if self.pinned_posts is not None:
            self.pinned_posts = [Post(pinned_post) for pinned_post in self.pinned_posts]

    def __repr__(self):
        return f"PostsResponse(data={self.data})"


class PostLikersResponse(Response):
    __slots__ = ("last_id", "users")

    def __init__(self, data: dict):
        super().__init__(data)

        self.last_id = data.get("last_id")

        self.users: List[User] = data.get("users")
        if self.users is not None:
            self.users = [User(user) for user in self.users]

    def __repr__(self):
        return f"PostLikersResponse(data={self.data})"


class PostTagsResponse(Response):
    __slots__ = "tags"

    def __init__(self, data: dict):
        super().__init__(data)

        self.tags: List[PostTag] = data.get("tags")
        if self.tags is not None:
            self.tags = [PostTag(tag) for tag in self.tags]

    def __repr__(self):
        return f"PostTagsResponse(data={self.data})"


class PromotionsResponse(Response):
    __slots__ = "promotions"

    def __init__(self, data: dict):
        super().__init__(data)

        self.promotions: List[Promotion] = data.get("promotions")
        if self.promotions is not None:
            self.promotions = [Promotion(promotion) for promotion in self.promotions]

    def __repr__(self):
        return f"PromotionsResponse(data={self.data})"


class LikePostsResponse(Response):
    __slots__ = "like_ids"

    def __init__(self, data: dict):
        super().__init__(data)

        self.like_ids = data.get("like_ids")

    def __repr__(self):
        return f"LikePostsResponse(data={self.data})"


class ValidationPostResponse(Response):
    __slots__ = "is_allow_to_post"

    def __init__(self, data: dict):
        super().__init__(data)

        self.is_allow_to_post = data.get("is_allow_to_post")

    def __repr__(self):
        return f"ValidationPostResponse(data={self.data})"


class RefreshCounterRequestsResponse(Response):
    __slots__ = "reset_counter_requests"

    def __init__(self, data: dict):
        super().__init__(data)

        self.reset_counter_requests: List[RefreshCounterRequest] = data.get(
            "reset_counter_requests"
        )
        if self.reset_counter_requests is not None:
            self.reset_counter_requests = [
                RefreshCounterRequest(reset_counter_request)
                for reset_counter_request in self.reset_counter_requests
            ]

    def __repr__(self):
        return f"RefreshCounterRequestsResponse(data={self.data})"


class ReviewsResponse(Response):
    __slots__ = ("reviews", "pinned_reviews")

    def __init__(self, data: dict):
        super().__init__(data)

        self.reviews: List[Review] = data.get("reviews")
        if self.reviews is not None:
            self.reviews = [Review(review) for review in self.reviews]

        self.pinned_reviews: List[Review] = data.get("pinned_reviews")
        if self.pinned_reviews is not None:
            self.pinned_reviews = [
                Review(pinned_review) for pinned_review in self.pinned_reviews
            ]

    def __repr__(self):
        return f"ReviewsResponse(data={self.data})"


class SocialShareUsersResponse(Response):
    __slots__ = "social_shared_users"

    def __init__(self, data: dict):
        super().__init__(data)

        self.social_shared_users: List[UserWrapper] = data.get("social_shared_users")
        if self.social_shared_users is not None:
            self.social_shared_users = [
                UserWrapper(social_shared_user)
                for social_shared_user in self.social_shared_users
            ]

    def __repr__(self):
        return f"SocialShareUsersResponse(data={self.data})"


class StickerPacksResponse(Response):
    __slots__ = "sticker_packs"

    def __init__(self, data: dict):
        super().__init__(data)

        self.sticker_packs: List[StickerPack] = data.get("sticker_packs")
        if self.sticker_packs is not None:
            self.sticker_packs = [
                StickerPack(sticker_pack) for sticker_pack in self.sticker_packs
            ]

    def __repr__(self):
        return f"StickerPacksResponse(data={self.data})"


class TokenResponse(Response):
    __slots__ = (
        "data",
        "user_id",
        "created_at",
        "access_token",
        "refresh_token",
        "expires_in",
    )

    def __init__(self, data: dict):
        super().__init__(data)

        self.user_id = data.get("id")
        self.created_at = data.get("created_at")
        self.access_token = data.get("access_token")
        self.refresh_token = data.get("refresh_token")
        self.expires_in = data.get("expires_in")

    def __repr__(self):
        return f"TokenResponse(data={self.data})"


class VipGameRewardUrlResponse(Response):
    __slots__ = "url"

    def __init__(self, data: dict):
        super().__init__(data)

        self.url = data.get("url")

    def __repr__(self):
        return f"VipGameRewardUrlResponse(data={self.data})"


class VoteSurveyResponse(Response):
    __slots__ = "survey"

    def __init__(self, data: dict):
        super().__init__(data)

        self.survey: Survey = data.get("survey")
        if self.survey is not None:
            self.survey = Survey(self.survey)

    def __repr__(self):
        return f"VoteSurveyResponse(data={self.data})"


class UnreadStatusResponse(Response):
    __slots__ = "is_unread"

    def __init__(self, data: dict):
        super().__init__(data)

        self.is_unread = data.get("is_unread")

    def __repr__(self):
        return f"UnreadStatusResponse(data={self.data})"


class UserResponse(Response):
    __slots__ = (
        "user",
        "masked_email",
        "twitter_id",
        "is_line_connected",
        "is_facebook_connected",
        "is_lobi_connected",
        "is_push_notification_on",
        "is_call_on",
        "is_video_on",
        "is_group_call_on",
        "is_group_video_on",
        "vip_until",
        "is_email_confirmed",
        "blocking_limit",
        "uuid",
        "birthdate",
        "gifting_ability",
    )

    def __init__(self, data: dict):
        super().__init__(data)

        self.user: User = data.get("user")
        if self.user is not None:
            self.user = User(self.user)

        self.masked_email = data.get("masked_email")
        self.twitter_id = data.get("twitter_id")
        self.is_line_connected = data.get("line_connected")
        self.is_facebook_connected = data.get("facebook_connected")
        self.is_lobi_connected = data.get("lobi_connected")
        self.is_push_notification_on = data.get("push_notification")
        self.is_call_on = data.get("phone_on")
        self.is_video_on = data.get("video_on")
        self.is_group_call_on = data.get("group_phone_on")
        self.is_group_video_on = data.get("group_video_on")
        self.vip_until = data.get("vip_until_seconds")
        self.is_email_confirmed = data.get("email_confirmed")
        self.blocking_limit = data.get("blocking_limit")
        self.uuid = data.get("uuid")
        self.birthdate = data.get("birth_date")

        self.gifting_ability: GiftingAbility = data.get("gifting_ability")
        if self.gifting_ability is not None:
            self.gifting_ability = GiftingAbility(self.gifting_ability)

    def __repr__(self):
        return f"UserResponse(data={self.data})"


class UsersResponse(Response):
    __slots__ = ("users", "next_page_value")

    def __init__(self, data: dict):
        super().__init__(data)

        self.users: List[User] = data.get("users")
        if self.users is not None:
            self.users = [User(user) for user in self.users]

        self.next_page_value = data.get("next_page_value")

    def __repr__(self):
        return f"UsersResponse(data={self.data})"


class RankingUsersResponse(Response):
    __slots__ = "users"

    def __init__(self, data: dict):
        super().__init__(data)

        self.users: List[User] = data.get("users")
        if self.users is not None:
            self.users = [User(user) for user in self.users]

    def __repr__(self):
        return f"RankingUsersResponse(data={self.data})"


class UserCustomDefinitionsResponse(Response):
    __slots__ = (
        "age",
        "followers_count",
        "followings_count",
        "created_at",
        "last_loggedin_at",
        "status",
        "reported_count",
    )

    def __init__(self, data: dict):
        super().__init__(data)

        self.age = data.get("age")
        self.followers_count = data.get("followers_count")
        self.followings_count = data.get("followings_count")
        self.created_at = data.get("created_at")
        self.last_loggedin_at = data.get("last_loggedin_at")
        self.status = data.get("status")
        self.reported_count = data.get("reported_count")

    def __repr__(self):
        return f"UserCustomDefinitionsResponse(data={self.data})"


class HimaUsersResponse(Response):
    __slots__ = "hima_users"

    def __init__(self, data: dict):
        super().__init__(data)

        self.hima_users: List[UserWrapper] = data.get("hima_users")
        if self.hima_users is not None:
            self.hima_users = [UserWrapper(hima_user) for hima_user in self.hima_users]

    def __repr__(self):
        return f"HimaUsersResponse(data={self.data})"


class UsersByTimestampResponse(Response):
    __slots__ = ("last_timestamp", "users")

    def __init__(self, data: dict):
        super().__init__(data)

        self.last_timestamp = data.get("last_timestamp")

        self.users: List[User] = data.get("users")
        if self.users is not None:
            self.users = [User(user) for user in self.users]

    def __repr__(self):
        return f"UsersByTimestampResponse(data={self.data})"


class UserTimestampResponse(Response):
    __slots__ = ("time", "ip_address", "country")

    def __init__(self, data: dict):
        super().__init__(data)

        self.time = data.get("time")
        self.ip_address = data.get("ip_address")
        self.country = data.get("country")

    def __repr__(self):
        return f"UserTimestampResponse(data={self.data})"


class WebSocketTokenResponse(Response):
    __slots__ = "token"

    def __init__(self, data: dict):
        super().__init__(data)

        self.token = data.get("token")

    def __repr__(self):
        return f"WebSocketTokenResponse(data={self.data})"


class ApplicationConfigResponse(Response):
    __slots__ = "app"

    def __init__(self, data: dict):
        super().__init__(data)

        self.app: ApplicationConfig = data.get("app")
        if self.app is not None:
            self.app = ApplicationConfig(self.app)

    def __repr__(self):
        return f"ApplicationConfigResponse(data={self.data})"


class BanWordsResponse(Response):
    __slots__ = "ban_words"

    def __init__(self, data: dict):
        super().__init__(data)

        self.ban_words: List[BanWord] = data.get("ban_words")
        if self.ban_words is not None:
            self.ban_words = [BanWord(ban_word) for ban_word in self.ban_words]

    def __repr__(self):
        return f"BanWordsResponse(data={self.data})"


class PopularWordsResponse(Response):
    __slots__ = "popular_words"

    def __init__(self, data: dict):
        super().__init__(data)

        self.popular_words: List[PopularWord] = data.get("popular_words")
        if self.popular_words is not None:
            self.popular_words = [
                PopularWord(popular_word) for popular_word in self.popular_words
            ]

    def __repr__(self):
        return f"PopularWordsResponse(data={self.data})"


class CallActionSignatureResponse(Response):
    __slots__ = "signature_payload"

    def __init__(self, data: dict):
        super().__init__(data)

        self.signature_payload: SignaturePayload = data.get("signature_payload")
        if self.signature_payload is not None:
            self.signature_payload = SignaturePayload(self.signature_payload)

    def __repr__(self):
        return f"CallActionSignatureResponse(data={self.data})"
