from .models import *


class ActiveFollowingsResponse:

    __slots__ = ("data", "last_loggedin_at", "users")

    def __init__(self, data):
        self.data = data
        self.last_loggedin_at = data.get("last_loggedin_at")

        self.users = data.get("users")
        if self.users is not None:
            self.users = [
                User(user) for user in self.users
            ]

    def __repr__(self):
        return f"ActiveFollowingsResponse(data={self.data})"


class AdditionalSettingsResponse:

    __slots__ = ("data", "settings")

    def __init__(self, data):
        self.data = data

        self.settings = data.get("settings")
        if self.settings is not None:
            self.settings = Settings(self.settings)

    def __repr__(self):
        return f"AdditionalSettingsResponse(data={self.data})"


class AppReviewStatusResponse:

    __slots__ = ("data", "is_app_reviewed")

    def __init__(self, data):
        self.data = data
        self.is_app_reviewed = data.get("is_app_reviewed")

    def __repr__(self):
        return f"AppReviewStatusResponse(data={self.data})"


class BgmsResponse:

    __slots__ = ("data", "bgm")

    def __init__(self, data):
        self.data = data

        self.bgm = data.get("bgm")
        if self.bgm is not None:
            self.bgm = [
                Bgm(bgm) for bgm in self.bgm
            ]

    def __repr__(self):
        return f"BgmsResponse(data={self.data})"


class BookmarkPostResponse:

    __slots__ = ("data", "is_bookmarked")

    def __init__(self, data):
        self.data = data
        self.is_bookmarked = data.get("is_bookmarked")

    def __repr__(self):
        return f"BookmarkPostResponse(data={self.data})"


class CallStatusResponse:

    __slots__ = ("data", "phone_status", "video_status", "room_url")

    def __init__(self, data):
        self.data = data
        self.phone_status = data.get("phone_status")
        self.video_status = data.get("video_status")
        self.room_url = data.get("room_url")

    def __repr__(self):
        return f"CallStatusResponse(data={self.data})"


class ChatRoomsResponse:

    __slots__ = ("data", "pinned_chat_rooms", "chat_rooms", "next_page_value")

    def __init__(self, data):
        self.data = data

        self.pinned_chat_rooms = data.get("pinned_chat_rooms")
        if self.pinned_chat_rooms is not None:
            self.pinned_chat_rooms = [
                ChatRoom(pinned_chat_room) for pinned_chat_room in self.pinned_chat_rooms
            ]

        self.chat_rooms = data.get("chat_rooms")
        if self.chat_rooms is not None:
            self.chat_rooms = [
                ChatRoom(chat_room) for chat_room in self.chat_rooms
            ]

        self.next_page_value = data.get("next_page_value")

    def __repr__(self):
        return f"ChatRoomsResponse(data={self.data})"


class ConferenceCallResponse:

    __slots__ = ("data", "conference_call")

    def __init__(self, data):
        self.data = data

        self.conference_call = data.get("conference_call")
        if self.conference_call is not None:
            self.conference_call = ConferenceCall(self.conference_call)

    def __repr__(self):
        return f"ConferenceCallResponse(data={self.data})"


class ContactStatusResponse:

    __slots__ = ("data", "contacts")

    def __init__(self, data):
        self.data = data
        self.contacts = data.get("contacts")

    def __repr__(self):
        return f"ContactStatusResponse(data={self.data})"


class CreatePostResponse:

    __slots__ = ("data", "conference_call", "post")

    def __init__(self, data):
        self.data = data

        self.conference_call = data.get("conference_call")
        if self.conference_call is not None:
            self.conference_call = ConferenceCall(self.conference_call)

        self.post = data.get("post")
        if self.post is not None:
            self.post = Post(self.post)

    def __repr__(self):
        return f"CreatePostResponse(data={self.data})"


class CreateUserResponse:

    __slots__ = ("data", "id", "access_token", "refresh_token", "expires_in")

    def __init__(self, data):
        self.data = data
        self.id = data.get("id")
        self.access_token = data.get("access_token")
        self.refresh_token = data.get("refresh_token")
        self.expires_in = data.get("expires_in")

    def __repr__(self):
        return f"CreateUserResponse(data={self.data})"


class EmailVerificationPresignedUrlResponse:

    __slots__ = ("data", "url")

    def __init__(self, data):
        self.data = data
        self.url = data.get("url")

    def __repr__(self):
        return f"EmailVerificationPresignedUrlResponse(data={self.data})"


class DefaultSettingsResponse:

    __slots__ = ("data", "timeline_settings")

    def __init__(self, data):
        self.data = data

        self.timeline_settings = data.get("timeline_settings")
        if self.timeline_settings is not None:
            self.timeline_settings = TimelineSettings(self.timeline_settings)

    def __repr__(self):
        return f"DefaultSettingsResponse(data={self.data})"


class FollowRecommendationsResponse:

    __slots__ = ("data", "total", "users", "next")

    def __init__(self, data):
        self.data = data
        self.total = data.get("total")

        self.users = data.get("users")
        if self.users is not None:
            self.users = [
                User(user) for user in self.users
            ]

        self.next = data.get("next")

    def __repr__(self):
        return f"FollowRecommendationsResponse(data={self.data})"


class FollowRequestCountResponse:

    __slots__ = ("data", "users_count")

    def __init__(self, data):
        self.data = data
        self.users_count = data.get("users_count")

    def __repr__(self):
        return f"FollowRequestCountResponse(data={self.data})"


class FollowUsersResponse:

    __slots__ = ("data", "last_follow_id", "users")

    def __init__(self, data):
        self.data = data
        self.last_follow_id = data.get("last_follow_id")

        self.users = data.get("users")
        if self.users is not None:
            self.users = [
                User(user) for user in self.users
            ]

    def __repr__(self):
        return f"FollowUsersResponse(data={self.data})"


class FootprintsResponse:

    __slots__ = ("data", "footprints")

    def __init__(self, data):
        self.data = data

        self.footprints = data.get("footprints")
        if self.footprints is not None:
            self.footprints = [
                Footprint(footprint) for footprint in self.footprints
            ]

    def __repr__(self):
        return f"FootprintsResponse(data={self.data})"


class GamesResponse:

    __slots__ = ("data", "games", "from_id")

    def __init__(self, data):
        self.data = data

        self.games = data.get("games")
        if self.games is not None:
            self.games = [
                Game(game) for game in self.games
            ]

        self.from_id = data.get("from_id")

    def __repr__(self):
        return f"GamesResponse(data={self.data})"


class GenresResponse:

    __slots__ = ("data", "genres", "next_page_value")

    def __init__(self, data):
        self.data = data

        self.genres = data.get("genres")
        if self.genres is not None:
            self.genres = [
                Genre(genre) for genre in self.genres
            ]

        self.next_page_value = data.get("next_page_value")

    def __repr__(self):
        return f"GenresResponse(data={self.data})"


class MessageResponse:

    __slots__ = ("data", "id", "conferenceCall")

    def __init__(self, data):
        self.data = data
        self.id = data.get("id")

        self.conference_call = data.get("conference_call")
        if self.conference_call is not None:
            self.conference_call = ConferenceCall(self.conference_call)

    def __repr__(self):
        return f"MessageResponse(data={self.data})"


class PostResponse:

    __slots__ = ("data", "post")

    def __init__(self, data):
        self.data = data

        self.post = data.get("post")
        if self.post is not None:
            self.post = Post(self.post)

    def __repr__(self):
        return f"PostResponse(data={self.data})"


class PostsResponse:

    __slots__ = ("data", "next_page_value", "posts", "pinned_posts")

    def __init__(self, data):
        self.data = data
        self.next_page_value = data.get("next_page_value")

        self.posts = data.get("posts")
        if self.posts is not None:
            self.posts = [
                Post(post) for post in self.posts
            ]

        self.pinned_posts = data.get("pinned_posts")
        if self.pinned_posts is not None:
            self.pinned_posts = [
                Post(pinned_post) for pinned_post in self.pinned_posts
            ]

    def __repr__(self):
        return f"PostsResponse(data={self.data})"


class PostLikersResponse:

    __slots__ = ("data", "last_id", "users")

    def __init__(self, data):
        self.data = data
        self.last_id = data.get("last_id")

        self.users = data.get("users")
        if self.users is not None:
            self.users = [
                User(user) for user in self.users
            ]

    def __repr__(self):
        return f"PostLikersResponse(data={self.data})"


class PostTagsResponse:

    __slots__ = ("data", "tags")

    def __init__(self, data):
        self.data = data

        self.tags = data.get("tags")
        if self.tags is not None:
            self.tags = [
                PostTag(tag) for tag in self.tags
            ]

    def __repr__(self):
        return f"PostTagsResponse(data={self.data})"


class LikePostsResponse:

    __slots__ = ("data", "like_ids")

    def __init__(self, data):
        self.data = data
        self.like_ids = data.get("like_ids")

    def __repr__(self):
        return f"LikePostsResponse(data={self.data})"


class ValidationPostResponse:

    __slots__ = ("data", "is_allow_to_post")

    def __init__(self, data):
        self.data = data
        self.is_allow_to_post = data.get("is_allow_to_post")

    def __repr__(self):
        return f"ValidationPostResponse(data={self.data})"


class RefreshCounterRequestsResponse:

    __slots__ = ("data", "reset_counter_requests")

    def __init__(self, data):
        self.data = data

        self.reset_counter_requests = data.get("reset_counter_requests")
        if self.reset_counter_requests is not None:
            self.reset_counter_requests = [
                RefreshCounterRequest(reset_counter_request) for reset_counter_request in self.reset_counter_requests
            ]

    def __repr__(self):
        return f"RefreshCounterRequestsResponse(data={self.data})"


class SocialShareUsersResponse:

    __slots__ = ("data", "social_shared_users")

    def __init__(self, data):
        self.data = data

        self.social_shared_users = data.get("social_shared_users")
        if self.social_shared_users is not None:
            self.social_shared_users = [
                UserWrapper(social_shared_user) for social_shared_user in self.social_shared_users
            ]

    def __repr__(self):
        return f"SocialShareUsersResponse(data={self.data})"


class VoteSurveyResponse:

    __slots__ = ("data", "survey")

    def __init__(self, data):
        self.data = data

        self.survey = data.get("survey")
        if self.survey is not None:
            self.survey = Survey(self.survey)

    def __repr__(self):
        return f"VoteSurveyResponse(data={self.data})"


class UserResponse:

    __slots__ = (
        "data", "user", "masked_email", "twitter_id", "is_line_connected",
        "is_facebook_connected", "is_lobi_connected", "is_push_notification_on",
        "is_call_on", "is_video_on", "is_group_call_on", "is_group_video_on", "vip_until",
        "is_email_confirmed", "blocking_limit", "uuid", "birthdate", "gifting_ability"
    )

    def __init__(self, data):
        self.data = data

        self.user = data.get("user")
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

        self.gifting_ability = data.get("gifting_ability")
        if self.gifting_ability is not None:
            self.gifting_ability = GiftingAbility(self.gifting_ability)

    def __repr__(self):
        return f"UserResponse(data={self.data})"


class UsersResponse:

    __slots__ = ("data", "users", "next_page_value")

    def __init__(self, data):
        self.data = data

        self.users = data.get("users")
        if self.users is not None:
            self.users = [
                User(user) for user in self.users
            ]

        self.next_page_value = data.get("next_page_value")

    def __repr__(self):
        return f"UsersResponse(data={self.data})"


class UserCustomDefinitionsResponse:

    __slots__ = (
        "age", "followers_count", "followings_count", "created_at",
        "last_loggedin_at", "status", "reported_count"
    )

    def __init__(self, data):
        self.data = data
        self.age = data.get("age")
        self.followers_count = data.get("followers_count")
        self.followings_count = data.get("followings_count")
        self.created_at = data.get("created_at")
        self.last_loggedin_at = data.get("last_loggedin_at")
        self.status = data.get("status")
        self.reported_count = data.get("reported_count")

    def __repr__(self):
        return f"UserCustomDefinitionsResponse(data={self.data})"


class UserEmailResponse:

    __slots__ = ("email")

    def __init__(self, data):
        self.data = data
        self.email = data.get("email")

    def __repr__(self):
        return f"UserEmailResponse(data={self.data})"


class HimaUsersResponse:

    __slots__ = ("data", "hima_users")

    def __init__(self, data):
        self.data = data

        self.hima_users = data.get("hima_users")
        if self.hima_users is not None:
            self.hima_users = [
                UserWrapper(hima_user) for hima_user in self.hima_users
            ]

    def __repr__(self):
        return f"HimaUsersResponse(data={self.data})"


class UsersByTimestampResponse:

    __slots__ = ("data", "last_timestamp", "users")

    def __init__(self, data):
        self.data = data
        self.last_timestamp = data.get("last_timestamp")

        self.users = data.get("users")
        if self.users is not None:
            self.users = [
                User(user) for user in self.users
            ]

    def __repr__(self):
        return f"UsersByTimestampResponse(data={self.data})"


class UserTimestampResponse:

    __slots__ = ("data", "time", "ip_address", "country")

    def __init__(self, data):
        self.data = data
        self.id = data.get("time")
        self.ip_address = data.get("ip_address")
        self.country = data.get("country")

    def __repr__(self):
        return f"UserTimestampResponse(data={self.data})"


class WebSocketTokenResponse:

    __slots__ = ("data", "token")

    def __init__(self, data):
        self.data = data
        self.token = data.get("token")

    def __repr__(self):
        return f"WebSocketTokenResponse(data={self.data})"
